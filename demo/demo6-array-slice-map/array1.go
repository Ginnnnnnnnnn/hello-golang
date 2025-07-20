package main

import "fmt"

// 数组-声明
func main() {

	// 第一种声明
	var array1 [3]int
	array1[0] = 0
	array1[1] = 1
	array1[2] = 2
	fmt.Println(array1)

	// 第二种声明
	array2 := [4]int{0, 1, 2, 3}
	fmt.Println(array2)

	// 第一种遍历
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	// 第二种遍历
	for index, value := range array2 {
		fmt.Println("index =", index, "value =", value)
	}

}
