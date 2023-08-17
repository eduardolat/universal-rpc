package schema

import (
	"errors"
)

func ParseSchema(unparsedSchema UnparsedSchema) (Schema, error) {
	var parsedSchema Schema

	version, ok := unparsedSchema["version"].(int)
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
