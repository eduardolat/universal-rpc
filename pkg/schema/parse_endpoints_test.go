package schema

import "testing"

func TestParseEndpoints(t *testing.T) {

	t.Run("Allow empty endpoints", func(t *testing.T) {

		unparsed := map[string]any{
			"foo": map[string]any{},
		}

		_, err := parseEndpoints(unparsed)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Allow endpoints with no Returns", func(t *testing.T) {

		unparsed := map[string]any{
			"foo": map[string]any{
				"Params": map[string]any{},
			},
		}

		_, err := parseEndpoints(unparsed)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Correctly parse single endpoint", func(t *testing.T) {

		unparsed := map[string]any{
			"foo": map[string]any{
				"Params": map[string]any{
					"foo": map[string]any{
						"Type": "string",
					},
					"bar": map[string]any{
						"Type": "number",
					},
				},
				"Returns": map[string]any{
					"baz": map[string]any{
						"Type": "string",
					},
				},
			},
		}

		endpoints, err := parseEndpoints(unparsed)
		if err != nil {
			t.Fatal(err)
		}

		if len(endpoints) != 1 {
			t.Fatalf("Expected 1 endpoint, got %d", len(endpoints))
		}

		// Check if has foo endpoint
		if _, ok := endpoints["foo"]; !ok {
			t.Fatalf("Expected endpoint foo, got %v", endpoints)
		}

	})

	t.Run("Correctly parse multiple endpoints", func(t *testing.T) {
		unparsed := map[string]any{
			"foo": map[string]any{
				"Params": map[string]any{
					"foo": map[string]any{
						"Type": "string",
					},
					"bar": map[string]any{
						"Type": "number",
					},
				},
				"Returns": map[string]any{
					"baz": map[string]any{
						"Type": "string",
					},
				},
			},
			"bar": map[string]any{
				"Params": map[string]any{
					"bar": map[string]any{
						"Type": "number",
					},
				},
				"Returns": map[string]any{
					"biz": map[string]any{
						"Type": "string",
					},
				},
			},
		}

		endpoints, err := parseEndpoints(unparsed)
		if err != nil {
			t.Fatal(err)
		}

		if len(endpoints) != 2 {
			t.Fatalf("Expected 2 endpoints, got %d", len(endpoints))
		}

		if _, ok := endpoints["foo"]; !ok {
			t.Fatalf("Expected endpoint foo")
		}

		if _, ok := endpoints["bar"]; !ok {
			t.Fatalf("Expected endpoint bar")
		}

	})

	t.Run("An endpoint can only have Params and Returns", func(t *testing.T) {

		unparsed := map[string]any{
			"foo": map[string]any{
				"Bar": map[string]any{},
			},
		}

		_, err := parseEndpoints(unparsed)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}

	})
}
