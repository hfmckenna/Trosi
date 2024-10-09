package main

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
