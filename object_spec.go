package vjson

// ObjectFieldSpec is a type used for parsing an ObjectField
type ObjectFieldSpec struct {
	Name     string                 `mapstructure:"name" json:"name"`
	Type     fieldType              `json:"type"`
	Required bool                   `mapstructure:"required" json:"required,omitempty"`
	Schema   map[string]interface{} `mapstructure:"schema" json:"schema,omitempty"`
}

// NewObject receives an ObjectFieldSpec and returns and ObjectField
func NewObject(spec ObjectFieldSpec, schema Schema) *ObjectField {
	return &ObjectField{
		name:     spec.Name,
		required: spec.Required,
		schema:   schema,
	}
}
