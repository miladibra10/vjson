package vjson

type Field interface {
	Validate(interface{}) error
}
