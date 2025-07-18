package main

import "fmt"

// 常量-声明
func main() {

	// 第一种方式
	const I1 int = 0
	fmt.Println("I1 =", I1)

	// 第二种方式，自动推导类型
	const I2 = 0
	fmt.Println("I2 =", I2)

}
