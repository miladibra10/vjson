package vjson

import (
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

			obj := struct{
				Age int `json:"age"`
			}{10}

			err := field.Validate(obj)
			assert.Nil(t, err)
		})
	})
}

func TestNewObject(t *testing.T) {
	s := Schema{}
	field := NewObject(ObjectFieldSpec{
		Name: "bar",
		Required: true,
	}, s)

	assert.NotNil(t, field)
	assert.Equal(t, "bar", field.name)
	assert.Equal(t, s, field.schema)
}