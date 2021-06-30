package vjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchema_Validate(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		str := `{"age": 10}`
		schema := Schema{
			Fields: []Field{
				Integer("age").Required().Range(0, 20).Range(50, 60),
			},
		}

		err := schema.ValidateString(str)
		assert.Nil(t, err)

		err = schema.ValidateBytes([]byte(str))
		assert.Nil(t, err)
	})
	t.Run("float", func(t *testing.T) {
		str := `{"height": 10}`
		schema := Schema{
			Fields: []Field{
				Float("height").Required().Range(0.5, 20.3).Range(50, 60),
			},
		}

		err := schema.ValidateString(str)
		assert.Nil(t, err)

		err = schema.ValidateBytes([]byte(str))
		assert.Nil(t, err)
	})
	t.Run("string", func(t *testing.T) {
		str := `{"name": "foo"}`
		schema := Schema{
			Fields: []Field{
				String("name").Required(),
			},
		}

		err := schema.ValidateString(str)
		assert.Nil(t, err)

		err = schema.ValidateBytes([]byte(str))
		assert.Nil(t, err)
	})
	t.Run("array", func(t *testing.T) {
		str := `{"scores": [10,20,15,18]}`
		schema := Schema{
			Fields: []Field{
				Array("scores", Integer("score").Required().Range(0, 20)),
			},
		}

		err := schema.ValidateString(str)
		assert.Nil(t, err)

		err = schema.ValidateBytes([]byte(str))
		assert.Nil(t, err)
	})
	t.Run("boolean", func(t *testing.T) {
		str := `{"isOk": true}`
		schema := Schema{
			Fields: []Field{
				Boolean("isOk").Required(),
			},
		}

		err := schema.ValidateString(str)
		assert.Nil(t, err)

		err = schema.ValidateBytes([]byte(str))
		assert.Nil(t, err)
	})
	t.Run("nested", func(t *testing.T) {
		str := `{"scores": [10,20,15,18], "object":{"age":19}}`
		str2 := `{"scores": [10,20,15,18], "object":{"age":21}}`
		objSchema := Schema{
			Fields: []Field{
				Integer("age").Required().Range(0, 20).Range(50, 60),
			},
		}
		schema := Schema{
			Fields: []Field{
				Array("scores", Integer("score").Required().Range(0, 20)).Required(),
				Object("object", objSchema),
			},
		}

		err := schema.ValidateString(str)
		assert.Nil(t, err)

		err = schema.ValidateBytes([]byte(str))
		assert.Nil(t, err)

		err = schema.ValidateString(str2)
		assert.NotNil(t, err)

		err = schema.ValidateBytes([]byte(str2))
		assert.NotNil(t, err)
	})
	t.Run("invalid", func(t *testing.T) {
		str := "{{"
		schema := Schema{
			Fields: []Field{
				Array("scores", Integer("score").Required().Range(0, 20)).Required(),
			},
		}

		err := schema.ValidateString(str)
		assert.NotNil(t, err)

		err = schema.ValidateBytes([]byte(str))
		assert.NotNil(t, err)

	})
}

