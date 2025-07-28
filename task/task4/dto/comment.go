package dto

import (
	"fmt"
	"task4/models"
)

func CommentAdd(comment models.Comment) int64 {
	result := DB.Create(&comment)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}
	return result.RowsAffected
}

func CommentList(postId uint) []models.Comment {
	var comments []models.Comment
	result := DB.Preload("User").Where("post_id", postId).Find(&comments)
	if result.Error != nil {
		fmt.Println(result.Error)
		return comments
	}
	return comments
}
