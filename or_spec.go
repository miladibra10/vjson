package vjson

// OrFieldSpec is a type used for parsing an OrField
type OrFieldSpec struct {
	Name     string  `mapstructure:"name" json:"name"`
	Type     string  `json:"type"`
	Required bool    `mapstructure:"required" json:"required,omitempty"`
	Fields   []map[string]interface{} `mapstructure:"fields" json:"fields"`
}

// NewOr receives an OrFieldSpec and returns an OrField
func NewOr(spec OrFieldSpec) *OrField {
	return &OrField{
		name:     spec.Name,
		required: spec.Required,
		fields:   []Field{},
	}
}
