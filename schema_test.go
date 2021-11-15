package vjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, true, schema.Fields[0].(*StringField).FieldValidateFormat)
			assert.Equal(t, true, schema.Fields[0].(*StringField).FieldValidateMaxLength)
			assert.Equal(t, true, schema.Fields[0].(*StringField).FieldValidateMinLength)
			assert.Equal(t, false, schema.Fields[0].(*StringField).FieldValidateChoices)
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
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).FieldRequired)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).FieldMinValidation)
			assert.Equal(t, 2, schema.Fields[0].(*IntegerField).FieldMin)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).FieldMaxValidation)
			assert.Equal(t, 100, schema.Fields[0].(*IntegerField).FieldMax)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).FieldSignValidation)
			assert.Equal(t, true, schema.Fields[0].(*IntegerField).FieldPositive)
			assert.Len(t, schema.Fields[0].(*IntegerField).FieldRanges, 1)

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
			assert.Equal(t, true, schema.Fields[0].(*FloatField).FieldRequired)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).FieldMinValidation)
			assert.Equal(t, 2.0, schema.Fields[0].(*FloatField).FieldMin)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).FieldMaxValidation)
			assert.Equal(t, 100.0, schema.Fields[0].(*FloatField).FieldMax)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).FieldSignValidation)
			assert.Equal(t, true, schema.Fields[0].(*FloatField).FieldPositive)
			assert.Len(t, schema.Fields[0].(*FloatField).FieldRanges, 1)

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
			assert.Equal(t, false, schema.Fields[0].(*ArrayField).FieldRequred)
			assert.Equal(t, 2, schema.Fields[0].(*ArrayField).FieldMinLength)
			assert.Equal(t, true, schema.Fields[0].(*ArrayField).FieldMinLengthValidation)
			assert.Equal(t, 10, schema.Fields[0].(*ArrayField).FieldMaxLength)
			assert.Equal(t, true, schema.Fields[0].(*ArrayField).FieldMaxLengthValidation)
			assert.Equal(t, "age", schema.Fields[0].(*ArrayField).Items.GetName())

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
			assert.Len(t, schema.Fields[0].(*ObjectField).FieldSchema.Fields, 1)
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
			assert.Equal(t, false, schema.Fields[0].(*BooleanField).FieldRequired)
			assert.Equal(t, true, schema.Fields[0].(*BooleanField).FieldValueValidation)
			assert.Equal(t, true, schema.Fields[0].(*BooleanField).Value)
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
			assert.Equal(t, "foo", schema.Fields[0].(*NullField).Name)
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

func BenchmarkSchema_ValidateString(b *testing.B) {
	s := NewSchema(
		String("first_name").Required().MinLength(2),
		String("last_name").Required().MinLength(2),
		Integer("age").Positive(),
		Array("cars", String("car").Choices("pride", "peugeot", "audi", "ford")),
		Float("avg").Range(0, 1),
		Boolean("is_active").ShouldBe(true),
		Object("test", NewSchema(
			String("field").Required(),
		)).Required(),
		Null("null"),
	)
	jsonStr := `{"first_name": "abbas", "last_name": "booazar", "age": 30, "cars":["pride", "audi"], "avg": 0.97, "is_active": true, "test":{"field": "yes"}}`
	for i := 0; i < b.N; i++ {
		_ = s.ValidateString(jsonStr)
	}
}

func BenchmarkSchema_ValidateBytes(b *testing.B) {
	s := NewSchema(
		String("first_name").Required().MinLength(2),
		String("last_name").Required().MinLength(2),
		Integer("age").Positive(),
		Array("cars", String("car").Choices("pride", "peugeot", "audi", "ford")),
		Float("avg").Range(0, 1),
		Boolean("is_active").ShouldBe(true),
		Object("test", NewSchema(
			String("field").Required(),
		)).Required(),
		Null("null"),
	)
	jsonBytes := []byte(`{"first_name": "abbas", "last_name": "booazar", "age": 30, "cars":["pride", "audi"], "avg": 0.97, "is_active": true, "test":{"field": "yes"}}`)
	for i := 0; i < b.N; i++ {
		_ = s.ValidateBytes(jsonBytes)
	}
}
