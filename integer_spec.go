package vjson

// IntRangeSpec is a type for parsing an integer field range
type IntRangeSpec struct {
	Start int `mapstructure:"start"`
	End   int `mapstructure:"end"`
}

// IntegerFieldSpec is a type used for parsing an IntegerField
type IntegerFieldSpec struct {
	Name     string         `mapstructure:"name"`
	Required bool           `mapstructure:"required"`
	Min      int            `mapstructure:"min"`
	Max      int            `mapstructure:"max"`
	Positive bool           `mapstructure:"positive"`
	Ranges   []IntRangeSpec `mapstructure:"ranges"`
}

// NewInteger receives an IntegerFieldSpec and returns and IntegerField
func NewInteger(spec IntegerFieldSpec, minValidation, maxValidation, signValidation, rangeValidation bool) *IntegerField {
	ranges := make([]intRange, 0, len(spec.Ranges))
	for _, rangeSpec := range spec.Ranges {
		ranges = append(ranges, intRange{
			Start: rangeSpec.Start,
			End:   rangeSpec.End,
		})
	}
	return &IntegerField{
		Name:                 spec.Name,
		FieldRequired:        spec.Required,
		FieldMin:             spec.Min,
		FieldMinValidation:   minValidation,
		FieldMax:             spec.Max,
		FieldMaxValidation:   maxValidation,
		FieldSignValidation:  signValidation,
		FieldPositive:        spec.Positive,
		FieldRangeValidation: rangeValidation,
		FieldRanges:          ranges,
	}
}
