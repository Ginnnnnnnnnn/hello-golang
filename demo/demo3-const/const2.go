package main

import "fmt"

// 第一种方式，一般用于全局
const (
	I1 int    = 1
	B1 bool   = true
	S1 string = "string"
)

// 多常量-声明
func main() {

	fmt.Println("I1 =", I1, "B1 =", B1, "S1 =", S1)

	// 第二种方式
	const I2, B2, S2 = 1, false, ""
	fmt.Println("I2 =", I2, "B2 =", B2, "S2 =", S2)

	// _ 的特殊作用
	const I3, B3, _ = 1, false, ""
	fmt.Println("I3 =", I3, "B3 =", B3)

}
