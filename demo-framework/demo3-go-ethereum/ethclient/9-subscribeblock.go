package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

// SubscribeBlock 订阅区块
func SubscribeBlock() {
	headers := make(chan *types.Header)
	sub, err := wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("区块哈希 =", block.Hash().Hex())
			fmt.Println("区块高度 =", block.Number().Uint64())
			fmt.Println("区块时间 =", block.Time())
			fmt.Println("区块随机数 =", block.Nonce())
			fmt.Println("区块交易数 =", len(block.Transactions()))
		}
	}
}
