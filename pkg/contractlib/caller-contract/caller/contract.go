// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package caller

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CallerABI is the input ABI used to generate the binding from.
const CallerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"adventure\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"approve_all\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_claimers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_need_approval\",\"type\":\"uint256[]\"}],\"name\":\"claim_gold\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_delvers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_need_approval\",\"type\":\"uint256[]\"}],\"name\":\"dungeon\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"is_approved\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_is_approved\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"level_up\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Caller is an auto generated Go binding around an Ethereum contract.
type Caller struct {
	CallerCaller     // Read-only binding to the contract
	CallerTransactor // Write-only binding to the contract
	CallerFilterer   // Log filterer for contract events
}

// CallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallerSession struct {
	Contract     *Caller           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallerCallerSession struct {
	Contract *CallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallerTransactorSession struct {
	Contract     *CallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallerRaw struct {
	Contract *Caller // Generic contract binding to access the raw methods on
}

// CallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallerCallerRaw struct {
	Contract *CallerCaller // Generic read-only contract binding to access the raw methods on
}

// CallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallerTransactorRaw struct {
	Contract *CallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaller creates a new instance of Caller, bound to a specific deployed contract.
func NewCaller(address common.Address, backend bind.ContractBackend) (*Caller, error) {
	contract, err := bindCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// NewCallerCaller creates a new read-only instance of Caller, bound to a specific deployed contract.
func NewCallerCaller(address common.Address, caller bind.ContractCaller) (*CallerCaller, error) {
	contract, err := bindCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallerCaller{contract: contract}, nil
}

// NewCallerTransactor creates a new write-only instance of Caller, bound to a specific deployed contract.
func NewCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CallerTransactor, error) {
	contract, err := bindCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallerTransactor{contract: contract}, nil
}

// NewCallerFilterer creates a new log filterer instance of Caller, bound to a specific deployed contract.
func NewCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CallerFilterer, error) {
	contract, err := bindCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallerFilterer{contract: contract}, nil
}

// bindCaller binds a generic wrapper to an already deployed contract.
func bindCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.CallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transact(opts, method, params...)
}

// IsApproved is a free data retrieval call binding the contract method 0xe708c494.
//
// Solidity: function is_approved(uint256[] _ids) view returns(bool[] _is_approved)
func (_Caller *CallerCaller) IsApproved(opts *bind.CallOpts, _ids []*big.Int) ([]bool, error) {
	var out []interface{}
	err := _Caller.contract.Call(opts, &out, "is_approved", _ids)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0xe708c494.
//
// Solidity: function is_approved(uint256[] _ids) view returns(bool[] _is_approved)
func (_Caller *CallerSession) IsApproved(_ids []*big.Int) ([]bool, error) {
	return _Caller.Contract.IsApproved(&_Caller.CallOpts, _ids)
}

// IsApproved is a free data retrieval call binding the contract method 0xe708c494.
//
// Solidity: function is_approved(uint256[] _ids) view returns(bool[] _is_approved)
func (_Caller *CallerCallerSession) IsApproved(_ids []*big.Int) ([]bool, error) {
	return _Caller.Contract.IsApproved(&_Caller.CallOpts, _ids)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Caller *CallerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Caller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Caller *CallerSession) Owner() (common.Address, error) {
	return _Caller.Contract.Owner(&_Caller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Caller *CallerCallerSession) Owner() (common.Address, error) {
	return _Caller.Contract.Owner(&_Caller.CallOpts)
}

// Adventure is a paid mutator transaction binding the contract method 0x3a091650.
//
// Solidity: function adventure(uint256[] _ids) payable returns()
func (_Caller *CallerTransactor) Adventure(opts *bind.TransactOpts, _ids []*big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "adventure", _ids)
}

// Adventure is a paid mutator transaction binding the contract method 0x3a091650.
//
// Solidity: function adventure(uint256[] _ids) payable returns()
func (_Caller *CallerSession) Adventure(_ids []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.Adventure(&_Caller.TransactOpts, _ids)
}

// Adventure is a paid mutator transaction binding the contract method 0x3a091650.
//
// Solidity: function adventure(uint256[] _ids) payable returns()
func (_Caller *CallerTransactorSession) Adventure(_ids []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.Adventure(&_Caller.TransactOpts, _ids)
}

// ApproveAll is a paid mutator transaction binding the contract method 0x33dc76f9.
//
// Solidity: function approve_all(uint256[] _ids) payable returns()
func (_Caller *CallerTransactor) ApproveAll(opts *bind.TransactOpts, _ids []*big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "approve_all", _ids)
}

// ApproveAll is a paid mutator transaction binding the contract method 0x33dc76f9.
//
// Solidity: function approve_all(uint256[] _ids) payable returns()
func (_Caller *CallerSession) ApproveAll(_ids []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.ApproveAll(&_Caller.TransactOpts, _ids)
}

