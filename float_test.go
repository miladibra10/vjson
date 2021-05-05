package vjson

import (
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
}

func TestNewFloat(t *testing.T) {
	field := NewFloat(FloatFieldSpec{
		Name: "bar",
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