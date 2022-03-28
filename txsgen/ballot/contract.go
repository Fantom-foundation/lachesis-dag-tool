// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ballot

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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"text\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Voiting\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"winnerName_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winningProposal_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516104673803806104678339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b5050505090500160405250505060008090505b815181101561012757600160405180604001604052808484815181106100e757fe5b60209081029190910181015182526000918101829052835460018181018655948352918190208351600290930201918255919091015190820155016100c5565b505061032f806101386000396000f3fe60806040526004361061005a5760003560e01c8063609ff1bd11610043578063609ff1bd146100c1578063a3ec138d146100e8578063e2ba53f0146101285761005a565b80630121b93f1461005f578063013cf08b1461007e575b600080fd5b61007c6004803603602081101561007557600080fd5b503561013d565b005b34801561008a57600080fd5b506100a8600480360360208110156100a157600080fd5b5035610229565b6040805192835260208301919091528051918290030190f35b3480156100cd57600080fd5b506100d6610254565b60408051918252519081900360200190f35b3480156100f457600080fd5b506100d66004803603602081101561010b57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166102bb565b34801561013457600080fd5b506100d66102cd565b33600090815260208190526040902054801561018057600180600183038154811061016457fe5b6000918252602090912060016002909202010180549190910390555b3360009081526020526001805481908490811061019957fe5b906000526020600020906002020160010160008282540192505081905550600182815481106101c457fe5b906000526020600020906002020160000154823373ffffffffffffffffffffffffffffffffffffffff167fee51d93e4784cfee8c8b14400fd5c768d31980d6a648124618b164b9b3dd69e1346040518082815260200191505060405180910390a45050565b6001818154811061023657fe5b60009182526020909120600290910201805460019091015490915082565b600080805b6001548110156102b657816001828154811061027157fe5b90600052602060002090600202016001015411156102ae576001818154811061029657fe5b90600052602060002090600202016001015491508092505b600101610259565b505090565b60006020819052908152604090205481565b600060016102d9610254565b815481106102e357fe5b90600052602060002090600202016000015490509056fea265627a7a723158207d2c2287159dd480a6fc9af9edeae1b086492978db89fbc32be265a54a9c00d764736f6c634300050c0032",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, proposalNames [][32]byte) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, proposalNames)
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

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Contract *ContractCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Contract *ContractSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Contract.Contract.Proposals(&_Contract.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Contract *ContractCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Contract.Contract.Proposals(&_Contract.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 vote)
func (_Contract *ContractCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 vote)
func (_Contract *ContractSession) Voters(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Voters(&_Contract.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 vote)
func (_Contract *ContractCallerSession) Voters(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Voters(&_Contract.CallOpts, arg0)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Contract *ContractCaller) WinnerName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "winnerName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Contract *ContractSession) WinnerName() ([32]byte, error) {
	return _Contract.Contract.WinnerName(&_Contract.CallOpts)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Contract *ContractCallerSession) WinnerName() ([32]byte, error) {
	return _Contract.Contract.WinnerName(&_Contract.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Contract *ContractCaller) WinningProposal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "winningProposal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Contract *ContractSession) WinningProposal() (*big.Int, error) {
	return _Contract.Contract.WinningProposal(&_Contract.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Contract *ContractCallerSession) WinningProposal() (*big.Int, error) {
	return _Contract.Contract.WinningProposal(&_Contract.CallOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) payable returns()
func (_Contract *ContractTransactor) Vote(opts *bind.TransactOpts, proposal *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) payable returns()
func (_Contract *ContractSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) payable returns()
func (_Contract *ContractTransactorSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, proposal)
}

// ContractVoitingIterator is returned from FilterVoiting and is used to iterate over the raw logs and unpacked data for Voiting events raised by the Contract contract.
type ContractVoitingIterator struct {
	Event *ContractVoiting // Event containing the contract specifics and raw log

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
func (it *ContractVoitingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVoiting)
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
		it.Event = new(ContractVoiting)
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
func (it *ContractVoitingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVoitingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVoiting represents a Voiting event raised by the Contract contract.
type ContractVoiting struct {
	Who    common.Address
	Num    *big.Int
	Text   [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVoiting is a free log retrieval operation binding the contract event 0xee51d93e4784cfee8c8b14400fd5c768d31980d6a648124618b164b9b3dd69e1.
//
// Solidity: event Voiting(address indexed who, uint256 indexed num, bytes32 indexed text, uint256 amount)
func (_Contract *ContractFilterer) FilterVoiting(opts *bind.FilterOpts, who []common.Address, num []*big.Int, text [][32]byte) (*ContractVoitingIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var numRule []interface{}
	for _, numItem := range num {
		numRule = append(numRule, numItem)
	}
	var textRule []interface{}
	for _, textItem := range text {
		textRule = append(textRule, textItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Voiting", whoRule, numRule, textRule)
	if err != nil {
		return nil, err
	}
	return &ContractVoitingIterator{contract: _Contract.contract, event: "Voiting", logs: logs, sub: sub}, nil
}

// WatchVoiting is a free log subscription operation binding the contract event 0xee51d93e4784cfee8c8b14400fd5c768d31980d6a648124618b164b9b3dd69e1.
//
// Solidity: event Voiting(address indexed who, uint256 indexed num, bytes32 indexed text, uint256 amount)
func (_Contract *ContractFilterer) WatchVoiting(opts *bind.WatchOpts, sink chan<- *ContractVoiting, who []common.Address, num []*big.Int, text [][32]byte) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var numRule []interface{}
	for _, numItem := range num {
		numRule = append(numRule, numItem)
	}
	var textRule []interface{}
	for _, textItem := range text {
		textRule = append(textRule, textItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Voiting", whoRule, numRule, textRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVoiting)
				if err := _Contract.contract.UnpackLog(event, "Voiting", log); err != nil {
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

// ParseVoiting is a log parse operation binding the contract event 0xee51d93e4784cfee8c8b14400fd5c768d31980d6a648124618b164b9b3dd69e1.
//
// Solidity: event Voiting(address indexed who, uint256 indexed num, bytes32 indexed text, uint256 amount)
func (_Contract *ContractFilterer) ParseVoiting(log types.Log) (*ContractVoiting, error) {
	event := new(ContractVoiting)
	if err := _Contract.contract.UnpackLog(event, "Voiting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
