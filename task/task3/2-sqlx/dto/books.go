package dto

import (
	"fmt"
	"github.com/shopspring/decimal"
	"sqlx/models"
)

func BooksFindGtPrice(price decimal.Decimal) *[]models.Books {
	var books []models.Books
	err := DB.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", price)
	if err != nil {
		fmt.Println(err)
	}
	return &books
}
