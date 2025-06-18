package vjson

import (
	"encoding/json"
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

	t.Run("invalid_length_values", func(t *testing.T) {
		t.Run("negative_min_length", func(t *testing.T) {
			field := Array("foo", Integer("bar")).MinLength(-1)

			// The implementation actually sets the validation to true and keeps the negative value
			assert.Equal(t, true, field.minLengthValidation)
			assert.Equal(t, -1, field.minLength)
		})

		t.Run("negative_max_length", func(t *testing.T) {
			field := Array("foo", Integer("bar")).MaxLength(-1)

			// The implementation actually sets the validation to true and keeps the negative value
			assert.Equal(t, true, field.maxLengthValidation)
			assert.Equal(t, -1, field.maxLength)
		})
	})

	t.Run("empty_array", func(t *testing.T) {
		t.Run("with_min_length", func(t *testing.T) {
			field := Array("foo", Integer("bar")).MinLength(1)

			err := field.Validate([]interface{}{})
			assert.NotNil(t, err)
		})

		t.Run("without_min_length", func(t *testing.T) {
			field := Array("foo", Integer("bar"))

			err := field.Validate([]interface{}{})
			assert.Nil(t, err)
		})
	})

	t.Run("combined_validations", func(t *testing.T) {
		field := Array("foo", Integer("bar")).MinLength(2).MaxLength(4)

		err := field.Validate([]interface{}{1})
		assert.NotNil(t, err) // Too short

		err = field.Validate([]interface{}{1, 2})
		assert.Nil(t, err)

		err = field.Validate([]interface{}{1, 2, 3, 4})
		assert.Nil(t, err)

		err = field.Validate([]interface{}{1, 2, 3, 4, 5})
		assert.NotNil(t, err) // Too long
	})

	t.Run("different_item_types", func(t *testing.T) {
		t.Run("string_array", func(t *testing.T) {
			field := Array("foo", String("bar"))

			err := field.Validate([]interface{}{"a", "b", "c"})
			assert.Nil(t, err)

			err = field.Validate([]interface{}{1, 2, 3})
			assert.NotNil(t, err) // Integers in a string array
		})

		t.Run("boolean_array", func(t *testing.T) {
			field := Array("foo", Boolean("bar"))

			err := field.Validate([]interface{}{true, false, true})
			assert.Nil(t, err)

			err = field.Validate([]interface{}{"true", "false"})
			assert.NotNil(t, err) // Strings in a boolean array
		})
	})

	t.Run("nested_arrays", func(t *testing.T) {
		innerArrayField := Array("inner", Integer("value"))
		outerArrayField := Array("outer", innerArrayField)

		// Valid nested array: [[1, 2], [3, 4]]
		validNestedArray := []interface{}{
			[]interface{}{1, 2},
			[]interface{}{3, 4},
		}
		err := outerArrayField.Validate(validNestedArray)
		assert.Nil(t, err)

		// Invalid nested array: [[1, 2], ["a", "b"]]
		invalidNestedArray := []interface{}{
			[]interface{}{1, 2},
			[]interface{}{"a", "b"}, // String values in an integer array
		}
		err = outerArrayField.Validate(invalidNestedArray)
		assert.NotNil(t, err)

		// Invalid nested array: [[1, 2], 3]
		invalidTypeArray := []interface{}{
			[]interface{}{1, 2},
			3, // Not an array
		}
		err = outerArrayField.Validate(invalidTypeArray)
		assert.NotNil(t, err)
	})
}

func TestArrayField_MarshalJSON(t *testing.T) {
	field := Array("foo", String("bar"))

	b, err := json.Marshal(field)
	assert.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(b, &data)
	assert.Nil(t, err)

	assert.Equal(t, "foo", data["name"])
	assert.Equal(t, string(arrayType), data["type"])
	assert.Equal(t, "bar", data["items"].(map[string]interface{})["name"])
}

func TestNewArray(t *testing.T) {
	field := NewArray(ArrayFieldSpec{
		Name:     "bar",
		Required: true,
	}, String("foo"), false, false)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, false, field.minLengthValidation)
	assert.Equal(t, false, field.maxLengthValidation)
}
