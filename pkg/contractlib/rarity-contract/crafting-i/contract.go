// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crafting_i

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

// CraftingIMetaData contains all meta data concerning the CraftingI contract.
var CraftingIMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"check\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"base_type\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"item_type\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"craft_i\",\"type\":\"uint256\"}],\"name\":\"Crafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SUMMMONER_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_base_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_item_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_crafting_materials\",\"type\":\"uint256\"}],\"name\":\"craft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dc\",\"type\":\"uint256\"}],\"name\":\"craft_skillcheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"crafted\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"check\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_item_id\",\"type\":\"uint256\"}],\"name\":\"get_armor_dc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dc\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_base_type\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_item_id\",\"type\":\"uint256\"}],\"name\":\"get_dc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dc\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_goods_dc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dc\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_base_type\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_item_type\",\"type\":\"uint256\"}],\"name\":\"get_item_cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_item\",\"type\":\"uint256\"}],\"name\":\"get_token_uri_armor\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_item\",\"type\":\"uint256\"}],\"name\":\"get_token_uri_goods\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_item\",\"type\":\"uint256\"}],\"name\":\"get_token_uri_weapon\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_type_id\",\"type\":\"uint256\"}],\"name\":\"get_type\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_type\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_item_id\",\"type\":\"uint256\"}],\"name\":\"get_weapon_dc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dc\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_base_type\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_item_type\",\"type\":\"uint256\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"base_type\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"item_type\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"crafted\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"crafter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_attribute\",\"type\":\"uint256\"}],\"name\":\"modifier_for_attribute\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"_modifier\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"next_item\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base_type\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_item_type\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_crafting_materials\",\"type\":\"uint256\"}],\"name\":\"simulate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"crafted\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"check\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dc\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_item\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CraftingIABI is the input ABI used to generate the binding from.
// Deprecated: Use CraftingIMetaData.ABI instead.
var CraftingIABI = CraftingIMetaData.ABI

// CraftingI is an auto generated Go binding around an Ethereum contract.
type CraftingI struct {
	CraftingICaller     // Read-only binding to the contract
	CraftingITransactor // Write-only binding to the contract
	CraftingIFilterer   // Log filterer for contract events
}

// CraftingICaller is an auto generated read-only Go binding around an Ethereum contract.
type CraftingICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingITransactor is an auto generated write-only Go binding around an Ethereum contract.
type CraftingITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CraftingIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CraftingISession struct {
	Contract     *CraftingI        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CraftingICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CraftingICallerSession struct {
	Contract *CraftingICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CraftingITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CraftingITransactorSession struct {
	Contract     *CraftingITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CraftingIRaw is an auto generated low-level Go binding around an Ethereum contract.
type CraftingIRaw struct {
	Contract *CraftingI // Generic contract binding to access the raw methods on
}

// CraftingICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CraftingICallerRaw struct {
	Contract *CraftingICaller // Generic read-only contract binding to access the raw methods on
}

// CraftingITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CraftingITransactorRaw struct {
	Contract *CraftingITransactor // Generic write-only contract binding to access the raw methods on
}

// NewCraftingI creates a new instance of CraftingI, bound to a specific deployed contract.
func NewCraftingI(address common.Address, backend bind.ContractBackend) (*CraftingI, error) {
	contract, err := bindCraftingI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CraftingI{CraftingICaller: CraftingICaller{contract: contract}, CraftingITransactor: CraftingITransactor{contract: contract}, CraftingIFilterer: CraftingIFilterer{contract: contract}}, nil
}

// NewCraftingICaller creates a new read-only instance of CraftingI, bound to a specific deployed contract.
func NewCraftingICaller(address common.Address, caller bind.ContractCaller) (*CraftingICaller, error) {
	contract, err := bindCraftingI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingICaller{contract: contract}, nil
}

// NewCraftingITransactor creates a new write-only instance of CraftingI, bound to a specific deployed contract.
func NewCraftingITransactor(address common.Address, transactor bind.ContractTransactor) (*CraftingITransactor, error) {
	contract, err := bindCraftingI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingITransactor{contract: contract}, nil
}

// NewCraftingIFilterer creates a new log filterer instance of CraftingI, bound to a specific deployed contract.
func NewCraftingIFilterer(address common.Address, filterer bind.ContractFilterer) (*CraftingIFilterer, error) {
	contract, err := bindCraftingI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CraftingIFilterer{contract: contract}, nil
}

// bindCraftingI binds a generic wrapper to an already deployed contract.
func bindCraftingI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CraftingIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftingI *CraftingIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftingI.Contract.CraftingICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftingI *CraftingIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingI.Contract.CraftingITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftingI *CraftingIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftingI.Contract.CraftingITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftingI *CraftingICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftingI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftingI *CraftingITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftingI *CraftingITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftingI.Contract.contract.Transact(opts, method, params...)
}

