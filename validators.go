package main

import (
	"encoding/json"
	"strings"
)

func inferType(value string) JSONSchemaType {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return TypeNull
	}
	if trimmed == "true" || trimmed == "false" {
		return TypeBoolean
	}
	if _, err := json.Number(trimmed).Int64(); err == nil {
		return TypeInteger
	}
	if _, err := json.Number(trimmed).Float64(); err == nil {
		return TypeNumber
	}
	return TypeString
}
