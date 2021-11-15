package vjson

import "github.com/pkg/errors"

// BooleanField is the type for validating booleans in a JSON
type BooleanField struct {
	Name                 string `json:"name"`
	FieldRequired        bool   `json:"required"`
	FieldValueValidation bool   `json:"valueValidation"`
	Value                bool   `json:"value"`
}

// To Force Implementing Field interface by BooleanField
var _ Field = (*BooleanField)(nil)

// GetName returns name of the field
func (b *BooleanField) GetName() string {
	return b.Name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (b *BooleanField) Validate(v interface{}) error {
	if v == nil {
		if !b.FieldRequired {
			return nil
		}
		return errors.Errorf("Value for %s field is required", b.Name)
	}

	value, ok := v.(bool)

	if !ok {
		return errors.Errorf("Value for %s should be a boolean", b.Name)
	}

	if b.FieldValueValidation {
		if value != b.Value {
			return errors.Errorf("Value for %s should be a %v", b.Name, b.Value)
		}
	}

	return nil
}

// Required is called to make a field required in a JSON
func (b *BooleanField) Required() *BooleanField {
	b.FieldRequired = true
	return b
}

// ShouldBe is called for setting a value for checking a boolean.
func (b *BooleanField) ShouldBe(value bool) *BooleanField {
	b.Value = value
	b.FieldValueValidation = true
	return b
}

// Boolean is the constructor of a boolean field
func Boolean(name string) *BooleanField {
	return &BooleanField{
		Name:          name,
		FieldRequired: false,
	}
}
