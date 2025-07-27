package dto

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"sql/models"
)

func AccountsFindById(tx *sql.Tx, aid int64) *models.Accounts {
	row := tx.QueryRow("SELECT id,balance FROM accounts WHERE id=?", aid)
	accounts := models.Accounts{}
	err := row.Scan(&accounts.Id, &accounts.Balance)
	if err != nil {
		return nil
	}
	return &accounts
}

func AccountsAddBalance(tx *sql.Tx, id int64, amount decimal.Decimal) int64 {
	result, err := tx.Exec(`
		UPDATE
			accounts
		SET
		    balance = balance + ?
		WHERE
		    id = ?
	`, amount, id)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	count, _ := result.RowsAffected()
	return count
}

func AccountsSubBalance(tx *sql.Tx, id int64, amount decimal.Decimal) int64 {
	result, err := tx.Exec(`
		UPDATE
			accounts
		SET
		    balance = balance - ?
		WHERE
		    id = ?
			AND balance - ? >= 0
	`, amount, id, amount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	count, _ := result.RowsAffected()
	return count
}
