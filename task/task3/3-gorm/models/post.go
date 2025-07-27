package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Post struct {
	Id            int64     `gorm:"primarykey;comment:文章ID"`
	UserId        int64     `gorm:"type:bigint(20);not null;comment:用户ID"`
	Title         string    `gorm:"type:varchar(255);not null;comment:文章标题"`
	CommentStatus int8      `gorm:"type:tinyint;not null;default:0;comment:评论状态 0-无评论 1-有评论"`
	Comments      []Comment `gorm:"foreignKey:PostId"`
}

func (p Post) TableName() string {
	return "posts"
}

func (p Post) TableComment() string {
	return "文章表"
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 查询数量
	var count int64
	result := tx.Model(&Post{}).Where("user_id = ?", p.UserId).Count(&count)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	// 更新用户文章数量
	result = tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", count)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return
}
