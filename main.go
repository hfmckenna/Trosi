package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
			schema := toSchemaProperty(node)
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
