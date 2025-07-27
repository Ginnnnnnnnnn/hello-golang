package main

import (
	"fmt"
	"task4/req"
)

// 用户注册
func postAdd() {
	url := "http://127.0.0.1:8080/post/add"
	user := req.PostAdd{
		Title:   "文章-1",
		Content: "文章内容-1",
	}
	resp := Post(url, &user)
	fmt.Println(resp)
}
