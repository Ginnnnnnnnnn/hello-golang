package service

import (
	"gorm.io/gorm"
	"task4/dto"
	"task4/models"
	"task4/req"
	"task4/resp"
)

func PostAdd(postAdd req.PostAdd) bool {
	post := models.Post{
		Title:   postAdd.Title,
		Content: postAdd.Content,
		UserID:  postAdd.UserID,
	}
	result := dto.PostAdd(post)
	return result > 0
}

func PostList() *[]resp.PostList {
	var postList []resp.PostList
	posts := dto.PostList()
	for _, post := range *posts {
		postList = append(postList, resp.PostList{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
		})
	}
	return &postList
}

func PostFindDetailById(id uint) *resp.PostDetail {
	post, count := dto.PostFindDetailById(id)
	if count < 1 {
		return nil
	}
	postDetail := resp.PostDetail{
		Id:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		UserDetail: resp.UserDetail{
			Id:       post.User.ID,
			Username: post.User.Username,
			Email:    post.User.Email,
		},
	}
	return &postDetail
}

func PostUpdate(postAdd req.PostUpdate) bool {
	post := models.Post{
		Model: gorm.Model{
			ID: postAdd.Id,
		},
		Title:   postAdd.Title,
		Content: postAdd.Content,
	}
	result := dto.PostUpdate(post)
	return result > 0
}

func PostDelete(id uint) bool {
	result := dto.PostDelete(id)
	return result > 0
}
