package main

import "github.com/gin-gonic/gin"

// go-modules go 1.14 正式发布
// 环境设置
// go env -w GO111MODULE=on 设置 go-modules 开启
// go env -w GOPROXY=https://goproxy.cn,direct 设置镜像仓库
// go env -w GOPRIVATE="*.example.com" 设置私有仓库
// 初始化项目
// go mod init 项目名
// 常用命令
// go mod init 生成 go.mod 文件
// go mod download 下载 go.mod 文件中指明的所有依赖
// go mod tidy 整理现有的依赖
// go mod graph 查看现有的依赖结构
// go mod edit 编辑 go.mod 文件
// go mod vendor 导出项目所有的依赖到vendor目录
// go mod verify 校验一个模块是否被篡改过
// go mod why 查看为什么需要依赖某模块
// go list -m all 查看依赖版本。1.go list -m github.com/gin-gonic/gin 查看指定依赖 2.go list -m '...sql...' 模糊查询
func main() {
	router := gin.Default()
	router.Run(":8080")
}
