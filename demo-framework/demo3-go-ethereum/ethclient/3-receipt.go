package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

// Receipt 查询收据
func Receipt() {
	fmt.Println("===========================查询收据===========================")
	// 方式一 :
	// 调用 BlockReceipts 方法就可以得到指定区块中所有的收据列表。通过区块哈希
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	for _, receipt := range receiptByHash {
		fmt.Println("收据状态 =", receipt.Status)
		fmt.Println("收据日志 =", receipt.Logs)
		fmt.Println("收据哈希 =", receipt.TxHash.Hex())
		fmt.Println("收据交易索引 =", receipt.TransactionIndex)
		break
	}
	// 方式二 :
	// 调用 BlockReceipts 方法就可以得到指定区块中所有的收据列表。通过区块高度
	blockNumber := big.NewInt(5671744)
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	for _, receipt := range receiptsByNum {
		fmt.Println("收据状态 =", receipt.Status)
		fmt.Println("收据日志 =", receipt.Logs)
		fmt.Println("收据哈希 =", receipt.TxHash.Hex())
		fmt.Println("收据交易索引 =", receipt.TransactionIndex)
		break
	}
	// 方式三 :
	// 调用 TransactionReceipt 方法，使用交易哈希查询，调用 TransactionReceipt 方法
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("收据状态 =", receipt.Status)
	fmt.Println("收据日志 =", receipt.Logs)
	fmt.Println("收据哈希 =", receipt.TxHash.Hex())
	fmt.Println("收据交易索引 =", receipt.TransactionIndex)
}
