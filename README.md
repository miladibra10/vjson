# vjson
[![codecov](https://codecov.io/gh/miladibra10/vjson/branch/main/graph/badge.svg)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/miladibra10/vjson)](https://goreportcard.com/report/github.com/miladibra10/vjson)
[![<miladibra10>](https://circleci.com/gh/miladibra10/vjson.svg?style=svg)](#)
[![Go Reference](https://pkg.go.dev/badge/github.com/miladibra10/vjson.svg)](https://pkg.go.dev/github.com/miladibra10/vjson)

**vjson**  is a Go package that helps to validate JSON objects in a declarative way.

# Getting Started

## Installing

For installing vjson, use command below:

```
go get -u github.com/miladibra10/vjson
```

# Concepts

There are two main concepts in `vjson` that are:

+ Schema
+ Field

## Schema

A schema is the holder of JSON object specifications. It contains the way of validation of a JSON object. a schema
consists of an array of fields.

## Field

A field contains characteristics of a specific field of a json object. multiple field types are supported by `vjson`.
list of supported types are:

+ `integer`
+ `float`
+ `string`
+ `boolean`
+ `array`
+ `object`
+ `null`

# How to create a Schema

There are two ways to create a schema.

+ Schema could be declared in code in a declarative manner.
+ Schema could be parsed from a file or string.

## Schema in code

Schema could be declared in code like this:

```go
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
```

`schema` object contains a string field, named `name`. This code validates `jsonString`.

## Parse Schema
Schema could be parsed from a file or a string. These methods help to parse schema.

+ [ReadFromString(input string)](#parse-schema): receives characteristics of a schema in a string format and returns a schema object and an error.
+ [ReadFromBytes(input []byte)](#parse-schema): receives characteristics of a schema in a byte array format and returns a schema object and an error.
+ [ReadFromFile(filePath string)](#parse-schema): receives file path of json file which contains characteristics of a schema, and returns a schema object and an error.

Format of schema for parsing should be a json like this:
```json
{
  "fields": [
    ...
  ]
}
```
`fields` should contain [field specifications](#fields). 

This code parses a schema from string:

```go
package main

import "github.com/miladibra10/vjson"

func main() {
	schemaStr := `
	{
    "strict": true,
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
```

`schemaStr` describes the schema and `vjson.ReadFromString(schemaStr)` parses the string and returns a schema object.

`schema` object contains a string field, named `name`. This code validates `jsonString`.

> **Note**: You could Marshal your schema as a json object for backup usages with `json.Marshal` function.



# Fields

## Integer
An integer field could be created in code like this:
```go
vjson.Integer("foo")
```
some validation characteristics could be added to an integer field with chaining some functions:

+ [Required()](#integer) sets the field as a required field. validation will return an error if a required field is not present in json object.
+ [Min(min int)](#integer) forces the integer field to be greater than `min` in validating json object.
+ [Max(max int)](#integer) forces the integer field to be lower than `max` in validating json object.
+ [Positive()](#integer) checks if the value of field is positive.
+ [Negative()](#integer) checks if the value of field is negative.
+ [Range(start, end int)](#integer) adds a range for integer field. the value of json field should be within this range.

integer field could be described by a json for schema parsing.
+ **`name`**: the name of the field
+ **`type`**: type value for integer field must be `integer`
+ `required`: whether the field is required or not
+ `min`: minimum value of field
+ `max`: maximum value of field
+ `positive`: a boolean that describes that a field is positive or negative (`true` for positive and `false` for negative)
+ `ranges`: an array of ranges to be checked in field validation.

### Example
an integer field, named `foo` which is required, minimum value should be 2, maximum value should be 10, should be positive and be within range [3,5] or [6,8] ,could be declared like this:

#### Code
```go
vjson.Integer("foo").Required().Min(2).Max(10).Positive().Range(3,5).Range(6,8)
```

#### File
```json
{
  "name": "foo",
  "type": "integer",
  "required": true,
  "min": 2,
  "max": 10,
  "positive": true,
  "ranges": [
    {
      "start": 3,
      "end": 5
    },
    {
      "start": 6,
      "end": 8
    }
  ]
}
```

## Float
A float field could be created in code like this:
```go
vjson.Float("foo")
```
some validation characteristics could be added to a float field with chaining some functions:

+ [Required()](#float) sets the field as a required field. validation will return an error if a required field is not present in json object.
+ [Min(min float64)](#float) forces the float field to be greater than `min` in validating json object.
+ [Max(max float64)](#float) forces the float field to be lower than `max` in validating json object.
+ [Positive()](#float) checks if the value of field is positive.
+ [Negative()](#float) checks if the value of field is negative.
+ [Range(start, end float64)](#float) adds a range for float field. the value of json field should be within this range.

float field could be described by a json for schema parsing.
+ **`name`**: the name of the field
+ **`type`**: type value for float field must be `float`
+ `required`: whether the field is required or not
+ `min`: minimum value of field
+ `max`: maximum value of field
+ `positive`: a boolean that describes that a field is positive or negative (`true` for positive and `false` for negative)
+ `ranges`: an array of ranges to be checked in field validation.

### Example
a float field, named `foo` which is required, minimum value should be 2.5, maximum value should be 10.5, should be positive and be within range [3,5] or [6,8] ,could be declared like this:

#### Code
```go
vjson.Float("foo").Required().Min(2.5).Max(10.5).Positive().Range(3,5).Range(6,8)
```

#### File
```json
{
  "name": "foo",
  "type": "float",
  "required": true,
  "min": 2.5,
  "max": 10.5,
  "positive": true,
  "ranges": [
    {
      "start": 3,
      "end": 5
    },
    {
      "start": 6,
      "end": 8
    }
  ]
}
```

## String
A string field could be created in code like this:
```go
vjson.String("foo")
```
some validation characteristics could be added to a string field with chaining some functions:

+ [Required()](#string) sets the field as a required field. validation will return an error if a required field is not present in json object.
+ [MinLength(min int)](#string) forces the length of string field to be greater than `min` in validating json object.
+ [MaxLength(max int)](#string) forces the length of string field to be lower than `max` in validating json object.
+ [Format(format string)](#string) gets a `regex` format and checks if value of json object matches the format.
+ [Choices(choice ...string)](#string) checks if the value of string field is equal to one of choices.

float field could be described by a json for schema parsing.
+ **`name`**: the name of the field
+ **`type`**: type value for string field must be `string`
+ `required`: whether the field is required or not
+ `min_length`: minimum length of string value of field
+ `max_length`: maximum length of string value of field
+ `format`: a `regex` format and checks if value of json object matches the format.
+ `choices`: a list of strings that value of field should be equal to one of them.

### Example
a string field, named `foo` which is required, minimum length value should be 2, maximum length value should be 10, should be Equal to one of these values: `first`, `second` could be declared like this:

#### Code
```go
vjson.String("foo").Required().MinLength(2).MaxLength(10).Choices("first", "second")
```

#### File
```json
{
  "name": "foo",
  "type": "string",
  "required": true,
  "min_length": 2,
  "max_length": 10,
  "choices": [
    "first",
    "second"
  ]
}
```

## Boolean
A boolean field could be created in code like this:
```go
vjson.Boolean("foo")
```
some validation characteristics could be added to a boolean field with chaining some functions:

+ [Required()](#boolean) sets the field as a required field. validation will return an error if a required field is not present in json object.
+ [ShouldBe(value bool)](#boolean) forces the value of field be equal to `value`

boolean field could be described by a json for schema parsing.
+ **`name`**: the name of the field
+ **`type`**: type value for boolean field must be `boolean`
+ `required`: whether the field is required or not
+ `value`: a boolean (same sa `ShouldBe` in code) that describes that the value of json field.

### Example
a boolean field, named `foo` which is required, and always should be false, could be declared like this:

#### Code
```go
vjson.Boolean("foo").Required().ShouldBe(false)
```

#### File
```json
{
  "name": "foo",
  "type": "boolean",
  "required": true,
  "value": false
}
```

## Array
An array field could be created in code like this:
```go
vjson.Array("foo", vjson.String("item"))
```
the first argument is the name of array field, and the second one is the field characteristics of each item of array.


some validation characteristics could be added to an array field with chaining some functions:

+ [Required()](#array) sets the field as a required field. validation will return an error if a required field is not present in json object.
+ [MinLength(min int)](#array) forces the length of array field to be greater than `min` in validating json object.
+ [MaxLength(max int)](#array) forces the length of array field to be lower than `max` in validating json object.

array field could be described by a json for schema parsing.
+ **`name`**: the name of the field
+ **`type`**: type value for array field must be `array`
+ `required`: whether the field is required or not
+ `min_length`: minimum length of array
+ `max_length`: maximum length of array
+ `items`: specifications of item fields. could be any field.

### Example
an array field, named `foo` with integer items between [0,20] range, which is required, and its length should be at least 2 and at last 10, could be declared like this:

#### Code
```go
vjson.Array("foo", vjson.Integer("item").Range(0,20)).Required().MinLength(2).MaxLength(10)
```

#### File
```json
{
  "name": "foo",
  "type": "array",
  "required": true,
  "min_length": 2,
  "max_length": 10,
  "items": {
    "name": "item",
    "type": "integer",
    "ranges": [
      {
        "start": 0,
        "end": 20
      }
    ]
  }
}
```

### Fixed Length Array

Each item has a different type in the array.

#### Code
```go
vjson.FixArray("foo", []Field{
  vjson.String("item"),
  vjson.Integer("item2")
})
```

#### File
```json
{
  "name": "foo",
  "type": "array",
  "required": true,
  "fix_items": [{
    "name": "item",
    "type": "string",
  }, {
    "name": "item2",
    "type": "integer",
  }]
}
```

## Object
An object field could be created in code like this:
```go
vjson.Object("foo", vjson.NewSchema(
        /// Fields	
	))
```
the first argument is the name of object field, and the second one is the schema of object value. this feature helps validation of nested json objects


some validation characteristics could be added to an array field with chaining some functions:

+ [Required()](#object) sets the field as a required field. validation will return an error if a required field is not present in json object.
+ [Strict()](#object)
When set strict mode is true, all fields in json object must validate. Default strict mode is `false`.


object field could be described by a json for schema parsing.
+ **`name`**: the name of the field
+ **`type`**: type value for object field must be `object`
+ `required`: whether the field is required or not
+ `schema`: schema of object value.

### Example
a required object field, named `foo` which its valid value is an object with `name` and `last_name` required strings, could be declared like this:

#### Code
```go
vjson.Object("foo", vjson.NewSchema(
	vjson.String("name").Required(),
	vjson.String("last_name").Required(),
	)).Required().Strict()
```

#### File
```json
{
  "name": "foo",
  "type": "object",
  "required": true,
  "schema": {
    "strict": true,
    "fields": [
      {
        "name": "name",
        "type": "string",
        "required": true
      },
      {
        "name": "last_name",
        "type": "string",
        "required": true
      }
    ]
  }
}
```

## Null
A null field (a field that its value should be null!) could be created in code like this:
```go
vjson.Null("foo")
```

null field could be described by a json for schema parsing.
+ **`name`**: the name of the field
+ **`type`**: type value for null field must be `null`

### Example
a null field, named `foo`, could be declared like this:

#### Code
```go
vjson.Null("foo")
```

#### File
```json
{
  "name": "foo",
  "type": "null"
}
```

# Validation
After creating a schema, you can validate your json objects with these methods:

+ [ValidateBytes(input []byte)](#validation): receives a byte array as a json input and validates it. this method returns an error. it would be `nil` if the object is valid, and it will return an error if the input object is not valid.
+ [ValidateString(input string)](#validation): acts like `ValidateBytes` but its argument is string.

# Example
This code validates an object that should have `name` and `age` fields.
```go
package main

import "github.com/miladibra10/vjson"

func main() {
	schema := vjson.NewSchema(
		vjson.String("name").Required(),
		vjson.Integer("age").Positive(),
	)

	jsonString := `
	{
		
  "required": true,
  "value": false"name": "James"
	}
	`

	err := schema.ValidateString(jsonString)
	if err != nil {
		panic(err) // Will not panic
	}

	jsonString = `
	{
		"name": "James",
        "age": 10
	}
	`

	err = schema.ValidateString(jsonString)
	if err != nil {
		panic(err) // Will not panic
	}

	jsonString = `
	{
        "age": 10
	}
	`

	err = schema.ValidateString(jsonString)
	if err != nil {
		panic(err) // Will panic because name field is missing in jsonString
	}
	
}
```

# Benchmarks

Results of benchmarking validation functions highly depends on types and number of fields.

two simple benchmarks (exists in `schema_test.go` file) with using all features of `vjson` gives this result:

```
goos: linux
goarch: amd64
pkg: github.com/miladibra10/vjson
BenchmarkSchema_ValidateString-8          416664              2792 ns/op
BenchmarkSchema_ValidateBytes-8           431734              2858 ns/op
PASS
ok      github.com/miladibra10/vjson    2.461s
```