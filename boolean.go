package vjson

import "github.com/pkg/errors"

// BooleanField is the type for validating booleans in a JSON
type BooleanField struct {
	name            string
	required        bool
	valueValidation bool
	value           bool
}

// To Force Implementing Field interface by BooleanField
var _ Field = (*BooleanField)(nil)

// GetName returns name of the field
func (b *BooleanField) GetName() string {
	return b.name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (b *BooleanField) Validate(v interface{}) error {
	if v == nil {
		if !b.required {
			return nil
		}
		return errors.Errorf("Value for %s field is required", b.name)
	}

	value, ok := v.(bool)

	if !ok {
		return errors.Errorf("Value for %s should be a boolean", b.name)
	}

	if b.valueValidation {
		if value != b.value {
			return errors.Errorf("Value for %s should be a %v", b.name, b.value)
		}
	}

	return nil
}

// Required is called to make a field required in a JSON
func (b *BooleanField) Required() *BooleanField {
	b.required = true
	return b
}

// ShouldBe is called for setting a value for checking a boolean.
func (b *BooleanField) ShouldBe(value bool) *BooleanField {
	b.value = value
	b.valueValidation = true
	return b
}

// Boolean is the constructor of a boolean field
func Boolean(name string) *BooleanField {
	return &BooleanField{
		name:     name,
		required: false,
	}
}
