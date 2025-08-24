package ethclient

import (
	"context"
	gocrypto "crypto"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

const (
	privateKeyHex = "privateKeyHex"
	rawUrl        = "rawUrl"
)

var client *ethclient.Client
var privateKey *ecdsa.PrivateKey
var publicKey gocrypto.PublicKey
var fromAddress common.Address
var chainId *big.Int

func init() {
	var err error
	client, err = ethclient.Dial(rawUrl)
	if err != nil {
		log.Fatal("初始化 eth client 失败", err)
	}
	_ = client
	// 加载账户信息
	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	publicKey = privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	// NetworkID
	chainId, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func waitSuccess(tx *types.Transaction) *types.Receipt {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil && !errors.Is(err, ethereum.NotFound) {
			log.Fatal("查询交易异常", err)
		}
		if receipt != nil && receipt.Status == 1 {
			return receipt
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
