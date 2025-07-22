package main

import "github.com/gin-gonic/gin"

// go-modules go 1.14 正式发布
// 环境设置
// go env -w GO111MODULE=on 设置 go-modules 开启
// go env -w GOPROXY=https://goproxy.cn,direct 设置镜像仓库
// go env -w GOPRIVATE="*.example.com" 设置私有仓库
// 初始化项目
// go mod init 项目名
func main() {
	router := gin.Default()
	router.Run(":8080")
}
