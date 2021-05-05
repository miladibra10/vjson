package vjson

type FieldType string

const (
	IntegerType FieldType = "integer"
	FloatType   FieldType = "float"
	StringType  FieldType = "string"
	ArrayType   FieldType = "array"
	BooleanType FieldType = "boolean"
	ObjectType  FieldType = "object"
)

const TypeKey = "type"
