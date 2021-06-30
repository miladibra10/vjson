package vjson

// ObjectFieldSpec is a type used for parsing an ObjectField
type ObjectFieldSpec struct {
	Name     string                 `mapstructure:"name"`
	Required bool                   `mapstructure:"required"`
	Schema   map[string]interface{} `mapstructure:"schema"`
}

// NewObject receives an ObjectFieldSpec and returns and ObjectField
func NewObject(spec ObjectFieldSpec, schema Schema) *ObjectField {
	return &ObjectField{
		name:     spec.Name,
		required: spec.Required,
		schema:   schema,
	}
}
