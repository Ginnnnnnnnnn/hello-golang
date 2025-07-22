package main

import "fmt"

// channel-监控多路
func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c1)
		}
		c2 <- 0
	}()

	test(c1, c2)
}

func test(c1, c2 chan int) {
	x, y := 0, 1
	for {
		select {
		case c1 <- x:
			x = x + y
		case <-c2:
			fmt.Println("结束")
			return
		}
	}
}
