package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// Online 用户上线
func (this *User) Online() {
	// 将用户加入 onlineMap
	this.server.MapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.MapLock.Unlock()

	// 广播用户上线
	this.server.BroadCast(this, "已上线")
}

// Offline 用户下线
func (this *User) Offline() {
	// 将用户删除 onlineMap
	this.server.MapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.MapLock.Unlock()

	// 广播用户下线
	this.server.BroadCast(this, "已下线")
}

// DoMessage 发送消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询在线用户指令
		this.server.MapLock.Lock()
		for _, val := range this.server.OnlineMap {
			onlineMsg := "[" + val.Addr + "]" + val.Name + ":" + "在线..."
			this.SendMsg(onlineMsg)
		}
		this.server.MapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 修改用户名
		newName := strings.Split(msg, "|")[1]
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名已存在")
		} else {
			this.server.MapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.MapLock.Unlock()

			this.Name = newName
			this.SendMsg("您的名称已修改为:" + newName)
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}

// listenMessage
// 监听 channel，有消息发给客户端
func (this *User) listenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg))
	}
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// NewUser
// 创建对象
func NewUser(conn net.Conn, server *Server) *User {
	user := &User{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听
	go user.listenMessage()

	return user
}
