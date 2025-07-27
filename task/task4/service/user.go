package service

import (
	"task4/dto"
	"task4/models"
	"task4/req"
)

func UserFindByUsername(username string) *models.User {
	return dto.UserCountByUsername(username)
}

func UserCountByEmail(username string) (count int64) {
	count = dto.UserCountByEmail(username)
	return
}

func UserAdd(userAdd *req.UserAdd) bool {
	user := models.User{
		Username: userAdd.Username,
		Password: userAdd.Password,
		Email:    userAdd.Email,
	}
	result := dto.UserAdd(&user)
	return result > 0
}
