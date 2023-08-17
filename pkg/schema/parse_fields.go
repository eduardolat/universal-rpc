package schema

import (
	"fmt"
	"slices"
)

func parseFields(unparsedFields map[string]any) (map[string]Field, error) {

	parsedFields := map[string]Field{}

	for propertyName, rawProperty := range unparsedFields {
		parsedProperty, err := parseField(rawProperty)
		if err != nil {
			return nil, err
		}

		parsedFields[propertyName] = parsedProperty
	}

	return parsedFields, nil
}

func parseField(rawUnparsedField any) (Field, error) {
	unparsedField, ok := rawUnparsedField.(map[string]any)
	if !ok {
		return Field{}, fmt.Errorf("field is not a map")
	}

	if _, ok := unparsedField["Type"].(string); !ok {
		return Field{}, fmt.Errorf("field is missing type")
	}

	isAllowedType := slices.Contains(PrimitiveTypes[:], unparsedField["Type"].(string))
	isAllowedStructure := slices.Contains(DataStructures[:], unparsedField["Type"].(string))
	if !isAllowedType && !isAllowedStructure {
		return Field{}, nil
	}

	// Should be required by default
	isRequired := true
	unparsedRequired, ok := unparsedField["Required"].(bool)
	if ok {
		isRequired = unparsedRequired
	}

	parsedField := Field{
		Required: isRequired,
		Type:     unparsedField["Type"].(string),
	}

	if unparsedField["Type"] == "array" {
		if _, ok := unparsedField["Items"]; !ok {
			return Field{}, fmt.Errorf("array field should include Items")
		}

		parsedItemsField, err := parseField(unparsedField["Items"])
		if err != nil {
			return Field{}, err
		}

		parsedField.Items = &parsedItemsField
	}

	if unparsedField["Type"] == "object" {
		rawProperties, ok := unparsedField["Properties"]
		if !ok {
			return Field{}, fmt.Errorf("object field should include Properties")
		}

		unparsedProperties, ok := rawProperties.(map[string]any)
		if !ok {
			return Field{}, fmt.Errorf("properties should be a map")
		}

		parsedPropertiesField := map[string]*Field{}

		for propertyName, rawProperty := range unparsedProperties {
			parsedProperty, err := parseField(rawProperty)
			if err != nil {
				return Field{}, err
			}

			parsedPropertiesField[propertyName] = &parsedProperty
		}

		parsedField.Properties = parsedPropertiesField

	}

	return parsedField, nil
}
