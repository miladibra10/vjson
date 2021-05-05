package vjson

type FloatRangeSpec struct {
	Start float64 `mapstructure:"start"`
	End   float64 `mapstructure:"end"`
}

type FloatFieldSpec struct {
	Name     string           `mapstructure:"name"`
	Required bool             `mapstructure:"required"`
	Min      float64          `mapstructure:"min"`
	Max      float64          `mapstructure:"max"`
	Positive bool             `mapstructure:"positive"`
	Ranges   []FloatRangeSpec `mapstructure:"ranges"`
}


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