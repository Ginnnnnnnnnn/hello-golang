package main

import "fmt"

// channel-持续读取数据
func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// 关闭 channel
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}

}
