package main

import (
	"demo1-gin/aop"
	"demo1-gin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin
// 1.基本路由
// 2.参数绑定 和 参数校验
// 3.响应格式 和 重定向
// 4.静态资源
// 5.中间件: 全局中间件 分组中间件 中间件
// 6.模板引擎
func main() {
	router := gin.Default()
	// 路由
	router.Any("/test", aop.Aop, api.Test)
	// 重定向
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	// 静态资源
	router.Static("/static", "./assets")
	router.StaticFile("/staticFile/txt1", "./assets/txt1.txt")
	// 启动监听,默认 0.0.0.0:8080
	router.Run()
}
