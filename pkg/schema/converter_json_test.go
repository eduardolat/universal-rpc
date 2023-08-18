package schema

import (
	"testing"
)

func TestJsonSchemaConverter(t *testing.T) {
	converter := JsonSchemaConverter{}

	// Test that the converter returns the same JSON that it was given.
	rawJson := []byte(`{"test":"test","test2":2,"test3":true,"test4":null,"test5":{"test6":"test6"},"test7":["test7"]}`)
	convertedJson, err := converter.ConvertToJson(rawJson)
	if err != nil {
		t.Errorf("Error converting JSON: %s", err)
	}

	if string(rawJson) != string(convertedJson) {
		t.Errorf("Converted JSON did not match raw JSON: %s != %s", string(rawJson), string(convertedJson))
	}
}
