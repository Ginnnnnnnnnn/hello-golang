package main

import (
	"demo2-gorm/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// gorm
// 1.连接数据库
// 2.定义模型
// 3.CRUD
func main() {
	// 初始化 mysql 连接
	db, err := initMysql()
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	db.DB()
	// CRUD
	user := models.User{Id: 1, Name: "小明", BaseModel: models.GetBaseModel()}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("插入失败", result.Error)
		return
	}
}

func initMysql() (db *gorm.DB, err error) {
	user := "xzdev"
	password := "123456"
	ip := "47.96.154.175"
	port := "3306"
	dbName := "test"
	// 创建连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
