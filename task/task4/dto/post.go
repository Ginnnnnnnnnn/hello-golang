package dto

import (
	"fmt"
	"task4/models"
)

func PostAdd(post models.Post) int64 {
	result := DB.Create(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}
	return result.RowsAffected
}

func PostList() *[]models.Post {
	var posts []models.Post
	result := DB.Find(&posts)
	if result.Error != nil {
		fmt.Println(result.Error)
		return &posts
	}
	return &posts
}

func PostFindDetailById(id uint) (models.Post, int64) {
	var post models.Post
	result := DB.Preload("User").Where("id = ?", id).First(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return post, result.RowsAffected
}

func PostUpdate(post models.Post) int64 {
	result := DB.Model(&post).Where("id = ?", post.ID).Updates(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}
	return result.RowsAffected
}

func PostDelete(id uint) int64 {
	result := DB.Where("id = ?", id).Delete(&models.Post{})
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}
	return result.RowsAffected
}
