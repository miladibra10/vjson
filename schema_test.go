package vjson

import (
	"encoding/json"
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
	schema, err := ReadFromString(`{"fields":[{"name":"bar","type": "string","required":true}]}`)
	assert.Nil(t, err)
	assert.Len(t, schema.Fields, 1)
}

func TestNewSchema(t *testing.T) {
	s := NewSchema(
		Integer("foo"),
		String("bar"),
	)
	assert.Len(t, s.Fields, 2)
}

func TestSchema_MarshalJSON(t *testing.T) {
	schema := NewSchema(
		Integer("foo"),
		String("bar").Required(),
	)
	schemaBytes, _ := json.Marshal(schema)

	var newSchema Schema

	err := json.Unmarshal(schemaBytes, &newSchema)
	assert.Nil(t, err)

	assert.Equal(t, schema.Fields[0], newSchema.Fields[0])
	assert.Equal(t, len(schema.Fields), len(newSchema.Fields))

}

func TestSchema_UnmarshalJSON(t *testing.T) {
	var s Schema
	err := s.UnmarshalJSON([]byte("{{"))
	assert.NotNil(t, err)
}

func TestSchema_MultipleFieldsSameType(t *testing.T) {
	schema := Schema{
		Fields: []Field{
			String("first_name").Required().MinLength(2),
			String("last_name").Required().MinLength(2),
			String("email").Required().Format("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"),
		},
	}

	// Valid input with all string fields
	validJSON := `{
		"first_name": "John",
		"last_name": "Doe",
		"email": "john.doe@example.com"
	}`
	err := schema.ValidateString(validJSON)
	assert.Nil(t, err)

	// Invalid first_name (too short)
	invalidFirstName := `{
		"first_name": "J",
		"last_name": "Doe",
		"email": "john.doe@example.com"
	}`
	err = schema.ValidateString(invalidFirstName)
	assert.NotNil(t, err)

	// Invalid email format
	invalidEmail := `{
		"first_name": "John",
		"last_name": "Doe",
		"email": "not-an-email"
	}`
	err = schema.ValidateString(invalidEmail)
	assert.NotNil(t, err)
}

func TestSchema_FieldsWithSameName(t *testing.T) {
	// This is an edge case - fields with the same name
	// The actual behavior is that all fields with the same name are validated
	schema := Schema{
		Fields: []Field{
			String("name").Required().MinLength(5),
			Integer("name").Required().Min(10),
		},
	}

	// Valid for integer field, invalid for string field
	validForInteger := `{"name": 15}`
	err := schema.ValidateString(validForInteger)
	assert.NotNil(t, err) // Fails because 15 is not a string

	// Valid for string field, invalid for integer field
	validForString := `{"name": "John Doe"}`
	err = schema.ValidateString(validForString)
	assert.NotNil(t, err) // Fails because "John Doe" is not an integer

	// Invalid for both fields
	invalidForBoth := `{"name": true}`
	err = schema.ValidateString(invalidForBoth)
	assert.NotNil(t, err)

	// Let's create a schema with fields that don't conflict
	nonConflictingSchema := Schema{
		Fields: []Field{
			String("name").Required(),
			Integer("age").Required().Min(10),
		},
	}

	// Valid for both fields
	validForBoth := `{"name": "John Doe", "age": 15}`
	err = nonConflictingSchema.ValidateString(validForBoth)
	assert.Nil(t, err)
}

func TestSchema_ValidateMap(t *testing.T) {
	schema := Schema{
		Fields: []Field{
			String("name").Required(),
			Integer("age").Required().Min(18),
		},
	}

	// Valid map
	validMap := map[string]interface{}{
		"name": "John Doe",
		"age":  25,
	}
	validJSON, err := json.Marshal(validMap)
	assert.Nil(t, err)
	err = schema.ValidateBytes(validJSON)
	assert.Nil(t, err)

	// Invalid map (missing required field)
	invalidMap1 := map[string]interface{}{
		"name": "John Doe",
	}
	invalidJSON1, err := json.Marshal(invalidMap1)
	assert.Nil(t, err)
	err = schema.ValidateBytes(invalidJSON1)
	assert.NotNil(t, err)

	// Invalid map (wrong type)
	invalidMap2 := map[string]interface{}{
		"name": "John Doe",
		"age":  "25", // String instead of integer
	}
	invalidJSON2, err := json.Marshal(invalidMap2)
	assert.Nil(t, err)
	err = schema.ValidateBytes(invalidJSON2)
	assert.NotNil(t, err)

	// Invalid map (below minimum)
	invalidMap3 := map[string]interface{}{
		"name": "John Doe",
		"age":  15,
	}
	invalidJSON3, err := json.Marshal(invalidMap3)
	assert.Nil(t, err)
	err = schema.ValidateBytes(invalidJSON3)
	assert.NotNil(t, err)
}

func TestSchema_EmptySchema(t *testing.T) {
	emptySchema := Schema{
		Fields: []Field{},
	}

	// Any JSON should be valid for an empty schema
	err := emptySchema.ValidateString(`{"name": "John", "age": 30}`)
	assert.Nil(t, err)

	err = emptySchema.ValidateString(`{}`)
	assert.Nil(t, err)

	err = emptySchema.ValidateString(`[]`)
	assert.Nil(t, err)

	// Even invalid JSON should fail
	err = emptySchema.ValidateString(`{invalid}`)
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
