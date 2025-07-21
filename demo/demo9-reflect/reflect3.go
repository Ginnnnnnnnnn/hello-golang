package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"名字"`
	Sex  string `info:"sex" doc:"性别"`
}

// 反射-结构体-标签
func main() {

	// 反射获取结构体标签，需要传递指针
	resume := Resume{
		Name: "小明",
		Sex:  "男",
	}
	reflectTag(&resume)

}

func reflectTag(arg interface{}) {

	// 反射获取结构体标签
	argElem := reflect.TypeOf(arg).Elem()
	for i := 0; i < argElem.NumField(); i++ {
		info := argElem.Field(i).Tag.Get("info")
		doc := argElem.Field(i).Tag.Get("doc")
		fmt.Println("info =", info, "doc =", doc)
	}

}