func TestReadFromFile(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			schema, err := ReadFromFile("test/string.json")
			assert.Nil(t, err)
			assert.Len(t, schema.Fields, 1)
			assert.Equal(t, true, schema.Fields[0].(*StringField).validateFormat)
			assert.Equal(t, true, schema.Fields[0].(*StringField).validateMaxLength)
			assert.Equal(t, true, schema.Fields[0].(*StringField).validateMinLength)
			assert.Equal(t, false, schema.Fields[0].(*StringField).validateChoices)
		})

		t.Run("invalid", func(t *testing.T) {
			schema, err := ReadFromFile("test/string_invalid.json")
			assert.NotNil(t, err)
			assert.Nil(t, schema)
		})
	})
	t.Run("integer", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			schema, err := ReadFromFile("test/integer.json")
			assert.Nil(t, err)
			assert.Len(t, schema.Fields, 1)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).required)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).minValidation)
			assert.Equal(t, 2, schema.Fields[0].(*IntegerField).min)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).maxValidation)
			assert.Equal(t, 100, schema.Fields[0].(*IntegerField).max)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).signValidation)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).positive)
			assert.Len(t, schema.Fields[0].(*IntegerField).ranges, 1)

		})

		t.Run("invalid", func(t *testing.T) {
			schema, err := ReadFromFile("test/integer_invalid.json")
			assert.NotNil(t, err)
			assert.Nil(t, schema)
		})
	})
	t.Run("float", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			schema, err := ReadFromFile("test/float.json")
			assert.Nil(t, err)
			assert.Len(t, schema.Fields, 1)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).required)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).minValidation)
			assert.Equal(t, 2.0, schema.Fields[0].(*FloatField).min)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).maxValidation)
			assert.Equal(t, 100.0, schema.Fields[0].(*FloatField).max)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).signValidation)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).positive)
			assert.Len(t, schema.Fields[0].(*FloatField).ranges, 1)

		})

		t.Run("invalid", func(t *testing.T) {
			schema, err := ReadFromFile("test/float_invalid.json")
			assert.NotNil(t, err)
			assert.Nil(t, schema)
		})
	})
	t.Run("array", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			schema, err := ReadFromFile("test/array.json")
			assert.Nil(t, err)
			assert.Len(t, schema.Fields, 1)
			assert.Equal(t, false, schema.Fields[0].(*ArrayField).required)
			assert.Equal(t, 2, schema.Fields[0].(*ArrayField).minLength)
			assert.Equal(t, true, schema.Fields[0].(*ArrayField).minLengthValidation)
			assert.Equal(t, 10, schema.Fields[0].(*ArrayField).maxLength)
			assert.Equal(t, true, schema.Fields[0].(*ArrayField).maxLengthValidation)
			assert.Equal(t, "age", schema.Fields[0].(*ArrayField).items.GetName())

		})

		t.Run("invalid", func(t *testing.T) {
			schema, err := ReadFromFile("test/array_invalid.json")
			assert.NotNil(t, err)
			assert.Nil(t, schema)
		})
	})
	t.Run("object", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			schema, err := ReadFromFile("test/object.json")
			assert.Nil(t, err)
			assert.Len(t, schema.Fields, 1)
			assert.Len(t, schema.Fields[0].(*ObjectField).schema.Fields, 1)
		})

		t.Run("invalid", func(t *testing.T) {
			schema, err := ReadFromFile("test/object_invalid.json")
			assert.NotNil(t, err)
			assert.Nil(t, schema)
		})
	})
	t.Run("boolean", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			schema, err := ReadFromFile("test/boolean.json")
			assert.Nil(t, err)
			assert.Len(t, schema.Fields, 1)
			assert.Equal(t, false, schema.Fields[0].(*BooleanField).required)
			assert.Equal(t, true, schema.Fields[0].(*BooleanField).valueValidation)
			assert.Equal(t, true, schema.Fields[0].(*BooleanField).value)
		})

		t.Run("invalid", func(t *testing.T) {
			schema, err := ReadFromFile("test/boolean_invalid.json")
			assert.NotNil(t, err)
			assert.Nil(t, schema)
		})
	})
	t.Run("null", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			schema, err := ReadFromFile("test/null.json")
			assert.Nil(t, err)
			assert.Len(t, schema.Fields, 1)
			assert.Equal(t, "foo", schema.Fields[0].(*NullField).name)
		})

		t.Run("invalid", func(t *testing.T) {
			schema, err := ReadFromFile("test/null_invalid.json")
			assert.NotNil(t, err)
			assert.Nil(t, schema)
		})
	})
	t.Run("invalid_type", func(t *testing.T) {
		schema, err := ReadFromFile("test/invalid_type.json")
		assert.NotNil(t, err)
		assert.Nil(t, schema)
	})
	t.Run("missing_type", func(t *testing.T) {
		schema, err := ReadFromFile("test/missing_type.json")
		assert.NotNil(t, err)
		assert.Nil(t, schema)
	})
	t.Run("wrong_type", func(t *testing.T) {
		schema, err := ReadFromFile("test/wrong_type.json")
		assert.NotNil(t, err)
		assert.Nil(t, schema)
	})
	t.Run("invalid_file", func(t *testing.T) {
		schema, err := ReadFromFile("test/not_existing.json")
		assert.NotNil(t, err)
		assert.Nil(t, schema)
	})
}

func TestReadFromString(t *testing.T) {
	schema, err := ReadFromString("{}")
	assert.Nil(t, err)
	assert.Len(t, schema.Fields, 0)
}

func TestNewSchema(t *testing.T) {
	s := NewSchema(
		Integer("foo"),
		String("bar"),
	)
	assert.Len(t, s.Fields, 2)
}

func TestSchema_UnmarshalJSON(t *testing.T) {
	var s Schema
	err := s.UnmarshalJSON([]byte("{{"))
	assert.NotNil(t, err)
}
