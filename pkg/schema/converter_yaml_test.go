package schema

import (
	"testing"
)

func TestYamlSchemaConverter(t *testing.T) {
	converter := YamlSchemaConverter{}

	rawYaml := []byte("test: test\ntest2: 2\ntest3: true\ntest4: null\ntest5:\n  test6: test6\ntest7:\n- test7")
	expectedJson := []byte(`{"test":"test","test2":2,"test3":true,"test4":null,"test5":{"test6":"test6"},"test7":["test7"]}`)

	convertedYaml, err := converter.ConvertToJson(rawYaml)
	if err != nil {
		t.Fatal("Error converting YAML")
	}

	if string(expectedJson) != string(convertedYaml) {
		t.Fatal("Converted YAML did not match expected JSON")
	}
}
