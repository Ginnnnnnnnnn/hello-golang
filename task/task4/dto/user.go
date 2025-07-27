package dto

import (
	"log"
	"task4/models"
)

func UserCountByUsername(username string) *models.User {
	var user models.User
	result := DB.Model(&models.User{}).Where("username = ?", username).First(&user)
	if result.RowsAffected == 0 {
		return nil
	}
	return &user
}

func UserCountByEmail(email string) (count int64) {
	DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return
}

func UserAdd(user *models.User) int64 {
	result := DB.Create(user)
	if result.Error != nil {
		log.Fatalf("添加用户错误: %v", result.Error)
		return 0
	}
	return result.RowsAffected
}
