package vjson

// BooleanFieldSpec is a type used for parsing an BooleanField
type BooleanFieldSpec struct {
	Name     string `mapstructure:"name"`
	Required bool   `mapstructure:"required"`
	Value    bool   `mapstructure:"value"`
}

// NewBoolean receives an BooleanFieldSpec and returns and BooleanField
func NewBoolean(spec BooleanFieldSpec, valueValidation bool) *BooleanField {
	return &BooleanField{
		name:            spec.Name,
		required:        spec.Required,
		valueValidation: valueValidation,
		value:           spec.Value,
	}
}
