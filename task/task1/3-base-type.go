package main

import "fmt"

func main() {
	// 加一
	plusOneDatas := [][]int{{1, 2, 3}, {4, 3, 2, 1}, {9}}
	for _, value := range plusOneDatas {
		result := plusOne(value)
		fmt.Println(result)
	}
}

func plusOne(digits []int) []int {
	var result []int
	baseNum := 1
	i := len(digits) - 1
	for ; i >= 0; i-- {
		num := digits[i] + baseNum
		if num > 9 {
			result = append([]int{0}, result...)
			baseNum = 1
		} else {
			result = append([]int{num}, result...)
			baseNum = 0
		}
	}
	if baseNum == 1 {
		result = append([]int{1}, result...)
	}
	return result
}
