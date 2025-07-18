package main

import "fmt"

// 第一种方式，一般用于全局
var (
	i1 int
	b1 bool
	s1 string
)

// 多变量-声明
func main() {

	fmt.Println("i1 =", i1, "b1 =", b1, "s1 =", s1)

	// 第二种方式
	var i2, b2, s2 int
	fmt.Println("i2 =", i2, "b2 =", b2, "s2 =", s2)

	// 第三种方式
	var i3, b3, s3 = 1, false, ""
	fmt.Println("i3 =", i3, "b3 =", b3, "s3 =", s3)

	// 第四种方式
	i4, b4, s4 := 1, false, ""
	fmt.Println("i4 =", i4, "b4 =", b4, "s4 =", s4)

	// _ 的特殊作用
	i5, b5, _ := 1, false, ""
	fmt.Println("i5 =", i5, "b5 =", b5)

}
