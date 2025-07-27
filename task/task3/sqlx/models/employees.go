package models

import "github.com/shopspring/decimal"

type Employees struct {
	Id         int64
	Name       string
	Department string
	Salary     decimal.Decimal
}
