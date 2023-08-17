package schema

import (
	"errors"
	"fmt"
)

// Function that parses a router and returns a Router struct including
// its queries, mutations and nested routers.
func parseRouter(unparsedRouter map[string]any) (Router, error) {
	var parsedRouter Router = Router{
		Routers:   nil,
		Queries:   nil,
		Mutations: nil,
	}

	for key, value := range unparsedRouter {

		switch key {

		case "queries":
			unparsedQueries, ok := value.(map[string]any)
			if !ok {
				return Router{}, errors.New("queries is not a map")
			}

			parsedQueries, err := parseEndpoints(unparsedQueries)
			if err != nil {
				return Router{}, err
			}

			parsedRouter.Queries = parsedQueries

		case "mutations":
			unparsedMutations, ok := value.(map[string]any)
			if !ok {
				return Router{}, errors.New("mutations is not a map")
			}

			parsedMutations, err := parseEndpoints(unparsedMutations)
			if err != nil {
				return Router{}, err
			}

			parsedRouter.Mutations = parsedMutations

		case "routers":
			// Should recursively call parseRouter for each nested router
			// and build a map of Router structs.
			unparsedNestedRouters, ok := value.(map[string]any)
			if !ok {
				return Router{}, errors.New("routers is not a map")
			}

			newInnerRouters := map[string]Router{}

			for nestedRouterName, nestedRouterRawValue := range unparsedNestedRouters {
				nestedRouterValue, ok := nestedRouterRawValue.(map[string]any)
				if !ok {
					return Router{}, errors.New("nested router is not a map")
				}

				parsedInnerRouter, err := parseRouter(nestedRouterValue)
				if err != nil {
					return Router{}, err
				}

				newInnerRouters[nestedRouterName] = parsedInnerRouter
			}

			parsedRouter.Routers = newInnerRouters

		default:
			if key != "version" {
				return Router{}, fmt.Errorf("unknown schema key: %s", key)
			}

		}

	}

	return parsedRouter, nil
}
