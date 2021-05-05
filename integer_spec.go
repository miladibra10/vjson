package vjson

type IntRangeSpec struct {
	Start int `mapstructure:"start"`
	End   int `mapstructure:"end"`
}

type IntegerFieldSpec struct {
	Name     string         `mapstructure:"name"`
	Required bool           `mapstructure:"required"`
	Min      int            `mapstructure:"min"`
	Max      int            `mapstructure:"max"`
	Positive bool           `mapstructure:"positive"`
	Ranges   []IntRangeSpec `mapstructure:"ranges"`
}

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
