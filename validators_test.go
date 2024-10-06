package main

import "testing"

func TestInferType(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expect JSONSchemaType
	}{
		{
			name:   "empty",
			input:  " ",
			expect: TypeString,
		},
		{
			name:   "string",
			input:  "test",
			expect: TypeString,
		},
		{
			name:   "number",
			input:  "1",
			expect: TypeNumber,
		},
		{
			name:   "float",
			input:  "1.1",
			expect: TypeNumber,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := inferType(tc.input)
			if result != tc.expect {
				t.Errorf("inferType(%q): expected %v, got %v", tc.input, tc.expect, result)
			}
		})
	}
}
