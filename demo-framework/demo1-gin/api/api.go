package api

import (
	"demo1-gin/mod"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	// 绑定参数
	user := mod.User{}
	c.Bind(&user)
	// 响应
	c.JSON(200, user)
}
