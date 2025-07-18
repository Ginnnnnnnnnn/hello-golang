package main

import "fmt"

// init 函数，优先被执行
func init() {
	fmt.Println("init")
}

// 主函数，在 init 函数之后执行
func main() {

	fmt.Println("main")

	// 单参数，单返回值
	a1 := func1("abc")
	fmt.Println(a1)

	// 多参数，多返回值，返回值无名称
	a2, b2 := func2("abc", "def")
	fmt.Println(a2, b2)

	// 多参数，多返回值，返回值有名称
	a3, b3 := func3("abc", "def")
	fmt.Println(a3, b3)

}

func func1(a string) string {
	return a
}

func func2(a, b string) (string, string) {
	return a, b
}

func func3(a string, b string) (r1, r2 string) {
	r1 = a
	r2 = b
	return
}
