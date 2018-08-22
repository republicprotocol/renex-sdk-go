// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x608060405234801561001057600080fd5b5061027c806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610082578063a9059cbb146100b0575b600080fd5b34801561006757600080fd5b506100706100f5565b60408051918252519081900360200190f35b34801561008e57600080fd5b5061007073ffffffffffffffffffffffffffffffffffffffff600435166100fb565b3480156100bc57600080fd5b506100e173ffffffffffffffffffffffffffffffffffffffff60043516602435610123565b604080519115158252519081900360200190f35b60015490565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b600073ffffffffffffffffffffffffffffffffffffffff8316151561014757600080fd5b3360009081526020819052604090205482111561016357600080fd5b33600090815260208190526040902054610183908363ffffffff61022b16565b336000908152602081905260408082209290925573ffffffffffffffffffffffffffffffffffffffff8516815220546101c2908363ffffffff61023d16565b73ffffffffffffffffffffffffffffffffffffffff8416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60008282111561023757fe5b50900390565b8181018281101561024a57fe5b929150505600a165627a7a723058203380cc6307c525bb245b03e37aa15c5d6e8f025e3bfc8d1c227bfd08aad6de020029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// BurnableTokenABI is the input ABI used to generate the binding from.
const BurnableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BurnableTokenBin is the compiled bytecode used for deploying new contracts.
const BurnableTokenBin = `0x608060405234801561001057600080fd5b50610361806100206000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461006657806342966c681461008d57806370a08231146100a7578063a9059cbb146100c8575b600080fd5b34801561007257600080fd5b5061007b610100565b60408051918252519081900360200190f35b34801561009957600080fd5b506100a5600435610106565b005b3480156100b357600080fd5b5061007b600160a060020a0360043516610113565b3480156100d457600080fd5b506100ec600160a060020a036004351660243561012e565b604080519115158252519081900360200190f35b60015490565b610110338261020f565b50565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a038316151561014557600080fd5b3360009081526020819052604090205482111561016157600080fd5b33600090815260208190526040902054610181908363ffffffff61031016565b3360009081526020819052604080822092909255600160a060020a038516815220546101b3908363ffffffff61032216565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b600160a060020a03821660009081526020819052604090205481111561023457600080fd5b600160a060020a03821660009081526020819052604090205461025d908263ffffffff61031016565b600160a060020a038316600090815260208190526040902055600154610289908263ffffffff61031016565b600155604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b60008282111561031c57fe5b50900390565b8181018281101561032f57fe5b929150505600a165627a7a723058208c240b47ec8c59dccb593c7831125d41919c48ee2b0da7340749261d0e0bf5bf0029`

// DeployBurnableToken deploys a new Ethereum contract, binding an instance of BurnableToken to it.
func DeployBurnableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// BurnableToken is an auto generated Go binding around an Ethereum contract.
type BurnableToken struct {
	BurnableTokenCaller     // Read-only binding to the contract
	BurnableTokenTransactor // Write-only binding to the contract
	BurnableTokenFilterer   // Log filterer for contract events
}

// BurnableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnableTokenSession struct {
	Contract     *BurnableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnableTokenCallerSession struct {
	Contract *BurnableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BurnableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnableTokenTransactorSession struct {
	Contract     *BurnableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BurnableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnableTokenRaw struct {
	Contract *BurnableToken // Generic contract binding to access the raw methods on
}

// BurnableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnableTokenCallerRaw struct {
	Contract *BurnableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BurnableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnableTokenTransactorRaw struct {
	Contract *BurnableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnableToken creates a new instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableToken(address common.Address, backend bind.ContractBackend) (*BurnableToken, error) {
	contract, err := bindBurnableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// NewBurnableTokenCaller creates a new read-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenCaller(address common.Address, caller bind.ContractCaller) (*BurnableTokenCaller, error) {
	contract, err := bindBurnableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenCaller{contract: contract}, nil
}

// NewBurnableTokenTransactor creates a new write-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnableTokenTransactor, error) {
	contract, err := bindBurnableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransactor{contract: contract}, nil
}

// NewBurnableTokenFilterer creates a new log filterer instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnableTokenFilterer, error) {
	contract, err := bindBurnableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenFilterer{contract: contract}, nil
}

// bindBurnableToken binds a generic wrapper to an already deployed contract.
func bindBurnableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.BurnableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BurnableToken *BurnableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BurnableToken *BurnableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// BurnableTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BurnableToken contract.
type BurnableTokenBurnIterator struct {
	Event *BurnableTokenBurn // Event containing the contract specifics and raw log

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
func (it *BurnableTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenBurn)
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
		it.Event = new(BurnableTokenBurn)
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
func (it *BurnableTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenBurn represents a Burn event raised by the BurnableToken contract.
type BurnableTokenBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*BurnableTokenBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenBurnIterator{contract: _BurnableToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BurnableTokenBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenBurn)
				if err := _BurnableToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// BurnableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnableToken contract.
type BurnableTokenTransferIterator struct {
	Event *BurnableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BurnableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenTransfer)
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
		it.Event = new(BurnableTokenTransfer)
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
func (it *BurnableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenTransfer represents a Transfer event raised by the BurnableToken contract.
type BurnableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransferIterator{contract: _BurnableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenTransfer)
				if err := _BurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// DarknodeRegistryABI is the input ABI used to generate the binding from.
const DarknodeRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isPendingRegistration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodesNextEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"},{\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"name\":\"_bond\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextMinimumBond\",\"type\":\"uint256\"}],\"name\":\"updateMinimumBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"getDarknodeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextSlasher\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isPendingDeregistration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_start\",\"type\":\"address\"},{\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getPreviousDarknodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumEpochInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumEpochInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prover\",\"type\":\"address\"},{\"name\":\"_challenger1\",\"type\":\"address\"},{\"name\":\"_challenger2\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRefundable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousEpoch\",\"outputs\":[{\"name\":\"epochhash\",\"type\":\"uint256\"},{\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextMinimumEpochInterval\",\"type\":\"uint256\"}],\"name\":\"updateMinimumEpochInterval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumPodSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodesPreviousEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"name\":\"epochhash\",\"type\":\"uint256\"},{\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRegisteredInPreviousEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isDeregistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextMinimumPodSize\",\"type\":\"uint256\"}],\"name\":\"updateMinimumPodSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"deregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"getDarknodePublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ren\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_slasher\",\"type\":\"address\"}],\"name\":\"updateSlasher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"getDarknodeBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferStoreOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumPodSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isDeregisterable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_start\",\"type\":\"address\"},{\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getDarknodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRefunded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_renAddress\",\"type\":\"address\"},{\"name\":\"_storeAddress\",\"type\":\"address\"},{\"name\":\"_minimumBond\",\"type\":\"uint256\"},{\"name\":\"_minimumPodSize\",\"type\":\"uint256\"},{\"name\":\"_minimumEpochInterval\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_darknodeID\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_bond\",\"type\":\"uint256\"}],\"name\":\"LogDarknodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"LogDarknodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogDarknodeOwnerRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogNewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousMinimumBond\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextMinimumBond\",\"type\":\"uint256\"}],\"name\":\"LogMinimumBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousMinimumPodSize\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextMinimumPodSize\",\"type\":\"uint256\"}],\"name\":\"LogMinimumPodSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousMinimumEpochInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextMinimumEpochInterval\",\"type\":\"uint256\"}],\"name\":\"LogMinimumEpochIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousSlasher\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextSlasher\",\"type\":\"address\"}],\"name\":\"LogSlasherUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DarknodeRegistryBin is the compiled bytecode used for deploying new contracts.
const DarknodeRegistryBin = `0x608060405234801561001057600080fd5b5060405160a0806125fa833981016040818152825160208085015183860151606087015160809097015160008054600160a060020a03199081163317825560118054600160a060020a039687169083161790556010805495909716941693909317909455600481905560085560058690556009959095556006829055600a919091558183019091524360001981014080845291909201829052600c55600d556001819055600281905560035561252f806100cb6000396000f3006080604052600436106101df5763ffffffff60e060020a600035041663040fa05181146101e45780630847e9fa146102195780630aeb6b40146102405780630ff9aafe146102725780631460e6031461028a5780631cedf8a31461029f57806321a2ad3a146102dc578063303ee989146102f15780634384607414610312578063455dc46d1461038657806355cacda51461039b578063563bf264146103b05780635aebd1cb146103dd5780635cdaab48146103fe57806360a22fe41461042c57806363b851b914610441578063702c25ee14610459578063715018a61461046e57806371740d161461048357806376671808146104985780637be266da146104ad5780638020fc1f146104ce57806380a0c461146104ef57806384ac33ec1461050757806384d2688c146105285780638a9b4067146105be5780638da5cb5b146105d3578063900cf0cf146105e8578063975057e7146105fd578063aa7517e114610612578063b134427114610627578063b3139d381461063c578063ba0f5b201461065d578063c2250a991461067e578063c3c5a5471461069f578063c7dbc2be146106c0578063e1878925146106d5578063ec5325c1146106f6578063f2fde38b1461071a578063fa89401a1461073b578063ffc9152e1461075c575b600080fd5b3480156101f057600080fd5b50610205600160a060020a036004351661077d565b604080519115158252519081900360200190f35b34801561022557600080fd5b5061022e61082b565b60408051918252519081900360200190f35b34801561024c57600080fd5b5061027060048035600160a060020a03169060248035908101910135604435610831565b005b34801561027e57600080fd5b50610270600435610bd0565b34801561029657600080fd5b5061022e610bec565b3480156102ab57600080fd5b506102c0600160a060020a0360043516610bf2565b60408051600160a060020a039092168252519081900360200190f35b3480156102e857600080fd5b506102c0610c7d565b3480156102fd57600080fd5b50610205600160a060020a0360043516610c8c565b34801561031e57600080fd5b50610336600160a060020a0360043516602435610cde565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561037257818101518382015260200161035a565b505050509050019250505060405180910390f35b34801561039257600080fd5b5061022e610d01565b3480156103a757600080fd5b5061022e610d07565b3480156103bc57600080fd5b50610270600160a060020a0360043581169060243581169060443516610d0d565b3480156103e957600080fd5b50610205600160a060020a036004351661119b565b34801561040a57600080fd5b50610413611236565b6040805192835260208301919091528051918290030190f35b34801561043857600080fd5b5061022e61123f565b34801561044d57600080fd5b50610270600435611245565b34801561046557600080fd5b5061022e611261565b34801561047a57600080fd5b50610270611267565b34801561048f57600080fd5b5061022e6112d3565b3480156104a457600080fd5b506104136112d9565b3480156104b957600080fd5b50610205600160a060020a03600435166112e2565b3480156104da57600080fd5b50610205600160a060020a0360043516611307565b3480156104fb57600080fd5b5061027060043561139c565b34801561051357600080fd5b50610270600160a060020a03600435166113b8565b34801561053457600080fd5b50610549600160a060020a03600435166115d0565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561058357818101518382015260200161056b565b50505050905090810190601f1680156105b05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105ca57600080fd5b506102c06116c9565b3480156105df57600080fd5b506102c06116d8565b3480156105f457600080fd5b506102706116e7565b34801561060957600080fd5b506102c061198f565b34801561061e57600080fd5b5061022e61199e565b34801561063357600080fd5b506102c06119a4565b34801561064857600080fd5b50610270600160a060020a03600435166119b3565b34801561066957600080fd5b5061022e600160a060020a03600435166119f9565b34801561068a57600080fd5b50610270600160a060020a0360043516611a64565b3480156106ab57600080fd5b50610205600160a060020a0360043516611afd565b3480156106cc57600080fd5b5061022e611b22565b3480156106e157600080fd5b50610205600160a060020a0360043516611b28565b34801561070257600080fd5b50610336600160a060020a0360043516602435611bbf565b34801561072657600080fd5b50610270600160a060020a0360043516611bda565b34801561074757600080fd5b50610270600160a060020a0360043516611bfd565b34801561076857600080fd5b50610205600160a060020a0360043516611f57565b601154604080517fe2d7c64c000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291516000938493169163e2d7c64c91602480830192602092919082900301818787803b1580156107e657600080fd5b505af11580156107fa573d6000803e3d6000fd5b505050506040513d602081101561081057600080fd5b5051905080158015906108245750600d5481115b9392505050565b60025481565b8361083b81611f57565b15156108b6576040805160e560020a62461bcd028152602060048201526024808201527f6d75737420626520726566756e646564206f72206e657665722072656769737460448201527f6572656400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600454821015610910576040805160e560020a62461bcd02815260206004820152601160248201527f696e73756666696369656e7420626f6e64000000000000000000000000000000604482015290519081900360640190fd5b601054604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b15801561098357600080fd5b505af1158015610997573d6000803e3d6000fd5b505050506040513d60208110156109ad57600080fd5b50511515610a05576040805160e560020a62461bcd02815260206004820152601360248201527f626f6e642074726173666572206661696c656400000000000000000000000000604482015290519081900360640190fd5b601054601154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015610a7757600080fd5b505af1158015610a8b573d6000803e3d6000fd5b505050506040513d6020811015610aa157600080fd5b5050601154600654600d546040517fa85ef579000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483019081523360248401819052604484018990529390940160848301819052600060a4840181905260c06064850190815260c485018b9052929096169563a85ef579958c95948a948d948d949093909290919060e401868680828437820191505098505050505050505050600060405180830381600087803b158015610b6457600080fd5b505af1158015610b78573d6000803e3d6000fd5b5050600280546001019055505060408051600160a060020a03871681526020810184905281517fd2819ba4c736158371edf0be38fd8d1fc435609832e392f118c4c79160e5bd7b929181900390910190a15050505050565b600054600160a060020a03163314610be757600080fd5b600855565b60015481565b601154604080516000805160206124e48339815191528152600160a060020a0384811660048301529151600093929092169163a30788159160248082019260209290919082900301818787803b158015610c4b57600080fd5b505af1158015610c5f573d6000803e3d6000fd5b505050506040513d6020811015610c7557600080fd5b505192915050565b600b54600160a060020a031681565b6011546040805160e160020a62c985b7028152600160a060020a0384811660048301529151600093849316916301930b6e91602480830192602092919082900301818787803b1580156107e657600080fd5b606081801515610ced57506003545b610cf984826001612082565b949350505050565b600a5481565b60065481565b6007546000908190600160a060020a03163314610d74576040805160e560020a62461bcd02815260206004820152600f60248201527f6d75737420626520736c61736865720000000000000000000000000000000000604482015290519081900360640190fd5b601154604080517fcad41357000000000000000000000000000000000000000000000000000000008152600160a060020a0388811660048301529151600293929092169163cad41357916024808201926020929091908290030181600087803b158015610de057600080fd5b505af1158015610df4573d6000803e3d6000fd5b505050506040513d6020811015610e0a57600080fd5b5051811515610e1557fe5b049150600482601154604080517ffbc402fc000000000000000000000000000000000000000000000000000000008152600160a060020a038a81166004830152602482018890529151949093049450169163fbc402fc9160448082019260009290919082900301818387803b158015610e8d57600080fd5b505af1158015610ea1573d6000803e3d6000fd5b50505050610eae85611b28565b15610f8657601154600654600d54604080517f3ac39d4b000000000000000000000000000000000000000000000000000000008152600160a060020a038a8116600483015292909301602484015251921691633ac39d4b9160448082019260009290919082900301818387803b158015610f2757600080fd5b505af1158015610f3b573d6000803e3d6000fd5b505060028054600019019055505060408051600160a060020a038716815290517f2dc89de5703d2c341a22ebfc7c4d3f197e5e1f0c19bc2e1135f387163cb927e49181900360200190a15b601054601154604080516000805160206124e48339815191528152600160a060020a03888116600483015291519382169363a9059cbb939092169163a3078815916024808201926020929091908290030181600087803b158015610fe957600080fd5b505af1158015610ffd573d6000803e3d6000fd5b505050506040513d602081101561101357600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a039092166004830152602482018590525160448083019260209291908290030181600087803b15801561106157600080fd5b505af1158015611075573d6000803e3d6000fd5b505050506040513d602081101561108b57600080fd5b5050601054601154604080516000805160206124e48339815191528152600160a060020a03878116600483015291519382169363a9059cbb939092169163a3078815916024808201926020929091908290030181600087803b1580156110f057600080fd5b505af1158015611104573d6000803e3d6000fd5b505050506040513d602081101561111a57600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a039092166004830152602482018590525160448083019260209291908290030181600087803b15801561116857600080fd5b505af115801561117c573d6000803e3d6000fd5b505050506040513d602081101561119257600080fd5b50505050505050565b60006111a682611307565b80156112305750600f546011546040805160e160020a62c985b7028152600160a060020a038681166004830152915191909216916301930b6e9160248083019260209291908290030181600087803b15801561120157600080fd5b505af1158015611215573d6000803e3d6000fd5b505050506040513d602081101561122b57600080fd5b505111155b92915050565b600e54600f5482565b60085481565b600054600160a060020a0316331461125c57600080fd5b600a55565b60095481565b600054600160a060020a0316331461127e57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60035481565b600c54600d5482565b60408051808201909152600e548152600f54602082015260009061123090839061230a565b6011546040805160e160020a62c985b7028152600160a060020a0384811660048301529151600093849316916301930b6e91602480830192602092919082900301818787803b15801561135957600080fd5b505af115801561136d573d6000803e3d6000fd5b505050506040513d602081101561138357600080fd5b5051905080158015906108245750600d54101592915050565b600054600160a060020a031633146113b357600080fd5b600955565b806113c281611b28565b1515611418576040805160e560020a62461bcd02815260206004820152601660248201527f6d757374206265206465726567697374657261626c6500000000000000000000604482015290519081900360640190fd5b601154604080516000805160206124e48339815191528152600160a060020a038086166004830152915185933393169163a30788159160248083019260209291908290030181600087803b15801561146f57600080fd5b505af1158015611483573d6000803e3d6000fd5b505050506040513d602081101561149957600080fd5b5051600160a060020a0316146114f9576040805160e560020a62461bcd02815260206004820152601660248201527f6d757374206265206461726b6e6f6465206f776e657200000000000000000000604482015290519081900360640190fd5b601154600654600d54604080517f3ac39d4b000000000000000000000000000000000000000000000000000000008152600160a060020a03888116600483015292909301602484015251921691633ac39d4b9160448082019260009290919082900301818387803b15801561156d57600080fd5b505af1158015611581573d6000803e3d6000fd5b505060028054600019019055505060408051600160a060020a038516815290517f2dc89de5703d2c341a22ebfc7c4d3f197e5e1f0c19bc2e1135f387163cb927e49181900360200190a1505050565b601154604080517fee594a50000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151606093929092169163ee594a509160248082019260009290919082900301818387803b15801561163b57600080fd5b505af115801561164f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561167857600080fd5b81019080805164010000000081111561169057600080fd5b820160208101848111156116a357600080fd5b81516401000000008111828201871017156116bd57600080fd5b50909695505050505050565b601054600160a060020a031681565b600054600160a060020a031681565b600f54600090151561175557600054600160a060020a03163314611755576040805160e560020a62461bcd02815260206004820152601d60248201527f6e6f7420617574686f7269736564202866697273742065706f63687329000000604482015290519081900360640190fd5b600654600d54014310156117b3576040805160e560020a62461bcd02815260206004820152601d60248201527f65706f636820696e74657276616c20686173206e6f7420706173736564000000604482015290519081900360640190fd5b50600c8054600e55600d8054600f5560408051808201909152436000198101408083526020909201819052928190559190556001805460035560025490556004546008541461184057600854600481905560408051828152602081019290925280517f7c77c94944e9e4e5b0d46f1297127d060020792687cd743401d782346c68f6559281900390910190a15b6005546009541461188f57600954600581905560408051828152602081019290925280517f6d520e46e5714982ddf8cb6216bcb3e1c1d5b79d337afc305335f819394f5d6a9281900390910190a15b600654600a54146118de57600a54600681905560408051828152602081019290925280517fb218cde2730b79a0667ddf869466ee66a12ef56fe65fa4986a590f8a7108c9de9281900390910190a15b600754600b54600160a060020a0390811691161461196357600b546007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039283169081179182905560408051929093168252602082015281517f933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8929181900390910190a15b6040517feff7e281fe3b4211ed1f0a5e6419bcc40f4552974f771357e66926421f0a58e890600090a150565b601154600160a060020a031681565b60045481565b600754600160a060020a031681565b600054600160a060020a031633146119ca57600080fd5b600b805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b601154604080517fcad41357000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093929092169163cad413579160248082019260209290919082900301818787803b158015610c4b57600080fd5b600054600160a060020a03163314611a7b57600080fd5b601154604080517ff2fde38b000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151919092169163f2fde38b91602480830192600092919082900301818387803b158015611ae257600080fd5b505af1158015611af6573d6000803e3d6000fd5b5050505050565b60408051808201909152600c548152600d54602082015260009061123090839061230a565b60055481565b6011546040805160e160020a62c985b7028152600160a060020a0384811660048301529151600093849316916301930b6e91602480830192602092919082900301818787803b158015611b7a57600080fd5b505af1158015611b8e573d6000803e3d6000fd5b505050506040513d6020811015611ba457600080fd5b50519050611bb183611afd565b801561082457501592915050565b606081801515611bce57506001545b610cf984826000612082565b600054600160a060020a03163314611bf157600080fd5b611bfa81612466565b50565b601154604080516000805160206124e48339815191528152600160a060020a03808516600483015291516000938593339391169163a30788159160248082019260209290919082900301818987803b158015611c5857600080fd5b505af1158015611c6c573d6000803e3d6000fd5b505050506040513d6020811015611c8257600080fd5b5051600160a060020a031614611ce2576040805160e560020a62461bcd02815260206004820152601660248201527f6d757374206265206461726b6e6f6465206f776e657200000000000000000000604482015290519081900360640190fd5b82611cec8161119b565b1515611d68576040805160e560020a62461bcd02815260206004820152602b60248201527f6d7573742062652064657265676973746572656420666f72206174206c65617360448201527f74206f6e652065706f6368000000000000000000000000000000000000000000606482015290519081900360840190fd5b601154604080517fcad41357000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301529151919092169163cad413579160248083019260209291908290030181600087803b158015611dd057600080fd5b505af1158015611de4573d6000803e3d6000fd5b505050506040513d6020811015611dfa57600080fd5b5051601154604080517f41b44392000000000000000000000000000000000000000000000000000000008152600160a060020a03888116600483015291519396509116916341b443929160248082019260009290919082900301818387803b158015611e6557600080fd5b505af1158015611e79573d6000803e3d6000fd5b5050601054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018890529051600160a060020a03909216935063a9059cbb92506044808201926020929091908290030181600087803b158015611eea57600080fd5b505af1158015611efe573d6000803e3d6000fd5b505050506040513d6020811015611f1457600080fd5b5050604080513381526020810185905281517f96ab9e56c79eee4a72db6e2879cbfbecdba5c65b411f4861824e66b89df19764929181900390910190a150505050565b601154604080517fe2d7c64c000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291516000938493849391169163e2d7c64c9160248082019260209290919082900301818787803b158015611fc457600080fd5b505af1158015611fd8573d6000803e3d6000fd5b505050506040513d6020811015611fee57600080fd5b50516011546040805160e160020a62c985b7028152600160a060020a03888116600483015291519395509116916301930b6e916024808201926020929091908290030181600087803b15801561204357600080fd5b505af1158015612057573d6000803e3d6000fd5b505050506040513d602081101561206d57600080fd5b5051905081158015610cf95750159392505050565b60608281600080808415156120975760015494505b846040519080825280602002602001820160405280156120c1578160200160208202803883390190505b50935060009250889150600160a060020a038216151561215d57601160009054906101000a9004600160a060020a0316600160a060020a0316631bce6ff36040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561212e57600080fd5b505af1158015612142573d6000803e3d6000fd5b505050506040513d602081101561215857600080fd5b505191505b848310156122fd57600160a060020a038216151561217a576122fd565b861561219057612189826112e2565b905061219c565b61219982611afd565b90505b80151561223e57601154604080517fab73e316000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301529151919092169163ab73e3169160248083019260209291908290030181600087803b15801561220b57600080fd5b505af115801561221f573d6000803e3d6000fd5b505050506040513d602081101561223557600080fd5b5051915061215d565b81848481518110151561224d57fe5b600160a060020a039283166020918202909201810191909152601154604080517fab73e31600000000000000000000000000000000000000000000000000000000815286851660048201529051919093169263ab73e3169260248083019391928290030181600087803b1580156122c357600080fd5b505af11580156122d7573d6000803e3d6000fd5b505050506040513d60208110156122ed57600080fd5b505160019390930192915061215d565b5091979650505050505050565b601154604080517fe2d7c64c000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015291516000938493849384938493929092169163e2d7c64c9160248082019260209290919082900301818787803b15801561237d57600080fd5b505af1158015612391573d6000803e3d6000fd5b505050506040513d60208110156123a757600080fd5b50516011546040805160e160020a62c985b7028152600160a060020a038b8116600483015291519397509116916301930b6e916024808201926020929091908290030181600087803b1580156123fc57600080fd5b505af1158015612410573d6000803e3d6000fd5b505050506040513d602081101561242657600080fd5b50519250831580159061243d575085602001518411155b915082158061244f5750856020015183115b905081801561245b5750805b979650505050505050565b600160a060020a038116151561247b57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a307881500000000000000000000000000000000000000000000000000000000a165627a7a7230582032b68ba423a138b646f1e71765b021d5c6df0c364ec4f66dfc966a981fbd9a170029`

// DeployDarknodeRegistry deploys a new Ethereum contract, binding an instance of DarknodeRegistry to it.
func DeployDarknodeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _renAddress common.Address, _storeAddress common.Address, _minimumBond *big.Int, _minimumPodSize *big.Int, _minimumEpochInterval *big.Int) (common.Address, *types.Transaction, *DarknodeRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DarknodeRegistryBin), backend, _renAddress, _storeAddress, _minimumBond, _minimumPodSize, _minimumEpochInterval)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DarknodeRegistry{DarknodeRegistryCaller: DarknodeRegistryCaller{contract: contract}, DarknodeRegistryTransactor: DarknodeRegistryTransactor{contract: contract}, DarknodeRegistryFilterer: DarknodeRegistryFilterer{contract: contract}}, nil
}

// DarknodeRegistry is an auto generated Go binding around an Ethereum contract.
type DarknodeRegistry struct {
	DarknodeRegistryCaller     // Read-only binding to the contract
	DarknodeRegistryTransactor // Write-only binding to the contract
	DarknodeRegistryFilterer   // Log filterer for contract events
}

// DarknodeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DarknodeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DarknodeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DarknodeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DarknodeRegistrySession struct {
	Contract     *DarknodeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DarknodeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DarknodeRegistryCallerSession struct {
	Contract *DarknodeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DarknodeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DarknodeRegistryTransactorSession struct {
	Contract     *DarknodeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DarknodeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DarknodeRegistryRaw struct {
	Contract *DarknodeRegistry // Generic contract binding to access the raw methods on
}

// DarknodeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DarknodeRegistryCallerRaw struct {
	Contract *DarknodeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DarknodeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DarknodeRegistryTransactorRaw struct {
	Contract *DarknodeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDarknodeRegistry creates a new instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistry(address common.Address, backend bind.ContractBackend) (*DarknodeRegistry, error) {
	contract, err := bindDarknodeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistry{DarknodeRegistryCaller: DarknodeRegistryCaller{contract: contract}, DarknodeRegistryTransactor: DarknodeRegistryTransactor{contract: contract}, DarknodeRegistryFilterer: DarknodeRegistryFilterer{contract: contract}}, nil
}

// NewDarknodeRegistryCaller creates a new read-only instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryCaller(address common.Address, caller bind.ContractCaller) (*DarknodeRegistryCaller, error) {
	contract, err := bindDarknodeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryCaller{contract: contract}, nil
}

// NewDarknodeRegistryTransactor creates a new write-only instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DarknodeRegistryTransactor, error) {
	contract, err := bindDarknodeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryTransactor{contract: contract}, nil
}

// NewDarknodeRegistryFilterer creates a new log filterer instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DarknodeRegistryFilterer, error) {
	contract, err := bindDarknodeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryFilterer{contract: contract}, nil
}

// bindDarknodeRegistry binds a generic wrapper to an already deployed contract.
func bindDarknodeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistry *DarknodeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistry.Contract.DarknodeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistry *DarknodeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.DarknodeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistry *DarknodeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.DarknodeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistry *DarknodeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistry *DarknodeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistry *DarknodeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.contract.Transact(opts, method, params...)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) CurrentEpoch(opts *bind.CallOpts) (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	ret := new(struct {
		Epochhash   *big.Int
		Blocknumber *big.Int
	})
	out := ret
	err := _DarknodeRegistry.contract.Call(opts, out, "currentEpoch")
	return *ret, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) CurrentEpoch() (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	return _DarknodeRegistry.Contract.CurrentEpoch(&_DarknodeRegistry.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) CurrentEpoch() (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	return _DarknodeRegistry.Contract.CurrentEpoch(&_DarknodeRegistry.CallOpts)
}

// GetDarknodeBond is a free data retrieval call binding the contract method 0xba0f5b20.
//
// Solidity: function getDarknodeBond(_darknodeID address) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodeBond(opts *bind.CallOpts, _darknodeID common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodeBond", _darknodeID)
	return *ret0, err
}

// GetDarknodeBond is a free data retrieval call binding the contract method 0xba0f5b20.
//
// Solidity: function getDarknodeBond(_darknodeID address) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodeBond(_darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistry.Contract.GetDarknodeBond(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeBond is a free data retrieval call binding the contract method 0xba0f5b20.
//
// Solidity: function getDarknodeBond(_darknodeID address) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodeBond(_darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistry.Contract.GetDarknodeBond(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x1cedf8a3.
//
// Solidity: function getDarknodeOwner(_darknodeID address) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodeOwner(opts *bind.CallOpts, _darknodeID common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodeOwner", _darknodeID)
	return *ret0, err
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x1cedf8a3.
//
// Solidity: function getDarknodeOwner(_darknodeID address) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodeOwner(_darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodeOwner(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x1cedf8a3.
//
// Solidity: function getDarknodeOwner(_darknodeID address) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodeOwner(_darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodeOwner(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodePublicKey is a free data retrieval call binding the contract method 0x84d2688c.
//
// Solidity: function getDarknodePublicKey(_darknodeID address) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodePublicKey(opts *bind.CallOpts, _darknodeID common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodePublicKey", _darknodeID)
	return *ret0, err
}

// GetDarknodePublicKey is a free data retrieval call binding the contract method 0x84d2688c.
//
// Solidity: function getDarknodePublicKey(_darknodeID address) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodePublicKey(_darknodeID common.Address) ([]byte, error) {
	return _DarknodeRegistry.Contract.GetDarknodePublicKey(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodePublicKey is a free data retrieval call binding the contract method 0x84d2688c.
//
// Solidity: function getDarknodePublicKey(_darknodeID address) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodePublicKey(_darknodeID common.Address) ([]byte, error) {
	return _DarknodeRegistry.Contract.GetDarknodePublicKey(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodes is a free data retrieval call binding the contract method 0xec5325c1.
//
// Solidity: function getDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodes(opts *bind.CallOpts, _start common.Address, _count *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodes", _start, _count)
	return *ret0, err
}

// GetDarknodes is a free data retrieval call binding the contract method 0xec5325c1.
//
// Solidity: function getDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// GetDarknodes is a free data retrieval call binding the contract method 0xec5325c1.
//
// Solidity: function getDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// GetPreviousDarknodes is a free data retrieval call binding the contract method 0x43846074.
//
// Solidity: function getPreviousDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCaller) GetPreviousDarknodes(opts *bind.CallOpts, _start common.Address, _count *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getPreviousDarknodes", _start, _count)
	return *ret0, err
}

// GetPreviousDarknodes is a free data retrieval call binding the contract method 0x43846074.
//
// Solidity: function getPreviousDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistrySession) GetPreviousDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetPreviousDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// GetPreviousDarknodes is a free data retrieval call binding the contract method 0x43846074.
//
// Solidity: function getPreviousDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetPreviousDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetPreviousDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// IsDeregisterable is a free data retrieval call binding the contract method 0xe1878925.
//
// Solidity: function isDeregisterable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsDeregisterable(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isDeregisterable", _darknodeID)
	return *ret0, err
}

// IsDeregisterable is a free data retrieval call binding the contract method 0xe1878925.
//
// Solidity: function isDeregisterable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsDeregisterable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregisterable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregisterable is a free data retrieval call binding the contract method 0xe1878925.
//
// Solidity: function isDeregisterable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsDeregisterable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregisterable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregistered is a free data retrieval call binding the contract method 0x8020fc1f.
//
// Solidity: function isDeregistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsDeregistered(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isDeregistered", _darknodeID)
	return *ret0, err
}

// IsDeregistered is a free data retrieval call binding the contract method 0x8020fc1f.
//
// Solidity: function isDeregistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsDeregistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregistered is a free data retrieval call binding the contract method 0x8020fc1f.
//
// Solidity: function isDeregistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsDeregistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsPendingDeregistration is a free data retrieval call binding the contract method 0x303ee989.
//
// Solidity: function isPendingDeregistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsPendingDeregistration(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isPendingDeregistration", _darknodeID)
	return *ret0, err
}

// IsPendingDeregistration is a free data retrieval call binding the contract method 0x303ee989.
//
// Solidity: function isPendingDeregistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsPendingDeregistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingDeregistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsPendingDeregistration is a free data retrieval call binding the contract method 0x303ee989.
//
// Solidity: function isPendingDeregistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsPendingDeregistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingDeregistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsPendingRegistration is a free data retrieval call binding the contract method 0x040fa051.
//
// Solidity: function isPendingRegistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsPendingRegistration(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isPendingRegistration", _darknodeID)
	return *ret0, err
}

// IsPendingRegistration is a free data retrieval call binding the contract method 0x040fa051.
//
// Solidity: function isPendingRegistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsPendingRegistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingRegistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsPendingRegistration is a free data retrieval call binding the contract method 0x040fa051.
//
// Solidity: function isPendingRegistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsPendingRegistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingRegistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefundable is a free data retrieval call binding the contract method 0x5aebd1cb.
//
// Solidity: function isRefundable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRefundable(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRefundable", _darknodeID)
	return *ret0, err
}

