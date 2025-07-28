package service

import (
	"task4/dto"
	"task4/models"
	"task4/req"
	"task4/resp"
)

func CommentAdd(commentAdd req.CommentAdd) bool {
	comment := models.Comment{
		PostID:  commentAdd.PostID,
		UserID:  commentAdd.UserID,
		Content: commentAdd.Content,
	}
	count := dto.CommentAdd(comment)
	return count > 0
}

func CommentList(postId uint) []resp.CommentList {
	comments := dto.CommentList(postId)
	var result []resp.CommentList
	for _, comment := range comments {
		result = append(result, resp.CommentList{
			Id:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			User: resp.UserDetail{
				Id:       comment.User.ID,
				Username: comment.User.Username,
				Email:    comment.User.Email,
			},
		})
	}
	return result
}
