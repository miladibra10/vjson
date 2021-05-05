package main

import "github.com/miladibra10/vjson"

func main() {
	schema := vjson.NewSchema(
		vjson.String("name").Required(),
		)

	jsonString := `
	{
		"name": "James"
	}
	`

	err := schema.ValidateString(jsonString)
	if err != nil {
		panic(err)
	}
}
