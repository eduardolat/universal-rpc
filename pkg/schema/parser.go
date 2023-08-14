package schema

import (
	"errors"
	"fmt"
)

func ParseSchema(unparsedSchema UnparsedSchema) (Schema, error) {
	var parsedSchema Schema

	version, ok := unparsedSchema["version"].(float64)
	if !ok {
		return Schema{}, errors.New("version not found or is not a number")
	}
	parsedSchema.Version = int(version)

	// Pass the top level map to parseRouter because it is a top level router
	// and it will recursively call parseRouter for each nested router.
	router, err := parseRouter(unparsedSchema)
	if err != nil {
		return Schema{}, err
	}

	parsedSchema.Routers = router.Routers
	parsedSchema.Queries = router.Queries
	parsedSchema.Mutations = router.Mutations

	return parsedSchema, nil
}

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

			parsedQueries, err := parseQueries(unparsedQueries)
			if err != nil {
				return Router{}, err
			}

			parsedRouter.Queries = parsedQueries

		case "mutations":
			unparsedMutations, ok := value.(map[string]any)
			if !ok {
				return Router{}, errors.New("mutations is not a map")
			}

			parsedMutations, err := parseMutations(unparsedMutations)
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

func parseQueries(unparsedQueries map[string]any) (map[string]Endpoint, error) {
	return nil, nil
}

func parseMutations(unparsedMutations map[string]any) (map[string]Endpoint, error) {
	return nil, nil
}