// IsRefundable is a free data retrieval call binding the contract method 0x5aebd1cb.
//
// Solidity: function isRefundable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRefundable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefundable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefundable is a free data retrieval call binding the contract method 0x5aebd1cb.
//
// Solidity: function isRefundable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRefundable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefundable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRefunded(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRefunded", _darknodeID)
	return *ret0, err
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRefunded(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefunded(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRefunded(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefunded(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRegistered(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRegistered", _darknodeID)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRegistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRegistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegisteredInPreviousEpoch is a free data retrieval call binding the contract method 0x7be266da.
//
// Solidity: function isRegisteredInPreviousEpoch(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRegisteredInPreviousEpoch(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRegisteredInPreviousEpoch", _darknodeID)
	return *ret0, err
}

// IsRegisteredInPreviousEpoch is a free data retrieval call binding the contract method 0x7be266da.
//
// Solidity: function isRegisteredInPreviousEpoch(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRegisteredInPreviousEpoch(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegisteredInPreviousEpoch(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegisteredInPreviousEpoch is a free data retrieval call binding the contract method 0x7be266da.
//
// Solidity: function isRegisteredInPreviousEpoch(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRegisteredInPreviousEpoch(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegisteredInPreviousEpoch(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumBond")
	return *ret0, err
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumBond(&_DarknodeRegistry.CallOpts)
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumBond(&_DarknodeRegistry.CallOpts)
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumEpochInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumEpochInterval")
	return *ret0, err
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumPodSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumPodSize")
	return *ret0, err
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumBond")
	return *ret0, err
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumBond(&_DarknodeRegistry.CallOpts)
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumBond(&_DarknodeRegistry.CallOpts)
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumEpochInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumEpochInterval")
	return *ret0, err
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumPodSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumPodSize")
	return *ret0, err
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextSlasher is a free data retrieval call binding the contract method 0x21a2ad3a.
//
// Solidity: function nextSlasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextSlasher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextSlasher")
	return *ret0, err
}

// NextSlasher is a free data retrieval call binding the contract method 0x21a2ad3a.
//
// Solidity: function nextSlasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) NextSlasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.NextSlasher(&_DarknodeRegistry.CallOpts)
}

// NextSlasher is a free data retrieval call binding the contract method 0x21a2ad3a.
//
// Solidity: function nextSlasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextSlasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.NextSlasher(&_DarknodeRegistry.CallOpts)
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodes")
	return *ret0, err
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodes() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodes(&_DarknodeRegistry.CallOpts)
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodes() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodes(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodesNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodesNextEpoch")
	return *ret0, err
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodesNextEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesNextEpoch(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodesNextEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesNextEpoch(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesPreviousEpoch is a free data retrieval call binding the contract method 0x71740d16.
//
// Solidity: function numDarknodesPreviousEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodesPreviousEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodesPreviousEpoch")
	return *ret0, err
}

// NumDarknodesPreviousEpoch is a free data retrieval call binding the contract method 0x71740d16.
//
// Solidity: function numDarknodesPreviousEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodesPreviousEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesPreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesPreviousEpoch is a free data retrieval call binding the contract method 0x71740d16.
//
// Solidity: function numDarknodesPreviousEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodesPreviousEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesPreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Owner() (common.Address, error) {
	return _DarknodeRegistry.Contract.Owner(&_DarknodeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Owner() (common.Address, error) {
	return _DarknodeRegistry.Contract.Owner(&_DarknodeRegistry.CallOpts)
}

// PreviousEpoch is a free data retrieval call binding the contract method 0x5cdaab48.
//
// Solidity: function previousEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) PreviousEpoch(opts *bind.CallOpts) (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	ret := new(struct {
		Epochhash   *big.Int
		Blocknumber *big.Int
	})
	out := ret
	err := _DarknodeRegistry.contract.Call(opts, out, "previousEpoch")
	return *ret, err
}

// PreviousEpoch is a free data retrieval call binding the contract method 0x5cdaab48.
//
// Solidity: function previousEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) PreviousEpoch() (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	return _DarknodeRegistry.Contract.PreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// PreviousEpoch is a free data retrieval call binding the contract method 0x5cdaab48.
//
// Solidity: function previousEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) PreviousEpoch() (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	return _DarknodeRegistry.Contract.PreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Ren(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "ren")
	return *ret0, err
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Ren() (common.Address, error) {
	return _DarknodeRegistry.Contract.Ren(&_DarknodeRegistry.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Ren() (common.Address, error) {
	return _DarknodeRegistry.Contract.Ren(&_DarknodeRegistry.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "slasher")
	return *ret0, err
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Slasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.Slasher(&_DarknodeRegistry.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Slasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.Slasher(&_DarknodeRegistry.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Store() (common.Address, error) {
	return _DarknodeRegistry.Contract.Store(&_DarknodeRegistry.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Store() (common.Address, error) {
	return _DarknodeRegistry.Contract.Store(&_DarknodeRegistry.CallOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Deregister(opts *bind.TransactOpts, _darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "deregister", _darknodeID)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Deregister(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Deregister(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Deregister(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Deregister(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Epoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "epoch")
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Epoch() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Epoch(&_DarknodeRegistry.TransactOpts)
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Epoch() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Epoch(&_DarknodeRegistry.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Refund(opts *bind.TransactOpts, _darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "refund", _darknodeID)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Refund(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Refund(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Refund(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Refund(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Register is a paid mutator transaction binding the contract method 0x0aeb6b40.
//
// Solidity: function register(_darknodeID address, _publicKey bytes, _bond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Register(opts *bind.TransactOpts, _darknodeID common.Address, _publicKey []byte, _bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "register", _darknodeID, _publicKey, _bond)
}

// Register is a paid mutator transaction binding the contract method 0x0aeb6b40.
//
// Solidity: function register(_darknodeID address, _publicKey bytes, _bond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Register(_darknodeID common.Address, _publicKey []byte, _bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Register(&_DarknodeRegistry.TransactOpts, _darknodeID, _publicKey, _bond)
}

// Register is a paid mutator transaction binding the contract method 0x0aeb6b40.
//
// Solidity: function register(_darknodeID address, _publicKey bytes, _bond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Register(_darknodeID common.Address, _publicKey []byte, _bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Register(&_DarknodeRegistry.TransactOpts, _darknodeID, _publicKey, _bond)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RenounceOwnership(&_DarknodeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RenounceOwnership(&_DarknodeRegistry.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x563bf264.
//
// Solidity: function slash(_prover address, _challenger1 address, _challenger2 address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Slash(opts *bind.TransactOpts, _prover common.Address, _challenger1 common.Address, _challenger2 common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "slash", _prover, _challenger1, _challenger2)
}

// Slash is a paid mutator transaction binding the contract method 0x563bf264.
//
// Solidity: function slash(_prover address, _challenger1 address, _challenger2 address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Slash(_prover common.Address, _challenger1 common.Address, _challenger2 common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Slash(&_DarknodeRegistry.TransactOpts, _prover, _challenger1, _challenger2)
}

// Slash is a paid mutator transaction binding the contract method 0x563bf264.
//
// Solidity: function slash(_prover address, _challenger1 address, _challenger2 address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Slash(_prover common.Address, _challenger1 common.Address, _challenger2 common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Slash(&_DarknodeRegistry.TransactOpts, _prover, _challenger1, _challenger2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// TransferStoreOwnership is a paid mutator transaction binding the contract method 0xc2250a99.
//
// Solidity: function transferStoreOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) TransferStoreOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "transferStoreOwnership", _newOwner)
}

// TransferStoreOwnership is a paid mutator transaction binding the contract method 0xc2250a99.
//
// Solidity: function transferStoreOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) TransferStoreOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferStoreOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// TransferStoreOwnership is a paid mutator transaction binding the contract method 0xc2250a99.
//
// Solidity: function transferStoreOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) TransferStoreOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferStoreOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumBond(opts *bind.TransactOpts, _nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumBond", _nextMinimumBond)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumBond(_nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumBond(&_DarknodeRegistry.TransactOpts, _nextMinimumBond)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumBond(_nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumBond(&_DarknodeRegistry.TransactOpts, _nextMinimumBond)
}

// UpdateMinimumEpochInterval is a paid mutator transaction binding the contract method 0x63b851b9.
//
// Solidity: function updateMinimumEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumEpochInterval(opts *bind.TransactOpts, _nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumEpochInterval", _nextMinimumEpochInterval)
}

// UpdateMinimumEpochInterval is a paid mutator transaction binding the contract method 0x63b851b9.
//
// Solidity: function updateMinimumEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumEpochInterval(_nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumEpochInterval(&_DarknodeRegistry.TransactOpts, _nextMinimumEpochInterval)
}

// UpdateMinimumEpochInterval is a paid mutator transaction binding the contract method 0x63b851b9.
//
// Solidity: function updateMinimumEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumEpochInterval(_nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumEpochInterval(&_DarknodeRegistry.TransactOpts, _nextMinimumEpochInterval)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumPodSize(opts *bind.TransactOpts, _nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumPodSize", _nextMinimumPodSize)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumPodSize(_nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumPodSize(&_DarknodeRegistry.TransactOpts, _nextMinimumPodSize)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumPodSize(_nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumPodSize(&_DarknodeRegistry.TransactOpts, _nextMinimumPodSize)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_slasher address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateSlasher", _slasher)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_slasher address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateSlasher(&_DarknodeRegistry.TransactOpts, _slasher)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_slasher address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateSlasher(&_DarknodeRegistry.TransactOpts, _slasher)
}

// DarknodeRegistryLogDarknodeDeregisteredIterator is returned from FilterLogDarknodeDeregistered and is used to iterate over the raw logs and unpacked data for LogDarknodeDeregistered events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeDeregisteredIterator struct {
	Event *DarknodeRegistryLogDarknodeDeregistered // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodeDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodeDeregistered)
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
		it.Event = new(DarknodeRegistryLogDarknodeDeregistered)
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
func (it *DarknodeRegistryLogDarknodeDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodeDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodeDeregistered represents a LogDarknodeDeregistered event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeDeregistered struct {
	DarknodeID common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeDeregistered is a free log retrieval operation binding the contract event 0x2dc89de5703d2c341a22ebfc7c4d3f197e5e1f0c19bc2e1135f387163cb927e4.
//
// Solidity: e LogDarknodeDeregistered(_darknodeID address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodeDeregistered(opts *bind.FilterOpts) (*DarknodeRegistryLogDarknodeDeregisteredIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodeDeregistered")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodeDeregisteredIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodeDeregistered", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeDeregistered is a free log subscription operation binding the contract event 0x2dc89de5703d2c341a22ebfc7c4d3f197e5e1f0c19bc2e1135f387163cb927e4.
//
// Solidity: e LogDarknodeDeregistered(_darknodeID address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodeDeregistered(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodeDeregistered) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodeDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodeDeregistered)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodeDeregistered", log); err != nil {
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

// DarknodeRegistryLogDarknodeOwnerRefundedIterator is returned from FilterLogDarknodeOwnerRefunded and is used to iterate over the raw logs and unpacked data for LogDarknodeOwnerRefunded events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeOwnerRefundedIterator struct {
	Event *DarknodeRegistryLogDarknodeOwnerRefunded // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodeOwnerRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodeOwnerRefunded)
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
		it.Event = new(DarknodeRegistryLogDarknodeOwnerRefunded)
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
func (it *DarknodeRegistryLogDarknodeOwnerRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodeOwnerRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodeOwnerRefunded represents a LogDarknodeOwnerRefunded event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeOwnerRefunded struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeOwnerRefunded is a free log retrieval operation binding the contract event 0x96ab9e56c79eee4a72db6e2879cbfbecdba5c65b411f4861824e66b89df19764.
//
// Solidity: e LogDarknodeOwnerRefunded(_owner address, _amount uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodeOwnerRefunded(opts *bind.FilterOpts) (*DarknodeRegistryLogDarknodeOwnerRefundedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodeOwnerRefunded")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodeOwnerRefundedIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodeOwnerRefunded", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeOwnerRefunded is a free log subscription operation binding the contract event 0x96ab9e56c79eee4a72db6e2879cbfbecdba5c65b411f4861824e66b89df19764.
//
// Solidity: e LogDarknodeOwnerRefunded(_owner address, _amount uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodeOwnerRefunded(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodeOwnerRefunded) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodeOwnerRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodeOwnerRefunded)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodeOwnerRefunded", log); err != nil {
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

// DarknodeRegistryLogDarknodeRegisteredIterator is returned from FilterLogDarknodeRegistered and is used to iterate over the raw logs and unpacked data for LogDarknodeRegistered events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeRegisteredIterator struct {
	Event *DarknodeRegistryLogDarknodeRegistered // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodeRegistered)
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
		it.Event = new(DarknodeRegistryLogDarknodeRegistered)
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
func (it *DarknodeRegistryLogDarknodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodeRegistered represents a LogDarknodeRegistered event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeRegistered struct {
	DarknodeID common.Address
	Bond       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeRegistered is a free log retrieval operation binding the contract event 0xd2819ba4c736158371edf0be38fd8d1fc435609832e392f118c4c79160e5bd7b.
//
// Solidity: e LogDarknodeRegistered(_darknodeID address, _bond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodeRegistered(opts *bind.FilterOpts) (*DarknodeRegistryLogDarknodeRegisteredIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodeRegistered")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodeRegisteredIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodeRegistered", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeRegistered is a free log subscription operation binding the contract event 0xd2819ba4c736158371edf0be38fd8d1fc435609832e392f118c4c79160e5bd7b.
//
// Solidity: e LogDarknodeRegistered(_darknodeID address, _bond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodeRegistered(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodeRegistered) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodeRegistered)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodeRegistered", log); err != nil {
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

// DarknodeRegistryLogMinimumBondUpdatedIterator is returned from FilterLogMinimumBondUpdated and is used to iterate over the raw logs and unpacked data for LogMinimumBondUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumBondUpdatedIterator struct {
	Event *DarknodeRegistryLogMinimumBondUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogMinimumBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogMinimumBondUpdated)
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
		it.Event = new(DarknodeRegistryLogMinimumBondUpdated)
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
func (it *DarknodeRegistryLogMinimumBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogMinimumBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogMinimumBondUpdated represents a LogMinimumBondUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumBondUpdated struct {
	PreviousMinimumBond *big.Int
	NextMinimumBond     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogMinimumBondUpdated is a free log retrieval operation binding the contract event 0x7c77c94944e9e4e5b0d46f1297127d060020792687cd743401d782346c68f655.
//
// Solidity: e LogMinimumBondUpdated(previousMinimumBond uint256, nextMinimumBond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogMinimumBondUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogMinimumBondUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogMinimumBondUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogMinimumBondUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogMinimumBondUpdated", logs: logs, sub: sub}, nil
}

// WatchLogMinimumBondUpdated is a free log subscription operation binding the contract event 0x7c77c94944e9e4e5b0d46f1297127d060020792687cd743401d782346c68f655.
//
// Solidity: e LogMinimumBondUpdated(previousMinimumBond uint256, nextMinimumBond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogMinimumBondUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogMinimumBondUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogMinimumBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogMinimumBondUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogMinimumBondUpdated", log); err != nil {
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

// DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator is returned from FilterLogMinimumEpochIntervalUpdated and is used to iterate over the raw logs and unpacked data for LogMinimumEpochIntervalUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator struct {
	Event *DarknodeRegistryLogMinimumEpochIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogMinimumEpochIntervalUpdated)
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
		it.Event = new(DarknodeRegistryLogMinimumEpochIntervalUpdated)
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
func (it *DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogMinimumEpochIntervalUpdated represents a LogMinimumEpochIntervalUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumEpochIntervalUpdated struct {
	PreviousMinimumEpochInterval *big.Int
	NextMinimumEpochInterval     *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterLogMinimumEpochIntervalUpdated is a free log retrieval operation binding the contract event 0xb218cde2730b79a0667ddf869466ee66a12ef56fe65fa4986a590f8a7108c9de.
//
// Solidity: e LogMinimumEpochIntervalUpdated(previousMinimumEpochInterval uint256, nextMinimumEpochInterval uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogMinimumEpochIntervalUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogMinimumEpochIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogMinimumEpochIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchLogMinimumEpochIntervalUpdated is a free log subscription operation binding the contract event 0xb218cde2730b79a0667ddf869466ee66a12ef56fe65fa4986a590f8a7108c9de.
//
// Solidity: e LogMinimumEpochIntervalUpdated(previousMinimumEpochInterval uint256, nextMinimumEpochInterval uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogMinimumEpochIntervalUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogMinimumEpochIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogMinimumEpochIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogMinimumEpochIntervalUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogMinimumEpochIntervalUpdated", log); err != nil {
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

// DarknodeRegistryLogMinimumPodSizeUpdatedIterator is returned from FilterLogMinimumPodSizeUpdated and is used to iterate over the raw logs and unpacked data for LogMinimumPodSizeUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumPodSizeUpdatedIterator struct {
	Event *DarknodeRegistryLogMinimumPodSizeUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogMinimumPodSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogMinimumPodSizeUpdated)
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
		it.Event = new(DarknodeRegistryLogMinimumPodSizeUpdated)
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
func (it *DarknodeRegistryLogMinimumPodSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogMinimumPodSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogMinimumPodSizeUpdated represents a LogMinimumPodSizeUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumPodSizeUpdated struct {
	PreviousMinimumPodSize *big.Int
	NextMinimumPodSize     *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterLogMinimumPodSizeUpdated is a free log retrieval operation binding the contract event 0x6d520e46e5714982ddf8cb6216bcb3e1c1d5b79d337afc305335f819394f5d6a.
//
// Solidity: e LogMinimumPodSizeUpdated(previousMinimumPodSize uint256, nextMinimumPodSize uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogMinimumPodSizeUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogMinimumPodSizeUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogMinimumPodSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogMinimumPodSizeUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogMinimumPodSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchLogMinimumPodSizeUpdated is a free log subscription operation binding the contract event 0x6d520e46e5714982ddf8cb6216bcb3e1c1d5b79d337afc305335f819394f5d6a.
//
// Solidity: e LogMinimumPodSizeUpdated(previousMinimumPodSize uint256, nextMinimumPodSize uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogMinimumPodSizeUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogMinimumPodSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogMinimumPodSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogMinimumPodSizeUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogMinimumPodSizeUpdated", log); err != nil {
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

// DarknodeRegistryLogNewEpochIterator is returned from FilterLogNewEpoch and is used to iterate over the raw logs and unpacked data for LogNewEpoch events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogNewEpochIterator struct {
	Event *DarknodeRegistryLogNewEpoch // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogNewEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogNewEpoch)
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
		it.Event = new(DarknodeRegistryLogNewEpoch)
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
func (it *DarknodeRegistryLogNewEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogNewEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogNewEpoch represents a LogNewEpoch event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogNewEpoch struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNewEpoch is a free log retrieval operation binding the contract event 0xeff7e281fe3b4211ed1f0a5e6419bcc40f4552974f771357e66926421f0a58e8.
//
// Solidity: e LogNewEpoch()
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogNewEpoch(opts *bind.FilterOpts) (*DarknodeRegistryLogNewEpochIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogNewEpoch")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogNewEpochIterator{contract: _DarknodeRegistry.contract, event: "LogNewEpoch", logs: logs, sub: sub}, nil
}

// WatchLogNewEpoch is a free log subscription operation binding the contract event 0xeff7e281fe3b4211ed1f0a5e6419bcc40f4552974f771357e66926421f0a58e8.
//
// Solidity: e LogNewEpoch()
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogNewEpoch(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogNewEpoch) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogNewEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogNewEpoch)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogNewEpoch", log); err != nil {
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

// DarknodeRegistryLogSlasherUpdatedIterator is returned from FilterLogSlasherUpdated and is used to iterate over the raw logs and unpacked data for LogSlasherUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogSlasherUpdatedIterator struct {
	Event *DarknodeRegistryLogSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogSlasherUpdated)
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
		it.Event = new(DarknodeRegistryLogSlasherUpdated)
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
func (it *DarknodeRegistryLogSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogSlasherUpdated represents a LogSlasherUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogSlasherUpdated struct {
	PreviousSlasher common.Address
	NextSlasher     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSlasherUpdated is a free log retrieval operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: e LogSlasherUpdated(previousSlasher address, nextSlasher address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogSlasherUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogSlasherUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogSlasherUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogSlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchLogSlasherUpdated is a free log subscription operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: e LogSlasherUpdated(previousSlasher address, nextSlasher address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogSlasherUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogSlasherUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogSlasherUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogSlasherUpdated", log); err != nil {
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

// DarknodeRegistryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipRenouncedIterator struct {
	Event *DarknodeRegistryOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryOwnershipRenounced)
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
		it.Event = new(DarknodeRegistryOwnershipRenounced)
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
func (it *DarknodeRegistryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryOwnershipRenounced represents a OwnershipRenounced event raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DarknodeRegistryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryOwnershipRenouncedIterator{contract: _DarknodeRegistry.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryOwnershipRenounced)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DarknodeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipTransferredIterator struct {
	Event *DarknodeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryOwnershipTransferred)
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
		it.Event = new(DarknodeRegistryOwnershipTransferred)
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
func (it *DarknodeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DarknodeRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryOwnershipTransferredIterator{contract: _DarknodeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryOwnershipTransferred)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DarknodeRegistryStoreABI is the input ABI used to generate the binding from.
const DarknodeRegistryStoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"}],\"name\":\"darknodeDeregisteredAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"begin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"},{\"name\":\"deregisteredAt\",\"type\":\"uint256\"}],\"name\":\"updateDarknodeDeregisteredAt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"}],\"name\":\"removeDarknode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ren\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"}],\"name\":\"darknodeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"address\"},{\"name\":\"_darknodeOwner\",\"type\":\"address\"},{\"name\":\"_bond\",\"type\":\"uint256\"},{\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"name\":\"_registeredAt\",\"type\":\"uint256\"},{\"name\":\"_deregisteredAt\",\"type\":\"uint256\"}],\"name\":\"appendDarknode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"}],\"name\":\"next\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"}],\"name\":\"darknodeBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"}],\"name\":\"darknodeRegisteredAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"}],\"name\":\"darknodePublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"darknodeID\",\"type\":\"address\"},{\"name\":\"bond\",\"type\":\"uint256\"}],\"name\":\"updateDarknodeBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ren\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DarknodeRegistryStoreBin is the compiled bytecode used for deploying new contracts.
const DarknodeRegistryStoreBin = `0x608060405234801561001057600080fd5b50604051602080610fc6833981016040525160008054600160a060020a0319908116331790915560038054600160a060020a0390931692909116919091179055610f678061005f6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301930b6e81146100df5780631bce6ff3146101125780633ac39d4b1461014357806341b4439214610169578063715018a61461018a5780638a9b40671461019f5780638da5cb5b146101b4578063a3078815146101c9578063a85ef579146101ea578063ab73e31614610229578063cad413571461024a578063e2d7c64c1461026b578063ee594a501461028c578063f2fde38b14610322578063fbc402fc14610343575b600080fd5b3480156100eb57600080fd5b50610100600160a060020a0360043516610367565b60408051918252519081900360200190f35b34801561011e57600080fd5b5061012761039e565b60408051600160a060020a039092168252519081900360200190f35b34801561014f57600080fd5b50610167600160a060020a03600435166024356103c6565b005b34801561017557600080fd5b50610167600160a060020a03600435166103fc565b34801561019657600080fd5b50610167610578565b3480156101ab57600080fd5b506101276105e4565b3480156101c057600080fd5b506101276105f3565b3480156101d557600080fd5b50610127600160a060020a0360043516610602565b3480156101f657600080fd5b50610167600160a060020a036004803582169160248035909116916044359160643590810191013560843560a435610639565b34801561023557600080fd5b50610127600160a060020a036004351661074c565b34801561025657600080fd5b50610100600160a060020a0360043516610775565b34801561027757600080fd5b50610100600160a060020a03600435166107ad565b34801561029857600080fd5b506102ad600160a060020a03600435166107e4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102e75781810151838201526020016102cf565b50505050905090810190601f1680156103145780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561032e57600080fd5b50610167600160a060020a03600435166108a9565b34801561034f57600080fd5b50610167600160a060020a03600435166024356108cc565b60008054600160a060020a0316331461037f57600080fd5b50600160a060020a031660009081526001602052604090206003015490565b60008054600160a060020a031633146103b657600080fd5b6103c06002610a0d565b90505b90565b600054600160a060020a031633146103dd57600080fd5b600160a060020a03909116600090815260016020526040902060030155565b60008054600160a060020a0316331461041457600080fd5b50600160a060020a038116600090815260016020819052604082209081018054825473ffffffffffffffffffffffffffffffffffffffff191683559083905560028201839055600382018390559161046f6004830182610e26565b505061047c600283610a2c565b60035460008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152602481018690529051919093169263a9059cbb9260448083019360209390929083900390910190829087803b1580156104f257600080fd5b505af1158015610506573d6000803e3d6000fd5b505050506040513d602081101561051c57600080fd5b50511515610574576040805160e560020a62461bcd02815260206004820152601460248201527f626f6e64207472616e73666572206661696c6564000000000000000000000000604482015290519081900360640190fd5b5050565b600054600160a060020a0316331461058f57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354600160a060020a031681565b600054600160a060020a031681565b60008054600160a060020a0316331461061a57600080fd5b50600160a060020a039081166000908152600160205260409020541690565b610641610e6a565b600054600160a060020a0316331461065857600080fd5b60a06040519081016040528088600160a060020a0316815260200187815260200184815260200183815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843750505092909352505050600160a060020a038981166000908152600160208181526040928390208551815473ffffffffffffffffffffffffffffffffffffffff191695169490941784558481015191840191909155908301516002830155606083015160038301556080830151805193945084936107339260048501920190610ea3565b50905050610742600289610b56565b5050505050505050565b60008054600160a060020a0316331461076457600080fd5b61076f600283610b69565b92915050565b60008054600160a060020a0316331461078d57600080fd5b50600160a060020a03166000908152600160208190526040909120015490565b60008054600160a060020a031633146107c557600080fd5b50600160a060020a031660009081526001602052604090206002015490565b600054606090600160a060020a031633146107fe57600080fd5b600160a060020a03821660009081526001602081815260409283902060040180548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561089d5780601f106108725761010080835404028352916020019161089d565b820191906000526020600020905b81548152906001019060200180831161088057829003601f168201915b50505050509050919050565b600054600160a060020a031633146108c057600080fd5b6108c981610bef565b50565b60008054600160a060020a031633146108e457600080fd5b50600160a060020a03821660009081526001602081905260409091200180549082905581811115610a085760035460008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015286860360248201529051919093169263a9059cbb9260448083019360209390929083900390910190829087803b15801561098657600080fd5b505af115801561099a573d6000803e3d6000fd5b505050506040513d60208110156109b057600080fd5b50511515610a08576040805160e560020a62461bcd02815260206004820152601460248201527f63616e6e6f74207472616e7366657220626f6e64000000000000000000000000604482015290519081900360640190fd5b505050565b60008080526020919091526040902060010154600160a060020a031690565b600080610a398484610c6c565b1515610a8f576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f7420696e206c697374000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0383161515610aa457610b50565b5050600160a060020a0381811660008181526020859052604080822080546001808301805461010093849004891680885286882090930180549190991673ffffffffffffffffffffffffffffffffffffffff199182168117909955888752948620805474ffffffffffffffffffffffffffffffffffffffff0019169383029390931790925594909352805474ffffffffffffffffffffffffffffffffffffffffff191690558154169055905b50505050565b61057482610b6384610c8b565b83610cad565b6000610b758383610c6c565b1515610bcb576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f7420696e206c697374000000000000000000000000000000000000000000604482015290519081900360640190fd5b50600160a060020a0390811660009081526020929092526040909120600101541690565b600160a060020a0381161515610c0457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a03166000908152602091909152604090205460ff1690565b6000808052602082905260409020546101009004600160a060020a0316919050565b6000610cb98483610c6c565b15610d0e576040805160e560020a62461bcd02815260206004820152600f60248201527f616c726561647920696e206c6973740000000000000000000000000000000000604482015290519081900360640190fd5b610d188484610c6c565b80610d2a5750600160a060020a038316155b1515610d80576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f7420696e206c697374000000000000000000000000000000000000000000604482015290519081900360640190fd5b50600160a060020a039182166000818152602094909452604080852060019081018054948616808852838820805474ffffffffffffffffffffffffffffffffffffffff001990811661010097880217825581850180549890991673ffffffffffffffffffffffffffffffffffffffff1998891681179099558354909716821790925595875291862080549094169285029290921790925591909252815460ff1916179055565b50805460018160011615610100020316600290046000825580601f10610e4c57506108c9565b601f0160209004906000526020600020908101906108c99190610f21565b60a0604051908101604052806000600160a060020a03168152602001600081526020016000815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ee457805160ff1916838001178555610f11565b82800160010185558215610f11579182015b82811115610f11578251825591602001919060010190610ef6565b50610f1d929150610f21565b5090565b6103c391905b80821115610f1d5760008155600101610f275600a165627a7a72305820dfbe37e559ca14802761a5d159d42ac426d05903b7bda11a93cdd59f5b3c0deb0029`

// DeployDarknodeRegistryStore deploys a new Ethereum contract, binding an instance of DarknodeRegistryStore to it.
func DeployDarknodeRegistryStore(auth *bind.TransactOpts, backend bind.ContractBackend, _ren common.Address) (common.Address, *types.Transaction, *DarknodeRegistryStore, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRegistryStoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DarknodeRegistryStoreBin), backend, _ren)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DarknodeRegistryStore{DarknodeRegistryStoreCaller: DarknodeRegistryStoreCaller{contract: contract}, DarknodeRegistryStoreTransactor: DarknodeRegistryStoreTransactor{contract: contract}, DarknodeRegistryStoreFilterer: DarknodeRegistryStoreFilterer{contract: contract}}, nil
}

// DarknodeRegistryStore is an auto generated Go binding around an Ethereum contract.
type DarknodeRegistryStore struct {
	DarknodeRegistryStoreCaller     // Read-only binding to the contract
	DarknodeRegistryStoreTransactor // Write-only binding to the contract
	DarknodeRegistryStoreFilterer   // Log filterer for contract events
}

// DarknodeRegistryStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type DarknodeRegistryStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DarknodeRegistryStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DarknodeRegistryStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DarknodeRegistryStoreSession struct {
	Contract     *DarknodeRegistryStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DarknodeRegistryStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DarknodeRegistryStoreCallerSession struct {
	Contract *DarknodeRegistryStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// DarknodeRegistryStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DarknodeRegistryStoreTransactorSession struct {
	Contract     *DarknodeRegistryStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// DarknodeRegistryStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type DarknodeRegistryStoreRaw struct {
	Contract *DarknodeRegistryStore // Generic contract binding to access the raw methods on
}

// DarknodeRegistryStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DarknodeRegistryStoreCallerRaw struct {
	Contract *DarknodeRegistryStoreCaller // Generic read-only contract binding to access the raw methods on
}

// DarknodeRegistryStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DarknodeRegistryStoreTransactorRaw struct {
	Contract *DarknodeRegistryStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDarknodeRegistryStore creates a new instance of DarknodeRegistryStore, bound to a specific deployed contract.
func NewDarknodeRegistryStore(address common.Address, backend bind.ContractBackend) (*DarknodeRegistryStore, error) {
	contract, err := bindDarknodeRegistryStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryStore{DarknodeRegistryStoreCaller: DarknodeRegistryStoreCaller{contract: contract}, DarknodeRegistryStoreTransactor: DarknodeRegistryStoreTransactor{contract: contract}, DarknodeRegistryStoreFilterer: DarknodeRegistryStoreFilterer{contract: contract}}, nil
}

// NewDarknodeRegistryStoreCaller creates a new read-only instance of DarknodeRegistryStore, bound to a specific deployed contract.
func NewDarknodeRegistryStoreCaller(address common.Address, caller bind.ContractCaller) (*DarknodeRegistryStoreCaller, error) {
	contract, err := bindDarknodeRegistryStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryStoreCaller{contract: contract}, nil
}

// NewDarknodeRegistryStoreTransactor creates a new write-only instance of DarknodeRegistryStore, bound to a specific deployed contract.
func NewDarknodeRegistryStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*DarknodeRegistryStoreTransactor, error) {
	contract, err := bindDarknodeRegistryStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryStoreTransactor{contract: contract}, nil
}

// NewDarknodeRegistryStoreFilterer creates a new log filterer instance of DarknodeRegistryStore, bound to a specific deployed contract.
func NewDarknodeRegistryStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*DarknodeRegistryStoreFilterer, error) {
	contract, err := bindDarknodeRegistryStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryStoreFilterer{contract: contract}, nil
}

// bindDarknodeRegistryStore binds a generic wrapper to an already deployed contract.
func bindDarknodeRegistryStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRegistryStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistryStore *DarknodeRegistryStoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistryStore.Contract.DarknodeRegistryStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistryStore *DarknodeRegistryStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.DarknodeRegistryStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistryStore *DarknodeRegistryStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.DarknodeRegistryStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistryStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.contract.Transact(opts, method, params...)
}

// Begin is a free data retrieval call binding the contract method 0x1bce6ff3.
//
// Solidity: function begin() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) Begin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "begin")
	return *ret0, err
}

// Begin is a free data retrieval call binding the contract method 0x1bce6ff3.
//
// Solidity: function begin() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) Begin() (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Begin(&_DarknodeRegistryStore.CallOpts)
}

// Begin is a free data retrieval call binding the contract method 0x1bce6ff3.
//
// Solidity: function begin() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) Begin() (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Begin(&_DarknodeRegistryStore.CallOpts)
}

// DarknodeBond is a free data retrieval call binding the contract method 0xcad41357.
//
// Solidity: function darknodeBond(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) DarknodeBond(opts *bind.CallOpts, darknodeID common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "darknodeBond", darknodeID)
	return *ret0, err
}

// DarknodeBond is a free data retrieval call binding the contract method 0xcad41357.
//
// Solidity: function darknodeBond(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) DarknodeBond(darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistryStore.Contract.DarknodeBond(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodeBond is a free data retrieval call binding the contract method 0xcad41357.
//
// Solidity: function darknodeBond(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) DarknodeBond(darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistryStore.Contract.DarknodeBond(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodeDeregisteredAt is a free data retrieval call binding the contract method 0x01930b6e.
//
// Solidity: function darknodeDeregisteredAt(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) DarknodeDeregisteredAt(opts *bind.CallOpts, darknodeID common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "darknodeDeregisteredAt", darknodeID)
	return *ret0, err
}

// DarknodeDeregisteredAt is a free data retrieval call binding the contract method 0x01930b6e.
//
// Solidity: function darknodeDeregisteredAt(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) DarknodeDeregisteredAt(darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistryStore.Contract.DarknodeDeregisteredAt(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodeDeregisteredAt is a free data retrieval call binding the contract method 0x01930b6e.
//
// Solidity: function darknodeDeregisteredAt(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) DarknodeDeregisteredAt(darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistryStore.Contract.DarknodeDeregisteredAt(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodeOwner is a free data retrieval call binding the contract method 0xa3078815.
//
// Solidity: function darknodeOwner(darknodeID address) constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) DarknodeOwner(opts *bind.CallOpts, darknodeID common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "darknodeOwner", darknodeID)
	return *ret0, err
}

// DarknodeOwner is a free data retrieval call binding the contract method 0xa3078815.
//
// Solidity: function darknodeOwner(darknodeID address) constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) DarknodeOwner(darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistryStore.Contract.DarknodeOwner(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodeOwner is a free data retrieval call binding the contract method 0xa3078815.
//
// Solidity: function darknodeOwner(darknodeID address) constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) DarknodeOwner(darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistryStore.Contract.DarknodeOwner(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodePublicKey is a free data retrieval call binding the contract method 0xee594a50.
//
// Solidity: function darknodePublicKey(darknodeID address) constant returns(bytes)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) DarknodePublicKey(opts *bind.CallOpts, darknodeID common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "darknodePublicKey", darknodeID)
	return *ret0, err
}

