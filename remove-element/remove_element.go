package remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slowIndex := 0

	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}

	return slowIndex
}
