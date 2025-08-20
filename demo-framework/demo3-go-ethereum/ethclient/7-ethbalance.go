package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"
)

// EthBalance ETH余额
func EthBalance() {
	account := common.HexToAddress("0x777b1ccFEc764859e135E9A818eD90Bd1319878b")
	// 方法一
	// 调用 ethclient 的 BalanceAt 方法，给它传递账户地址和可选的区块号。将区块号设置为 nil 将返回最新的余额。
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("余额 =", balance)
	// 方法二
	// 传区块高度能读取指定区块时的账户余额，区块高度必须是 big.Int 类型。
	blockNumber := big.NewInt(5532993)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("余额 =", balance)
	// 待处理的余额
	// 例如，在提交或等待交易确认后。客户端提供了类似 BalanceAt 的方法，名为 PendingBalanceAt，它接收账户地址作为参数。
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("待处理余额 =", pendingBalance)
	// 转为位数
	fBalance := new(big.Float)
	fBalance.SetString(pendingBalance.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("余额(转换) =", ethValue)
}
