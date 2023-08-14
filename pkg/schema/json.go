package schema

import (
	"encoding/json"
	"os"
)

func LoadJsonSchema(b []byte) (UnparsedSchema, error) {
	var parsedJson UnparsedSchema

	err := json.Unmarshal(b, &parsedJson)
	if err != nil {
		return nil, err
	}

	return parsedJson, nil
}

func LoadJsonSchemaFromFile(filename string) (UnparsedSchema, error) {
	jsonFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return LoadJsonSchema(jsonFile)
}
