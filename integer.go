package vjson

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type intRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// IntegerField is the type for validating integers in a JSON
type IntegerField struct {
	Name          string `json:"name"`
	FieldRequired bool   `json:"required"`

	FieldMin           int  `json:"min"`
	FieldMinValidation bool `json:"minValidation"`

	FieldMax           int  `json:"max"`
	FieldMaxValidation bool `json:"maxValidation"`

	FieldSignValidation bool `json:"signValidation"`
	FieldPositive       bool `json:"positive"`

	FieldRangeValidation bool       `json:"rangeValidation"`
	FieldRanges          []intRange `json:"ranges"`
}

// To Force Implementing Field interface by IntegerField
var _ Field = (*IntegerField)(nil)

// GetName returns name of the field
func (i *IntegerField) GetName() string {
	return i.Name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (i *IntegerField) Validate(v interface{}) error {
	if v == nil {
		if !i.FieldRequired {
			return nil
		}
		return errors.Errorf("Value for %s field is required", i.Name)
	}
	var value int
	var intOK bool

	//gjson library returns float64 when field value is number
	floatValue, floatOK := v.(float64)

	value, intOK = v.(int)

	if !floatOK && !intOK {
		return errors.Errorf("Value for %s should be a number", i.Name)
	}

	if floatOK {
		value = int(floatValue)
	}

	var result error
	if i.FieldSignValidation && i.FieldPositive {
		if value < 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a positive integer", i.Name))
		}
	} else if i.FieldSignValidation && !i.FieldPositive {
		if value > 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a negative integer", i.Name))
		}
	}

	if i.FieldMinValidation {
		if value < i.FieldMin {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at least %d", i.Name, i.FieldMin))
		}
	}

	if i.FieldMaxValidation {
		if value > i.FieldMax {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at most %d", i.Name, i.FieldMax))
		}
	}

	if i.FieldRangeValidation {
		inRange := false
		for _, r := range i.FieldRanges {
			if value >= r.Start && value <= r.End {
				inRange = true
				break
			}
		}

		if !inRange {
			var ranges strings.Builder
			for _, r := range i.FieldRanges {
				ranges.WriteString(fmt.Sprintf("[%d,%d] ", r.Start, r.End))
			}
			result = multierror.Append(result, errors.Errorf("Value for %s should be in one of these ranges: %s", i.Name, ranges.String()))
		}
	}

	return result
}

// Required is called to make a field required in a JSON
func (i *IntegerField) Required() *IntegerField {
	i.FieldRequired = true
	return i
}

// Positive is called when we want to force the value to be positive in validation.
func (i *IntegerField) Positive() *IntegerField {
	i.FieldSignValidation = true
	i.FieldPositive = true
	return i
}

// Negative is called when we want to force the value to be negative in validation.
func (i *IntegerField) Negative() *IntegerField {
	i.FieldSignValidation = true
	i.FieldPositive = false
	return i
}

// Min is called when we want to set a minimum value for an integer value in validation.
func (i *IntegerField) Min(value int) *IntegerField {
	i.FieldMin = value
	i.FieldMinValidation = true
	return i
}

// Max is called when we want to set a maximum value for an integer value in validation.
func (i *IntegerField) Max(value int) *IntegerField {
	i.FieldMax = value
	i.FieldMaxValidation = true
	return i
}

// Range is called when we want to define valid ranges for an integer value in validation.
func (i *IntegerField) Range(start, end int) *IntegerField {
	i.FieldRanges = append(i.FieldRanges, intRange{Start: start, End: end})
	i.FieldRangeValidation = true
	return i
}

// Integer is the constructor of an integer field
func Integer(name string) *IntegerField {
	return &IntegerField{
		Name:                 name,
		FieldRequired:        false,
		FieldMin:             0,
		FieldMinValidation:   false,
		FieldMax:             0,
		FieldMaxValidation:   false,
		FieldSignValidation:  false,
		FieldPositive:        false,
		FieldRangeValidation: false,
		FieldRanges:          []intRange{},
	}
}
