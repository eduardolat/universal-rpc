package schema

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Implementation of IRawSchemaConverter required by the schema loader.
type YamlSchemaConverter struct{}

// Function that converts the raw YAML schema to JSON.
func (s *YamlSchemaConverter) ConvertToJson(b []byte) ([]byte, error) {
	var parsedYaml map[string]any

	err := yaml.Unmarshal(b, &parsedYaml)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(parsedYaml)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}
