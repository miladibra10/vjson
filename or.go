package vjson

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// OrField is the type for validating if a value matches at least one of the provided field types
type OrField struct {
	name     string
	required bool
	fields   []Field
}

// To Force Implementing Field interface by OrField
var _ Field = (*OrField)(nil)

// GetName returns name of the field
func (o *OrField) GetName() string {
	return o.name
}

// Validate is used for validating a value. It returns an error if the value is invalid for all field types.
func (o *OrField) Validate(v interface{}) error {
	if v == nil {
		if !o.required {
			return nil
		}
		return errors.Errorf("Value for %s field is required", o.name)
	}

	if len(o.fields) == 0 {
		return errors.Errorf("No field types provided for %s", o.name)
	}

	for _, field := range o.fields {
		err := field.Validate(v)
		if err == nil {
			// If any field validation passes, the Or validation passes
			return nil
		}
	}

	// If we get here, all validations failed
	return errors.Errorf("Value for %s does not match any of the allowed types", o.name)
}

// Required is called to make a field required in a JSON
func (o *OrField) Required() *OrField {
	o.required = true
	return o
}

// WithFields is called to set the field types to check
func (o *OrField) WithFields(fields ...Field) *OrField {
	o.fields = fields
	return o
}

func (o *OrField) MarshalJSON() ([]byte, error) {
	// Convert []Field to []map[string]interface{}
	fieldsJSON := make([]map[string]interface{}, 0, len(o.fields))
	for _, field := range o.fields {
		fieldBytes, err := json.Marshal(field)
		if err != nil {
			return nil, errors.Wrapf(err, "could not marshal field of or field: %s", o.name)
		}

		fieldMap := make(map[string]interface{})
		err = json.Unmarshal(fieldBytes, &fieldMap)
		if err != nil {
			return nil, errors.Wrapf(err, "could not unmarshal field of or field: %s", o.name)
		}

		fieldsJSON = append(fieldsJSON, fieldMap)
	}

	return json.Marshal(OrFieldSpec{
		Name:     o.name,
		Type:     "or",
		Required: o.required,
		Fields:   fieldsJSON,
	})
}

// Or is the constructor of an Or field
func Or(name string, fields ...Field) *OrField {
	return &OrField{
		name:     name,
		required: false,
		fields:   fields,
	}
}
