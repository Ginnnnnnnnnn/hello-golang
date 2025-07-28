package controller

import (
	"github.com/gin-gonic/gin"
	"task4/req"
	"task4/service"
)

func CommentApi() {
	userGroup := Router.Group("/comment", auth)
	userGroup.POST("/add", CommentAdd)
	userGroup.GET("/list", CommentList)
}

func CommentAdd(c *gin.Context) {
	var commentAdd req.CommentAdd
	if err := c.ShouldBind(&commentAdd); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	commentAdd.UserID = getUserId(c)
	service.CommentAdd(commentAdd)
	Ok(c, nil)
}

func CommentList(c *gin.Context) {
	var commentList req.CommentList
	if err := c.ShouldBind(&commentList); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	result := service.CommentList(commentList.PostID)
	Ok(c, result)
}
