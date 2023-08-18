package schema

import (
	"encoding/json"
	"os"
)

/*
	⚠️ The schema must always be loaded from a string in JSON format, in this way we ensure
	that it is always loaded in exactly the same way, that's why before the loader a
	converter has to be executed.

	Read here to see the standard output format that loader returns
	https://pkg.go.dev/encoding/json#Unmarshal
*/

type IRawSchemaConverter interface {
	ConvertToJson(b []byte) ([]byte, error)
}

type schemaLoader struct {
	converter IRawSchemaConverter
}

func NewSchemaLoader(converter IRawSchemaConverter) *schemaLoader {
	return &schemaLoader{
		converter: converter,
	}
}

func (s *schemaLoader) LoadSchema(b []byte) (UnparsedSchema, error) {
	// Must be converted to JSON before loading to ensure consistency in the load
	// process independently of the format in which the schema is written
	jsonBytes, err := s.converter.ConvertToJson(b)
	if err != nil {
		return nil, err
	}

	var parsedJson map[string]any
	err = json.Unmarshal(jsonBytes, &parsedJson)
	if err != nil {
		return nil, err
	}

	return parsedJson, nil
}

func (s *schemaLoader) LoadSchemaFromFile(filepath string) (UnparsedSchema, error) {
	jsonFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return s.LoadSchema(jsonFile)
}
