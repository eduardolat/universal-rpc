package main

import (
	"fmt"

	"github.com/eduardolat/universal-rpc/pkg/schema"
)

func main() {

	unparsedSchema, err := schema.LoadJsonSchemaFromFile("test-schema.json")
	if err != nil {
		panic(err)
	}

	parsedSchema, err := schema.ParseSchema(unparsedSchema)
	if err != nil {
		panic(err)
	}

	fmt.Println(parsedSchema)

}
