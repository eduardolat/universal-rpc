package schema

import (
	"errors"
)

func ParseSchema(unparsedSchema UnparsedSchema) (Schema, error) {
	var parsedSchema Schema

	version, err := parseVersion(unparsedSchema)
	if err != nil {
		return Schema{}, err
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

func parseVersion(unparsedSchema UnparsedSchema) (int, error) {

	var version int

	if intVersion, ok := unparsedSchema["version"].(int); ok {
		version = intVersion
	} else if floatVersion, ok := unparsedSchema["version"].(float64); ok {
		version = int(floatVersion)
	} else {
		return 0, errors.New("version not found or is not a number")
	}

	return version, nil

}
