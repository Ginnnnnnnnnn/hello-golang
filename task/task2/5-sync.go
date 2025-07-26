package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	lock sync.Mutex
	num1 int

	num2 int32
)

func main() {
	// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				lock.Lock()
				num1++
				lock.Unlock()
			}
		}()
	}
	// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&num2, 1)
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(num1)
	fmt.Println(num2)
}
