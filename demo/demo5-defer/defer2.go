package main

import "fmt"

// defer
func main() {
	// 执行顺序 return -> defer
	test()
}

func test() int {
	defer deferFunc()
	return returnFunc()
}

func returnFunc() int {
	fmt.Println("return func")
	return 0
}

func deferFunc() {
	fmt.Println("defer func")
}
