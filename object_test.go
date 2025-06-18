package vjson

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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

	t.Run("invalid_json", func(t *testing.T) {
		field := Object("foo", objSchema)

		err := field.Validate(`{"age":10`)  // Missing closing brace
		assert.NotNil(t, err)

		err = field.Validate(`{age:10}`)  // Missing quotes around field name
		assert.NotNil(t, err)
	})

	t.Run("missing_required_fields", func(t *testing.T) {
		field := Object("foo", objSchema)

		err := field.Validate(`{}`)  // Missing required age field
		assert.NotNil(t, err)

		err = field.Validate(`{"name":"John"}`)  // Missing required age field, has other fields
		assert.NotNil(t, err)
	})

	t.Run("invalid_field_values", func(t *testing.T) {
		field := Object("foo", objSchema)

		err := field.Validate(`{"age":-10}`)  // Age is negative, should be 0-90
		assert.NotNil(t, err)

		err = field.Validate(`{"age":100}`)  // Age is > 90, should be 0-90
		assert.NotNil(t, err)

		err = field.Validate(`{"age":"ten"}`)  // Age is string, should be integer
		assert.NotNil(t, err)
	})

	t.Run("multiple_fields", func(t *testing.T) {
		multiFieldSchema := Schema{
			Fields: []Field{
				Integer("age").Min(0).Max(90).Required(),
				String("name").MinLength(2).Required(),
				Boolean("active").ShouldBe(true),
			},
		}
		field := Object("foo", multiFieldSchema)

		// All required fields present and valid
		err := field.Validate(`{"age":30, "name":"John", "active":true}`)
		assert.Nil(t, err)

		// Missing one required field
		err = field.Validate(`{"age":30, "active":true}`)
		assert.NotNil(t, err)

		// One field with invalid value
		err = field.Validate(`{"age":30, "name":"John", "active":false}`)
		assert.NotNil(t, err)

		// Multiple invalid fields
		err = field.Validate(`{"age":100, "name":"J", "active":false}`)
		assert.NotNil(t, err)
	})

	t.Run("nested_objects", func(t *testing.T) {
		addressSchema := Schema{
			Fields: []Field{
				String("street").Required(),
				Integer("zipcode").Required(),
			},
		}

		personSchema := Schema{
			Fields: []Field{
				String("name").Required(),
				Integer("age").Required(),
				Object("address", addressSchema).Required(),
			},
		}

		field := Object("person", personSchema)

		// Valid nested object
		err := field.Validate(`{
			"name": "John",
			"age": 30,
			"address": {
				"street": "Main St",
				"zipcode": 12345
			}
		}`)
		assert.Nil(t, err)

		// Invalid nested object (missing required field)
		err = field.Validate(`{
			"name": "John",
			"age": 30,
			"address": {
				"street": "Main St"
			}
		}`)
		assert.NotNil(t, err)

		// Invalid nested object (wrong type)
		err = field.Validate(`{
			"name": "John",
			"age": 30,
			"address": "Main St"
		}`)
		assert.NotNil(t, err)

		// Multiple levels of nesting
		companySchema := Schema{
			Fields: []Field{
				String("name").Required(),
				Object("ceo", personSchema).Required(),
			},
		}

		companyField := Object("company", companySchema)

		// Valid multi-level nested object
		err = companyField.Validate(`{
			"name": "Acme Inc",
			"ceo": {
				"name": "John",
				"age": 45,
				"address": {
					"street": "Executive Ave",
					"zipcode": 54321
				}
			}
		}`)
		assert.Nil(t, err)

		// Invalid multi-level nested object
		err = companyField.Validate(`{
			"name": "Acme Inc",
			"ceo": {
				"name": "John",
				"age": 45,
				"address": {
					"zipcode": 54321
				}
			}
		}`)
		assert.NotNil(t, err)
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
