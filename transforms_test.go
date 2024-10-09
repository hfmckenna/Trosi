package main

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestParseXML(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		input    string
		expected XMLNode
		wantErr  bool
	}{
		{
			name:  "Valid XML",
			input: "<root><child>value</child></root>",
			expected: XMLNode{
				XMLName: xml.Name{Local: "root"},
				Nodes: []XMLNode{
					{
						XMLName: xml.Name{Local: "child"},
						Content: "value",
					},
				},
			},
			wantErr: false,
		},
		{
			name:     "Invalid XML",
			input:    "<root><child>value</child>",
			expected: XMLNode{},
			wantErr:  true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseXML(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("parseXML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if (!reflect.DeepEqual(result, tc.expected)) != tc.wantErr {
				t.Errorf("parseXML() = %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestGenerateJSONSchema(t *testing.T) {
	testCases := []struct {
		name     string
		input    JSONSchemaProperty
		expected string
	}{
		{
			name:     "number",
			input:    JSONSchemaProperty{Type: TypeNumber},
			expected: readFile(RealFileSystem{}, "test/number.json"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := generateJSONSchema(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("generateJSONSchema() = %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestToSchemaProperty(t *testing.T) {
	testCases := []struct {
		name     string
		input    XMLNode
		expected JSONSchemaProperty
	}{
		{
			name: "string",
			input: XMLNode{
				XMLName: xml.Name{Local: "root"},
				Content: "test",
				Nodes:   nil,
			},
			expected: JSONSchemaProperty{
				Type:       "string",
				Properties: nil,
			},
		},
		{
			name: "number",
			input: XMLNode{
				XMLName: xml.Name{Local: "root"},
				Content: "1",
				Nodes:   nil,
			},
			expected: JSONSchemaProperty{
				Type:       "integer",
				Properties: nil,
			},
		},
		{
			name: "object",
			input: XMLNode{
				XMLName: xml.Name{Local: "root"},
				Content: "",
				Nodes:   []XMLNode{{XMLName: xml.Name{Local: "child"}, Content: "value"}},
			},
			expected: JSONSchemaProperty{
				Type:       "object",
				Properties: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toSchemaProperty(tc.input)
			if !reflect.DeepEqual(result.Type, tc.expected.Type) {
				t.Errorf("toSchemaProperty() = %v, want %v", result, tc.expected)
			}
		})
	}
}
