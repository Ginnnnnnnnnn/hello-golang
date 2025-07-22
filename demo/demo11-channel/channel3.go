package main

import (
	"fmt"
)

// channel-关闭
// 关闭 channel 后，⽆无法向 channel 再发送数据
// 关闭 channel 后，可以继续从 channel 接收数据
func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// 关闭 channel
		close(c)
	}()

	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}

}
