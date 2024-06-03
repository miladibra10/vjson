package main

import "github.com/thanharrow/vjson"

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
