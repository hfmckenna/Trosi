package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

type XMLNode struct {
	XMLName xml.Name
	Content string    `xml:",chardata"`
	Nodes   []XMLNode `xml:",any"`
}

type JSONSchemaType string

type JSONSchemaProperty struct {
	Type       JSONSchemaType                `json:"type"`
	Properties map[string]JSONSchemaProperty `json:"properties,omitempty"`
}

func extractSchema(node XMLNode) JSONSchemaProperty {
	schema := JSONSchemaProperty{
		Properties: make(map[string]JSONSchemaProperty),
	}
	if len(node.Nodes) == 0 {
		schema.Type = inferType(node.Content)
	} else {
		schema.Type = TypeObject
		for _, child := range node.Nodes {
			schema.Properties[child.XMLName.Local] = extractSchema(child)
		}
	}
	return schema
}

func mergeSchemas(schemas []JSONSchemaProperty) JSONSchemaProperty {
	merged := JSONSchemaProperty{
		Type:       TypeObject,
		Properties: make(map[string]JSONSchemaProperty),
	}
	for _, schema := range schemas {
		for name, prop := range schema.Properties {
			if existing, ok := merged.Properties[name]; ok {
				merged.Properties[name] = mergeProp(existing, prop)
			} else {
				merged.Properties[name] = prop
			}
		}
	}
	return merged
}

func mergeProp(p1, p2 JSONSchemaProperty) JSONSchemaProperty {
	if p1.Type != p2.Type {
		return JSONSchemaProperty{Type: TypeString} // Default to string if types don't match
	}
	if p1.Type == TypeObject {
		return mergeSchemas([]JSONSchemaProperty{p1, p2})
	}
	return p1 // If same type and not object, just return one of them
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

func main() {
	files, err := os.ReadDir(".")
	var schemas []JSONSchemaProperty
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".xml") {
			contents := readFile(RealFileSystem{}, file.Name())
			node, err := parseXML(contents)
			if err != nil {
				log.Fatal("Error when parsing XML")
			}
			schema := extractSchema(node)
			schemas = append(schemas, schema)
		}
	}
	mergedSchema := mergeSchemas(schemas)
	jsonSchema, err := generateJSONSchema(mergedSchema)
	if err != nil {
		fmt.Println("Error generating JSON Schema:", err)
		return
	}
	fmt.Println(jsonSchema)
}
