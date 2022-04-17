// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20SC

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

// ERC20SCMetaData contains all meta data concerning the ERC20SC contract.
var ERC20SCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferToPlayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161072d38038061072d8339818101604052602081101561003357600080fd5b505160038190553360008181526001602052604081209290925581546001600160a01b0319161790556106c28061006b6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806353897ba01161006657806353897ba0146101b857806370a08231146101d557806395d89b41146101fb578063a9059cbb14610203578063dd62ed3e1461022f5761009e565b806306fdde03146100a3578063095ea7b31461012057806318160ddd1461016057806323b872dd1461017a578063313ce567146101b0575b600080fd5b6100ab61025d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e55781810151838201526020016100cd565b50505050905090810190601f1680156101125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61014c6004803603604081101561013657600080fd5b506001600160a01b038135169060200135610283565b604080519115158252519081900360200190f35b6101686102e9565b60408051918252519081900360200190f35b61014c6004803603606081101561019057600080fd5b506001600160a01b038135811691602081013590911690604001356102ef565b61016861044f565b61014c600480360360208110156101ce57600080fd5b5035610454565b610168600480360360208110156101eb57600080fd5b50356001600160a01b0316610534565b6100ab61054f565b61014c6004803603604081101561021957600080fd5b506001600160a01b03813516906020013561056e565b6101686004803603604081101561024557600080fd5b506001600160a01b038135811691602001351661063a565b6040518060400160405280600a81526020016922a92199182a37b5b2b760b11b81525081565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035490565b6001600160a01b03831660009081526001602052604081205482111561031457600080fd5b6001600160a01b038416600090815260026020908152604080832033845290915290205482111561034457600080fd5b6001600160a01b03841660009081526001602052604090205461036d908363ffffffff61066516565b6001600160a01b03851660009081526001602090815260408083209390935560028152828220338352905220546103aa908363ffffffff61066516565b6001600160a01b0380861660009081526002602090815260408083203384528252808320949094559186168152600190915220546103ee908363ffffffff61067716565b6001600160a01b0380851660008181526001602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b601281565b600080546001600160a01b031681526001602052604081205482111561047957600080fd5b600080546001600160a01b03168152600160205260409020546104a2908363ffffffff61066516565b600080546001600160a01b03168152600160205260408082209290925533815220546104d4908363ffffffff61067716565b3360008181526001602090815260408083209490945590548351868152935192936001600160a01b03909116927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001919050565b6001600160a01b031660009081526001602052604090205490565b6040518060400160405280600381526020016204532360ec1b81525081565b3360009081526001602052604081205482111561058a57600080fd5b336000908152600160205260409020546105aa908363ffffffff61066516565b33600090815260016020526040808220929092556001600160a01b038516815220546105dc908363ffffffff61067716565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561067157fe5b50900390565b60008282018381101561068657fe5b939250505056fea265627a7a723158202cef67fd40e49d4fe67eb93a0b5ab9d90a82b27a4bea65ebdaf3ae1e880a831364736f6c63430005110032",
}

// ERC20SCABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20SCMetaData.ABI instead.
var ERC20SCABI = ERC20SCMetaData.ABI

// ERC20SCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20SCMetaData.Bin instead.
var ERC20SCBin = ERC20SCMetaData.Bin

// DeployERC20SC deploys a new Ethereum contract, binding an instance of ERC20SC to it.
func DeployERC20SC(auth *bind.TransactOpts, backend bind.ContractBackend, total *big.Int) (common.Address, *types.Transaction, *ERC20SC, error) {
	parsed, err := ERC20SCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20SCBin), backend, total)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20SC{ERC20SCCaller: ERC20SCCaller{contract: contract}, ERC20SCTransactor: ERC20SCTransactor{contract: contract}, ERC20SCFilterer: ERC20SCFilterer{contract: contract}}, nil
}

// ERC20SC is an auto generated Go binding around an Ethereum contract.
type ERC20SC struct {
	ERC20SCCaller     // Read-only binding to the contract
	ERC20SCTransactor // Write-only binding to the contract
	ERC20SCFilterer   // Log filterer for contract events
}

