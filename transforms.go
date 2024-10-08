package main

import (
	"encoding/json"
	"encoding/xml"
)

func parseXML(xmlString string) (XMLNode, error) {
	var node XMLNode
	err := xml.Unmarshal([]byte(xmlString), &node)
	return node, err
}

func generateJSONSchema(schema JSONSchemaProperty) (string, error) {
	jsonSchema := map[string]interface{}{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type":    schema.Type,
	}
	if len(schema.Properties) > 0 {
		jsonSchema["properties"] = schema.Properties
	}
	jsonBytes, err := json.MarshalIndent(jsonSchema, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func toSchemaProperty(node XMLNode) JSONSchemaProperty {
	schema := JSONSchemaProperty{
		Properties: make(map[string]JSONSchemaProperty),
	}
	if len(node.Nodes) == 0 {
		schema.Type = inferType(node.Content)
	} else {
		schema.Type = TypeObject
		for _, child := range node.Nodes {
			schema.Properties[child.XMLName.Local] = toSchemaProperty(child)
		}
	}
	return schema
}
