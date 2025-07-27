package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"sqlx/dto"
	"sqlx/models"
)

func main() {
	// 关闭链接
	defer dto.DB.Close()
	// 题目1：使用SQL扩展库进行查询
	sqlx1()
	// 题目2：实现类型安全映射
	sqlx2()
}

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
func sqlx1() {
	// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	employees := dto.EmployeesFindByDepartment("技术部")
	fmt.Println(employees)
	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	employee := dto.EmployeesMaxSalary()
	fmt.Println(employee)
}

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
func sqlx2() {
	// 定义一个 Book 结构体，包含与 books 表对应的字段。
	book := models.Books{
		Id:     1,
		Title:  "书籍",
		Author: "作者",
		Price:  decimal.Decimal{},
	}
	fmt.Println(book)
	// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	books := dto.BooksFindGtPrice(decimal.NewFromFloat(50))
	fmt.Println(books)
}
