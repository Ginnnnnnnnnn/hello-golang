package main

import (
	"flag"
	"fmt"
)

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务IP地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务端口地址(默认8888)")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>连接服务器失败...")
		return
	}
	fmt.Println(">>>连接服务器成功...")
	go client.DealResponse()
	client.Run()
}
