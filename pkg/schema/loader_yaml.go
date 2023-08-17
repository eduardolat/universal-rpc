package schema

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YamlSchemaLoader struct{}

func (s *YamlSchemaLoader) LoadSchema(b []byte) (UnparsedSchema, error) {
	var parsedYaml UnparsedSchema

	err := yaml.Unmarshal(b, &parsedYaml)
	if err != nil {
		return nil, err
	}

	return parsedYaml, nil
}

func (s *YamlSchemaLoader) LoadSchemaFromFile(filepath string) (UnparsedSchema, error) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return s.LoadSchema(yamlFile)
}
