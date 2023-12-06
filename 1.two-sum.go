package main

import "fmt"

// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	num2Index := make(map[int]int)
	for i, num := range nums {
		second := target - num
		if index, ok := num2Index[second]; ok {
			return []int{index, i}
		}
		num2Index[num] = i
	}
	return nil
}

func main() {
	fmt.Printf("%+v\n", twoSum([]int{3, 2, 4}, 6))
}
