package vjson

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"strings"
)

type intRange struct {
	start int
	end   int
}

type IntegerField struct {
	name     string
	required bool

	min           int
	minValidation bool

	max           int
	maxValidation bool

	signValidation bool
	positive       bool

	rangeValidation bool
	ranges          []intRange
}

// To Force Implementing Field interface by IntegerField
var _ Field = (*IntegerField)(nil)

func (i *IntegerField) GetName() string {
	return i.name
}

func (i *IntegerField) Validate(v interface{}) error {
	if v == nil {
		if !i.required {
			return nil
		} else {
			return errors.Errorf("Value for %s field is required", i.name)
		}
	}
	var value int
	var intOK bool

	//gjson library returns float64 when field value is number
	floatValue, floatOK := v.(float64)

	value, intOK = v.(int)

	if !floatOK && !intOK {
		return errors.Errorf("Value for %s should be a number", i.name)
	}

	if floatOK {
		value = int(floatValue)
	}

	var result error
	if i.signValidation && i.positive {
		if value < 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a positive integer", i.name))
		}
	} else if i.signValidation && !i.positive {
		if value > 0 {
			result = multierror.Append(result, errors.Errorf("Value for %s should be a negative integer", i.name))
		}
	}

	if i.minValidation {
		if value < i.min {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at least %d", i.name, i.min))
		}
	}

	if i.maxValidation {
		if value > i.max {
			result = multierror.Append(result, errors.Errorf("Value for %s should be at most %d", i.name, i.max))
		}
	}

	if i.rangeValidation {
		inRange := false
		for _, r := range i.ranges {
			if value >= r.start && value <= r.end {
				inRange = true
				break
			}
		}

		if !inRange {
			var ranges strings.Builder
			for _, r := range i.ranges {
				ranges.WriteString(fmt.Sprintf("[%d,%d] ", r.start, r.end))
			}
			result = multierror.Append(result, errors.Errorf("Value for %s should be in one of these ranges: %s", i.name, ranges.String()))
		}
	}

	return result
}

func (i *IntegerField) Required() *IntegerField {
	i.required = true
	return i
}

func (i *IntegerField) Positive() *IntegerField {
	i.signValidation = true
	i.positive = true
	return i
}

func (i *IntegerField) Negative() *IntegerField {
	i.signValidation = true
	i.positive = false
	return i
}

func (i *IntegerField) Min(value int) *IntegerField {
	i.min = value
	i.minValidation = true
	return i
}

func (i *IntegerField) Max(value int) *IntegerField {
	i.max = value
	i.maxValidation = true
	return i
}

func (i *IntegerField) Range(start, end int) *IntegerField {
	i.ranges = append(i.ranges, intRange{start: start, end: end})
	i.rangeValidation = true
	return i
}

func Integer(name string) *IntegerField {
	return &IntegerField{
		name:            name,
		required:        false,
		min:             0,
		minValidation:   false,
		max:             0,
		maxValidation:   false,
		signValidation:  false,
		positive:        false,
		rangeValidation: false,
		ranges:          []intRange{},
	}
}
