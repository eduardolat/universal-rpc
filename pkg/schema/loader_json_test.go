package schema

import (
	"testing"
)

func TestJsonLoadSchema(t *testing.T) {
	loader := &JsonSchemaLoader{}
	schema := []byte(`{"str": "value", "int": 1, "float": 1.1, "bool": true }`)

	result, err := loader.LoadSchema(schema)
	if err != nil {
		t.Errorf("Error al cargar el esquema: %s", err)
	}

	if result["str"] != "value" {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["str"], "value")
	}

	if result["int"] != float64(1) {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["int"], float64(1))
	}

	if result["float"] != float64(1.1) {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["float"], float64(1.1))
	}

	if result["bool"] != true {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["bool"], true)
	}
}
