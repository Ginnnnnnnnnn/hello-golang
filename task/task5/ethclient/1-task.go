package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

// BlockReadAndWrite 任务 1：区块链读写 任务目标
func BlockReadAndWrite() {
	Block()
	EthTransfer()
}

// Block 查询区块
// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
// 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
// 输出查询结果到控制台。
func Block() {
	blockNumber := big.NewInt(5671744)
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

// EthTransfer 发送交易
// 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
// 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
// 对交易进行签名，并将签名后的交易发送到网络。
// 输出交易的哈希值。
func EthTransfer() {
	// 获取随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 获取交易目标地址
	toAddress := common.HexToAddress("0x14727CFE397bcF3c258f559a6D3AF0d8F54615b6")
	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 设置交易信息
	var data []byte
	txData := types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    big.NewInt(1000000000000000),
		Gas:      uint64(21000),
		GasPrice: gasPrice,
		Data:     data,
	}
	tx := types.NewTx(&txData)
	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
