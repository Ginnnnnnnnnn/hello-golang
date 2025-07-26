package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 只出现一次的数字
	singleNumberDatas := [][]int{{2, 2, 1}, {4, 1, 2, 1, 2}, {1}}
	for _, value := range singleNumberDatas {
		num := singleNumber(value)
		fmt.Println(num)
	}
	//  回文数
	isPalindromeDatas := []int{121, -121, 10}
	for _, value := range isPalindromeDatas {
		result := isPalindrome(value)
		fmt.Println(result)
	}
}

func singleNumber(nums []int) int {
	countMap := map[int]int{}
	for _, value := range nums {
		_, result := countMap[value]
		if !result {
			countMap[value] = 1
		} else {
			countMap[value] += 1
		}
	}
	for key, val := range countMap {
		if val == 1 {
			return key
		}
	}
	return 0
}

func isPalindrome(x int) bool {
	strArray := strconv.Itoa(x)
	j := len(strArray) - 1
	for _, value := range strArray {
		if value != int32(strArray[j]) {
			return false
		}
		j--
	}
	return true
}
