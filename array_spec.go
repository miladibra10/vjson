package vjson

type ArrayFieldSpec struct {
	Name      string                 `mapstructure:"name"`
	Required  bool                   `mapstructure:"required"`
	Items     map[string]interface{} `mapstructure:"items"`
	MinLength int                    `mapstructure:"min_length"`
	MaxLength int                    `mapstructure:"max_length"`
}

func NewArray(spec ArrayFieldSpec, itemField Field, minLengthValidation, maxLengthValidation bool) *ArrayField {
	return &ArrayField{
		name:                spec.Name,
		required:            spec.Required,
		items:               itemField,
		minLength:           spec.MinLength,
		minLengthValidation: minLengthValidation,
		maxLength:           spec.MaxLength,
		maxLengthValidation: maxLengthValidation,
	}
}
