package main

import "fmt"

// 切片-使用
func main() {

	slice1 := []int{0}

	// 获取
	item := slice1[0]
	fmt.Println(item)

	// 添加
	slice1 = append(slice1, 1)

	// 截取，截取后的切片一样是引用传递
	slice2 := slice1[0:1]
	slice2 = slice1[:1]
	slice2 = slice1[0:]

	// copy
	slice3 := make([]int, len(slice2), 5)
	copy(slice3, slice2)

	// 遍历
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
	for index, value := range slice1 {
		fmt.Println(index, value)
	}

}
