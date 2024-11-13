package main

import "encoding/xml"

type XMLNode struct {
	XMLName xml.Name
	Content string    `xml:",chardata"`
	Nodes   []XMLNode `xml:",any"`
}

type JSONSchemaType string

type JSONSchemaProperty struct {
	Type        JSONSchemaType                `json:"type"`
	Properties  map[string]JSONSchemaProperty `json:"properties,omitempty"`
	PrefixItems []map[string]JSONSchemaType   `json:"prefixItems,omitempty"`
	Items       map[string]JSONSchemaType     `json:"items,omitempty"`
	Defs        map[string]JSONSchemaProperty `json:"$defs,omitempty"`
}

const (
	TypeString  JSONSchemaType = "string"
	TypeNumber  JSONSchemaType = "number"
	TypeInteger JSONSchemaType = "integer"
	TypeBoolean JSONSchemaType = "boolean"
	TypeNull    JSONSchemaType = "null"
	TypeObject  JSONSchemaType = "object"
	TypeArray   JSONSchemaType = "array"
)
