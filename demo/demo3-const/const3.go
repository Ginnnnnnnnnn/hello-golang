package main

import (
	"fmt"
	"unsafe"
)

// 常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过。
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
	// d = test()
)

func test() int {
	return 1
}

// iota
const (
	_     = iota              // 0
	iota1                     // 1
	iota2                     // 2
	iota3                     // 3
	iota4 = iota - iota2 + 10 // 3 - 2 + 10 = 11
	iota5                     // 4 - 2 + 10 = 12
	iota6                     // 5 - 2 + 10 = 13
)
const (
	iota11, iota111 = iota + 1, iota + 2 // 1, 2
	iota22, iota222                      // 2, 3
	iota33, iota333                      // 3, 4
)

// 常量-杂项
func main() {

	fmt.Println("a =", 1, " b =", b, "c =", c)

}
