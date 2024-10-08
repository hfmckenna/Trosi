package main

import "encoding/xml"

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
