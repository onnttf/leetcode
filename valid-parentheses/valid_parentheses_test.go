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
		{"Test1", "()", true},
		{"Test2", "()[]{}", true},
		{"Test3", "(]", false},
		{"Test4", "([])", true},
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
