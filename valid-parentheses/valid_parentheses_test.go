package valid_parentheses

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		// Add your test cases here
		{"Test 1", "()", true},
		{"Test 2", "()[]{}", true},
		{"Test 3", "(]", false},
		{"Test 4", "([])", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("validParentheses() = %v, want %v", got, tt.output)
			}
		})
	}
}
