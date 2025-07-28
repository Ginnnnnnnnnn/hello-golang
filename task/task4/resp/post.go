package resp

import "time"

type PostList struct {
	Id        uint
	Title     string
	Content   string
	CreatedAt time.Time
}

type PostDetail struct {
	Id         uint
	Title      string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserDetail UserDetail
}
