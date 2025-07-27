package dto

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	user := "root"
	password := "123456"
	ip := "localhost"
	port := "3306"
	dbName := "hello-golang"
	// 创建连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库链接失败:", err)
	}
	// 测试
	err = DB.Ping()
	if err != nil {
		fmt.Println("数据库链接失败:", err)
	}
}
