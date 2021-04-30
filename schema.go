package vjson

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type Schema struct {
	Fields []Field `json:"fields"`
}

func (s *Schema) ValidateBytes(input []byte) error {
	if gjson.ValidBytes(input) {
		json := gjson.ParseBytes(input)
		return s.validateJSON(json)
	}
	return errors.Errorf("could not parse json input.")
}

func (s *Schema) ValidateString(input string) error {
	if gjson.Valid(input) {
		json := gjson.Parse(input)
		return s.validateJSON(json)
	}
	return errors.Errorf("could not parse json input.")
}

func (s *Schema) validateJSON(json gjson.Result) error {
	var result error
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

func ReadFromString(input string) (*Schema, error) {
	return nil, nil
}

func ReadFromBytes(input []byte) (*Schema, error) {
	return nil, nil
}

func ReadFromFile(filePath string) (*Schema, error) {
	return nil, nil
}
