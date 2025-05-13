package roman_to_integer

import (
	"reflect"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		// Add your test cases here
		{
			name:   "Test 1",
			input:  "III",
			output: 3,
		},
		{
			name:   "Test 2",
			input:  "IV",
			output: 4,
		},
		{
			name:   "Test 3",
			input:  "IX",
			output: 9,
		},
		{
			name:   "Test 4",
			input:  "LVIII",
			output: 58,
		},
		{
			name:   "Test 5",
			input:  "MCMXCIV",
			output: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("romanToInt() = %v, want %v", got, tt.output)
			}
		})
	}
}
