package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户
	MapLock   sync.RWMutex
	OnlineMap map[string]*User

	// 消息广播的 channel
	MessageC chan string
}

// Start
// 启动监听
func (this *Server) Start() {
	// 启动监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("启动服务监听失败", err)
	}
	// 结束关闭
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("关闭服务监听失败", err)
		}
	}(listener)
	// 启动 messageC
	go this.ListenMessageC()
	// 监听连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接建立失败", err)
		}
		go this.handler(conn)
	}
}

// handler
// 处理连接
func (this *Server) handler(conn net.Conn) {
	// 创建用户
	user := NewUser(conn, this)
	// 用户上线
	user.Online()
	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				user.Offline()
				return
			}
			if n == 0 {
				user.Offline()
				return
			}
			// 读取消息内容（去掉'\n'）
			msg := string(buf[:n-1])
			// 广播
			user.DoMessage(msg)
		}
	}()
	// 当前 handler 阻塞
	select {}
}

// BroadCast
// 写入广播 MessageC
// user 发送者
// msg 消息内容
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.MessageC <- sendMsg
}

// ListenMessageC
// 监听 MessageC 发送在线用户
func (this *Server) ListenMessageC() {
	for {
		msg := <-this.MessageC

		// 把消息发给所有在线用户
		this.MapLock.RLock()
		for _, val := range this.OnlineMap {
			val.C <- msg
		}
		this.MapLock.RUnlock()
	}
}

// NewServer
// 创建对象
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		MessageC:  make(chan string),
	}
	return server
}
