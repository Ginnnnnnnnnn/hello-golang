package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

// Transaction 查询交易
func Transaction() {
	fmt.Println("===========================查询交易===========================")
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 方式一 :
	// 使用 BlockByNumber 方法获取到完整的区块信息之后，可以调用区块实例的 Transactions 方法来读取块中的交易。
	// 该方法返回一个 Transaction 类型的列表，循环遍历集合并获取交易的信息。
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	for _, tx := range block.Transactions() {
		fmt.Println("交易哈希 =", tx.Hash().Hex())
		fmt.Println("交易数值 =", tx.Value().String())
		fmt.Println("交易Gas =", tx.Gas())
		fmt.Println("交易Gas价格 =", tx.GasPrice().Uint64())
		fmt.Println("交易随机数 =", tx.Nonce())
		fmt.Println("交易数据 =", tx.Data())
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("交易发起地址 =", sender.Hex())
		}
		fmt.Println("交易目标地址 =", tx.To().Hex())
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("交易状态 =", receipt.Status)
		break
	}
	// 方式二 :
	// 使用客户端的 TransactionInBlock 方法。此方法仅接受块哈希和块内事务的索引值。调用 TransactionCount 来了解块中有多少个事务。
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	for i := uint(0); i < count; i++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("交易哈希 =", tx.Hash().Hex())
		if i == 0 {
			break
		}
	}
	// 方式三 :
	// 使用 TransactionByHash 在给定具体事务哈希值的情况下直接查询单个事务。
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("交易哈希 =", tx.Hash().Hex())
	fmt.Println("交易isPending =", isPending)
}
