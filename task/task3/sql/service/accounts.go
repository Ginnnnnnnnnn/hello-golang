package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"sql/dto"
)

func AToB(aid, bid int64, amount decimal.Decimal) {
	// 开启事务
	tx, _ := dto.DB.Begin()
	// 查询A账户信息
	aaccounts := dto.AccountsFindById(tx, aid)
	if aaccounts == nil {
		fmt.Println("查询账户异常:", aid)
		tx.Rollback()
		return
	}
	if aaccounts.Balance.LessThan(amount) {
		fmt.Println("余额不足:", bid)
		tx.Rollback()
		return
	}
	// 查询B账户信息
	baccounts := dto.AccountsFindById(tx, bid)
	if baccounts == nil {
		fmt.Println("查询账户异常:", bid)
		tx.Rollback()
		return
	}
	// 扣减A账户余额
	count := dto.AccountsSubBalance(tx, aid, amount)
	if count < 1 {
		fmt.Println("扣减账户余额失败:", bid)
		tx.Rollback()
		return
	}
	// 增加B账户余额
	count = dto.AccountsAddBalance(tx, bid, amount)
	if count < 1 {
		fmt.Println("增加账户余额失败:", bid)
		tx.Rollback()
		return
	}
	// 添加流水
	count = dto.TransactionsAdd(tx, aid, bid, amount)
	if count < 1 {
		fmt.Println("添加转账记录失败:", bid, bid, amount)
		tx.Rollback()
		return
	}
	// 提交事务
	tx.Commit()
}
