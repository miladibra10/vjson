package vjson

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ObjectField is the type for validating another JSON object in a JSON
type ObjectField struct {
	Name          string `json:"name"`
	FieldRequired bool   `json:"required"`
	FieldSchema   Schema `json:"schema"`
}

// To Force Implementing Field interface by ObjectField
var _ Field = (*ObjectField)(nil)

// GetName returns name of the field
func (o *ObjectField) GetName() string {
	return o.Name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (o *ObjectField) Validate(v interface{}) error {
	if v == nil {
		if !o.FieldRequired {
			return nil
		}
		return errors.Errorf("Value for %s field is required", o.Name)
	}

	// The input is either string or an interface{} object
	value, ok := v.(string)

	var err error
	var jsonBytes []byte
	if !ok {
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			return errors.Errorf("Value for %s should be an object", o.Name)
		}
	} else {
		return o.FieldSchema.ValidateString(value)
	}

	return o.FieldSchema.ValidateBytes(jsonBytes)
}

// Required is called to make a field required in a JSON
func (o *ObjectField) Required() *ObjectField {
	o.FieldRequired = true
	return o
}

// Object is the constructor of an object field
func Object(name string, schema Schema) *ObjectField {
	return &ObjectField{
		Name:          name,
		FieldRequired: false,
		FieldSchema:   schema,
	}
}
