package main

import (
	"fmt"
	"gorm/dto"
	"gorm/models"
	"strconv"
	"time"
)

func main() {
	dto.CreateTable()
	// 题目1：模型定义
	gorm1()
	// 题目2：关联查询
	gorm2()
	// 题目3：钩子函数
	gorm3()
}

// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
func gorm1() {
	// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	user := &models.User{
		Id:   1,
		Name: "用户",
		Posts: []models.Post{
			{
				Id:     1,
				UserId: 1,
				Title:  "文章",
				Comments: []models.Comment{
					{
						Id:      1,
						PostId:  1,
						Content: "评论",
					},
				},
			},
		},
	}
	fmt.Println(user)
	// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
	dto.CreateTable()
}

// 基于上述博客系统的模型定义。
func gorm2() {
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	user := dto.UserFindById(1)
	fmt.Println(user)
	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	post := dto.PostMaxByComments()
	fmt.Println(post)
}

// 继续使用博客系统的模型。
func gorm3() {
	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	post := models.Post{
		UserId: 1,
		Title:  "文章" + strconv.Itoa(int(time.Now().Unix())),
	}
	dto.PostCreate(&post)
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	comment := models.Comment{
		Id: 1,
	}
	dto.CommentDelete(&comment)
}
