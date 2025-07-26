package main

import (
	"fmt"
	"time"
)

func main() {
	// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	goroutine1()
	// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	funcs := []func(){
		func() {
			startTime := time.Now()
			time.Sleep(1 * time.Second)
			endTime := time.Now()
			fmt.Println("任务1 =", endTime.Sub(startTime))
		},
		func() {
			startTime := time.Now()
			time.Sleep(2 * time.Second)
			endTime := time.Now()
			fmt.Println("任务2 =", endTime.Sub(startTime))
		},
		func() {
			startTime := time.Now()
			time.Sleep(3 * time.Second)
			endTime := time.Now()
			fmt.Println("任务3 =", endTime.Sub(startTime))
		},
	}
	goroutine2(funcs)
	// 等待
	time.Sleep(10 * time.Second)
}

func goroutine1() {
	go func() {
		for i := 1; i < 10; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()
	go func() {
		for i := 1; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
}

func goroutine2(funcs []func()) {
	for _, value := range funcs {
		go value()
	}
}
