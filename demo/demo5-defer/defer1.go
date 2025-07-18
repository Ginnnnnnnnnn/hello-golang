package main

import "fmt"

// defer
func main() {
	// 执行顺序 3 -> 2 -> 1
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
}
