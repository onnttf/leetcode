package longest_common_prefix

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output string
	}{
		// Add your test cases here
		{name: "Test 1", input: []string{"flower", "flow", "flight"}, output: "fl"},
		{name: "Test 2", input: []string{"dog", "racecar", "car"}, output: ""},
		{name: "Test 3", input: []string{"a"}, output: "a"},
		{name: "Test 4", input: []string{"ab", "a"}, output: "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.output)
			}
		})
	}
}
