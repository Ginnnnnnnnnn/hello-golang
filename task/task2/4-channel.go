package main

import (
	"fmt"
	"time"
)

func main() {
	// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	channel1()
	// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	channel2()
	// 等待
	time.Sleep(10 * time.Second)
}

func channel1() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	go func() {
		for {
			v := <-c
			fmt.Println("channel1 =", v)
		}
	}()
}

func channel2() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()
	go func() {
		for {
			v := <-c
			fmt.Println("channel2 =", v)
		}
	}()
}
