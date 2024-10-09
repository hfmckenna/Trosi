package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
