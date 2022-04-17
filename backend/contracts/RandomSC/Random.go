// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RandomSC

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
)

// RandomSCMetaData contains all meta data concerning the RandomSC contract.
var RandomSCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"random_number\",\"type\":\"uint256\"}],\"name\":\"randMod\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051806060016040528060405180604001604052806004815260200163726f636b60e01b8152508152602001604051806040016040528060058152602001643830b832b960d91b81525081526020016040518060400160405280600881526020016773636973736f727360c01b815250815250600090600361009592919061009b565b506101f9565b8280548282559060005260206000209081019282156100e8579160200282015b828111156100e857825180516100d89184916020909101906100f8565b50916020019190600101906100bb565b506100f4929150610172565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013957805160ff1916838001178555610166565b82800160010185558215610166579182015b8281111561016657825182559160200191906001019061014b565b506100f4929150610198565b61019591905b808211156100f457600061018c82826101b2565b50600101610178565b90565b61019591905b808211156100f4576000815560010161019e565b50805460018160011615610100020316600290046000825580601f106101d857506101f6565b601f0160209004906000526020600020908101906101f69190610198565b50565b6101dc806102086000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063805553e814610030575b600080fd5b61004d6004803603602081101561004657600080fd5b50356100c2565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561008757818101518382015260200161006f565b50505050905090810190601f1680156100b45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051602080820184905242828401524460608381019190915233811b60808401528351607481850301815260949093019093528151910120600090600390068154811061010d57fe5b600091825260209182902001805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561019b5780601f106101705761010080835404028352916020019161019b565b820191906000526020600020905b81548152906001019060200180831161017e57829003601f168201915b5050505050905091905056fea265627a7a72315820ebf647144b0482acb67f4bb985de5f9be7d45056de4bf64c324a7d60b5d4e96464736f6c63430005110032",
}

// RandomSCABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomSCMetaData.ABI instead.
var RandomSCABI = RandomSCMetaData.ABI

// RandomSCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RandomSCMetaData.Bin instead.
var RandomSCBin = RandomSCMetaData.Bin

// DeployRandomSC deploys a new Ethereum contract, binding an instance of RandomSC to it.
func DeployRandomSC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RandomSC, error) {
	parsed, err := RandomSCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RandomSCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RandomSC{RandomSCCaller: RandomSCCaller{contract: contract}, RandomSCTransactor: RandomSCTransactor{contract: contract}, RandomSCFilterer: RandomSCFilterer{contract: contract}}, nil
}

// RandomSC is an auto generated Go binding around an Ethereum contract.
type RandomSC struct {
	RandomSCCaller     // Read-only binding to the contract
	RandomSCTransactor // Write-only binding to the contract
	RandomSCFilterer   // Log filterer for contract events
}

// RandomSCCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomSCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomSCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomSCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomSCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomSCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomSCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomSCSession struct {
	Contract     *RandomSC         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomSCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomSCCallerSession struct {
	Contract *RandomSCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RandomSCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomSCTransactorSession struct {
	Contract     *RandomSCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RandomSCRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomSCRaw struct {
	Contract *RandomSC // Generic contract binding to access the raw methods on
}

// RandomSCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomSCCallerRaw struct {
	Contract *RandomSCCaller // Generic read-only contract binding to access the raw methods on
}

// RandomSCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomSCTransactorRaw struct {
	Contract *RandomSCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomSC creates a new instance of RandomSC, bound to a specific deployed contract.
func NewRandomSC(address common.Address, backend bind.ContractBackend) (*RandomSC, error) {
	contract, err := bindRandomSC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomSC{RandomSCCaller: RandomSCCaller{contract: contract}, RandomSCTransactor: RandomSCTransactor{contract: contract}, RandomSCFilterer: RandomSCFilterer{contract: contract}}, nil
}

// NewRandomSCCaller creates a new read-only instance of RandomSC, bound to a specific deployed contract.
func NewRandomSCCaller(address common.Address, caller bind.ContractCaller) (*RandomSCCaller, error) {
	contract, err := bindRandomSC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomSCCaller{contract: contract}, nil
}

// NewRandomSCTransactor creates a new write-only instance of RandomSC, bound to a specific deployed contract.
func NewRandomSCTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomSCTransactor, error) {
	contract, err := bindRandomSC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomSCTransactor{contract: contract}, nil
}

// NewRandomSCFilterer creates a new log filterer instance of RandomSC, bound to a specific deployed contract.
func NewRandomSCFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomSCFilterer, error) {
	contract, err := bindRandomSC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomSCFilterer{contract: contract}, nil
}

// bindRandomSC binds a generic wrapper to an already deployed contract.
func bindRandomSC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomSCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomSC *RandomSCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomSC.Contract.RandomSCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomSC *RandomSCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomSC.Contract.RandomSCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomSC *RandomSCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomSC.Contract.RandomSCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomSC *RandomSCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomSC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomSC *RandomSCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomSC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomSC *RandomSCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomSC.Contract.contract.Transact(opts, method, params...)
}

// RandMod is a free data retrieval call binding the contract method 0x805553e8.
//
// Solidity: function randMod(uint256 random_number) view returns(string)
func (_RandomSC *RandomSCCaller) RandMod(opts *bind.CallOpts, random_number *big.Int) (string, error) {
	var out []interface{}
	err := _RandomSC.contract.Call(opts, &out, "randMod", random_number)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RandMod is a free data retrieval call binding the contract method 0x805553e8.
//
// Solidity: function randMod(uint256 random_number) view returns(string)
func (_RandomSC *RandomSCSession) RandMod(random_number *big.Int) (string, error) {
	return _RandomSC.Contract.RandMod(&_RandomSC.CallOpts, random_number)
}

// RandMod is a free data retrieval call binding the contract method 0x805553e8.
//
// Solidity: function randMod(uint256 random_number) view returns(string)
func (_RandomSC *RandomSCCallerSession) RandMod(random_number *big.Int) (string, error) {
	return _RandomSC.Contract.RandMod(&_RandomSC.CallOpts, random_number)
}
