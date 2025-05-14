package remove_element

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		val    int
		output int
	}{
		// Add your test cases here
		{name: "Test 1", nums: []int{3, 2, 2, 3}, val: 3, output: 2},
		{name: "Test 2", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, output: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("removeElement() = %v, want %v", got, tt.output)
			}
		})
	}
}
