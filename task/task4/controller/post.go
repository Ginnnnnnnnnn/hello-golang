package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"task4/req"
	"task4/service"
)

func PostApi() {
	userGroup := Router.Group("/post", auth)
	userGroup.POST("/add", PostAdd)
	userGroup.GET("/list", PostList)
	userGroup.GET("/detail", PostDetail)
	userGroup.POST("/update", PostUpdate)
	userGroup.POST("/delete", PostDelete)
}

func PostAdd(c *gin.Context) {
	var postAdd req.PostAdd
	if err := c.ShouldBind(&postAdd); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	postAdd.UserID = getUserId(c)
	service.PostAdd(postAdd)
	Ok(c, nil)
}

func PostList(c *gin.Context) {
	posts := service.PostList()
	Ok(c, posts)
}

func PostDetail(c *gin.Context) {
	var postDetail req.PostDetail
	if err := c.ShouldBind(&postDetail); err != nil {
		fmt.Println(err)
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	post := service.PostFindDetailById(postDetail.Id)
	Ok(c, post)
}

func PostUpdate(c *gin.Context) {
	var postUpdate req.PostUpdate
	if err := c.ShouldBind(&postUpdate); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	// 查询信息
	postDetail := service.PostFindDetailById(postUpdate.Id)
	if postDetail == nil {
		Err(c, PARAM_ERR, "文章信息不存在")
		return
	}
	id := postDetail.Id
	if id != getUserId(c) {
		Err(c, PARAM_ERR, "文章作者才能修改")
		return
	}
	service.PostUpdate(postUpdate)
	Ok(c, nil)
}

func PostDelete(c *gin.Context) {
	var postDelete req.PostDelete
	if err := c.ShouldBind(&postDelete); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	// 查询信息
	postDetail := service.PostFindDetailById(postDelete.Id)
	if postDetail == nil {
		Err(c, PARAM_ERR, "文章信息不存在")
		return
	}
	id := postDetail.Id
	if id != getUserId(c) {
		Err(c, PARAM_ERR, "文章作者才能删除")
		return
	}
	service.PostDelete(postDelete.Id)
	Ok(c, nil)
}
