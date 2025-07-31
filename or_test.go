package vjson

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrField_GetName(t *testing.T) {
	field := Or("foo")
	assert.Equal(t, "foo", field.GetName())
}

func TestOrField_Validate(t *testing.T) {
	t.Run("no_fields_provided", func(t *testing.T) {
		field := Or("foo")
		err := field.Validate("test")
		assert.NotNil(t, err)
	})

	t.Run("not_required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Or("foo").WithFields(String("bar"), Boolean("baz"))
			err := field.Validate(nil)
			assert.Nil(t, err)
		})

		t.Run("valid_string_value", func(t *testing.T) {
			field := Or("foo").WithFields(String("bar"), Boolean("baz"))
			err := field.Validate("test")
			assert.Nil(t, err)
		})

		t.Run("valid_boolean_value", func(t *testing.T) {
			field := Or("foo").WithFields(String("bar"), Boolean("baz"))
			err := field.Validate(true)
			assert.Nil(t, err)
		})

		t.Run("invalid_value", func(t *testing.T) {
			field := Or("foo").WithFields(String("bar"), Boolean("baz"))
			err := field.Validate(123)
			assert.NotNil(t, err)
		})
	})

	t.Run("required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Or("foo").Required().WithFields(String("bar"), Boolean("baz"))
			err := field.Validate(nil)
			assert.NotNil(t, err)
		})

		t.Run("valid_string_value", func(t *testing.T) {
			field := Or("foo").Required().WithFields(String("bar"), Boolean("baz"))
			err := field.Validate("test")
			assert.Nil(t, err)
		})

		t.Run("valid_boolean_value", func(t *testing.T) {
			field := Or("foo").Required().WithFields(String("bar"), Boolean("baz"))
			err := field.Validate(true)
			assert.Nil(t, err)
		})

		t.Run("invalid_value", func(t *testing.T) {
			field := Or("foo").Required().WithFields(String("bar"), Boolean("baz"))
			err := field.Validate(123)
			assert.NotNil(t, err)
		})
	})
}

func TestOrField_MarshalJSON(t *testing.T) {
	field := Or("foo").WithFields(String("bar"), Boolean("baz"))

	b, err := json.Marshal(field)
	assert.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(b, &data)
	assert.Nil(t, err)

	assert.Equal(t, "foo", data["name"])
	assert.Equal(t, "or", data["type"])

	// Check that the fields property exists and is an array
	fields, ok := data["fields"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, fields, 2)
}

func TestNewOr(t *testing.T) {
	// Create a simple field spec with fields
	fieldSpec := OrFieldSpec{
		Name:     "bar",
		Required: true,
		Fields: []map[string]interface{}{
			{
				"name": "string_field",
				"type": "string",
			},
			{
				"name": "boolean_field",
				"type": "boolean",
			},
		},
	}

	field := NewOr(fieldSpec)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, true, field.required)

	// The fields should be initialized but empty since we don't process them in NewOr
	// They are processed in getOrField in schema.go
	assert.Empty(t, field.fields)
}
