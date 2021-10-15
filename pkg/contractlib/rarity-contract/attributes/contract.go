// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package attributes

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

// AttributesMetaData contains all meta data concerning the Attributes contract.
var AttributesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"strength\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dexterity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"constitution\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"intelligence\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"wisdom\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"charisma\",\"type\":\"uint32\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"leveler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"strength\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dexterity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"constitution\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"intelligence\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"wisdom\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"charisma\",\"type\":\"uint32\"}],\"name\":\"Leveled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current_level\",\"type\":\"uint256\"}],\"name\":\"abilities_by_level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ability_scores\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"strength\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dexterity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"constitution\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"intelligence\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"wisdom\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"charisma\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"calc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_str\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_const\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_int\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wis\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cha\",\"type\":\"uint256\"}],\"name\":\"calculate_point_buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"character_created\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_charisma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_constitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_dexterity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_intelligence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_strength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_wisdom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"level_points_spent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_str\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_dex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_const\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_int\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_wis\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_cha\",\"type\":\"uint32\"}],\"name\":\"point_buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AttributesABI is the input ABI used to generate the binding from.
// Deprecated: Use AttributesMetaData.ABI instead.
var AttributesABI = AttributesMetaData.ABI

// Attributes is an auto generated Go binding around an Ethereum contract.
type Attributes struct {
	AttributesCaller     // Read-only binding to the contract
	AttributesTransactor // Write-only binding to the contract
	AttributesFilterer   // Log filterer for contract events
}

// AttributesCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttributesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttributesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttributesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttributesSession struct {
	Contract     *Attributes       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttributesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttributesCallerSession struct {
	Contract *AttributesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AttributesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttributesTransactorSession struct {
	Contract     *AttributesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AttributesRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttributesRaw struct {
	Contract *Attributes // Generic contract binding to access the raw methods on
}

// AttributesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttributesCallerRaw struct {
	Contract *AttributesCaller // Generic read-only contract binding to access the raw methods on
}

// AttributesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttributesTransactorRaw struct {
	Contract *AttributesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttributes creates a new instance of Attributes, bound to a specific deployed contract.
func NewAttributes(address common.Address, backend bind.ContractBackend) (*Attributes, error) {
	contract, err := bindAttributes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attributes{AttributesCaller: AttributesCaller{contract: contract}, AttributesTransactor: AttributesTransactor{contract: contract}, AttributesFilterer: AttributesFilterer{contract: contract}}, nil
}

// NewAttributesCaller creates a new read-only instance of Attributes, bound to a specific deployed contract.
func NewAttributesCaller(address common.Address, caller bind.ContractCaller) (*AttributesCaller, error) {
	contract, err := bindAttributes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttributesCaller{contract: contract}, nil
}

// NewAttributesTransactor creates a new write-only instance of Attributes, bound to a specific deployed contract.
func NewAttributesTransactor(address common.Address, transactor bind.ContractTransactor) (*AttributesTransactor, error) {
	contract, err := bindAttributes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttributesTransactor{contract: contract}, nil
}

// NewAttributesFilterer creates a new log filterer instance of Attributes, bound to a specific deployed contract.
func NewAttributesFilterer(address common.Address, filterer bind.ContractFilterer) (*AttributesFilterer, error) {
	contract, err := bindAttributes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttributesFilterer{contract: contract}, nil
}

// bindAttributes binds a generic wrapper to an already deployed contract.
func bindAttributes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttributesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attributes *AttributesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attributes.Contract.AttributesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attributes *AttributesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attributes.Contract.AttributesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attributes *AttributesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attributes.Contract.AttributesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attributes *AttributesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attributes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attributes *AttributesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attributes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attributes *AttributesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attributes.Contract.contract.Transact(opts, method, params...)
}

