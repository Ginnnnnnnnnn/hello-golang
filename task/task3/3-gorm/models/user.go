package models

type User struct {
	Id        int64  `gorm:"primarykey;comment:用户ID"`
	Name      string `gorm:"type:varchar(255);not null;comment:用户名称"`
	PostCount int64  `gorm:"type:bigint(20);not null;default:0;comment:文章数量"`
	Posts     []Post `gorm:"foreignKey:UserId"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) TableComment() string {
	return "用户表"
}
