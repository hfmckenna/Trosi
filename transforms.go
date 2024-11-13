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

func toSchemaProperty(node XMLNode, parent string) JSONSchemaProperty {
	schema := JSONSchemaProperty{
		Properties: make(map[string]JSONSchemaProperty),
	}
	if len(node.Nodes) == 0 {
		schema.Type = inferType(node.Content)
	} else if isArray(node) {
		var types []map[string]JSONSchemaType
		var objectTypes []JSONSchemaProperty
		var currentType JSONSchemaType
		isSameType := true
		for _, n := range node.Nodes {
			s := toSchemaProperty(n, toNestedName(parent, n))
			if s.Type == TypeObject {
				objectTypes = append(objectTypes, s)
			} else {
				types = append(types, map[string]JSONSchemaType{
					"type": s.Type,
				})
			}
			if s.Type != currentType {
				isSameType = false
			}
			currentType = s.Type
		}
		schema.Type = TypeArray
		if len(objectTypes) > 0 {
			schema.Defs = mergeSchemas(objectTypes).Properties
		}
		if !isSameType {
			schema.PrefixItems = types
		} else {
			schema.Items = types[0]
		}
	} else {
		schema.Type = TypeObject
		for _, child := range node.Nodes {
			schema.Properties[child.XMLName.Local] = toSchemaProperty(child, toNestedName(parent, child))
		}
	}
	return schema
}

func isArray(node XMLNode) bool {
	name := node.Nodes[0].XMLName.Local
	for _, n := range node.Nodes {
		if name != n.XMLName.Local {
			return false
		}
		name = n.XMLName.Local
	}
	return true
}

func toNestedName(parent string, node XMLNode) string {
	return parent + "/" + node.XMLName.Local
}
