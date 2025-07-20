package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this Cat) Sleep() {
	fmt.Println(this.color, "Cat Sleep")
}

func (this Cat) GetColor() string {
	return this.color
}

func (this Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this Dog) Sleep() {
	fmt.Println(this.color, "Dog Sleep")
}

func (this Dog) GetColor() string {
	return this.color
}

func (this Dog) GetType() string {
	return "Dog"
}

// 面向对象-多态，实现接口的全部方法，通过 &Cat{} 实现
func main() {

	animal1 := &Cat{color: "red"}
	showSleep(animal1)

	animal2 := &Dog{color: "pink"}
	showSleep(animal2)

}

func showSleep(animal AnimalIF) {
	animal.Sleep()
}
