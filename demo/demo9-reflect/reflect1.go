package main

import (
	"fmt"
	"reflect"
)

// 反射-基础类型
func main() {

	// 基础类型
	var i int = 1
	reflectBaseType(i)
}

func reflectBaseType(arg interface{}) {

	// TypeOf 用来动态获取参数的类型，如果参数为空返回 nil
	fmt.Println("type =", reflect.TypeOf(arg))

	// ValueOf 用来获取参数的值，如果参数为空返回 0
	fmt.Println("value =", reflect.ValueOf(arg))

}
