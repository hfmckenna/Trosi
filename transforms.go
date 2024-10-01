package main

import (
	"encoding/xml"
)

func parseXML(xmlString string) (XMLNode, error) {
	var node XMLNode
	err := xml.Unmarshal([]byte(xmlString), &node)
	return node, err
}
