package vjson

// IntRangeSpec is a type for parsing an integer field range
type IntRangeSpec struct {
	Start int `mapstructure:"start" json:"start"`
	End   int `mapstructure:"end" json:"end"`
}

// IntegerFieldSpec is a type used for parsing an IntegerField
type IntegerFieldSpec struct {
	Name     string         `mapstructure:"name" json:"name"`
	Type     fieldType      `json:"type"`
	Required bool           `mapstructure:"required" json:"required,omitempty"`
	Min      int            `mapstructure:"min" json:"min,omitempty"`
	Max      int            `mapstructure:"max" json:"max,omitempty"`
	Positive bool           `mapstructure:"positive" json:"positive,omitempty"`
	Ranges   []IntRangeSpec `mapstructure:"ranges" json:"ranges,omitempty"`
}

// NewInteger receives an IntegerFieldSpec and returns and IntegerField
func NewInteger(spec IntegerFieldSpec, minValidation, maxValidation, signValidation, rangeValidation bool) *IntegerField {
	ranges := make([]intRange, 0, len(spec.Ranges))
	for _, rangeSpec := range spec.Ranges {
		ranges = append(ranges, intRange{
			start: rangeSpec.Start,
			end:   rangeSpec.End,
		})
	}
	return &IntegerField{
		name:            spec.Name,
		required:        spec.Required,
		min:             spec.Min,
		minValidation:   minValidation,
		max:             spec.Max,
		maxValidation:   maxValidation,
		signValidation:  signValidation,
		positive:        spec.Positive,
		rangeValidation: rangeValidation,
		ranges:          ranges,
	}
}
