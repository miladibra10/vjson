package vjson

type fieldType string

const (
	integerType fieldType = "integer"
	floatType   fieldType = "float"
	stringType  fieldType = "string"
	arrayType   fieldType = "array"
	booleanType fieldType = "boolean"
	objectType  fieldType = "object"
	nullType    fieldType = "null"
	orType      fieldType = "or"
)

const typeKey = "type"
