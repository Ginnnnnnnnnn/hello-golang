package models

import "github.com/shopspring/decimal"

type Transactions struct {
	Id            int64
	FromAccountId int64
	ToAccountId   int64
	Amount        decimal.Decimal
}
