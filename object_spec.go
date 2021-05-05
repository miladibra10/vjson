package vjson

type ObjectFieldSpec struct {
	Name     string                 `mapstructure:"name"`
	Required bool                   `mapstructure:"required"`
	Schema   map[string]interface{} `mapstructure:"schema"`
}

func NewObject(spec ObjectFieldSpec, schema Schema) *ObjectField {
	return &ObjectField{
		name:     spec.Name,
		required: spec.Required,
		schema:   schema,
	}
}