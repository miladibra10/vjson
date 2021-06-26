package vjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullField_GetName(t *testing.T) {
	field := Null("foo")
	assert.Equal(t, "foo", field.GetName())
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

