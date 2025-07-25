package main

import "fmt"

func main() {
	// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	i := 10
	fmt.Println("前 =", i)
	pointer1(&i)
	fmt.Println("后 =", i)
	// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	u := []uint{1, 2, 3}
	fmt.Println("前 =", u)
	pointer2(u)
	fmt.Println("后 =", u)
}

func pointer1(i *int) {
	*i = *i + 10
}

func pointer2(u []uint) {
	for i := 0; i < len(u); i++ {
		u[i] = u[i] * 2
	}
}
