package models

import "github.com/shopspring/decimal"

type Accounts struct {
	Id      int64
	Balance decimal.Decimal
}
