package go_json_validator

type Schema struct {
	Fields []Field `json:"fields"`
}

func (s *Schema) ValidateBytes(input []byte) error {
	return nil
}


func (s *Schema) ValidateString(input string) error {
	return nil
}

func ReadFromString(input string) (*Schema, error) {
	return nil, nil
}

func ReadFromBytes(input []byte) (*Schema, error) {
	return nil, nil
}

func ReadFromFile(filePath string) (*Schema, error) {
	return nil, nil
}