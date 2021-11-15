package vjson

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type floatRange struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

// FloatField is the type for validating floats in a JSON
type FloatField struct {
	Name          string `json:"name"`
	FieldRequired bool   `json:"required"`

	FieldMin           float64 `json:"min"`
	FieldMinValidation bool    `json:"minValidation"`

	FieldMax           float64 `json:"max"`
	FieldMaxValidation bool    `json:"maxValidation"`

	FieldSignValidation bool `json:"signValidation"`
	FieldPositive       bool `json:"positive"`

	FieldRangeValidation bool         `json:"rangeValidation"`
	FieldRanges          []floatRange `json:"ranges"`
}

// To Force Implementing Field interface by IntegerField
var _ Field = (*FloatField)(nil)

// GetName returns name of the field
func (f *FloatField) GetName() string {
	return f.Name
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (f *FloatField) Validate(v interface{}) error {
	if v == nil {
		if !f.FieldRequired {
			return nil
		}
		return errors.Errorf("Value for %s field is required", f.Name)
	}

	value, ok := v.(float64)

	if !ok {
		return errors.Errorf("Value for %s should be a float number", f.Name)
	}

	var result error
	if f.FieldSignValidation && f.FieldPositive {
		if value < 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a positive float", f.Name))
		}
	} else if f.FieldSignValidation && !f.FieldPositive {
		if value > 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a negative float", f.Name))
		}
	}

	if f.FieldMinValidation {
		if value < f.FieldMin {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at least %f", f.Name, f.FieldMin))
		}
	}

	if f.FieldMaxValidation {
		if value > f.FieldMax {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at most %f", f.Name, f.FieldMax))
		}
	}

	if f.FieldRangeValidation {
		inRange := false
		for _, r := range f.FieldRanges {
			if value >= r.Start && value <= r.End {
				inRange = true
				break
			}
		}

		if !inRange {
			var ranges strings.Builder
			for _, r := range f.FieldRanges {
				ranges.WriteString(fmt.Sprintf("[%f,%f] ", r.Start, r.End))
			}
			result = multierror.Append(result, errors.Errorf("Value for %s should be in one of these ranges: %s", f.Name, ranges.String()))
		}
	}

	return result
}

// Required is called to make a field required in a JSON
func (f *FloatField) Required() *FloatField {
	f.FieldRequired = true
	return f
}

// Positive is called when we want to force the value to be positive in validation.
func (f *FloatField) Positive() *FloatField {
	f.FieldSignValidation = true
	f.FieldPositive = true
	return f
}

// Negative is called when we want to force the value to be negative in validation.
func (f *FloatField) Negative() *FloatField {
	f.FieldSignValidation = true
	f.FieldPositive = false
	return f
}

// Min is called when we want to set a minimum value for a float value in validation.
func (f *FloatField) Min(value float64) *FloatField {
	f.FieldMin = value
	f.FieldMinValidation = true
	return f
}

// Max is called when we want to set a maximum value for a float value in validation.
func (f *FloatField) Max(value float64) *FloatField {
	f.FieldMax = value
	f.FieldMaxValidation = true
	return f
}

// Range is called when we want to define valid ranges for a float value in validation.
func (f *FloatField) Range(start, end float64) *FloatField {
	f.FieldRanges = append(f.FieldRanges, floatRange{Start: start, End: end})
	f.FieldRangeValidation = true
	return f
}

// Float is the constructor of a float field
func Float(name string) *FloatField {
	return &FloatField{
		Name:                 name,
		FieldRequired:        false,
		FieldMin:             0,
		FieldMinValidation:   false,
		FieldMax:             0,
		FieldMaxValidation:   false,
		FieldSignValidation:  false,
		FieldPositive:        false,
		FieldRangeValidation: false,
		FieldRanges:          []floatRange{},
	}
}
