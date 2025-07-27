package dto

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm/models"
)

var DB *gorm.DB

func init() {
	user := "root"
	password := "123456"
	ip := "localhost"
	port := "3306"
	dbName := "hello-golang"
	// 创建连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("数据库链接失败:", err)
	}
}

func CreateTable() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
	if err != nil {
		fmt.Println("创建表结构失败", err)
	}
}
