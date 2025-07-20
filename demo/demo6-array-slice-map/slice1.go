package main

import "fmt"

// 切片-声明
func main() {

	// 第一种
	slice1 := []int{0, 1, 2}
	fmt.Println(slice1)

	// 第二种
	var slice2 []int
	slice2 = make([]int, 3)
	fmt.Println(slice2)

	// 第三种
	var slice3 []int = make([]int, 3)
	fmt.Println(slice3)

	// 第四种
	slice4 := make([]int, 3)
	fmt.Println(slice4)

}
