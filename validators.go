package main

import (
	"encoding/json"
	"strings"
)

const (
	TypeString JSONSchemaType = "string"
	TypeNumber JSONSchemaType = "number"
	TypeObject JSONSchemaType = "object"
)

func inferType(value string) JSONSchemaType {
	if strings.TrimSpace(value) == "" {
		return TypeString
	}
	if _, err := json.Number(value).Int64(); err == nil {
		return TypeNumber
	}
	if _, err := json.Number(value).Float64(); err == nil {
		return TypeNumber
	}
	return TypeString
}
