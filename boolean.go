package vjson

import "github.com/pkg/errors"

type BooleanField struct {
	name            string
	required        bool
	valueValidation bool
	value           bool
}

func (b *BooleanField) Validate(v interface{}) error {
	if v == nil {
		if !b.required {
			return nil
		} else {
			return errors.Errorf("Value for %s field is required", b.name)
		}
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

func (b *BooleanField) Required() *BooleanField {
	b.required = true
	return b
}

func (b *BooleanField) ShouldBe(value bool) *BooleanField {
	b.value = value
	b.valueValidation = true
	return b
}

func Boolean(name string) *BooleanField {
	return &BooleanField{
		name:     name,
		required: false,
	}
}