// DarknodePublicKey is a free data retrieval call binding the contract method 0xee594a50.
//
// Solidity: function darknodePublicKey(darknodeID address) constant returns(bytes)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) DarknodePublicKey(darknodeID common.Address) ([]byte, error) {
	return _DarknodeRegistryStore.Contract.DarknodePublicKey(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodePublicKey is a free data retrieval call binding the contract method 0xee594a50.
//
// Solidity: function darknodePublicKey(darknodeID address) constant returns(bytes)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) DarknodePublicKey(darknodeID common.Address) ([]byte, error) {
	return _DarknodeRegistryStore.Contract.DarknodePublicKey(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodeRegisteredAt is a free data retrieval call binding the contract method 0xe2d7c64c.
//
// Solidity: function darknodeRegisteredAt(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) DarknodeRegisteredAt(opts *bind.CallOpts, darknodeID common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "darknodeRegisteredAt", darknodeID)
	return *ret0, err
}

// DarknodeRegisteredAt is a free data retrieval call binding the contract method 0xe2d7c64c.
//
// Solidity: function darknodeRegisteredAt(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) DarknodeRegisteredAt(darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistryStore.Contract.DarknodeRegisteredAt(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// DarknodeRegisteredAt is a free data retrieval call binding the contract method 0xe2d7c64c.
//
// Solidity: function darknodeRegisteredAt(darknodeID address) constant returns(uint256)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) DarknodeRegisteredAt(darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistryStore.Contract.DarknodeRegisteredAt(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// Next is a free data retrieval call binding the contract method 0xab73e316.
//
// Solidity: function next(darknodeID address) constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) Next(opts *bind.CallOpts, darknodeID common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "next", darknodeID)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0xab73e316.
//
// Solidity: function next(darknodeID address) constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) Next(darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Next(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// Next is a free data retrieval call binding the contract method 0xab73e316.
//
// Solidity: function next(darknodeID address) constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) Next(darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Next(&_DarknodeRegistryStore.CallOpts, darknodeID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) Owner() (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Owner(&_DarknodeRegistryStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) Owner() (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Owner(&_DarknodeRegistryStore.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCaller) Ren(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistryStore.contract.Call(opts, out, "ren")
	return *ret0, err
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) Ren() (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Ren(&_DarknodeRegistryStore.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreCallerSession) Ren() (common.Address, error) {
	return _DarknodeRegistryStore.Contract.Ren(&_DarknodeRegistryStore.CallOpts)
}

// AppendDarknode is a paid mutator transaction binding the contract method 0xa85ef579.
//
// Solidity: function appendDarknode(_darknodeID address, _darknodeOwner address, _bond uint256, _publicKey bytes, _registeredAt uint256, _deregisteredAt uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactor) AppendDarknode(opts *bind.TransactOpts, _darknodeID common.Address, _darknodeOwner common.Address, _bond *big.Int, _publicKey []byte, _registeredAt *big.Int, _deregisteredAt *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.contract.Transact(opts, "appendDarknode", _darknodeID, _darknodeOwner, _bond, _publicKey, _registeredAt, _deregisteredAt)
}

// AppendDarknode is a paid mutator transaction binding the contract method 0xa85ef579.
//
// Solidity: function appendDarknode(_darknodeID address, _darknodeOwner address, _bond uint256, _publicKey bytes, _registeredAt uint256, _deregisteredAt uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) AppendDarknode(_darknodeID common.Address, _darknodeOwner common.Address, _bond *big.Int, _publicKey []byte, _registeredAt *big.Int, _deregisteredAt *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.AppendDarknode(&_DarknodeRegistryStore.TransactOpts, _darknodeID, _darknodeOwner, _bond, _publicKey, _registeredAt, _deregisteredAt)
}

// AppendDarknode is a paid mutator transaction binding the contract method 0xa85ef579.
//
// Solidity: function appendDarknode(_darknodeID address, _darknodeOwner address, _bond uint256, _publicKey bytes, _registeredAt uint256, _deregisteredAt uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorSession) AppendDarknode(_darknodeID common.Address, _darknodeOwner common.Address, _bond *big.Int, _publicKey []byte, _registeredAt *big.Int, _deregisteredAt *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.AppendDarknode(&_DarknodeRegistryStore.TransactOpts, _darknodeID, _darknodeOwner, _bond, _publicKey, _registeredAt, _deregisteredAt)
}

// RemoveDarknode is a paid mutator transaction binding the contract method 0x41b44392.
//
// Solidity: function removeDarknode(darknodeID address) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactor) RemoveDarknode(opts *bind.TransactOpts, darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistryStore.contract.Transact(opts, "removeDarknode", darknodeID)
}

// RemoveDarknode is a paid mutator transaction binding the contract method 0x41b44392.
//
// Solidity: function removeDarknode(darknodeID address) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) RemoveDarknode(darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.RemoveDarknode(&_DarknodeRegistryStore.TransactOpts, darknodeID)
}

// RemoveDarknode is a paid mutator transaction binding the contract method 0x41b44392.
//
// Solidity: function removeDarknode(darknodeID address) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorSession) RemoveDarknode(darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.RemoveDarknode(&_DarknodeRegistryStore.TransactOpts, darknodeID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistryStore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.RenounceOwnership(&_DarknodeRegistryStore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.RenounceOwnership(&_DarknodeRegistryStore.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistryStore.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.TransferOwnership(&_DarknodeRegistryStore.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.TransferOwnership(&_DarknodeRegistryStore.TransactOpts, _newOwner)
}

// UpdateDarknodeBond is a paid mutator transaction binding the contract method 0xfbc402fc.
//
// Solidity: function updateDarknodeBond(darknodeID address, bond uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactor) UpdateDarknodeBond(opts *bind.TransactOpts, darknodeID common.Address, bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.contract.Transact(opts, "updateDarknodeBond", darknodeID, bond)
}

// UpdateDarknodeBond is a paid mutator transaction binding the contract method 0xfbc402fc.
//
// Solidity: function updateDarknodeBond(darknodeID address, bond uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) UpdateDarknodeBond(darknodeID common.Address, bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.UpdateDarknodeBond(&_DarknodeRegistryStore.TransactOpts, darknodeID, bond)
}

// UpdateDarknodeBond is a paid mutator transaction binding the contract method 0xfbc402fc.
//
// Solidity: function updateDarknodeBond(darknodeID address, bond uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorSession) UpdateDarknodeBond(darknodeID common.Address, bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.UpdateDarknodeBond(&_DarknodeRegistryStore.TransactOpts, darknodeID, bond)
}

// UpdateDarknodeDeregisteredAt is a paid mutator transaction binding the contract method 0x3ac39d4b.
//
// Solidity: function updateDarknodeDeregisteredAt(darknodeID address, deregisteredAt uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactor) UpdateDarknodeDeregisteredAt(opts *bind.TransactOpts, darknodeID common.Address, deregisteredAt *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.contract.Transact(opts, "updateDarknodeDeregisteredAt", darknodeID, deregisteredAt)
}

// UpdateDarknodeDeregisteredAt is a paid mutator transaction binding the contract method 0x3ac39d4b.
//
// Solidity: function updateDarknodeDeregisteredAt(darknodeID address, deregisteredAt uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreSession) UpdateDarknodeDeregisteredAt(darknodeID common.Address, deregisteredAt *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.UpdateDarknodeDeregisteredAt(&_DarknodeRegistryStore.TransactOpts, darknodeID, deregisteredAt)
}

// UpdateDarknodeDeregisteredAt is a paid mutator transaction binding the contract method 0x3ac39d4b.
//
// Solidity: function updateDarknodeDeregisteredAt(darknodeID address, deregisteredAt uint256) returns()
func (_DarknodeRegistryStore *DarknodeRegistryStoreTransactorSession) UpdateDarknodeDeregisteredAt(darknodeID common.Address, deregisteredAt *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistryStore.Contract.UpdateDarknodeDeregisteredAt(&_DarknodeRegistryStore.TransactOpts, darknodeID, deregisteredAt)
}

// DarknodeRegistryStoreOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the DarknodeRegistryStore contract.
type DarknodeRegistryStoreOwnershipRenouncedIterator struct {
	Event *DarknodeRegistryStoreOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryStoreOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryStoreOwnershipRenounced)
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
		it.Event = new(DarknodeRegistryStoreOwnershipRenounced)
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
func (it *DarknodeRegistryStoreOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryStoreOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryStoreOwnershipRenounced represents a OwnershipRenounced event raised by the DarknodeRegistryStore contract.
type DarknodeRegistryStoreOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DarknodeRegistryStoreOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRegistryStore.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryStoreOwnershipRenouncedIterator{contract: _DarknodeRegistryStore.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryStoreOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRegistryStore.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryStoreOwnershipRenounced)
				if err := _DarknodeRegistryStore.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DarknodeRegistryStoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DarknodeRegistryStore contract.
type DarknodeRegistryStoreOwnershipTransferredIterator struct {
	Event *DarknodeRegistryStoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryStoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryStoreOwnershipTransferred)
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
		it.Event = new(DarknodeRegistryStoreOwnershipTransferred)
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
func (it *DarknodeRegistryStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryStoreOwnershipTransferred represents a OwnershipTransferred event raised by the DarknodeRegistryStore contract.
type DarknodeRegistryStoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DarknodeRegistryStoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistryStore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryStoreOwnershipTransferredIterator{contract: _DarknodeRegistryStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistryStore *DarknodeRegistryStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryStoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistryStore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryStoreOwnershipTransferred)
				if err := _DarknodeRegistryStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DarknodeRewardVaultABI is the input ABI used to generate the binding from.
const DarknodeRewardVaultABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"darknodeBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknode\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"darknodeRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"updateDarknodeRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETHEREUM\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknode\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_darknodeRegistry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousDarknodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"LogDarknodeRegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DarknodeRewardVaultBin is the compiled bytecode used for deploying new contracts.
const DarknodeRewardVaultBin = `0x608060405234801561001057600080fd5b50604051602080610908833981016040525160008054600160a060020a0319908116331790915560018054600160a060020a03909316929091169190911790556108a98061005f6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166370324b77811461009d578063715018a6146100d65780638340f549146100ed5780638da5cb5b1461010a5780639e45e0d01461013b578063aaff096d14610150578063f2fde38b14610171578063f7cdf47c14610192578063f940e385146101a7575b600080fd5b3480156100a957600080fd5b506100c4600160a060020a03600435811690602435166101ce565b60408051918252519081900360200190f35b3480156100e257600080fd5b506100eb6101eb565b005b6100eb600160a060020a0360043581169060243516604435610257565b34801561011657600080fd5b5061011f610485565b60408051600160a060020a039092168252519081900360200190f35b34801561014757600080fd5b5061011f610494565b34801561015c57600080fd5b506100eb600160a060020a03600435166104a3565b34801561017d57600080fd5b506100eb600160a060020a0360043516610531565b34801561019e57600080fd5b5061011f610554565b3480156101b357600080fd5b506100eb600160a060020a036004358116906024351661056c565b600260209081526000928352604080842090915290825290205481565b600054600160a060020a0316331461020257600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600160a060020a03821673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14156102d8573481146102d3576040805160e560020a62461bcd02815260206004820152601360248201527f6d69736d6174636865642074782076616c756500000000000000000000000000604482015290519081900360640190fd5b61041e565b341561032e576040805160e560020a62461bcd02815260206004820152601960248201527f756e6578706563746564206574686572207472616e7366657200000000000000604482015290519081900360640190fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018390529051600160a060020a038416916323b872dd9160648083019260209291908290030181600087803b15801561039c57600080fd5b505af11580156103b0573d6000803e3d6000fd5b505050506040513d60208110156103c657600080fd5b5051151561041e576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b600160a060020a03808416600090815260026020908152604080832093861683529290522054610454908263ffffffff6107ed16565b600160a060020a03938416600090815260026020908152604080832095909616825293909352929091209190915550565b600054600160a060020a031681565b600154600160a060020a031681565b600054600160a060020a031633146104ba57600080fd5b60015460408051600160a060020a039283168152918316602083015280517ff9f6dd5c784f63cc27c1079c73574a73485a6c2e7f7e2181c5eb2be8c693cfb79281900390910190a16001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461054857600080fd5b61055181610800565b50565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b600154604080517f1cedf8a3000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152915160009384931691631cedf8a391602480830192602092919082900301818787803b1580156105d557600080fd5b505af11580156105e9573d6000803e3d6000fd5b505050506040513d60208110156105ff57600080fd5b50519150600160a060020a0382161515610663576040805160e560020a62461bcd02815260206004820152601660248201527f696e76616c6964206461726b6e6f6465206f776e657200000000000000000000604482015290519081900360640190fd5b50600160a060020a03838116600090815260026020908152604080832093861680845293909152812080549190559073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14156106e957604051600160a060020a0383169082156108fc029083906000818181858888f193505050501580156106e3573d6000803e3d6000fd5b506107e7565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561076557600080fd5b505af1158015610779573d6000803e3d6000fd5b505050506040513d602081101561078f57600080fd5b505115156107e7576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b50505050565b818101828110156107fa57fe5b92915050565b600160a060020a038116151561081557600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582089cc1bc2395c5f047ecf2f86b9d5152c34456669a91aa48575d347a9fb5f3e3a0029`

// DeployDarknodeRewardVault deploys a new Ethereum contract, binding an instance of DarknodeRewardVault to it.
func DeployDarknodeRewardVault(auth *bind.TransactOpts, backend bind.ContractBackend, _darknodeRegistry common.Address) (common.Address, *types.Transaction, *DarknodeRewardVault, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRewardVaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DarknodeRewardVaultBin), backend, _darknodeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DarknodeRewardVault{DarknodeRewardVaultCaller: DarknodeRewardVaultCaller{contract: contract}, DarknodeRewardVaultTransactor: DarknodeRewardVaultTransactor{contract: contract}, DarknodeRewardVaultFilterer: DarknodeRewardVaultFilterer{contract: contract}}, nil
}

// DarknodeRewardVault is an auto generated Go binding around an Ethereum contract.
type DarknodeRewardVault struct {
	DarknodeRewardVaultCaller     // Read-only binding to the contract
	DarknodeRewardVaultTransactor // Write-only binding to the contract
	DarknodeRewardVaultFilterer   // Log filterer for contract events
}

// DarknodeRewardVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type DarknodeRewardVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRewardVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DarknodeRewardVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRewardVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DarknodeRewardVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRewardVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DarknodeRewardVaultSession struct {
	Contract     *DarknodeRewardVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DarknodeRewardVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DarknodeRewardVaultCallerSession struct {
	Contract *DarknodeRewardVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DarknodeRewardVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DarknodeRewardVaultTransactorSession struct {
	Contract     *DarknodeRewardVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DarknodeRewardVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type DarknodeRewardVaultRaw struct {
	Contract *DarknodeRewardVault // Generic contract binding to access the raw methods on
}

// DarknodeRewardVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DarknodeRewardVaultCallerRaw struct {
	Contract *DarknodeRewardVaultCaller // Generic read-only contract binding to access the raw methods on
}

// DarknodeRewardVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DarknodeRewardVaultTransactorRaw struct {
	Contract *DarknodeRewardVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDarknodeRewardVault creates a new instance of DarknodeRewardVault, bound to a specific deployed contract.
func NewDarknodeRewardVault(address common.Address, backend bind.ContractBackend) (*DarknodeRewardVault, error) {
	contract, err := bindDarknodeRewardVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DarknodeRewardVault{DarknodeRewardVaultCaller: DarknodeRewardVaultCaller{contract: contract}, DarknodeRewardVaultTransactor: DarknodeRewardVaultTransactor{contract: contract}, DarknodeRewardVaultFilterer: DarknodeRewardVaultFilterer{contract: contract}}, nil
}

// NewDarknodeRewardVaultCaller creates a new read-only instance of DarknodeRewardVault, bound to a specific deployed contract.
func NewDarknodeRewardVaultCaller(address common.Address, caller bind.ContractCaller) (*DarknodeRewardVaultCaller, error) {
	contract, err := bindDarknodeRewardVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRewardVaultCaller{contract: contract}, nil
}

// NewDarknodeRewardVaultTransactor creates a new write-only instance of DarknodeRewardVault, bound to a specific deployed contract.
func NewDarknodeRewardVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*DarknodeRewardVaultTransactor, error) {
	contract, err := bindDarknodeRewardVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRewardVaultTransactor{contract: contract}, nil
}

// NewDarknodeRewardVaultFilterer creates a new log filterer instance of DarknodeRewardVault, bound to a specific deployed contract.
func NewDarknodeRewardVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*DarknodeRewardVaultFilterer, error) {
	contract, err := bindDarknodeRewardVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DarknodeRewardVaultFilterer{contract: contract}, nil
}

// bindDarknodeRewardVault binds a generic wrapper to an already deployed contract.
func bindDarknodeRewardVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRewardVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRewardVault *DarknodeRewardVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRewardVault.Contract.DarknodeRewardVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRewardVault *DarknodeRewardVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.DarknodeRewardVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRewardVault *DarknodeRewardVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.DarknodeRewardVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRewardVault *DarknodeRewardVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRewardVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRewardVault *DarknodeRewardVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRewardVault *DarknodeRewardVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.contract.Transact(opts, method, params...)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultCaller) ETHEREUM(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRewardVault.contract.Call(opts, out, "ETHEREUM")
	return *ret0, err
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultSession) ETHEREUM() (common.Address, error) {
	return _DarknodeRewardVault.Contract.ETHEREUM(&_DarknodeRewardVault.CallOpts)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultCallerSession) ETHEREUM() (common.Address, error) {
	return _DarknodeRewardVault.Contract.ETHEREUM(&_DarknodeRewardVault.CallOpts)
}

// DarknodeBalances is a free data retrieval call binding the contract method 0x70324b77.
//
// Solidity: function darknodeBalances( address,  address) constant returns(uint256)
func (_DarknodeRewardVault *DarknodeRewardVaultCaller) DarknodeBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRewardVault.contract.Call(opts, out, "darknodeBalances", arg0, arg1)
	return *ret0, err
}

// DarknodeBalances is a free data retrieval call binding the contract method 0x70324b77.
//
// Solidity: function darknodeBalances( address,  address) constant returns(uint256)
func (_DarknodeRewardVault *DarknodeRewardVaultSession) DarknodeBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DarknodeRewardVault.Contract.DarknodeBalances(&_DarknodeRewardVault.CallOpts, arg0, arg1)
}

