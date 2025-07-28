package req

type PostAdd struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

type PostDetail struct {
	Id uint `form:"id" binding:"required"`
}

type PostUpdate struct {
	Id      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostDelete struct {
	Id uint `json:"id" binding:"required"`
}
