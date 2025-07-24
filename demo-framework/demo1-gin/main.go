package main

import (
	"demo1-gin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin
func main() {
	router := gin.Default()
	// 路由
	router.Any("/test", api.Test)
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
