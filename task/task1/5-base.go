package main

import "fmt"

func main() {
	// 两数之和
	result1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result1)
	result2 := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result2)
	result3 := twoSum([]int{3, 3}, 6)
	fmt.Println(result3)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
