package schema

import "testing"

func TestParseFields(t *testing.T) {
	// This only tests a bit of the functionality to ensure that this
	// function is working, the internal functionality is tested in
	// TestParseField

	unparsedFields := map[string]any{
		"foo": map[string]any{
			"Type": "string",
		},
		"bar": map[string]any{
			"Type":     "boolean",
			"Required": false,
		},
		"aaa": map[string]any{
			"Type": "number",
		},
		"bbb": map[string]any{
			"Type": "null",
		},
		"ccc": map[string]any{
			"Type": "any",
		},
		"baz": map[string]any{
			"Type": "array",
			"Items": map[string]any{
				"Type":     "string",
				"Required": false,
			},
		},
		"qux": map[string]any{
			"Type": "object",
			"Properties": map[string]any{
				"foo": map[string]any{
					"Type": "string",
				},
				"bar": map[string]any{
					"Type": "boolean",
				},
			},
		},
	}

	parsedFields, err := parseFields(unparsedFields)

	if err != nil {
		t.Error(err)
	}

	if len(parsedFields) != len(unparsedFields) {
		t.Error("Expected 7 fields")
	}

	if parsedFields["foo"].Type != "string" {
		t.Error("Expected foo to be a string")
	}

	if parsedFields["bar"].Type != "boolean" {
		t.Error("Expected bar to be a boolean")
	}

	if parsedFields["bar"].Required {
		t.Error("Expected bar to be not required")
	}

	if parsedFields["aaa"].Type != "number" {
		t.Error("Expected aaa to be a number")
	}

	if parsedFields["bbb"].Type != "null" {
		t.Error("Expected bbb to be a null")
	}

	if parsedFields["ccc"].Type != "any" {
		t.Error("Expected ccc to be a any")
	}

	if parsedFields["baz"].Type != "array" {
		t.Error("Expected baz to be a array")
	}

	if parsedFields["baz"].Items == nil {
		t.Error("Expected baz to have Items")
	}

	if parsedFields["baz"].Items.Type != "string" {
		t.Error("Expected baz to have Items of type string")
	}

	if parsedFields["baz"].Items.Required {
		t.Error("Expected baz to have Items not required")
	}

	if parsedFields["qux"].Type != "object" {
		t.Error("Expected qux to be a object")
	}

	if parsedFields["qux"].Properties == nil {
		t.Error("Expected qux to have Properties")
	}

	if parsedFields["qux"].Properties["foo"].Type != "string" {
		t.Error("Expected qux to have Properties with foo of type string")
	}

	if parsedFields["qux"].Properties["bar"].Type != "boolean" {
		t.Error("Expected qux to have Properties with bar of type boolean")
	}

}

func TestParseField(t *testing.T) {

	t.Run("Must be a map", func(t *testing.T) {
		_, err := parseField("test")
		if err == nil {
			t.Error("Expected error")
		}
	})

	t.Run("Must have at least Type field", func(t *testing.T) {
		_, err := parseField(map[string]any{})
		if err == nil {
			t.Error("Expected error")
		}
	})

	t.Run("Must parse basic field", func(t *testing.T) {
		_, err := parseField(map[string]any{
			"Type": "string",
		})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Must be required by default and if specified respect the value", func(t *testing.T) {
		parsedField, err := parseField(map[string]any{
			"Type": "string",
		})
		if err != nil {
			t.Error(err)
		}
		if !parsedField.Required {
			t.Error("Expected field to be required if not specified")
		}

		parsedField, err = parseField(map[string]any{
			"Type":     "string",
			"Required": true,
		})
		if err != nil {
			t.Error(err)
		}
		if !parsedField.Required {
			t.Error("Expected field to be required if is explicitly specified")
		}

		parsedField, err = parseField(map[string]any{
			"Type":     "string",
			"Required": false,
		})
		if err != nil {
			t.Error(err)
		}
		if parsedField.Required {
			t.Error("Expected field to be not required if is explicitly specified")
		}
	})

	t.Run("Should ignore fields with unknown types", func(t *testing.T) {
		parsedField, err := parseField(map[string]any{
			"Type": "Other-Random-Thing",
		})
		if err != nil {
			t.Error(err)
		}

		if parsedField.Type != "" {
			t.Error("Field should be empty")
		}
	})

	t.Run("If the field is an array it should have the Items and parse correctly", func(t *testing.T) {
		_, err := parseField(map[string]any{
			"Type": "array",
		})
		if err == nil {
			t.Error("Expected error because Items is not specified")
		}

		parsedField, err := parseField(map[string]any{
			"Type": "array",
			"Items": map[string]any{
				"Type": "string",
			},
		})
		if err != nil {
			t.Error(err)
		}

		if parsedField.Items == nil {
			t.Error("Expected Items to be parsed")
		}

		if parsedField.Items.Type != "string" {
			t.Error("Expected Items to be parsed")
		}
	})

	t.Run("If the field is a object it should have the Properties and parse correctly", func(t *testing.T) {
		_, err := parseField(map[string]any{
			"Type": "object",
		})
		if err == nil {
			t.Error("Expected error because Properties is not specified")
		}

		parsedField, err := parseField(map[string]any{
			"Type": "object",
			"Properties": map[string]any{
				"foo": map[string]any{
					"Type": "string",
				},
				"bar": map[string]any{
					"Type": "boolean",
				},
			},
		})
		if err != nil {
			t.Error(err)
		}

		if parsedField.Properties == nil {
			t.Error("Expected Properties to be parsed")
		}

		if parsedField.Properties["foo"].Type != "string" {
			t.Error("Expected Properties to be parsed")
		}

		if parsedField.Properties["bar"].Type != "boolean" {
			t.Error("Expected Properties to be parsed")
		}
	})

}
