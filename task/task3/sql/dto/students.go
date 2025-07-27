package dto

import (
	"fmt"
	"sql/models"
)

func StudentsInstall() {
	_, err := DB.Exec("INSERT INTO students( name, age, grade ) VALUE ( ?, ?, ? )", "张三", 20, "三年级")
	if err != nil {
		fmt.Println("插入失败:", err)
	}
}

func StudentsQuery() {
	rows, err := DB.Query("SELECT id, name, age, grade FROM students WHERE age > ?", 18)
	if err != nil {
		fmt.Println("查询失败:", err)
	}
	for rows.Next() {
		students := models.Students{}
		rows.Scan(&students.Id, &students.Name, &students.Age, &students.Gender)
		fmt.Println(students)
	}
}

func StudentsUpdate() {
	_, err := DB.Exec("UPDATE students SET grade = ? WHERE name = ?", "四年级", "张三")
	if err != nil {
		fmt.Println("更新失败:", err)
	}
}

func StudentsDelete() {
	_, err := DB.Exec("DELETE FROM students WHERE age < ?", 15)
	if err != nil {
		fmt.Println("删除失败:", err)
	}
}
