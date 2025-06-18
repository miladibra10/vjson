package vjson

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatField_GetName(t *testing.T) {
	field := Float("foo")
	assert.Equal(t, "foo", field.GetName())
}

func TestFloatField_Validate(t *testing.T) {
	t.Run("invalid_input", func(t *testing.T) {
		field := Float("foo")

		err := field.Validate("Hi")
		assert.NotNil(t, err)
	})
	t.Run("not_required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Float("foo")

			err := field.Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("valid_value_float", func(t *testing.T) {
			field := Float("foo")

			err := field.Validate(float64(2))
			assert.Nil(t, err)
		})
	})
	t.Run("required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Float("foo").Required()

			err := field.Validate(nil)
			assert.NotNil(t, err)
		})
		t.Run("valid_value_float", func(t *testing.T) {
			field := Float("foo").Required()
			err := field.Validate(float64(2))
			assert.Nil(t, err)
		})
	})
	t.Run("positive", func(t *testing.T) {
		field := Float("foo").Required().Positive()

		err := field.Validate(float64(1))
		assert.Nil(t, err)

		err = field.Validate(float64(-1))
		assert.NotNil(t, err)
	})
	t.Run("negative", func(t *testing.T) {
		field := Float("foo").Required().Negative()

		err := field.Validate(float64(1))
		assert.NotNil(t, err)

		err = field.Validate(float64(-1))
		assert.Nil(t, err)
	})
	t.Run("min", func(t *testing.T) {
		field := Float("foo").Required().Min(10)

		err := field.Validate(float64(12))
		assert.Nil(t, err)

		err = field.Validate(float64(2))
		assert.NotNil(t, err)
	})
	t.Run("max", func(t *testing.T) {
		field := Float("foo").Required().Max(10)

		err := field.Validate(float64(9))
		assert.Nil(t, err)

		err = field.Validate(float64(13))
		assert.NotNil(t, err)
	})
	t.Run("ranges", func(t *testing.T) {
		field := Float("foo").Required().Range(-10, 10).Range(20, 30)

		err := field.Validate(float64(2))
		assert.Nil(t, err)

		err = field.Validate(float64(25))
		assert.Nil(t, err)

		err = field.Validate(float64(100))
		assert.NotNil(t, err)
	})

	t.Run("invalid_range", func(t *testing.T) {
		field := Float("foo").Required().Range(10.5, 5.5) // Start > End

		err := field.Validate(float64(7.5))
		assert.NotNil(t, err)

		err = field.Validate(float64(12.5))
		assert.NotNil(t, err)

		err = field.Validate(float64(3.5))
		assert.NotNil(t, err)
	})

	t.Run("overlapping_ranges", func(t *testing.T) {
		field := Float("foo").Required().Range(1.5, 10.5).Range(5.5, 15.5)

		err := field.Validate(float64(3.5))
		assert.Nil(t, err)

		err = field.Validate(float64(7.5))
		assert.Nil(t, err) // In both ranges

		err = field.Validate(float64(12.5))
		assert.Nil(t, err)

		err = field.Validate(float64(20.5))
		assert.NotNil(t, err)
	})

	t.Run("zero_validation", func(t *testing.T) {
		t.Run("positive", func(t *testing.T) {
			field := Float("foo").Required().Positive()

			// The implementation actually allows zero as positive
			err := field.Validate(float64(0))
			assert.Nil(t, err)

			err = field.Validate(float64(0.1))
			assert.Nil(t, err)
		})

		t.Run("negative", func(t *testing.T) {
			field := Float("foo").Required().Negative()

			// The implementation actually allows zero as negative
			err := field.Validate(float64(0))
			assert.Nil(t, err)

			err = field.Validate(float64(-0.1))
			assert.Nil(t, err)
		})
	})

	t.Run("integer_values", func(t *testing.T) {
		field := Float("foo").Required().Range(1.5, 10.5)

		// The implementation doesn't accept integer values for float fields
		err := field.Validate(5)
		assert.NotNil(t, err)

		// But it would accept float values
		err = field.Validate(float64(5))
		assert.Nil(t, err)

		err = field.Validate(float64(15))
		assert.NotNil(t, err)
	})

	t.Run("combined_validations", func(t *testing.T) {
		t.Run("positive_min_max", func(t *testing.T) {
			field := Float("foo").Required().Positive().Min(5.5).Max(10.5)

			err := field.Validate(float64(3.5))
			assert.NotNil(t, err) // Less than min

			err = field.Validate(float64(7.5))
			assert.Nil(t, err)

			err = field.Validate(float64(12.5))
			assert.NotNil(t, err) // Greater than max

			err = field.Validate(float64(-3.5))
			assert.NotNil(t, err) // Not positive
		})

		t.Run("negative_min_max", func(t *testing.T) {
			field := Float("foo").Required().Negative().Min(-10.5).Max(-5.5)

			err := field.Validate(float64(-12.5))
			assert.NotNil(t, err) // Less than min

			err = field.Validate(float64(-7.5))
			assert.Nil(t, err)

			err = field.Validate(float64(-3.5))
			assert.NotNil(t, err) // Greater than max

			err = field.Validate(float64(3.5))
			assert.NotNil(t, err) // Not negative
		})

		t.Run("range_and_min_max", func(t *testing.T) {
			field := Float("foo").Required().Range(1.5, 10.5).Min(5.5).Max(15.5)

			err := field.Validate(float64(3.5))
			assert.NotNil(t, err) // Less than min

			err = field.Validate(float64(7.5))
			assert.Nil(t, err)

			// The implementation requires values to be within the specified ranges,
			// even if they're within min/max
			err = field.Validate(float64(12.5))
			assert.NotNil(t, err) // Within max but outside range

			err = field.Validate(float64(20.5))
			assert.NotNil(t, err) // Greater than max
		})
	})
}

func TestFloatField_MarshalJSON(t *testing.T) {
	field := Float("foo").Range(10, 20)
	b, err := json.Marshal(field)
	assert.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(b, &data)
	assert.Nil(t, err)

	assert.Equal(t, "foo", data["name"])
	assert.Equal(t, string(floatType), data["type"])
	assert.Equal(t, float64(10), data["ranges"].([]interface{})[0].(map[string]interface{})["start"])
	assert.Equal(t, float64(20), data["ranges"].([]interface{})[0].(map[string]interface{})["end"])
}

func TestNewFloat(t *testing.T) {
	field := NewFloat(FloatFieldSpec{
		Name:     "bar",
		Required: true,
		Ranges: []FloatRangeSpec{
			{
				Start: 0,
				End:   20,
			},
		},
	}, false, false, false, true)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, false, field.minValidation)
	assert.Equal(t, false, field.maxValidation)
	assert.Equal(t, false, field.signValidation)
	assert.Equal(t, true, field.rangeValidation)
}
