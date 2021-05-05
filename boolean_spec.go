package vjson

type BooleanFieldSpec struct {
	Name     string `mapstructure:"name"`
	Required bool   `mapstructure:"required"`
	Value    bool   `mapstructure:"value"`
}


func NewBoolean(spec BooleanFieldSpec, valueValidation bool) *BooleanField {
	return &BooleanField{
		name:            spec.Name,
		required:        spec.Required,
		valueValidation: valueValidation,
		value:           spec.Value,
	}
}
