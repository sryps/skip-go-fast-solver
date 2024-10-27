// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregation_ism

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

// AggregationIsmMetaData contains all meta data concerning the AggregationIsm contract.
var AggregationIsmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"moduleType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"modulesAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AggregationIsmABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregationIsmMetaData.ABI instead.
var AggregationIsmABI = AggregationIsmMetaData.ABI

// AggregationIsm is an auto generated Go binding around an Ethereum contract.
type AggregationIsm struct {
	AggregationIsmCaller     // Read-only binding to the contract
	AggregationIsmTransactor // Write-only binding to the contract
	AggregationIsmFilterer   // Log filterer for contract events
}

// AggregationIsmCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregationIsmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationIsmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregationIsmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationIsmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregationIsmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationIsmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregationIsmSession struct {
	Contract     *AggregationIsm   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregationIsmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregationIsmCallerSession struct {
	Contract *AggregationIsmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AggregationIsmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregationIsmTransactorSession struct {
	Contract     *AggregationIsmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AggregationIsmRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregationIsmRaw struct {
	Contract *AggregationIsm // Generic contract binding to access the raw methods on
}

// AggregationIsmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregationIsmCallerRaw struct {
	Contract *AggregationIsmCaller // Generic read-only contract binding to access the raw methods on
}

// AggregationIsmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregationIsmTransactorRaw struct {
	Contract *AggregationIsmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregationIsm creates a new instance of AggregationIsm, bound to a specific deployed contract.
func NewAggregationIsm(address common.Address, backend bind.ContractBackend) (*AggregationIsm, error) {
	contract, err := bindAggregationIsm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregationIsm{AggregationIsmCaller: AggregationIsmCaller{contract: contract}, AggregationIsmTransactor: AggregationIsmTransactor{contract: contract}, AggregationIsmFilterer: AggregationIsmFilterer{contract: contract}}, nil
}

// NewAggregationIsmCaller creates a new read-only instance of AggregationIsm, bound to a specific deployed contract.
func NewAggregationIsmCaller(address common.Address, caller bind.ContractCaller) (*AggregationIsmCaller, error) {
	contract, err := bindAggregationIsm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationIsmCaller{contract: contract}, nil
}

// NewAggregationIsmTransactor creates a new write-only instance of AggregationIsm, bound to a specific deployed contract.
func NewAggregationIsmTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregationIsmTransactor, error) {
	contract, err := bindAggregationIsm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationIsmTransactor{contract: contract}, nil
}

// NewAggregationIsmFilterer creates a new log filterer instance of AggregationIsm, bound to a specific deployed contract.
func NewAggregationIsmFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregationIsmFilterer, error) {
	contract, err := bindAggregationIsm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregationIsmFilterer{contract: contract}, nil
}

// bindAggregationIsm binds a generic wrapper to an already deployed contract.
func bindAggregationIsm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregationIsmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationIsm *AggregationIsmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationIsm.Contract.AggregationIsmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationIsm *AggregationIsmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationIsm.Contract.AggregationIsmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationIsm *AggregationIsmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationIsm.Contract.AggregationIsmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationIsm *AggregationIsmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationIsm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationIsm *AggregationIsmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationIsm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationIsm *AggregationIsmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationIsm.Contract.contract.Transact(opts, method, params...)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_AggregationIsm *AggregationIsmCaller) ModuleType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AggregationIsm.contract.Call(opts, &out, "moduleType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_AggregationIsm *AggregationIsmSession) ModuleType() (uint8, error) {
	return _AggregationIsm.Contract.ModuleType(&_AggregationIsm.CallOpts)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_AggregationIsm *AggregationIsmCallerSession) ModuleType() (uint8, error) {
	return _AggregationIsm.Contract.ModuleType(&_AggregationIsm.CallOpts)
}

// ModulesAndThreshold is a free data retrieval call binding the contract method 0x6f72df75.
//
// Solidity: function modulesAndThreshold(bytes _message) view returns(address[] modules, uint8 threshold)
func (_AggregationIsm *AggregationIsmCaller) ModulesAndThreshold(opts *bind.CallOpts, _message []byte) (struct {
	Modules   []common.Address
	Threshold uint8
}, error) {
	var out []interface{}
	err := _AggregationIsm.contract.Call(opts, &out, "modulesAndThreshold", _message)

	outstruct := new(struct {
		Modules   []common.Address
		Threshold uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Modules = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Threshold = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// ModulesAndThreshold is a free data retrieval call binding the contract method 0x6f72df75.
//
// Solidity: function modulesAndThreshold(bytes _message) view returns(address[] modules, uint8 threshold)
func (_AggregationIsm *AggregationIsmSession) ModulesAndThreshold(_message []byte) (struct {
	Modules   []common.Address
	Threshold uint8
}, error) {
	return _AggregationIsm.Contract.ModulesAndThreshold(&_AggregationIsm.CallOpts, _message)
}

// ModulesAndThreshold is a free data retrieval call binding the contract method 0x6f72df75.
//
// Solidity: function modulesAndThreshold(bytes _message) view returns(address[] modules, uint8 threshold)
func (_AggregationIsm *AggregationIsmCallerSession) ModulesAndThreshold(_message []byte) (struct {
	Modules   []common.Address
	Threshold uint8
}, error) {
	return _AggregationIsm.Contract.ModulesAndThreshold(&_AggregationIsm.CallOpts, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_AggregationIsm *AggregationIsmTransactor) Verify(opts *bind.TransactOpts, _metadata []byte, _message []byte) (*types.Transaction, error) {
	return _AggregationIsm.contract.Transact(opts, "verify", _metadata, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_AggregationIsm *AggregationIsmSession) Verify(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _AggregationIsm.Contract.Verify(&_AggregationIsm.TransactOpts, _metadata, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_AggregationIsm *AggregationIsmTransactorSession) Verify(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _AggregationIsm.Contract.Verify(&_AggregationIsm.TransactOpts, _metadata, _message)
}
