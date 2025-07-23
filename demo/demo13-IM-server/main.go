package main

// v1.1 构建基础 server
// v1.2 用户上线功能
// v1.3 用户消息广播机制
// v1.4 用户业务层封装
// v1.5 在线用户查询
// v1.6 修改用户名
// v1.7 超时强踢功能
// v1.8 私聊功能
func main() {

	// 开启监听
	myServer := NewServer("0.0.0.0", 8888)
	myServer.Start()

}
