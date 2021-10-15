// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gold

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

// GoldMetaData contains all meta data concerning the Gold contract.
var GoldMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spender\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"wealth_by_level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"wealth\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// GoldABI is the input ABI used to generate the binding from.
// Deprecated: Use GoldMetaData.ABI instead.
var GoldABI = GoldMetaData.ABI

// Gold is an auto generated Go binding around an Ethereum contract.
type Gold struct {
	GoldCaller     // Read-only binding to the contract
	GoldTransactor // Write-only binding to the contract
	GoldFilterer   // Log filterer for contract events
}

// GoldCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoldSession struct {
	Contract     *Gold             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoldCallerSession struct {
	Contract *GoldCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GoldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoldTransactorSession struct {
	Contract     *GoldTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoldRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoldRaw struct {
	Contract *Gold // Generic contract binding to access the raw methods on
}

// GoldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoldCallerRaw struct {
	Contract *GoldCaller // Generic read-only contract binding to access the raw methods on
}

// GoldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoldTransactorRaw struct {
	Contract *GoldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGold creates a new instance of Gold, bound to a specific deployed contract.
func NewGold(address common.Address, backend bind.ContractBackend) (*Gold, error) {
	contract, err := bindGold(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gold{GoldCaller: GoldCaller{contract: contract}, GoldTransactor: GoldTransactor{contract: contract}, GoldFilterer: GoldFilterer{contract: contract}}, nil
}

// NewGoldCaller creates a new read-only instance of Gold, bound to a specific deployed contract.
func NewGoldCaller(address common.Address, caller bind.ContractCaller) (*GoldCaller, error) {
	contract, err := bindGold(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoldCaller{contract: contract}, nil
}

// NewGoldTransactor creates a new write-only instance of Gold, bound to a specific deployed contract.
func NewGoldTransactor(address common.Address, transactor bind.ContractTransactor) (*GoldTransactor, error) {
	contract, err := bindGold(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoldTransactor{contract: contract}, nil
}

// NewGoldFilterer creates a new log filterer instance of Gold, bound to a specific deployed contract.
func NewGoldFilterer(address common.Address, filterer bind.ContractFilterer) (*GoldFilterer, error) {
	contract, err := bindGold(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoldFilterer{contract: contract}, nil
}

// bindGold binds a generic wrapper to an already deployed contract.
func bindGold(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gold *GoldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gold.Contract.GoldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gold *GoldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gold.Contract.GoldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gold *GoldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gold.Contract.GoldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gold *GoldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gold.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gold *GoldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gold.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gold *GoldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gold.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xcca16fa8.
//
// Solidity: function allowance(uint256 , uint256 ) view returns(uint256)
func (_Gold *GoldCaller) Allowance(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xcca16fa8.
//
// Solidity: function allowance(uint256 , uint256 ) view returns(uint256)
func (_Gold *GoldSession) Allowance(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Gold.Contract.Allowance(&_Gold.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xcca16fa8.
//
// Solidity: function allowance(uint256 , uint256 ) view returns(uint256)
func (_Gold *GoldCallerSession) Allowance(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Gold.Contract.Allowance(&_Gold.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_Gold *GoldCaller) BalanceOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_Gold *GoldSession) BalanceOf(arg0 *big.Int) (*big.Int, error) {
	return _Gold.Contract.BalanceOf(&_Gold.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_Gold *GoldCallerSession) BalanceOf(arg0 *big.Int) (*big.Int, error) {
	return _Gold.Contract.BalanceOf(&_Gold.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 summoner) view returns(uint256 amount)
func (_Gold *GoldCaller) Claimable(opts *bind.CallOpts, summoner *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "claimable", summoner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 summoner) view returns(uint256 amount)
func (_Gold *GoldSession) Claimable(summoner *big.Int) (*big.Int, error) {
	return _Gold.Contract.Claimable(&_Gold.CallOpts, summoner)
}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 summoner) view returns(uint256 amount)
func (_Gold *GoldCallerSession) Claimable(summoner *big.Int) (*big.Int, error) {
	return _Gold.Contract.Claimable(&_Gold.CallOpts, summoner)
}

// Claimed is a free data retrieval call binding the contract method 0xdbe7e3bd.
//
// Solidity: function claimed(uint256 ) view returns(uint256)
func (_Gold *GoldCaller) Claimed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "claimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xdbe7e3bd.
//
// Solidity: function claimed(uint256 ) view returns(uint256)
func (_Gold *GoldSession) Claimed(arg0 *big.Int) (*big.Int, error) {
	return _Gold.Contract.Claimed(&_Gold.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0xdbe7e3bd.
//
// Solidity: function claimed(uint256 ) view returns(uint256)
func (_Gold *GoldCallerSession) Claimed(arg0 *big.Int) (*big.Int, error) {
	return _Gold.Contract.Claimed(&_Gold.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gold *GoldCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gold *GoldSession) Decimals() (uint8, error) {
	return _Gold.Contract.Decimals(&_Gold.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gold *GoldCallerSession) Decimals() (uint8, error) {
	return _Gold.Contract.Decimals(&_Gold.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gold *GoldCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gold *GoldSession) Name() (string, error) {
	return _Gold.Contract.Name(&_Gold.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gold *GoldCallerSession) Name() (string, error) {
	return _Gold.Contract.Name(&_Gold.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gold *GoldCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gold *GoldSession) Symbol() (string, error) {
	return _Gold.Contract.Symbol(&_Gold.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gold *GoldCallerSession) Symbol() (string, error) {
	return _Gold.Contract.Symbol(&_Gold.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gold *GoldCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gold *GoldSession) TotalSupply() (*big.Int, error) {
	return _Gold.Contract.TotalSupply(&_Gold.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gold *GoldCallerSession) TotalSupply() (*big.Int, error) {
	return _Gold.Contract.TotalSupply(&_Gold.CallOpts)
}

// WealthByLevel is a free data retrieval call binding the contract method 0xb373f4a3.
//
// Solidity: function wealth_by_level(uint256 level) pure returns(uint256 wealth)
func (_Gold *GoldCaller) WealthByLevel(opts *bind.CallOpts, level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gold.contract.Call(opts, &out, "wealth_by_level", level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WealthByLevel is a free data retrieval call binding the contract method 0xb373f4a3.
//
// Solidity: function wealth_by_level(uint256 level) pure returns(uint256 wealth)
func (_Gold *GoldSession) WealthByLevel(level *big.Int) (*big.Int, error) {
	return _Gold.Contract.WealthByLevel(&_Gold.CallOpts, level)
}

// WealthByLevel is a free data retrieval call binding the contract method 0xb373f4a3.
//
// Solidity: function wealth_by_level(uint256 level) pure returns(uint256 wealth)
func (_Gold *GoldCallerSession) WealthByLevel(level *big.Int) (*big.Int, error) {
	return _Gold.Contract.WealthByLevel(&_Gold.CallOpts, level)
}

// Approve is a paid mutator transaction binding the contract method 0xb866c8a4.
//
// Solidity: function approve(uint256 from, uint256 spender, uint256 amount) returns(bool)
func (_Gold *GoldTransactor) Approve(opts *bind.TransactOpts, from *big.Int, spender *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.contract.Transact(opts, "approve", from, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb866c8a4.
//
// Solidity: function approve(uint256 from, uint256 spender, uint256 amount) returns(bool)
func (_Gold *GoldSession) Approve(from *big.Int, spender *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.Approve(&_Gold.TransactOpts, from, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb866c8a4.
//
// Solidity: function approve(uint256 from, uint256 spender, uint256 amount) returns(bool)
func (_Gold *GoldTransactorSession) Approve(from *big.Int, spender *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.Approve(&_Gold.TransactOpts, from, spender, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 summoner) returns()
func (_Gold *GoldTransactor) Claim(opts *bind.TransactOpts, summoner *big.Int) (*types.Transaction, error) {
	return _Gold.contract.Transact(opts, "claim", summoner)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 summoner) returns()
func (_Gold *GoldSession) Claim(summoner *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.Claim(&_Gold.TransactOpts, summoner)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 summoner) returns()
func (_Gold *GoldTransactorSession) Claim(summoner *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.Claim(&_Gold.TransactOpts, summoner)
}

// Transfer is a paid mutator transaction binding the contract method 0x90dd2627.
//
// Solidity: function transfer(uint256 from, uint256 to, uint256 amount) returns(bool)
func (_Gold *GoldTransactor) Transfer(opts *bind.TransactOpts, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.contract.Transact(opts, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x90dd2627.
//
// Solidity: function transfer(uint256 from, uint256 to, uint256 amount) returns(bool)
func (_Gold *GoldSession) Transfer(from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.Transfer(&_Gold.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x90dd2627.
//
// Solidity: function transfer(uint256 from, uint256 to, uint256 amount) returns(bool)
func (_Gold *GoldTransactorSession) Transfer(from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.Transfer(&_Gold.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x8856f779.
//
// Solidity: function transferFrom(uint256 executor, uint256 from, uint256 to, uint256 amount) returns(bool)
func (_Gold *GoldTransactor) TransferFrom(opts *bind.TransactOpts, executor *big.Int, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.contract.Transact(opts, "transferFrom", executor, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x8856f779.
//
// Solidity: function transferFrom(uint256 executor, uint256 from, uint256 to, uint256 amount) returns(bool)
func (_Gold *GoldSession) TransferFrom(executor *big.Int, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.TransferFrom(&_Gold.TransactOpts, executor, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x8856f779.
//
// Solidity: function transferFrom(uint256 executor, uint256 from, uint256 to, uint256 amount) returns(bool)
func (_Gold *GoldTransactorSession) TransferFrom(executor *big.Int, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Gold.Contract.TransferFrom(&_Gold.TransactOpts, executor, from, to, amount)
}

// GoldApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Gold contract.
type GoldApprovalIterator struct {
	Event *GoldApproval // Event containing the contract specifics and raw log

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
func (it *GoldApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldApproval)
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
		it.Event = new(GoldApproval)
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
func (it *GoldApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldApproval represents a Approval event raised by the Gold contract.
type GoldApproval struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x08aaf4f7dd1adfa5bfe7067dea5b4ebd7e119d43257438a9189f37d7044eb09a.
//
// Solidity: event Approval(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_Gold *GoldFilterer) FilterApproval(opts *bind.FilterOpts, from []*big.Int, to []*big.Int) (*GoldApprovalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gold.contract.FilterLogs(opts, "Approval", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GoldApprovalIterator{contract: _Gold.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x08aaf4f7dd1adfa5bfe7067dea5b4ebd7e119d43257438a9189f37d7044eb09a.
//
// Solidity: event Approval(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_Gold *GoldFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GoldApproval, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gold.contract.WatchLogs(opts, "Approval", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldApproval)
				if err := _Gold.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x08aaf4f7dd1adfa5bfe7067dea5b4ebd7e119d43257438a9189f37d7044eb09a.
//
// Solidity: event Approval(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_Gold *GoldFilterer) ParseApproval(log types.Log) (*GoldApproval, error) {
	event := new(GoldApproval)
	if err := _Gold.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GoldTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Gold contract.
type GoldTransferIterator struct {
	Event *GoldTransfer // Event containing the contract specifics and raw log

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
func (it *GoldTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoldTransfer)
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
		it.Event = new(GoldTransfer)
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
func (it *GoldTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GoldTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoldTransfer represents a Transfer event raised by the Gold contract.
type GoldTransfer struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_Gold *GoldFilterer) FilterTransfer(opts *bind.FilterOpts, from []*big.Int, to []*big.Int) (*GoldTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gold.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GoldTransferIterator{contract: _Gold.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_Gold *GoldFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GoldTransfer, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gold.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GoldTransfer)
				if err := _Gold.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_Gold *GoldFilterer) ParseTransfer(log types.Log) (*GoldTransfer, error) {
	event := new(GoldTransfer)
	if err := _Gold.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
