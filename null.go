package vjson

import "github.com/pkg/errors"

type NullField struct {
	name string
}

// To Force Implementing Field interface by NullField
var _ Field = (*NullField)(nil)

func (n *NullField) GetName() string {
	return n.name
}

func (n *NullField) Validate(input interface{}) error {
	if input == nil {
		return nil
	}
	return errors.Errorf("Value for %s should be null", n.name)
}

func Null(name string) *NullField {
	return &NullField{
		name: name,
	}
}
