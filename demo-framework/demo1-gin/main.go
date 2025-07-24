package main

import (
	"demo1-gin/api"
	"github.com/gin-gonic/gin"
)

// gin
func main() {
	router := gin.Default()
	router.GET("/test", api.Test)
	router.Run()
}
