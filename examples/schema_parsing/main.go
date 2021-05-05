package main

import "github.com/miladibra10/vjson"

func main() {
	schemaStr := `
	{
		"fields": [
			{
				"name": "name",
				"type": "string"
				"required": true
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
		"name": "James"
	}
	`

	err = schema.ValidateString(jsonString)
	if err != nil {
		panic(err)
	}
}
