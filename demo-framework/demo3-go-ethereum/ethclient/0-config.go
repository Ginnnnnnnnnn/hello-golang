package ethclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var client *ethclient.Client
var wsClient *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/qresbFoK5F1IM8pqYSDUs")
	if err != nil {
		log.Fatal("初始化 eth client 失败", err)
	}
	_ = client
	wsClient, err = ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/qresbFoK5F1IM8pqYSDUs")
	if err != nil {
		log.Fatal(err)
	}
	_ = wsClient
}
