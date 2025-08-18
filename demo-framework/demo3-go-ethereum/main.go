package main

import (
	"demo3-go-ethereum/ethclient"
)

func main() {
	// 查询区块
	ethclient.Block()
	// 查询交易
	ethclient.Transaction()
	// 查询收据
	ethclient.Receipt()
}
