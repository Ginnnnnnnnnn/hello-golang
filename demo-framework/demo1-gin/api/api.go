package api

import (
	"demo1-gin/mod"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	fmt.Println("func test")
	// 绑定参数
	//var user mod.User
	user := mod.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	// 响应
	c.JSON(200, user)
}
