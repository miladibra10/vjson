package vjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooleanField_GetName(t *testing.T) {
	field := Boolean("foo")
	assert.Equal(t, "foo", field.GetName())
}

func TestBooleanField_Validate(t *testing.T) {
	t.Run("invalid_input", func(t *testing.T) {
		field := Boolean("foo")

		err := field.Validate("true")
		assert.NotNil(t, err)
	})
	t.Run("not_required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Boolean("foo")

			err := field.Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Boolean("foo")
			err := field.Validate(true)
			assert.Nil(t, err)
		})
	})
	t.Run("required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Boolean("foo").Required()

			err := field.Validate(nil)
			assert.NotNil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Boolean("foo").Required()

			err := field.Validate(false)
			assert.Nil(t, err)
		})
	})
	t.Run("should_be", func(t *testing.T) {
		field := Boolean("foo").Required().ShouldBe(true)

		err := field.Validate(true)
		assert.Nil(t, err)

		err = field.Validate(false)
		assert.NotNil(t, err)
	})
}

func TestNewBoolean(t *testing.T) {
	field := NewBoolean(BooleanFieldSpec{
		Name: "bar",
		Required: true,
	}, false)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, false, field.valueValidation)
}