// AbilitiesByLevel is a free data retrieval call binding the contract method 0x51c23226.
//
// Solidity: function abilities_by_level(uint256 current_level) pure returns(uint256)
func (_Attributes *AttributesCaller) AbilitiesByLevel(opts *bind.CallOpts, current_level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Attributes.contract.Call(opts, &out, "abilities_by_level", current_level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AbilitiesByLevel is a free data retrieval call binding the contract method 0x51c23226.
//
// Solidity: function abilities_by_level(uint256 current_level) pure returns(uint256)
func (_Attributes *AttributesSession) AbilitiesByLevel(current_level *big.Int) (*big.Int, error) {
	return _Attributes.Contract.AbilitiesByLevel(&_Attributes.CallOpts, current_level)
}

// AbilitiesByLevel is a free data retrieval call binding the contract method 0x51c23226.
//
// Solidity: function abilities_by_level(uint256 current_level) pure returns(uint256)
func (_Attributes *AttributesCallerSession) AbilitiesByLevel(current_level *big.Int) (*big.Int, error) {
	return _Attributes.Contract.AbilitiesByLevel(&_Attributes.CallOpts, current_level)
}

// AbilityScores is a free data retrieval call binding the contract method 0x77d9e11c.
//
// Solidity: function ability_scores(uint256 ) view returns(uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesCaller) AbilityScores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}, error) {
	var out []interface{}
	err := _Attributes.contract.Call(opts, &out, "ability_scores", arg0)

	outstruct := new(struct {
		Strength     uint32
		Dexterity    uint32
		Constitution uint32
		Intelligence uint32
		Wisdom       uint32
		Charisma     uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Strength = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Dexterity = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Constitution = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Intelligence = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Wisdom = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Charisma = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// AbilityScores is a free data retrieval call binding the contract method 0x77d9e11c.
//
// Solidity: function ability_scores(uint256 ) view returns(uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesSession) AbilityScores(arg0 *big.Int) (struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}, error) {
	return _Attributes.Contract.AbilityScores(&_Attributes.CallOpts, arg0)
}

// AbilityScores is a free data retrieval call binding the contract method 0x77d9e11c.
//
// Solidity: function ability_scores(uint256 ) view returns(uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesCallerSession) AbilityScores(arg0 *big.Int) (struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}, error) {
	return _Attributes.Contract.AbilityScores(&_Attributes.CallOpts, arg0)
}

// Calc is a free data retrieval call binding the contract method 0x38c9027a.
//
// Solidity: function calc(uint256 score) pure returns(uint256)
func (_Attributes *AttributesCaller) Calc(opts *bind.CallOpts, score *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Attributes.contract.Call(opts, &out, "calc", score)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Calc is a free data retrieval call binding the contract method 0x38c9027a.
//
// Solidity: function calc(uint256 score) pure returns(uint256)
func (_Attributes *AttributesSession) Calc(score *big.Int) (*big.Int, error) {
	return _Attributes.Contract.Calc(&_Attributes.CallOpts, score)
}

// Calc is a free data retrieval call binding the contract method 0x38c9027a.
//
// Solidity: function calc(uint256 score) pure returns(uint256)
func (_Attributes *AttributesCallerSession) Calc(score *big.Int) (*big.Int, error) {
	return _Attributes.Contract.Calc(&_Attributes.CallOpts, score)
}

// CalculatePointBuy is a free data retrieval call binding the contract method 0x9ed24c8e.
//
// Solidity: function calculate_point_buy(uint256 _str, uint256 _dex, uint256 _const, uint256 _int, uint256 _wis, uint256 _cha) pure returns(uint256)
func (_Attributes *AttributesCaller) CalculatePointBuy(opts *bind.CallOpts, _str *big.Int, _dex *big.Int, _const *big.Int, _int *big.Int, _wis *big.Int, _cha *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Attributes.contract.Call(opts, &out, "calculate_point_buy", _str, _dex, _const, _int, _wis, _cha)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePointBuy is a free data retrieval call binding the contract method 0x9ed24c8e.
//
// Solidity: function calculate_point_buy(uint256 _str, uint256 _dex, uint256 _const, uint256 _int, uint256 _wis, uint256 _cha) pure returns(uint256)
func (_Attributes *AttributesSession) CalculatePointBuy(_str *big.Int, _dex *big.Int, _const *big.Int, _int *big.Int, _wis *big.Int, _cha *big.Int) (*big.Int, error) {
	return _Attributes.Contract.CalculatePointBuy(&_Attributes.CallOpts, _str, _dex, _const, _int, _wis, _cha)
}

// CalculatePointBuy is a free data retrieval call binding the contract method 0x9ed24c8e.
//
// Solidity: function calculate_point_buy(uint256 _str, uint256 _dex, uint256 _const, uint256 _int, uint256 _wis, uint256 _cha) pure returns(uint256)
func (_Attributes *AttributesCallerSession) CalculatePointBuy(_str *big.Int, _dex *big.Int, _const *big.Int, _int *big.Int, _wis *big.Int, _cha *big.Int) (*big.Int, error) {
	return _Attributes.Contract.CalculatePointBuy(&_Attributes.CallOpts, _str, _dex, _const, _int, _wis, _cha)
}

// CharacterCreated is a free data retrieval call binding the contract method 0xf8955432.
//
// Solidity: function character_created(uint256 ) view returns(bool)
func (_Attributes *AttributesCaller) CharacterCreated(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Attributes.contract.Call(opts, &out, "character_created", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CharacterCreated is a free data retrieval call binding the contract method 0xf8955432.
//
// Solidity: function character_created(uint256 ) view returns(bool)
func (_Attributes *AttributesSession) CharacterCreated(arg0 *big.Int) (bool, error) {
	return _Attributes.Contract.CharacterCreated(&_Attributes.CallOpts, arg0)
}

// CharacterCreated is a free data retrieval call binding the contract method 0xf8955432.
//
// Solidity: function character_created(uint256 ) view returns(bool)
func (_Attributes *AttributesCallerSession) CharacterCreated(arg0 *big.Int) (bool, error) {
	return _Attributes.Contract.CharacterCreated(&_Attributes.CallOpts, arg0)
}

// LevelPointsSpent is a free data retrieval call binding the contract method 0x849ada74.
//
// Solidity: function level_points_spent(uint256 ) view returns(uint256)
func (_Attributes *AttributesCaller) LevelPointsSpent(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Attributes.contract.Call(opts, &out, "level_points_spent", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LevelPointsSpent is a free data retrieval call binding the contract method 0x849ada74.
//
// Solidity: function level_points_spent(uint256 ) view returns(uint256)
func (_Attributes *AttributesSession) LevelPointsSpent(arg0 *big.Int) (*big.Int, error) {
	return _Attributes.Contract.LevelPointsSpent(&_Attributes.CallOpts, arg0)
}

// LevelPointsSpent is a free data retrieval call binding the contract method 0x849ada74.
//
// Solidity: function level_points_spent(uint256 ) view returns(uint256)
func (_Attributes *AttributesCallerSession) LevelPointsSpent(arg0 *big.Int) (*big.Int, error) {
	return _Attributes.Contract.LevelPointsSpent(&_Attributes.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_Attributes *AttributesCaller) TokenURI(opts *bind.CallOpts, _summoner *big.Int) (string, error) {
	var out []interface{}
	err := _Attributes.contract.Call(opts, &out, "tokenURI", _summoner)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_Attributes *AttributesSession) TokenURI(_summoner *big.Int) (string, error) {
	return _Attributes.Contract.TokenURI(&_Attributes.CallOpts, _summoner)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_Attributes *AttributesCallerSession) TokenURI(_summoner *big.Int) (string, error) {
	return _Attributes.Contract.TokenURI(&_Attributes.CallOpts, _summoner)
}

// IncreaseCharisma is a paid mutator transaction binding the contract method 0xbf2bf895.
//
// Solidity: function increase_charisma(uint256 _summoner) returns()
func (_Attributes *AttributesTransactor) IncreaseCharisma(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.contract.Transact(opts, "increase_charisma", _summoner)
}

// IncreaseCharisma is a paid mutator transaction binding the contract method 0xbf2bf895.
//
// Solidity: function increase_charisma(uint256 _summoner) returns()
func (_Attributes *AttributesSession) IncreaseCharisma(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseCharisma(&_Attributes.TransactOpts, _summoner)
}

// IncreaseCharisma is a paid mutator transaction binding the contract method 0xbf2bf895.
//
// Solidity: function increase_charisma(uint256 _summoner) returns()
func (_Attributes *AttributesTransactorSession) IncreaseCharisma(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseCharisma(&_Attributes.TransactOpts, _summoner)
}

// IncreaseConstitution is a paid mutator transaction binding the contract method 0xe0d92c4a.
//
// Solidity: function increase_constitution(uint256 _summoner) returns()
func (_Attributes *AttributesTransactor) IncreaseConstitution(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.contract.Transact(opts, "increase_constitution", _summoner)
}

// IncreaseConstitution is a paid mutator transaction binding the contract method 0xe0d92c4a.
//
// Solidity: function increase_constitution(uint256 _summoner) returns()
func (_Attributes *AttributesSession) IncreaseConstitution(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseConstitution(&_Attributes.TransactOpts, _summoner)
}

// IncreaseConstitution is a paid mutator transaction binding the contract method 0xe0d92c4a.
//
// Solidity: function increase_constitution(uint256 _summoner) returns()
func (_Attributes *AttributesTransactorSession) IncreaseConstitution(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseConstitution(&_Attributes.TransactOpts, _summoner)
}

// IncreaseDexterity is a paid mutator transaction binding the contract method 0x05934d9c.
//
// Solidity: function increase_dexterity(uint256 _summoner) returns()
func (_Attributes *AttributesTransactor) IncreaseDexterity(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.contract.Transact(opts, "increase_dexterity", _summoner)
}

// IncreaseDexterity is a paid mutator transaction binding the contract method 0x05934d9c.
//
// Solidity: function increase_dexterity(uint256 _summoner) returns()
func (_Attributes *AttributesSession) IncreaseDexterity(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseDexterity(&_Attributes.TransactOpts, _summoner)
}

// IncreaseDexterity is a paid mutator transaction binding the contract method 0x05934d9c.
//
// Solidity: function increase_dexterity(uint256 _summoner) returns()
func (_Attributes *AttributesTransactorSession) IncreaseDexterity(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseDexterity(&_Attributes.TransactOpts, _summoner)
}

// IncreaseIntelligence is a paid mutator transaction binding the contract method 0x96cf4c4b.
//
// Solidity: function increase_intelligence(uint256 _summoner) returns()
func (_Attributes *AttributesTransactor) IncreaseIntelligence(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.contract.Transact(opts, "increase_intelligence", _summoner)
}

// IncreaseIntelligence is a paid mutator transaction binding the contract method 0x96cf4c4b.
//
// Solidity: function increase_intelligence(uint256 _summoner) returns()
func (_Attributes *AttributesSession) IncreaseIntelligence(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseIntelligence(&_Attributes.TransactOpts, _summoner)
}

// IncreaseIntelligence is a paid mutator transaction binding the contract method 0x96cf4c4b.
//
// Solidity: function increase_intelligence(uint256 _summoner) returns()
func (_Attributes *AttributesTransactorSession) IncreaseIntelligence(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseIntelligence(&_Attributes.TransactOpts, _summoner)
}

// IncreaseStrength is a paid mutator transaction binding the contract method 0xde999039.
//
// Solidity: function increase_strength(uint256 _summoner) returns()
func (_Attributes *AttributesTransactor) IncreaseStrength(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.contract.Transact(opts, "increase_strength", _summoner)
}

// IncreaseStrength is a paid mutator transaction binding the contract method 0xde999039.
//
// Solidity: function increase_strength(uint256 _summoner) returns()
func (_Attributes *AttributesSession) IncreaseStrength(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseStrength(&_Attributes.TransactOpts, _summoner)
}

// IncreaseStrength is a paid mutator transaction binding the contract method 0xde999039.
//
// Solidity: function increase_strength(uint256 _summoner) returns()
func (_Attributes *AttributesTransactorSession) IncreaseStrength(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseStrength(&_Attributes.TransactOpts, _summoner)
}

// IncreaseWisdom is a paid mutator transaction binding the contract method 0xfe6676b3.
//
// Solidity: function increase_wisdom(uint256 _summoner) returns()
func (_Attributes *AttributesTransactor) IncreaseWisdom(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.contract.Transact(opts, "increase_wisdom", _summoner)
}

// IncreaseWisdom is a paid mutator transaction binding the contract method 0xfe6676b3.
//
// Solidity: function increase_wisdom(uint256 _summoner) returns()
func (_Attributes *AttributesSession) IncreaseWisdom(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseWisdom(&_Attributes.TransactOpts, _summoner)
}

// IncreaseWisdom is a paid mutator transaction binding the contract method 0xfe6676b3.
//
// Solidity: function increase_wisdom(uint256 _summoner) returns()
func (_Attributes *AttributesTransactorSession) IncreaseWisdom(_summoner *big.Int) (*types.Transaction, error) {
	return _Attributes.Contract.IncreaseWisdom(&_Attributes.TransactOpts, _summoner)
}

// PointBuy is a paid mutator transaction binding the contract method 0xc3c2407c.
//
// Solidity: function point_buy(uint256 _summoner, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_Attributes *AttributesTransactor) PointBuy(opts *bind.TransactOpts, _summoner *big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _Attributes.contract.Transact(opts, "point_buy", _summoner, _str, _dex, _const, _int, _wis, _cha)
}

// PointBuy is a paid mutator transaction binding the contract method 0xc3c2407c.
//
// Solidity: function point_buy(uint256 _summoner, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_Attributes *AttributesSession) PointBuy(_summoner *big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _Attributes.Contract.PointBuy(&_Attributes.TransactOpts, _summoner, _str, _dex, _const, _int, _wis, _cha)
}

// PointBuy is a paid mutator transaction binding the contract method 0xc3c2407c.
//
// Solidity: function point_buy(uint256 _summoner, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_Attributes *AttributesTransactorSession) PointBuy(_summoner *big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _Attributes.Contract.PointBuy(&_Attributes.TransactOpts, _summoner, _str, _dex, _const, _int, _wis, _cha)
}

// AttributesCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Attributes contract.
type AttributesCreatedIterator struct {
	Event *AttributesCreated // Event containing the contract specifics and raw log

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
func (it *AttributesCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributesCreated)
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
		it.Event = new(AttributesCreated)
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
func (it *AttributesCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributesCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributesCreated represents a Created event raised by the Attributes contract.
type AttributesCreated struct {
	Creator      common.Address
	Summoner     *big.Int
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x471b798bfd4e4756e197917dd3ba2cbf782f09c2deef48cd5fa694cdba976328.
//
// Solidity: event Created(address indexed creator, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesFilterer) FilterCreated(opts *bind.FilterOpts, creator []common.Address) (*AttributesCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Attributes.contract.FilterLogs(opts, "Created", creatorRule)
	if err != nil {
		return nil, err
	}
	return &AttributesCreatedIterator{contract: _Attributes.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x471b798bfd4e4756e197917dd3ba2cbf782f09c2deef48cd5fa694cdba976328.
//
// Solidity: event Created(address indexed creator, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *AttributesCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Attributes.contract.WatchLogs(opts, "Created", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributesCreated)
				if err := _Attributes.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x471b798bfd4e4756e197917dd3ba2cbf782f09c2deef48cd5fa694cdba976328.
//
// Solidity: event Created(address indexed creator, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesFilterer) ParseCreated(log types.Log) (*AttributesCreated, error) {
	event := new(AttributesCreated)
	if err := _Attributes.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttributesLeveledIterator is returned from FilterLeveled and is used to iterate over the raw logs and unpacked data for Leveled events raised by the Attributes contract.
type AttributesLeveledIterator struct {
	Event *AttributesLeveled // Event containing the contract specifics and raw log

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
func (it *AttributesLeveledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributesLeveled)
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
		it.Event = new(AttributesLeveled)
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
func (it *AttributesLeveledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributesLeveledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributesLeveled represents a Leveled event raised by the Attributes contract.
type AttributesLeveled struct {
	Leveler      common.Address
	Summoner     *big.Int
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLeveled is a free log retrieval operation binding the contract event 0x207b63f0a8632134b76c8dd10b38a52179fc29b3fa01eb51e4662e4c720ec8d4.
//
// Solidity: event Leveled(address indexed leveler, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesFilterer) FilterLeveled(opts *bind.FilterOpts, leveler []common.Address) (*AttributesLeveledIterator, error) {

	var levelerRule []interface{}
	for _, levelerItem := range leveler {
		levelerRule = append(levelerRule, levelerItem)
	}

	logs, sub, err := _Attributes.contract.FilterLogs(opts, "Leveled", levelerRule)
	if err != nil {
		return nil, err
	}
	return &AttributesLeveledIterator{contract: _Attributes.contract, event: "Leveled", logs: logs, sub: sub}, nil
}

// WatchLeveled is a free log subscription operation binding the contract event 0x207b63f0a8632134b76c8dd10b38a52179fc29b3fa01eb51e4662e4c720ec8d4.
//
// Solidity: event Leveled(address indexed leveler, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesFilterer) WatchLeveled(opts *bind.WatchOpts, sink chan<- *AttributesLeveled, leveler []common.Address) (event.Subscription, error) {

	var levelerRule []interface{}
	for _, levelerItem := range leveler {
		levelerRule = append(levelerRule, levelerItem)
	}

	logs, sub, err := _Attributes.contract.WatchLogs(opts, "Leveled", levelerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributesLeveled)
				if err := _Attributes.contract.UnpackLog(event, "Leveled", log); err != nil {
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

// ParseLeveled is a log parse operation binding the contract event 0x207b63f0a8632134b76c8dd10b38a52179fc29b3fa01eb51e4662e4c720ec8d4.
//
// Solidity: event Leveled(address indexed leveler, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_Attributes *AttributesFilterer) ParseLeveled(log types.Log) (*AttributesLeveled, error) {
	event := new(AttributesLeveled)
	if err := _Attributes.contract.UnpackLog(event, "Leveled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
