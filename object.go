package vjson

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type ObjectField struct {
	name     string
	required bool
	schema   Schema
}

// To Force Implementing Field interface by ObjectField
var _ Field = (*ObjectField)(nil)

func (o *ObjectField) GetName() string {
	return o.name
}

func (o *ObjectField) Validate(v interface{}) error {
	if v == nil {
		if !o.required {
			return nil
		} else {
			return errors.Errorf("Value for %s field is required", o.name)
		}
	}

	// The input is either string or an interface{} object
	value, ok := v.(string)

	var err error
	var jsonBytes []byte
	if !ok {
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			return errors.Errorf("Value for %s should be an object", o.name)
		}
	} else {
		return o.schema.ValidateString(value)
	}

	return o.schema.ValidateBytes(jsonBytes)
}

func (o *ObjectField) Required() *ObjectField {
	o.required = true
	return o
}

func Object(name string, schema Schema) *ObjectField {
	return &ObjectField{
		name:     name,
		required: false,
		schema:   schema,
	}
}
