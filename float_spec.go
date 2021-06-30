package vjson

// FloatRangeSpec is a type for parsing a float field range
type FloatRangeSpec struct {
	Start float64 `mapstructure:"start"`
	End   float64 `mapstructure:"end"`
}

// FloatFieldSpec is a type used for parsing an FloatField
type FloatFieldSpec struct {
	Name     string           `mapstructure:"name"`
	Required bool             `mapstructure:"required"`
	Min      float64          `mapstructure:"min"`
	Max      float64          `mapstructure:"max"`
	Positive bool             `mapstructure:"positive"`
	Ranges   []FloatRangeSpec `mapstructure:"ranges"`
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
