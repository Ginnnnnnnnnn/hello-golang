package req

type CommentAdd struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

type CommentList struct {
	PostID uint `form:"post_id" binding:"required"`
}
