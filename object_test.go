package vjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectField_GetName(t *testing.T) {
	field := Object("foo", Schema{})
	assert.Equal(t, "foo", field.GetName())
}

func TestObjectField_Validate(t *testing.T) {
	objSchema := Schema{
		Fields: []Field{
			Integer("age").Min(0).Max(90).Required(),
		},
	}
	t.Run("invalid_input", func(t *testing.T) {
		field := Object("foo", objSchema)

		err := field.Validate(1)
		assert.NotNil(t, err)
	})
	t.Run("not_required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Object("foo", objSchema)

			err := field.Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Object("foo", objSchema)

			err := field.Validate(`{"age":10}`)
			assert.Nil(t, err)
		})
	})
	t.Run("required_field", func(t *testing.T) {
		t.Run("nil_value", func(t *testing.T) {
			field := Object("foo", objSchema).Required()

			err := field.Validate(nil)
			assert.NotNil(t, err)
		})
		t.Run("valid_value", func(t *testing.T) {
			field := Object("foo", objSchema)

			err := field.Validate(`{"age":10}`)
			assert.Nil(t, err)
		})
		t.Run("valid_struct_value", func(t *testing.T) {
			field := Object("foo", objSchema)

			obj := struct {
				Age int `json:"age"`
			}{10}

			err := field.Validate(obj)
			assert.Nil(t, err)
		})
	})

	t.Run("strict_field", func(t *testing.T) {
		t.Run("strict_off", func(t *testing.T) {
			field := Object("foo", objSchema).Required()

			err := field.Validate(`{"age":10, "name": "john"}`)
			assert.Nil(t, err)
		})

		t.Run("strict_on", func(t *testing.T) {
			field := Object("foo", objSchema).Required().Strict()

			err := field.Validate(`{"age":10, "name": "john"}`)
			assert.NotNil(t, err)
		})

		t.Run("strict_off_and_not_required", func(t *testing.T) {
			field := Object("foo", objSchema)

			err := field.Validate(`{"age":10, "name": "john"}`)
			assert.Nil(t, err)
		})

		t.Run("strict_on_and_not_required", func(t *testing.T) {
			field := Object("foo", objSchema).Strict()

			err := field.Validate(`{"age":10, "name": "john"}`)
			assert.NotNil(t, err)
		})

		t.Run("strict_on_and_not_required_for_nil", func(t *testing.T) {
			field := Object("foo", objSchema).Strict()

			err := field.Validate(nil)
			assert.Nil(t, err)
		})

		t.Run("strict_on_struct", func(t *testing.T) {
			objSchemaStrict := Schema{
				Fields: []Field{
					Integer("age").Min(0).Max(90).Required(),
				},
			}
			objSchemaStrict.StrictMode = true
			field := Object("foo", objSchemaStrict)

			err := field.Validate(`{"age":10, "name": "john"}`)
			assert.NotNil(t, err)
		})
	})
}

func TestObjectField_MarshalJSON(t *testing.T) {
	field := Object("foo", NewSchema(String("bar")))

	b, err := json.Marshal(field)
	assert.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(b, &data)
	assert.Nil(t, err)

	assert.Equal(t, "foo", data["name"])
	assert.Equal(t, string(objectType), data["type"])
	assert.Equal(t, "bar", data["schema"].(map[string]interface{})["fields"].([]interface{})[0].(map[string]interface{})["name"])
}

func TestNewObject(t *testing.T) {
	s := Schema{}
	field := NewObject(ObjectFieldSpec{
		Name:     "bar",
		Required: true,
	}, s)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, s, field.schema)
}
