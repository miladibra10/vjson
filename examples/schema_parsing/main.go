package main

import (
	"github.com/miladibra10/vjson"
)

func main() {
	schemaStr := `
	{
		"strict": true,
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
			},
			{
				"name": "person",
				"type": "object",
				"required": true,
				"schema": {
					"strict": true,
					"fields": [
					  {
						"name": "name",
						"type": "object",
						"required": true,
						"schema": {
							"strict": true,
							"fields": [{
								"name": "first",
								"type": "string",
								"required": true
							},
							{
								"name": "last",
								"type": "string",
								"required": true
							}]
						}
					  },
					  {
						"name": "gender",
						"type": "string",
						"required": true
					  }
					]
				}
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
		"name": ["hello", 123],
		"person": {
			"name": {
				"first": "asg",
				"last": "4234"
			},
			"gender": "male"
		}
	}
	`

	err = schema.ValidateString(jsonString)
	if err != nil {
		panic(err)
	}
}
