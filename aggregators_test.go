package main

import (
	"reflect"
	"testing"
)

func TestMergeSchemas(t *testing.T) {
	testCases := []struct {
		name     string
		input    []JSONSchemaProperty
		expected JSONSchemaProperty
	}{
		{
			name:     "mismatch",
			input:    []JSONSchemaProperty{{Type: TypeObject, Properties: map[string]JSONSchemaProperty{"a": {Type: TypeObject}}}, {Type: TypeObject, Properties: map[string]JSONSchemaProperty{"a": {Type: TypeNumber}}}},
			expected: JSONSchemaProperty{Type: TypeObject, Properties: map[string]JSONSchemaProperty{"a": {Type: TypeString}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeSchemas(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("toSchemaProperty() = %v, want %v", result, tc.expected)
			}
		})
	}
}
