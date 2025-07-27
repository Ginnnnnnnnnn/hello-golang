package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	Id      int64  `gorm:"primarykey;comment:评论ID"`
	PostId  int64  `gorm:"type:bigint(20);not null;comment:文章ID"`
	Content string `gorm:"type:varchar(1024);not null;comment:评论内容"`
}

func (c Comment) TableName() string {
	return "comments"
}

func (c Comment) TableComment() string {
	return "评论表"
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	// 查询信息
	tx.First(c)
	// 查询数量
	var count int64
	result := tx.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	count--
	// 更新数据
	commentStatus := 0
	if count > 0 {
		commentStatus = 1
	}
	result = tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", commentStatus)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return
}
