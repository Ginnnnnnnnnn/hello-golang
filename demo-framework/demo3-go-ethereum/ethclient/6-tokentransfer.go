package ethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

// TokenTransfer token转账
func TokenTransfer() {
	toAddress := common.HexToAddress("0x14727CFE397bcF3c258f559a6D3AF0d8F54615b6")
	tokenAddress := common.HexToAddress("0x644Ea8f07c225Cb94426AB71af7538489139ce04")
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("privateKeyHex")
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 获取地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("address =", fromAddress.Hex())
	// 获取随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 设置Data
	var data []byte
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte("transfer(address,uint256)"))
	methodID := hash.Sum(nil)[:4]
	data = append(data, methodID...)
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	data = append(data, paddedAddress...)
	amount := new(big.Int)
	amount.SetString("1", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, paddedAmount...)
	// 获取均价gas
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 设置交易信息
	txData := types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    big.NewInt(0),
		Gas:      uint64(60000),
		GasPrice: gasPrice,
		Data:     data,
	}
	tx := types.NewTx(&txData)
	// 签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
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
