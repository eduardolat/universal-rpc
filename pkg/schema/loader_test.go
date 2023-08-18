package schema

import (
	"testing"
)

func TestSchemaLoader(t *testing.T) {
	converter := JsonSchemaConverter{}
	loader := NewSchemaLoader(&converter)

	schema := []byte(`{"str": "value", "int": 1, "float": 1.1, "negative_float": -15.134, "bool": true }`)

	result, err := loader.LoadSchema(schema)
	if err != nil {
		t.Errorf("Error al cargar el esquema: %s", err)
	}

	if result["str"] != "value" {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["str"], "value")
	}

	// Why float64? https://pkg.go.dev/encoding/json#Unmarshal
	if result["int"] != float64(1) {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["int"], float64(1))
	}

	// Why float64? https://pkg.go.dev/encoding/json#Unmarshal
	if result["float"] != float64(1.1) {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["float"], float64(1.1))
	}

	// Why float64? https://pkg.go.dev/encoding/json#Unmarshal
	if result["negative_float"] != float64(-15.134) {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["negative_float"], float64(-15.134))
	}

	if result["bool"] != true {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["bool"], true)
	}
}
