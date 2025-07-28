package resp

import "time"

type CommentList struct {
	Id        uint
	Content   string
	CreatedAt time.Time
	User      UserDetail
}
