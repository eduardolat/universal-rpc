package schema

// Implementation of IRawSchemaConverter required by the schema loader.
type JsonSchemaConverter struct{}

// Function that converts the raw JSON schema to JSON.
// This is just a requirement of the schema loader, so we just return the raw JSON.
func (s *JsonSchemaConverter) ConvertToJson(b []byte) ([]byte, error) {
	return b, nil
}
