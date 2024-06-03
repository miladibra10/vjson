package vjson

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

// Schema is the type for declaring a JSON schema and validating a json object.
type Schema struct {
	Fields     []Field `json:"fields"`
	StrictMode bool    `json:"strict"`
}

// SchemaSpec is used for parsing a Schema
type SchemaSpec struct {
	Fields     []map[string]interface{} `json:"fields"`
	StrictMode bool                     `json:"strict"`
}

// UnmarshalJSON is implemented for parsing a Schema. it overrides json.Unmarshal behaviour.
func (s *Schema) UnmarshalJSON(bytes []byte) error {
	var schemaSpec SchemaSpec
	err := json.Unmarshal(bytes, &schemaSpec)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal to SchemaSpec")
	}
	s.Fields = make([]Field, 0, len(schemaSpec.Fields))
	s.StrictMode = schemaSpec.StrictMode

	var result error

	for _, fieldSpec := range schemaSpec.Fields {
		field, err := s.getField(fieldSpec)
		if err != nil {
			result = multierror.Append(result, err)
			continue
		}
		s.Fields = append(s.Fields, field)
	}

	return result
}

func (s *Schema) getField(fieldSpec map[string]interface{}) (Field, error) {
	fieldTypeRaw, found := fieldSpec[typeKey]
	if found {
		fieldTypeStr, ok := fieldTypeRaw.(string)
		if ok {
			fieldType := fieldType(fieldTypeStr)
			switch fieldType {
			case integerType:
				{
					field, err := s.getIntegerField(fieldSpec)
					if err != nil {
						return nil, err
					}
					return field, nil
				}
			case floatType:
				{
					field, err := s.getFloatField(fieldSpec)
					if err != nil {
						return nil, err
					}
					return field, nil
				}
			case stringType:
				{
					field, err := s.getStringField(fieldSpec)
					if err != nil {
						return nil, err
					}
					return field, nil
				}
			case arrayType:
				{
					field, err := s.getArrayField(fieldSpec)
					if err != nil {
						return nil, err
					}
					return field, nil
				}
			case booleanType:
				{
					field, err := s.getBooleanField(fieldSpec)
					if err != nil {
						return nil, err
					}
					return field, nil
				}
			case objectType:
				{
					field, err := s.getObjectField(fieldSpec)
					if err != nil {
						return nil, err
					}
					return field, nil
				}
			case nullType:
				{
					field, err := s.getNullField(fieldSpec)
					if err != nil {
						return nil, err
					}
					return field, nil
				}
			default:
				{
					return nil, errors.Errorf("Invalid type: %s", fieldType)
				}
			}
		}
		return nil, errors.Errorf("invalid field type")
	}
	return nil, errors.Errorf("field type not found")
}

func (s *Schema) getIntegerField(fieldSpec map[string]interface{}) (*IntegerField, error) {
	var intSpec IntegerFieldSpec
	err := mapstructure.Decode(fieldSpec, &intSpec)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode integer field to IntegerFieldSpec")
	}
	if intSpec.Name == "" {
		return nil, errors.Errorf("name field is required for an integer field")
	}
	_, minValidation := fieldSpec["min"]
	_, maxValidation := fieldSpec["max"]
	_, signValidation := fieldSpec["positive"]
	_, rangeValidation := fieldSpec["ranges"]

	intField := NewInteger(intSpec, minValidation, maxValidation, signValidation, rangeValidation)

	return intField, nil
}

func (s *Schema) getFloatField(fieldSpec map[string]interface{}) (*FloatField, error) {
	var floatSpec FloatFieldSpec
	err := mapstructure.Decode(fieldSpec, &floatSpec)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode float field to IntegerFieldSpec")
	}

	if floatSpec.Name == "" {
		return nil, errors.Errorf("name field is required for a float field")
	}

	_, minValidation := fieldSpec["min"]
	_, maxValidation := fieldSpec["max"]
	_, signValidation := fieldSpec["positive"]
	_, rangeValidation := fieldSpec["ranges"]

	floatField := NewFloat(floatSpec, minValidation, maxValidation, signValidation, rangeValidation)

	return floatField, nil
}

func (s *Schema) getStringField(fieldSpec map[string]interface{}) (*StringField, error) {
	var stringSpec StringFieldSpec
	err := mapstructure.Decode(fieldSpec, &stringSpec)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode float field to IntegerFieldSpec")
	}
	if stringSpec.Name == "" {
		return nil, errors.Errorf("name field is required for a string field")
	}
	_, minLenValidation := fieldSpec["min_length"]
	_, maxLenValidation := fieldSpec["max_length"]
	_, formatValidation := fieldSpec["format"]
	_, choiceValidation := fieldSpec["choices"]

	stringField := NewString(stringSpec, minLenValidation, maxLenValidation, formatValidation, choiceValidation)

	return stringField, nil
}

