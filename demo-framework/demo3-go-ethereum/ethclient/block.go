package ethclient

import (
	"context"
	"fmt"
	"log"
	"math/big"
)

// Block 查询区块
func Block() {
	fmt.Println("===========================查询区块===========================")
	blockNumber := big.NewInt(5671744)
	// 查询区块头,若传入 nil，它将返回最新的区块头。
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块编号 =", header.Number.Uint64())
	fmt.Println("区块时间 =", header.Time)
	fmt.Println("区块难度 =", header.Difficulty.Uint64())
	fmt.Println("区块哈希 =", header.Hash().Hex())
	// 查询区块交易数量
	count, err := client.TransactionCount(context.Background(), header.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块交易数量 = ", count)
	// 查询完整区块
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块编号 =", block.Number().Uint64())
	fmt.Println("区块时间 =", block.Time())
	fmt.Println("区块难度 =", block.Difficulty().Uint64())
	fmt.Println("区块哈希 =", block.Hash().Hex())
	fmt.Println("区块交易数量 =", len(block.Transactions()))
}
