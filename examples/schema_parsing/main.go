package main

import (
	"fmt"

	"github.com/thanharrow/vjson"
)

func main() {
	schemaStr := `
	{
		"fields": [
			{
				"name": "name",
				"type": "string",
				"required": true
			}
		]
	}
	`
	schema, err := vjson.ReadFromString(schemaStr)
	if err != nil {
		panic(err)
	}

	schema.SetStrict(false) // default false

	// field age is not validate, so validate is false
	jsonString := `
	{
		"name": "James",
    "age": 10
	}
	`

	err = schema.ValidateString(jsonString)
	if err != nil {
		fmt.Printf("schema is not valid: " + err.Error())
	}
}
