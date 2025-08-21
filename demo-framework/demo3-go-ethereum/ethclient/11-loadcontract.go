package ethclient

import (
	"demo3-go-ethereum/contracts/store"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

// LoadContract 加载合约
func LoadContract() *store.Store {
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	return storeContract
}
