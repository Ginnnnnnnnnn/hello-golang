package models

import "time"

type BaseModel struct {
	CreateTime time.Time
	UpdateTime time.Time
	Deleted    int8
}

func GetBaseModel() BaseModel {
	return BaseModel{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Deleted:    0,
	}
}
