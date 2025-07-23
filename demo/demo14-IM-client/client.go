package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	conn       net.Conn
	// 用户模式
	flag int
	// 私聊目标
	remoteName string
}

func (this *Client) Run() {
	for this.flag != 0 {
		for this.menu() != true {
		}
		switch this.flag {
		case 1:
			// 公聊模式
			this.publicCat()
			break
		case 2:
			// 私聊模式
			this.privateChat()
			break
		case 3:
			// 更改名称
			this.updateName()
			break
		case 4:
			// 用户列表
			this.selectUsers()
			break
		}
	}
}

// Menu 菜单
func (this *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改名称")
	fmt.Println("4.用户列表")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag < 0 || flag > 4 {
		fmt.Println("请输入正确菜单ID")
		return false
	}
	this.flag = flag
	return true
}

// 私聊模式
func (this *Client) privateChat() {
	fmt.Println(">>>请输入用户对象,exit退出")
	fmt.Scanln(&this.remoteName)
	this.publicCat()
}

// 公聊模式
func (this *Client) publicCat() {
	fmt.Println(">>>请输入聊天内容,exit退出")
	var chatMsg string
	fmt.Scanln(&chatMsg)
	// 判断消息内容是否为空,为空重新输入
	if chatMsg == "" {
		this.publicCat()
		return
	}
	// 判断退出
	if chatMsg == "exit" {
		this.remoteName = ""
		return
	}
	// 判断是否是私聊
	if this.remoteName != "" {
		chatMsg = "to|" + this.remoteName + "|" + chatMsg
	}
	// 发送消息
	_, err := this.conn.Write([]byte(chatMsg))
	if err != nil {
		fmt.Println(err)
	}
	// 递归调用
	this.publicCat()
}

// 更改名称
func (this *Client) updateName() {
	fmt.Println(">>>请输入用户名:")
	var newName string
	fmt.Scanln(&newName)
	sendMsg := "rename|" + newName
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(err)
	}
}

// 用户列表
func (this Client) selectUsers() {
	sendMsg := "who"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// DealResponse 处理响应
func (this *Client) DealResponse() {
	_, err := io.Copy(os.Stdout, this.conn)
	if err != nil {
		fmt.Println("服务器连接中断...")
	}
}

// NewClient 创建客户端
func NewClient(ServerIp string, serverPort int) *Client {
	// 创建结构体
	client := &Client{
		ServerIp:   ServerIp,
		ServerPort: serverPort,
		flag:       999,
	}
	// 建立连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ServerIp, serverPort))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	client.conn = conn
	// 返回
	return client
}
