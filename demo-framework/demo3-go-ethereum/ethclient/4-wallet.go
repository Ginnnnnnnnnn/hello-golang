package ethclient

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

// Wallet 创建新钱包
func Wallet() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("私钥 =", hexutil.Encode(privateKeyBytes)[2:]) // 去掉 0x
	// 现有私钥
	//privateKey, err := crypto.HexToECDSA("privateKeyHex")
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("公钥 =", hexutil.Encode(publicKeyBytes)[4:]) // 去掉 0x04
	// 获取钱包地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("钱包地址 =", address)
	// 获取钱包地址,第二种方式
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("钱包地址(完整):", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println("钱包地址[12:]:", hexutil.Encode(hash.Sum(nil)[12:]))
}
