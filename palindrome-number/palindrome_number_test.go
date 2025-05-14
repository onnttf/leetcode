package palindrome_number

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output any
	}{
		// Add your test cases here
		{name: "Test 1", input: 121, output: true},
		{name: "Test 2", input: -121, output: false},
		{name: "Test 3", input: 10, output: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.output)
			}
		})
	}
}
