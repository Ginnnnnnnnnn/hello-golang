package controller

import (
	"github.com/gin-gonic/gin"
	"task4/req"
)

func PostAdd(c *gin.Context) {
	var postAdd req.PostAdd
	if err := c.ShouldBind(&postAdd); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	Ok(c, nil)
}
