package vjson

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullField_GetName(t *testing.T) {
	field := Null("foo")
	assert.Equal(t, "foo", field.GetName())
}

func TestNullField_MarshalJSON(t *testing.T) {
	field := Null("foo")

	b, err := json.Marshal(field)
	assert.Nil(t, err)

	data := map[string]string{}
	err = json.Unmarshal(b, &data)
	assert.Nil(t, err)

	assert.Equal(t, "foo", data["name"])
	assert.Equal(t, string(nullType), data["type"])
}

func TestNullField_Validate(t *testing.T) {
	t.Run("invalid_input", func(t *testing.T) {
		field := Null("foo")

		err := field.Validate(1)
		assert.NotNil(t, err)
	})
	t.Run("valid_input", func(t *testing.T) {
		field := Null("foo")

		err := field.Validate(nil)
		assert.Nil(t, err)
	})
}
