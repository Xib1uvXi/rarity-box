// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package craft_i

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

// CraftIMetaData contains all meta data concerning the CraftI contract.
var CraftIMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"adventure\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"adventurers_log\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spender\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dex\",\"type\":\"uint256\"}],\"name\":\"armor_class\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_str\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"attack_bonus\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"}],\"name\":\"base_attack_bonus_by_class\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"attack\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"base_attack_bonus_by_class_and_level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_str\",\"type\":\"uint256\"}],\"name\":\"damage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dungeon_armor_class\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dungeon_damage\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dungeon_health\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dungeon_to_hit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"}],\"name\":\"health_by_class\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"health\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_const\",\"type\":\"uint32\"}],\"name\":\"health_by_class_and_level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"health\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_attribute\",\"type\":\"uint256\"}],\"name\":\"modifier_for_attribute\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"_modifier\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"scout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_attack_bonus\",\"type\":\"int256\"}],\"name\":\"to_hit_ac\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CraftIABI is the input ABI used to generate the binding from.
// Deprecated: Use CraftIMetaData.ABI instead.
var CraftIABI = CraftIMetaData.ABI

// CraftI is an auto generated Go binding around an Ethereum contract.
type CraftI struct {
	CraftICaller     // Read-only binding to the contract
	CraftITransactor // Write-only binding to the contract
	CraftIFilterer   // Log filterer for contract events
}

// CraftICaller is an auto generated read-only Go binding around an Ethereum contract.
type CraftICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftITransactor is an auto generated write-only Go binding around an Ethereum contract.
type CraftITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CraftIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CraftISession struct {
	Contract     *CraftI           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CraftICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CraftICallerSession struct {
	Contract *CraftICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CraftITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CraftITransactorSession struct {
	Contract     *CraftITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CraftIRaw is an auto generated low-level Go binding around an Ethereum contract.
type CraftIRaw struct {
	Contract *CraftI // Generic contract binding to access the raw methods on
}

// CraftICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CraftICallerRaw struct {
	Contract *CraftICaller // Generic read-only contract binding to access the raw methods on
}

// CraftITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CraftITransactorRaw struct {
	Contract *CraftITransactor // Generic write-only contract binding to access the raw methods on
}

// NewCraftI creates a new instance of CraftI, bound to a specific deployed contract.
func NewCraftI(address common.Address, backend bind.ContractBackend) (*CraftI, error) {
	contract, err := bindCraftI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CraftI{CraftICaller: CraftICaller{contract: contract}, CraftITransactor: CraftITransactor{contract: contract}, CraftIFilterer: CraftIFilterer{contract: contract}}, nil
}

// NewCraftICaller creates a new read-only instance of CraftI, bound to a specific deployed contract.
func NewCraftICaller(address common.Address, caller bind.ContractCaller) (*CraftICaller, error) {
	contract, err := bindCraftI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CraftICaller{contract: contract}, nil
}

// NewCraftITransactor creates a new write-only instance of CraftI, bound to a specific deployed contract.
func NewCraftITransactor(address common.Address, transactor bind.ContractTransactor) (*CraftITransactor, error) {
	contract, err := bindCraftI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CraftITransactor{contract: contract}, nil
}

// NewCraftIFilterer creates a new log filterer instance of CraftI, bound to a specific deployed contract.
func NewCraftIFilterer(address common.Address, filterer bind.ContractFilterer) (*CraftIFilterer, error) {
	contract, err := bindCraftI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CraftIFilterer{contract: contract}, nil
}

// bindCraftI binds a generic wrapper to an already deployed contract.
func bindCraftI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CraftIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftI *CraftIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftI.Contract.CraftICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftI *CraftIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftI.Contract.CraftITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftI *CraftIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftI.Contract.CraftITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftI *CraftICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftI *CraftITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftI *CraftITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftI.Contract.contract.Transact(opts, method, params...)
}