// ApproveAll is a paid mutator transaction binding the contract method 0x33dc76f9.
//
// Solidity: function approve_all(uint256[] _ids) payable returns()
func (_Caller *CallerTransactorSession) ApproveAll(_ids []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.ApproveAll(&_Caller.TransactOpts, _ids)
}

// ClaimGold is a paid mutator transaction binding the contract method 0x6ba5d979.
//
// Solidity: function claim_gold(uint256[] _claimers, uint256[] _need_approval) payable returns()
func (_Caller *CallerTransactor) ClaimGold(opts *bind.TransactOpts, _claimers []*big.Int, _need_approval []*big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "claim_gold", _claimers, _need_approval)
}

// ClaimGold is a paid mutator transaction binding the contract method 0x6ba5d979.
//
// Solidity: function claim_gold(uint256[] _claimers, uint256[] _need_approval) payable returns()
func (_Caller *CallerSession) ClaimGold(_claimers []*big.Int, _need_approval []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.ClaimGold(&_Caller.TransactOpts, _claimers, _need_approval)
}

// ClaimGold is a paid mutator transaction binding the contract method 0x6ba5d979.
//
// Solidity: function claim_gold(uint256[] _claimers, uint256[] _need_approval) payable returns()
func (_Caller *CallerTransactorSession) ClaimGold(_claimers []*big.Int, _need_approval []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.ClaimGold(&_Caller.TransactOpts, _claimers, _need_approval)
}

// Dungeon is a paid mutator transaction binding the contract method 0xbe606afd.
//
// Solidity: function dungeon(uint256[] _delvers, uint256[] _need_approval) payable returns()
func (_Caller *CallerTransactor) Dungeon(opts *bind.TransactOpts, _delvers []*big.Int, _need_approval []*big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "dungeon", _delvers, _need_approval)
}

// Dungeon is a paid mutator transaction binding the contract method 0xbe606afd.
//
// Solidity: function dungeon(uint256[] _delvers, uint256[] _need_approval) payable returns()
func (_Caller *CallerSession) Dungeon(_delvers []*big.Int, _need_approval []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.Dungeon(&_Caller.TransactOpts, _delvers, _need_approval)
}

// Dungeon is a paid mutator transaction binding the contract method 0xbe606afd.
//
// Solidity: function dungeon(uint256[] _delvers, uint256[] _need_approval) payable returns()
func (_Caller *CallerTransactorSession) Dungeon(_delvers []*big.Int, _need_approval []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.Dungeon(&_Caller.TransactOpts, _delvers, _need_approval)
}

// LevelUp is a paid mutator transaction binding the contract method 0x13217f0e.
//
// Solidity: function level_up(uint256[] _ids) payable returns()
func (_Caller *CallerTransactor) LevelUp(opts *bind.TransactOpts, _ids []*big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "level_up", _ids)
}

// LevelUp is a paid mutator transaction binding the contract method 0x13217f0e.
//
// Solidity: function level_up(uint256[] _ids) payable returns()
func (_Caller *CallerSession) LevelUp(_ids []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.LevelUp(&_Caller.TransactOpts, _ids)
}

// LevelUp is a paid mutator transaction binding the contract method 0x13217f0e.
//
// Solidity: function level_up(uint256[] _ids) payable returns()
func (_Caller *CallerTransactorSession) LevelUp(_ids []*big.Int) (*types.Transaction, error) {
	return _Caller.Contract.LevelUp(&_Caller.TransactOpts, _ids)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Caller *CallerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Caller *CallerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Caller.Contract.RenounceOwnership(&_Caller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Caller *CallerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Caller.Contract.RenounceOwnership(&_Caller.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Caller *CallerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Caller *CallerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Caller.Contract.TransferOwnership(&_Caller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Caller *CallerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Caller.Contract.TransferOwnership(&_Caller.TransactOpts, newOwner)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_Caller *CallerTransactor) WithdrawFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "withdrawFees", to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_Caller *CallerSession) WithdrawFees(to common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WithdrawFees(&_Caller.TransactOpts, to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_Caller *CallerTransactorSession) WithdrawFees(to common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WithdrawFees(&_Caller.TransactOpts, to)
}

// CallerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Caller contract.
type CallerOwnershipTransferredIterator struct {
	Event *CallerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CallerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallerOwnershipTransferred)
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
		it.Event = new(CallerOwnershipTransferred)
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
func (it *CallerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallerOwnershipTransferred represents a OwnershipTransferred event raised by the Caller contract.
type CallerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Caller *CallerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CallerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Caller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CallerOwnershipTransferredIterator{contract: _Caller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Caller *CallerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CallerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Caller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallerOwnershipTransferred)
				if err := _Caller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Caller *CallerFilterer) ParseOwnershipTransferred(log types.Log) (*CallerOwnershipTransferred, error) {
	event := new(CallerOwnershipTransferred)
	if err := _Caller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
