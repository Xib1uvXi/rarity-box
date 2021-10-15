// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package skills

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

// SkillsMetaData contains all meta data concerning the Skills contract.
var SkillsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"}],\"name\":\"base_per_class\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint8[36]\",\"name\":\"_skills\",\"type\":\"uint8[36]\"}],\"name\":\"calculate_points_for_set\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"}],\"name\":\"class_skills\",\"outputs\":[{\"internalType\":\"bool[36]\",\"name\":\"_skills\",\"type\":\"bool[36]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"}],\"name\":\"class_skills_by_name\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"get_skills\",\"outputs\":[{\"internalType\":\"uint8[36]\",\"name\":\"\",\"type\":\"uint8[36]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint8[36]\",\"name\":\"_skills\",\"type\":\"uint8[36]\"}],\"name\":\"is_valid_set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_attribute\",\"type\":\"uint256\"}],\"name\":\"modifier_for_attribute\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"_modifier\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint8[36]\",\"name\":\"_skills\",\"type\":\"uint8[36]\"}],\"name\":\"set_skills\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"skills\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_int\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"skills_per_level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// SkillsABI is the input ABI used to generate the binding from.
// Deprecated: Use SkillsMetaData.ABI instead.
var SkillsABI = SkillsMetaData.ABI

// Skills is an auto generated Go binding around an Ethereum contract.
type Skills struct {
	SkillsCaller     // Read-only binding to the contract
	SkillsTransactor // Write-only binding to the contract
	SkillsFilterer   // Log filterer for contract events
}

// SkillsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SkillsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SkillsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SkillsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SkillsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SkillsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SkillsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SkillsSession struct {
	Contract     *Skills           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SkillsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SkillsCallerSession struct {
	Contract *SkillsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SkillsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SkillsTransactorSession struct {
	Contract     *SkillsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SkillsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SkillsRaw struct {
	Contract *Skills // Generic contract binding to access the raw methods on
}

// SkillsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SkillsCallerRaw struct {
	Contract *SkillsCaller // Generic read-only contract binding to access the raw methods on
}

// SkillsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SkillsTransactorRaw struct {
	Contract *SkillsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSkills creates a new instance of Skills, bound to a specific deployed contract.
func NewSkills(address common.Address, backend bind.ContractBackend) (*Skills, error) {
	contract, err := bindSkills(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Skills{SkillsCaller: SkillsCaller{contract: contract}, SkillsTransactor: SkillsTransactor{contract: contract}, SkillsFilterer: SkillsFilterer{contract: contract}}, nil
}

// NewSkillsCaller creates a new read-only instance of Skills, bound to a specific deployed contract.
func NewSkillsCaller(address common.Address, caller bind.ContractCaller) (*SkillsCaller, error) {
	contract, err := bindSkills(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SkillsCaller{contract: contract}, nil
}

// NewSkillsTransactor creates a new write-only instance of Skills, bound to a specific deployed contract.
func NewSkillsTransactor(address common.Address, transactor bind.ContractTransactor) (*SkillsTransactor, error) {
	contract, err := bindSkills(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SkillsTransactor{contract: contract}, nil
}

// NewSkillsFilterer creates a new log filterer instance of Skills, bound to a specific deployed contract.
func NewSkillsFilterer(address common.Address, filterer bind.ContractFilterer) (*SkillsFilterer, error) {
	contract, err := bindSkills(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SkillsFilterer{contract: contract}, nil
}

// bindSkills binds a generic wrapper to an already deployed contract.
func bindSkills(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SkillsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Skills *SkillsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Skills.Contract.SkillsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Skills *SkillsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Skills.Contract.SkillsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Skills *SkillsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Skills.Contract.SkillsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Skills *SkillsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Skills.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Skills *SkillsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Skills.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Skills *SkillsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Skills.Contract.contract.Transact(opts, method, params...)
}

// BasePerClass is a free data retrieval call binding the contract method 0x5938b63a.
//
// Solidity: function base_per_class(uint256 _class) pure returns(uint256 base)
func (_Skills *SkillsCaller) BasePerClass(opts *bind.CallOpts, _class *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "base_per_class", _class)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePerClass is a free data retrieval call binding the contract method 0x5938b63a.
//
// Solidity: function base_per_class(uint256 _class) pure returns(uint256 base)
func (_Skills *SkillsSession) BasePerClass(_class *big.Int) (*big.Int, error) {
	return _Skills.Contract.BasePerClass(&_Skills.CallOpts, _class)
}

// BasePerClass is a free data retrieval call binding the contract method 0x5938b63a.
//
// Solidity: function base_per_class(uint256 _class) pure returns(uint256 base)
func (_Skills *SkillsCallerSession) BasePerClass(_class *big.Int) (*big.Int, error) {
	return _Skills.Contract.BasePerClass(&_Skills.CallOpts, _class)
}

// CalculatePointsForSet is a free data retrieval call binding the contract method 0x67193047.
//
// Solidity: function calculate_points_for_set(uint256 _class, uint8[36] _skills) pure returns(uint256 points)
func (_Skills *SkillsCaller) CalculatePointsForSet(opts *bind.CallOpts, _class *big.Int, _skills [36]uint8) (*big.Int, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "calculate_points_for_set", _class, _skills)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePointsForSet is a free data retrieval call binding the contract method 0x67193047.
//
// Solidity: function calculate_points_for_set(uint256 _class, uint8[36] _skills) pure returns(uint256 points)
func (_Skills *SkillsSession) CalculatePointsForSet(_class *big.Int, _skills [36]uint8) (*big.Int, error) {
	return _Skills.Contract.CalculatePointsForSet(&_Skills.CallOpts, _class, _skills)
}

// CalculatePointsForSet is a free data retrieval call binding the contract method 0x67193047.
//
// Solidity: function calculate_points_for_set(uint256 _class, uint8[36] _skills) pure returns(uint256 points)
func (_Skills *SkillsCallerSession) CalculatePointsForSet(_class *big.Int, _skills [36]uint8) (*big.Int, error) {
	return _Skills.Contract.CalculatePointsForSet(&_Skills.CallOpts, _class, _skills)
}

// ClassSkills is a free data retrieval call binding the contract method 0x9a672982.
//
// Solidity: function class_skills(uint256 _class) pure returns(bool[36] _skills)
func (_Skills *SkillsCaller) ClassSkills(opts *bind.CallOpts, _class *big.Int) ([36]bool, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "class_skills", _class)

	if err != nil {
		return *new([36]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([36]bool)).(*[36]bool)

	return out0, err

}

// ClassSkills is a free data retrieval call binding the contract method 0x9a672982.
//
// Solidity: function class_skills(uint256 _class) pure returns(bool[36] _skills)
func (_Skills *SkillsSession) ClassSkills(_class *big.Int) ([36]bool, error) {
	return _Skills.Contract.ClassSkills(&_Skills.CallOpts, _class)
}

// ClassSkills is a free data retrieval call binding the contract method 0x9a672982.
//
// Solidity: function class_skills(uint256 _class) pure returns(bool[36] _skills)
func (_Skills *SkillsCallerSession) ClassSkills(_class *big.Int) ([36]bool, error) {
	return _Skills.Contract.ClassSkills(&_Skills.CallOpts, _class)
}

// ClassSkillsByName is a free data retrieval call binding the contract method 0xda0ee354.
//
// Solidity: function class_skills_by_name(uint256 _class) view returns(string[])
func (_Skills *SkillsCaller) ClassSkillsByName(opts *bind.CallOpts, _class *big.Int) ([]string, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "class_skills_by_name", _class)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ClassSkillsByName is a free data retrieval call binding the contract method 0xda0ee354.
//
// Solidity: function class_skills_by_name(uint256 _class) view returns(string[])
func (_Skills *SkillsSession) ClassSkillsByName(_class *big.Int) ([]string, error) {
	return _Skills.Contract.ClassSkillsByName(&_Skills.CallOpts, _class)
}

// ClassSkillsByName is a free data retrieval call binding the contract method 0xda0ee354.
//
// Solidity: function class_skills_by_name(uint256 _class) view returns(string[])
func (_Skills *SkillsCallerSession) ClassSkillsByName(_class *big.Int) ([]string, error) {
	return _Skills.Contract.ClassSkillsByName(&_Skills.CallOpts, _class)
}

// GetSkills is a free data retrieval call binding the contract method 0xaa198143.
//
// Solidity: function get_skills(uint256 _summoner) view returns(uint8[36])
func (_Skills *SkillsCaller) GetSkills(opts *bind.CallOpts, _summoner *big.Int) ([36]uint8, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "get_skills", _summoner)

	if err != nil {
		return *new([36]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([36]uint8)).(*[36]uint8)

	return out0, err

}

// GetSkills is a free data retrieval call binding the contract method 0xaa198143.
//
// Solidity: function get_skills(uint256 _summoner) view returns(uint8[36])
func (_Skills *SkillsSession) GetSkills(_summoner *big.Int) ([36]uint8, error) {
	return _Skills.Contract.GetSkills(&_Skills.CallOpts, _summoner)
}

// GetSkills is a free data retrieval call binding the contract method 0xaa198143.
//
// Solidity: function get_skills(uint256 _summoner) view returns(uint8[36])
func (_Skills *SkillsCallerSession) GetSkills(_summoner *big.Int) ([36]uint8, error) {
	return _Skills.Contract.GetSkills(&_Skills.CallOpts, _summoner)
}

// IsValidSet is a free data retrieval call binding the contract method 0xe67d77e8.
//
// Solidity: function is_valid_set(uint256 _summoner, uint8[36] _skills) view returns(bool)
func (_Skills *SkillsCaller) IsValidSet(opts *bind.CallOpts, _summoner *big.Int, _skills [36]uint8) (bool, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "is_valid_set", _summoner, _skills)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSet is a free data retrieval call binding the contract method 0xe67d77e8.
//
// Solidity: function is_valid_set(uint256 _summoner, uint8[36] _skills) view returns(bool)
func (_Skills *SkillsSession) IsValidSet(_summoner *big.Int, _skills [36]uint8) (bool, error) {
	return _Skills.Contract.IsValidSet(&_Skills.CallOpts, _summoner, _skills)
}

// IsValidSet is a free data retrieval call binding the contract method 0xe67d77e8.
//
// Solidity: function is_valid_set(uint256 _summoner, uint8[36] _skills) view returns(bool)
func (_Skills *SkillsCallerSession) IsValidSet(_summoner *big.Int, _skills [36]uint8) (bool, error) {
	return _Skills.Contract.IsValidSet(&_Skills.CallOpts, _summoner, _skills)
}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_Skills *SkillsCaller) ModifierForAttribute(opts *bind.CallOpts, _attribute *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "modifier_for_attribute", _attribute)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_Skills *SkillsSession) ModifierForAttribute(_attribute *big.Int) (*big.Int, error) {
	return _Skills.Contract.ModifierForAttribute(&_Skills.CallOpts, _attribute)
}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_Skills *SkillsCallerSession) ModifierForAttribute(_attribute *big.Int) (*big.Int, error) {
	return _Skills.Contract.ModifierForAttribute(&_Skills.CallOpts, _attribute)
}

// Skills is a free data retrieval call binding the contract method 0xd48464c4.
//
// Solidity: function skills(uint256 , uint256 ) view returns(uint8)
func (_Skills *SkillsCaller) Skills(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "skills", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Skills is a free data retrieval call binding the contract method 0xd48464c4.
//
// Solidity: function skills(uint256 , uint256 ) view returns(uint8)
func (_Skills *SkillsSession) Skills(arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	return _Skills.Contract.Skills(&_Skills.CallOpts, arg0, arg1)
}

// Skills is a free data retrieval call binding the contract method 0xd48464c4.
//
// Solidity: function skills(uint256 , uint256 ) view returns(uint8)
func (_Skills *SkillsCallerSession) Skills(arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	return _Skills.Contract.Skills(&_Skills.CallOpts, arg0, arg1)
}

// SkillsPerLevel is a free data retrieval call binding the contract method 0xbffe71c6.
//
// Solidity: function skills_per_level(int256 _int, uint256 _class, uint256 _level) pure returns(uint256 points)
func (_Skills *SkillsCaller) SkillsPerLevel(opts *bind.CallOpts, _int *big.Int, _class *big.Int, _level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Skills.contract.Call(opts, &out, "skills_per_level", _int, _class, _level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SkillsPerLevel is a free data retrieval call binding the contract method 0xbffe71c6.
//
// Solidity: function skills_per_level(int256 _int, uint256 _class, uint256 _level) pure returns(uint256 points)
func (_Skills *SkillsSession) SkillsPerLevel(_int *big.Int, _class *big.Int, _level *big.Int) (*big.Int, error) {
	return _Skills.Contract.SkillsPerLevel(&_Skills.CallOpts, _int, _class, _level)
}

// SkillsPerLevel is a free data retrieval call binding the contract method 0xbffe71c6.
//
// Solidity: function skills_per_level(int256 _int, uint256 _class, uint256 _level) pure returns(uint256 points)
func (_Skills *SkillsCallerSession) SkillsPerLevel(_int *big.Int, _class *big.Int, _level *big.Int) (*big.Int, error) {
	return _Skills.Contract.SkillsPerLevel(&_Skills.CallOpts, _int, _class, _level)
}

// SetSkills is a paid mutator transaction binding the contract method 0x54d3652a.
//
// Solidity: function set_skills(uint256 _summoner, uint8[36] _skills) returns()
func (_Skills *SkillsTransactor) SetSkills(opts *bind.TransactOpts, _summoner *big.Int, _skills [36]uint8) (*types.Transaction, error) {
	return _Skills.contract.Transact(opts, "set_skills", _summoner, _skills)
}

// SetSkills is a paid mutator transaction binding the contract method 0x54d3652a.
//
// Solidity: function set_skills(uint256 _summoner, uint8[36] _skills) returns()
func (_Skills *SkillsSession) SetSkills(_summoner *big.Int, _skills [36]uint8) (*types.Transaction, error) {
	return _Skills.Contract.SetSkills(&_Skills.TransactOpts, _summoner, _skills)
}

// SetSkills is a paid mutator transaction binding the contract method 0x54d3652a.
//
// Solidity: function set_skills(uint256 _summoner, uint8[36] _skills) returns()
func (_Skills *SkillsTransactorSession) SetSkills(_summoner *big.Int, _skills [36]uint8) (*types.Transaction, error) {
	return _Skills.Contract.SetSkills(&_Skills.TransactOpts, _summoner, _skills)
}
