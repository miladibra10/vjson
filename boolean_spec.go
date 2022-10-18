package vjson

// BooleanFieldSpec is a type used for parsing an BooleanField
type BooleanFieldSpec struct {
	Name     string    `mapstructure:"name" json:"name"`
	Type     fieldType `json:"type"`
	Required bool      `mapstructure:"required" json:"required,omitempty"`
	Value    bool      `mapstructure:"value" json:"value,omitempty"`
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