func (s *Schema) getArrayField(fieldSpec map[string]interface{}) (*ArrayField, error) {
	var arraySpec ArrayFieldSpec
	err := mapstructure.Decode(fieldSpec, &arraySpec)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode array field to ArrayFieldSpec")
	}
	if arraySpec.Name == "" {
		return nil, errors.Errorf("name field is required for an array field")
	}

	var itemField Field
	var fixItemFields []Field

	itemsFieldSpecRaw, found := fieldSpec["items"]
	fixItemsFieldSpecRaw, foundFixItem := fieldSpec["fix_items"]
	if !found && !foundFixItem {
		return nil, errors.Errorf("items key is missing for array field name: %s", arraySpec.Name)
	}
	if found && foundFixItem {
		return nil, errors.Errorf("could not both using items key and fix items for array field name: %s", arraySpec.Name)
	}

	if itemsFieldSpec, ok := itemsFieldSpecRaw.(map[string]interface{}); ok {
		itemField, err = s.getField(itemsFieldSpec)
		if err != nil {
			return nil, errors.Wrapf(err, "could not get item field of array field name: %s", arraySpec.Name)
		}
	} else if fixItemsFieldSpec, ok := fixItemsFieldSpecRaw.([]interface{}); ok {
		fixItemFields = []Field{}
		for _, fixFieldSpec := range fixItemsFieldSpec {
			fixItemField, err := s.getField(fixFieldSpec.(map[string]interface{}))
			if err != nil {
				return nil, errors.Wrapf(err, "could not get item field of array field name: %s", arraySpec.Name)
			}
			fixItemFields = append(fixItemFields, fixItemField)
		}
	} else {
		return nil, errors.Errorf("invalid format for items key for array field name: %s", arraySpec.Name)
	}

	_, minLenValidation := fieldSpec["min_length"]
	_, maxLenValidation := fieldSpec["max_length"]

	arrayField := NewArray(arraySpec, itemField, fixItemFields, minLenValidation, maxLenValidation)
	return arrayField, nil
}

func (s *Schema) getBooleanField(fieldSpec map[string]interface{}) (*BooleanField, error) {
	var booleanSpec BooleanFieldSpec
	err := mapstructure.Decode(fieldSpec, &booleanSpec)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode boolean field to BooleanFieldSpec")
	}
	if booleanSpec.Name == "" {
		return nil, errors.Errorf("name field is required for a boolean field")
	}

	_, valueValidation := fieldSpec["value"]

	booleanField := NewBoolean(booleanSpec, valueValidation)

	return booleanField, nil
}

func (s *Schema) getObjectField(fieldSpec map[string]interface{}) (*ObjectField, error) {
	var objectSpec ObjectFieldSpec
	err := mapstructure.Decode(fieldSpec, &objectSpec)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode object field to ObjectFieldSpec")
	}
	if objectSpec.Name == "" {
		return nil, errors.Errorf("name field is required for an object field")
	}

	schemaSpecRaw, found := fieldSpec["schema"]
	if !found {
		return nil, errors.Errorf("schema key is missing for object field name: %s", objectSpec.Name)
	}
	schemaSpec, ok := schemaSpecRaw.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("invalid format for schema key for object field name: %s", objectSpec.Name)
	}

	jsonSchemaSpec, err := json.Marshal(schemaSpec)
	if err != nil {
		return nil, errors.Errorf("could not marshal schema to json for object field name: %s", objectSpec.Name)
	}

	var schema Schema
	err = json.Unmarshal(jsonSchemaSpec, &schema)
	if err != nil {
		return nil, errors.Errorf("could not unmarshal schema spec to schema for object field name: %s", objectSpec.Name)
	}

	objectField := NewObject(objectSpec, schema)
	return objectField, nil
}

func (s *Schema) getNullField(fieldSpec map[string]interface{}) (*NullField, error) {
	var nullSpec NullFieldSpec
	err := mapstructure.Decode(fieldSpec, &nullSpec)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode null field to NullFieldSpec")
	}
	if nullSpec.Name == "" {
		return nil, errors.Errorf("name field is required for a null field")
	}

	nullField := NewNull(nullSpec)

	return nullField, nil
}

// ValidateBytes receives a byte array of a json object and validates it according to the specified Schema.
// it returns an error if the input is invalid.
func (s *Schema) ValidateBytes(input []byte) error {
	if gjson.ValidBytes(input) {
		jsonObject := gjson.ParseBytes(input)
		return s.validateJSON(jsonObject)
	}
	return errors.Errorf("could not parse json input.")
}

// ValidateString is like ValidateBytes but it receives the json object as string input.
func (s *Schema) ValidateString(input string) error {
	if gjson.Valid(input) {
		jsonObject := gjson.Parse(input)
		return s.validateJSON(jsonObject)
	}
	return errors.Errorf("could not parse json input.")
}

func (s *Schema) validateJSON(json gjson.Result) error {
	var result error
	if json.IsObject() {
		if s.StrictMode {
			jsonMap := json.Map()
			for jsonField := range jsonMap {
				fieldFound := false
				for _, field := range s.Fields {
					if jsonField == field.GetName() {
						fieldFound = true
					}
				}
				if !fieldFound {
					result = multierror.Append(result, errors.Errorf("Field %s is not validata", jsonField))
				}
			}
			if result != nil {
				return result
			}
		}
	}

	for _, field := range s.Fields {
		fieldName := field.GetName()
		fieldValue := json.Get(fieldName).Value()
		err := field.Validate(fieldValue)
		if err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "Field %s is invalid.", fieldName))
		}
	}
	return result
}

// NewSchema is the constructor for Schema. it receives a list of Field in its arguments.
func NewSchema(fields ...Field) Schema {
	return Schema{Fields: fields}
}

// ReadFromString is for parsing a Schema from a string input.
func ReadFromString(input string) (*Schema, error) {
	return ReadFromBytes([]byte(input))
}

// ReadFromBytes is for parsing a Schema from a byte array input.
func ReadFromBytes(input []byte) (*Schema, error) {
	var s Schema
	err := json.Unmarshal(input, &s)
	if err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal file given to Schema")
	}
	return &s, nil
}

// ReadFromFile is for parsing a Schema from a file input.
func ReadFromFile(filePath string) (*Schema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not open file given, path: %s", filePath)
	}
	input, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read file given, path: %s", filePath)
	}
	return ReadFromBytes(input)
}
