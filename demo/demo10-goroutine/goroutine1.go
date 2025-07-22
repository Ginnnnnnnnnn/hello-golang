package main

import (
	"fmt"
	"time"
)

// goroutine 主线程退出，所有 goroutine 也会退出
func main() {

	go newGoroutine()
	i := 0
	for {
		i++
		fmt.Println("main goroutine i =", i)
		time.Sleep(1 * time.Second)
	}

}

func newGoroutine() {
	i := 0
	for {
		i++
		fmt.Println("new goroutine i =", i)
		time.Sleep(1 * time.Second)
	}
}
