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
	// Remember that the rawUnparsedField is json like format (first letter is lowercase)
	unparsedField, ok := rawUnparsedField.(map[string]any)
	if !ok {
		return Field{}, fmt.Errorf("field is not a map")
	}

	if _, ok := unparsedField["type"].(string); !ok {
		return Field{}, fmt.Errorf("field is missing type")
	}

	isAllowedType := slices.Contains(PrimitiveTypes[:], unparsedField["type"].(string))
	isAllowedStructure := slices.Contains(DataStructures[:], unparsedField["type"].(string))
	if !isAllowedType && !isAllowedStructure {
		return Field{}, nil
	}

	// Should be required by default
	isRequired := true
	unparsedRequired, ok := unparsedField["required"].(bool)
	if ok {
		isRequired = unparsedRequired
	}

	parsedField := Field{
		Required: isRequired,
		Type:     unparsedField["type"].(string),
	}

	if unparsedField["type"] == "array" {
		if _, ok := unparsedField["items"]; !ok {
			return Field{}, fmt.Errorf("array field should include items")
		}

		parsedItemsField, err := parseField(unparsedField["items"])
		if err != nil {
			return Field{}, err
		}

		parsedField.Items = &parsedItemsField
	}

	if unparsedField["type"] == "object" {
		rawProperties, ok := unparsedField["properties"]
		if !ok {
			return Field{}, fmt.Errorf("object field should include properties")
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
