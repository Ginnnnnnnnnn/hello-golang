package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	// 删除有序数组中的重复项
	removeDuplicatesDatas := [][]int{{1, 1, 2}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}
	for _, value := range removeDuplicatesDatas {
		count := removeDuplicates(value)
		fmt.Println(count)
		reuslt := "[ "
		for i := 0; i < count; i++ {
			reuslt = reuslt + strconv.Itoa(value[i]) + " "
		}
		fmt.Println(reuslt + "]")
	}
	// 合并区间
	mergeData := [][][]int{{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, {{1, 4}, {4, 5}}, {{1, 4}, {0, 4}}}
	for _, value := range mergeData {
		result := merge(value)
		fmt.Println(result)
	}
}

func removeDuplicates(nums []int) int {
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(item1 []int, item2 []int) int {
		return item1[0] - item2[0]
	})
	results := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		intervals := intervals[i]
		if i > 0 && intervals[0] <= results[len(results)-1][1] {
			results[len(results)-1][1] = max(results[len(results)-1][1], intervals[1])
		} else {
			results = append(results, intervals)
		}
	}
	return results
}
