package schema

import (
	"encoding/json"
	"os"
)

type JsonSchemaLoader struct{}

func (s *JsonSchemaLoader) LoadSchema(b []byte) (UnparsedSchema, error) {
	var parsedJson UnparsedSchema

	err := json.Unmarshal(b, &parsedJson)
	if err != nil {
		return nil, err
	}

	return parsedJson, nil
}

func (s *JsonSchemaLoader) LoadSchemaFromFile(filepath string) (UnparsedSchema, error) {
	jsonFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return s.LoadSchema(jsonFile)
}
