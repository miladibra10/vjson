package vjson

// NullFieldSpec is a type used for parsing an NullField
type NullFieldSpec struct {
	Name string `mapstructure:"name"`
}

// NewNull receives an NullFieldSpec and returns and NullField
func NewNull(spec NullFieldSpec) *NullField {
	return &NullField{
		Name: spec.Name,
	}
}
