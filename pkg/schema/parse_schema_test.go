package schema

import "testing"

func TestParseSchema(t *testing.T) {

	t.Run("should throw an error if schema is not correctly formatted", func(t *testing.T) {
		unparsedSchema := map[string]any{
			"version":   1,
			"incorrect": "something",
		}

		_, err := ParseSchema(unparsedSchema)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("should parse a router correctly", func(t *testing.T) {

		unparsedSchema := map[string]any{
			"version": 1,
			"routers": map[string]any{
				"router1": map[string]any{
					"queries": map[string]any{
						"query1": map[string]any{
							"params": map[string]any{
								"param1": map[string]any{
									"type": "string",
								},
							},
							"returns": map[string]any{
								"return1": map[string]any{
									"type": "string",
								},
							},
						},
					},
					"mutations": map[string]any{
						"mutation1": map[string]any{
							"params": map[string]any{
								"param1": map[string]any{
									"type": "string",
								},
							},
							"returns": map[string]any{
								"return1": map[string]any{
									"type": "string",
								},
							},
						},
					},
					"routers": map[string]any{
						"subrouter2": map[string]any{
							"queries": map[string]any{
								"query2": map[string]any{
									"params": map[string]any{
										"param1": map[string]any{
											"type": "string",
										},
									},
								},
							},
						},
					},
				},
				"emptyRouter": map[string]any{},
			},
		}

		parsedSchema, err := ParseSchema(unparsedSchema)
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}

		if parsedSchema.Version != 1 {
			t.Fatalf("expected version 1, got %v", parsedSchema.Version)
		}

		if parsedSchema.Routers["router1"].Queries["query1"].Params["param1"].Type != "string" {
			t.Fatalf("expected query1.param1.type to be string, got %v", parsedSchema.Routers["router1"].Queries["query1"].Params["param1"].Type)
		}

		if parsedSchema.Routers["router1"].Queries["query1"].Returns["return1"].Type != "string" {
			t.Fatalf("expected query1.return1.type to be string, got %v", parsedSchema.Routers["router1"].Queries["query1"].Returns["return1"].Type)
		}

		if parsedSchema.Routers["router1"].Mutations["mutation1"].Params["param1"].Type != "string" {
			t.Fatalf("expected mutation1.param1.type to be string, got %v", parsedSchema.Routers["router1"].Mutations["mutation1"].Params["param1"].Type)
		}

		if parsedSchema.Routers["router1"].Mutations["mutation1"].Returns["return1"].Type != "string" {
			t.Fatalf("expected mutation1.return1.type to be string, got %v", parsedSchema.Routers["router1"].Mutations["mutation1"].Returns["return1"].Type)
		}

		if parsedSchema.Routers["router1"].Routers["subrouter2"].Queries["query2"].Params["param1"].Type != "string" {
			t.Fatalf("expected query2.param1.type to be string, got %v", parsedSchema.Routers["router1"].Routers["subrouter2"].Queries["query2"].Params["param1"].Type)
		}

	})

	t.Run("should accept top level queries and mutations", func(t *testing.T) {
		unparsedSchema := map[string]any{
			"version": 1,
			"queries": map[string]any{
				"query1": map[string]any{
					"params": map[string]any{
						"param1": map[string]any{
							"type": "string",
						},
					},
				},
			},
			"mutations": map[string]any{
				"mutation1": map[string]any{
					"params": map[string]any{
						"param1": map[string]any{
							"type": "string",
						},
					},
				},
			},
		}

		parsedSchema, err := ParseSchema(unparsedSchema)
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}

		if parsedSchema.Queries["query1"].Params["param1"].Type != "string" {
			t.Fatalf("expected query1.param1.type to be string, got %v", parsedSchema.Queries["query1"].Params["param1"].Type)
		}

		if parsedSchema.Mutations["mutation1"].Params["param1"].Type != "string" {
			t.Fatalf("expected mutation1.param1.type to be string, got %v", parsedSchema.Mutations["mutation1"].Params["param1"].Type)
		}

	})

}

func TestParseVersion(t *testing.T) {
	t.Run("should parse a int version correctly", func(t *testing.T) {
		unparsed := map[string]any{
			"version": 1,
		}

		version, err := parseVersion(unparsed)
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}

		if version != 1 {
			t.Fatalf("expected 1, got %v", version)
		}
	})

	t.Run("should parse a float64 version correctly", func(t *testing.T) {
		unparsed := map[string]any{
			"version": 1.1234,
		}

		version, err := parseVersion(unparsed)
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}

		if version != 1 {
			t.Fatalf("expected 1, got %v", version)
		}
	})

	t.Run("should throw an error if version is not a number", func(t *testing.T) {
		unparsed := map[string]any{
			"version": "1",
		}

		_, err := parseVersion(unparsed)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
