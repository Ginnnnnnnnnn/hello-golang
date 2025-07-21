package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
}

func (this User) Call() {
	fmt.Println("user call ...")
}

// 反射-结构体
func main() {

	// 结构体
	user := User{
		id:   1,
		name: "小明",
	}
	reflectStruct(user)

}

func reflectStruct(arg interface{}) {

	// TypeOf 用来动态获取参数的类型，如果参数为空返回 nil
	argType := reflect.TypeOf(arg)
	fmt.Println("type =", argType.Name())

	// ValueOf 用来获取参数的值，如果参数为空返回 0
	argValue := reflect.ValueOf(arg)
	fmt.Println("value =", argValue)

	// 反射获取字段
	for i := 0; i < argType.NumField(); i++ {
		fieldType := argType.Field(i)
		fieldValue := argValue.Field(i)
		fmt.Println("fieldType =", fieldType.Name, "fieldValue =", fieldValue)
	}

	// 反射获取方法
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Println("methodName =", method.Name, "methodType =", method.Type)
	}

}
