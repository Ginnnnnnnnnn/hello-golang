package main

import "fmt"

func main() {

	a, b := 100, 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	swap1(a, b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)

	fmt.Println("========================================")

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	swap2(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)

}

func swap1(a int, b int) {
	temp := a
	a = b
	b = temp
}

func swap2(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
