package vjson

// FloatRangeSpec is a type for parsing a float field range
type FloatRangeSpec struct {
	Start float64 `mapstructure:"start" json:"start"`
	End   float64 `mapstructure:"end" json:"end"`
}

// FloatFieldSpec is a type used for parsing an FloatField
type FloatFieldSpec struct {
	Name     string           `mapstructure:"name" json:"name"`
	Type     fieldType        `json:"type"`
	Required bool             `mapstructure:"required" json:"required,omitempty"`
	Min      float64          `mapstructure:"min" json:"min,omitempty"`
	Max      float64          `mapstructure:"max" json:"max,omitempty"`
	Positive bool             `mapstructure:"positive" json:"positive,omitempty"`
	Ranges   []FloatRangeSpec `mapstructure:"ranges" json:"ranges,omitempty"`
}

// NewFloat receives an FloatFieldSpec and returns and FloatField
func NewFloat(spec FloatFieldSpec, minValidation, maxValidation, signValidation, rangeValidation bool) *FloatField {
	ranges := make([]floatRange, 0, len(spec.Ranges))
	for _, rangeSpec := range spec.Ranges {
		ranges = append(ranges, floatRange{
			start: rangeSpec.Start,
			end:   rangeSpec.End,
		})
	}
	return &FloatField{
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
