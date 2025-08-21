package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
)

var key = "0x0000000000000000000000000000000000000000000000000000000000001111"
var value = "0x0000000000000000000000000000000000000000000000000000000000002222"

// FindContractEvent 合约事件(查询)
func FindContractEvent() {
	contractAddress := common.HexToAddress(contractAddr)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(9033000),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(storeContractAbi))
	if err != nil {
		log.Fatal(err)
	}
	for _, vLog := range logs {
		fmt.Println("区块哈希 =", vLog.BlockHash.Hex())
		fmt.Println("区块高度 =", vLog.BlockNumber)
		fmt.Println("交易哈希 =", vLog.TxHash.Hex())
		// 解析Data
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("key =", common.Bytes2Hex(event.Key[:]))
		fmt.Println("value =", common.Bytes2Hex(event.Value[:]))
		// 解析Topic
		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex())
		}
		fmt.Println("topics[0]=", topics[0])
		if len(topics) > 1 {
			fmt.Println("indexed topics:", topics[1:])
		}
	}
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("signature topics=", hash.Hex())
}

// SubscribeContractEvent 合约事件(订阅)
func SubscribeContractEvent() {
	contractAddress := common.HexToAddress(contractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := wsClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(storeContractAbi))
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("区块哈希 =", vLog.BlockHash.Hex())
			fmt.Println("区块高度 =", vLog.BlockNumber)
			fmt.Println("交易哈希 =", vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
}
