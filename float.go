package vjson


import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"strings"
)

type floatRange struct {
	start float64
	end   float64
}

type FloatField struct {
	name     string
	required bool

	min           float64
	minValidation bool

	max           float64
	maxValidation bool

	signValidation bool
	positive       bool

	rangeValidation bool
	ranges          []floatRange
}

// To Force Implementing Field interface by IntegerField
var _ Field = (*FloatField)(nil)

func (f *FloatField) Validate(v interface{}) error {
	if v == nil {
		if !f.required {
			return nil
		} else {
			return errors.Errorf("Value for %s field is required", f.name)
		}
	}

	value, ok := v.(float64)

	if !ok {
		return errors.Errorf("Value for %s should be a float number", f.name)
	}


	var result error
	if f.signValidation && f.positive {
		if value < 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a positive float", f.name))
		}
	} else if f.signValidation && !f.positive {
		if value > 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a negative float", f.name))
		}
	}

	if f.minValidation {
		if value < f.min {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at least %f", f.name, f.min))
		}
	}

	if f.maxValidation {
		if value > f.max {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at most %f", f.name, f.max))
		}
	}

	if f.rangeValidation {
		inRange := false
		for _, r := range f.ranges {
			if value >= r.start && value <= r.end {
				inRange = true
				break
			}
		}

		if !inRange {
			var ranges strings.Builder
			for _, r := range f.ranges {
				ranges.WriteString(fmt.Sprintf("[%f,%f] ", r.start, r.end))
			}
			result = multierror.Append(result, errors.Errorf("Value for %s should be in one of these ranges: %s", f.name, ranges.String()))
		}
	}

	return result
}

func (f *FloatField) Required() *FloatField {
	f.required = true
	return f
}

func (f *FloatField) Positive() *FloatField {
	f.signValidation = true
	f.positive = true
	return f
}

func (f *FloatField) Negative() *FloatField {
	f.signValidation = true
	f.positive = false
	return f
}

func (f *FloatField) Min(value float64) *FloatField {
	f.min = value
	f.minValidation = true
	return f
}

func (f *FloatField) Max(value float64) *FloatField {
	f.max = value
	f.maxValidation = true
	return f
}

func (f *FloatField) Range(start, end float64) *FloatField {
	f.ranges = append(f.ranges, floatRange{start: start, end: end})
	f.rangeValidation = true
	return f
}

func Float(name string) *FloatField {
	return &FloatField{
		name:            name,
		required:        false,
		min:             0,
		minValidation:   false,
		max:             0,
		maxValidation:   false,
		signValidation:  false,
		positive:        false,
		rangeValidation: false,
		ranges:          []floatRange{},
	}
}

