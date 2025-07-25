package models

import (
	"fmt"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	Id   int64 `gorm:"primaryKey,autoIncrement"`
	Name string
	// 嵌套结构
	// 对于匿名字段，GORM 会将其字段包含在父结构体中
	// 对于正常的结构体字段，你也可以通过标签 embedded 将其嵌入
	BaseModel BaseModel `gorm:"embedded"`
}

// TableName 设置表名
func (this *User) TableName() string {
	return "user"
}

// BeforeDelete 删除前钩子
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("Before Delete")
	return
}
