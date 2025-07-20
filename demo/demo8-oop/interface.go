package main

import "fmt"

type Book struct {
	name string
}

// interface 万能数据类型，可以通过 .() 进行类型判断
func main() {

	myFunc(Book{})
	myFunc(0)
	myFunc("abc")
	myFunc(true)

}

func myFunc(arg interface{}) {
	fmt.Println("call myFunc")
	if val, ok := arg.(Book); ok {
		fmt.Printf("参数类型 = %T\n", val)
	} else if val, ok := arg.(int); ok {
		fmt.Printf("参数类型 = %T\n", val)
	} else if val, ok := arg.(string); ok {
		fmt.Printf("参数类型 = %T\n", val)
	} else if val, ok := arg.(bool); ok {
		fmt.Printf("参数类型 = %T\n", val)
	}
}
