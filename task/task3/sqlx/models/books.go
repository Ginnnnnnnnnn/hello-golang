package models

import "github.com/shopspring/decimal"

type Books struct {
	Id     int64           `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  decimal.Decimal `db:"price"`
}
