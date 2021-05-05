package vjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringField_GetName(t *testing.T) {
	field := String("foo")
	assert.Equal(t, "foo", field.GetName())
}

func TestStringField_Validate(t *testing.T) {
	t.Run("invalid_input", func(t *testing.T) {
		field := String("foo")

		err := field.Validate(1)
		assert.NotNil(t, err)
	})
	t.Run("not_required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := String("foo")

			err := field.Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := String("foo")

			err := field.Validate("Hi")
			assert.Nil(t, err)
		})
	})
	t.Run("required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := String("foo").Required()

			err := field.Validate(nil)
			assert.NotNil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := String("foo").Required()

			err := field.Validate("Hi")
			assert.Nil(t, err)
		})
	})
	t.Run("min_length", func(t *testing.T) {
		t.Run("invalid_length", func(t *testing.T) {
			field := String("foo").MinLength(-1)

			assert.Equal(t, false, field.validateMinLength)
			assert.Equal(t, 0, field.minLength)
		})
		t.Run("valid_length", func(t *testing.T) {
			field := String("foo").MinLength(5)

			t.Run("valid_input", func(t *testing.T) {
				err := field.Validate("12345")
				assert.Nil(t, err)
			})
			t.Run("invalid_input", func(t *testing.T) {
				err := field.Validate("1234")
				assert.NotNil(t, err)
			})
		})
	})
	t.Run("max_length", func(t *testing.T) {
		t.Run("invalid_length", func(t *testing.T) {
			field := String("foo").MaxLength(-1)

			assert.Equal(t, false, field.validateMaxLength)
			assert.Equal(t, 0, field.maxLength)
		})
		t.Run("valid_length", func(t *testing.T) {
			field := String("foo").MaxLength(5)

			t.Run("valid_input", func(t *testing.T) {
				err := field.Validate("123")
				assert.Nil(t, err)
			})
			t.Run("invalid_input", func(t *testing.T) {
				err := field.Validate("123456")
				assert.NotNil(t, err)
			})
		})
	})
	t.Run("choices", func(t *testing.T) {
		t.Run("valid_choice", func(t *testing.T) {
			field := String("foo").Choices("A", "B")

			err := field.Validate("A")
			assert.Nil(t, err)

			err = field.Validate("B")
			assert.Nil(t, err)

			err = field.Validate("AB")
			assert.NotNil(t, err)
		})
	})
	t.Run("format", func(t *testing.T) {
		t.Run("valid_format", func(t *testing.T) {

			field := String("foo").Format("p([a-z]+)ch")

			err := field.Validate("peach")
			assert.Nil(t, err)

			err = field.Validate("pach")
			assert.Nil(t, err)

			err = field.Validate("HI")
			assert.NotNil(t, err)

			err = field.Validate("foo")
			assert.NotNil(t, err)
		})
		t.Run("invalid_format", func(t *testing.T) {

			field := String("foo").Format(")(")

			err := field.Validate("peach")
			assert.NotNil(t, err)

		})

	})
}

func TestNewString(t *testing.T) {
	field := NewString(StringFieldSpec{
		Name: "bar",
		Required: true,
	}, false, false, false, false)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, false, field.validateMinLength)
	assert.Equal(t, false, field.validateMaxLength)
	assert.Equal(t, false, field.validateFormat)
	assert.Equal(t, false, field.validateChoices)
}
