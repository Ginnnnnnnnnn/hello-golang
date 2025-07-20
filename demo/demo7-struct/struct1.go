package main

import "fmt"

// 声明一个类型
type myint int

// User 声明一个结构体
type User struct {
	id   int
	name string
}

// 结构体-声明
func main() {

	var myInt myint
	fmt.Printf("myint = %T\n", myInt)

	var user User
	user.id = 1
	user.name = "小明"
	fmt.Printf("user = %v\n", user)

	setUser1(user)
	fmt.Printf("user = %v\n", user)
	setUser2(&user)
	fmt.Printf("user = %v\n", user)
}

func setUser1(user User) {
	user.id = 2
}

func setUser2(user *User) {
	user.id = 2
}