// DarknodeBalances is a free data retrieval call binding the contract method 0x70324b77.
//
// Solidity: function darknodeBalances( address,  address) constant returns(uint256)
func (_DarknodeRewardVault *DarknodeRewardVaultCallerSession) DarknodeBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DarknodeRewardVault.Contract.DarknodeBalances(&_DarknodeRewardVault.CallOpts, arg0, arg1)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultCaller) DarknodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRewardVault.contract.Call(opts, out, "darknodeRegistry")
	return *ret0, err
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultSession) DarknodeRegistry() (common.Address, error) {
	return _DarknodeRewardVault.Contract.DarknodeRegistry(&_DarknodeRewardVault.CallOpts)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultCallerSession) DarknodeRegistry() (common.Address, error) {
	return _DarknodeRewardVault.Contract.DarknodeRegistry(&_DarknodeRewardVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRewardVault.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultSession) Owner() (common.Address, error) {
	return _DarknodeRewardVault.Contract.Owner(&_DarknodeRewardVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRewardVault *DarknodeRewardVaultCallerSession) Owner() (common.Address, error) {
	return _DarknodeRewardVault.Contract.Owner(&_DarknodeRewardVault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(_darknode address, _token address, _value uint256) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactor) Deposit(opts *bind.TransactOpts, _darknode common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DarknodeRewardVault.contract.Transact(opts, "deposit", _darknode, _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(_darknode address, _token address, _value uint256) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultSession) Deposit(_darknode common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.Deposit(&_DarknodeRewardVault.TransactOpts, _darknode, _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(_darknode address, _token address, _value uint256) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactorSession) Deposit(_darknode common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.Deposit(&_DarknodeRewardVault.TransactOpts, _darknode, _token, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRewardVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRewardVault *DarknodeRewardVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.RenounceOwnership(&_DarknodeRewardVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.RenounceOwnership(&_DarknodeRewardVault.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.TransferOwnership(&_DarknodeRewardVault.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.TransferOwnership(&_DarknodeRewardVault.TransactOpts, _newOwner)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactor) UpdateDarknodeRegistry(opts *bind.TransactOpts, _newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.contract.Transact(opts, "updateDarknodeRegistry", _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.UpdateDarknodeRegistry(&_DarknodeRewardVault.TransactOpts, _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactorSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.UpdateDarknodeRegistry(&_DarknodeRewardVault.TransactOpts, _newDarknodeRegistry)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(_darknode address, _token address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactor) Withdraw(opts *bind.TransactOpts, _darknode common.Address, _token common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.contract.Transact(opts, "withdraw", _darknode, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(_darknode address, _token address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultSession) Withdraw(_darknode common.Address, _token common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.Withdraw(&_DarknodeRewardVault.TransactOpts, _darknode, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(_darknode address, _token address) returns()
func (_DarknodeRewardVault *DarknodeRewardVaultTransactorSession) Withdraw(_darknode common.Address, _token common.Address) (*types.Transaction, error) {
	return _DarknodeRewardVault.Contract.Withdraw(&_DarknodeRewardVault.TransactOpts, _darknode, _token)
}

// DarknodeRewardVaultLogDarknodeRegistryUpdatedIterator is returned from FilterLogDarknodeRegistryUpdated and is used to iterate over the raw logs and unpacked data for LogDarknodeRegistryUpdated events raised by the DarknodeRewardVault contract.
type DarknodeRewardVaultLogDarknodeRegistryUpdatedIterator struct {
	Event *DarknodeRewardVaultLogDarknodeRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRewardVaultLogDarknodeRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRewardVaultLogDarknodeRegistryUpdated)
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
		it.Event = new(DarknodeRewardVaultLogDarknodeRegistryUpdated)
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
func (it *DarknodeRewardVaultLogDarknodeRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRewardVaultLogDarknodeRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRewardVaultLogDarknodeRegistryUpdated represents a LogDarknodeRegistryUpdated event raised by the DarknodeRewardVault contract.
type DarknodeRewardVaultLogDarknodeRegistryUpdated struct {
	PreviousDarknodeRegistry common.Address
	NextDarknodeRegistry     common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeRegistryUpdated is a free log retrieval operation binding the contract event 0xf9f6dd5c784f63cc27c1079c73574a73485a6c2e7f7e2181c5eb2be8c693cfb7.
//
// Solidity: e LogDarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_DarknodeRewardVault *DarknodeRewardVaultFilterer) FilterLogDarknodeRegistryUpdated(opts *bind.FilterOpts) (*DarknodeRewardVaultLogDarknodeRegistryUpdatedIterator, error) {

	logs, sub, err := _DarknodeRewardVault.contract.FilterLogs(opts, "LogDarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRewardVaultLogDarknodeRegistryUpdatedIterator{contract: _DarknodeRewardVault.contract, event: "LogDarknodeRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeRegistryUpdated is a free log subscription operation binding the contract event 0xf9f6dd5c784f63cc27c1079c73574a73485a6c2e7f7e2181c5eb2be8c693cfb7.
//
// Solidity: e LogDarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_DarknodeRewardVault *DarknodeRewardVaultFilterer) WatchLogDarknodeRegistryUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRewardVaultLogDarknodeRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRewardVault.contract.WatchLogs(opts, "LogDarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRewardVaultLogDarknodeRegistryUpdated)
				if err := _DarknodeRewardVault.contract.UnpackLog(event, "LogDarknodeRegistryUpdated", log); err != nil {
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

// DarknodeRewardVaultOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the DarknodeRewardVault contract.
type DarknodeRewardVaultOwnershipRenouncedIterator struct {
	Event *DarknodeRewardVaultOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DarknodeRewardVaultOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRewardVaultOwnershipRenounced)
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
		it.Event = new(DarknodeRewardVaultOwnershipRenounced)
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
func (it *DarknodeRewardVaultOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRewardVaultOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRewardVaultOwnershipRenounced represents a OwnershipRenounced event raised by the DarknodeRewardVault contract.
type DarknodeRewardVaultOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRewardVault *DarknodeRewardVaultFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DarknodeRewardVaultOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRewardVault.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRewardVaultOwnershipRenouncedIterator{contract: _DarknodeRewardVault.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRewardVault *DarknodeRewardVaultFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DarknodeRewardVaultOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRewardVault.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRewardVaultOwnershipRenounced)
				if err := _DarknodeRewardVault.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DarknodeRewardVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DarknodeRewardVault contract.
type DarknodeRewardVaultOwnershipTransferredIterator struct {
	Event *DarknodeRewardVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DarknodeRewardVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRewardVaultOwnershipTransferred)
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
		it.Event = new(DarknodeRewardVaultOwnershipTransferred)
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
func (it *DarknodeRewardVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRewardVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRewardVaultOwnershipTransferred represents a OwnershipTransferred event raised by the DarknodeRewardVault contract.
type DarknodeRewardVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRewardVault *DarknodeRewardVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DarknodeRewardVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRewardVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRewardVaultOwnershipTransferredIterator{contract: _DarknodeRewardVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRewardVault *DarknodeRewardVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DarknodeRewardVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRewardVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRewardVaultOwnershipTransferred)
				if err := _DarknodeRewardVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ECRecoveryABI is the input ABI used to generate the binding from.
const ECRecoveryABI = "[]"

// ECRecoveryBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582051ce5400e9db2ec99f60568a05f33dba86f40cf2f0c53aa9337bccedec093bd10029`

// DeployECRecovery deploys a new Ethereum contract, binding an instance of ECRecovery to it.
func DeployECRecovery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
	ECRecoveryFilterer   // Log filterer for contract events
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverySession struct {
	Contract     *ECRecovery       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryCallerSession struct {
	Contract *ECRecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ECRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryTransactorSession struct {
	Contract     *ECRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryRaw struct {
	Contract *ECRecovery // Generic contract binding to access the raw methods on
}

// ECRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryCallerRaw struct {
	Contract *ECRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryTransactorRaw struct {
	Contract *ECRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecovery creates a new instance of ECRecovery, bound to a specific deployed contract.
func NewECRecovery(address common.Address, backend bind.ContractBackend) (*ECRecovery, error) {
	contract, err := bindECRecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// NewECRecoveryFilterer creates a new log filterer instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryFilterer, error) {
	contract, err := bindECRecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryFilterer{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.ECRecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// LinkedListABI is the input ABI used to generate the binding from.
const LinkedListABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LinkedListBin is the compiled bytecode used for deploying new contracts.
const LinkedListBin = `0x60ba61002f600b82828239805160001a6073146000811461001f57610021565bfe5b5030600052607381538281f300730000000000000000000000000000000000000000301460806040526004361060555763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f26be3fc8114605a575b600080fd5b60606089565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000815600a165627a7a72305820f17661e9d2530ea23166320277c90b5b8e3145ec59e553c2963721d0514e53030029`

// DeployLinkedList deploys a new Ethereum contract, binding an instance of LinkedList to it.
func DeployLinkedList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LinkedList, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkedListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LinkedListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LinkedList{LinkedListCaller: LinkedListCaller{contract: contract}, LinkedListTransactor: LinkedListTransactor{contract: contract}, LinkedListFilterer: LinkedListFilterer{contract: contract}}, nil
}

// LinkedList is an auto generated Go binding around an Ethereum contract.
type LinkedList struct {
	LinkedListCaller     // Read-only binding to the contract
	LinkedListTransactor // Write-only binding to the contract
	LinkedListFilterer   // Log filterer for contract events
}

// LinkedListCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinkedListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkedListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinkedListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkedListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinkedListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkedListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinkedListSession struct {
	Contract     *LinkedList       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LinkedListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinkedListCallerSession struct {
	Contract *LinkedListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LinkedListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinkedListTransactorSession struct {
	Contract     *LinkedListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LinkedListRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinkedListRaw struct {
	Contract *LinkedList // Generic contract binding to access the raw methods on
}

// LinkedListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinkedListCallerRaw struct {
	Contract *LinkedListCaller // Generic read-only contract binding to access the raw methods on
}

// LinkedListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinkedListTransactorRaw struct {
	Contract *LinkedListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinkedList creates a new instance of LinkedList, bound to a specific deployed contract.
func NewLinkedList(address common.Address, backend bind.ContractBackend) (*LinkedList, error) {
	contract, err := bindLinkedList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkedList{LinkedListCaller: LinkedListCaller{contract: contract}, LinkedListTransactor: LinkedListTransactor{contract: contract}, LinkedListFilterer: LinkedListFilterer{contract: contract}}, nil
}

// NewLinkedListCaller creates a new read-only instance of LinkedList, bound to a specific deployed contract.
func NewLinkedListCaller(address common.Address, caller bind.ContractCaller) (*LinkedListCaller, error) {
	contract, err := bindLinkedList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkedListCaller{contract: contract}, nil
}

// NewLinkedListTransactor creates a new write-only instance of LinkedList, bound to a specific deployed contract.
func NewLinkedListTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkedListTransactor, error) {
	contract, err := bindLinkedList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkedListTransactor{contract: contract}, nil
}

// NewLinkedListFilterer creates a new log filterer instance of LinkedList, bound to a specific deployed contract.
func NewLinkedListFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkedListFilterer, error) {
	contract, err := bindLinkedList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkedListFilterer{contract: contract}, nil
}

// bindLinkedList binds a generic wrapper to an already deployed contract.
func bindLinkedList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkedListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkedList *LinkedListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkedList.Contract.LinkedListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkedList *LinkedListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkedList.Contract.LinkedListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkedList *LinkedListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkedList.Contract.LinkedListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkedList *LinkedListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkedList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkedList *LinkedListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkedList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkedList *LinkedListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkedList.Contract.contract.Transact(opts, method, params...)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() constant returns(address)
func (_LinkedList *LinkedListCaller) NULL(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LinkedList.contract.Call(opts, out, "NULL")
	return *ret0, err
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() constant returns(address)
func (_LinkedList *LinkedListSession) NULL() (common.Address, error) {
	return _LinkedList.Contract.NULL(&_LinkedList.CallOpts)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() constant returns(address)
func (_LinkedList *LinkedListCallerSession) NULL() (common.Address, error) {
	return _LinkedList.Contract.NULL(&_LinkedList.CallOpts)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582099d2c55290e91916c6ac98bafd7658732ecb091e6e5a64012b6f00e6c13e62b50029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OrderbookABI is the input ABI used to generate the binding from.
const OrderbookABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderConfirmer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"},{\"name\":\"_matchedOrderID\",\"type\":\"bytes32\"}],\"name\":\"confirmOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ordersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"openSellOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"openBuyOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"buyOrderAtIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"orderOpeningFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ren\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOrderOpeningFee\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"darknodeRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderDepth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderBroker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"updateDarknodeRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderMatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderTrader\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"orderPriority\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"sellOrderAtIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"orderAtIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_orderOpeningFee\",\"type\":\"uint256\"},{\"name\":\"_renAddress\",\"type\":\"address\"},{\"name\":\"_darknodeRegistry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextFee\",\"type\":\"uint256\"}],\"name\":\"LogFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousDarknodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"LogDarknodeRegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OrderbookBin is the compiled bytecode used for deploying new contracts.
const OrderbookBin = `0x608060405234801561001057600080fd5b5060405160608061161383398101604090815281516020830151919092015160008054600160a060020a0319908116331790915560019390935560028054600160a060020a039384169085161790556003805492909116919092161790556115968061007d6000396000f3006080604052600436106101485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631107c3f7811461014d5780632a0401f01461018157806335daa7311461019e57806339b0d677146101c55780635060340b146101e9578063574ed6c11461020d57806360f7b5eb14610231578063715018a6146102495780637241d9f61461025e57806389895d53146102735780638a9b40671461028b5780638da5cb5b146102a05780638f72fc77146102b55780639012c4a8146103ae5780639e45e0d0146103c6578063a188fcb8146103db578063a5181661146103f3578063aab14d041461040b578063aaff096d14610447578063af3e8a4014610468578063b1a0801014610480578063b248e4e114610498578063e00c65ae146104b0578063f2fde38b146104c8578063f70ffc3a146104e9575b600080fd5b34801561015957600080fd5b50610165600435610501565b60408051600160a060020a039092168252519081900360200190f35b34801561018d57600080fd5b5061019c600435602435610522565b005b3480156101aa57600080fd5b506101b3610771565b60408051918252519081900360200190f35b3480156101d157600080fd5b5061019c60246004803582810192910135903561077b565b3480156101f557600080fd5b5061019c6024600480358281019291013590356107fd565b34801561021957600080fd5b5061019c60246004803582810192910135903561087f565b34801561023d57600080fd5b506101b3600435610a3b565b34801561025557600080fd5b5061019c610a6f565b34801561026a57600080fd5b506101b3610adb565b34801561027f57600080fd5b506101b3600435610ae1565b34801561029757600080fd5b50610165610af6565b3480156102ac57600080fd5b50610165610b05565b3480156102c157600080fd5b506102d0600435602435610b14565b60405180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015610318578181015183820152602001610300565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561035757818101518382015260200161033f565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561039657818101518382015260200161037e565b50505050905001965050505050505060405180910390f35b3480156103ba57600080fd5b5061019c600435610cbe565b3480156103d257600080fd5b50610165610d17565b3480156103e757600080fd5b506101b3600435610d26565b3480156103ff57600080fd5b50610165600435610d5e565b34801561041757600080fd5b50610423600435610d7c565b6040518082600381111561043357fe5b60ff16815260200191505060405180910390f35b34801561045357600080fd5b5061019c600160a060020a0360043516610d91565b34801561047457600080fd5b506101b3600435610e1f565b34801561048c57600080fd5b50610165600435610e34565b3480156104a457600080fd5b506101b3600435610e54565b3480156104bc57600080fd5b506101b3600435610e69565b3480156104d457600080fd5b5061019c600160a060020a0360043516610e8b565b3480156104f557600080fd5b506101b3600435610eae565b600081815260076020526040902060020154600160a060020a03165b919050565b600354604080517fc3c5a547000000000000000000000000000000000000000000000000000000008152336004820181905291519192600160a060020a03169163c3c5a547916024808201926020929091908290030181600087803b15801561058a57600080fd5b505af115801561059e573d6000803e3d6000fd5b505050506040513d60208110156105b457600080fd5b5051151561060c576040805160e560020a62461bcd02815260206004820152601b60248201527f6d7573742062652072656769737465726564206461726b6e6f64650000000000604482015290519081900360640190fd5b600160008481526007602052604090205460ff16600381111561062b57fe5b14610680576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c6964206f7264657220737461747573000000000000000000000000604482015290519081900360640190fd5b600160008381526007602052604090205460ff16600381111561069f57fe5b146106f4576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c6964206f7264657220737461747573000000000000000000000000604482015290519081900360640190fd5b50600082815260076020526040808220805460ff199081166002908117835580830180543373ffffffffffffffffffffffffffffffffffffffff1991821681179092556005808601899055436004968701819055988852959096208054909316821783559082018054909516179093559082019390935590910155565b6005546004540190565b6107b683838080601f01602080910402602001604051908101604052809392919081815260200183838082843750879450610ed09350505050565b600580546001810182557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00182905554600091825260076020526040909120600301555050565b61083883838080601f01602080910402602001604051908101604052809392919081815260200183838082843750879450610ed09350505050565b600480546001810182557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0182905554600091825260076020526040909120600301555050565b60606000600160008481526007602052604090205460ff1660038111156108a257fe5b14156109a357604080517f52657075626c69632050726f746f636f6c3a2063616e63656c3a200000000000602080830191909152603b80830187905283518084039091018152601f88018290049091028201607b908101909352605b82018781529094506109259285928991899182910183828082843750611131945050505050565b600084815260076020526040902054909150600160a060020a03808316610100909204161461099e576040805160e560020a62461bcd02815260206004820152601160248201527f696e76616c6964207369676e6174757265000000000000000000000000000000604482015290519081900360640190fd5b610a15565b60008381526007602052604081205460ff1660038111156109c057fe5b14610a15576040805160e560020a62461bcd02815260206004820152601360248201527f696e76616c6964206f7264657220737461746500000000000000000000000000604482015290519081900360640190fd5b50506000908152600760205260409020805460ff19166003178155436004909101555050565b6004546000908210610a4f5750600061051d565b6004805483908110610a5d57fe5b90600052602060002001549050919050565b600054600160a060020a03163314610a8657600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60015481565b60009081526007602052604090206004015490565b600254600160a060020a031681565b600054600160a060020a031681565b6060806060600060608060606000806006805490508b101515610b3657610cb1565b6006548a96508b87011115610b4e576006548b900395505b85604051908082528060200260200182016040528015610b78578160200160208202803883390190505b50945085604051908082528060200260200182016040528015610ba5578160200160208202803883390190505b50935085604051908082528060200260200182016040528015610bd2578160200160208202803883390190505b509250600091505b85821015610ca75760068054838d01908110610bf257fe5b90600052602060002001549050808583815181101515610c0e57fe5b6020908102919091018101919091526000828152600790915260409020548451610100909104600160a060020a031690859084908110610c4a57fe5b600160a060020a03909216602092830290910182015260008281526007909152604090205460ff166003811115610c7d57fe5b8383815181101515610c8b57fe5b60ff909216602092830290910190910152600190910190610bda565b8484849850985098505b5050505050509250925092565b600054600160a060020a03163314610cd557600080fd5b600154604080519182526020820183905280517f6308fd23d4b466bc07d9c9ef31631f5b96c7bac2efcb3d214fe3cc678e7ae00a9281900390910190a1600155565b600354600160a060020a031681565b6000818152600760205260408120600401541515610d465750600061051d565b50600090815260076020526040902060040154430390565b600090815260076020526040902060010154600160a060020a031690565b60009081526007602052604090205460ff1690565b600054600160a060020a03163314610da857600080fd5b60035460408051600160a060020a039283168152918316602083015280517ff9f6dd5c784f63cc27c1079c73574a73485a6c2e7f7e2181c5eb2be8c693cfb79281900390910190a16003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60009081526007602052604090206005015490565b6000908152600760205260409020546101009004600160a060020a031690565b60009081526007602052604090206003015490565b6005546000908210610e7d5750600061051d565b6005805483908110610a5d57fe5b600054600160a060020a03163314610ea257600080fd5b610eab816112eb565b50565b6006546000908210610ec25750600061051d565b6006805483908110610a5d57fe5b600254600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019290925251606092600092600160a060020a03909116916323b872dd9160648082019260209290919082900301818787803b158015610f4b57600080fd5b505af1158015610f5f573d6000803e3d6000fd5b505050506040513d6020811015610f7557600080fd5b50511515610fcd576040805160e560020a62461bcd02815260206004820152601360248201527f666565207472616e73666572206661696c656400000000000000000000000000604482015290519081900360640190fd5b60008381526007602052604081205460ff166003811115610fea57fe5b1461103f576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c6964206f7264657220737461747573000000000000000000000000604482015290519081900360640190fd5b604080517f52657075626c69632050726f746f636f6c3a206f70656e3a200000000000000060208201526039808201869052825180830390910181526059909101909152915061108f8285611131565b60008481526007602052604081208054600160a060020a039390931661010002600160ff19909416841774ffffffffffffffffffffffffffffffffffffffff001916178155828101805473ffffffffffffffffffffffffffffffffffffffff1916331790554360049091015560068054928301815590527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0192909255505050565b600060608060006040805190810160405280601a81526020017f19457468657265756d205369676e6564204d6573736167653a0a00000000000081525092508261117b8751611368565b876040516020018084805190602001908083835b602083106111ae5780518252601f19909201916020918201910161118f565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106111f65780518252601f1990920191602091820191016111d7565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061123e5780518252601f19909201916020918201910161121f565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040529150816040518082805190602001908083835b602083106112a65780518252601f199092019160209182019101611287565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506112df818661149a565b93505b50505092915050565b600160a060020a038116151561130057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60608160008083818415156113b25760408051808201909152600181527f300000000000000000000000000000000000000000000000000000000000000060208201529550611490565b600093508492505b60008311156113d457600a830492506001840193506113ba565b836040519080825280601f01601f191660200182016040528015611402578160200160208202803883390190505b509150600090505b8381101561148c57600a85066030017f0100000000000000000000000000000000000000000000000000000000000000028260018387030381518110151561144e57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a8504945060010161140a565b8195505b5050505050919050565b600080600080845160411415156114b457600093506112e2565b50505060208201516040830151606084015160001a601b60ff821610156114d957601b015b8060ff16601b141580156114f157508060ff16601c14155b156114ff57600093506112e2565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015611559573d6000803e3d6000fd5b5050506020604051035193506112e25600a165627a7a72305820c1e21657f58db2d0eb2271c200f0c2b695d7d560ab3d59a367396fc5f1e512030029`

// DeployOrderbook deploys a new Ethereum contract, binding an instance of Orderbook to it.
func DeployOrderbook(auth *bind.TransactOpts, backend bind.ContractBackend, _orderOpeningFee *big.Int, _renAddress common.Address, _darknodeRegistry common.Address) (common.Address, *types.Transaction, *Orderbook, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderbookBin), backend, _orderOpeningFee, _renAddress, _darknodeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// Orderbook is an auto generated Go binding around an Ethereum contract.
type Orderbook struct {
	OrderbookCaller     // Read-only binding to the contract
	OrderbookTransactor // Write-only binding to the contract
	OrderbookFilterer   // Log filterer for contract events
}

// OrderbookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderbookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderbookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderbookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderbookSession struct {
	Contract     *Orderbook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderbookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderbookCallerSession struct {
	Contract *OrderbookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderbookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderbookTransactorSession struct {
	Contract     *OrderbookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderbookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderbookRaw struct {
	Contract *Orderbook // Generic contract binding to access the raw methods on
}

// OrderbookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderbookCallerRaw struct {
	Contract *OrderbookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderbookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderbookTransactorRaw struct {
	Contract *OrderbookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderbook creates a new instance of Orderbook, bound to a specific deployed contract.
func NewOrderbook(address common.Address, backend bind.ContractBackend) (*Orderbook, error) {
	contract, err := bindOrderbook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// NewOrderbookCaller creates a new read-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookCaller(address common.Address, caller bind.ContractCaller) (*OrderbookCaller, error) {
	contract, err := bindOrderbook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookCaller{contract: contract}, nil
}

// NewOrderbookTransactor creates a new write-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderbookTransactor, error) {
	contract, err := bindOrderbook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookTransactor{contract: contract}, nil
}

// NewOrderbookFilterer creates a new log filterer instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderbookFilterer, error) {
	contract, err := bindOrderbook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderbookFilterer{contract: contract}, nil
}

// bindOrderbook binds a generic wrapper to an already deployed contract.
func bindOrderbook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.OrderbookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transact(opts, method, params...)
}

// BuyOrderAtIndex is a free data retrieval call binding the contract method 0x60f7b5eb.
//
// Solidity: function buyOrderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) BuyOrderAtIndex(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "buyOrderAtIndex", _index)
	return *ret0, err
}

// BuyOrderAtIndex is a free data retrieval call binding the contract method 0x60f7b5eb.
//
// Solidity: function buyOrderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookSession) BuyOrderAtIndex(_index *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.BuyOrderAtIndex(&_Orderbook.CallOpts, _index)
}

// BuyOrderAtIndex is a free data retrieval call binding the contract method 0x60f7b5eb.
//
// Solidity: function buyOrderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) BuyOrderAtIndex(_index *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.BuyOrderAtIndex(&_Orderbook.CallOpts, _index)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_Orderbook *OrderbookCaller) DarknodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "darknodeRegistry")
	return *ret0, err
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_Orderbook *OrderbookSession) DarknodeRegistry() (common.Address, error) {
	return _Orderbook.Contract.DarknodeRegistry(&_Orderbook.CallOpts)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_Orderbook *OrderbookCallerSession) DarknodeRegistry() (common.Address, error) {
	return _Orderbook.Contract.DarknodeRegistry(&_Orderbook.CallOpts)
}

// GetOrders is a free data retrieval call binding the contract method 0x8f72fc77.
//
// Solidity: function getOrders(_offset uint256, _limit uint256) constant returns(bytes32[], address[], uint8[])
func (_Orderbook *OrderbookCaller) GetOrders(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([][32]byte, []common.Address, []uint8, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([]common.Address)
		ret2 = new([]uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Orderbook.contract.Call(opts, out, "getOrders", _offset, _limit)
	return *ret0, *ret1, *ret2, err
}

// GetOrders is a free data retrieval call binding the contract method 0x8f72fc77.
//
// Solidity: function getOrders(_offset uint256, _limit uint256) constant returns(bytes32[], address[], uint8[])
func (_Orderbook *OrderbookSession) GetOrders(_offset *big.Int, _limit *big.Int) ([][32]byte, []common.Address, []uint8, error) {
	return _Orderbook.Contract.GetOrders(&_Orderbook.CallOpts, _offset, _limit)
}

// GetOrders is a free data retrieval call binding the contract method 0x8f72fc77.
//
// Solidity: function getOrders(_offset uint256, _limit uint256) constant returns(bytes32[], address[], uint8[])
func (_Orderbook *OrderbookCallerSession) GetOrders(_offset *big.Int, _limit *big.Int) ([][32]byte, []common.Address, []uint8, error) {
	return _Orderbook.Contract.GetOrders(&_Orderbook.CallOpts, _offset, _limit)
}

// OrderAtIndex is a free data retrieval call binding the contract method 0xf70ffc3a.
//
// Solidity: function orderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) OrderAtIndex(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderAtIndex", _index)
	return *ret0, err
}

// OrderAtIndex is a free data retrieval call binding the contract method 0xf70ffc3a.
//
// Solidity: function orderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookSession) OrderAtIndex(_index *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.OrderAtIndex(&_Orderbook.CallOpts, _index)
}

// OrderAtIndex is a free data retrieval call binding the contract method 0xf70ffc3a.
//
// Solidity: function orderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) OrderAtIndex(_index *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.OrderAtIndex(&_Orderbook.CallOpts, _index)
}

// OrderBlockNumber is a free data retrieval call binding the contract method 0x89895d53.
//
// Solidity: function orderBlockNumber(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrderBlockNumber(opts *bind.CallOpts, _orderID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderBlockNumber", _orderID)
	return *ret0, err
}

// OrderBlockNumber is a free data retrieval call binding the contract method 0x89895d53.
//
// Solidity: function orderBlockNumber(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookSession) OrderBlockNumber(_orderID [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderBlockNumber(&_Orderbook.CallOpts, _orderID)
}

// OrderBlockNumber is a free data retrieval call binding the contract method 0x89895d53.
//
// Solidity: function orderBlockNumber(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrderBlockNumber(_orderID [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderBlockNumber(&_Orderbook.CallOpts, _orderID)
}

// OrderBroker is a free data retrieval call binding the contract method 0xa5181661.
//
// Solidity: function orderBroker(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderBroker(opts *bind.CallOpts, _orderID [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderBroker", _orderID)
	return *ret0, err
}

// OrderBroker is a free data retrieval call binding the contract method 0xa5181661.
//
// Solidity: function orderBroker(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderBroker(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderBroker(&_Orderbook.CallOpts, _orderID)
}

// OrderBroker is a free data retrieval call binding the contract method 0xa5181661.
//
// Solidity: function orderBroker(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderBroker(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderBroker(&_Orderbook.CallOpts, _orderID)
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderConfirmer(opts *bind.CallOpts, _orderID [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderConfirmer", _orderID)
	return *ret0, err
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderConfirmer(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderConfirmer(&_Orderbook.CallOpts, _orderID)
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderConfirmer(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderConfirmer(&_Orderbook.CallOpts, _orderID)
}

// OrderDepth is a free data retrieval call binding the contract method 0xa188fcb8.
//
// Solidity: function orderDepth(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrderDepth(opts *bind.CallOpts, _orderID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderDepth", _orderID)
	return *ret0, err
}

// OrderDepth is a free data retrieval call binding the contract method 0xa188fcb8.
//
// Solidity: function orderDepth(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookSession) OrderDepth(_orderID [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderDepth(&_Orderbook.CallOpts, _orderID)
}

// OrderDepth is a free data retrieval call binding the contract method 0xa188fcb8.
//
// Solidity: function orderDepth(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrderDepth(_orderID [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderDepth(&_Orderbook.CallOpts, _orderID)
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderID bytes32) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) OrderMatch(opts *bind.CallOpts, _orderID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderMatch", _orderID)
	return *ret0, err
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderID bytes32) constant returns(bytes32)
func (_Orderbook *OrderbookSession) OrderMatch(_orderID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.OrderMatch(&_Orderbook.CallOpts, _orderID)
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderID bytes32) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) OrderMatch(_orderID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.OrderMatch(&_Orderbook.CallOpts, _orderID)
}

// OrderOpeningFee is a free data retrieval call binding the contract method 0x7241d9f6.
//
// Solidity: function orderOpeningFee() constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrderOpeningFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderOpeningFee")
	return *ret0, err
}

// OrderOpeningFee is a free data retrieval call binding the contract method 0x7241d9f6.
//
// Solidity: function orderOpeningFee() constant returns(uint256)
func (_Orderbook *OrderbookSession) OrderOpeningFee() (*big.Int, error) {
	return _Orderbook.Contract.OrderOpeningFee(&_Orderbook.CallOpts)
}

// OrderOpeningFee is a free data retrieval call binding the contract method 0x7241d9f6.
//
// Solidity: function orderOpeningFee() constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrderOpeningFee() (*big.Int, error) {
	return _Orderbook.Contract.OrderOpeningFee(&_Orderbook.CallOpts)
}

// OrderPriority is a free data retrieval call binding the contract method 0xb248e4e1.
//
// Solidity: function orderPriority(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrderPriority(opts *bind.CallOpts, _orderID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderPriority", _orderID)
	return *ret0, err
}

// OrderPriority is a free data retrieval call binding the contract method 0xb248e4e1.
//
// Solidity: function orderPriority(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookSession) OrderPriority(_orderID [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderPriority(&_Orderbook.CallOpts, _orderID)
}

// OrderPriority is a free data retrieval call binding the contract method 0xb248e4e1.
//
// Solidity: function orderPriority(_orderID bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrderPriority(_orderID [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderPriority(&_Orderbook.CallOpts, _orderID)
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderID bytes32) constant returns(uint8)
func (_Orderbook *OrderbookCaller) OrderState(opts *bind.CallOpts, _orderID [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderState", _orderID)
	return *ret0, err
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderID bytes32) constant returns(uint8)
func (_Orderbook *OrderbookSession) OrderState(_orderID [32]byte) (uint8, error) {
	return _Orderbook.Contract.OrderState(&_Orderbook.CallOpts, _orderID)
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderID bytes32) constant returns(uint8)
func (_Orderbook *OrderbookCallerSession) OrderState(_orderID [32]byte) (uint8, error) {
	return _Orderbook.Contract.OrderState(&_Orderbook.CallOpts, _orderID)
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderTrader(opts *bind.CallOpts, _orderID [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderTrader", _orderID)
	return *ret0, err
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderTrader(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderTrader(&_Orderbook.CallOpts, _orderID)
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderID bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderTrader(_orderID [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderTrader(&_Orderbook.CallOpts, _orderID)
}

// OrdersCount is a free data retrieval call binding the contract method 0x35daa731.
//
// Solidity: function ordersCount() constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrdersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "ordersCount")
	return *ret0, err
}

// OrdersCount is a free data retrieval call binding the contract method 0x35daa731.
//
// Solidity: function ordersCount() constant returns(uint256)
func (_Orderbook *OrderbookSession) OrdersCount() (*big.Int, error) {
	return _Orderbook.Contract.OrdersCount(&_Orderbook.CallOpts)
}

// OrdersCount is a free data retrieval call binding the contract method 0x35daa731.
//
// Solidity: function ordersCount() constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrdersCount() (*big.Int, error) {
	return _Orderbook.Contract.OrdersCount(&_Orderbook.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Orderbook *OrderbookCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Orderbook *OrderbookSession) Owner() (common.Address, error) {
	return _Orderbook.Contract.Owner(&_Orderbook.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Orderbook *OrderbookCallerSession) Owner() (common.Address, error) {
	return _Orderbook.Contract.Owner(&_Orderbook.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_Orderbook *OrderbookCaller) Ren(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "ren")
	return *ret0, err
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_Orderbook *OrderbookSession) Ren() (common.Address, error) {
	return _Orderbook.Contract.Ren(&_Orderbook.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_Orderbook *OrderbookCallerSession) Ren() (common.Address, error) {
	return _Orderbook.Contract.Ren(&_Orderbook.CallOpts)
}

// SellOrderAtIndex is a free data retrieval call binding the contract method 0xe00c65ae.
//
// Solidity: function sellOrderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) SellOrderAtIndex(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "sellOrderAtIndex", _index)
	return *ret0, err
}

// SellOrderAtIndex is a free data retrieval call binding the contract method 0xe00c65ae.
//
// Solidity: function sellOrderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookSession) SellOrderAtIndex(_index *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.SellOrderAtIndex(&_Orderbook.CallOpts, _index)
}

// SellOrderAtIndex is a free data retrieval call binding the contract method 0xe00c65ae.
//
// Solidity: function sellOrderAtIndex(_index uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) SellOrderAtIndex(_index *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.SellOrderAtIndex(&_Orderbook.CallOpts, _index)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x574ed6c1.
//
// Solidity: function cancelOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookTransactor) CancelOrder(opts *bind.TransactOpts, _signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "cancelOrder", _signature, _orderID)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x574ed6c1.
//
// Solidity: function cancelOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookSession) CancelOrder(_signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.CancelOrder(&_Orderbook.TransactOpts, _signature, _orderID)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x574ed6c1.
//
// Solidity: function cancelOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookTransactorSession) CancelOrder(_signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.CancelOrder(&_Orderbook.TransactOpts, _signature, _orderID)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x2a0401f0.
//
// Solidity: function confirmOrder(_orderID bytes32, _matchedOrderID bytes32) returns()
func (_Orderbook *OrderbookTransactor) ConfirmOrder(opts *bind.TransactOpts, _orderID [32]byte, _matchedOrderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "confirmOrder", _orderID, _matchedOrderID)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x2a0401f0.
//
// Solidity: function confirmOrder(_orderID bytes32, _matchedOrderID bytes32) returns()
func (_Orderbook *OrderbookSession) ConfirmOrder(_orderID [32]byte, _matchedOrderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.ConfirmOrder(&_Orderbook.TransactOpts, _orderID, _matchedOrderID)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x2a0401f0.
//
// Solidity: function confirmOrder(_orderID bytes32, _matchedOrderID bytes32) returns()
func (_Orderbook *OrderbookTransactorSession) ConfirmOrder(_orderID [32]byte, _matchedOrderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.ConfirmOrder(&_Orderbook.TransactOpts, _orderID, _matchedOrderID)
}

// OpenBuyOrder is a paid mutator transaction binding the contract method 0x5060340b.
//
// Solidity: function openBuyOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookTransactor) OpenBuyOrder(opts *bind.TransactOpts, _signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "openBuyOrder", _signature, _orderID)
}

// OpenBuyOrder is a paid mutator transaction binding the contract method 0x5060340b.
//
// Solidity: function openBuyOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookSession) OpenBuyOrder(_signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenBuyOrder(&_Orderbook.TransactOpts, _signature, _orderID)
}

// OpenBuyOrder is a paid mutator transaction binding the contract method 0x5060340b.
//
// Solidity: function openBuyOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookTransactorSession) OpenBuyOrder(_signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenBuyOrder(&_Orderbook.TransactOpts, _signature, _orderID)
}

// OpenSellOrder is a paid mutator transaction binding the contract method 0x39b0d677.
//
// Solidity: function openSellOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookTransactor) OpenSellOrder(opts *bind.TransactOpts, _signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "openSellOrder", _signature, _orderID)
}

// OpenSellOrder is a paid mutator transaction binding the contract method 0x39b0d677.
//
// Solidity: function openSellOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookSession) OpenSellOrder(_signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenSellOrder(&_Orderbook.TransactOpts, _signature, _orderID)
}

// OpenSellOrder is a paid mutator transaction binding the contract method 0x39b0d677.
//
// Solidity: function openSellOrder(_signature bytes, _orderID bytes32) returns()
func (_Orderbook *OrderbookTransactorSession) OpenSellOrder(_signature []byte, _orderID [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenSellOrder(&_Orderbook.TransactOpts, _signature, _orderID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderbook *OrderbookTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderbook *OrderbookSession) RenounceOwnership() (*types.Transaction, error) {
	return _Orderbook.Contract.RenounceOwnership(&_Orderbook.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderbook *OrderbookTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Orderbook.Contract.RenounceOwnership(&_Orderbook.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Orderbook *OrderbookTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Orderbook *OrderbookSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.TransferOwnership(&_Orderbook.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Orderbook *OrderbookTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.TransferOwnership(&_Orderbook.TransactOpts, _newOwner)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_Orderbook *OrderbookTransactor) UpdateDarknodeRegistry(opts *bind.TransactOpts, _newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "updateDarknodeRegistry", _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_Orderbook *OrderbookSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateDarknodeRegistry(&_Orderbook.TransactOpts, _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_Orderbook *OrderbookTransactorSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateDarknodeRegistry(&_Orderbook.TransactOpts, _newDarknodeRegistry)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newOrderOpeningFee uint256) returns()
func (_Orderbook *OrderbookTransactor) UpdateFee(opts *bind.TransactOpts, _newOrderOpeningFee *big.Int) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "updateFee", _newOrderOpeningFee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newOrderOpeningFee uint256) returns()
func (_Orderbook *OrderbookSession) UpdateFee(_newOrderOpeningFee *big.Int) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateFee(&_Orderbook.TransactOpts, _newOrderOpeningFee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newOrderOpeningFee uint256) returns()
func (_Orderbook *OrderbookTransactorSession) UpdateFee(_newOrderOpeningFee *big.Int) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateFee(&_Orderbook.TransactOpts, _newOrderOpeningFee)
}

// OrderbookLogDarknodeRegistryUpdatedIterator is returned from FilterLogDarknodeRegistryUpdated and is used to iterate over the raw logs and unpacked data for LogDarknodeRegistryUpdated events raised by the Orderbook contract.
type OrderbookLogDarknodeRegistryUpdatedIterator struct {
	Event *OrderbookLogDarknodeRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *OrderbookLogDarknodeRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookLogDarknodeRegistryUpdated)
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
		it.Event = new(OrderbookLogDarknodeRegistryUpdated)
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
func (it *OrderbookLogDarknodeRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookLogDarknodeRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookLogDarknodeRegistryUpdated represents a LogDarknodeRegistryUpdated event raised by the Orderbook contract.
type OrderbookLogDarknodeRegistryUpdated struct {
	PreviousDarknodeRegistry common.Address
	NextDarknodeRegistry     common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeRegistryUpdated is a free log retrieval operation binding the contract event 0xf9f6dd5c784f63cc27c1079c73574a73485a6c2e7f7e2181c5eb2be8c693cfb7.
//
// Solidity: e LogDarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_Orderbook *OrderbookFilterer) FilterLogDarknodeRegistryUpdated(opts *bind.FilterOpts) (*OrderbookLogDarknodeRegistryUpdatedIterator, error) {

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "LogDarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &OrderbookLogDarknodeRegistryUpdatedIterator{contract: _Orderbook.contract, event: "LogDarknodeRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeRegistryUpdated is a free log subscription operation binding the contract event 0xf9f6dd5c784f63cc27c1079c73574a73485a6c2e7f7e2181c5eb2be8c693cfb7.
//
// Solidity: e LogDarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_Orderbook *OrderbookFilterer) WatchLogDarknodeRegistryUpdated(opts *bind.WatchOpts, sink chan<- *OrderbookLogDarknodeRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "LogDarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookLogDarknodeRegistryUpdated)
				if err := _Orderbook.contract.UnpackLog(event, "LogDarknodeRegistryUpdated", log); err != nil {
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

// OrderbookLogFeeUpdatedIterator is returned from FilterLogFeeUpdated and is used to iterate over the raw logs and unpacked data for LogFeeUpdated events raised by the Orderbook contract.
type OrderbookLogFeeUpdatedIterator struct {
	Event *OrderbookLogFeeUpdated // Event containing the contract specifics and raw log

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
func (it *OrderbookLogFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookLogFeeUpdated)
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
		it.Event = new(OrderbookLogFeeUpdated)
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
func (it *OrderbookLogFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookLogFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookLogFeeUpdated represents a LogFeeUpdated event raised by the Orderbook contract.
type OrderbookLogFeeUpdated struct {
	PreviousFee *big.Int
	NextFee     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogFeeUpdated is a free log retrieval operation binding the contract event 0x6308fd23d4b466bc07d9c9ef31631f5b96c7bac2efcb3d214fe3cc678e7ae00a.
//
// Solidity: e LogFeeUpdated(previousFee uint256, nextFee uint256)
func (_Orderbook *OrderbookFilterer) FilterLogFeeUpdated(opts *bind.FilterOpts) (*OrderbookLogFeeUpdatedIterator, error) {

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "LogFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &OrderbookLogFeeUpdatedIterator{contract: _Orderbook.contract, event: "LogFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchLogFeeUpdated is a free log subscription operation binding the contract event 0x6308fd23d4b466bc07d9c9ef31631f5b96c7bac2efcb3d214fe3cc678e7ae00a.
//
// Solidity: e LogFeeUpdated(previousFee uint256, nextFee uint256)
func (_Orderbook *OrderbookFilterer) WatchLogFeeUpdated(opts *bind.WatchOpts, sink chan<- *OrderbookLogFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "LogFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookLogFeeUpdated)
				if err := _Orderbook.contract.UnpackLog(event, "LogFeeUpdated", log); err != nil {
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

// OrderbookOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Orderbook contract.
type OrderbookOwnershipRenouncedIterator struct {
	Event *OrderbookOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OrderbookOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookOwnershipRenounced)
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
		it.Event = new(OrderbookOwnershipRenounced)
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
func (it *OrderbookOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookOwnershipRenounced represents a OwnershipRenounced event raised by the Orderbook contract.
type OrderbookOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Orderbook *OrderbookFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OrderbookOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrderbookOwnershipRenouncedIterator{contract: _Orderbook.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Orderbook *OrderbookFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OrderbookOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookOwnershipRenounced)
				if err := _Orderbook.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OrderbookOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Orderbook contract.
type OrderbookOwnershipTransferredIterator struct {
	Event *OrderbookOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OrderbookOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookOwnershipTransferred)
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
		it.Event = new(OrderbookOwnershipTransferred)
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
func (it *OrderbookOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookOwnershipTransferred represents a OwnershipTransferred event raised by the Orderbook contract.
type OrderbookOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Orderbook *OrderbookFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OrderbookOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrderbookOwnershipTransferredIterator{contract: _Orderbook.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Orderbook *OrderbookFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OrderbookOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookOwnershipTransferred)
				if err := _Orderbook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561020b806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610130565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661013f565b600054600160a060020a031633146100db57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461015657600080fd5b61015f81610162565b50565b600160a060020a038116151561017757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820d07f4d5221378e0ccfc5230ace30fcd5568360922dc2baa242decb0ba9afd4100029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x608060405260008054600160a860020a031916331790556103c4806100256000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a811461007c5780635c975abb14610093578063715018a6146100bc5780638456cb59146100d15780638da5cb5b146100e6578063f2fde38b14610117575b600080fd5b34801561008857600080fd5b50610091610138565b005b34801561009f57600080fd5b506100a86101bf565b604080519115158252519081900360200190f35b3480156100c857600080fd5b506100916101e0565b3480156100dd57600080fd5b5061009161024c565b3480156100f257600080fd5b506100fb6102e9565b60408051600160a060020a039092168252519081900360200190f35b34801561012357600080fd5b50610091600160a060020a03600435166102f8565b600054600160a060020a0316331461014f57600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561017857600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031633146101f757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461026357600080fd5b60005474010000000000000000000000000000000000000000900460ff161561028b57600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b600054600160a060020a0316331461030f57600080fd5b6103188161031b565b50565b600160a060020a038116151561033057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058207d0a7a42de164bf3e33f827ae05b6303ba2a10e703b3c9ca4bead9f8150273bf0029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Pausable *PausableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Pausable contract.
type PausableOwnershipRenouncedIterator struct {
	Event *PausableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipRenounced)
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
		it.Event = new(PausableOwnershipRenounced)
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
func (it *PausableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipRenounced represents a OwnershipRenounced event raised by the Pausable contract.
type PausableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Pausable *PausableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PausableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipRenouncedIterator{contract: _Pausable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Pausable *PausableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PausableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipRenounced)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// PausableTokenABI is the input ABI used to generate the binding from.
const PausableTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PausableTokenBin is the compiled bytecode used for deploying new contracts.
const PausableTokenBin = `0x608060405260038054600160a860020a03191633179055610a84806100256000396000f3006080604052600436106100cf5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100d457806318160ddd1461010c57806323b872dd146101335780633f4ba83a1461015d5780635c975abb14610174578063661884631461018957806370a08231146101ad578063715018a6146101ce5780638456cb59146101e35780638da5cb5b146101f8578063a9059cbb14610229578063d73dd6231461024d578063dd62ed3e14610271578063f2fde38b14610298575b600080fd5b3480156100e057600080fd5b506100f8600160a060020a03600435166024356102b9565b604080519115158252519081900360200190f35b34801561011857600080fd5b506101216102e4565b60408051918252519081900360200190f35b34801561013f57600080fd5b506100f8600160a060020a03600435811690602435166044356102ea565b34801561016957600080fd5b50610172610317565b005b34801561018057600080fd5b506100f861038f565b34801561019557600080fd5b506100f8600160a060020a036004351660243561039f565b3480156101b957600080fd5b50610121600160a060020a03600435166103c3565b3480156101da57600080fd5b506101726103de565b3480156101ef57600080fd5b5061017261044c565b34801561020457600080fd5b5061020d6104c9565b60408051600160a060020a039092168252519081900360200190f35b34801561023557600080fd5b506100f8600160a060020a03600435166024356104d8565b34801561025957600080fd5b506100f8600160a060020a03600435166024356104fc565b34801561027d57600080fd5b50610121600160a060020a0360043581169060243516610520565b3480156102a457600080fd5b50610172600160a060020a036004351661054b565b60035460009060a060020a900460ff16156102d357600080fd5b6102dd838361056e565b9392505050565b60015490565b60035460009060a060020a900460ff161561030457600080fd5b61030f8484846105d4565b949350505050565b600354600160a060020a0316331461032e57600080fd5b60035460a060020a900460ff16151561034657600080fd5b6003805474ff0000000000000000000000000000000000000000191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b60035460a060020a900460ff1681565b60035460009060a060020a900460ff16156103b957600080fd5b6102dd838361074b565b600160a060020a031660009081526020819052604090205490565b600354600160a060020a031633146103f557600080fd5b600354604051600160a060020a03909116907ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482090600090a26003805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354600160a060020a0316331461046357600080fd5b60035460a060020a900460ff161561047a57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600354600160a060020a031681565b60035460009060a060020a900460ff16156104f257600080fd5b6102dd838361083b565b60035460009060a060020a900460ff161561051657600080fd5b6102dd838361091c565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a0316331461056257600080fd5b61056b816109b5565b50565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b6000600160a060020a03831615156105eb57600080fd5b600160a060020a03841660009081526020819052604090205482111561061057600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561064057600080fd5b600160a060020a038416600090815260208190526040902054610669908363ffffffff610a3316565b600160a060020a03808616600090815260208190526040808220939093559085168152205461069e908363ffffffff610a4516565b600160a060020a038085166000908152602081815260408083209490945591871681526002825282812033825290915220546106e0908363ffffffff610a3316565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b336000908152600260209081526040808320600160a060020a0386168452909152812054808311156107a057336000908152600260209081526040808320600160a060020a03881684529091528120556107d5565b6107b0818463ffffffff610a3316565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6000600160a060020a038316151561085257600080fd5b3360009081526020819052604090205482111561086e57600080fd5b3360009081526020819052604090205461088e908363ffffffff610a3316565b3360009081526020819052604080822092909255600160a060020a038516815220546108c0908363ffffffff610a4516565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a0386168452909152812054610950908363ffffffff610a4516565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03811615156109ca57600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610a3f57fe5b50900390565b81810182811015610a5257fe5b929150505600a165627a7a72305820b74990bbc2c31603ee33b72f704b7d064e6449244ddb06fbfd184f6df591549f0029`

// DeployPausableToken deploys a new Ethereum contract, binding an instance of PausableToken to it.
func DeployPausableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PausableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// PausableToken is an auto generated Go binding around an Ethereum contract.
type PausableToken struct {
	PausableTokenCaller     // Read-only binding to the contract
	PausableTokenTransactor // Write-only binding to the contract
	PausableTokenFilterer   // Log filterer for contract events
}

// PausableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableTokenSession struct {
	Contract     *PausableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableTokenCallerSession struct {
	Contract *PausableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PausableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTokenTransactorSession struct {
	Contract     *PausableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PausableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableTokenRaw struct {
	Contract *PausableToken // Generic contract binding to access the raw methods on
}

// PausableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableTokenCallerRaw struct {
	Contract *PausableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTokenTransactorRaw struct {
	Contract *PausableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableToken creates a new instance of PausableToken, bound to a specific deployed contract.
func NewPausableToken(address common.Address, backend bind.ContractBackend) (*PausableToken, error) {
	contract, err := bindPausableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// NewPausableTokenCaller creates a new read-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenCaller(address common.Address, caller bind.ContractCaller) (*PausableTokenCaller, error) {
	contract, err := bindPausableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenCaller{contract: contract}, nil
}

// NewPausableTokenTransactor creates a new write-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTokenTransactor, error) {
	contract, err := bindPausableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransactor{contract: contract}, nil
}

// NewPausableTokenFilterer creates a new log filterer instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableTokenFilterer, error) {
	contract, err := bindPausableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableTokenFilterer{contract: contract}, nil
}

// bindPausableToken binds a generic wrapper to an already deployed contract.
func bindPausableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.PausableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PausableToken *PausableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PausableToken *PausableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCallerSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCallerSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PausableToken *PausableTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PausableToken *PausableTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _PausableToken.Contract.RenounceOwnership(&_PausableToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PausableToken *PausableTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PausableToken.Contract.RenounceOwnership(&_PausableToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PausableToken *PausableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PausableToken *PausableTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PausableToken *PausableTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// PausableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PausableToken contract.
type PausableTokenApprovalIterator struct {
	Event *PausableTokenApproval // Event containing the contract specifics and raw log

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
func (it *PausableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenApproval)
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
		it.Event = new(PausableTokenApproval)
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
func (it *PausableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenApproval represents a Approval event raised by the PausableToken contract.
type PausableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PausableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenApprovalIterator{contract: _PausableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PausableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenApproval)
				if err := _PausableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PausableTokenOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PausableToken contract.
type PausableTokenOwnershipRenouncedIterator struct {
	Event *PausableTokenOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PausableTokenOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenOwnershipRenounced)
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
		it.Event = new(PausableTokenOwnershipRenounced)
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
func (it *PausableTokenOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenOwnershipRenounced represents a OwnershipRenounced event raised by the PausableToken contract.
type PausableTokenOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PausableToken *PausableTokenFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PausableTokenOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenOwnershipRenouncedIterator{contract: _PausableToken.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PausableToken *PausableTokenFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PausableTokenOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenOwnershipRenounced)
				if err := _PausableToken.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PausableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PausableToken contract.
type PausableTokenOwnershipTransferredIterator struct {
	Event *PausableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenOwnershipTransferred)
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
		it.Event = new(PausableTokenOwnershipTransferred)
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
func (it *PausableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the PausableToken contract.
type PausableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenOwnershipTransferredIterator{contract: _PausableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenOwnershipTransferred)
				if err := _PausableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PausableToken contract.
type PausableTokenPauseIterator struct {
	Event *PausableTokenPause // Event containing the contract specifics and raw log

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
func (it *PausableTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenPause)
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
		it.Event = new(PausableTokenPause)
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
func (it *PausableTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenPause represents a Pause event raised by the PausableToken contract.
type PausableTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PausableToken *PausableTokenFilterer) FilterPause(opts *bind.FilterOpts) (*PausableTokenPauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenPauseIterator{contract: _PausableToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PausableToken *PausableTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausableTokenPause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenPause)
				if err := _PausableToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PausableToken contract.
type PausableTokenTransferIterator struct {
	Event *PausableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PausableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenTransfer)
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
		it.Event = new(PausableTokenTransfer)
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
func (it *PausableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenTransfer represents a Transfer event raised by the PausableToken contract.
type PausableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PausableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransferIterator{contract: _PausableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PausableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenTransfer)
				if err := _PausableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// PausableTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PausableToken contract.
type PausableTokenUnpauseIterator struct {
	Event *PausableTokenUnpause // Event containing the contract specifics and raw log

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
func (it *PausableTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenUnpause)
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
		it.Event = new(PausableTokenUnpause)
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
func (it *PausableTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenUnpause represents a Unpause event raised by the PausableToken contract.
type PausableTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PausableToken *PausableTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableTokenUnpauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenUnpauseIterator{contract: _PausableToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PausableToken *PausableTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenUnpause)
				if err := _PausableToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// RenExBalancesABI is the input ABI used to generate the binding from.
const RenExBalancesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"brokerVerifierContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSettlementContract\",\"type\":\"address\"}],\"name\":\"updateRenExSettlementContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardVaultContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_traderFrom\",\"type\":\"address\"},{\"name\":\"_traderTo\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_feePayee\",\"type\":\"address\"}],\"name\":\"transferBalanceWithFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newBrokerVerifierContract\",\"type\":\"address\"}],\"name\":\"updateBrokerVerifierContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"signalBackupWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SIGNAL_DELAY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRewardVaultContract\",\"type\":\"address\"}],\"name\":\"updateRewardVaultContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"settlementContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETHEREUM\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderWithdrawalSignals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rewardVaultContract\",\"type\":\"address\"},{\"name\":\"_brokerVerifierContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogBalanceDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogBalanceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRenExSettlementContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newRenExSettlementContract\",\"type\":\"address\"}],\"name\":\"LogRenExSettlementContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRewardVaultContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newRewardVaultContract\",\"type\":\"address\"}],\"name\":\"LogRewardVaultContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousBrokerVerifierContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newBrokerVerifierContract\",\"type\":\"address\"}],\"name\":\"LogBrokerVerifierContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExBalancesBin is the compiled bytecode used for deploying new contracts.
const RenExBalancesBin = `0x608060405234801561001057600080fd5b506040516040806112b783398101604052805160209091015160008054600160a060020a0319908116331790915560038054600160a060020a0394851690831617905560028054939092169216919091179055611245806100726000396000f3006080604052600436106100d75763ffffffff60e060020a600035041663080bdfa881146100dc5780631d65551d1461010d57806323017a3a1461013057806331f092651461014557806334814e581461017657806347e7ef24146101b157806367aa50ae146101c8578063715018a6146101e95780638da5cb5b146101fe57806395b8cf5514610213578063c43c633b14610234578063e12cbb3c1461026d578063e1c3aedc14610282578063ea42418b146102a3578063f2fde38b146102b8578063f7cdf47c146102d9578063fc257baa146102ee575b600080fd5b3480156100e857600080fd5b506100f1610315565b60408051600160a060020a039092168252519081900360200190f35b34801561011957600080fd5b5061012e600160a060020a0360043516610324565b005b34801561013c57600080fd5b506100f16103b2565b34801561015157600080fd5b5061012e60048035600160a060020a03169060248035916044359182019101356103c1565b34801561018257600080fd5b5061012e600160a060020a03600435811690602435811690604435811690606435906084359060a43516610873565b61012e600160a060020a0360043516602435610b61565b3480156101d457600080fd5b5061012e600160a060020a0360043516610d63565b3480156101f557600080fd5b5061012e610df1565b34801561020a57600080fd5b506100f1610e5d565b34801561021f57600080fd5b5061012e600160a060020a0360043516610e6c565b34801561024057600080fd5b5061025b600160a060020a0360043581169060243516610e95565b60408051918252519081900360200190f35b34801561027957600080fd5b5061025b610eb2565b34801561028e57600080fd5b5061012e600160a060020a0360043516610eb9565b3480156102af57600080fd5b506100f1610f47565b3480156102c457600080fd5b5061012e600160a060020a0360043516610f56565b3480156102e557600080fd5b506100f1610f79565b3480156102fa57600080fd5b5061025b600160a060020a0360043581169060243516610f91565b600254600160a060020a031681565b600054600160a060020a0316331461033b57600080fd5b60015460408051600160a060020a039283168152918316602083015280517f8108e388ca125ce52d732c3507bc30e14194e147d133753ca55d6b5f109467ae9281900390910190a16001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600160a060020a031681565b60008483838080601f0160208091040260200160405190810160405280939291908181526020018383808284375050600254604080517fc043df8c00000000000000000000000000000000000000000000000000000000815233600482018181526024830193845289516044840152895191985060009750879650600160a060020a03909416945063c043df8c9388938a93919290916064019060208501908083838c5b8381101561047d578181015183820152602001610465565b50505050905090810190601f1680156104aa5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156104ca57600080fd5b505af11580156104de573d6000803e3d6000fd5b505050506040513d60208110156104f457600080fd5b50511561065457339550610509868b8b610fae565b600160a060020a038a1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141561056a57604051600160a060020a038716908a156108fc02908b906000818181858888f19350505050158015610564573d6000803e3d6000fd5b5061064f565b89600160a060020a031663a9059cbb878b6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156105cd57600080fd5b505af11580156105e1573d6000803e3d6000fd5b505050506040513d60208110156105f757600080fd5b5051151561064f576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b610867565b5050600160a060020a0381811660009081526005602090815260408083209387168352929052205480158015916202a3004291909103119082906106955750805b15156106eb576040805160e560020a62461bcd02815260206004820152600d60248201527f6e6f74207369676e616c6c656400000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038084166000908152600560209081526040808320938916835292905290812055339550610721868b8b610fae565b600160a060020a038a1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141561078257604051600160a060020a038716908a156108fc02908b906000818181858888f1935050505015801561077c573d6000803e3d6000fd5b50610867565b89600160a060020a031663a9059cbb878b6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156107e557600080fd5b505af11580156107f9573d6000803e3d6000fd5b505050506040513d602081101561080f57600080fd5b50511515610867576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b50505050505050505050565b600154600160a060020a031633146108d5576040805160e560020a62461bcd02815260206004820152600e60248201527f6e6f7420617574686f7269736564000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03808716600090815260046020908152604080832093881683529290522054821115610952576040805160e560020a62461bcd02815260206004820152601a60248201527f696e73756666696369656e742066756e647320666f7220666565000000000000604482015290519081900360640190fd5b600160a060020a03841673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415610a0e57600354604080517f8340f549000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015287811660248301526044820186905291519190921691638340f54991859160648082019260009290919082900301818588803b1580156109f057600080fd5b505af1158015610a04573d6000803e3d6000fd5b5050505050610b38565b600354604080517f095ea7b3000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810185905290519186169163095ea7b3916044808201926020929091908290030181600087803b158015610a7d57600080fd5b505af1158015610a91573d6000803e3d6000fd5b505050506040513d6020811015610aa757600080fd5b5050600354604080517f8340f549000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015287811660248301526044820186905291519190921691638340f54991606480830192600092919082900301818387803b158015610b1f57600080fd5b505af1158015610b33573d6000803e3d6000fd5b505050505b610b458685848601610fae565b6000831115610b5957610b598585856110d1565b505050505050565b33600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415610c0957348214610c04576040805160e560020a62461bcd02815260206004820152602760248201527f6d69736d6174636865642076616c756520706172616d6574657220616e64207460448201527f782076616c756500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610d53565b3415610c5f576040805160e560020a62461bcd02815260206004820152601960248201527f756e6578706563746564206574686572207472616e7366657200000000000000604482015290519081900360640190fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a038381166004830152306024830152604482018590529151918516916323b872dd916064808201926020929091908290030181600087803b158015610cd157600080fd5b505af1158015610ce5573d6000803e3d6000fd5b505050506040513d6020811015610cfb57600080fd5b50511515610d53576040805160e560020a62461bcd02815260206004820152601460248201527f746f6b656e2074726173666572206661696c6564000000000000000000000000604482015290519081900360640190fd5b610d5e8184846110d1565b505050565b600054600160a060020a03163314610d7a57600080fd5b60025460408051600160a060020a039283168152918316602083015280517f320ee0620117fe5b31d2b2ca97b2a711e2045489d6eb5290a73f04747f1819be9281900390910190a16002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a03163314610e0857600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b336000908152600560209081526040808320600160a060020a0394909416835292905220429055565b600460209081526000928352604080842090915290825290205481565b6202a30081565b600054600160a060020a03163314610ed057600080fd5b60035460408051600160a060020a039283168152918316602083015280517f6c348498f095f3d7eb84de1c0bf7fd7db8217d2fdd2af573ad0fa3642901c2459281900390910190a16003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600054600160a060020a03163314610f6d57600080fd5b610f7681611177565b50565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b600560209081526000928352604080842090915290825290205481565b600160a060020a0380841660009081526004602090815260408083209386168352929052205481111561102b576040805160e560020a62461bcd02815260206004820152601260248201527f696e73756666696369656e742066756e64730000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03808416600090815260046020908152604080832093861683529290522054611061908263ffffffff6111f416565b600160a060020a038085166000818152600460209081526040808320948816808452948252918290209490945580519182529281019190915280820183905290517f2622669645a3d6b14dc5d6134367eb988c1b617ea2df59f8aa0f102ba049977c9181900360600190a1505050565b600160a060020a03808416600090815260046020908152604080832093861683529290522054611107908263ffffffff61120616565b600160a060020a038085166000818152600460209081526040808320948816808452948252918290209490945580519182529281019190915280820183905290517f8d6dc51e8945c5e6a4ddb872612ec2b1f5e375a97daebdbb6a03a64e7c6592099181900360600190a1505050565b600160a060020a038116151561118c57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282111561120057fe5b50900390565b8181018281101561121357fe5b929150505600a165627a7a723058209d0d3bd3e31ee8810ccdffda571a66866f7ea8a3f48d45637c5d8506f73b6bd70029`

// DeployRenExBalances deploys a new Ethereum contract, binding an instance of RenExBalances to it.
func DeployRenExBalances(auth *bind.TransactOpts, backend bind.ContractBackend, _rewardVaultContract common.Address, _brokerVerifierContract common.Address) (common.Address, *types.Transaction, *RenExBalances, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBalancesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExBalancesBin), backend, _rewardVaultContract, _brokerVerifierContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExBalances{RenExBalancesCaller: RenExBalancesCaller{contract: contract}, RenExBalancesTransactor: RenExBalancesTransactor{contract: contract}, RenExBalancesFilterer: RenExBalancesFilterer{contract: contract}}, nil
}

// RenExBalances is an auto generated Go binding around an Ethereum contract.
type RenExBalances struct {
	RenExBalancesCaller     // Read-only binding to the contract
	RenExBalancesTransactor // Write-only binding to the contract
	RenExBalancesFilterer   // Log filterer for contract events
}

// RenExBalancesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExBalancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExBalancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExBalancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExBalancesSession struct {
	Contract     *RenExBalances    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExBalancesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExBalancesCallerSession struct {
	Contract *RenExBalancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RenExBalancesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExBalancesTransactorSession struct {
	Contract     *RenExBalancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RenExBalancesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExBalancesRaw struct {
	Contract *RenExBalances // Generic contract binding to access the raw methods on
}

// RenExBalancesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExBalancesCallerRaw struct {
	Contract *RenExBalancesCaller // Generic read-only contract binding to access the raw methods on
}

// RenExBalancesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExBalancesTransactorRaw struct {
	Contract *RenExBalancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExBalances creates a new instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalances(address common.Address, backend bind.ContractBackend) (*RenExBalances, error) {
	contract, err := bindRenExBalances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExBalances{RenExBalancesCaller: RenExBalancesCaller{contract: contract}, RenExBalancesTransactor: RenExBalancesTransactor{contract: contract}, RenExBalancesFilterer: RenExBalancesFilterer{contract: contract}}, nil
}

// NewRenExBalancesCaller creates a new read-only instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesCaller(address common.Address, caller bind.ContractCaller) (*RenExBalancesCaller, error) {
	contract, err := bindRenExBalances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesCaller{contract: contract}, nil
}

// NewRenExBalancesTransactor creates a new write-only instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExBalancesTransactor, error) {
	contract, err := bindRenExBalances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesTransactor{contract: contract}, nil
}

// NewRenExBalancesFilterer creates a new log filterer instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExBalancesFilterer, error) {
	contract, err := bindRenExBalances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesFilterer{contract: contract}, nil
}

// bindRenExBalances binds a generic wrapper to an already deployed contract.
func bindRenExBalances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBalancesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBalances *RenExBalancesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBalances.Contract.RenExBalancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBalances *RenExBalancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.Contract.RenExBalancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBalances *RenExBalancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBalances.Contract.RenExBalancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBalances *RenExBalancesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBalances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBalances *RenExBalancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBalances *RenExBalancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBalances.Contract.contract.Transact(opts, method, params...)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) ETHEREUM(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "ETHEREUM")
	return *ret0, err
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RenExBalances *RenExBalancesSession) ETHEREUM() (common.Address, error) {
	return _RenExBalances.Contract.ETHEREUM(&_RenExBalances.CallOpts)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) ETHEREUM() (common.Address, error) {
	return _RenExBalances.Contract.ETHEREUM(&_RenExBalances.CallOpts)
}

// SIGNALDELAY is a free data retrieval call binding the contract method 0xe12cbb3c.
//
// Solidity: function SIGNAL_DELAY() constant returns(uint256)
func (_RenExBalances *RenExBalancesCaller) SIGNALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "SIGNAL_DELAY")
	return *ret0, err
}

// SIGNALDELAY is a free data retrieval call binding the contract method 0xe12cbb3c.
//
// Solidity: function SIGNAL_DELAY() constant returns(uint256)
func (_RenExBalances *RenExBalancesSession) SIGNALDELAY() (*big.Int, error) {
	return _RenExBalances.Contract.SIGNALDELAY(&_RenExBalances.CallOpts)
}

// SIGNALDELAY is a free data retrieval call binding the contract method 0xe12cbb3c.
//
// Solidity: function SIGNAL_DELAY() constant returns(uint256)
func (_RenExBalances *RenExBalancesCallerSession) SIGNALDELAY() (*big.Int, error) {
	return _RenExBalances.Contract.SIGNALDELAY(&_RenExBalances.CallOpts)
}

// BrokerVerifierContract is a free data retrieval call binding the contract method 0x080bdfa8.
//
// Solidity: function brokerVerifierContract() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) BrokerVerifierContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "brokerVerifierContract")
	return *ret0, err
}

// BrokerVerifierContract is a free data retrieval call binding the contract method 0x080bdfa8.
//
// Solidity: function brokerVerifierContract() constant returns(address)
func (_RenExBalances *RenExBalancesSession) BrokerVerifierContract() (common.Address, error) {
	return _RenExBalances.Contract.BrokerVerifierContract(&_RenExBalances.CallOpts)
}

// BrokerVerifierContract is a free data retrieval call binding the contract method 0x080bdfa8.
//
// Solidity: function brokerVerifierContract() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) BrokerVerifierContract() (common.Address, error) {
	return _RenExBalances.Contract.BrokerVerifierContract(&_RenExBalances.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBalances *RenExBalancesSession) Owner() (common.Address, error) {
	return _RenExBalances.Contract.Owner(&_RenExBalances.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) Owner() (common.Address, error) {
	return _RenExBalances.Contract.Owner(&_RenExBalances.CallOpts)
}

// RewardVaultContract is a free data retrieval call binding the contract method 0x23017a3a.
//
// Solidity: function rewardVaultContract() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) RewardVaultContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "rewardVaultContract")
	return *ret0, err
}

// RewardVaultContract is a free data retrieval call binding the contract method 0x23017a3a.
//
// Solidity: function rewardVaultContract() constant returns(address)
func (_RenExBalances *RenExBalancesSession) RewardVaultContract() (common.Address, error) {
	return _RenExBalances.Contract.RewardVaultContract(&_RenExBalances.CallOpts)
}

// RewardVaultContract is a free data retrieval call binding the contract method 0x23017a3a.
//
// Solidity: function rewardVaultContract() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) RewardVaultContract() (common.Address, error) {
	return _RenExBalances.Contract.RewardVaultContract(&_RenExBalances.CallOpts)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) SettlementContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "settlementContract")
	return *ret0, err
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesSession) SettlementContract() (common.Address, error) {
	return _RenExBalances.Contract.SettlementContract(&_RenExBalances.CallOpts)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) SettlementContract() (common.Address, error) {
	return _RenExBalances.Contract.SettlementContract(&_RenExBalances.CallOpts)
}

// TraderBalances is a free data retrieval call binding the contract method 0xc43c633b.
//
// Solidity: function traderBalances( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesCaller) TraderBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "traderBalances", arg0, arg1)
	return *ret0, err
}

// TraderBalances is a free data retrieval call binding the contract method 0xc43c633b.
//
// Solidity: function traderBalances( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesSession) TraderBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RenExBalances.Contract.TraderBalances(&_RenExBalances.CallOpts, arg0, arg1)
}

// TraderBalances is a free data retrieval call binding the contract method 0xc43c633b.
//
// Solidity: function traderBalances( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesCallerSession) TraderBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RenExBalances.Contract.TraderBalances(&_RenExBalances.CallOpts, arg0, arg1)
}

// TraderWithdrawalSignals is a free data retrieval call binding the contract method 0xfc257baa.
//
// Solidity: function traderWithdrawalSignals( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesCaller) TraderWithdrawalSignals(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "traderWithdrawalSignals", arg0, arg1)
	return *ret0, err
}

// TraderWithdrawalSignals is a free data retrieval call binding the contract method 0xfc257baa.
//
// Solidity: function traderWithdrawalSignals( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesSession) TraderWithdrawalSignals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RenExBalances.Contract.TraderWithdrawalSignals(&_RenExBalances.CallOpts, arg0, arg1)
}

// TraderWithdrawalSignals is a free data retrieval call binding the contract method 0xfc257baa.
//
// Solidity: function traderWithdrawalSignals( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesCallerSession) TraderWithdrawalSignals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RenExBalances.Contract.TraderWithdrawalSignals(&_RenExBalances.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "deposit", _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesSession) Deposit(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.Deposit(&_RenExBalances.TransactOpts, _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactorSession) Deposit(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.Deposit(&_RenExBalances.TransactOpts, _token, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBalances *RenExBalancesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBalances *RenExBalancesSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBalances.Contract.RenounceOwnership(&_RenExBalances.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBalances *RenExBalancesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBalances.Contract.RenounceOwnership(&_RenExBalances.TransactOpts)
}

// SignalBackupWithdraw is a paid mutator transaction binding the contract method 0x95b8cf55.
//
// Solidity: function signalBackupWithdraw(_token address) returns()
func (_RenExBalances *RenExBalancesTransactor) SignalBackupWithdraw(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "signalBackupWithdraw", _token)
}

// SignalBackupWithdraw is a paid mutator transaction binding the contract method 0x95b8cf55.
//
// Solidity: function signalBackupWithdraw(_token address) returns()
func (_RenExBalances *RenExBalancesSession) SignalBackupWithdraw(_token common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.SignalBackupWithdraw(&_RenExBalances.TransactOpts, _token)
}

// SignalBackupWithdraw is a paid mutator transaction binding the contract method 0x95b8cf55.
//
// Solidity: function signalBackupWithdraw(_token address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) SignalBackupWithdraw(_token common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.SignalBackupWithdraw(&_RenExBalances.TransactOpts, _token)
}

// TransferBalanceWithFee is a paid mutator transaction binding the contract method 0x34814e58.
//
// Solidity: function transferBalanceWithFee(_traderFrom address, _traderTo address, _token address, _value uint256, _fee uint256, _feePayee address) returns()
func (_RenExBalances *RenExBalancesTransactor) TransferBalanceWithFee(opts *bind.TransactOpts, _traderFrom common.Address, _traderTo common.Address, _token common.Address, _value *big.Int, _fee *big.Int, _feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "transferBalanceWithFee", _traderFrom, _traderTo, _token, _value, _fee, _feePayee)
}

// TransferBalanceWithFee is a paid mutator transaction binding the contract method 0x34814e58.
//
// Solidity: function transferBalanceWithFee(_traderFrom address, _traderTo address, _token address, _value uint256, _fee uint256, _feePayee address) returns()
func (_RenExBalances *RenExBalancesSession) TransferBalanceWithFee(_traderFrom common.Address, _traderTo common.Address, _token common.Address, _value *big.Int, _fee *big.Int, _feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferBalanceWithFee(&_RenExBalances.TransactOpts, _traderFrom, _traderTo, _token, _value, _fee, _feePayee)
}

// TransferBalanceWithFee is a paid mutator transaction binding the contract method 0x34814e58.
//
// Solidity: function transferBalanceWithFee(_traderFrom address, _traderTo address, _token address, _value uint256, _fee uint256, _feePayee address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) TransferBalanceWithFee(_traderFrom common.Address, _traderTo common.Address, _token common.Address, _value *big.Int, _fee *big.Int, _feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferBalanceWithFee(&_RenExBalances.TransactOpts, _traderFrom, _traderTo, _token, _value, _fee, _feePayee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBalances *RenExBalancesTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBalances *RenExBalancesSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferOwnership(&_RenExBalances.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferOwnership(&_RenExBalances.TransactOpts, _newOwner)
}

// UpdateBrokerVerifierContract is a paid mutator transaction binding the contract method 0x67aa50ae.
//
// Solidity: function updateBrokerVerifierContract(_newBrokerVerifierContract address) returns()
func (_RenExBalances *RenExBalancesTransactor) UpdateBrokerVerifierContract(opts *bind.TransactOpts, _newBrokerVerifierContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "updateBrokerVerifierContract", _newBrokerVerifierContract)
}

// UpdateBrokerVerifierContract is a paid mutator transaction binding the contract method 0x67aa50ae.
//
// Solidity: function updateBrokerVerifierContract(_newBrokerVerifierContract address) returns()
func (_RenExBalances *RenExBalancesSession) UpdateBrokerVerifierContract(_newBrokerVerifierContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateBrokerVerifierContract(&_RenExBalances.TransactOpts, _newBrokerVerifierContract)
}

// UpdateBrokerVerifierContract is a paid mutator transaction binding the contract method 0x67aa50ae.
//
// Solidity: function updateBrokerVerifierContract(_newBrokerVerifierContract address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) UpdateBrokerVerifierContract(_newBrokerVerifierContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateBrokerVerifierContract(&_RenExBalances.TransactOpts, _newBrokerVerifierContract)
}

// UpdateRenExSettlementContract is a paid mutator transaction binding the contract method 0x1d65551d.
//
// Solidity: function updateRenExSettlementContract(_newSettlementContract address) returns()
func (_RenExBalances *RenExBalancesTransactor) UpdateRenExSettlementContract(opts *bind.TransactOpts, _newSettlementContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "updateRenExSettlementContract", _newSettlementContract)
}

// UpdateRenExSettlementContract is a paid mutator transaction binding the contract method 0x1d65551d.
//
// Solidity: function updateRenExSettlementContract(_newSettlementContract address) returns()
func (_RenExBalances *RenExBalancesSession) UpdateRenExSettlementContract(_newSettlementContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRenExSettlementContract(&_RenExBalances.TransactOpts, _newSettlementContract)
}

// UpdateRenExSettlementContract is a paid mutator transaction binding the contract method 0x1d65551d.
//
// Solidity: function updateRenExSettlementContract(_newSettlementContract address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) UpdateRenExSettlementContract(_newSettlementContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRenExSettlementContract(&_RenExBalances.TransactOpts, _newSettlementContract)
}

// UpdateRewardVaultContract is a paid mutator transaction binding the contract method 0xe1c3aedc.
//
// Solidity: function updateRewardVaultContract(_newRewardVaultContract address) returns()
func (_RenExBalances *RenExBalancesTransactor) UpdateRewardVaultContract(opts *bind.TransactOpts, _newRewardVaultContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "updateRewardVaultContract", _newRewardVaultContract)
}

// UpdateRewardVaultContract is a paid mutator transaction binding the contract method 0xe1c3aedc.
//
// Solidity: function updateRewardVaultContract(_newRewardVaultContract address) returns()
func (_RenExBalances *RenExBalancesSession) UpdateRewardVaultContract(_newRewardVaultContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRewardVaultContract(&_RenExBalances.TransactOpts, _newRewardVaultContract)
}

// UpdateRewardVaultContract is a paid mutator transaction binding the contract method 0xe1c3aedc.
//
// Solidity: function updateRewardVaultContract(_newRewardVaultContract address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) UpdateRewardVaultContract(_newRewardVaultContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRewardVaultContract(&_RenExBalances.TransactOpts, _newRewardVaultContract)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31f09265.
//
// Solidity: function withdraw(_token address, _value uint256, _signature bytes) returns()
func (_RenExBalances *RenExBalancesTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _value *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "withdraw", _token, _value, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31f09265.
//
// Solidity: function withdraw(_token address, _value uint256, _signature bytes) returns()
func (_RenExBalances *RenExBalancesSession) Withdraw(_token common.Address, _value *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RenExBalances.Contract.Withdraw(&_RenExBalances.TransactOpts, _token, _value, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31f09265.
//
// Solidity: function withdraw(_token address, _value uint256, _signature bytes) returns()
func (_RenExBalances *RenExBalancesTransactorSession) Withdraw(_token common.Address, _value *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RenExBalances.Contract.Withdraw(&_RenExBalances.TransactOpts, _token, _value, _signature)
}

// RenExBalancesLogBalanceDecreasedIterator is returned from FilterLogBalanceDecreased and is used to iterate over the raw logs and unpacked data for LogBalanceDecreased events raised by the RenExBalances contract.
type RenExBalancesLogBalanceDecreasedIterator struct {
	Event *RenExBalancesLogBalanceDecreased // Event containing the contract specifics and raw log

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
func (it *RenExBalancesLogBalanceDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesLogBalanceDecreased)
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
		it.Event = new(RenExBalancesLogBalanceDecreased)
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
func (it *RenExBalancesLogBalanceDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesLogBalanceDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesLogBalanceDecreased represents a LogBalanceDecreased event raised by the RenExBalances contract.
type RenExBalancesLogBalanceDecreased struct {
	Trader common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBalanceDecreased is a free log retrieval operation binding the contract event 0x2622669645a3d6b14dc5d6134367eb988c1b617ea2df59f8aa0f102ba049977c.
//
// Solidity: e LogBalanceDecreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) FilterLogBalanceDecreased(opts *bind.FilterOpts) (*RenExBalancesLogBalanceDecreasedIterator, error) {

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "LogBalanceDecreased")
	if err != nil {
		return nil, err
	}
	return &RenExBalancesLogBalanceDecreasedIterator{contract: _RenExBalances.contract, event: "LogBalanceDecreased", logs: logs, sub: sub}, nil
}

// WatchLogBalanceDecreased is a free log subscription operation binding the contract event 0x2622669645a3d6b14dc5d6134367eb988c1b617ea2df59f8aa0f102ba049977c.
//
// Solidity: e LogBalanceDecreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) WatchLogBalanceDecreased(opts *bind.WatchOpts, sink chan<- *RenExBalancesLogBalanceDecreased) (event.Subscription, error) {

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "LogBalanceDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesLogBalanceDecreased)
				if err := _RenExBalances.contract.UnpackLog(event, "LogBalanceDecreased", log); err != nil {
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

// RenExBalancesLogBalanceIncreasedIterator is returned from FilterLogBalanceIncreased and is used to iterate over the raw logs and unpacked data for LogBalanceIncreased events raised by the RenExBalances contract.
type RenExBalancesLogBalanceIncreasedIterator struct {
	Event *RenExBalancesLogBalanceIncreased // Event containing the contract specifics and raw log

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
func (it *RenExBalancesLogBalanceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesLogBalanceIncreased)
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
		it.Event = new(RenExBalancesLogBalanceIncreased)
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
func (it *RenExBalancesLogBalanceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesLogBalanceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesLogBalanceIncreased represents a LogBalanceIncreased event raised by the RenExBalances contract.
type RenExBalancesLogBalanceIncreased struct {
	Trader common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBalanceIncreased is a free log retrieval operation binding the contract event 0x8d6dc51e8945c5e6a4ddb872612ec2b1f5e375a97daebdbb6a03a64e7c659209.
//
// Solidity: e LogBalanceIncreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) FilterLogBalanceIncreased(opts *bind.FilterOpts) (*RenExBalancesLogBalanceIncreasedIterator, error) {

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "LogBalanceIncreased")
	if err != nil {
		return nil, err
	}
	return &RenExBalancesLogBalanceIncreasedIterator{contract: _RenExBalances.contract, event: "LogBalanceIncreased", logs: logs, sub: sub}, nil
}

// WatchLogBalanceIncreased is a free log subscription operation binding the contract event 0x8d6dc51e8945c5e6a4ddb872612ec2b1f5e375a97daebdbb6a03a64e7c659209.
//
// Solidity: e LogBalanceIncreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) WatchLogBalanceIncreased(opts *bind.WatchOpts, sink chan<- *RenExBalancesLogBalanceIncreased) (event.Subscription, error) {

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "LogBalanceIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesLogBalanceIncreased)
				if err := _RenExBalances.contract.UnpackLog(event, "LogBalanceIncreased", log); err != nil {
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

// RenExBalancesLogBrokerVerifierContractUpdatedIterator is returned from FilterLogBrokerVerifierContractUpdated and is used to iterate over the raw logs and unpacked data for LogBrokerVerifierContractUpdated events raised by the RenExBalances contract.
type RenExBalancesLogBrokerVerifierContractUpdatedIterator struct {
	Event *RenExBalancesLogBrokerVerifierContractUpdated // Event containing the contract specifics and raw log

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
func (it *RenExBalancesLogBrokerVerifierContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesLogBrokerVerifierContractUpdated)
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
		it.Event = new(RenExBalancesLogBrokerVerifierContractUpdated)
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
func (it *RenExBalancesLogBrokerVerifierContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesLogBrokerVerifierContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesLogBrokerVerifierContractUpdated represents a LogBrokerVerifierContractUpdated event raised by the RenExBalances contract.
type RenExBalancesLogBrokerVerifierContractUpdated struct {
	PreviousBrokerVerifierContract common.Address
	NewBrokerVerifierContract      common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterLogBrokerVerifierContractUpdated is a free log retrieval operation binding the contract event 0x320ee0620117fe5b31d2b2ca97b2a711e2045489d6eb5290a73f04747f1819be.
//
// Solidity: e LogBrokerVerifierContractUpdated(previousBrokerVerifierContract address, newBrokerVerifierContract address)
func (_RenExBalances *RenExBalancesFilterer) FilterLogBrokerVerifierContractUpdated(opts *bind.FilterOpts) (*RenExBalancesLogBrokerVerifierContractUpdatedIterator, error) {

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "LogBrokerVerifierContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExBalancesLogBrokerVerifierContractUpdatedIterator{contract: _RenExBalances.contract, event: "LogBrokerVerifierContractUpdated", logs: logs, sub: sub}, nil
}

// WatchLogBrokerVerifierContractUpdated is a free log subscription operation binding the contract event 0x320ee0620117fe5b31d2b2ca97b2a711e2045489d6eb5290a73f04747f1819be.
//
// Solidity: e LogBrokerVerifierContractUpdated(previousBrokerVerifierContract address, newBrokerVerifierContract address)
func (_RenExBalances *RenExBalancesFilterer) WatchLogBrokerVerifierContractUpdated(opts *bind.WatchOpts, sink chan<- *RenExBalancesLogBrokerVerifierContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "LogBrokerVerifierContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesLogBrokerVerifierContractUpdated)
				if err := _RenExBalances.contract.UnpackLog(event, "LogBrokerVerifierContractUpdated", log); err != nil {
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

// RenExBalancesLogRenExSettlementContractUpdatedIterator is returned from FilterLogRenExSettlementContractUpdated and is used to iterate over the raw logs and unpacked data for LogRenExSettlementContractUpdated events raised by the RenExBalances contract.
type RenExBalancesLogRenExSettlementContractUpdatedIterator struct {
	Event *RenExBalancesLogRenExSettlementContractUpdated // Event containing the contract specifics and raw log

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
func (it *RenExBalancesLogRenExSettlementContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesLogRenExSettlementContractUpdated)
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
		it.Event = new(RenExBalancesLogRenExSettlementContractUpdated)
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
func (it *RenExBalancesLogRenExSettlementContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesLogRenExSettlementContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesLogRenExSettlementContractUpdated represents a LogRenExSettlementContractUpdated event raised by the RenExBalances contract.
type RenExBalancesLogRenExSettlementContractUpdated struct {
	PreviousRenExSettlementContract common.Address
	NewRenExSettlementContract      common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterLogRenExSettlementContractUpdated is a free log retrieval operation binding the contract event 0x8108e388ca125ce52d732c3507bc30e14194e147d133753ca55d6b5f109467ae.
//
// Solidity: e LogRenExSettlementContractUpdated(previousRenExSettlementContract address, newRenExSettlementContract address)
func (_RenExBalances *RenExBalancesFilterer) FilterLogRenExSettlementContractUpdated(opts *bind.FilterOpts) (*RenExBalancesLogRenExSettlementContractUpdatedIterator, error) {

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "LogRenExSettlementContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExBalancesLogRenExSettlementContractUpdatedIterator{contract: _RenExBalances.contract, event: "LogRenExSettlementContractUpdated", logs: logs, sub: sub}, nil
}

// WatchLogRenExSettlementContractUpdated is a free log subscription operation binding the contract event 0x8108e388ca125ce52d732c3507bc30e14194e147d133753ca55d6b5f109467ae.
//
// Solidity: e LogRenExSettlementContractUpdated(previousRenExSettlementContract address, newRenExSettlementContract address)
func (_RenExBalances *RenExBalancesFilterer) WatchLogRenExSettlementContractUpdated(opts *bind.WatchOpts, sink chan<- *RenExBalancesLogRenExSettlementContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "LogRenExSettlementContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesLogRenExSettlementContractUpdated)
				if err := _RenExBalances.contract.UnpackLog(event, "LogRenExSettlementContractUpdated", log); err != nil {
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

// RenExBalancesLogRewardVaultContractUpdatedIterator is returned from FilterLogRewardVaultContractUpdated and is used to iterate over the raw logs and unpacked data for LogRewardVaultContractUpdated events raised by the RenExBalances contract.
type RenExBalancesLogRewardVaultContractUpdatedIterator struct {
	Event *RenExBalancesLogRewardVaultContractUpdated // Event containing the contract specifics and raw log

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
func (it *RenExBalancesLogRewardVaultContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesLogRewardVaultContractUpdated)
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
		it.Event = new(RenExBalancesLogRewardVaultContractUpdated)
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
func (it *RenExBalancesLogRewardVaultContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesLogRewardVaultContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesLogRewardVaultContractUpdated represents a LogRewardVaultContractUpdated event raised by the RenExBalances contract.
type RenExBalancesLogRewardVaultContractUpdated struct {
	PreviousRewardVaultContract common.Address
	NewRewardVaultContract      common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterLogRewardVaultContractUpdated is a free log retrieval operation binding the contract event 0x6c348498f095f3d7eb84de1c0bf7fd7db8217d2fdd2af573ad0fa3642901c245.
//
// Solidity: e LogRewardVaultContractUpdated(previousRewardVaultContract address, newRewardVaultContract address)
func (_RenExBalances *RenExBalancesFilterer) FilterLogRewardVaultContractUpdated(opts *bind.FilterOpts) (*RenExBalancesLogRewardVaultContractUpdatedIterator, error) {

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "LogRewardVaultContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExBalancesLogRewardVaultContractUpdatedIterator{contract: _RenExBalances.contract, event: "LogRewardVaultContractUpdated", logs: logs, sub: sub}, nil
}

// WatchLogRewardVaultContractUpdated is a free log subscription operation binding the contract event 0x6c348498f095f3d7eb84de1c0bf7fd7db8217d2fdd2af573ad0fa3642901c245.
//
// Solidity: e LogRewardVaultContractUpdated(previousRewardVaultContract address, newRewardVaultContract address)
func (_RenExBalances *RenExBalancesFilterer) WatchLogRewardVaultContractUpdated(opts *bind.WatchOpts, sink chan<- *RenExBalancesLogRewardVaultContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "LogRewardVaultContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesLogRewardVaultContractUpdated)
				if err := _RenExBalances.contract.UnpackLog(event, "LogRewardVaultContractUpdated", log); err != nil {
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

// RenExBalancesOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExBalances contract.
type RenExBalancesOwnershipRenouncedIterator struct {
	Event *RenExBalancesOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExBalancesOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesOwnershipRenounced)
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
		it.Event = new(RenExBalancesOwnershipRenounced)
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
func (it *RenExBalancesOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesOwnershipRenounced represents a OwnershipRenounced event raised by the RenExBalances contract.
type RenExBalancesOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExBalancesOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesOwnershipRenouncedIterator{contract: _RenExBalances.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExBalancesOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesOwnershipRenounced)
				if err := _RenExBalances.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExBalancesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExBalances contract.
type RenExBalancesOwnershipTransferredIterator struct {
	Event *RenExBalancesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExBalancesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesOwnershipTransferred)
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
		it.Event = new(RenExBalancesOwnershipTransferred)
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
func (it *RenExBalancesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesOwnershipTransferred represents a OwnershipTransferred event raised by the RenExBalances contract.
type RenExBalancesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExBalancesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesOwnershipTransferredIterator{contract: _RenExBalances.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExBalancesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesOwnershipTransferred)
				if err := _RenExBalances.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExBrokerVerifierABI is the input ABI used to generate the binding from.
const RenExBrokerVerifierABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_broker\",\"type\":\"address\"}],\"name\":\"deregisterBroker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderNonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"brokers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trader\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"verifyWithdrawSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_broker\",\"type\":\"address\"}],\"name\":\"registerBroker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"broker\",\"type\":\"address\"}],\"name\":\"LogBrokerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"broker\",\"type\":\"address\"}],\"name\":\"LogBrokerDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExBrokerVerifierBin is the compiled bytecode used for deploying new contracts.
const RenExBrokerVerifierBin = `0x608060405260008054600160a060020a031916331790556109f6806100256000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663490618d18114610092578063506ee1ef146100b557806366874cc5146100e8578063715018a61461011d5780638da5cb5b14610132578063c043df8c14610163578063c684286814610190578063f2fde38b146101b1575b600080fd5b34801561009e57600080fd5b506100b3600160a060020a03600435166101d2565b005b3480156100c157600080fd5b506100d6600160a060020a03600435166102ca565b60408051918252519081900360200190f35b3480156100f457600080fd5b50610109600160a060020a03600435166102dc565b604080519115158252519081900360200190f35b34801561012957600080fd5b506100b36102f1565b34801561013e57600080fd5b5061014761035d565b60408051600160a060020a039092168252519081900360200190f35b34801561016f57600080fd5b5061010960048035600160a060020a0316906024803590810191013561036c565b34801561019c57600080fd5b506100b3600160a060020a0360043516610473565b3480156101bd57600080fd5b506100b3600160a060020a036004351661056e565b600054600160a060020a031633146101e957600080fd5b600160a060020a03811660009081526001602052604090205460ff16151561027257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6e6f742072656769737465726564000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038116600081815260016020908152604091829020805460ff19169055815192835290517fe470a29f46ba9a09f7ec358ae2eb422a5a8f941f128ed7d8f5cf35278ab216409281900390910190a150565b60026020526000908152604090205481565b60016020526000908152604090205460ff1681565b600054600160a060020a0316331461030857600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600160a060020a03831660008181526002602090815260408083205481517f52657075626c69632050726f746f636f6c3a2077697468647261773a20000000818501526c01000000000000000000000000909502603d860152605180860191909152815180860390910181526091601f87018490049093028501830190915260718401858152929390928492610418928592918991899182910183828082843750610591945050505050565b600160a060020a03811660009081526001602052604090205490915060ff161561046557600160a060020a038616600090815260026020526040902080546001908101909155925061046a565b600092505b50509392505050565b600054600160a060020a0316331461048a57600080fd5b600160a060020a03811660009081526001602052604090205460ff161561051257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f616c726561647920726567697374657265640000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038116600081815260016020818152604092839020805460ff1916909217909155815192835290517fd4ba9549a2404d1e5bedd0a4ae90c79e2b41ce4dea6bef98dc999fec1f2784939281900390910190a150565b600054600160a060020a0316331461058557600080fd5b61058e8161074b565b50565b600060608060006040805190810160405280601a81526020017f19457468657265756d205369676e6564204d6573736167653a0a0000000000008152509250826105db87516107c8565b876040516020018084805190602001908083835b6020831061060e5780518252601f1990920191602091820191016105ef565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106106565780518252601f199092019160209182019101610637565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061069e5780518252601f19909201916020918201910161067f565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040529150816040518082805190602001908083835b602083106107065780518252601f1990920191602091820191016106e7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061073f81866108fa565b93505b50505092915050565b600160a060020a038116151561076057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60608160008083818415156108125760408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015295506108f0565b600093508492505b600083111561083457600a8304925060018401935061081a565b836040519080825280601f01601f191660200182016040528015610862578160200160208202803883390190505b509150600090505b838110156108ec57600a85066030017f010000000000000000000000000000000000000000000000000000000000000002826001838703038151811015156108ae57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a8504945060010161086a565b8195505b5050505050919050565b600080600080845160411415156109145760009350610742565b50505060208201516040830151606084015160001a601b60ff8216101561093957601b015b8060ff16601b1415801561095157508060ff16601c14155b1561095f5760009350610742565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af11580156109b9573d6000803e3d6000fd5b5050506020604051035193506107425600a165627a7a72305820c6ce4d5e7cceff131c145f756c9b42d950d0a4d7376a1547d75a30d918818eb10029`

// DeployRenExBrokerVerifier deploys a new Ethereum contract, binding an instance of RenExBrokerVerifier to it.
func DeployRenExBrokerVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RenExBrokerVerifier, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBrokerVerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExBrokerVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExBrokerVerifier{RenExBrokerVerifierCaller: RenExBrokerVerifierCaller{contract: contract}, RenExBrokerVerifierTransactor: RenExBrokerVerifierTransactor{contract: contract}, RenExBrokerVerifierFilterer: RenExBrokerVerifierFilterer{contract: contract}}, nil
}

// RenExBrokerVerifier is an auto generated Go binding around an Ethereum contract.
type RenExBrokerVerifier struct {
	RenExBrokerVerifierCaller     // Read-only binding to the contract
	RenExBrokerVerifierTransactor // Write-only binding to the contract
	RenExBrokerVerifierFilterer   // Log filterer for contract events
}

// RenExBrokerVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExBrokerVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBrokerVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExBrokerVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBrokerVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExBrokerVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBrokerVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExBrokerVerifierSession struct {
	Contract     *RenExBrokerVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RenExBrokerVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExBrokerVerifierCallerSession struct {
	Contract *RenExBrokerVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RenExBrokerVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExBrokerVerifierTransactorSession struct {
	Contract     *RenExBrokerVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RenExBrokerVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExBrokerVerifierRaw struct {
	Contract *RenExBrokerVerifier // Generic contract binding to access the raw methods on
}

// RenExBrokerVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExBrokerVerifierCallerRaw struct {
	Contract *RenExBrokerVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// RenExBrokerVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExBrokerVerifierTransactorRaw struct {
	Contract *RenExBrokerVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExBrokerVerifier creates a new instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifier(address common.Address, backend bind.ContractBackend) (*RenExBrokerVerifier, error) {
	contract, err := bindRenExBrokerVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifier{RenExBrokerVerifierCaller: RenExBrokerVerifierCaller{contract: contract}, RenExBrokerVerifierTransactor: RenExBrokerVerifierTransactor{contract: contract}, RenExBrokerVerifierFilterer: RenExBrokerVerifierFilterer{contract: contract}}, nil
}

// NewRenExBrokerVerifierCaller creates a new read-only instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifierCaller(address common.Address, caller bind.ContractCaller) (*RenExBrokerVerifierCaller, error) {
	contract, err := bindRenExBrokerVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierCaller{contract: contract}, nil
}

// NewRenExBrokerVerifierTransactor creates a new write-only instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExBrokerVerifierTransactor, error) {
	contract, err := bindRenExBrokerVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierTransactor{contract: contract}, nil
}

// NewRenExBrokerVerifierFilterer creates a new log filterer instance of RenExBrokerVerifier, bound to a specific deployed contract.
func NewRenExBrokerVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExBrokerVerifierFilterer, error) {
	contract, err := bindRenExBrokerVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierFilterer{contract: contract}, nil
}

// bindRenExBrokerVerifier binds a generic wrapper to an already deployed contract.
func bindRenExBrokerVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBrokerVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBrokerVerifier *RenExBrokerVerifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBrokerVerifier.Contract.RenExBrokerVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBrokerVerifier *RenExBrokerVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenExBrokerVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBrokerVerifier *RenExBrokerVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenExBrokerVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBrokerVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.contract.Transact(opts, method, params...)
}

// Brokers is a free data retrieval call binding the contract method 0x66874cc5.
//
// Solidity: function brokers( address) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) Brokers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "brokers", arg0)
	return *ret0, err
}

// Brokers is a free data retrieval call binding the contract method 0x66874cc5.
//
// Solidity: function brokers( address) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) Brokers(arg0 common.Address) (bool, error) {
	return _RenExBrokerVerifier.Contract.Brokers(&_RenExBrokerVerifier.CallOpts, arg0)
}

// Brokers is a free data retrieval call binding the contract method 0x66874cc5.
//
// Solidity: function brokers( address) constant returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) Brokers(arg0 common.Address) (bool, error) {
	return _RenExBrokerVerifier.Contract.Brokers(&_RenExBrokerVerifier.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) Owner() (common.Address, error) {
	return _RenExBrokerVerifier.Contract.Owner(&_RenExBrokerVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) Owner() (common.Address, error) {
	return _RenExBrokerVerifier.Contract.Owner(&_RenExBrokerVerifier.CallOpts)
}

// TraderNonces is a free data retrieval call binding the contract method 0x506ee1ef.
//
// Solidity: function traderNonces( address) constant returns(uint256)
func (_RenExBrokerVerifier *RenExBrokerVerifierCaller) TraderNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExBrokerVerifier.contract.Call(opts, out, "traderNonces", arg0)
	return *ret0, err
}

// TraderNonces is a free data retrieval call binding the contract method 0x506ee1ef.
//
// Solidity: function traderNonces( address) constant returns(uint256)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) TraderNonces(arg0 common.Address) (*big.Int, error) {
	return _RenExBrokerVerifier.Contract.TraderNonces(&_RenExBrokerVerifier.CallOpts, arg0)
}

// TraderNonces is a free data retrieval call binding the contract method 0x506ee1ef.
//
// Solidity: function traderNonces( address) constant returns(uint256)
func (_RenExBrokerVerifier *RenExBrokerVerifierCallerSession) TraderNonces(arg0 common.Address) (*big.Int, error) {
	return _RenExBrokerVerifier.Contract.TraderNonces(&_RenExBrokerVerifier.CallOpts, arg0)
}

// DeregisterBroker is a paid mutator transaction binding the contract method 0x490618d1.
//
// Solidity: function deregisterBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) DeregisterBroker(opts *bind.TransactOpts, _broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "deregisterBroker", _broker)
}

// DeregisterBroker is a paid mutator transaction binding the contract method 0x490618d1.
//
// Solidity: function deregisterBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) DeregisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.DeregisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// DeregisterBroker is a paid mutator transaction binding the contract method 0x490618d1.
//
// Solidity: function deregisterBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) DeregisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.DeregisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// RegisterBroker is a paid mutator transaction binding the contract method 0xc6842868.
//
// Solidity: function registerBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) RegisterBroker(opts *bind.TransactOpts, _broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "registerBroker", _broker)
}

// RegisterBroker is a paid mutator transaction binding the contract method 0xc6842868.
//
// Solidity: function registerBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) RegisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RegisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// RegisterBroker is a paid mutator transaction binding the contract method 0xc6842868.
//
// Solidity: function registerBroker(_broker address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) RegisterBroker(_broker common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RegisterBroker(&_RenExBrokerVerifier.TransactOpts, _broker)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenounceOwnership(&_RenExBrokerVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.RenounceOwnership(&_RenExBrokerVerifier.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.TransferOwnership(&_RenExBrokerVerifier.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.TransferOwnership(&_RenExBrokerVerifier.TransactOpts, _newOwner)
}

// VerifyWithdrawSignature is a paid mutator transaction binding the contract method 0xc043df8c.
//
// Solidity: function verifyWithdrawSignature(_trader address, _signature bytes) returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactor) VerifyWithdrawSignature(opts *bind.TransactOpts, _trader common.Address, _signature []byte) (*types.Transaction, error) {
	return _RenExBrokerVerifier.contract.Transact(opts, "verifyWithdrawSignature", _trader, _signature)
}

// VerifyWithdrawSignature is a paid mutator transaction binding the contract method 0xc043df8c.
//
// Solidity: function verifyWithdrawSignature(_trader address, _signature bytes) returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierSession) VerifyWithdrawSignature(_trader common.Address, _signature []byte) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.VerifyWithdrawSignature(&_RenExBrokerVerifier.TransactOpts, _trader, _signature)
}

// VerifyWithdrawSignature is a paid mutator transaction binding the contract method 0xc043df8c.
//
// Solidity: function verifyWithdrawSignature(_trader address, _signature bytes) returns(bool)
func (_RenExBrokerVerifier *RenExBrokerVerifierTransactorSession) VerifyWithdrawSignature(_trader common.Address, _signature []byte) (*types.Transaction, error) {
	return _RenExBrokerVerifier.Contract.VerifyWithdrawSignature(&_RenExBrokerVerifier.TransactOpts, _trader, _signature)
}

// RenExBrokerVerifierLogBrokerDeregisteredIterator is returned from FilterLogBrokerDeregistered and is used to iterate over the raw logs and unpacked data for LogBrokerDeregistered events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerDeregisteredIterator struct {
	Event *RenExBrokerVerifierLogBrokerDeregistered // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierLogBrokerDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierLogBrokerDeregistered)
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
		it.Event = new(RenExBrokerVerifierLogBrokerDeregistered)
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
func (it *RenExBrokerVerifierLogBrokerDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierLogBrokerDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierLogBrokerDeregistered represents a LogBrokerDeregistered event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerDeregistered struct {
	Broker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBrokerDeregistered is a free log retrieval operation binding the contract event 0xe470a29f46ba9a09f7ec358ae2eb422a5a8f941f128ed7d8f5cf35278ab21640.
//
// Solidity: e LogBrokerDeregistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterLogBrokerDeregistered(opts *bind.FilterOpts) (*RenExBrokerVerifierLogBrokerDeregisteredIterator, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "LogBrokerDeregistered")
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierLogBrokerDeregisteredIterator{contract: _RenExBrokerVerifier.contract, event: "LogBrokerDeregistered", logs: logs, sub: sub}, nil
}

// WatchLogBrokerDeregistered is a free log subscription operation binding the contract event 0xe470a29f46ba9a09f7ec358ae2eb422a5a8f941f128ed7d8f5cf35278ab21640.
//
// Solidity: e LogBrokerDeregistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchLogBrokerDeregistered(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierLogBrokerDeregistered) (event.Subscription, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "LogBrokerDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierLogBrokerDeregistered)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "LogBrokerDeregistered", log); err != nil {
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

// RenExBrokerVerifierLogBrokerRegisteredIterator is returned from FilterLogBrokerRegistered and is used to iterate over the raw logs and unpacked data for LogBrokerRegistered events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerRegisteredIterator struct {
	Event *RenExBrokerVerifierLogBrokerRegistered // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierLogBrokerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierLogBrokerRegistered)
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
		it.Event = new(RenExBrokerVerifierLogBrokerRegistered)
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
func (it *RenExBrokerVerifierLogBrokerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierLogBrokerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierLogBrokerRegistered represents a LogBrokerRegistered event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierLogBrokerRegistered struct {
	Broker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBrokerRegistered is a free log retrieval operation binding the contract event 0xd4ba9549a2404d1e5bedd0a4ae90c79e2b41ce4dea6bef98dc999fec1f278493.
//
// Solidity: e LogBrokerRegistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterLogBrokerRegistered(opts *bind.FilterOpts) (*RenExBrokerVerifierLogBrokerRegisteredIterator, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "LogBrokerRegistered")
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierLogBrokerRegisteredIterator{contract: _RenExBrokerVerifier.contract, event: "LogBrokerRegistered", logs: logs, sub: sub}, nil
}

// WatchLogBrokerRegistered is a free log subscription operation binding the contract event 0xd4ba9549a2404d1e5bedd0a4ae90c79e2b41ce4dea6bef98dc999fec1f278493.
//
// Solidity: e LogBrokerRegistered(broker address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchLogBrokerRegistered(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierLogBrokerRegistered) (event.Subscription, error) {

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "LogBrokerRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierLogBrokerRegistered)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "LogBrokerRegistered", log); err != nil {
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

// RenExBrokerVerifierOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipRenouncedIterator struct {
	Event *RenExBrokerVerifierOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierOwnershipRenounced)
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
		it.Event = new(RenExBrokerVerifierOwnershipRenounced)
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
func (it *RenExBrokerVerifierOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierOwnershipRenounced represents a OwnershipRenounced event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExBrokerVerifierOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierOwnershipRenouncedIterator{contract: _RenExBrokerVerifier.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierOwnershipRenounced)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExBrokerVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipTransferredIterator struct {
	Event *RenExBrokerVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExBrokerVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBrokerVerifierOwnershipTransferred)
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
		it.Event = new(RenExBrokerVerifierOwnershipTransferred)
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
func (it *RenExBrokerVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBrokerVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBrokerVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the RenExBrokerVerifier contract.
type RenExBrokerVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExBrokerVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBrokerVerifierOwnershipTransferredIterator{contract: _RenExBrokerVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBrokerVerifier *RenExBrokerVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExBrokerVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBrokerVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBrokerVerifierOwnershipTransferred)
				if err := _RenExBrokerVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExSettlementABI is the input ABI used to generate the binding from.
const RenExSettlementABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRenExTokensContract\",\"type\":\"address\"}],\"name\":\"updateRenExTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renExTokensContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submissionGasPriceLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"getMatchDetails\",\"outputs\":[{\"name\":\"settled\",\"type\":\"bool\"},{\"name\":\"priorityVolume\",\"type\":\"uint256\"},{\"name\":\"secondaryVolume\",\"type\":\"uint256\"},{\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"name\":\"priorityAddress\",\"type\":\"address\"},{\"name\":\"secondaryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOrderbookContract\",\"type\":\"address\"}],\"name\":\"updateOrderbook\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSubmissionGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"updateSubmissionGasPriceLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DARKNODE_FEES_DENOMINATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderSubmitter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RENEX_ATOMIC_SETTLEMENT_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderDetails\",\"outputs\":[{\"name\":\"settlementID\",\"type\":\"uint64\"},{\"name\":\"tokens\",\"type\":\"uint64\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"minimumVolume\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"matchTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DARKNODE_FEES_NUMERATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSlasherAddress\",\"type\":\"address\"}],\"name\":\"updateSlasher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prefix\",\"type\":\"bytes\"},{\"name\":\"_settlementID\",\"type\":\"uint64\"},{\"name\":\"_tokens\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_volume\",\"type\":\"uint256\"},{\"name\":\"_minimumVolume\",\"type\":\"uint256\"}],\"name\":\"submitOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"orderbookContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prefix\",\"type\":\"bytes\"},{\"name\":\"_settlementID\",\"type\":\"uint64\"},{\"name\":\"_tokens\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_volume\",\"type\":\"uint256\"},{\"name\":\"_minimumVolume\",\"type\":\"uint256\"}],\"name\":\"hashOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RENEX_SETTLEMENT_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyID\",\"type\":\"bytes32\"},{\"name\":\"_sellID\",\"type\":\"bytes32\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slasherAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renExBalancesContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRenExBalancesContract\",\"type\":\"address\"}],\"name\":\"updateRenExBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_guiltyOrderID\",\"type\":\"bytes32\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_orderbookContract\",\"type\":\"address\"},{\"name\":\"_renExTokensContract\",\"type\":\"address\"},{\"name\":\"_renExBalancesContract\",\"type\":\"address\"},{\"name\":\"_slasherAddress\",\"type\":\"address\"},{\"name\":\"_submissionGasPriceLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousOrderbook\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextOrderbook\",\"type\":\"address\"}],\"name\":\"LogOrderbookUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRenExTokens\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextRenExTokens\",\"type\":\"address\"}],\"name\":\"LogRenExTokensUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRenExBalances\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextRenExBalances\",\"type\":\"address\"}],\"name\":\"LogRenExBalancesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousSubmissionGasPriceLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextSubmissionGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"LogSubmissionGasPriceLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousSlasher\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextSlasher\",\"type\":\"address\"}],\"name\":\"LogSlasherUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExSettlementBin is the compiled bytecode used for deploying new contracts.
const RenExSettlementBin = `0x608060405234801561001057600080fd5b5060405160a08061265f8339810160409081528151602083015191830151606084015160809094015160008054600160a060020a0319908116331790915560018054600160a060020a0395861690831617905560028054958516958216959095179094556003805492841692851692909217909155600480549290941691909216179091556005556125b8806100a76000396000f30060806040526004361061013a5763ffffffff60e060020a6000350416632dff692d811461013f57806334106c891461017b5780634015e83b1461019e5780634c3052de146101cf5780634f54f4d8146101f65780636074b80614610254578063675df16f146102755780636af630d11461028d578063715018a6146102a257806375d4115e146102b75780638da5cb5b146102cf578063943a5e28146102e4578063a3c50b3214610312578063ab0fd37314610362578063aba880371461037d578063b3139d3814610392578063b86f602c146103b3578063bbe7221e146103f1578063ce0f92b714610406578063d23df2d114610444578063d32a9cd914610459578063d53c61bf14610474578063ddf25ce914610489578063ee0715ed1461049e578063f2fde38b146104bf578063f415ed14146104e0575b600080fd5b34801561014b57600080fd5b506101576004356104f8565b6040518082600381111561016757fe5b60ff16815260200191505060405180910390f35b34801561018757600080fd5b5061019c600160a060020a036004351661050d565b005b3480156101aa57600080fd5b506101b361058e565b60408051600160a060020a039092168252519081900360200190f35b3480156101db57600080fd5b506101e461059d565b60408051918252519081900360200190f35b34801561020257600080fd5b5061020e6004356105a3565b60408051971515885260208801969096528686019490945260608601929092526080850152600160a060020a0390811660a08501521660c0830152519081900360e00190f35b34801561026057600080fd5b5061019c600160a060020a0360043516610732565b34801561028157600080fd5b5061019c6004356107b3565b34801561029957600080fd5b506101e461080c565b3480156102ae57600080fd5b5061019c610812565b3480156102c357600080fd5b506101b3600435610871565b3480156102db57600080fd5b506101b361088c565b3480156102f057600080fd5b506102f961089b565b6040805163ffffffff9092168252519081900360200190f35b34801561031e57600080fd5b5061032a6004356108a0565b6040805167ffffffffffffffff9687168152949095166020850152838501929092526060830152608082015290519081900360a00190f35b34801561036e57600080fd5b506101e46004356024356108e0565b34801561038957600080fd5b506101e461089b565b34801561039e57600080fd5b5061019c600160a060020a03600435166108fd565b3480156103bf57600080fd5b5061019c602460048035828101929101359067ffffffffffffffff90358116906044351660643560843560a43561097e565b3480156103fd57600080fd5b506101b3610c84565b34801561041257600080fd5b506101e4602460048035828101929101359067ffffffffffffffff90358116906044351660643560843560a435610c93565b34801561045057600080fd5b506102f9610d18565b34801561046557600080fd5b5061019c600435602435610d1d565b34801561048057600080fd5b506101b36113ea565b34801561049557600080fd5b506101b36113f9565b3480156104aa57600080fd5b5061019c600160a060020a0360043516611408565b3480156104cb57600080fd5b5061019c600160a060020a0360043516611489565b3480156104ec57600080fd5b5061019c6004356114ac565b60086020526000908152604090205460ff1681565b600054600160a060020a0316331461052457600080fd5b60025460408051600160a060020a039283168152918316602083015280517fc44a7f49dd4281e6c3ed47edb754b69b064653d53ed217e1354e79e8fe4b06a09281900390910190a160028054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a031681565b60055481565b6000806000806000806000806000806105ba612482565b6105c26124a8565b600160009054906101000a9004600160a060020a0316600160a060020a031663af3e8a408e6040518263ffffffff1660e060020a028152600401808260001916600019168152602001915050602060405180830381600087803b15801561062857600080fd5b505af115801561063c573d6000803e3d6000fd5b505050506040513d602081101561065257600080fd5b5051945061065f8d61192b565b61066a57848d61066d565b8c855b60008281526006602052604090205491955093506106a09068010000000000000000900467ffffffffffffffff16611969565b91506106ad848484611b33565b9050600260008e81526008602052604090205460ff1660038111156106ce57fe5b14806106f65750600360008e81526008602052604090205460ff1660038111156106f457fe5b145b816000015182602001518360400151846060015185608001518660a001519b509b509b509b509b509b509b505050505050919395979092949650565b600054600160a060020a0316331461074957600080fd5b60015460408051600160a060020a039283168152918316602083015280517ff7af59918b82b5e13957d357d0fcc86f12a806b0d2e826bc24a0f13ae85e45989281900390910190a160018054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a031633146107ca57600080fd5b600554604080519182526020820183905280517fd0ef246766073915a6813492cff2a021d7cd4bf7d4feff3ef74327c7f4940e969281900390910190a1600555565b6103e881565b600054600160a060020a0316331461082957600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a260008054600160a060020a0319169055565b600760205260009081526040902054600160a060020a031681565b600054600160a060020a031681565b600281565b600660205260009081526040902080546001820154600283015460039093015467ffffffffffffffff808416946801000000000000000090940416929085565b600960209081526000928352604080842090915290825290205481565b600054600160a060020a0316331461091457600080fd5b60045460408051600160a060020a039283168152918316602083015280517f933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c89281900390910190a160048054600160a060020a031916600160a060020a0392909216919091179055565b6109866124f1565b6000600554803a111515156109e5576040805160e560020a62461bcd02815260206004820152601260248201527f67617320707269636520746f6f20686967680000000000000000000000000000604482015290519081900360640190fd5b60a0604051908101604052808967ffffffffffffffff1681526020018867ffffffffffffffff168152602001878152602001868152602001858152509250610a5e8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843750899450611cba9350505050565b91506000808381526008602052604090205460ff166003811115610a7e57fe5b14610ad3576040805160e560020a62461bcd02815260206004820152601760248201527f6f7264657220616c7265616479207375626d6974746564000000000000000000604482015290519081900360640190fd5b600154604080517faab14d04000000000000000000000000000000000000000000000000000000008152600481018590529051600292600160a060020a03169163aab14d049160248083019260209291908290030181600087803b158015610b3a57600080fd5b505af1158015610b4e573d6000803e3d6000fd5b505050506040513d6020811015610b6457600080fd5b50516003811115610b7157fe5b14610bc6576040805160e560020a62461bcd02815260206004820152601160248201527f756e636f6e6669726d6564206f72646572000000000000000000000000000000604482015290519081900360640190fd5b60008281526007602090815260408083208054600160a060020a031916331790556008909152902080546001919060ff19168280021790555050600090815260066020908152604091829020835181549285015167ffffffffffffffff1990931667ffffffffffffffff918216176fffffffffffffffff000000000000000019166801000000000000000091909316029190911781559082015160018201556060820151600282015560809091015160039091015550505050505050565b600154600160a060020a031681565b6000610d0c88888080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505060a0604051908101604052808967ffffffffffffffff1681526020018867ffffffffffffffff16815260200187815260200186815260200185815250611cba565b98975050505050505050565b600181565b610d25612482565b600080600160008681526008602052604090205460ff166003811115610d4757fe5b14610d9c576040805160e560020a62461bcd02815260206004820152601260248201527f696e76616c696420627579207374617475730000000000000000000000000000604482015290519081900360640190fd5b600160008581526008602052604090205460ff166003811115610dbb57fe5b14610e10576040805160e560020a62461bcd02815260206004820152601360248201527f696e76616c69642073656c6c2073746174757300000000000000000000000000604482015290519081900360640190fd5b60008581526006602052604090205467ffffffffffffffff1660021480610e4f575060008581526006602052604090205467ffffffffffffffff166001145b1515610ea5576040805160e560020a62461bcd02815260206004820152601560248201527f696e76616c696420736574746c656d656e742069640000000000000000000000604482015290519081900360640190fd5b610fff60066000876000191660001916815260200190815260200160002060a060405190810160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152505060066000876000191660001916815260200190815260200160002060a060405190810160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050611e22565b1515611055576040805160e560020a62461bcd02815260206004820152601360248201527f696e636f6d70617469626c65206f726465727300000000000000000000000000604482015290519081900360640190fd5b600154604080517faf3e8a400000000000000000000000000000000000000000000000000000000081526004810188905290518692600160a060020a03169163af3e8a409160248083019260209291908290030181600087803b1580156110bb57600080fd5b505af11580156110cf573d6000803e3d6000fd5b505050506040513d60208110156110e557600080fd5b50511461113c576040805160e560020a62461bcd02815260206004820152601260248201527f756e636f6e6669726d6564206f72646572730000000000000000000000000000604482015290519081900360640190fd5b60008581526006602052604090205461116a9068010000000000000000900467ffffffffffffffff16611969565b80516040015190935015156111c9576040805160e560020a62461bcd02815260206004820152601b60248201527f756e72656769737465726564207072696f7269747920746f6b656e0000000000604482015290519081900360640190fd5b8260200151604001511515611228576040805160e560020a62461bcd02815260206004820152601c60248201527f756e72656769737465726564207365636f6e6461727920746f6b656e00000000604482015290519081900360640190fd5b6001546040805160008051602061256d8339815191528152600481018890529051600160a060020a039092169163b1a08010916024808201926020929091908290030181600087803b15801561127d57600080fd5b505af1158015611291573d6000803e3d6000fd5b505050506040513d60208110156112a757600080fd5b50516001546040805160008051602061256d8339815191528152600481018890529051929450600160a060020a039091169163b1a08010916024808201926020929091908290030181600087803b15801561130157600080fd5b505af1158015611315573d6000803e3d6000fd5b505050506040513d602081101561132b57600080fd5b50519050600160a060020a038281169082161415611393576040805160e560020a62461bcd02815260206004820152601760248201527f6f72646572732066726f6d2073616d6520747261646572000000000000000000604482015290519081900360640190fd5b6113a08585848487611eb4565b50505060008281526009602090815260408083208484528252808320429055938252600890528281208054600260ff19918216811790925592825292902080549091169091179055565b600454600160a060020a031681565b600354600160a060020a031681565b600054600160a060020a0316331461141f57600080fd5b60035460408051600160a060020a039283168152918316602083015280517f28e85eee30dd92456f8c6864bcdaadb36644672cee6f285d571b1e58c08adca19281900390910190a160038054600160a060020a031916600160a060020a0392909216919091179055565b600054600160a060020a031633146114a057600080fd5b6114a981612068565b50565b60008060006114b9612482565b6114c16124a8565b600454600160a060020a03163314611523576040805160e560020a62461bcd02815260206004820152600c60248201527f756e617574686f72697365640000000000000000000000000000000000000000604482015290519081900360640190fd5b60008681526006602052604090205467ffffffffffffffff16600214611593576040805160e560020a62461bcd02815260206004820152601960248201527f736c617368696e67206e6f6e2d61746f6d696320747261646500000000000000604482015290519081900360640190fd5b600154604080517faf3e8a40000000000000000000000000000000000000000000000000000000008152600481018990529051600160a060020a039092169163af3e8a40916024808201926020929091908290030181600087803b1580156115fa57600080fd5b505af115801561160e573d6000803e3d6000fd5b505050506040513d602081101561162457600080fd5b50519450600260008781526008602052604090205460ff16600381111561164757fe5b1461169c576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c6964206f7264657220737461747573000000000000000000000000604482015290519081900360640190fd5b600260008681526008602052604090205460ff1660038111156116bb57fe5b14611710576040805160e560020a62461bcd02815260206004820152601460248201527f696e76616c6964206f7264657220737461747573000000000000000000000000604482015290519081900360640190fd5b6000868152600860205260409020805460ff191660031790556117328661192b565b61173d578486611740565b85855b60008281526006602052604090205491955093506117739068010000000000000000900467ffffffffffffffff16611969565b91506117808484846120d8565b6003546001546040805160008051602061256d8339815191528152600481018b90529051939450600160a060020a03928316936334814e58939092169163b1a08010916024808201926020929091908290030181600087803b1580156117e557600080fd5b505af11580156117f9573d6000803e3d6000fd5b505050506040513d602081101561180f57600080fd5b50516001546040805160008051602061256d8339815191528152600481018b90529051600160a060020a039092169163b1a08010916024808201926020929091908290030181600087803b15801561186657600080fd5b505af115801561187a573d6000803e3d6000fd5b505050506040513d602081101561189057600080fd5b5051608085015160408087015160048054835160e060020a63ffffffff8a16028152600160a060020a0397881692810192909252948616602482015292851660448401526064830181905260848301529290911660a4820152905160c480830192600092919082900301818387803b15801561190b57600080fd5b505af115801561191f573d6000803e3d6000fd5b50505050505050505050565b60009081526006602052604090205468010000000000000000900463ffffffff81811664010000000067ffffffffffffffff90931692909204161090565b611971612482565b600254604080517ffbb6272d00000000000000000000000000000000000000000000000000000000815263ffffffff64010000000067ffffffffffffffff871604166004820152905160009283928392839283928392600160a060020a039092169163fbb6272d9160248082019260609290919082900301818787803b1580156119fa57600080fd5b505af1158015611a0e573d6000803e3d6000fd5b505050506040513d6060811015611a2457600080fd5b508051602082015160409283015160025484517ffbb6272d00000000000000000000000000000000000000000000000000000000815263ffffffff8e1660048201529451939a509198509650600160a060020a03169163fbb6272d9160248083019260609291908290030181600087803b158015611aa157600080fd5b505af1158015611ab5573d6000803e3d6000fd5b505050506040513d6060811015611acb57600080fd5b508051602080830151604093840151845160a081018652600160a060020a039b8c1681870190815260ff9b8c166060808401919091529a1515608083015281528551998a0186529a909316885297909716868801521515908501525050509082015292915050565b611b3b6124a8565b611b43612535565b6000806000611b50612535565b611b58612535565b6040805190810160405280600660008c6000191660001916815260200190815260200160002060010154600660008e600019166000191681526020019081526020016000206001015401815260200160028152509550611bf4600660008c6000191660001916815260200190815260200160002060020154600660008c6000191660001916815260200190815260200160002060020154612331565b8651909550611c2990611c0e90879063ffffffff61234916565b8760200151600c808c600001516020015160ff160303612372565b9350611c45856001600c8b602001516020015160ff1603612372565b9250611c50846123d7565b9150611c5b836123d7565b6040805160c08101825284518152825160208083019190915280860151928201929092528183015160608201528a5151600160a060020a039081166080830152918b01515190911660a0820152975090505b5050505050509392505050565b600082826000015183602001518460400151856060015186608001516040516020018087805190602001908083835b60208310611d085780518252601f199092019160209182019101611ce9565b6001836020036101000a0380198251168184511680821785525050505050509050018667ffffffffffffffff1667ffffffffffffffff1678010000000000000000000000000000000000000000000000000281526008018567ffffffffffffffff1667ffffffffffffffff16780100000000000000000000000000000000000000000000000002815260080184815260200183815260200182815260200196505050505050506040516020818303038152906040526040518082805190602001908083835b60208310611dec5780518252601f199092019160209182019101611dcd565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090505b92915050565b6000611e3683602001518360200151612401565b1515611e4457506000611e1c565b816040015183604001511015611e5c57506000611e1c565b816080015183606001511015611e7457506000611e1c565b826080015182606001511015611e8c57506000611e1c565b8151835167ffffffffffffffff908116911614611eab57506000611e1c565b50600192915050565b611ebc6124a8565b60008681526006602052604090205467ffffffffffffffff16600214611eef57611ee7868684611b33565b905080611efe565b611efa8686846120d8565b9050805b6003546080820151825160408085015160008c815260076020528281205483517f34814e58000000000000000000000000000000000000000000000000000000008152600160a060020a038d811660048301528c8116602483015296871660448201526064810195909552608485019290925290841660a4840152905194955091909216926334814e589260c480820193929182900301818387803b158015611fa657600080fd5b505af1158015611fba573d6000803e3d6000fd5b505060035460a0840151602080860151606087015160008c8152600790935260408084205481517f34814e58000000000000000000000000000000000000000000000000000000008152600160a060020a038d811660048301528e8116602483015296871660448201526064810194909452608484019290925290841660a4830152519290931694506334814e58935060c48084019391929182900301818387803b15801561190b57600080fd5b600160a060020a038116151561207d57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a360008054600160a060020a031916600160a060020a0392909216919091179055565b6120e06124a8565b6120e8612535565b6000806120f3612535565b60006120fd612535565b6040805190810160405280600660008c6000191660001916815260200190815260200160002060010154600660008e600019166000191681526020019081526020016000206001015401815260200160028152509550612199600660008c6000191660001916815260200190815260200160002060020154600660008c6000191660001916815260200190815260200160002060020154612331565b94506121ac886020015160000151612474565b15612239576121cb856001600c8b602001516020015160ff1603612372565b93506121d6846123d7565b925060c06040519081016040528060008152602001600081526020018460200151815260200184602001518152602001896020015160000151600160a060020a03168152602001896020015160000151600160a060020a03168152509650611cad565b87515161224590612474565b156122bb57855161226190611c0e90879063ffffffff61234916565b915061226c826123d7565b6040805160c0810182526000808252602080830191909152830180519282019290925290516060820152895151600160a060020a0390811660808301528a51511660a082015297509050611cad565b6040805160e560020a62461bcd02815260206004820152602660248201527f6e6f6e2d6574682061746f6d696320737761707320617265206e6f742073757060448201527f706f727465640000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008183106123405781612342565b825b9392505050565b600082151561235a57506000611e1c565b5081810281838281151561236a57fe5b0414611e1c57fe5b6000808260010b1215156123b857604d600183900b131561238f57fe5b826123a785600185900b600a0a63ffffffff61234916565b8115156123b057fe5b049050612342565b8160000360010b600a0a83858115156123cd57fe5b048115156123b057fe5b6123df612535565b50604080518082019091526103e86103e6830204808252909103602082015290565b600064010000000067ffffffffffffffff83160463ffffffff908116908416148015612447575064010000000067ffffffffffffffff84160463ffffffff908116908316145b80156123425750505063ffffffff81811664010000000067ffffffffffffffff9093169290920416111590565b600160a060020a0316151590565b60c06040519081016040528061249661254c565b81526020016124a361254c565b905290565b60c060405190810160405280600081526020016000815260200160008152602001600081526020016000600160a060020a031681526020016000600160a060020a031681525090565b60a060405190810160405280600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000815260200160008152602001600081525090565b604080518082019091526000808252602082015290565b6040805160608101825260008082526020820181905291810191909152905600b1a0801000000000000000000000000000000000000000000000000000000000a165627a7a72305820e04a012bc3ea37dd95dfd18e9d010ab7ea8479ab906568281c14b116169271c10029`

// DeployRenExSettlement deploys a new Ethereum contract, binding an instance of RenExSettlement to it.
func DeployRenExSettlement(auth *bind.TransactOpts, backend bind.ContractBackend, _orderbookContract common.Address, _renExTokensContract common.Address, _renExBalancesContract common.Address, _slasherAddress common.Address, _submissionGasPriceLimit *big.Int) (common.Address, *types.Transaction, *RenExSettlement, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExSettlementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExSettlementBin), backend, _orderbookContract, _renExTokensContract, _renExBalancesContract, _slasherAddress, _submissionGasPriceLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExSettlement{RenExSettlementCaller: RenExSettlementCaller{contract: contract}, RenExSettlementTransactor: RenExSettlementTransactor{contract: contract}, RenExSettlementFilterer: RenExSettlementFilterer{contract: contract}}, nil
}

// RenExSettlement is an auto generated Go binding around an Ethereum contract.
type RenExSettlement struct {
	RenExSettlementCaller     // Read-only binding to the contract
	RenExSettlementTransactor // Write-only binding to the contract
	RenExSettlementFilterer   // Log filterer for contract events
}

// RenExSettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExSettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExSettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExSettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExSettlementSession struct {
	Contract     *RenExSettlement  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExSettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExSettlementCallerSession struct {
	Contract *RenExSettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RenExSettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExSettlementTransactorSession struct {
	Contract     *RenExSettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RenExSettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExSettlementRaw struct {
	Contract *RenExSettlement // Generic contract binding to access the raw methods on
}

// RenExSettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExSettlementCallerRaw struct {
	Contract *RenExSettlementCaller // Generic read-only contract binding to access the raw methods on
}

// RenExSettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExSettlementTransactorRaw struct {
	Contract *RenExSettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExSettlement creates a new instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlement(address common.Address, backend bind.ContractBackend) (*RenExSettlement, error) {
	contract, err := bindRenExSettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExSettlement{RenExSettlementCaller: RenExSettlementCaller{contract: contract}, RenExSettlementTransactor: RenExSettlementTransactor{contract: contract}, RenExSettlementFilterer: RenExSettlementFilterer{contract: contract}}, nil
}

// NewRenExSettlementCaller creates a new read-only instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementCaller(address common.Address, caller bind.ContractCaller) (*RenExSettlementCaller, error) {
	contract, err := bindRenExSettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementCaller{contract: contract}, nil
}

// NewRenExSettlementTransactor creates a new write-only instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExSettlementTransactor, error) {
	contract, err := bindRenExSettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementTransactor{contract: contract}, nil
}

// NewRenExSettlementFilterer creates a new log filterer instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExSettlementFilterer, error) {
	contract, err := bindRenExSettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementFilterer{contract: contract}, nil
}

// bindRenExSettlement binds a generic wrapper to an already deployed contract.
func bindRenExSettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExSettlementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExSettlement *RenExSettlementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExSettlement.Contract.RenExSettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExSettlement *RenExSettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenExSettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExSettlement *RenExSettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenExSettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExSettlement *RenExSettlementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExSettlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExSettlement *RenExSettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExSettlement *RenExSettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExSettlement.Contract.contract.Transact(opts, method, params...)
}

// DARKNODEFEESDENOMINATOR is a free data retrieval call binding the contract method 0x6af630d1.
//
// Solidity: function DARKNODE_FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) DARKNODEFEESDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "DARKNODE_FEES_DENOMINATOR")
	return *ret0, err
}

// DARKNODEFEESDENOMINATOR is a free data retrieval call binding the contract method 0x6af630d1.
//
// Solidity: function DARKNODE_FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) DARKNODEFEESDENOMINATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESDENOMINATOR(&_RenExSettlement.CallOpts)
}

// DARKNODEFEESDENOMINATOR is a free data retrieval call binding the contract method 0x6af630d1.
//
// Solidity: function DARKNODE_FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) DARKNODEFEESDENOMINATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESDENOMINATOR(&_RenExSettlement.CallOpts)
}

// DARKNODEFEESNUMERATOR is a free data retrieval call binding the contract method 0xaba88037.
//
// Solidity: function DARKNODE_FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) DARKNODEFEESNUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "DARKNODE_FEES_NUMERATOR")
	return *ret0, err
}

// DARKNODEFEESNUMERATOR is a free data retrieval call binding the contract method 0xaba88037.
//
// Solidity: function DARKNODE_FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) DARKNODEFEESNUMERATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESNUMERATOR(&_RenExSettlement.CallOpts)
}

// DARKNODEFEESNUMERATOR is a free data retrieval call binding the contract method 0xaba88037.
//
// Solidity: function DARKNODE_FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) DARKNODEFEESNUMERATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.DARKNODEFEESNUMERATOR(&_RenExSettlement.CallOpts)
}

// RENEXATOMICSETTLEMENTID is a free data retrieval call binding the contract method 0x943a5e28.
//
// Solidity: function RENEX_ATOMIC_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCaller) RENEXATOMICSETTLEMENTID(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "RENEX_ATOMIC_SETTLEMENT_ID")
	return *ret0, err
}

// RENEXATOMICSETTLEMENTID is a free data retrieval call binding the contract method 0x943a5e28.
//
// Solidity: function RENEX_ATOMIC_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementSession) RENEXATOMICSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXATOMICSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// RENEXATOMICSETTLEMENTID is a free data retrieval call binding the contract method 0x943a5e28.
//
// Solidity: function RENEX_ATOMIC_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCallerSession) RENEXATOMICSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXATOMICSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// RENEXSETTLEMENTID is a free data retrieval call binding the contract method 0xd23df2d1.
//
// Solidity: function RENEX_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCaller) RENEXSETTLEMENTID(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "RENEX_SETTLEMENT_ID")
	return *ret0, err
}

// RENEXSETTLEMENTID is a free data retrieval call binding the contract method 0xd23df2d1.
//
// Solidity: function RENEX_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementSession) RENEXSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// RENEXSETTLEMENTID is a free data retrieval call binding the contract method 0xd23df2d1.
//
// Solidity: function RENEX_SETTLEMENT_ID() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCallerSession) RENEXSETTLEMENTID() (uint32, error) {
	return _RenExSettlement.Contract.RENEXSETTLEMENTID(&_RenExSettlement.CallOpts)
}

// GetMatchDetails is a free data retrieval call binding the contract method 0x4f54f4d8.
//
// Solidity: function getMatchDetails(_orderID bytes32) constant returns(settled bool, priorityVolume uint256, secondaryVolume uint256, priorityFee uint256, secondaryFee uint256, priorityAddress address, secondaryAddress address)
func (_RenExSettlement *RenExSettlementCaller) GetMatchDetails(opts *bind.CallOpts, _orderID [32]byte) (struct {
	Settled          bool
	PriorityVolume   *big.Int
	SecondaryVolume  *big.Int
	PriorityFee      *big.Int
	SecondaryFee     *big.Int
	PriorityAddress  common.Address
	SecondaryAddress common.Address
}, error) {
	ret := new(struct {
		Settled          bool
		PriorityVolume   *big.Int
		SecondaryVolume  *big.Int
		PriorityFee      *big.Int
		SecondaryFee     *big.Int
		PriorityAddress  common.Address
		SecondaryAddress common.Address
	})
	out := ret
	err := _RenExSettlement.contract.Call(opts, out, "getMatchDetails", _orderID)
	return *ret, err
}

// GetMatchDetails is a free data retrieval call binding the contract method 0x4f54f4d8.
//
// Solidity: function getMatchDetails(_orderID bytes32) constant returns(settled bool, priorityVolume uint256, secondaryVolume uint256, priorityFee uint256, secondaryFee uint256, priorityAddress address, secondaryAddress address)
func (_RenExSettlement *RenExSettlementSession) GetMatchDetails(_orderID [32]byte) (struct {
	Settled          bool
	PriorityVolume   *big.Int
	SecondaryVolume  *big.Int
	PriorityFee      *big.Int
	SecondaryFee     *big.Int
	PriorityAddress  common.Address
	SecondaryAddress common.Address
}, error) {
	return _RenExSettlement.Contract.GetMatchDetails(&_RenExSettlement.CallOpts, _orderID)
}

// GetMatchDetails is a free data retrieval call binding the contract method 0x4f54f4d8.
//
// Solidity: function getMatchDetails(_orderID bytes32) constant returns(settled bool, priorityVolume uint256, secondaryVolume uint256, priorityFee uint256, secondaryFee uint256, priorityAddress address, secondaryAddress address)
func (_RenExSettlement *RenExSettlementCallerSession) GetMatchDetails(_orderID [32]byte) (struct {
	Settled          bool
	PriorityVolume   *big.Int
	SecondaryVolume  *big.Int
	PriorityFee      *big.Int
	SecondaryFee     *big.Int
	PriorityAddress  common.Address
	SecondaryAddress common.Address
}, error) {
	return _RenExSettlement.Contract.GetMatchDetails(&_RenExSettlement.CallOpts, _orderID)
}

// HashOrder is a free data retrieval call binding the contract method 0xce0f92b7.
//
// Solidity: function hashOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) constant returns(bytes32)
func (_RenExSettlement *RenExSettlementCaller) HashOrder(opts *bind.CallOpts, _prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "hashOrder", _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
	return *ret0, err
}

// HashOrder is a free data retrieval call binding the contract method 0xce0f92b7.
//
// Solidity: function hashOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) constant returns(bytes32)
func (_RenExSettlement *RenExSettlementSession) HashOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) ([32]byte, error) {
	return _RenExSettlement.Contract.HashOrder(&_RenExSettlement.CallOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// HashOrder is a free data retrieval call binding the contract method 0xce0f92b7.
//
// Solidity: function hashOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) constant returns(bytes32)
func (_RenExSettlement *RenExSettlementCallerSession) HashOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) ([32]byte, error) {
	return _RenExSettlement.Contract.HashOrder(&_RenExSettlement.CallOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// MatchTimestamp is a free data retrieval call binding the contract method 0xab0fd373.
//
// Solidity: function matchTimestamp( bytes32,  bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) MatchTimestamp(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "matchTimestamp", arg0, arg1)
	return *ret0, err
}

// MatchTimestamp is a free data retrieval call binding the contract method 0xab0fd373.
//
// Solidity: function matchTimestamp( bytes32,  bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) MatchTimestamp(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _RenExSettlement.Contract.MatchTimestamp(&_RenExSettlement.CallOpts, arg0, arg1)
}

// MatchTimestamp is a free data retrieval call binding the contract method 0xab0fd373.
//
// Solidity: function matchTimestamp( bytes32,  bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) MatchTimestamp(arg0 [32]byte, arg1 [32]byte) (*big.Int, error) {
	return _RenExSettlement.Contract.MatchTimestamp(&_RenExSettlement.CallOpts, arg0, arg1)
}

// OrderDetails is a free data retrieval call binding the contract method 0xa3c50b32.
//
// Solidity: function orderDetails( bytes32) constant returns(settlementID uint64, tokens uint64, price uint256, volume uint256, minimumVolume uint256)
func (_RenExSettlement *RenExSettlementCaller) OrderDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SettlementID  uint64
	Tokens        uint64
	Price         *big.Int
	Volume        *big.Int
	MinimumVolume *big.Int
}, error) {
	ret := new(struct {
		SettlementID  uint64
		Tokens        uint64
		Price         *big.Int
		Volume        *big.Int
		MinimumVolume *big.Int
	})
	out := ret
	err := _RenExSettlement.contract.Call(opts, out, "orderDetails", arg0)
	return *ret, err
}

// OrderDetails is a free data retrieval call binding the contract method 0xa3c50b32.
//
// Solidity: function orderDetails( bytes32) constant returns(settlementID uint64, tokens uint64, price uint256, volume uint256, minimumVolume uint256)
func (_RenExSettlement *RenExSettlementSession) OrderDetails(arg0 [32]byte) (struct {
	SettlementID  uint64
	Tokens        uint64
	Price         *big.Int
	Volume        *big.Int
	MinimumVolume *big.Int
}, error) {
	return _RenExSettlement.Contract.OrderDetails(&_RenExSettlement.CallOpts, arg0)
}

// OrderDetails is a free data retrieval call binding the contract method 0xa3c50b32.
//
// Solidity: function orderDetails( bytes32) constant returns(settlementID uint64, tokens uint64, price uint256, volume uint256, minimumVolume uint256)
func (_RenExSettlement *RenExSettlementCallerSession) OrderDetails(arg0 [32]byte) (struct {
	SettlementID  uint64
	Tokens        uint64
	Price         *big.Int
	Volume        *big.Int
	MinimumVolume *big.Int
}, error) {
	return _RenExSettlement.Contract.OrderDetails(&_RenExSettlement.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus( bytes32) constant returns(uint8)
func (_RenExSettlement *RenExSettlementCaller) OrderStatus(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "orderStatus", arg0)
	return *ret0, err
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus( bytes32) constant returns(uint8)
func (_RenExSettlement *RenExSettlementSession) OrderStatus(arg0 [32]byte) (uint8, error) {
	return _RenExSettlement.Contract.OrderStatus(&_RenExSettlement.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus( bytes32) constant returns(uint8)
func (_RenExSettlement *RenExSettlementCallerSession) OrderStatus(arg0 [32]byte) (uint8, error) {
	return _RenExSettlement.Contract.OrderStatus(&_RenExSettlement.CallOpts, arg0)
}

// OrderSubmitter is a free data retrieval call binding the contract method 0x75d4115e.
//
// Solidity: function orderSubmitter( bytes32) constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) OrderSubmitter(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "orderSubmitter", arg0)
	return *ret0, err
}

// OrderSubmitter is a free data retrieval call binding the contract method 0x75d4115e.
//
// Solidity: function orderSubmitter( bytes32) constant returns(address)
func (_RenExSettlement *RenExSettlementSession) OrderSubmitter(arg0 [32]byte) (common.Address, error) {
	return _RenExSettlement.Contract.OrderSubmitter(&_RenExSettlement.CallOpts, arg0)
}

// OrderSubmitter is a free data retrieval call binding the contract method 0x75d4115e.
//
// Solidity: function orderSubmitter( bytes32) constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) OrderSubmitter(arg0 [32]byte) (common.Address, error) {
	return _RenExSettlement.Contract.OrderSubmitter(&_RenExSettlement.CallOpts, arg0)
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) OrderbookContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "orderbookContract")
	return *ret0, err
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) OrderbookContract() (common.Address, error) {
	return _RenExSettlement.Contract.OrderbookContract(&_RenExSettlement.CallOpts)
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) OrderbookContract() (common.Address, error) {
	return _RenExSettlement.Contract.OrderbookContract(&_RenExSettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) Owner() (common.Address, error) {
	return _RenExSettlement.Contract.Owner(&_RenExSettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) Owner() (common.Address, error) {
	return _RenExSettlement.Contract.Owner(&_RenExSettlement.CallOpts)
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) RenExBalancesContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "renExBalancesContract")
	return *ret0, err
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) RenExBalancesContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExBalancesContract(&_RenExSettlement.CallOpts)
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) RenExBalancesContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExBalancesContract(&_RenExSettlement.CallOpts)
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) RenExTokensContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "renExTokensContract")
	return *ret0, err
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) RenExTokensContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExTokensContract(&_RenExSettlement.CallOpts)
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) RenExTokensContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExTokensContract(&_RenExSettlement.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) SlasherAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "slasherAddress")
	return *ret0, err
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) SlasherAddress() (common.Address, error) {
	return _RenExSettlement.Contract.SlasherAddress(&_RenExSettlement.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) SlasherAddress() (common.Address, error) {
	return _RenExSettlement.Contract.SlasherAddress(&_RenExSettlement.CallOpts)
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) SubmissionGasPriceLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "submissionGasPriceLimit")
	return *ret0, err
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) SubmissionGasPriceLimit() (*big.Int, error) {
	return _RenExSettlement.Contract.SubmissionGasPriceLimit(&_RenExSettlement.CallOpts)
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) SubmissionGasPriceLimit() (*big.Int, error) {
	return _RenExSettlement.Contract.SubmissionGasPriceLimit(&_RenExSettlement.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenounceOwnership(&_RenExSettlement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenounceOwnership(&_RenExSettlement.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0xd32a9cd9.
//
// Solidity: function settle(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactor) Settle(opts *bind.TransactOpts, _buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "settle", _buyID, _sellID)
}

// Settle is a paid mutator transaction binding the contract method 0xd32a9cd9.
//
// Solidity: function settle(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementSession) Settle(_buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Settle(&_RenExSettlement.TransactOpts, _buyID, _sellID)
}

// Settle is a paid mutator transaction binding the contract method 0xd32a9cd9.
//
// Solidity: function settle(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) Settle(_buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Settle(&_RenExSettlement.TransactOpts, _buyID, _sellID)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(_guiltyOrderID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactor) Slash(opts *bind.TransactOpts, _guiltyOrderID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "slash", _guiltyOrderID)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(_guiltyOrderID bytes32) returns()
func (_RenExSettlement *RenExSettlementSession) Slash(_guiltyOrderID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Slash(&_RenExSettlement.TransactOpts, _guiltyOrderID)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(_guiltyOrderID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) Slash(_guiltyOrderID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.Slash(&_RenExSettlement.TransactOpts, _guiltyOrderID)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0xb86f602c.
//
// Solidity: function submitOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) returns()
func (_RenExSettlement *RenExSettlementTransactor) SubmitOrder(opts *bind.TransactOpts, _prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "submitOrder", _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0xb86f602c.
//
// Solidity: function submitOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) returns()
func (_RenExSettlement *RenExSettlementSession) SubmitOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitOrder(&_RenExSettlement.TransactOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0xb86f602c.
//
// Solidity: function submitOrder(_prefix bytes, _settlementID uint64, _tokens uint64, _price uint256, _volume uint256, _minimumVolume uint256) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) SubmitOrder(_prefix []byte, _settlementID uint64, _tokens uint64, _price *big.Int, _volume *big.Int, _minimumVolume *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitOrder(&_RenExSettlement.TransactOpts, _prefix, _settlementID, _tokens, _price, _volume, _minimumVolume)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TransferOwnership(&_RenExSettlement.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TransferOwnership(&_RenExSettlement.TransactOpts, _newOwner)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateOrderbook(opts *bind.TransactOpts, _newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateOrderbook", _newOrderbookContract)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateOrderbook(_newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateOrderbook(&_RenExSettlement.TransactOpts, _newOrderbookContract)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateOrderbook(_newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateOrderbook(&_RenExSettlement.TransactOpts, _newOrderbookContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateRenExBalances(opts *bind.TransactOpts, _newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateRenExBalances", _newRenExBalancesContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateRenExBalances(_newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExBalances(&_RenExSettlement.TransactOpts, _newRenExBalancesContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateRenExBalances(_newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExBalances(&_RenExSettlement.TransactOpts, _newRenExBalancesContract)
}

// UpdateRenExTokens is a paid mutator transaction binding the contract method 0x34106c89.
//
// Solidity: function updateRenExTokens(_newRenExTokensContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateRenExTokens(opts *bind.TransactOpts, _newRenExTokensContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateRenExTokens", _newRenExTokensContract)
}

// UpdateRenExTokens is a paid mutator transaction binding the contract method 0x34106c89.
//
// Solidity: function updateRenExTokens(_newRenExTokensContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateRenExTokens(_newRenExTokensContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExTokens(&_RenExSettlement.TransactOpts, _newRenExTokensContract)
}

// UpdateRenExTokens is a paid mutator transaction binding the contract method 0x34106c89.
//
// Solidity: function updateRenExTokens(_newRenExTokensContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateRenExTokens(_newRenExTokensContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExTokens(&_RenExSettlement.TransactOpts, _newRenExTokensContract)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_newSlasherAddress address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateSlasher(opts *bind.TransactOpts, _newSlasherAddress common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateSlasher", _newSlasherAddress)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_newSlasherAddress address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateSlasher(_newSlasherAddress common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSlasher(&_RenExSettlement.TransactOpts, _newSlasherAddress)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_newSlasherAddress address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateSlasher(_newSlasherAddress common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSlasher(&_RenExSettlement.TransactOpts, _newSlasherAddress)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateSubmissionGasPriceLimit(opts *bind.TransactOpts, _newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateSubmissionGasPriceLimit", _newSubmissionGasPriceLimit)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSubmissionGasPriceLimit(&_RenExSettlement.TransactOpts, _newSubmissionGasPriceLimit)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSubmissionGasPriceLimit(&_RenExSettlement.TransactOpts, _newSubmissionGasPriceLimit)
}

// RenExSettlementLogOrderbookUpdatedIterator is returned from FilterLogOrderbookUpdated and is used to iterate over the raw logs and unpacked data for LogOrderbookUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogOrderbookUpdatedIterator struct {
	Event *RenExSettlementLogOrderbookUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogOrderbookUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogOrderbookUpdated)
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
		it.Event = new(RenExSettlementLogOrderbookUpdated)
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
func (it *RenExSettlementLogOrderbookUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogOrderbookUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogOrderbookUpdated represents a LogOrderbookUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogOrderbookUpdated struct {
	PreviousOrderbook common.Address
	NextOrderbook     common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogOrderbookUpdated is a free log retrieval operation binding the contract event 0xf7af59918b82b5e13957d357d0fcc86f12a806b0d2e826bc24a0f13ae85e4598.
//
// Solidity: e LogOrderbookUpdated(previousOrderbook address, nextOrderbook address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogOrderbookUpdated(opts *bind.FilterOpts) (*RenExSettlementLogOrderbookUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogOrderbookUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogOrderbookUpdatedIterator{contract: _RenExSettlement.contract, event: "LogOrderbookUpdated", logs: logs, sub: sub}, nil
}

// WatchLogOrderbookUpdated is a free log subscription operation binding the contract event 0xf7af59918b82b5e13957d357d0fcc86f12a806b0d2e826bc24a0f13ae85e4598.
//
// Solidity: e LogOrderbookUpdated(previousOrderbook address, nextOrderbook address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogOrderbookUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogOrderbookUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogOrderbookUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogOrderbookUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogOrderbookUpdated", log); err != nil {
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

// RenExSettlementLogRenExBalancesUpdatedIterator is returned from FilterLogRenExBalancesUpdated and is used to iterate over the raw logs and unpacked data for LogRenExBalancesUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogRenExBalancesUpdatedIterator struct {
	Event *RenExSettlementLogRenExBalancesUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogRenExBalancesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogRenExBalancesUpdated)
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
		it.Event = new(RenExSettlementLogRenExBalancesUpdated)
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
func (it *RenExSettlementLogRenExBalancesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogRenExBalancesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogRenExBalancesUpdated represents a LogRenExBalancesUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogRenExBalancesUpdated struct {
	PreviousRenExBalances common.Address
	NextRenExBalances     common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLogRenExBalancesUpdated is a free log retrieval operation binding the contract event 0x28e85eee30dd92456f8c6864bcdaadb36644672cee6f285d571b1e58c08adca1.
//
// Solidity: e LogRenExBalancesUpdated(previousRenExBalances address, nextRenExBalances address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogRenExBalancesUpdated(opts *bind.FilterOpts) (*RenExSettlementLogRenExBalancesUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogRenExBalancesUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogRenExBalancesUpdatedIterator{contract: _RenExSettlement.contract, event: "LogRenExBalancesUpdated", logs: logs, sub: sub}, nil
}

// WatchLogRenExBalancesUpdated is a free log subscription operation binding the contract event 0x28e85eee30dd92456f8c6864bcdaadb36644672cee6f285d571b1e58c08adca1.
//
// Solidity: e LogRenExBalancesUpdated(previousRenExBalances address, nextRenExBalances address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogRenExBalancesUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogRenExBalancesUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogRenExBalancesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogRenExBalancesUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogRenExBalancesUpdated", log); err != nil {
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

// RenExSettlementLogRenExTokensUpdatedIterator is returned from FilterLogRenExTokensUpdated and is used to iterate over the raw logs and unpacked data for LogRenExTokensUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogRenExTokensUpdatedIterator struct {
	Event *RenExSettlementLogRenExTokensUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogRenExTokensUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogRenExTokensUpdated)
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
		it.Event = new(RenExSettlementLogRenExTokensUpdated)
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
func (it *RenExSettlementLogRenExTokensUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogRenExTokensUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogRenExTokensUpdated represents a LogRenExTokensUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogRenExTokensUpdated struct {
	PreviousRenExTokens common.Address
	NextRenExTokens     common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogRenExTokensUpdated is a free log retrieval operation binding the contract event 0xc44a7f49dd4281e6c3ed47edb754b69b064653d53ed217e1354e79e8fe4b06a0.
//
// Solidity: e LogRenExTokensUpdated(previousRenExTokens address, nextRenExTokens address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogRenExTokensUpdated(opts *bind.FilterOpts) (*RenExSettlementLogRenExTokensUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogRenExTokensUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogRenExTokensUpdatedIterator{contract: _RenExSettlement.contract, event: "LogRenExTokensUpdated", logs: logs, sub: sub}, nil
}

// WatchLogRenExTokensUpdated is a free log subscription operation binding the contract event 0xc44a7f49dd4281e6c3ed47edb754b69b064653d53ed217e1354e79e8fe4b06a0.
//
// Solidity: e LogRenExTokensUpdated(previousRenExTokens address, nextRenExTokens address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogRenExTokensUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogRenExTokensUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogRenExTokensUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogRenExTokensUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogRenExTokensUpdated", log); err != nil {
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

// RenExSettlementLogSlasherUpdatedIterator is returned from FilterLogSlasherUpdated and is used to iterate over the raw logs and unpacked data for LogSlasherUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogSlasherUpdatedIterator struct {
	Event *RenExSettlementLogSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogSlasherUpdated)
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
		it.Event = new(RenExSettlementLogSlasherUpdated)
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
func (it *RenExSettlementLogSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogSlasherUpdated represents a LogSlasherUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogSlasherUpdated struct {
	PreviousSlasher common.Address
	NextSlasher     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSlasherUpdated is a free log retrieval operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: e LogSlasherUpdated(previousSlasher address, nextSlasher address)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogSlasherUpdated(opts *bind.FilterOpts) (*RenExSettlementLogSlasherUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogSlasherUpdatedIterator{contract: _RenExSettlement.contract, event: "LogSlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchLogSlasherUpdated is a free log subscription operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: e LogSlasherUpdated(previousSlasher address, nextSlasher address)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogSlasherUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogSlasherUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogSlasherUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogSlasherUpdated", log); err != nil {
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

// RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator is returned from FilterLogSubmissionGasPriceLimitUpdated and is used to iterate over the raw logs and unpacked data for LogSubmissionGasPriceLimitUpdated events raised by the RenExSettlement contract.
type RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator struct {
	Event *RenExSettlementLogSubmissionGasPriceLimitUpdated // Event containing the contract specifics and raw log

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
func (it *RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementLogSubmissionGasPriceLimitUpdated)
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
		it.Event = new(RenExSettlementLogSubmissionGasPriceLimitUpdated)
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
func (it *RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementLogSubmissionGasPriceLimitUpdated represents a LogSubmissionGasPriceLimitUpdated event raised by the RenExSettlement contract.
type RenExSettlementLogSubmissionGasPriceLimitUpdated struct {
	PreviousSubmissionGasPriceLimit *big.Int
	NextSubmissionGasPriceLimit     *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterLogSubmissionGasPriceLimitUpdated is a free log retrieval operation binding the contract event 0xd0ef246766073915a6813492cff2a021d7cd4bf7d4feff3ef74327c7f4940e96.
//
// Solidity: e LogSubmissionGasPriceLimitUpdated(previousSubmissionGasPriceLimit uint256, nextSubmissionGasPriceLimit uint256)
func (_RenExSettlement *RenExSettlementFilterer) FilterLogSubmissionGasPriceLimitUpdated(opts *bind.FilterOpts) (*RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "LogSubmissionGasPriceLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementLogSubmissionGasPriceLimitUpdatedIterator{contract: _RenExSettlement.contract, event: "LogSubmissionGasPriceLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchLogSubmissionGasPriceLimitUpdated is a free log subscription operation binding the contract event 0xd0ef246766073915a6813492cff2a021d7cd4bf7d4feff3ef74327c7f4940e96.
//
// Solidity: e LogSubmissionGasPriceLimitUpdated(previousSubmissionGasPriceLimit uint256, nextSubmissionGasPriceLimit uint256)
func (_RenExSettlement *RenExSettlementFilterer) WatchLogSubmissionGasPriceLimitUpdated(opts *bind.WatchOpts, sink chan<- *RenExSettlementLogSubmissionGasPriceLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "LogSubmissionGasPriceLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementLogSubmissionGasPriceLimitUpdated)
				if err := _RenExSettlement.contract.UnpackLog(event, "LogSubmissionGasPriceLimitUpdated", log); err != nil {
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

// RenExSettlementOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExSettlement contract.
type RenExSettlementOwnershipRenouncedIterator struct {
	Event *RenExSettlementOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExSettlementOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementOwnershipRenounced)
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
		it.Event = new(RenExSettlementOwnershipRenounced)
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
func (it *RenExSettlementOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementOwnershipRenounced represents a OwnershipRenounced event raised by the RenExSettlement contract.
type RenExSettlementOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExSettlementOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementOwnershipRenouncedIterator{contract: _RenExSettlement.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExSettlementOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementOwnershipRenounced)
				if err := _RenExSettlement.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExSettlementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExSettlement contract.
type RenExSettlementOwnershipTransferredIterator struct {
	Event *RenExSettlementOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExSettlementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementOwnershipTransferred)
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
		it.Event = new(RenExSettlementOwnershipTransferred)
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
func (it *RenExSettlementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementOwnershipTransferred represents a OwnershipTransferred event raised by the RenExSettlement contract.
type RenExSettlementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExSettlementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementOwnershipTransferredIterator{contract: _RenExSettlement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExSettlementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementOwnershipTransferred)
				if err := _RenExSettlement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExTokensABI is the input ABI used to generate the binding from.
const RenExTokensABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenCode\",\"type\":\"uint32\"}],\"name\":\"deregisterToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenCode\",\"type\":\"uint32\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"registerToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"registered\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"LogTokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenCode\",\"type\":\"uint32\"}],\"name\":\"LogTokenDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExTokensBin is the compiled bytecode used for deploying new contracts.
const RenExTokensBin = `0x608060405260008054600160a060020a03191633179055610724806100256000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166325fc575a811461007c578063715018a61461009c5780638da5cb5b146100b15780639e20a9a0146100e2578063f2fde38b14610112578063fbb6272d14610133575b600080fd5b34801561008857600080fd5b5061009a63ffffffff6004351661017d565b005b3480156100a857600080fd5b5061009a610286565b3480156100bd57600080fd5b506100c66102f2565b60408051600160a060020a039092168252519081900360200190f35b3480156100ee57600080fd5b5061009a63ffffffff60043516600160a060020a036024351660ff60443516610301565b34801561011e57600080fd5b5061009a600160a060020a0360043516610604565b34801561013f57600080fd5b5061015163ffffffff60043516610627565b60408051600160a060020a03909416845260ff9092166020840152151582820152519081900360600190f35b600054600160a060020a0316331461019457600080fd5b63ffffffff81166000908152600160205260409020547501000000000000000000000000000000000000000000900460ff16151561021c576040805160e560020a62461bcd02815260206004820152600e60248201527f6e6f742072656769737465726564000000000000000000000000000000000000604482015290519081900360640190fd5b63ffffffff8116600081815260016020908152604091829020805475ff00000000000000000000000000000000000000000019169055815192835290517fa521f8871580a5e5b43fa9c01cd9c3a22f4be51fb276e322656db351dbdef3209281900390910190a150565b600054600160a060020a0316331461029d57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461031857600080fd5b63ffffffff83166000908152600160205260409020547501000000000000000000000000000000000000000000900460ff161561039f576040805160e560020a62461bcd02815260206004820152601260248201527f616c726561647920726567697374657265640000000000000000000000000000604482015290519081900360640190fd5b63ffffffff831660009081526002602052604090205460ff16156104c45763ffffffff8316600090815260016020526040902054600160a060020a03838116911614610435576040805160e560020a62461bcd02815260206004820152601160248201527f646966666572656e742061646472657373000000000000000000000000000000604482015290519081900360640190fd5b63ffffffff831660009081526001602052604090205460ff8281167401000000000000000000000000000000000000000090920416146104bf576040805160e560020a62461bcd02815260206004820152601260248201527f646966666572656e7420646563696d616c730000000000000000000000000000604482015290519081900360640190fd5b6104e5565b63ffffffff83166000908152600260205260409020805460ff191660011790555b6040805160608181018352600160a060020a0380861680845260ff80871660208087018281526001888a0181815263ffffffff8e166000818152928552918b902099518a5493519151151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001992909716740100000000000000000000000000000000000000000274ff0000000000000000000000000000000000000000199190991673ffffffffffffffffffffffffffffffffffffffff19909416939093179290921696909617169290921790955585519283528201528084019290925291517f3b2dff39857f1432a4160494b13a3d9c1bd8781453d9d6f56153b69d89792a35929181900390910190a1505050565b600054600160a060020a0316331461061b57600080fd5b6106248161067b565b50565b600160205260009081526040902054600160a060020a0381169060ff740100000000000000000000000000000000000000008204811691750100000000000000000000000000000000000000000090041683565b600160a060020a038116151561069057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582050a64a2194b18f945a36bb167d673265fc3b050a8a98d291de48dfb1c5bac5720029`

// DeployRenExTokens deploys a new Ethereum contract, binding an instance of RenExTokens to it.
func DeployRenExTokens(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RenExTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExTokensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExTokensBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExTokens{RenExTokensCaller: RenExTokensCaller{contract: contract}, RenExTokensTransactor: RenExTokensTransactor{contract: contract}, RenExTokensFilterer: RenExTokensFilterer{contract: contract}}, nil
}

// RenExTokens is an auto generated Go binding around an Ethereum contract.
type RenExTokens struct {
	RenExTokensCaller     // Read-only binding to the contract
	RenExTokensTransactor // Write-only binding to the contract
	RenExTokensFilterer   // Log filterer for contract events
}

// RenExTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExTokensSession struct {
	Contract     *RenExTokens      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExTokensCallerSession struct {
	Contract *RenExTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RenExTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExTokensTransactorSession struct {
	Contract     *RenExTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RenExTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExTokensRaw struct {
	Contract *RenExTokens // Generic contract binding to access the raw methods on
}

// RenExTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExTokensCallerRaw struct {
	Contract *RenExTokensCaller // Generic read-only contract binding to access the raw methods on
}

// RenExTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExTokensTransactorRaw struct {
	Contract *RenExTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExTokens creates a new instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokens(address common.Address, backend bind.ContractBackend) (*RenExTokens, error) {
	contract, err := bindRenExTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExTokens{RenExTokensCaller: RenExTokensCaller{contract: contract}, RenExTokensTransactor: RenExTokensTransactor{contract: contract}, RenExTokensFilterer: RenExTokensFilterer{contract: contract}}, nil
}

// NewRenExTokensCaller creates a new read-only instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensCaller(address common.Address, caller bind.ContractCaller) (*RenExTokensCaller, error) {
	contract, err := bindRenExTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExTokensCaller{contract: contract}, nil
}

// NewRenExTokensTransactor creates a new write-only instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExTokensTransactor, error) {
	contract, err := bindRenExTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExTokensTransactor{contract: contract}, nil
}

// NewRenExTokensFilterer creates a new log filterer instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExTokensFilterer, error) {
	contract, err := bindRenExTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExTokensFilterer{contract: contract}, nil
}

// bindRenExTokens binds a generic wrapper to an already deployed contract.
func bindRenExTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExTokens *RenExTokensRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExTokens.Contract.RenExTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExTokens *RenExTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.Contract.RenExTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExTokens *RenExTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExTokens.Contract.RenExTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExTokens *RenExTokensCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExTokens *RenExTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExTokens *RenExTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExTokens.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExTokens.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensSession) Owner() (common.Address, error) {
	return _RenExTokens.Contract.Owner(&_RenExTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensCallerSession) Owner() (common.Address, error) {
	return _RenExTokens.Contract.Owner(&_RenExTokens.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xfbb6272d.
//
// Solidity: function tokens( uint32) constant returns(addr address, decimals uint8, registered bool)
func (_RenExTokens *RenExTokensCaller) Tokens(opts *bind.CallOpts, arg0 uint32) (struct {
	Addr       common.Address
	Decimals   uint8
	Registered bool
}, error) {
	ret := new(struct {
		Addr       common.Address
		Decimals   uint8
		Registered bool
	})
	out := ret
	err := _RenExTokens.contract.Call(opts, out, "tokens", arg0)
	return *ret, err
}

// Tokens is a free data retrieval call binding the contract method 0xfbb6272d.
//
// Solidity: function tokens( uint32) constant returns(addr address, decimals uint8, registered bool)
func (_RenExTokens *RenExTokensSession) Tokens(arg0 uint32) (struct {
	Addr       common.Address
	Decimals   uint8
	Registered bool
}, error) {
	return _RenExTokens.Contract.Tokens(&_RenExTokens.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xfbb6272d.
//
// Solidity: function tokens( uint32) constant returns(addr address, decimals uint8, registered bool)
func (_RenExTokens *RenExTokensCallerSession) Tokens(arg0 uint32) (struct {
	Addr       common.Address
	Decimals   uint8
	Registered bool
}, error) {
	return _RenExTokens.Contract.Tokens(&_RenExTokens.CallOpts, arg0)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensTransactor) DeregisterToken(opts *bind.TransactOpts, _tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "deregisterToken", _tokenCode)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensSession) DeregisterToken(_tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.Contract.DeregisterToken(&_RenExTokens.TransactOpts, _tokenCode)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensTransactorSession) DeregisterToken(_tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.Contract.DeregisterToken(&_RenExTokens.TransactOpts, _tokenCode)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensTransactor) RegisterToken(opts *bind.TransactOpts, _tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "registerToken", _tokenCode, _tokenAddress, _tokenDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensSession) RegisterToken(_tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.Contract.RegisterToken(&_RenExTokens.TransactOpts, _tokenCode, _tokenAddress, _tokenDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensTransactorSession) RegisterToken(_tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.Contract.RegisterToken(&_RenExTokens.TransactOpts, _tokenCode, _tokenAddress, _tokenDecimals)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExTokens.Contract.RenounceOwnership(&_RenExTokens.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExTokens.Contract.RenounceOwnership(&_RenExTokens.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.Contract.TransferOwnership(&_RenExTokens.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.Contract.TransferOwnership(&_RenExTokens.TransactOpts, _newOwner)
}

// RenExTokensLogTokenDeregisteredIterator is returned from FilterLogTokenDeregistered and is used to iterate over the raw logs and unpacked data for LogTokenDeregistered events raised by the RenExTokens contract.
type RenExTokensLogTokenDeregisteredIterator struct {
	Event *RenExTokensLogTokenDeregistered // Event containing the contract specifics and raw log

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
func (it *RenExTokensLogTokenDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensLogTokenDeregistered)
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
		it.Event = new(RenExTokensLogTokenDeregistered)
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
func (it *RenExTokensLogTokenDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensLogTokenDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensLogTokenDeregistered represents a LogTokenDeregistered event raised by the RenExTokens contract.
type RenExTokensLogTokenDeregistered struct {
	TokenCode uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTokenDeregistered is a free log retrieval operation binding the contract event 0xa521f8871580a5e5b43fa9c01cd9c3a22f4be51fb276e322656db351dbdef320.
//
// Solidity: e LogTokenDeregistered(tokenCode uint32)
func (_RenExTokens *RenExTokensFilterer) FilterLogTokenDeregistered(opts *bind.FilterOpts) (*RenExTokensLogTokenDeregisteredIterator, error) {

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "LogTokenDeregistered")
	if err != nil {
		return nil, err
	}
	return &RenExTokensLogTokenDeregisteredIterator{contract: _RenExTokens.contract, event: "LogTokenDeregistered", logs: logs, sub: sub}, nil
}

// WatchLogTokenDeregistered is a free log subscription operation binding the contract event 0xa521f8871580a5e5b43fa9c01cd9c3a22f4be51fb276e322656db351dbdef320.
//
// Solidity: e LogTokenDeregistered(tokenCode uint32)
func (_RenExTokens *RenExTokensFilterer) WatchLogTokenDeregistered(opts *bind.WatchOpts, sink chan<- *RenExTokensLogTokenDeregistered) (event.Subscription, error) {

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "LogTokenDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensLogTokenDeregistered)
				if err := _RenExTokens.contract.UnpackLog(event, "LogTokenDeregistered", log); err != nil {
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

// RenExTokensLogTokenRegisteredIterator is returned from FilterLogTokenRegistered and is used to iterate over the raw logs and unpacked data for LogTokenRegistered events raised by the RenExTokens contract.
type RenExTokensLogTokenRegisteredIterator struct {
	Event *RenExTokensLogTokenRegistered // Event containing the contract specifics and raw log

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
func (it *RenExTokensLogTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensLogTokenRegistered)
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
		it.Event = new(RenExTokensLogTokenRegistered)
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
func (it *RenExTokensLogTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensLogTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensLogTokenRegistered represents a LogTokenRegistered event raised by the RenExTokens contract.
type RenExTokensLogTokenRegistered struct {
	TokenCode     uint32
	TokenAddress  common.Address
	TokenDecimals uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogTokenRegistered is a free log retrieval operation binding the contract event 0x3b2dff39857f1432a4160494b13a3d9c1bd8781453d9d6f56153b69d89792a35.
//
// Solidity: e LogTokenRegistered(tokenCode uint32, tokenAddress address, tokenDecimals uint8)
func (_RenExTokens *RenExTokensFilterer) FilterLogTokenRegistered(opts *bind.FilterOpts) (*RenExTokensLogTokenRegisteredIterator, error) {

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return &RenExTokensLogTokenRegisteredIterator{contract: _RenExTokens.contract, event: "LogTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchLogTokenRegistered is a free log subscription operation binding the contract event 0x3b2dff39857f1432a4160494b13a3d9c1bd8781453d9d6f56153b69d89792a35.
//
// Solidity: e LogTokenRegistered(tokenCode uint32, tokenAddress address, tokenDecimals uint8)
func (_RenExTokens *RenExTokensFilterer) WatchLogTokenRegistered(opts *bind.WatchOpts, sink chan<- *RenExTokensLogTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensLogTokenRegistered)
				if err := _RenExTokens.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
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

// RenExTokensOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExTokens contract.
type RenExTokensOwnershipRenouncedIterator struct {
	Event *RenExTokensOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExTokensOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensOwnershipRenounced)
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
		it.Event = new(RenExTokensOwnershipRenounced)
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
func (it *RenExTokensOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensOwnershipRenounced represents a OwnershipRenounced event raised by the RenExTokens contract.
type RenExTokensOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExTokensOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExTokensOwnershipRenouncedIterator{contract: _RenExTokens.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExTokensOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensOwnershipRenounced)
				if err := _RenExTokens.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExTokensOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExTokens contract.
type RenExTokensOwnershipTransferredIterator struct {
	Event *RenExTokensOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExTokensOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensOwnershipTransferred)
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
		it.Event = new(RenExTokensOwnershipTransferred)
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
func (it *RenExTokensOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensOwnershipTransferred represents a OwnershipTransferred event raised by the RenExTokens contract.
type RenExTokensOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExTokensOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExTokensOwnershipTransferredIterator{contract: _RenExTokens.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExTokensOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensOwnershipTransferred)
				if err := _RenExTokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RepublicTokenABI is the input ABI used to generate the binding from.
const RepublicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// RepublicTokenBin is the compiled bytecode used for deploying new contracts.
const RepublicTokenBin = `0x60806040526003805460a060020a60ff021916905534801561002057600080fd5b5060038054600160a060020a031916339081179091556b033b2e3c9fd0803ce8000000600181905560009182526020829052604090912055610ea6806100676000396000f3006080604052600436106101115763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610116578063095ea7b3146101a057806318160ddd146101d857806323b872dd146101ff5780632ff2e9dc14610229578063313ce5671461023e5780633f4ba83a1461026957806342966c68146102805780635c975abb1461029857806366188463146102ad57806370a08231146102d1578063715018a6146102f25780638456cb59146103075780638da5cb5b1461031c57806395d89b411461034d578063a9059cbb14610362578063bec3fa1714610386578063d73dd623146103aa578063dd62ed3e146103ce578063f2fde38b146103f5575b600080fd5b34801561012257600080fd5b5061012b610416565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016557818101518382015260200161014d565b50505050905090810190601f1680156101925780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ac57600080fd5b506101c4600160a060020a036004351660243561044d565b604080519115158252519081900360200190f35b3480156101e457600080fd5b506101ed610478565b60408051918252519081900360200190f35b34801561020b57600080fd5b506101c4600160a060020a036004358116906024351660443561047e565b34801561023557600080fd5b506101ed6104c1565b34801561024a57600080fd5b506102536104d1565b6040805160ff9092168252519081900360200190f35b34801561027557600080fd5b5061027e6104d6565b005b34801561028c57600080fd5b5061027e60043561054e565b3480156102a457600080fd5b506101c461055b565b3480156102b957600080fd5b506101c4600160a060020a036004351660243561056b565b3480156102dd57600080fd5b506101ed600160a060020a036004351661058f565b3480156102fe57600080fd5b5061027e6105aa565b34801561031357600080fd5b5061027e610618565b34801561032857600080fd5b50610331610695565b60408051600160a060020a039092168252519081900360200190f35b34801561035957600080fd5b5061012b6106a4565b34801561036e57600080fd5b506101c4600160a060020a03600435166024356106db565b34801561039257600080fd5b506101c4600160a060020a0360043516602435610715565b3480156103b657600080fd5b506101c4600160a060020a03600435166024356107ed565b3480156103da57600080fd5b506101ed600160a060020a0360043581169060243516610811565b34801561040157600080fd5b5061027e600160a060020a036004351661083c565b60408051808201909152600e81527f52657075626c696320546f6b656e000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff161561046757600080fd5b610471838361085c565b9392505050565b60015490565b60035460009060a060020a900460ff161561049857600080fd5b600160a060020a0383163014156104ae57600080fd5b6104b98484846108c2565b949350505050565b6b033b2e3c9fd0803ce800000081565b601281565b600354600160a060020a031633146104ed57600080fd5b60035460a060020a900460ff16151561050557600080fd5b6003805474ff0000000000000000000000000000000000000000191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b61055833826108e7565b50565b60035460a060020a900460ff1681565b60035460009060a060020a900460ff161561058557600080fd5b61047183836109d6565b600160a060020a031660009081526020819052604090205490565b600354600160a060020a031633146105c157600080fd5b600354604051600160a060020a03909116907ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482090600090a26003805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354600160a060020a0316331461062f57600080fd5b60035460a060020a900460ff161561064657600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600354600160a060020a031681565b60408051808201909152600381527f52454e0000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156106f557600080fd5b600160a060020a03831630141561070b57600080fd5b6104718383610ac6565b600354600090600160a060020a0316331461072f57600080fd5b6000821161073c57600080fd5b600354600160a060020a0316600090815260208190526040902054610767908363ffffffff610aea16565b600354600160a060020a03908116600090815260208190526040808220939093559085168152205461079f908363ffffffff610afc16565b600160a060020a038085166000818152602081815260409182902094909455600354815187815291519294931692600080516020610e5b83398151915292918290030190a350600192915050565b60035460009060a060020a900460ff161561080757600080fd5b6104718383610b0f565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a0316331461085357600080fd5b61055881610ba8565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035460009060a060020a900460ff16156108dc57600080fd5b6104b9848484610c26565b600160a060020a03821660009081526020819052604090205481111561090c57600080fd5b600160a060020a038216600090815260208190526040902054610935908263ffffffff610aea16565b600160a060020a038316600090815260208190526040902055600154610961908263ffffffff610aea16565b600155604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a03851691600080516020610e5b8339815191529181900360200190a35050565b336000908152600260209081526040808320600160a060020a038616845290915281205480831115610a2b57336000908152600260209081526040808320600160a060020a0388168452909152812055610a60565b610a3b818463ffffffff610aea16565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60035460009060a060020a900460ff1615610ae057600080fd5b6104718383610d8b565b600082821115610af657fe5b50900390565b81810182811015610b0957fe5b92915050565b336000908152600260209081526040808320600160a060020a0386168452909152812054610b43908363ffffffff610afc16565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a0381161515610bbd57600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a0383161515610c3d57600080fd5b600160a060020a038416600090815260208190526040902054821115610c6257600080fd5b600160a060020a0384166000908152600260209081526040808320338452909152902054821115610c9257600080fd5b600160a060020a038416600090815260208190526040902054610cbb908363ffffffff610aea16565b600160a060020a038086166000908152602081905260408082209390935590851681522054610cf0908363ffffffff610afc16565b600160a060020a03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610d32908363ffffffff610aea16565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610e5b833981519152929181900390910190a35060019392505050565b6000600160a060020a0383161515610da257600080fd5b33600090815260208190526040902054821115610dbe57600080fd5b33600090815260208190526040902054610dde908363ffffffff610aea16565b3360009081526020819052604080822092909255600160a060020a03851681522054610e10908363ffffffff610afc16565b600160a060020a03841660008181526020818152604091829020939093558051858152905191923392600080516020610e5b8339815191529281900390910190a3506001929150505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a72305820fdd541e4eaa35cb26285bd43c8ea0d5a24cf6a1d6203a28ffc4354e886fb0de40029`

// DeployRepublicToken deploys a new Ethereum contract, binding an instance of RepublicToken to it.
func DeployRepublicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RepublicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(RepublicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RepublicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RepublicToken{RepublicTokenCaller: RepublicTokenCaller{contract: contract}, RepublicTokenTransactor: RepublicTokenTransactor{contract: contract}, RepublicTokenFilterer: RepublicTokenFilterer{contract: contract}}, nil
}

// RepublicToken is an auto generated Go binding around an Ethereum contract.
type RepublicToken struct {
	RepublicTokenCaller     // Read-only binding to the contract
	RepublicTokenTransactor // Write-only binding to the contract
	RepublicTokenFilterer   // Log filterer for contract events
}

// RepublicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RepublicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepublicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RepublicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepublicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RepublicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepublicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RepublicTokenSession struct {
	Contract     *RepublicToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepublicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RepublicTokenCallerSession struct {
	Contract *RepublicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RepublicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RepublicTokenTransactorSession struct {
	Contract     *RepublicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RepublicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RepublicTokenRaw struct {
	Contract *RepublicToken // Generic contract binding to access the raw methods on
}

// RepublicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RepublicTokenCallerRaw struct {
	Contract *RepublicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RepublicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RepublicTokenTransactorRaw struct {
	Contract *RepublicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRepublicToken creates a new instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicToken(address common.Address, backend bind.ContractBackend) (*RepublicToken, error) {
	contract, err := bindRepublicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RepublicToken{RepublicTokenCaller: RepublicTokenCaller{contract: contract}, RepublicTokenTransactor: RepublicTokenTransactor{contract: contract}, RepublicTokenFilterer: RepublicTokenFilterer{contract: contract}}, nil
}

// NewRepublicTokenCaller creates a new read-only instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicTokenCaller(address common.Address, caller bind.ContractCaller) (*RepublicTokenCaller, error) {
	contract, err := bindRepublicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenCaller{contract: contract}, nil
}

// NewRepublicTokenTransactor creates a new write-only instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RepublicTokenTransactor, error) {
	contract, err := bindRepublicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenTransactor{contract: contract}, nil
}

// NewRepublicTokenFilterer creates a new log filterer instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RepublicTokenFilterer, error) {
	contract, err := bindRepublicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenFilterer{contract: contract}, nil
}

// bindRepublicToken binds a generic wrapper to an already deployed contract.
func bindRepublicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RepublicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepublicToken *RepublicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepublicToken.Contract.RepublicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepublicToken *RepublicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.Contract.RepublicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepublicToken *RepublicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepublicToken.Contract.RepublicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepublicToken *RepublicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepublicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepublicToken *RepublicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepublicToken *RepublicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepublicToken.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _RepublicToken.Contract.INITIALSUPPLY(&_RepublicToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _RepublicToken.Contract.INITIALSUPPLY(&_RepublicToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.Allowance(&_RepublicToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.Allowance(&_RepublicToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.BalanceOf(&_RepublicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.BalanceOf(&_RepublicToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RepublicToken *RepublicTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RepublicToken *RepublicTokenSession) Decimals() (uint8, error) {
	return _RepublicToken.Contract.Decimals(&_RepublicToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RepublicToken *RepublicTokenCallerSession) Decimals() (uint8, error) {
	return _RepublicToken.Contract.Decimals(&_RepublicToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepublicToken *RepublicTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepublicToken *RepublicTokenSession) Name() (string, error) {
	return _RepublicToken.Contract.Name(&_RepublicToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepublicToken *RepublicTokenCallerSession) Name() (string, error) {
	return _RepublicToken.Contract.Name(&_RepublicToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RepublicToken *RepublicTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RepublicToken *RepublicTokenSession) Owner() (common.Address, error) {
	return _RepublicToken.Contract.Owner(&_RepublicToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RepublicToken *RepublicTokenCallerSession) Owner() (common.Address, error) {
	return _RepublicToken.Contract.Owner(&_RepublicToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepublicToken *RepublicTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepublicToken *RepublicTokenSession) Paused() (bool, error) {
	return _RepublicToken.Contract.Paused(&_RepublicToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepublicToken *RepublicTokenCallerSession) Paused() (bool, error) {
	return _RepublicToken.Contract.Paused(&_RepublicToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepublicToken *RepublicTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepublicToken *RepublicTokenSession) Symbol() (string, error) {
	return _RepublicToken.Contract.Symbol(&_RepublicToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepublicToken *RepublicTokenCallerSession) Symbol() (string, error) {
	return _RepublicToken.Contract.Symbol(&_RepublicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) TotalSupply() (*big.Int, error) {
	return _RepublicToken.Contract.TotalSupply(&_RepublicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RepublicToken.Contract.TotalSupply(&_RepublicToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Approve(&_RepublicToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Approve(&_RepublicToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_RepublicToken *RepublicTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_RepublicToken *RepublicTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Burn(&_RepublicToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_RepublicToken *RepublicTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Burn(&_RepublicToken.TransactOpts, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.DecreaseApproval(&_RepublicToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.DecreaseApproval(&_RepublicToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.IncreaseApproval(&_RepublicToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.IncreaseApproval(&_RepublicToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepublicToken *RepublicTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepublicToken *RepublicTokenSession) Pause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Pause(&_RepublicToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepublicToken *RepublicTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Pause(&_RepublicToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RepublicToken *RepublicTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RepublicToken *RepublicTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _RepublicToken.Contract.RenounceOwnership(&_RepublicToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RepublicToken *RepublicTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RepublicToken.Contract.RenounceOwnership(&_RepublicToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Transfer(&_RepublicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Transfer(&_RepublicToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferFrom(&_RepublicToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferFrom(&_RepublicToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RepublicToken *RepublicTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RepublicToken *RepublicTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferOwnership(&_RepublicToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RepublicToken *RepublicTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferOwnership(&_RepublicToken.TransactOpts, _newOwner)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(beneficiary address, amount uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) TransferTokens(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transferTokens", beneficiary, amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(beneficiary address, amount uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) TransferTokens(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferTokens(&_RepublicToken.TransactOpts, beneficiary, amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(beneficiary address, amount uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) TransferTokens(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferTokens(&_RepublicToken.TransactOpts, beneficiary, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepublicToken *RepublicTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepublicToken *RepublicTokenSession) Unpause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Unpause(&_RepublicToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepublicToken *RepublicTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Unpause(&_RepublicToken.TransactOpts)
}

// RepublicTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RepublicToken contract.
type RepublicTokenApprovalIterator struct {
	Event *RepublicTokenApproval // Event containing the contract specifics and raw log

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
func (it *RepublicTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenApproval)
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
		it.Event = new(RepublicTokenApproval)
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
func (it *RepublicTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenApproval represents a Approval event raised by the RepublicToken contract.
type RepublicTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RepublicTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenApprovalIterator{contract: _RepublicToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RepublicTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenApproval)
				if err := _RepublicToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// RepublicTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the RepublicToken contract.
type RepublicTokenBurnIterator struct {
	Event *RepublicTokenBurn // Event containing the contract specifics and raw log

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
func (it *RepublicTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenBurn)
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
		it.Event = new(RepublicTokenBurn)
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
func (it *RepublicTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenBurn represents a Burn event raised by the RepublicToken contract.
type RepublicTokenBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*RepublicTokenBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenBurnIterator{contract: _RepublicToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *RepublicTokenBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenBurn)
				if err := _RepublicToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// RepublicTokenOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RepublicToken contract.
type RepublicTokenOwnershipRenouncedIterator struct {
	Event *RepublicTokenOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RepublicTokenOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenOwnershipRenounced)
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
		it.Event = new(RepublicTokenOwnershipRenounced)
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
func (it *RepublicTokenOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenOwnershipRenounced represents a OwnershipRenounced event raised by the RepublicToken contract.
type RepublicTokenOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RepublicTokenOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenOwnershipRenouncedIterator{contract: _RepublicToken.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RepublicTokenOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenOwnershipRenounced)
				if err := _RepublicToken.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RepublicTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RepublicToken contract.
type RepublicTokenOwnershipTransferredIterator struct {
	Event *RepublicTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RepublicTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenOwnershipTransferred)
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
		it.Event = new(RepublicTokenOwnershipTransferred)
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
func (it *RepublicTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenOwnershipTransferred represents a OwnershipTransferred event raised by the RepublicToken contract.
type RepublicTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RepublicTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenOwnershipTransferredIterator{contract: _RepublicToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RepublicTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenOwnershipTransferred)
				if err := _RepublicToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RepublicTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the RepublicToken contract.
type RepublicTokenPauseIterator struct {
	Event *RepublicTokenPause // Event containing the contract specifics and raw log

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
func (it *RepublicTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenPause)
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
		it.Event = new(RepublicTokenPause)
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
func (it *RepublicTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenPause represents a Pause event raised by the RepublicToken contract.
type RepublicTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_RepublicToken *RepublicTokenFilterer) FilterPause(opts *bind.FilterOpts) (*RepublicTokenPauseIterator, error) {

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &RepublicTokenPauseIterator{contract: _RepublicToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_RepublicToken *RepublicTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *RepublicTokenPause) (event.Subscription, error) {

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenPause)
				if err := _RepublicToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// RepublicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RepublicToken contract.
type RepublicTokenTransferIterator struct {
	Event *RepublicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RepublicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenTransfer)
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
		it.Event = new(RepublicTokenTransfer)
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
func (it *RepublicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenTransfer represents a Transfer event raised by the RepublicToken contract.
type RepublicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RepublicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenTransferIterator{contract: _RepublicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RepublicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenTransfer)
				if err := _RepublicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RepublicTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the RepublicToken contract.
type RepublicTokenUnpauseIterator struct {
	Event *RepublicTokenUnpause // Event containing the contract specifics and raw log

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
func (it *RepublicTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenUnpause)
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
		it.Event = new(RepublicTokenUnpause)
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
func (it *RepublicTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenUnpause represents a Unpause event raised by the RepublicToken contract.
type RepublicTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_RepublicToken *RepublicTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*RepublicTokenUnpauseIterator, error) {

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &RepublicTokenUnpauseIterator{contract: _RepublicToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_RepublicToken *RepublicTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *RepublicTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenUnpause)
				if err := _RepublicToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a02a84ac4f3af3b8d9db41a87667de0285507cc2e99581d2d45fbe62a74f203d0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SettlementUtilsABI is the input ABI used to generate the binding from.
const SettlementUtilsABI = "[]"

// SettlementUtilsBin is the compiled bytecode used for deploying new contracts.
const SettlementUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820069fc84ac797d4274be86c2a908788e32c7cebf4dcaab4ae096e021e42818d010029`

// DeploySettlementUtils deploys a new Ethereum contract, binding an instance of SettlementUtils to it.
func DeploySettlementUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SettlementUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SettlementUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SettlementUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SettlementUtils{SettlementUtilsCaller: SettlementUtilsCaller{contract: contract}, SettlementUtilsTransactor: SettlementUtilsTransactor{contract: contract}, SettlementUtilsFilterer: SettlementUtilsFilterer{contract: contract}}, nil
}

// SettlementUtils is an auto generated Go binding around an Ethereum contract.
type SettlementUtils struct {
	SettlementUtilsCaller     // Read-only binding to the contract
	SettlementUtilsTransactor // Write-only binding to the contract
	SettlementUtilsFilterer   // Log filterer for contract events
}

// SettlementUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SettlementUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SettlementUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SettlementUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SettlementUtilsSession struct {
	Contract     *SettlementUtils  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SettlementUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SettlementUtilsCallerSession struct {
	Contract *SettlementUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SettlementUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SettlementUtilsTransactorSession struct {
	Contract     *SettlementUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SettlementUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SettlementUtilsRaw struct {
	Contract *SettlementUtils // Generic contract binding to access the raw methods on
}

// SettlementUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SettlementUtilsCallerRaw struct {
	Contract *SettlementUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SettlementUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SettlementUtilsTransactorRaw struct {
	Contract *SettlementUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSettlementUtils creates a new instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtils(address common.Address, backend bind.ContractBackend) (*SettlementUtils, error) {
	contract, err := bindSettlementUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SettlementUtils{SettlementUtilsCaller: SettlementUtilsCaller{contract: contract}, SettlementUtilsTransactor: SettlementUtilsTransactor{contract: contract}, SettlementUtilsFilterer: SettlementUtilsFilterer{contract: contract}}, nil
}

// NewSettlementUtilsCaller creates a new read-only instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtilsCaller(address common.Address, caller bind.ContractCaller) (*SettlementUtilsCaller, error) {
	contract, err := bindSettlementUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementUtilsCaller{contract: contract}, nil
}

// NewSettlementUtilsTransactor creates a new write-only instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SettlementUtilsTransactor, error) {
	contract, err := bindSettlementUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementUtilsTransactor{contract: contract}, nil
}

// NewSettlementUtilsFilterer creates a new log filterer instance of SettlementUtils, bound to a specific deployed contract.
func NewSettlementUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SettlementUtilsFilterer, error) {
	contract, err := bindSettlementUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SettlementUtilsFilterer{contract: contract}, nil
}

// bindSettlementUtils binds a generic wrapper to an already deployed contract.
func bindSettlementUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SettlementUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SettlementUtils *SettlementUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SettlementUtils.Contract.SettlementUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SettlementUtils *SettlementUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SettlementUtils.Contract.SettlementUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SettlementUtils *SettlementUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SettlementUtils.Contract.SettlementUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SettlementUtils *SettlementUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SettlementUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SettlementUtils *SettlementUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SettlementUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SettlementUtils *SettlementUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SettlementUtils.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x608060405234801561001057600080fd5b506106b3806100206000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100ca57806323b872dd146100f1578063661884631461011b57806370a082311461013f578063a9059cbb14610160578063d73dd62314610184578063dd62ed3e146101a8575b600080fd5b34801561009e57600080fd5b506100b6600160a060020a03600435166024356101cf565b604080519115158252519081900360200190f35b3480156100d657600080fd5b506100df610235565b60408051918252519081900360200190f35b3480156100fd57600080fd5b506100b6600160a060020a036004358116906024351660443561023b565b34801561012757600080fd5b506100b6600160a060020a03600435166024356103b2565b34801561014b57600080fd5b506100df600160a060020a03600435166104a2565b34801561016c57600080fd5b506100b6600160a060020a03600435166024356104bd565b34801561019057600080fd5b506100b6600160a060020a036004351660243561059e565b3480156101b457600080fd5b506100df600160a060020a0360043581169060243516610637565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b6000600160a060020a038316151561025257600080fd5b600160a060020a03841660009081526020819052604090205482111561027757600080fd5b600160a060020a03841660009081526002602090815260408083203384529091529020548211156102a757600080fd5b600160a060020a0384166000908152602081905260409020546102d0908363ffffffff61066216565b600160a060020a038086166000908152602081905260408082209390935590851681522054610305908363ffffffff61067416565b600160a060020a03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610347908363ffffffff61066216565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b336000908152600260209081526040808320600160a060020a03861684529091528120548083111561040757336000908152600260209081526040808320600160a060020a038816845290915281205561043c565b610417818463ffffffff61066216565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a03831615156104d457600080fd5b336000908152602081905260409020548211156104f057600080fd5b33600090815260208190526040902054610510908363ffffffff61066216565b3360009081526020819052604080822092909255600160a060020a03851681522054610542908363ffffffff61067416565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546105d2908363ffffffff61067416565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561066e57fe5b50900390565b8181018281101561068157fe5b929150505600a165627a7a7230582002a8ac0e1fb78e81f525672fea77888ca221a22f23d1d03ea17a5f794f4c07950029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582001a72e96e30964cc9ea0e910abb2eca43143887c0c9e7633f2a66ce369e3e03d0029`

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}
