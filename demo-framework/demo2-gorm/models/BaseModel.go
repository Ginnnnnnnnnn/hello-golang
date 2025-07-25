package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreateTime time.Time
	UpdateTime time.Time
	Deleted    gorm.DeletedAt
}

func GetBaseModel() BaseModel {
	return BaseModel{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Deleted:    gorm.DeletedAt{Time: time.Now(), Valid: true},
	}
}
