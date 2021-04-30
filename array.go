package vjson

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type ArrayField struct {
	name     string
	required bool
	items    Field

	minLength           int
	minLengthValidation bool

	maxLength           int
	maxLengthValidation bool
}

// To Force Implementing Field interface by ArrayField
var _ Field = (*ArrayField)(nil)

func (a *ArrayField) Validate(v interface{}) error {
	if v == nil {
		if !a.required {
			return nil
		} else {
			return errors.Errorf("Value for %s field is required", a.name)
		}
	}

	values, ok := v.([]interface{})
	if !ok {
		return errors.Errorf("Value of %s should be array", a.name)
	}

	var result error
	if a.minLengthValidation {
		if len(values) < a.minLength {
			result = multierror.Append(result, errors.Errorf("length of %s array should be at least %d", a.name, a.minLength))
		}
	}

	if a.maxLengthValidation {
		if len(values) > a.maxLength {
			result = multierror.Append(result, errors.Errorf("length of %s array should be at most %d", a.name, a.maxLength))
		}
	}

	for _, value := range values {
		err := a.items.Validate(value)
		if err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "%v item is invalid in %s array", value, a.name))
		}
	}
	return result
}

func (a *ArrayField) Required() *ArrayField {
	a.required = true
	return a
}

func (a *ArrayField) MinLength(length int) *ArrayField {
	a.minLength = length
	a.minLengthValidation = true
	return a
}

func (a *ArrayField) MaxLength(length int) *ArrayField {
	a.maxLength = length
	a.maxLengthValidation = true
	return a
}

func Array(name string, itemField Field) *ArrayField {
	return &ArrayField{
		name:     name,
		required: false,
		items:    itemField,
	}
}
