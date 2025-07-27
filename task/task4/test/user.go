package main

import (
	"fmt"
	"task4/req"
)

// 用户注册
func register() {
	url := "http://127.0.0.1:8080/user/register"
	user := req.UserAdd{
		Username: "gin-1",
		Password: "123456",
		Email:    "2440016524@qq.com",
	}
	resp := Post(url, &user)
	fmt.Println(resp)
}

func login() {
	url := "http://127.0.0.1:8080/user/login"
	user := req.UserLogin{
		Username: "gin",
		Password: "123456",
	}
	resp := Post(url, &user)
	fmt.Println(resp)
}
