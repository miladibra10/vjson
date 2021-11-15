package vjson

import "github.com/pkg/errors"

// NullField is the type for validating floats in a JSON
type NullField struct {
	Name string `json:"name"`
}

// To Force Implementing Field interface by NullField
var _ Field = (*NullField)(nil)

// GetName returns name of the field
func (n *NullField) GetName() string {
	return n.Name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (n *NullField) Validate(input interface{}) error {
	if input == nil {
		return nil
	}
	return errors.Errorf("Value for %s should be null", n.Name)
}

// Null is the constructor of a null field in a JSON.
func Null(name string) *NullField {
	return &NullField{
		Name: name,
	}
}
