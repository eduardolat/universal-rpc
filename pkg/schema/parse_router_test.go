package schema

import "testing"

func TestParseRouter(t *testing.T) {

	t.Run("should throw an error if router is not correctly formatted", func(t *testing.T) {

		unparsedRouter := map[string]any{
			"incorrect": "something",
		}

		_, err := parseRouter(unparsedRouter)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

	})

	t.Run("should parse a router correctly", func(t *testing.T) {

		unparsedRouter := map[string]any{
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
		}

		parsedRouter, err := parseRouter(unparsedRouter)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if parsedRouter.Queries["query1"].Params["param1"].Type != "string" {
			t.Errorf("expected query1.param1.type to be string, got %v", parsedRouter.Queries["query1"].Params["param1"].Type)
		}

		if parsedRouter.Queries["query1"].Returns["return1"].Type != "string" {
			t.Errorf("expected query1.return1.type to be string, got %v", parsedRouter.Queries["query1"].Returns["return1"].Type)
		}

		if parsedRouter.Mutations["mutation1"].Params["param1"].Type != "string" {
			t.Errorf("expected mutation1.param1.type to be string, got %v", parsedRouter.Mutations["mutation1"].Params["param1"].Type)
		}

		if parsedRouter.Mutations["mutation1"].Returns["return1"].Type != "string" {
			t.Errorf("expected mutation1.return1.type to be string, got %v", parsedRouter.Mutations["mutation1"].Returns["return1"].Type)
		}

		if parsedRouter.Routers["subrouter2"].Queries["query2"].Params["param1"].Type != "string" {
			t.Errorf("expected subrouter2.query2.param1.type to be string, got %v", parsedRouter.Routers["subrouter2"].Queries["query2"].Params["param1"].Type)
		}

	})

}
