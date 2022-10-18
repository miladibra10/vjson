package vjson

// StringFieldSpec is a type used for parsing an StringField
type StringFieldSpec struct {
	Name      string    `mapstructure:"name" json:"name"`
	Type      fieldType `json:"type"`
	Required  bool      `mapstructure:"required" json:"required,omitempty"`
	MinLength int       `mapstructure:"min_length" json:"minLength,omitempty"`
	MaxLength int       `mapstructure:"max_length" json:"maxLength,omitempty"`
	Format    string    `mapstructure:"format" json:"format,omitempty"`
	Choices   []string  `mapstructure:"choices" json:"choices,omitempty"`
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
