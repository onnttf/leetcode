package remove_duplicates_from_sorted_array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		// Add your test cases here
		{name: "Test 1", input: []int{1, 1, 2}, output: 2},
		{name: "Test 2", input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, output: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("removeDuplicatesFromSortedArray() = %v, want %v", got, tt.output)
			}
		})
	}
}
