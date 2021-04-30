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
				Integer("age").Required().Range(0,20).Range(50,60),
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
				Array("scores", Integer("score").Required().Range(0,20)),
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
				Integer("age").Required().Range(0,20).Range(50,60),
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

