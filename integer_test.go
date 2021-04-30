package vjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerField_Validate(t *testing.T) {
	t.Run("invalid_input", func(t *testing.T) {
		field := Integer("foo")

		err := field.Validate("Hi")
		assert.NotNil(t, err)
	})
	t.Run("not_required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Integer("foo")

			err := field.Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Integer("foo")

			// We asume we get real float values because be use gjson library
			err := field.Validate(2)
			assert.Nil(t, err)
		})
		t.Run("valid_value_float", func(t *testing.T) {
			field := Integer("foo")

			// We asume we get real float values because be use gjson library
			err := field.Validate(float64(2))
			assert.Nil(t, err)
		})
	})
	t.Run("required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Integer("foo").Required()

			err := field.Validate(nil)
			assert.NotNil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Integer("foo").Required()

			// We asume we get real float values because be use gjson library
			err := field.Validate(2)
			assert.Nil(t, err)
		})
		t.Run("valid_value_float", func(t *testing.T) {
			field := Integer("foo").Required()

			// We asume we get real float values because be use gjson library
			err := field.Validate(float64(2))
			assert.Nil(t, err)
		})
	})
	t.Run("positive", func(t *testing.T) {
		field := Integer("foo").Required().Positive()

		err := field.Validate(1)
		assert.Nil(t, err)


		err = field.Validate(-1)
		assert.NotNil(t, err)
	})
	t.Run("negative", func(t *testing.T) {
		field := Integer("foo").Required().Negative()

		err := field.Validate(1)
		assert.NotNil(t, err)


		err = field.Validate(-1)
		assert.Nil(t, err)
	})
	t.Run("min", func(t *testing.T) {
		field := Integer("foo").Required().Min(10)

		err := field.Validate(12)
		assert.Nil(t, err)


		err = field.Validate(2)
		assert.NotNil(t, err)
	})
	t.Run("max", func(t *testing.T) {
		field := Integer("foo").Required().Max(10)

		err := field.Validate(9)
		assert.Nil(t, err)


		err = field.Validate(13)
		assert.NotNil(t, err)
	})
	t.Run("ranges", func(t *testing.T) {
		field := Integer("foo").Required().Range(-10, 10).Range(20, 30)

		err := field.Validate(float64(2))
		assert.Nil(t, err)

		err = field.Validate(2)
		assert.Nil(t, err)

		err = field.Validate(float64(25))
		assert.Nil(t, err)

		err = field.Validate(25)
		assert.Nil(t, err)

		err = field.Validate(float64(100))
		assert.NotNil(t, err)

		err = field.Validate(100)
		assert.NotNil(t, err)
	})
}
