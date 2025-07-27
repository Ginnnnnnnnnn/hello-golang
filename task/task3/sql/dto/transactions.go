package dto

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
)

func TransactionsAdd(tx *sql.Tx, aid, bid int64, amount decimal.Decimal) int64 {
	result, err := tx.Exec("INSERT INTO transactions( from_account_id, to_account_id, amount ) VALUE ( ?, ?, ? )", aid, bid, amount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	count, _ := result.RowsAffected()
	return count
}
