package main

import (
	"demo2-gorm/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

// gorm
// 1.连接数据库
// 2.定义模型
// 3.CRUD
// 4.Hook
// 5.关联查询
// 6.逻辑删除
// 7.原生SQL
func main() {
	// 初始化 mysql 连接
	db, err := initMysql()
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	db.DB()
	// CRUD
	createUser := models.User{Name: "小明", BaseModel: models.GetBaseModel()}
	createResult := db.Create(&createUser)
	if createResult.Error != nil {
		fmt.Println("插入失败", createResult.Error)
		return
	}

	findUser := models.User{}
	findResult := db.Distinct().Where("1 = ?", 1).Order("id desc").Limit(1).Find(&findUser)
	if findResult.Error != nil {
		fmt.Println("查询失败", findResult.Error)
		return
	}
	fmt.Println(findUser)

	updateUser := models.User{Id: 1, Name: "小明-已修改", BaseModel: models.GetBaseModel()}
	db.Save(&updateUser)
	db.Model(&models.User{}).Where("id = ?", 4).Update("name", "123")

	deleteUser := models.User{Id: 1}
	db.Delete(&deleteUser)
	db.Where("id = ?", 3).Delete(&models.User{})

	// 原生SQL
	rawUser := models.User{}
	db.Raw("SELECT id, name FROM user WHERE id = ?", 1).Scan(&rawUser)
	fmt.Println(rawUser)
	var rawUsers []models.User
	db.Raw("SELECT id, name FROM user WHERE id > ?", 1).Scan(&rawUsers)
	fmt.Println(rawUsers)
	db.Exec("DROP TABLE user")
}

func initMysql() (db *gorm.DB, err error) {
	user := "abckids"
	password := "a38ynfiim7C51ips"
	ip := "rm-bp18w5i1649sl1z20do.mysql.rds.aliyuncs.com"
	port := "3306"
	dbName := "test"
	// 创建连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
