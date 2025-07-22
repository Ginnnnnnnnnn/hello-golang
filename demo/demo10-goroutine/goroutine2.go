package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine 匿名函数
func main() {

	// 匿名无参无返回值函数
	go func() {
		defer fmt.Println("defer A")
		func() {
			defer fmt.Println("defer B")
			// 退出 goroutine
			runtime.Goexit()
			fmt.Println("B")
		}() // () 指调用
		fmt.Println("A")
	}()

	// 匿名有参有返回值函数
	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(1, 2)

	for {
		time.Sleep(1 * time.Second)
	}

}
