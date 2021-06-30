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
		name:              spec.Name,
		required:          spec.Required,
		validateMinLength: minLengthValidation,
		minLength:         spec.MinLength,
		validateMaxLength: maxLengthValidation,
		maxLength:         spec.MaxLength,
		validateFormat:    formatValidation,
		format:            spec.Format,
		validateChoices:   choiceValidation,
		choices:           spec.Choices,
	}
}