// AdventurersLog is a free data retrieval call binding the contract method 0xeed25028.
//
// Solidity: function adventurers_log(uint256 ) view returns(uint256)
func (_CraftI *CraftICaller) AdventurersLog(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "adventurers_log", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdventurersLog is a free data retrieval call binding the contract method 0xeed25028.
//
// Solidity: function adventurers_log(uint256 ) view returns(uint256)
func (_CraftI *CraftISession) AdventurersLog(arg0 *big.Int) (*big.Int, error) {
	return _CraftI.Contract.AdventurersLog(&_CraftI.CallOpts, arg0)
}

// AdventurersLog is a free data retrieval call binding the contract method 0xeed25028.
//
// Solidity: function adventurers_log(uint256 ) view returns(uint256)
func (_CraftI *CraftICallerSession) AdventurersLog(arg0 *big.Int) (*big.Int, error) {
	return _CraftI.Contract.AdventurersLog(&_CraftI.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xcca16fa8.
//
// Solidity: function allowance(uint256 , uint256 ) view returns(uint256)
func (_CraftI *CraftICaller) Allowance(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xcca16fa8.
//
// Solidity: function allowance(uint256 , uint256 ) view returns(uint256)
func (_CraftI *CraftISession) Allowance(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CraftI.Contract.Allowance(&_CraftI.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xcca16fa8.
//
// Solidity: function allowance(uint256 , uint256 ) view returns(uint256)
func (_CraftI *CraftICallerSession) Allowance(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CraftI.Contract.Allowance(&_CraftI.CallOpts, arg0, arg1)
}

// ArmorClass is a free data retrieval call binding the contract method 0x90e44b51.
//
// Solidity: function armor_class(uint256 _dex) pure returns(int256)
func (_CraftI *CraftICaller) ArmorClass(opts *bind.CallOpts, _dex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "armor_class", _dex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArmorClass is a free data retrieval call binding the contract method 0x90e44b51.
//
// Solidity: function armor_class(uint256 _dex) pure returns(int256)
func (_CraftI *CraftISession) ArmorClass(_dex *big.Int) (*big.Int, error) {
	return _CraftI.Contract.ArmorClass(&_CraftI.CallOpts, _dex)
}

// ArmorClass is a free data retrieval call binding the contract method 0x90e44b51.
//
// Solidity: function armor_class(uint256 _dex) pure returns(int256)
func (_CraftI *CraftICallerSession) ArmorClass(_dex *big.Int) (*big.Int, error) {
	return _CraftI.Contract.ArmorClass(&_CraftI.CallOpts, _dex)
}

// AttackBonus is a free data retrieval call binding the contract method 0xa2b8a4df.
//
// Solidity: function attack_bonus(uint256 _class, uint256 _str, uint256 _level) pure returns(int256)
func (_CraftI *CraftICaller) AttackBonus(opts *bind.CallOpts, _class *big.Int, _str *big.Int, _level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "attack_bonus", _class, _str, _level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttackBonus is a free data retrieval call binding the contract method 0xa2b8a4df.
//
// Solidity: function attack_bonus(uint256 _class, uint256 _str, uint256 _level) pure returns(int256)
func (_CraftI *CraftISession) AttackBonus(_class *big.Int, _str *big.Int, _level *big.Int) (*big.Int, error) {
	return _CraftI.Contract.AttackBonus(&_CraftI.CallOpts, _class, _str, _level)
}

// AttackBonus is a free data retrieval call binding the contract method 0xa2b8a4df.
//
// Solidity: function attack_bonus(uint256 _class, uint256 _str, uint256 _level) pure returns(int256)
func (_CraftI *CraftICallerSession) AttackBonus(_class *big.Int, _str *big.Int, _level *big.Int) (*big.Int, error) {
	return _CraftI.Contract.AttackBonus(&_CraftI.CallOpts, _class, _str, _level)
}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_CraftI *CraftICaller) BalanceOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_CraftI *CraftISession) BalanceOf(arg0 *big.Int) (*big.Int, error) {
	return _CraftI.Contract.BalanceOf(&_CraftI.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x9cc7f708.
//
// Solidity: function balanceOf(uint256 ) view returns(uint256)
func (_CraftI *CraftICallerSession) BalanceOf(arg0 *big.Int) (*big.Int, error) {
	return _CraftI.Contract.BalanceOf(&_CraftI.CallOpts, arg0)
}

// BaseAttackBonusByClass is a free data retrieval call binding the contract method 0x69be4917.
//
// Solidity: function base_attack_bonus_by_class(uint256 _class) pure returns(uint256 attack)
func (_CraftI *CraftICaller) BaseAttackBonusByClass(opts *bind.CallOpts, _class *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "base_attack_bonus_by_class", _class)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseAttackBonusByClass is a free data retrieval call binding the contract method 0x69be4917.
//
// Solidity: function base_attack_bonus_by_class(uint256 _class) pure returns(uint256 attack)
func (_CraftI *CraftISession) BaseAttackBonusByClass(_class *big.Int) (*big.Int, error) {
	return _CraftI.Contract.BaseAttackBonusByClass(&_CraftI.CallOpts, _class)
}

// BaseAttackBonusByClass is a free data retrieval call binding the contract method 0x69be4917.
//
// Solidity: function base_attack_bonus_by_class(uint256 _class) pure returns(uint256 attack)
func (_CraftI *CraftICallerSession) BaseAttackBonusByClass(_class *big.Int) (*big.Int, error) {
	return _CraftI.Contract.BaseAttackBonusByClass(&_CraftI.CallOpts, _class)
}

// BaseAttackBonusByClassAndLevel is a free data retrieval call binding the contract method 0xaf28f75a.
//
// Solidity: function base_attack_bonus_by_class_and_level(uint256 _class, uint256 _level) pure returns(uint256)
func (_CraftI *CraftICaller) BaseAttackBonusByClassAndLevel(opts *bind.CallOpts, _class *big.Int, _level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "base_attack_bonus_by_class_and_level", _class, _level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseAttackBonusByClassAndLevel is a free data retrieval call binding the contract method 0xaf28f75a.
//
// Solidity: function base_attack_bonus_by_class_and_level(uint256 _class, uint256 _level) pure returns(uint256)
func (_CraftI *CraftISession) BaseAttackBonusByClassAndLevel(_class *big.Int, _level *big.Int) (*big.Int, error) {
	return _CraftI.Contract.BaseAttackBonusByClassAndLevel(&_CraftI.CallOpts, _class, _level)
}

// BaseAttackBonusByClassAndLevel is a free data retrieval call binding the contract method 0xaf28f75a.
//
// Solidity: function base_attack_bonus_by_class_and_level(uint256 _class, uint256 _level) pure returns(uint256)
func (_CraftI *CraftICallerSession) BaseAttackBonusByClassAndLevel(_class *big.Int, _level *big.Int) (*big.Int, error) {
	return _CraftI.Contract.BaseAttackBonusByClassAndLevel(&_CraftI.CallOpts, _class, _level)
}

// Damage is a free data retrieval call binding the contract method 0x573265dd.
//
// Solidity: function damage(uint256 _str) pure returns(uint256)
func (_CraftI *CraftICaller) Damage(opts *bind.CallOpts, _str *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "damage", _str)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Damage is a free data retrieval call binding the contract method 0x573265dd.
//
// Solidity: function damage(uint256 _str) pure returns(uint256)
func (_CraftI *CraftISession) Damage(_str *big.Int) (*big.Int, error) {
	return _CraftI.Contract.Damage(&_CraftI.CallOpts, _str)
}

// Damage is a free data retrieval call binding the contract method 0x573265dd.
//
// Solidity: function damage(uint256 _str) pure returns(uint256)
func (_CraftI *CraftICallerSession) Damage(_str *big.Int) (*big.Int, error) {
	return _CraftI.Contract.Damage(&_CraftI.CallOpts, _str)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CraftI *CraftICaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CraftI *CraftISession) Decimals() (uint8, error) {
	return _CraftI.Contract.Decimals(&_CraftI.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CraftI *CraftICallerSession) Decimals() (uint8, error) {
	return _CraftI.Contract.Decimals(&_CraftI.CallOpts)
}

// DungeonArmorClass is a free data retrieval call binding the contract method 0x5a76e14e.
//
// Solidity: function dungeon_armor_class() view returns(int256)
func (_CraftI *CraftICaller) DungeonArmorClass(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "dungeon_armor_class")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DungeonArmorClass is a free data retrieval call binding the contract method 0x5a76e14e.
//
// Solidity: function dungeon_armor_class() view returns(int256)
func (_CraftI *CraftISession) DungeonArmorClass() (*big.Int, error) {
	return _CraftI.Contract.DungeonArmorClass(&_CraftI.CallOpts)
}

// DungeonArmorClass is a free data retrieval call binding the contract method 0x5a76e14e.
//
// Solidity: function dungeon_armor_class() view returns(int256)
func (_CraftI *CraftICallerSession) DungeonArmorClass() (*big.Int, error) {
	return _CraftI.Contract.DungeonArmorClass(&_CraftI.CallOpts)
}

// DungeonDamage is a free data retrieval call binding the contract method 0x052922e2.
//
// Solidity: function dungeon_damage() view returns(int256)
func (_CraftI *CraftICaller) DungeonDamage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "dungeon_damage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DungeonDamage is a free data retrieval call binding the contract method 0x052922e2.
//
// Solidity: function dungeon_damage() view returns(int256)
func (_CraftI *CraftISession) DungeonDamage() (*big.Int, error) {
	return _CraftI.Contract.DungeonDamage(&_CraftI.CallOpts)
}

// DungeonDamage is a free data retrieval call binding the contract method 0x052922e2.
//
// Solidity: function dungeon_damage() view returns(int256)
func (_CraftI *CraftICallerSession) DungeonDamage() (*big.Int, error) {
	return _CraftI.Contract.DungeonDamage(&_CraftI.CallOpts)
}

// DungeonHealth is a free data retrieval call binding the contract method 0x5d0f1d77.
//
// Solidity: function dungeon_health() view returns(int256)
func (_CraftI *CraftICaller) DungeonHealth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "dungeon_health")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DungeonHealth is a free data retrieval call binding the contract method 0x5d0f1d77.
//
// Solidity: function dungeon_health() view returns(int256)
func (_CraftI *CraftISession) DungeonHealth() (*big.Int, error) {
	return _CraftI.Contract.DungeonHealth(&_CraftI.CallOpts)
}

// DungeonHealth is a free data retrieval call binding the contract method 0x5d0f1d77.
//
// Solidity: function dungeon_health() view returns(int256)
func (_CraftI *CraftICallerSession) DungeonHealth() (*big.Int, error) {
	return _CraftI.Contract.DungeonHealth(&_CraftI.CallOpts)
}

// DungeonToHit is a free data retrieval call binding the contract method 0x4a87f992.
//
// Solidity: function dungeon_to_hit() view returns(int256)
func (_CraftI *CraftICaller) DungeonToHit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "dungeon_to_hit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DungeonToHit is a free data retrieval call binding the contract method 0x4a87f992.
//
// Solidity: function dungeon_to_hit() view returns(int256)
func (_CraftI *CraftISession) DungeonToHit() (*big.Int, error) {
	return _CraftI.Contract.DungeonToHit(&_CraftI.CallOpts)
}

// DungeonToHit is a free data retrieval call binding the contract method 0x4a87f992.
//
// Solidity: function dungeon_to_hit() view returns(int256)
func (_CraftI *CraftICallerSession) DungeonToHit() (*big.Int, error) {
	return _CraftI.Contract.DungeonToHit(&_CraftI.CallOpts)
}

// HealthByClass is a free data retrieval call binding the contract method 0x1bc1c1ce.
//
// Solidity: function health_by_class(uint256 _class) pure returns(uint256 health)
func (_CraftI *CraftICaller) HealthByClass(opts *bind.CallOpts, _class *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "health_by_class", _class)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthByClass is a free data retrieval call binding the contract method 0x1bc1c1ce.
//
// Solidity: function health_by_class(uint256 _class) pure returns(uint256 health)
func (_CraftI *CraftISession) HealthByClass(_class *big.Int) (*big.Int, error) {
	return _CraftI.Contract.HealthByClass(&_CraftI.CallOpts, _class)
}

// HealthByClass is a free data retrieval call binding the contract method 0x1bc1c1ce.
//
// Solidity: function health_by_class(uint256 _class) pure returns(uint256 health)
func (_CraftI *CraftICallerSession) HealthByClass(_class *big.Int) (*big.Int, error) {
	return _CraftI.Contract.HealthByClass(&_CraftI.CallOpts, _class)
}

// HealthByClassAndLevel is a free data retrieval call binding the contract method 0x6808eae7.
//
// Solidity: function health_by_class_and_level(uint256 _class, uint256 _level, uint32 _const) pure returns(uint256 health)
func (_CraftI *CraftICaller) HealthByClassAndLevel(opts *bind.CallOpts, _class *big.Int, _level *big.Int, _const uint32) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "health_by_class_and_level", _class, _level, _const)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthByClassAndLevel is a free data retrieval call binding the contract method 0x6808eae7.
//
// Solidity: function health_by_class_and_level(uint256 _class, uint256 _level, uint32 _const) pure returns(uint256 health)
func (_CraftI *CraftISession) HealthByClassAndLevel(_class *big.Int, _level *big.Int, _const uint32) (*big.Int, error) {
	return _CraftI.Contract.HealthByClassAndLevel(&_CraftI.CallOpts, _class, _level, _const)
}

// HealthByClassAndLevel is a free data retrieval call binding the contract method 0x6808eae7.
//
// Solidity: function health_by_class_and_level(uint256 _class, uint256 _level, uint32 _const) pure returns(uint256 health)
func (_CraftI *CraftICallerSession) HealthByClassAndLevel(_class *big.Int, _level *big.Int, _const uint32) (*big.Int, error) {
	return _CraftI.Contract.HealthByClassAndLevel(&_CraftI.CallOpts, _class, _level, _const)
}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_CraftI *CraftICaller) ModifierForAttribute(opts *bind.CallOpts, _attribute *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "modifier_for_attribute", _attribute)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_CraftI *CraftISession) ModifierForAttribute(_attribute *big.Int) (*big.Int, error) {
	return _CraftI.Contract.ModifierForAttribute(&_CraftI.CallOpts, _attribute)
}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_CraftI *CraftICallerSession) ModifierForAttribute(_attribute *big.Int) (*big.Int, error) {
	return _CraftI.Contract.ModifierForAttribute(&_CraftI.CallOpts, _attribute)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CraftI *CraftICaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CraftI *CraftISession) Name() (string, error) {
	return _CraftI.Contract.Name(&_CraftI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CraftI *CraftICallerSession) Name() (string, error) {
	return _CraftI.Contract.Name(&_CraftI.CallOpts)
}

// Scout is a free data retrieval call binding the contract method 0x4bec27bd.
//
// Solidity: function scout(uint256 _summoner) view returns(uint256 reward)
func (_CraftI *CraftICaller) Scout(opts *bind.CallOpts, _summoner *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "scout", _summoner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scout is a free data retrieval call binding the contract method 0x4bec27bd.
//
// Solidity: function scout(uint256 _summoner) view returns(uint256 reward)
func (_CraftI *CraftISession) Scout(_summoner *big.Int) (*big.Int, error) {
	return _CraftI.Contract.Scout(&_CraftI.CallOpts, _summoner)
}

// Scout is a free data retrieval call binding the contract method 0x4bec27bd.
//
// Solidity: function scout(uint256 _summoner) view returns(uint256 reward)
func (_CraftI *CraftICallerSession) Scout(_summoner *big.Int) (*big.Int, error) {
	return _CraftI.Contract.Scout(&_CraftI.CallOpts, _summoner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CraftI *CraftICaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CraftI *CraftISession) Symbol() (string, error) {
	return _CraftI.Contract.Symbol(&_CraftI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CraftI *CraftICallerSession) Symbol() (string, error) {
	return _CraftI.Contract.Symbol(&_CraftI.CallOpts)
}

// ToHitAc is a free data retrieval call binding the contract method 0xfabb5737.
//
// Solidity: function to_hit_ac(int256 _attack_bonus) pure returns(bool)
func (_CraftI *CraftICaller) ToHitAc(opts *bind.CallOpts, _attack_bonus *big.Int) (bool, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "to_hit_ac", _attack_bonus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ToHitAc is a free data retrieval call binding the contract method 0xfabb5737.
//
// Solidity: function to_hit_ac(int256 _attack_bonus) pure returns(bool)
func (_CraftI *CraftISession) ToHitAc(_attack_bonus *big.Int) (bool, error) {
	return _CraftI.Contract.ToHitAc(&_CraftI.CallOpts, _attack_bonus)
}

// ToHitAc is a free data retrieval call binding the contract method 0xfabb5737.
//
// Solidity: function to_hit_ac(int256 _attack_bonus) pure returns(bool)
func (_CraftI *CraftICallerSession) ToHitAc(_attack_bonus *big.Int) (bool, error) {
	return _CraftI.Contract.ToHitAc(&_CraftI.CallOpts, _attack_bonus)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CraftI *CraftICaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftI.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CraftI *CraftISession) TotalSupply() (*big.Int, error) {
	return _CraftI.Contract.TotalSupply(&_CraftI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CraftI *CraftICallerSession) TotalSupply() (*big.Int, error) {
	return _CraftI.Contract.TotalSupply(&_CraftI.CallOpts)
}

// Adventure is a paid mutator transaction binding the contract method 0xb00b52f1.
//
// Solidity: function adventure(uint256 _summoner) returns(uint256 reward)
func (_CraftI *CraftITransactor) Adventure(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _CraftI.contract.Transact(opts, "adventure", _summoner)
}

// Adventure is a paid mutator transaction binding the contract method 0xb00b52f1.
//
// Solidity: function adventure(uint256 _summoner) returns(uint256 reward)
func (_CraftI *CraftISession) Adventure(_summoner *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.Adventure(&_CraftI.TransactOpts, _summoner)
}

// Adventure is a paid mutator transaction binding the contract method 0xb00b52f1.
//
// Solidity: function adventure(uint256 _summoner) returns(uint256 reward)
func (_CraftI *CraftITransactorSession) Adventure(_summoner *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.Adventure(&_CraftI.TransactOpts, _summoner)
}

// Approve is a paid mutator transaction binding the contract method 0xb866c8a4.
//
// Solidity: function approve(uint256 from, uint256 spender, uint256 amount) returns(bool)
func (_CraftI *CraftITransactor) Approve(opts *bind.TransactOpts, from *big.Int, spender *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.contract.Transact(opts, "approve", from, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb866c8a4.
//
// Solidity: function approve(uint256 from, uint256 spender, uint256 amount) returns(bool)
func (_CraftI *CraftISession) Approve(from *big.Int, spender *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.Approve(&_CraftI.TransactOpts, from, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb866c8a4.
//
// Solidity: function approve(uint256 from, uint256 spender, uint256 amount) returns(bool)
func (_CraftI *CraftITransactorSession) Approve(from *big.Int, spender *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.Approve(&_CraftI.TransactOpts, from, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x90dd2627.
//
// Solidity: function transfer(uint256 from, uint256 to, uint256 amount) returns(bool)
func (_CraftI *CraftITransactor) Transfer(opts *bind.TransactOpts, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.contract.Transact(opts, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x90dd2627.
//
// Solidity: function transfer(uint256 from, uint256 to, uint256 amount) returns(bool)
func (_CraftI *CraftISession) Transfer(from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.Transfer(&_CraftI.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x90dd2627.
//
// Solidity: function transfer(uint256 from, uint256 to, uint256 amount) returns(bool)
func (_CraftI *CraftITransactorSession) Transfer(from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.Transfer(&_CraftI.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x8856f779.
//
// Solidity: function transferFrom(uint256 executor, uint256 from, uint256 to, uint256 amount) returns(bool)
func (_CraftI *CraftITransactor) TransferFrom(opts *bind.TransactOpts, executor *big.Int, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.contract.Transact(opts, "transferFrom", executor, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x8856f779.
//
// Solidity: function transferFrom(uint256 executor, uint256 from, uint256 to, uint256 amount) returns(bool)
func (_CraftI *CraftISession) TransferFrom(executor *big.Int, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.TransferFrom(&_CraftI.TransactOpts, executor, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x8856f779.
//
// Solidity: function transferFrom(uint256 executor, uint256 from, uint256 to, uint256 amount) returns(bool)
func (_CraftI *CraftITransactorSession) TransferFrom(executor *big.Int, from *big.Int, to *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CraftI.Contract.TransferFrom(&_CraftI.TransactOpts, executor, from, to, amount)
}

// CraftIApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CraftI contract.
type CraftIApprovalIterator struct {
	Event *CraftIApproval // Event containing the contract specifics and raw log

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
func (it *CraftIApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftIApproval)
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
		it.Event = new(CraftIApproval)
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
func (it *CraftIApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftIApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftIApproval represents a Approval event raised by the CraftI contract.
type CraftIApproval struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x08aaf4f7dd1adfa5bfe7067dea5b4ebd7e119d43257438a9189f37d7044eb09a.
//
// Solidity: event Approval(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_CraftI *CraftIFilterer) FilterApproval(opts *bind.FilterOpts, from []*big.Int, to []*big.Int) (*CraftIApprovalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CraftI.contract.FilterLogs(opts, "Approval", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CraftIApprovalIterator{contract: _CraftI.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x08aaf4f7dd1adfa5bfe7067dea5b4ebd7e119d43257438a9189f37d7044eb09a.
//
// Solidity: event Approval(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_CraftI *CraftIFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CraftIApproval, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CraftI.contract.WatchLogs(opts, "Approval", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftIApproval)
				if err := _CraftI.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CraftI *CraftIFilterer) ParseApproval(log types.Log) (*CraftIApproval, error) {
	event := new(CraftIApproval)
	if err := _CraftI.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftITransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CraftI contract.
type CraftITransferIterator struct {
	Event *CraftITransfer // Event containing the contract specifics and raw log

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
func (it *CraftITransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftITransfer)
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
		it.Event = new(CraftITransfer)
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
func (it *CraftITransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftITransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftITransfer represents a Transfer event raised by the CraftI contract.
type CraftITransfer struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_CraftI *CraftIFilterer) FilterTransfer(opts *bind.FilterOpts, from []*big.Int, to []*big.Int) (*CraftITransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CraftI.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CraftITransferIterator{contract: _CraftI.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 indexed from, uint256 indexed to, uint256 amount)
func (_CraftI *CraftIFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CraftITransfer, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CraftI.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftITransfer)
				if err := _CraftI.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CraftI *CraftIFilterer) ParseTransfer(log types.Log) (*CraftITransfer, error) {
	event := new(CraftITransfer)
	if err := _CraftI.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