// SUMMMONERID is a free data retrieval call binding the contract method 0x2420b7ad.
//
// Solidity: function SUMMMONER_ID() view returns(uint256)
func (_CraftingI *CraftingICaller) SUMMMONERID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "SUMMMONER_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUMMMONERID is a free data retrieval call binding the contract method 0x2420b7ad.
//
// Solidity: function SUMMMONER_ID() view returns(uint256)
func (_CraftingI *CraftingISession) SUMMMONERID() (*big.Int, error) {
	return _CraftingI.Contract.SUMMMONERID(&_CraftingI.CallOpts)
}

// SUMMMONERID is a free data retrieval call binding the contract method 0x2420b7ad.
//
// Solidity: function SUMMMONER_ID() view returns(uint256)
func (_CraftingI *CraftingICallerSession) SUMMMONERID() (*big.Int, error) {
	return _CraftingI.Contract.SUMMMONERID(&_CraftingI.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CraftingI *CraftingICaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CraftingI *CraftingISession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CraftingI.Contract.BalanceOf(&_CraftingI.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CraftingI *CraftingICallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CraftingI.Contract.BalanceOf(&_CraftingI.CallOpts, owner)
}

// CraftSkillcheck is a free data retrieval call binding the contract method 0xdae1da31.
//
// Solidity: function craft_skillcheck(uint256 _summoner, uint256 _dc) view returns(bool crafted, int256 check)
func (_CraftingI *CraftingICaller) CraftSkillcheck(opts *bind.CallOpts, _summoner *big.Int, _dc *big.Int) (struct {
	Crafted bool
	Check   *big.Int
}, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "craft_skillcheck", _summoner, _dc)

	outstruct := new(struct {
		Crafted bool
		Check   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Crafted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Check = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CraftSkillcheck is a free data retrieval call binding the contract method 0xdae1da31.
//
// Solidity: function craft_skillcheck(uint256 _summoner, uint256 _dc) view returns(bool crafted, int256 check)
func (_CraftingI *CraftingISession) CraftSkillcheck(_summoner *big.Int, _dc *big.Int) (struct {
	Crafted bool
	Check   *big.Int
}, error) {
	return _CraftingI.Contract.CraftSkillcheck(&_CraftingI.CallOpts, _summoner, _dc)
}

// CraftSkillcheck is a free data retrieval call binding the contract method 0xdae1da31.
//
// Solidity: function craft_skillcheck(uint256 _summoner, uint256 _dc) view returns(bool crafted, int256 check)
func (_CraftingI *CraftingICallerSession) CraftSkillcheck(_summoner *big.Int, _dc *big.Int) (struct {
	Crafted bool
	Check   *big.Int
}, error) {
	return _CraftingI.Contract.CraftSkillcheck(&_CraftingI.CallOpts, _summoner, _dc)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CraftingI *CraftingICaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CraftingI *CraftingISession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CraftingI.Contract.GetApproved(&_CraftingI.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CraftingI *CraftingICallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CraftingI.Contract.GetApproved(&_CraftingI.CallOpts, tokenId)
}

// GetArmorDc is a free data retrieval call binding the contract method 0xb4e2a2eb.
//
// Solidity: function get_armor_dc(uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingICaller) GetArmorDc(opts *bind.CallOpts, _item_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_armor_dc", _item_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArmorDc is a free data retrieval call binding the contract method 0xb4e2a2eb.
//
// Solidity: function get_armor_dc(uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingISession) GetArmorDc(_item_id *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetArmorDc(&_CraftingI.CallOpts, _item_id)
}

// GetArmorDc is a free data retrieval call binding the contract method 0xb4e2a2eb.
//
// Solidity: function get_armor_dc(uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingICallerSession) GetArmorDc(_item_id *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetArmorDc(&_CraftingI.CallOpts, _item_id)
}

// GetDc is a free data retrieval call binding the contract method 0xe9f5ba75.
//
// Solidity: function get_dc(uint256 _base_type, uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingICaller) GetDc(opts *bind.CallOpts, _base_type *big.Int, _item_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_dc", _base_type, _item_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDc is a free data retrieval call binding the contract method 0xe9f5ba75.
//
// Solidity: function get_dc(uint256 _base_type, uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingISession) GetDc(_base_type *big.Int, _item_id *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetDc(&_CraftingI.CallOpts, _base_type, _item_id)
}

// GetDc is a free data retrieval call binding the contract method 0xe9f5ba75.
//
// Solidity: function get_dc(uint256 _base_type, uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingICallerSession) GetDc(_base_type *big.Int, _item_id *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetDc(&_CraftingI.CallOpts, _base_type, _item_id)
}

// GetGoodsDc is a free data retrieval call binding the contract method 0xaa1b44e7.
//
// Solidity: function get_goods_dc() pure returns(uint256 dc)
func (_CraftingI *CraftingICaller) GetGoodsDc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_goods_dc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGoodsDc is a free data retrieval call binding the contract method 0xaa1b44e7.
//
// Solidity: function get_goods_dc() pure returns(uint256 dc)
func (_CraftingI *CraftingISession) GetGoodsDc() (*big.Int, error) {
	return _CraftingI.Contract.GetGoodsDc(&_CraftingI.CallOpts)
}

// GetGoodsDc is a free data retrieval call binding the contract method 0xaa1b44e7.
//
// Solidity: function get_goods_dc() pure returns(uint256 dc)
func (_CraftingI *CraftingICallerSession) GetGoodsDc() (*big.Int, error) {
	return _CraftingI.Contract.GetGoodsDc(&_CraftingI.CallOpts)
}

// GetItemCost is a free data retrieval call binding the contract method 0x6b1b111f.
//
// Solidity: function get_item_cost(uint256 _base_type, uint256 _item_type) pure returns(uint256 cost)
func (_CraftingI *CraftingICaller) GetItemCost(opts *bind.CallOpts, _base_type *big.Int, _item_type *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_item_cost", _base_type, _item_type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetItemCost is a free data retrieval call binding the contract method 0x6b1b111f.
//
// Solidity: function get_item_cost(uint256 _base_type, uint256 _item_type) pure returns(uint256 cost)
func (_CraftingI *CraftingISession) GetItemCost(_base_type *big.Int, _item_type *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetItemCost(&_CraftingI.CallOpts, _base_type, _item_type)
}

// GetItemCost is a free data retrieval call binding the contract method 0x6b1b111f.
//
// Solidity: function get_item_cost(uint256 _base_type, uint256 _item_type) pure returns(uint256 cost)
func (_CraftingI *CraftingICallerSession) GetItemCost(_base_type *big.Int, _item_type *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetItemCost(&_CraftingI.CallOpts, _base_type, _item_type)
}

// GetTokenUriArmor is a free data retrieval call binding the contract method 0xf3d7e2e5.
//
// Solidity: function get_token_uri_armor(uint256 _item) view returns(string output)
func (_CraftingI *CraftingICaller) GetTokenUriArmor(opts *bind.CallOpts, _item *big.Int) (string, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_token_uri_armor", _item)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenUriArmor is a free data retrieval call binding the contract method 0xf3d7e2e5.
//
// Solidity: function get_token_uri_armor(uint256 _item) view returns(string output)
func (_CraftingI *CraftingISession) GetTokenUriArmor(_item *big.Int) (string, error) {
	return _CraftingI.Contract.GetTokenUriArmor(&_CraftingI.CallOpts, _item)
}

// GetTokenUriArmor is a free data retrieval call binding the contract method 0xf3d7e2e5.
//
// Solidity: function get_token_uri_armor(uint256 _item) view returns(string output)
func (_CraftingI *CraftingICallerSession) GetTokenUriArmor(_item *big.Int) (string, error) {
	return _CraftingI.Contract.GetTokenUriArmor(&_CraftingI.CallOpts, _item)
}

// GetTokenUriGoods is a free data retrieval call binding the contract method 0x070daa6b.
//
// Solidity: function get_token_uri_goods(uint256 _item) view returns(string output)
func (_CraftingI *CraftingICaller) GetTokenUriGoods(opts *bind.CallOpts, _item *big.Int) (string, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_token_uri_goods", _item)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenUriGoods is a free data retrieval call binding the contract method 0x070daa6b.
//
// Solidity: function get_token_uri_goods(uint256 _item) view returns(string output)
func (_CraftingI *CraftingISession) GetTokenUriGoods(_item *big.Int) (string, error) {
	return _CraftingI.Contract.GetTokenUriGoods(&_CraftingI.CallOpts, _item)
}

// GetTokenUriGoods is a free data retrieval call binding the contract method 0x070daa6b.
//
// Solidity: function get_token_uri_goods(uint256 _item) view returns(string output)
func (_CraftingI *CraftingICallerSession) GetTokenUriGoods(_item *big.Int) (string, error) {
	return _CraftingI.Contract.GetTokenUriGoods(&_CraftingI.CallOpts, _item)
}

// GetTokenUriWeapon is a free data retrieval call binding the contract method 0xc32825d7.
//
// Solidity: function get_token_uri_weapon(uint256 _item) view returns(string output)
func (_CraftingI *CraftingICaller) GetTokenUriWeapon(opts *bind.CallOpts, _item *big.Int) (string, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_token_uri_weapon", _item)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenUriWeapon is a free data retrieval call binding the contract method 0xc32825d7.
//
// Solidity: function get_token_uri_weapon(uint256 _item) view returns(string output)
func (_CraftingI *CraftingISession) GetTokenUriWeapon(_item *big.Int) (string, error) {
	return _CraftingI.Contract.GetTokenUriWeapon(&_CraftingI.CallOpts, _item)
}

// GetTokenUriWeapon is a free data retrieval call binding the contract method 0xc32825d7.
//
// Solidity: function get_token_uri_weapon(uint256 _item) view returns(string output)
func (_CraftingI *CraftingICallerSession) GetTokenUriWeapon(_item *big.Int) (string, error) {
	return _CraftingI.Contract.GetTokenUriWeapon(&_CraftingI.CallOpts, _item)
}

// GetType is a free data retrieval call binding the contract method 0x74b0b6cf.
//
// Solidity: function get_type(uint256 _type_id) pure returns(string _type)
func (_CraftingI *CraftingICaller) GetType(opts *bind.CallOpts, _type_id *big.Int) (string, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_type", _type_id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0x74b0b6cf.
//
// Solidity: function get_type(uint256 _type_id) pure returns(string _type)
func (_CraftingI *CraftingISession) GetType(_type_id *big.Int) (string, error) {
	return _CraftingI.Contract.GetType(&_CraftingI.CallOpts, _type_id)
}

// GetType is a free data retrieval call binding the contract method 0x74b0b6cf.
//
// Solidity: function get_type(uint256 _type_id) pure returns(string _type)
func (_CraftingI *CraftingICallerSession) GetType(_type_id *big.Int) (string, error) {
	return _CraftingI.Contract.GetType(&_CraftingI.CallOpts, _type_id)
}

// GetWeaponDc is a free data retrieval call binding the contract method 0x16cb2b52.
//
// Solidity: function get_weapon_dc(uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingICaller) GetWeaponDc(opts *bind.CallOpts, _item_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "get_weapon_dc", _item_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWeaponDc is a free data retrieval call binding the contract method 0x16cb2b52.
//
// Solidity: function get_weapon_dc(uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingISession) GetWeaponDc(_item_id *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetWeaponDc(&_CraftingI.CallOpts, _item_id)
}

// GetWeaponDc is a free data retrieval call binding the contract method 0x16cb2b52.
//
// Solidity: function get_weapon_dc(uint256 _item_id) pure returns(uint256 dc)
func (_CraftingI *CraftingICallerSession) GetWeaponDc(_item_id *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.GetWeaponDc(&_CraftingI.CallOpts, _item_id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CraftingI *CraftingICaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CraftingI *CraftingISession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CraftingI.Contract.IsApprovedForAll(&_CraftingI.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CraftingI *CraftingICallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CraftingI.Contract.IsApprovedForAll(&_CraftingI.CallOpts, owner, operator)
}

// IsValid is a free data retrieval call binding the contract method 0xad507ea7.
//
// Solidity: function isValid(uint256 _base_type, uint256 _item_type) pure returns(bool)
func (_CraftingI *CraftingICaller) IsValid(opts *bind.CallOpts, _base_type *big.Int, _item_type *big.Int) (bool, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "isValid", _base_type, _item_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0xad507ea7.
//
// Solidity: function isValid(uint256 _base_type, uint256 _item_type) pure returns(bool)
func (_CraftingI *CraftingISession) IsValid(_base_type *big.Int, _item_type *big.Int) (bool, error) {
	return _CraftingI.Contract.IsValid(&_CraftingI.CallOpts, _base_type, _item_type)
}

// IsValid is a free data retrieval call binding the contract method 0xad507ea7.
//
// Solidity: function isValid(uint256 _base_type, uint256 _item_type) pure returns(bool)
func (_CraftingI *CraftingICallerSession) IsValid(_base_type *big.Int, _item_type *big.Int) (bool, error) {
	return _CraftingI.Contract.IsValid(&_CraftingI.CallOpts, _base_type, _item_type)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint8 base_type, uint8 item_type, uint32 crafted, uint256 crafter)
func (_CraftingI *CraftingICaller) Items(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BaseType uint8
	ItemType uint8
	Crafted  uint32
	Crafter  *big.Int
}, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "items", arg0)

	outstruct := new(struct {
		BaseType uint8
		ItemType uint8
		Crafted  uint32
		Crafter  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.ItemType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Crafted = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Crafter = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint8 base_type, uint8 item_type, uint32 crafted, uint256 crafter)
func (_CraftingI *CraftingISession) Items(arg0 *big.Int) (struct {
	BaseType uint8
	ItemType uint8
	Crafted  uint32
	Crafter  *big.Int
}, error) {
	return _CraftingI.Contract.Items(&_CraftingI.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint8 base_type, uint8 item_type, uint32 crafted, uint256 crafter)
func (_CraftingI *CraftingICallerSession) Items(arg0 *big.Int) (struct {
	BaseType uint8
	ItemType uint8
	Crafted  uint32
	Crafter  *big.Int
}, error) {
	return _CraftingI.Contract.Items(&_CraftingI.CallOpts, arg0)
}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_CraftingI *CraftingICaller) ModifierForAttribute(opts *bind.CallOpts, _attribute *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "modifier_for_attribute", _attribute)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_CraftingI *CraftingISession) ModifierForAttribute(_attribute *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.ModifierForAttribute(&_CraftingI.CallOpts, _attribute)
}

// ModifierForAttribute is a free data retrieval call binding the contract method 0x927330db.
//
// Solidity: function modifier_for_attribute(uint256 _attribute) pure returns(int256 _modifier)
func (_CraftingI *CraftingICallerSession) ModifierForAttribute(_attribute *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.ModifierForAttribute(&_CraftingI.CallOpts, _attribute)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CraftingI *CraftingICaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CraftingI *CraftingISession) Name() (string, error) {
	return _CraftingI.Contract.Name(&_CraftingI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CraftingI *CraftingICallerSession) Name() (string, error) {
	return _CraftingI.Contract.Name(&_CraftingI.CallOpts)
}

// NextItem is a free data retrieval call binding the contract method 0x0bf3bcba.
//
// Solidity: function next_item() view returns(uint256)
func (_CraftingI *CraftingICaller) NextItem(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "next_item")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextItem is a free data retrieval call binding the contract method 0x0bf3bcba.
//
// Solidity: function next_item() view returns(uint256)
func (_CraftingI *CraftingISession) NextItem() (*big.Int, error) {
	return _CraftingI.Contract.NextItem(&_CraftingI.CallOpts)
}

// NextItem is a free data retrieval call binding the contract method 0x0bf3bcba.
//
// Solidity: function next_item() view returns(uint256)
func (_CraftingI *CraftingICallerSession) NextItem() (*big.Int, error) {
	return _CraftingI.Contract.NextItem(&_CraftingI.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CraftingI *CraftingICaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CraftingI *CraftingISession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CraftingI.Contract.OwnerOf(&_CraftingI.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CraftingI *CraftingICallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CraftingI.Contract.OwnerOf(&_CraftingI.CallOpts, tokenId)
}

// Simulate is a free data retrieval call binding the contract method 0xc1428354.
//
// Solidity: function simulate(uint256 _summoner, uint256 _base_type, uint256 _item_type, uint256 _crafting_materials) view returns(bool crafted, int256 check, uint256 cost, uint256 dc)
func (_CraftingI *CraftingICaller) Simulate(opts *bind.CallOpts, _summoner *big.Int, _base_type *big.Int, _item_type *big.Int, _crafting_materials *big.Int) (struct {
	Crafted bool
	Check   *big.Int
	Cost    *big.Int
	Dc      *big.Int
}, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "simulate", _summoner, _base_type, _item_type, _crafting_materials)

	outstruct := new(struct {
		Crafted bool
		Check   *big.Int
		Cost    *big.Int
		Dc      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Crafted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Check = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Cost = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Dc = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Simulate is a free data retrieval call binding the contract method 0xc1428354.
//
// Solidity: function simulate(uint256 _summoner, uint256 _base_type, uint256 _item_type, uint256 _crafting_materials) view returns(bool crafted, int256 check, uint256 cost, uint256 dc)
func (_CraftingI *CraftingISession) Simulate(_summoner *big.Int, _base_type *big.Int, _item_type *big.Int, _crafting_materials *big.Int) (struct {
	Crafted bool
	Check   *big.Int
	Cost    *big.Int
	Dc      *big.Int
}, error) {
	return _CraftingI.Contract.Simulate(&_CraftingI.CallOpts, _summoner, _base_type, _item_type, _crafting_materials)
}

// Simulate is a free data retrieval call binding the contract method 0xc1428354.
//
// Solidity: function simulate(uint256 _summoner, uint256 _base_type, uint256 _item_type, uint256 _crafting_materials) view returns(bool crafted, int256 check, uint256 cost, uint256 dc)
func (_CraftingI *CraftingICallerSession) Simulate(_summoner *big.Int, _base_type *big.Int, _item_type *big.Int, _crafting_materials *big.Int) (struct {
	Crafted bool
	Check   *big.Int
	Cost    *big.Int
	Dc      *big.Int
}, error) {
	return _CraftingI.Contract.Simulate(&_CraftingI.CallOpts, _summoner, _base_type, _item_type, _crafting_materials)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CraftingI *CraftingICaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CraftingI *CraftingISession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CraftingI.Contract.SupportsInterface(&_CraftingI.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CraftingI *CraftingICallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CraftingI.Contract.SupportsInterface(&_CraftingI.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CraftingI *CraftingICaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CraftingI *CraftingISession) Symbol() (string, error) {
	return _CraftingI.Contract.Symbol(&_CraftingI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CraftingI *CraftingICallerSession) Symbol() (string, error) {
	return _CraftingI.Contract.Symbol(&_CraftingI.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CraftingI *CraftingICaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CraftingI *CraftingISession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.TokenByIndex(&_CraftingI.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CraftingI *CraftingICallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.TokenByIndex(&_CraftingI.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CraftingI *CraftingICaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CraftingI *CraftingISession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.TokenOfOwnerByIndex(&_CraftingI.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CraftingI *CraftingICallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _CraftingI.Contract.TokenOfOwnerByIndex(&_CraftingI.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _item) view returns(string uri)
func (_CraftingI *CraftingICaller) TokenURI(opts *bind.CallOpts, _item *big.Int) (string, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "tokenURI", _item)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _item) view returns(string uri)
func (_CraftingI *CraftingISession) TokenURI(_item *big.Int) (string, error) {
	return _CraftingI.Contract.TokenURI(&_CraftingI.CallOpts, _item)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _item) view returns(string uri)
func (_CraftingI *CraftingICallerSession) TokenURI(_item *big.Int) (string, error) {
	return _CraftingI.Contract.TokenURI(&_CraftingI.CallOpts, _item)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CraftingI *CraftingICaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CraftingI.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CraftingI *CraftingISession) TotalSupply() (*big.Int, error) {
	return _CraftingI.Contract.TotalSupply(&_CraftingI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CraftingI *CraftingICallerSession) TotalSupply() (*big.Int, error) {
	return _CraftingI.Contract.TotalSupply(&_CraftingI.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CraftingI *CraftingITransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CraftingI *CraftingISession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.Approve(&_CraftingI.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CraftingI *CraftingITransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.Approve(&_CraftingI.TransactOpts, to, tokenId)
}

// Craft is a paid mutator transaction binding the contract method 0x65f1bfc4.
//
// Solidity: function craft(uint256 _summoner, uint8 _base_type, uint8 _item_type, uint256 _crafting_materials) returns()
func (_CraftingI *CraftingITransactor) Craft(opts *bind.TransactOpts, _summoner *big.Int, _base_type uint8, _item_type uint8, _crafting_materials *big.Int) (*types.Transaction, error) {
	return _CraftingI.contract.Transact(opts, "craft", _summoner, _base_type, _item_type, _crafting_materials)
}

// Craft is a paid mutator transaction binding the contract method 0x65f1bfc4.
//
// Solidity: function craft(uint256 _summoner, uint8 _base_type, uint8 _item_type, uint256 _crafting_materials) returns()
func (_CraftingI *CraftingISession) Craft(_summoner *big.Int, _base_type uint8, _item_type uint8, _crafting_materials *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.Craft(&_CraftingI.TransactOpts, _summoner, _base_type, _item_type, _crafting_materials)
}

// Craft is a paid mutator transaction binding the contract method 0x65f1bfc4.
//
// Solidity: function craft(uint256 _summoner, uint8 _base_type, uint8 _item_type, uint256 _crafting_materials) returns()
func (_CraftingI *CraftingITransactorSession) Craft(_summoner *big.Int, _base_type uint8, _item_type uint8, _crafting_materials *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.Craft(&_CraftingI.TransactOpts, _summoner, _base_type, _item_type, _crafting_materials)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CraftingI *CraftingITransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CraftingI *CraftingISession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.SafeTransferFrom(&_CraftingI.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CraftingI *CraftingITransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.SafeTransferFrom(&_CraftingI.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CraftingI *CraftingITransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CraftingI.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CraftingI *CraftingISession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CraftingI.Contract.SafeTransferFrom0(&_CraftingI.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CraftingI *CraftingITransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CraftingI.Contract.SafeTransferFrom0(&_CraftingI.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CraftingI *CraftingITransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CraftingI.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CraftingI *CraftingISession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CraftingI.Contract.SetApprovalForAll(&_CraftingI.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CraftingI *CraftingITransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CraftingI.Contract.SetApprovalForAll(&_CraftingI.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CraftingI *CraftingITransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CraftingI *CraftingISession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.TransferFrom(&_CraftingI.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CraftingI *CraftingITransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CraftingI.Contract.TransferFrom(&_CraftingI.TransactOpts, from, to, tokenId)
}

// CraftingIApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CraftingI contract.
type CraftingIApprovalIterator struct {
	Event *CraftingIApproval // Event containing the contract specifics and raw log

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
func (it *CraftingIApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingIApproval)
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
		it.Event = new(CraftingIApproval)
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
func (it *CraftingIApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingIApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingIApproval represents a Approval event raised by the CraftingI contract.
type CraftingIApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CraftingI *CraftingIFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CraftingIApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CraftingI.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CraftingIApprovalIterator{contract: _CraftingI.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CraftingI *CraftingIFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CraftingIApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CraftingI.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingIApproval)
				if err := _CraftingI.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CraftingI *CraftingIFilterer) ParseApproval(log types.Log) (*CraftingIApproval, error) {
	event := new(CraftingIApproval)
	if err := _CraftingI.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingIApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CraftingI contract.
type CraftingIApprovalForAllIterator struct {
	Event *CraftingIApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CraftingIApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingIApprovalForAll)
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
		it.Event = new(CraftingIApprovalForAll)
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
func (it *CraftingIApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingIApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingIApprovalForAll represents a ApprovalForAll event raised by the CraftingI contract.
type CraftingIApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CraftingI *CraftingIFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CraftingIApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CraftingI.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CraftingIApprovalForAllIterator{contract: _CraftingI.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CraftingI *CraftingIFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CraftingIApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CraftingI.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingIApprovalForAll)
				if err := _CraftingI.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CraftingI *CraftingIFilterer) ParseApprovalForAll(log types.Log) (*CraftingIApprovalForAll, error) {
	event := new(CraftingIApprovalForAll)
	if err := _CraftingI.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingICraftedIterator is returned from FilterCrafted and is used to iterate over the raw logs and unpacked data for Crafted events raised by the CraftingI contract.
type CraftingICraftedIterator struct {
	Event *CraftingICrafted // Event containing the contract specifics and raw log

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
func (it *CraftingICraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingICrafted)
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
		it.Event = new(CraftingICrafted)
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
func (it *CraftingICraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingICraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingICrafted represents a Crafted event raised by the CraftingI contract.
type CraftingICrafted struct {
	Owner    common.Address
	Check    *big.Int
	Summoner *big.Int
	BaseType *big.Int
	ItemType *big.Int
	Gold     *big.Int
	CraftI   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCrafted is a free log retrieval operation binding the contract event 0x9adacb380363bfeed846fc4d134cba010cd0082bc0981f54b5edf7353ab0ea62.
//
// Solidity: event Crafted(address indexed owner, uint256 check, uint256 summoner, uint256 base_type, uint256 item_type, uint256 gold, uint256 craft_i)
func (_CraftingI *CraftingIFilterer) FilterCrafted(opts *bind.FilterOpts, owner []common.Address) (*CraftingICraftedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CraftingI.contract.FilterLogs(opts, "Crafted", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CraftingICraftedIterator{contract: _CraftingI.contract, event: "Crafted", logs: logs, sub: sub}, nil
}

// WatchCrafted is a free log subscription operation binding the contract event 0x9adacb380363bfeed846fc4d134cba010cd0082bc0981f54b5edf7353ab0ea62.
//
// Solidity: event Crafted(address indexed owner, uint256 check, uint256 summoner, uint256 base_type, uint256 item_type, uint256 gold, uint256 craft_i)
func (_CraftingI *CraftingIFilterer) WatchCrafted(opts *bind.WatchOpts, sink chan<- *CraftingICrafted, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CraftingI.contract.WatchLogs(opts, "Crafted", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingICrafted)
				if err := _CraftingI.contract.UnpackLog(event, "Crafted", log); err != nil {
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

// ParseCrafted is a log parse operation binding the contract event 0x9adacb380363bfeed846fc4d134cba010cd0082bc0981f54b5edf7353ab0ea62.
//
// Solidity: event Crafted(address indexed owner, uint256 check, uint256 summoner, uint256 base_type, uint256 item_type, uint256 gold, uint256 craft_i)
func (_CraftingI *CraftingIFilterer) ParseCrafted(log types.Log) (*CraftingICrafted, error) {
	event := new(CraftingICrafted)
	if err := _CraftingI.contract.UnpackLog(event, "Crafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingITransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CraftingI contract.
type CraftingITransferIterator struct {
	Event *CraftingITransfer // Event containing the contract specifics and raw log

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
func (it *CraftingITransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingITransfer)
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
		it.Event = new(CraftingITransfer)
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
func (it *CraftingITransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingITransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingITransfer represents a Transfer event raised by the CraftingI contract.
type CraftingITransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CraftingI *CraftingIFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CraftingITransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CraftingI.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CraftingITransferIterator{contract: _CraftingI.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CraftingI *CraftingIFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CraftingITransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CraftingI.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingITransfer)
				if err := _CraftingI.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CraftingI *CraftingIFilterer) ParseTransfer(log types.Log) (*CraftingITransfer, error) {
	event := new(CraftingITransfer)
	if err := _CraftingI.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
