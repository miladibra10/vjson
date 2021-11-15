package vjson

// StringFieldSpec is a type used for parsing an StringField
type StringFieldSpec struct {
	Name      string   `mapstructure:"name"`
	Required  bool     `mapstructure:"required"`
	MinLength int      `mapstructure:"min_length"`
	MaxLength int      `mapstructure:"max_length"`
	Format    string   `mapstructure:"format"`
	Choices   []string `mapstructure:"choices"`
}

// NewString receives an StringFieldSpec and returns and StringField
func NewString(spec StringFieldSpec, minLengthValidation, maxLengthValidation, formatValidation, choiceValidation bool) *StringField {
	return &StringField{
		Name:                   spec.Name,
		FieldRequired:          spec.Required,
		FieldValidateMinLength: minLengthValidation,
		FieldMinLength:         spec.MinLength,
		FieldValidateMaxLength: maxLengthValidation,
		FieldMaxLength:         spec.MaxLength,
		FieldValidateFormat:    formatValidation,
		FieldFormat:            spec.Format,
		FieldValidateChoices:   choiceValidation,
		FieldChoices:           spec.Choices,
	}
}
