package vjson

import "encoding/json"

// Field is the abstraction on a field in a json.
// different field types can be implemented with implementing this interface.
type Field interface {
	json.Marshaler
	GetName() string
	Validate(interface{}) error
}
