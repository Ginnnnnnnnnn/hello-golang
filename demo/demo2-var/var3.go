package main

import "fmt"

// 变量-类型
func main() {

	// int 整数
	var (
		i1 int   = 1 // int 是一种有符号整数类型，其大小至少为32位。然而，它是一个独立的类型，并不是（例如）int32的别名。
		i2 int8  = 2
		i3 int16 = 3
		i4 int32 = -2
		i5 int64 = -1
	)
	fmt.Println("i1 =", i1, "i2 =", i2, "i3 =", i3, "i4 =", i4, "i5 =", i5)

	// uint 正整数
	var (
		ui1 uint   = 1 // uint 是一种无符号整数类型，其大小至少为 32 位。然而，它是一个独立的类型，并不是（例如）uint32的别名。
		ui2 uint8  = 2
		ui3 uint16 = 3
		ui4 uint32 = 4
		ui5 uint64 = 5
	)
	fmt.Println("ui1 =", ui1, "ui2 =", ui2, "ui3 =", ui3, "ui4 =", ui4, "ui5 =", ui5)

	// float 小数
	var (
		f1 float32 = 3.14
		f2 float64 = 3.1415
	)
	fmt.Println("f1 =", f1, "f2 =", f2)

	// complex 复数
	var (
		c1 complex64  = 1 + 2i
		c2 complex128 = 1 + 2i + 3i
	)
	fmt.Println("c1 =", c1, "c2 =", c2)

	// bool 布尔
	var b bool = true
	fmt.Println("b = ", b)

	// string 字符串
	var s string = "hello world"
	fmt.Println("s =", s)
}
