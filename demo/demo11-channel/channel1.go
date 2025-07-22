package main

import (
	"fmt"
	"time"
)

// 无缓存 channel
// main程未读取到数据时-阻塞
// go程发送的数据未被读取走时-阻塞
func main() {

	// 定义 channel
	c := make(chan int)

	// 开启 goroutine
	go func() {
		defer fmt.Println("go程结束")
		// 666 发送给 c
		fmt.Println("go程开始发送")
		c <- 666
	}()
	time.Sleep(1 * time.Second)

	// 从 c 中读取数据赋值给 num
	fmt.Println("main程开始读取")
	num := <-c
	fmt.Println(num)
	time.Sleep(1 * time.Second)
	fmt.Println("main程结束")

	for {
		time.Sleep(3 * time.Second)
	}
}
