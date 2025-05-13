package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	k := len(nums)
	if k <= 1 {
		return k
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
