package main

import "fmt"

// Man 父类
type Man struct {
	id   int
	name string
}

func (this *Man) Eat() {
	fmt.Println("man eat ...")
}

func (this *Man) Work() {
	fmt.Println("man eat ...")
}

// SuperMan 子类
type SuperMan struct {
	Man
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("superman eat ...")
}

// 面向对象-继承，在结构体中添加父结构体
func main() {
	sm := SuperMan{
		Man: Man{
			id:   1,
			name: "小明",
		},
		level: 1,
	}
	sm.Work()
	sm.Eat()
}
