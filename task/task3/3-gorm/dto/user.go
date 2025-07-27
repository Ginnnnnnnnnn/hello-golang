package dto

import "gorm/models"

func UserFindById(id int64) *models.User {
	var user models.User
	DB.Preload("Posts").Preload("Posts.Comments").First(&user, id)
	return &user
}
