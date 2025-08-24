// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package countnumber

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CountnumberMetaData contains all meta data concerning the Countnumber contract.
var CountnumberMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061015b8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80634f2be91f146100385780636d4ce63c14610042575b5f5ffd5b610040610060565b005b61004a610078565b6040516100579190610098565b60405180910390f35b5f5f815480929190610071906100de565b9190505550565b5f5f54905090565b5f819050919050565b61009281610080565b82525050565b5f6020820190506100ab5f830184610089565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6100e882610080565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361011a576101196100b1565b5b60018201905091905056fea26469706673582212202d2cc50daf4cb23965255fb10ebff15973efcdd40502a5896e39cf65d404069764736f6c634300081e0033",
}

// CountnumberABI is the input ABI used to generate the binding from.
// Deprecated: Use CountnumberMetaData.ABI instead.
var CountnumberABI = CountnumberMetaData.ABI

// CountnumberBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountnumberMetaData.Bin instead.
var CountnumberBin = CountnumberMetaData.Bin

// DeployCountnumber deploys a new Ethereum contract, binding an instance of Countnumber to it.
func DeployCountnumber(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Countnumber, error) {
	parsed, err := CountnumberMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountnumberBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Countnumber{CountnumberCaller: CountnumberCaller{contract: contract}, CountnumberTransactor: CountnumberTransactor{contract: contract}, CountnumberFilterer: CountnumberFilterer{contract: contract}}, nil
}

// Countnumber is an auto generated Go binding around an Ethereum contract.
type Countnumber struct {
	CountnumberCaller     // Read-only binding to the contract
	CountnumberTransactor // Write-only binding to the contract
	CountnumberFilterer   // Log filterer for contract events
}

// CountnumberCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountnumberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountnumberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountnumberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountnumberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountnumberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountnumberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountnumberSession struct {
	Contract     *Countnumber      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountnumberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountnumberCallerSession struct {
	Contract *CountnumberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CountnumberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountnumberTransactorSession struct {
	Contract     *CountnumberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CountnumberRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountnumberRaw struct {
	Contract *Countnumber // Generic contract binding to access the raw methods on
}

// CountnumberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountnumberCallerRaw struct {
	Contract *CountnumberCaller // Generic read-only contract binding to access the raw methods on
}

// CountnumberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountnumberTransactorRaw struct {
	Contract *CountnumberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCountnumber creates a new instance of Countnumber, bound to a specific deployed contract.
func NewCountnumber(address common.Address, backend bind.ContractBackend) (*Countnumber, error) {
	contract, err := bindCountnumber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Countnumber{CountnumberCaller: CountnumberCaller{contract: contract}, CountnumberTransactor: CountnumberTransactor{contract: contract}, CountnumberFilterer: CountnumberFilterer{contract: contract}}, nil
}

// NewCountnumberCaller creates a new read-only instance of Countnumber, bound to a specific deployed contract.
func NewCountnumberCaller(address common.Address, caller bind.ContractCaller) (*CountnumberCaller, error) {
	contract, err := bindCountnumber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountnumberCaller{contract: contract}, nil
}

// NewCountnumberTransactor creates a new write-only instance of Countnumber, bound to a specific deployed contract.
func NewCountnumberTransactor(address common.Address, transactor bind.ContractTransactor) (*CountnumberTransactor, error) {
	contract, err := bindCountnumber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountnumberTransactor{contract: contract}, nil
}

// NewCountnumberFilterer creates a new log filterer instance of Countnumber, bound to a specific deployed contract.
func NewCountnumberFilterer(address common.Address, filterer bind.ContractFilterer) (*CountnumberFilterer, error) {
	contract, err := bindCountnumber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountnumberFilterer{contract: contract}, nil
}

// bindCountnumber binds a generic wrapper to an already deployed contract.
func bindCountnumber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CountnumberMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Countnumber *CountnumberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Countnumber.Contract.CountnumberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Countnumber *CountnumberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Countnumber.Contract.CountnumberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Countnumber *CountnumberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Countnumber.Contract.CountnumberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Countnumber *CountnumberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Countnumber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Countnumber *CountnumberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Countnumber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Countnumber *CountnumberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Countnumber.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Countnumber *CountnumberCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Countnumber.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Countnumber *CountnumberSession) Get() (*big.Int, error) {
	return _Countnumber.Contract.Get(&_Countnumber.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Countnumber *CountnumberCallerSession) Get() (*big.Int, error) {
	return _Countnumber.Contract.Get(&_Countnumber.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Countnumber *CountnumberTransactor) Add(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Countnumber.contract.Transact(opts, "add")
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Countnumber *CountnumberSession) Add() (*types.Transaction, error) {
	return _Countnumber.Contract.Add(&_Countnumber.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Countnumber *CountnumberTransactorSession) Add() (*types.Transaction, error) {
	return _Countnumber.Contract.Add(&_Countnumber.TransactOpts)
}
