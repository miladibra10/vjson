package vjson

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ObjectField is the type for validating another JSON object in a JSON
type ObjectField struct {
	name     string
	required bool
	schema   Schema
}

// To Force Implementing Field interface by ObjectField
var _ Field = (*ObjectField)(nil)

// GetName returns name of the field
func (o *ObjectField) GetName() string {
	return o.name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (o *ObjectField) Validate(v interface{}) error {
	if v == nil {
		if !o.required {
			return nil
		}
		return errors.Errorf("Value for %s field is required", o.name)
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

// Required is called to make a field required in a JSON
func (o *ObjectField) Required() *ObjectField {
	o.required = true
	return o
}

func (o *ObjectField) Strict() *ObjectField {
	o.schema.StrictMode = true
	return o
}

func (o *ObjectField) MarshalJSON() ([]byte, error) {
	schemaRaw, err := json.Marshal(o.schema)
	if err != nil {
		return nil, errors.Wrapf(err, "could not marshal schema field of object field: %s", o.name)
	}

	schema := make(map[string]interface{})
	err = json.Unmarshal(schemaRaw, &schema)
	if err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal schema field of array field: %s", o.name)
	}

	return json.Marshal(ObjectFieldSpec{
		Name:     o.name,
		Type:     objectType,
		Required: o.required,
		Schema:   schema,
	})
}

// Object is the constructor of an object field
func Object(name string, schema Schema) *ObjectField {
	return &ObjectField{
		name:     name,
		required: false,
		schema:   schema,
	}
}
