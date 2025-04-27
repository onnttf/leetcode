package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		output []int
	}{
		// Add your test cases here
		{
			name:   "Test1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			name:   "Test2",
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			name:   "Test3",
			nums:   []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.output)
			}
		})
	}
}
