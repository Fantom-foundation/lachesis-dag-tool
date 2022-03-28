// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package logemitter

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"str1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"str2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"str3\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"str4\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"str5\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"str6\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"str7\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"str8\",\"type\":\"bytes32\"}],\"name\":\"Logging\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"str1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"str2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"str3\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"str4\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"str5\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"str6\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"str7\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"str8\",\"type\":\"bytes32\"}],\"name\":\"emitLogs\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506101038061001f6000396000f3fe608060405260043610601c5760003560e01c8063fe74db15146021575b600080fd5b60666004803603610100811015603657600080fd5b5080359060208101359060408101359060608101359060808101359060a08101359060c08101359060e001356068565b005b6040805187815260208101879052808201869052606081018590526080810184905260a08101839052905188918a9133917ff26c641f7770083375ce931e80394fa6883abdee6a053d72c3241735b52b04ae919081900360c00190a4505050505050505056fea265627a7a72315820b5d22fb927b376479b434e657f5cf46808e6641d3a36275624de4151d5e42b5a64736f6c634300050c0032",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// EmitLogs is a paid mutator transaction binding the contract method 0xfe74db15.
//
// Solidity: function emitLogs(bytes32 str1, bytes32 str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8) payable returns()
func (_Contract *ContractTransactor) EmitLogs(opts *bind.TransactOpts, str1 [32]byte, str2 [32]byte, str3 [32]byte, str4 [32]byte, str5 [32]byte, str6 [32]byte, str7 [32]byte, str8 [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "emitLogs", str1, str2, str3, str4, str5, str6, str7, str8)
}

// EmitLogs is a paid mutator transaction binding the contract method 0xfe74db15.
//
// Solidity: function emitLogs(bytes32 str1, bytes32 str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8) payable returns()
func (_Contract *ContractSession) EmitLogs(str1 [32]byte, str2 [32]byte, str3 [32]byte, str4 [32]byte, str5 [32]byte, str6 [32]byte, str7 [32]byte, str8 [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.EmitLogs(&_Contract.TransactOpts, str1, str2, str3, str4, str5, str6, str7, str8)
}

// EmitLogs is a paid mutator transaction binding the contract method 0xfe74db15.
//
// Solidity: function emitLogs(bytes32 str1, bytes32 str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8) payable returns()
func (_Contract *ContractTransactorSession) EmitLogs(str1 [32]byte, str2 [32]byte, str3 [32]byte, str4 [32]byte, str5 [32]byte, str6 [32]byte, str7 [32]byte, str8 [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.EmitLogs(&_Contract.TransactOpts, str1, str2, str3, str4, str5, str6, str7, str8)
}

// ContractLoggingIterator is returned from FilterLogging and is used to iterate over the raw logs and unpacked data for Logging events raised by the Contract contract.
type ContractLoggingIterator struct {
	Event *ContractLogging // Event containing the contract specifics and raw log

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
func (it *ContractLoggingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogging)
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
		it.Event = new(ContractLogging)
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
func (it *ContractLoggingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLoggingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogging represents a Logging event raised by the Contract contract.
type ContractLogging struct {
	Sender common.Address
	Str1   [32]byte
	Str2   [32]byte
	Str3   [32]byte
	Str4   [32]byte
	Str5   [32]byte
	Str6   [32]byte
	Str7   [32]byte
	Str8   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogging is a free log retrieval operation binding the contract event 0xf26c641f7770083375ce931e80394fa6883abdee6a053d72c3241735b52b04ae.
//
// Solidity: event Logging(address indexed sender, bytes32 indexed str1, bytes32 indexed str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8)
func (_Contract *ContractFilterer) FilterLogging(opts *bind.FilterOpts, sender []common.Address, str1 [][32]byte, str2 [][32]byte) (*ContractLoggingIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var str1Rule []interface{}
	for _, str1Item := range str1 {
		str1Rule = append(str1Rule, str1Item)
	}
	var str2Rule []interface{}
	for _, str2Item := range str2 {
		str2Rule = append(str2Rule, str2Item)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Logging", senderRule, str1Rule, str2Rule)
	if err != nil {
		return nil, err
	}
	return &ContractLoggingIterator{contract: _Contract.contract, event: "Logging", logs: logs, sub: sub}, nil
}

// WatchLogging is a free log subscription operation binding the contract event 0xf26c641f7770083375ce931e80394fa6883abdee6a053d72c3241735b52b04ae.
//
// Solidity: event Logging(address indexed sender, bytes32 indexed str1, bytes32 indexed str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8)
func (_Contract *ContractFilterer) WatchLogging(opts *bind.WatchOpts, sink chan<- *ContractLogging, sender []common.Address, str1 [][32]byte, str2 [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var str1Rule []interface{}
	for _, str1Item := range str1 {
		str1Rule = append(str1Rule, str1Item)
	}
	var str2Rule []interface{}
	for _, str2Item := range str2 {
		str2Rule = append(str2Rule, str2Item)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Logging", senderRule, str1Rule, str2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogging)
				if err := _Contract.contract.UnpackLog(event, "Logging", log); err != nil {
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

// ParseLogging is a log parse operation binding the contract event 0xf26c641f7770083375ce931e80394fa6883abdee6a053d72c3241735b52b04ae.
//
// Solidity: event Logging(address indexed sender, bytes32 indexed str1, bytes32 indexed str2, bytes32 str3, bytes32 str4, bytes32 str5, bytes32 str6, bytes32 str7, bytes32 str8)
func (_Contract *ContractFilterer) ParseLogging(log types.Log) (*ContractLogging, error) {
	event := new(ContractLogging)
	if err := _Contract.contract.UnpackLog(event, "Logging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
