// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchain_security_module

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
	_ = abi.ConvertType
)

// InterchainSecurityModuleMetaData contains all meta data concerning the InterchainSecurityModule contract.
var InterchainSecurityModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"moduleType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InterchainSecurityModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainSecurityModuleMetaData.ABI instead.
var InterchainSecurityModuleABI = InterchainSecurityModuleMetaData.ABI

// InterchainSecurityModule is an auto generated Go binding around an Ethereum contract.
type InterchainSecurityModule struct {
	InterchainSecurityModuleCaller     // Read-only binding to the contract
	InterchainSecurityModuleTransactor // Write-only binding to the contract
	InterchainSecurityModuleFilterer   // Log filterer for contract events
}

// InterchainSecurityModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainSecurityModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainSecurityModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainSecurityModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainSecurityModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainSecurityModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainSecurityModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainSecurityModuleSession struct {
	Contract     *InterchainSecurityModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InterchainSecurityModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainSecurityModuleCallerSession struct {
	Contract *InterchainSecurityModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// InterchainSecurityModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainSecurityModuleTransactorSession struct {
	Contract     *InterchainSecurityModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// InterchainSecurityModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainSecurityModuleRaw struct {
	Contract *InterchainSecurityModule // Generic contract binding to access the raw methods on
}

// InterchainSecurityModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainSecurityModuleCallerRaw struct {
	Contract *InterchainSecurityModuleCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainSecurityModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainSecurityModuleTransactorRaw struct {
	Contract *InterchainSecurityModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainSecurityModule creates a new instance of InterchainSecurityModule, bound to a specific deployed contract.
func NewInterchainSecurityModule(address common.Address, backend bind.ContractBackend) (*InterchainSecurityModule, error) {
	contract, err := bindInterchainSecurityModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainSecurityModule{InterchainSecurityModuleCaller: InterchainSecurityModuleCaller{contract: contract}, InterchainSecurityModuleTransactor: InterchainSecurityModuleTransactor{contract: contract}, InterchainSecurityModuleFilterer: InterchainSecurityModuleFilterer{contract: contract}}, nil
}

// NewInterchainSecurityModuleCaller creates a new read-only instance of InterchainSecurityModule, bound to a specific deployed contract.
func NewInterchainSecurityModuleCaller(address common.Address, caller bind.ContractCaller) (*InterchainSecurityModuleCaller, error) {
	contract, err := bindInterchainSecurityModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainSecurityModuleCaller{contract: contract}, nil
}

// NewInterchainSecurityModuleTransactor creates a new write-only instance of InterchainSecurityModule, bound to a specific deployed contract.
func NewInterchainSecurityModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainSecurityModuleTransactor, error) {
	contract, err := bindInterchainSecurityModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainSecurityModuleTransactor{contract: contract}, nil
}

// NewInterchainSecurityModuleFilterer creates a new log filterer instance of InterchainSecurityModule, bound to a specific deployed contract.
func NewInterchainSecurityModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainSecurityModuleFilterer, error) {
	contract, err := bindInterchainSecurityModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainSecurityModuleFilterer{contract: contract}, nil
}

// bindInterchainSecurityModule binds a generic wrapper to an already deployed contract.
func bindInterchainSecurityModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainSecurityModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainSecurityModule *InterchainSecurityModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainSecurityModule.Contract.InterchainSecurityModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainSecurityModule *InterchainSecurityModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainSecurityModule.Contract.InterchainSecurityModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainSecurityModule *InterchainSecurityModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainSecurityModule.Contract.InterchainSecurityModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainSecurityModule *InterchainSecurityModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainSecurityModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainSecurityModule *InterchainSecurityModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainSecurityModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainSecurityModule *InterchainSecurityModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainSecurityModule.Contract.contract.Transact(opts, method, params...)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_InterchainSecurityModule *InterchainSecurityModuleCaller) ModuleType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _InterchainSecurityModule.contract.Call(opts, &out, "moduleType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_InterchainSecurityModule *InterchainSecurityModuleSession) ModuleType() (uint8, error) {
	return _InterchainSecurityModule.Contract.ModuleType(&_InterchainSecurityModule.CallOpts)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_InterchainSecurityModule *InterchainSecurityModuleCallerSession) ModuleType() (uint8, error) {
	return _InterchainSecurityModule.Contract.ModuleType(&_InterchainSecurityModule.CallOpts)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_InterchainSecurityModule *InterchainSecurityModuleTransactor) Verify(opts *bind.TransactOpts, _metadata []byte, _message []byte) (*types.Transaction, error) {
	return _InterchainSecurityModule.contract.Transact(opts, "verify", _metadata, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_InterchainSecurityModule *InterchainSecurityModuleSession) Verify(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _InterchainSecurityModule.Contract.Verify(&_InterchainSecurityModule.TransactOpts, _metadata, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_InterchainSecurityModule *InterchainSecurityModuleTransactorSession) Verify(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _InterchainSecurityModule.Contract.Verify(&_InterchainSecurityModule.TransactOpts, _metadata, _message)
}
