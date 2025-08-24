package ethclient

import (
	"demo3-go-ethereum/contracts/erc20"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

// TokenBalance Token余额
// 1.solcjs --abi IERC20Metadata.sol 生成ABI文件
// 2.solcjs --bin Store.sol 生成二进制文件
// solcjs 通过 npm install -g solc 安装
// 1.abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go
// 2.abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=store.go
// abigen 通过 go install github.com/ethereum/go-ethereum/cmd/abigen@latest 安装
func TokenBalance() {
	tokenAddress := common.HexToAddress("0x644Ea8f07c225Cb94426AB71af7538489139ce04")
	instance, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x777b1ccFEc764859e135E9A818eD90Bd1319878b")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("balance: %s\n", bal)
}
