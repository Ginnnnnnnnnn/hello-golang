package main

import "fmt"

// pair 变量的结构
// 变量
// type	value <- 这一对叫做 pair
// static type		concrete type
// int、string...	interface所指向的具体数据类型，系统看得见的类型
func main() {

	// pair <type: string, value: hello>
	var s1 string = "hello"
	fmt.Println(s1)

	// pair <type: string, value: hello>
	var if1 interface{} = s1
	fmt.Println(if1)

}
