package main

import (
	"fmt"
	"time"
)

// 有缓存 channel
// 缓存有空间-不阻塞
// 缓存无空间-阻塞
func main() {

	c := make(chan int, 3)
	fmt.Println("channel len =", len(c), "cap =", cap(c))

	go func() {
		defer fmt.Println("go程结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("go程正在发送 =", i, "len =", len(c), "cap =", cap(c))
		}
	}()

	time.Sleep(3 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("main程正在读取 =", num, "len =", len(c), "cap =", cap(c))
	}

	fmt.Println("main程结束")
}
