package dto

import (
	"fmt"
	"gorm/bo"
	"gorm/models"
)

func PostCreate(post *models.Post) {
	result := DB.Create(post)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func PostMaxByComments() *bo.PostMaxByComments {
	var postMaxByComments bo.PostMaxByComments
	err := DB.Model(&models.Post{}).
		Select("posts.id, posts.user_id, posts.title, COUNT( DISTINCT comments.ID ) AS CommentsCount").
		Joins("LEFT JOIN comments comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("CommentsCount DESC").
		First(&postMaxByComments).
		Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &postMaxByComments
}
