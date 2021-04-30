package vjson

type Field interface {
	GetName() string
	Validate(interface{}) error
}
