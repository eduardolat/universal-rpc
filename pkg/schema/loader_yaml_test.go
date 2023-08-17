package schema

import (
	"testing"
)

func TestYamlLoadSchema(t *testing.T) {
	loader := &YamlSchemaLoader{}
	schema := []byte("str: \"value\"\nint: 1\nfloat: 1.1\nbool: true")

	result, err := loader.LoadSchema(schema)
	if err != nil {
		t.Errorf("Error al cargar el esquema: %s", err)
	}

	if result["str"] != "value" {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["str"], "value")
	}

	if float64(result["int"].(int)) != float64(1) {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["int"], float64(1))
	}

	if result["float"] != float64(1.1) {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["float"], float64(1.1))
	}

	if result["bool"] != true {
		t.Errorf("Resultado inesperado: obtuvo %v, esperaba %v", result["bool"], true)
	}
}
