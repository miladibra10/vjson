package vjson

type NullFieldSpec struct {
	Name      string   `mapstructure:"name"`
}

func NewNull(spec NullFieldSpec) *NullField {
	return &NullField{
		name:              spec.Name,
	}
}

