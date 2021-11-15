package vjson

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// ArrayField is the type for validating arrays in a JSON
type ArrayField struct {
	Name         string `json:"name"`
	FieldRequred bool   `json:"required"`
	Items        Field  `json:"items"`

	FieldMinLength           int  `json:"minLength"`
	FieldMinLengthValidation bool `json:"minLengthValidation"`

	FieldMaxLength           int  `json:"maxLength"`
	FieldMaxLengthValidation bool `json:"maxLengthValidation"`
}

// To Force Implementing Field interface by ArrayField
var _ Field = (*ArrayField)(nil)

// GetName returns name of the field
func (a *ArrayField) GetName() string {
	return a.Name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (a *ArrayField) Validate(v interface{}) error {
	if v == nil {
		if !a.FieldRequred {
			return nil
		}
		return errors.Errorf("Value for %s field is required", a.Name)
	}

	values, ok := v.([]interface{})
	if !ok {
		return errors.Errorf("Value of %s should be array", a.Name)
	}

	var result error
	if a.FieldMinLengthValidation {
		if len(values) < a.FieldMinLength {
			result = multierror.Append(result, errors.Errorf("length of %s array should be at least %d", a.Name, a.FieldMinLength))
		}
	}

	if a.FieldMaxLengthValidation {
		if len(values) > a.FieldMaxLength {
			result = multierror.Append(result, errors.Errorf("length of %s array should be at most %d", a.Name, a.FieldMaxLength))
		}
	}

	for _, value := range values {
		err := a.Items.Validate(value)
		if err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "%v item is invalid in %s array", value, a.Name))
		}
	}
	return result
}

// Required is called to make a field required in a JSON
func (a *ArrayField) Required() *ArrayField {
	a.FieldRequred = true
	return a
}

// MinLength is called to set minimum length for an array field in a JSON
func (a *ArrayField) MinLength(length int) *ArrayField {
	a.FieldMinLength = length
	a.FieldMinLengthValidation = true
	return a
}

// MaxLength is called to set maximum length for an array field in a JSON
func (a *ArrayField) MaxLength(length int) *ArrayField {
	a.FieldMaxLength = length
	a.FieldMaxLengthValidation = true
	return a
}

// Array is the constructor of an array field.
func Array(name string, itemField Field) *ArrayField {
	return &ArrayField{
		Name:         name,
		FieldRequred: false,
		Items:        itemField,
	}
}
