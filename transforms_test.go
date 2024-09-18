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
