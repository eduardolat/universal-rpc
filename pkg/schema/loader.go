package schema

type SchemaLoader interface {
	LoadSchema(b []byte) (UnparsedSchema, error)
	LoadSchemaFromFile(filepath string) (UnparsedSchema, error)
}
