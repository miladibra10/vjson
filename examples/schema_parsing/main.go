package main

import (
	"github.com/miladibra10/vjson"
)

func main() {
	schemaStr := `
	{
		"fields": [
			{
				"name": "name",
				"type": "array",
				"required": true,
				"fix_items": [{
					"name": "11",
					"type": "string"
				}, {
					"name": "12",
					"type": "integer"
				}]
			}
		]
	}
	`
	schema, err := vjson.ReadFromString(schemaStr)
	if err != nil {
		panic(err)
	}

	jsonString := `
	{
		"name": ["hello", 123]
	}
	`

	err = schema.ValidateString(jsonString)
	if err != nil {
		panic(err)
	}
}
