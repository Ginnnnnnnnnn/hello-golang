package main

import (
	"encoding/json"
	"fmt"
	"task4/req"
)

// 添加文章
func postAdd() {
	url := "http://127.0.0.1:8080/post/add"
	post := req.PostAdd{
		Title:   "文章-1",
		Content: "文章内容-1",
	}
	resp := Post(url, &post)
	fmt.Println(resp)
}

// 文章列表
func postList() {
	url := "http://127.0.0.1:8080/post/list"
	resp := Get(url, nil)
	body, _ := json.Marshal(resp)
	fmt.Println(string(body))
}

// 文章详情
func postDetail() {
	url := "http://127.0.0.1:8080/post/detail"
	resp := Get(url, map[string][]string{
		"id": {"3"},
	})
	body, _ := json.Marshal(resp)
	fmt.Println(string(body))
}

// 文章更新
func postUpdate() {
	url := "http://127.0.0.1:8080/post/update"
	post := req.PostUpdate{
		Id:      3,
		Title:   "文章-1(已修改)",
		Content: "文章内容-1(已修改)",
	}
	resp := Post(url, &post)
	fmt.Println(resp)
}

// 文章删除
func postDelete() {
	url := "http://127.0.0.1:8080/post/delete"
	post := req.PostDelete{
		Id: 4,
	}
	resp := Post(url, &post)
	fmt.Println(resp)
}
