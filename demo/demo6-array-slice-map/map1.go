package main

import "fmt"

// map-声明
func main() {

	// 第一种
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["1"] = "a"
	map1["2"] = "b"
	map1["3"] = "c"
	fmt.Printf("map1 = %v\n", map1)

	// 第二种
	map2 := make(map[string]string)
	map2["1"] = "a"
	map2["2"] = "b"
	map2["3"] = "c"
	fmt.Printf("map2 = %v\n", map2)

	// 第三种
	mp3 := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
	}
	fmt.Printf("mp3 = %v\n", mp3)

}
