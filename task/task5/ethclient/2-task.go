package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"task5/contracts"
)

// ContractGeneration 任务 2：合约代码生成 任务目标
// 使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
// 具体任务
// 1.编写智能合约
// ·使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
// ·编译智能合约，生成 ABI 和字节码文件。
// 2.使用 abigen 生成 Go 绑定代码
// ·安装 abigen 工具。
// ·使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
// 3.使用生成的 Go 绑定代码与合约交互
// ·编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
// ·调用合约的方法，例如增加计数器的值。
// ·输出调用结果。
func ContractGeneration() {
	// 部署合约
	address, tx := deploy()
	waitSuccess(tx)
	// 加载合约
	contract := load(address)
	// 调用方法-累加数字
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	tx, err = contract.Add(transactOpts)
	if err != nil {
		log.Fatal("Add()调用异常", err)
	}
	waitSuccess(tx)
	fmt.Println("Add() =", tx.Hash())
	// 调用方法-查询数字
	callOpt := &bind.CallOpts{Context: context.Background()}
	number, err := contract.Get(callOpt)
	if err != nil {
		log.Fatal("Get()调用异常", err)
	}
	fmt.Println("number =", number)
}

func deploy() (common.Address, *types.Transaction) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	address, tx, instance, err := countnumber.DeployCountnumber(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("合约地址 =", address.Hex())
	fmt.Println("交易哈希 =", tx.Hash().Hex())

	_ = instance

	return address, tx
}

// LoadContract 加载合约
func load(contractAddr common.Address) *countnumber.Countnumber {
	storeContract, err := countnumber.NewCountnumber(contractAddr, client)
	if err != nil {
		log.Fatal("加载合约异常", err)
	}
	return storeContract
}
