package main

import "fmt"

// 变量-声明
func main() {

	// 第一种方式
	var i1 int
	fmt.Println("i1 =", i1)

	// 第二种方式
	var i2 int = 0
	fmt.Println("i2 =", i2)

	// 第三种方式，自动推导类型
	var i3 = 0
	fmt.Println("i3 =", i3)

	// 第四种方式，自动推导类型，省略 var
	i4 := 0
	fmt.Println("i4 =", i4)

}
