package main

import "fmt"

// 数组-使用
func main() {
	array1 := [4]int{0, 1, 2, 3}
	array2 := [5]int{0, 1, 2, 3, 4}

	// 获取
	item := array1[0]
	fmt.Println(item)

	// 数组参数严格校验长度
	arrayType(array1)

	// 数组参数是值 copy 和 slice 存在差异，slice是引用传递（num[i]仅限此操作,改变长度的操作不生效）
	fmt.Printf("修改前 %v\n", array2)
	arrayParam(array2)
	fmt.Printf("修改后 %v\n", array2)
}

func arrayType(array [4]int) {
	fmt.Printf("类型 = %T\n", array)
}

func arrayParam(array [5]int) {
	array[0] = array[0] * -1
}
