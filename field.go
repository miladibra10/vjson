package go_json_validator

type Field interface {
	Validate(interface{}) error
}
