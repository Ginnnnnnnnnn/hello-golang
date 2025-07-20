package main

import "fmt"

// map-使用
func main() {

	map1 := make(map[int]string)

	// 添加
	map1[1] = "a"
	map1[2] = "b"
	map1[3] = "c"
	map1[4] = "abc"

	// 遍历
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	// 删除
	delete(map1, 3)

	// 修改
	map1[4] = "已修改"

}
