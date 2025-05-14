package merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		name   string
		list1  *ListNode
		list2  *ListNode
		output *ListNode
	}{
		// Add your test cases here
		{
			name:   "Test 1",
			list1:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			list2:  &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			output: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
		},
		{
			name:   "Test 2",
			list1:  nil,
			list2:  nil,
			output: nil,
		},
		{
			name:   "Test 3",
			list1:  nil,
			list2:  &ListNode{Val: 0},
			output: &ListNode{Val: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.output)
			}
		})
	}
}
