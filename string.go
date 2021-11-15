package vjson

import (
	"regexp"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// StringField is the type for validating strings in a JSON
type StringField struct {
	Name          string `json:"name"`
	FieldRequired bool   `json:"required"`

	FieldValidateMinLength bool `json:"validateMinLength"`
	FieldMinLength         int  `json:"minLength"`

	FieldValidateMaxLength bool `json:"validateMaxLength"`
	FieldMaxLength         int  `json:"maxLength"`

	FieldValidateFormat bool   `json:"validateFormat"`
	FieldFormat         string `json:"format"`

	FieldValidateChoices bool     `json:"validateChoices"`
	FieldChoices         []string `json:"choices"`
}

// To Force Implementing Field interface by StringField
var _ Field = (*StringField)(nil)

// GetName returns name of the field
func (s *StringField) GetName() string {
	return s.Name
}

// Required is called to make a field required in a JSON
func (s *StringField) Required() *StringField {
	s.FieldRequired = true
	return s
}

// MinLength is called to set a minimum length to a string field
func (s *StringField) MinLength(length int) *StringField {
	if length < 0 {
		return s
	}
	s.FieldMinLength = length
	s.FieldValidateMinLength = true
	return s
}

// MaxLength is called to set a maximum length to a string field
func (s *StringField) MaxLength(length int) *StringField {
	if length < 0 {
		return s
	}
	s.FieldMaxLength = length
	s.FieldValidateMaxLength = true
	return s
}

// Format is called to set a regex format for validation of a string field
func (s *StringField) Format(format string) *StringField {
	s.FieldFormat = format
	s.FieldValidateFormat = true
	return s
}

// Choices is called to set valid choices of a string field in validation
func (s *StringField) Choices(choices ...string) *StringField {
	s.FieldChoices = choices
	s.FieldValidateChoices = true
	return s
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (s *StringField) Validate(value interface{}) error {
	if value == nil {
		if !s.FieldRequired {
			return nil
		}
		return errors.Errorf("Value for %s field is required", s.Name)
	}

	stringValue, ok := value.(string)

	if !ok {
		return errors.Errorf("Value for %s should be a string", s.Name)
	}

	var result error

	if s.FieldValidateMinLength {
		if len(stringValue) < s.FieldMinLength {
			result = multierror.Append(result, errors.Errorf("Value for %s field should have at least %d characters", s.Name, s.FieldMinLength))
		}
	}

	if s.FieldValidateMaxLength {
		if len(stringValue) > s.FieldMaxLength {
			result = multierror.Append(result, errors.Errorf("Value for %s field should have at most %d characters", s.Name, s.FieldMaxLength))
		}
	}

	if s.FieldValidateChoices {
		for _, choice := range s.FieldChoices {
			if stringValue == choice {
				return nil
			}
		}
		result = multierror.Append(result, errors.Errorf("Value for %s field should be one of: [%s] values", s.Name, strings.Join(s.FieldChoices, ",")))
	}

	if s.FieldValidateFormat {
		r, err := regexp.Compile(s.FieldFormat)
		if err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "Invalid StringField format string for field %s", s.Name))
			return result
		}

		isValidFormat := r.MatchString(stringValue)

		if !isValidFormat {
			result = multierror.Append(result, errors.Wrapf(err, "Invalid StringField format string for field %s", s.Name))
		}
	}

	return result
}

// String is the constructor of a string field
func String(name string) *StringField {
	return &StringField{
		Name:          name,
		FieldRequired: false,
		FieldChoices:  []string{},
	}
}
