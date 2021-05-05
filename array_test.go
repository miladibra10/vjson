package vjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayField_GetName(t *testing.T) {
	field := Array("foo", Integer("bar"))
	assert.Equal(t, "foo", field.GetName())
}

func TestArrayField_Validate(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Array("foo", Integer("bar")).Required()

			err := field.Validate(nil)
			assert.NotNil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Array("foo", Integer("bar")).Required()

			// input should be an array of interface
			err := field.Validate([]interface{}{1, 2})
			assert.Nil(t, err)
		})
		t.Run("invalid_value", func(t *testing.T) {
			t.Run("int_array", func(t *testing.T) {
				field := Array("foo", Integer("bar")).Required()

				// input should be an array of interface
				err := field.Validate([]int{1, 2})
				assert.NotNil(t, err)
			})
			t.Run("str_array", func(t *testing.T) {

				field := Array("foo", Integer("bar")).Required()

				err := field.Validate([]interface{}{"1", "2"})
				assert.NotNil(t, err)
			})
		})
	})
	t.Run("not_required", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Array("foo", Integer("bar"))

			err := field.Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Array("foo", Integer("bar"))

			err := field.Validate([]interface{}{1, 2})
			assert.Nil(t, err)
		})
		t.Run("invalid_value", func(t *testing.T) {
			t.Run("int_array", func(t *testing.T) {
				field := Array("foo", Integer("bar"))

				err := field.Validate([]int{1, 2})
				assert.NotNil(t, err)
			})
			t.Run("str_array", func(t *testing.T) {
				field := Array("foo", Integer("bar"))

				err := field.Validate([]interface{}{"1", "2"})
				assert.NotNil(t, err)
			})
		})
	})
	t.Run("min_length", func(t *testing.T) {
		field := Array("foo", Integer("bar")).MinLength(3)

		err := field.Validate([]interface{}{1, 2, 3})
		assert.Nil(t, err)

		err = field.Validate([]interface{}{1, 2, 3, 4})
		assert.Nil(t, err)

		err = field.Validate([]interface{}{1, 2})
		assert.NotNil(t, err)
	})
	t.Run("max_length", func(t *testing.T) {
		field := Array("foo", Integer("bar")).MaxLength(3)

		err := field.Validate([]interface{}{1, 2, 3})
		assert.Nil(t, err)

		err = field.Validate([]interface{}{1, 2, 3, 4})
		assert.NotNil(t, err)

		err = field.Validate([]interface{}{1, 2})
		assert.Nil(t, err)
	})
}

func TestNewArray(t *testing.T) {
	field := NewArray(ArrayFieldSpec{
		Name: "bar",
		Required: true,
	}, String("foo"), false, false)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, false, field.minLengthValidation)
	assert.Equal(t, false, field.maxLengthValidation)
}