// ERC20SCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20SCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20SCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20SCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20SCSession struct {
	Contract     *ERC20SC          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20SCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20SCCallerSession struct {
	Contract *ERC20SCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC20SCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20SCTransactorSession struct {
	Contract     *ERC20SCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC20SCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20SCRaw struct {
	Contract *ERC20SC // Generic contract binding to access the raw methods on
}

// ERC20SCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20SCCallerRaw struct {
	Contract *ERC20SCCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20SCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20SCTransactorRaw struct {
	Contract *ERC20SCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20SC creates a new instance of ERC20SC, bound to a specific deployed contract.
func NewERC20SC(address common.Address, backend bind.ContractBackend) (*ERC20SC, error) {
	contract, err := bindERC20SC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20SC{ERC20SCCaller: ERC20SCCaller{contract: contract}, ERC20SCTransactor: ERC20SCTransactor{contract: contract}, ERC20SCFilterer: ERC20SCFilterer{contract: contract}}, nil
}

// NewERC20SCCaller creates a new read-only instance of ERC20SC, bound to a specific deployed contract.
func NewERC20SCCaller(address common.Address, caller bind.ContractCaller) (*ERC20SCCaller, error) {
	contract, err := bindERC20SC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SCCaller{contract: contract}, nil
}

// NewERC20SCTransactor creates a new write-only instance of ERC20SC, bound to a specific deployed contract.
func NewERC20SCTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20SCTransactor, error) {
	contract, err := bindERC20SC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SCTransactor{contract: contract}, nil
}

// NewERC20SCFilterer creates a new log filterer instance of ERC20SC, bound to a specific deployed contract.
func NewERC20SCFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20SCFilterer, error) {
	contract, err := bindERC20SC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20SCFilterer{contract: contract}, nil
}

// bindERC20SC binds a generic wrapper to an already deployed contract.
func bindERC20SC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20SC *ERC20SCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20SC.Contract.ERC20SCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20SC *ERC20SCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SC.Contract.ERC20SCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20SC *ERC20SCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20SC.Contract.ERC20SCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20SC *ERC20SCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20SC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20SC *ERC20SCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20SC *ERC20SCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20SC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_ERC20SC *ERC20SCCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20SC.contract.Call(opts, &out, "allowance", owner, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_ERC20SC *ERC20SCSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _ERC20SC.Contract.Allowance(&_ERC20SC.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_ERC20SC *ERC20SCCallerSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _ERC20SC.Contract.Allowance(&_ERC20SC.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_ERC20SC *ERC20SCCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20SC.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_ERC20SC *ERC20SCSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20SC.Contract.BalanceOf(&_ERC20SC.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_ERC20SC *ERC20SCCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20SC.Contract.BalanceOf(&_ERC20SC.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ERC20SC *ERC20SCCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20SC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ERC20SC *ERC20SCSession) Decimals() (*big.Int, error) {
	return _ERC20SC.Contract.Decimals(&_ERC20SC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ERC20SC *ERC20SCCallerSession) Decimals() (*big.Int, error) {
	return _ERC20SC.Contract.Decimals(&_ERC20SC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20SC *ERC20SCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20SC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20SC *ERC20SCSession) Name() (string, error) {
	return _ERC20SC.Contract.Name(&_ERC20SC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20SC *ERC20SCCallerSession) Name() (string, error) {
	return _ERC20SC.Contract.Name(&_ERC20SC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20SC *ERC20SCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20SC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20SC *ERC20SCSession) Symbol() (string, error) {
	return _ERC20SC.Contract.Symbol(&_ERC20SC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20SC *ERC20SCCallerSession) Symbol() (string, error) {
	return _ERC20SC.Contract.Symbol(&_ERC20SC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20SC *ERC20SCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20SC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20SC *ERC20SCSession) TotalSupply() (*big.Int, error) {
	return _ERC20SC.Contract.TotalSupply(&_ERC20SC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20SC *ERC20SCCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20SC.Contract.TotalSupply(&_ERC20SC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.contract.Transact(opts, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.Approve(&_ERC20SC.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactorSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.Approve(&_ERC20SC.TransactOpts, delegate, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.Transfer(&_ERC20SC.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.Transfer(&_ERC20SC.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.contract.Transact(opts, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.TransferFrom(&_ERC20SC.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.TransferFrom(&_ERC20SC.TransactOpts, owner, buyer, numTokens)
}

// TransferToPlayer is a paid mutator transaction binding the contract method 0x53897ba0.
//
// Solidity: function transferToPlayer(uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactor) TransferToPlayer(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.contract.Transact(opts, "transferToPlayer", numTokens)
}

// TransferToPlayer is a paid mutator transaction binding the contract method 0x53897ba0.
//
// Solidity: function transferToPlayer(uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCSession) TransferToPlayer(numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.TransferToPlayer(&_ERC20SC.TransactOpts, numTokens)
}

// TransferToPlayer is a paid mutator transaction binding the contract method 0x53897ba0.
//
// Solidity: function transferToPlayer(uint256 numTokens) returns(bool)
func (_ERC20SC *ERC20SCTransactorSession) TransferToPlayer(numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20SC.Contract.TransferToPlayer(&_ERC20SC.TransactOpts, numTokens)
}

// ERC20SCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20SC contract.
type ERC20SCApprovalIterator struct {
	Event *ERC20SCApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20SCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SCApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20SCApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20SCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SCApproval represents a Approval event raised by the ERC20SC contract.
type ERC20SCApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_ERC20SC *ERC20SCFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*ERC20SCApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20SC.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SCApprovalIterator{contract: _ERC20SC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_ERC20SC *ERC20SCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20SCApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20SC.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SCApproval)
				if err := _ERC20SC.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_ERC20SC *ERC20SCFilterer) ParseApproval(log types.Log) (*ERC20SCApproval, error) {
	event := new(ERC20SCApproval)
	if err := _ERC20SC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20SC contract.
type ERC20SCTransferIterator struct {
	Event *ERC20SCTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20SCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SCTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20SCTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20SCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SCTransfer represents a Transfer event raised by the ERC20SC contract.
type ERC20SCTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_ERC20SC *ERC20SCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20SCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20SC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SCTransferIterator{contract: _ERC20SC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_ERC20SC *ERC20SCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20SCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20SC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SCTransfer)
				if err := _ERC20SC.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_ERC20SC *ERC20SCFilterer) ParseTransfer(log types.Log) (*ERC20SCTransfer, error) {
	event := new(ERC20SCTransfer)
	if err := _ERC20SC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
