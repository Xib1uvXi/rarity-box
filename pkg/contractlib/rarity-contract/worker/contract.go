// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package worker

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

// WorkerMetaData contains all meta data concerning the Worker contract.
var WorkerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"adventureALL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"levelUpALL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32\",\"name\":\"_str\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_dex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_const\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_int\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_wis\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_cha\",\"type\":\"uint32\"}],\"name\":\"pointBuyALL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"adventureCraftingALL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"claimGoldALL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WorkerABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerMetaData.ABI instead.
var WorkerABI = WorkerMetaData.ABI

// Worker is an auto generated Go binding around an Ethereum contract.
type Worker struct {
	WorkerCaller     // Read-only binding to the contract
	WorkerTransactor // Write-only binding to the contract
	WorkerFilterer   // Log filterer for contract events
}

// WorkerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerSession struct {
	Contract     *Worker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerCallerSession struct {
	Contract *WorkerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WorkerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerTransactorSession struct {
	Contract     *WorkerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerRaw struct {
	Contract *Worker // Generic contract binding to access the raw methods on
}

// WorkerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerCallerRaw struct {
	Contract *WorkerCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerTransactorRaw struct {
	Contract *WorkerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorker creates a new instance of Worker, bound to a specific deployed contract.
func NewWorker(address common.Address, backend bind.ContractBackend) (*Worker, error) {
	contract, err := bindWorker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Worker{WorkerCaller: WorkerCaller{contract: contract}, WorkerTransactor: WorkerTransactor{contract: contract}, WorkerFilterer: WorkerFilterer{contract: contract}}, nil
}

// NewWorkerCaller creates a new read-only instance of Worker, bound to a specific deployed contract.
func NewWorkerCaller(address common.Address, caller bind.ContractCaller) (*WorkerCaller, error) {
	contract, err := bindWorker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerCaller{contract: contract}, nil
}

// NewWorkerTransactor creates a new write-only instance of Worker, bound to a specific deployed contract.
func NewWorkerTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerTransactor, error) {
	contract, err := bindWorker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerTransactor{contract: contract}, nil
}

// NewWorkerFilterer creates a new log filterer instance of Worker, bound to a specific deployed contract.
func NewWorkerFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerFilterer, error) {
	contract, err := bindWorker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerFilterer{contract: contract}, nil
}

// bindWorker binds a generic wrapper to an already deployed contract.
func bindWorker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WorkerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Worker *WorkerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Worker.Contract.WorkerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Worker *WorkerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Worker.Contract.WorkerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Worker *WorkerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Worker.Contract.WorkerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Worker *WorkerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Worker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Worker *WorkerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Worker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Worker *WorkerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Worker.Contract.contract.Transact(opts, method, params...)
}

// AdventureALL is a paid mutator transaction binding the contract method 0x3f1e3f76.
//
// Solidity: function adventureALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactor) AdventureALL(opts *bind.TransactOpts, _ids []*big.Int) (*types.Transaction, error) {
	return _Worker.contract.Transact(opts, "adventureALL", _ids)
}

// AdventureALL is a paid mutator transaction binding the contract method 0x3f1e3f76.
//
// Solidity: function adventureALL(uint256[] _ids) returns()
func (_Worker *WorkerSession) AdventureALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.AdventureALL(&_Worker.TransactOpts, _ids)
}

// AdventureALL is a paid mutator transaction binding the contract method 0x3f1e3f76.
//
// Solidity: function adventureALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactorSession) AdventureALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.AdventureALL(&_Worker.TransactOpts, _ids)
}

// AdventureCraftingALL is a paid mutator transaction binding the contract method 0x81531c14.
//
// Solidity: function adventureCraftingALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactor) AdventureCraftingALL(opts *bind.TransactOpts, _ids []*big.Int) (*types.Transaction, error) {
	return _Worker.contract.Transact(opts, "adventureCraftingALL", _ids)
}

// AdventureCraftingALL is a paid mutator transaction binding the contract method 0x81531c14.
//
// Solidity: function adventureCraftingALL(uint256[] _ids) returns()
func (_Worker *WorkerSession) AdventureCraftingALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.AdventureCraftingALL(&_Worker.TransactOpts, _ids)
}

// AdventureCraftingALL is a paid mutator transaction binding the contract method 0x81531c14.
//
// Solidity: function adventureCraftingALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactorSession) AdventureCraftingALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.AdventureCraftingALL(&_Worker.TransactOpts, _ids)
}

// ClaimGoldALL is a paid mutator transaction binding the contract method 0x6e3c9496.
//
// Solidity: function claimGoldALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactor) ClaimGoldALL(opts *bind.TransactOpts, _ids []*big.Int) (*types.Transaction, error) {
	return _Worker.contract.Transact(opts, "claimGoldALL", _ids)
}

// ClaimGoldALL is a paid mutator transaction binding the contract method 0x6e3c9496.
//
// Solidity: function claimGoldALL(uint256[] _ids) returns()
func (_Worker *WorkerSession) ClaimGoldALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.ClaimGoldALL(&_Worker.TransactOpts, _ids)
}

// ClaimGoldALL is a paid mutator transaction binding the contract method 0x6e3c9496.
//
// Solidity: function claimGoldALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactorSession) ClaimGoldALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.ClaimGoldALL(&_Worker.TransactOpts, _ids)
}

// LevelUpALL is a paid mutator transaction binding the contract method 0xc5a500a4.
//
// Solidity: function levelUpALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactor) LevelUpALL(opts *bind.TransactOpts, _ids []*big.Int) (*types.Transaction, error) {
	return _Worker.contract.Transact(opts, "levelUpALL", _ids)
}

// LevelUpALL is a paid mutator transaction binding the contract method 0xc5a500a4.
//
// Solidity: function levelUpALL(uint256[] _ids) returns()
func (_Worker *WorkerSession) LevelUpALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.LevelUpALL(&_Worker.TransactOpts, _ids)
}

// LevelUpALL is a paid mutator transaction binding the contract method 0xc5a500a4.
//
// Solidity: function levelUpALL(uint256[] _ids) returns()
func (_Worker *WorkerTransactorSession) LevelUpALL(_ids []*big.Int) (*types.Transaction, error) {
	return _Worker.Contract.LevelUpALL(&_Worker.TransactOpts, _ids)
}

// PointBuyALL is a paid mutator transaction binding the contract method 0xfa288f8f.
//
// Solidity: function pointBuyALL(uint256[] _ids, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_Worker *WorkerTransactor) PointBuyALL(opts *bind.TransactOpts, _ids []*big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _Worker.contract.Transact(opts, "pointBuyALL", _ids, _str, _dex, _const, _int, _wis, _cha)
}

// PointBuyALL is a paid mutator transaction binding the contract method 0xfa288f8f.
//
// Solidity: function pointBuyALL(uint256[] _ids, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_Worker *WorkerSession) PointBuyALL(_ids []*big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _Worker.Contract.PointBuyALL(&_Worker.TransactOpts, _ids, _str, _dex, _const, _int, _wis, _cha)
}

// PointBuyALL is a paid mutator transaction binding the contract method 0xfa288f8f.
//
// Solidity: function pointBuyALL(uint256[] _ids, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_Worker *WorkerTransactorSession) PointBuyALL(_ids []*big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _Worker.Contract.PointBuyALL(&_Worker.TransactOpts, _ids, _str, _dex, _const, _int, _wis, _cha)
}
