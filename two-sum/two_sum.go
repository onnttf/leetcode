package two_sum

func TwoSum(nums []int, target int) []int {
	// return v1(nums, target)
	return v2(nums, target)
}

func v2(nums []int, target int) []int {
	value2Index := make(map[int]int, len(nums))
	for i, v := range nums {
		if index, ok := value2Index[target-v]; ok {
			return []int{index, i}
		}
		value2Index[v] = i
	}
	return nil
}

func v1(nums []int, target int) []int {
	for i, current := range nums {
		for j := i + 1; j < len(nums); j++ {
			next := nums[j]
			if current+next == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 1}
}
