package dto

import (
	"fmt"
	"sqlx/models"
)

func EmployeesFindByDepartment(department string) *[]models.Employees {
	row, err := DB.Query("SELECT id, name, department, salary FROM employees WHERE department = ?", department)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var employees []models.Employees
	for row.Next() {
		employee := models.Employees{}
		err := row.Scan(&employee.Id, &employee.Name, &employee.Department, &employee.Salary)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		employees = append(employees, employee)
	}
	return &employees
}

func EmployeesMaxSalary() *models.Employees {
	row := DB.QueryRow("SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	employees := models.Employees{}
	err := row.Scan(&employees.Id, &employees.Name, &employees.Department, &employees.Salary)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &employees
}
