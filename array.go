package vjson

import (
	"encoding/json"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// ArrayField is the type for validating arrays in a JSON
type ArrayField struct {
	name     string
	required bool
	items    Field
	fixItems []Field

	minLength           int
	minLengthValidation bool

	maxLength           int
	maxLengthValidation bool
}

// To Force Implementing Field interface by ArrayField
var _ Field = (*ArrayField)(nil)

// GetName returns name of the field
func (a *ArrayField) GetName() string {
	return a.name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (a *ArrayField) Validate(v interface{}) error {
	if v == nil {
		if !a.required {
			return nil
		}
		return errors.Errorf("Value for %s field is required", a.name)
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

	if a.items != nil && a.fixItems != nil {
		result = multierror.Append(result, errors.Errorf("could not using both items key and fix items for array %s", a.name))
	}

	if a.items != nil {
		for _, value := range values {
			err := a.items.Validate(value)
			if err != nil {
				result = multierror.Append(result, errors.Wrapf(err, "%v item is invalid in %s array", value, a.name))
			}
		}
	}
	if a.fixItems != nil {
		if len(a.fixItems) != len(values) {
			result = multierror.Append(result, errors.Errorf("length of %s array should is equal %d", a.name, len(a.fixItems)))
			return result
		}
		for i, value := range values {
			err := a.fixItems[i].Validate(value)
			if err != nil {
				result = multierror.Append(result, errors.Wrapf(err, "%v item is invalid in %s array", value, a.name))
			}
		}
	}
	return result
}

// Required is called to make a field required in a JSON
func (a *ArrayField) Required() *ArrayField {
	a.required = true
	return a
}

// MinLength is called to set minimum length for an array field in a JSON
func (a *ArrayField) MinLength(length int) *ArrayField {
	a.minLength = length
	a.minLengthValidation = true
	return a
}

// MaxLength is called to set maximum length for an array field in a JSON
func (a *ArrayField) MaxLength(length int) *ArrayField {
	a.maxLength = length
	a.maxLengthValidation = true
	return a
}

func (a *ArrayField) MarshalJSON() ([]byte, error) {
	var items map[string]interface{}
	if a.items != nil {
		itemsRaw, err := json.Marshal(a.items)
		if err != nil {
			return nil, errors.Wrapf(err, "could not marshal items field of array field: %s", a.name)
		}

		items = make(map[string]interface{})
		err = json.Unmarshal(itemsRaw, &items)
		if err != nil {
			return nil, errors.Wrapf(err, "could not unmarshal items field of array field: %s", a.name)
		}
	}

	var fixItems []map[string]interface{}
	if a.fixItems != nil {
		itemsRaw, err := json.Marshal(a.fixItems)
		if err != nil {
			return nil, errors.Wrapf(err, "could not marshal fix items field of array field: %s", a.name)
		}

		fixItems = []map[string]interface{}{}
		err = json.Unmarshal(itemsRaw, &fixItems)
		if err != nil {
			return nil, errors.Wrapf(err, "could not unmarshal fix items field of array field: %s", a.name)
		}
	}

	return json.Marshal(ArrayFieldSpec{
		Name:      a.name,
		Type:      arrayType,
		Required:  a.required,
		Items:     items,
		FixItems:  fixItems,
		MinLength: a.minLength,
		MaxLength: a.maxLength,
	})
}

// Array is the constructor of an array field.
func Array(name string, itemField Field) *ArrayField {
	return &ArrayField{
		name:     name,
		required: false,
		items:    itemField,
	}
}

// Array is the constructor of an array field.
func FixArray(name string, itemFields []Field) *ArrayField {
	return &ArrayField{
		name:     name,
		required: false,
		fixItems: itemFields,
	}
}
