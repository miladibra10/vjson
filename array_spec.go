package vjson

// ArrayFieldSpec is a type used for parsing an ArrayField
type ArrayFieldSpec struct {
	Name      string                 `mapstructure:"name" json:"name"`
	Type      fieldType              `json:"type"`
	Required  bool                   `mapstructure:"required" json:"required,omitempty"`
	Items     map[string]interface{} `mapstructure:"items" json:"items,omitempty"`
	MinLength int                    `mapstructure:"min_length" json:"minLength,omitempty"`
	MaxLength int                    `mapstructure:"max_length" json:"maxLength,omitempty"`

	FixItems []map[string]interface{} `mapstructure:"fix_items" json:"fix_items,omitempty"`
}

// NewArray receives an ArrayFieldSpec and returns and ArrayField
func NewArray(spec ArrayFieldSpec, itemField Field, fixItemsField []Field, minLengthValidation, maxLengthValidation bool) *ArrayField {
	return &ArrayField{
		name:                spec.Name,
		required:            spec.Required,
		items:               itemField,
		fixItems:            fixItemsField,
		minLength:           spec.MinLength,
		minLengthValidation: minLengthValidation,
		maxLength:           spec.MaxLength,
		maxLengthValidation: maxLengthValidation,
	}
}
