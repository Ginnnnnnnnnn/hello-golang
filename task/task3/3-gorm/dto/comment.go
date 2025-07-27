package dto

import (
	"fmt"
	"gorm/models"
)

func CommentDelete(comment *models.Comment) {
	result := DB.Delete(comment)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
