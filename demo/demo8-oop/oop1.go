package main

import "fmt"

type User struct {
	id   int
	name string
}

func (this *User) GetId() int {
	return this.id
}

func (this *User) SetId(id int) {
	this.id = id
}

func (this *User) GetName() string {
	return this.name
}

func (this *User) SetName(name string) {
	this.name = name
}

func (this *User) ToString() {
	fmt.Println("[ id =", this.id, " name =", this.name, "]")
}

// 面向对象-封装，通过 (this *User) 的方式来进行定义
func main() {
	user := User{
		id:   1,
		name: "小明",
	}
	user.ToString()
	user.SetName("小明（修改后）")
	user.ToString()
}
