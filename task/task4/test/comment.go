package main

import (
	"encoding/json"
	"fmt"
	"task4/req"
)

func commentAdd() {
	url := "http://127.0.0.1:8080/comment/add"
	commentAdd := req.CommentAdd{
		PostID:  3,
		Content: "评论1",
	}
	resp := Post(url, &commentAdd)
	fmt.Println(resp)
}

func commentList() {
	url := "http://127.0.0.1:8080/comment/list"
	resp := Get(url, map[string][]string{
		"post_id": {"3"},
	})
	body, _ := json.Marshal(resp)
	fmt.Println(string(body))
}
