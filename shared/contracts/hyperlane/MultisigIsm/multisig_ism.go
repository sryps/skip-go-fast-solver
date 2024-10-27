// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisig_ism

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

// MultisigIsmMetaData contains all meta data concerning the MultisigIsm contract.
var MultisigIsmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"moduleType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"validatorsAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MultisigIsmABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisigIsmMetaData.ABI instead.
var MultisigIsmABI = MultisigIsmMetaData.ABI

// MultisigIsm is an auto generated Go binding around an Ethereum contract.
type MultisigIsm struct {
	MultisigIsmCaller     // Read-only binding to the contract
	MultisigIsmTransactor // Write-only binding to the contract
	MultisigIsmFilterer   // Log filterer for contract events
}

// MultisigIsmCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigIsmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigIsmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigIsmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigIsmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigIsmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigIsmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigIsmSession struct {
	Contract     *MultisigIsm      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigIsmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigIsmCallerSession struct {
	Contract *MultisigIsmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MultisigIsmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigIsmTransactorSession struct {
	Contract     *MultisigIsmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MultisigIsmRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigIsmRaw struct {
	Contract *MultisigIsm // Generic contract binding to access the raw methods on
}

// MultisigIsmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigIsmCallerRaw struct {
	Contract *MultisigIsmCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigIsmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigIsmTransactorRaw struct {
	Contract *MultisigIsmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisigIsm creates a new instance of MultisigIsm, bound to a specific deployed contract.
func NewMultisigIsm(address common.Address, backend bind.ContractBackend) (*MultisigIsm, error) {
	contract, err := bindMultisigIsm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultisigIsm{MultisigIsmCaller: MultisigIsmCaller{contract: contract}, MultisigIsmTransactor: MultisigIsmTransactor{contract: contract}, MultisigIsmFilterer: MultisigIsmFilterer{contract: contract}}, nil
}

// NewMultisigIsmCaller creates a new read-only instance of MultisigIsm, bound to a specific deployed contract.
func NewMultisigIsmCaller(address common.Address, caller bind.ContractCaller) (*MultisigIsmCaller, error) {
	contract, err := bindMultisigIsm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigIsmCaller{contract: contract}, nil
}

// NewMultisigIsmTransactor creates a new write-only instance of MultisigIsm, bound to a specific deployed contract.
func NewMultisigIsmTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigIsmTransactor, error) {
	contract, err := bindMultisigIsm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigIsmTransactor{contract: contract}, nil
}

// NewMultisigIsmFilterer creates a new log filterer instance of MultisigIsm, bound to a specific deployed contract.
func NewMultisigIsmFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigIsmFilterer, error) {
	contract, err := bindMultisigIsm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigIsmFilterer{contract: contract}, nil
}

// bindMultisigIsm binds a generic wrapper to an already deployed contract.
func bindMultisigIsm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultisigIsmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultisigIsm *MultisigIsmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultisigIsm.Contract.MultisigIsmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultisigIsm *MultisigIsmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultisigIsm.Contract.MultisigIsmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultisigIsm *MultisigIsmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultisigIsm.Contract.MultisigIsmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultisigIsm *MultisigIsmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultisigIsm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultisigIsm *MultisigIsmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultisigIsm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultisigIsm *MultisigIsmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultisigIsm.Contract.contract.Transact(opts, method, params...)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_MultisigIsm *MultisigIsmCaller) ModuleType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MultisigIsm.contract.Call(opts, &out, "moduleType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_MultisigIsm *MultisigIsmSession) ModuleType() (uint8, error) {
	return _MultisigIsm.Contract.ModuleType(&_MultisigIsm.CallOpts)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_MultisigIsm *MultisigIsmCallerSession) ModuleType() (uint8, error) {
	return _MultisigIsm.Contract.ModuleType(&_MultisigIsm.CallOpts)
}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[] validators, uint8 threshold)
func (_MultisigIsm *MultisigIsmCaller) ValidatorsAndThreshold(opts *bind.CallOpts, _message []byte) (struct {
	Validators []common.Address
	Threshold  uint8
}, error) {
	var out []interface{}
	err := _MultisigIsm.contract.Call(opts, &out, "validatorsAndThreshold", _message)

	outstruct := new(struct {
		Validators []common.Address
		Threshold  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Threshold = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[] validators, uint8 threshold)
func (_MultisigIsm *MultisigIsmSession) ValidatorsAndThreshold(_message []byte) (struct {
	Validators []common.Address
	Threshold  uint8
}, error) {
	return _MultisigIsm.Contract.ValidatorsAndThreshold(&_MultisigIsm.CallOpts, _message)
}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[] validators, uint8 threshold)
func (_MultisigIsm *MultisigIsmCallerSession) ValidatorsAndThreshold(_message []byte) (struct {
	Validators []common.Address
	Threshold  uint8
}, error) {
	return _MultisigIsm.Contract.ValidatorsAndThreshold(&_MultisigIsm.CallOpts, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_MultisigIsm *MultisigIsmTransactor) Verify(opts *bind.TransactOpts, _metadata []byte, _message []byte) (*types.Transaction, error) {
	return _MultisigIsm.contract.Transact(opts, "verify", _metadata, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_MultisigIsm *MultisigIsmSession) Verify(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _MultisigIsm.Contract.Verify(&_MultisigIsm.TransactOpts, _metadata, _message)
}

// Verify is a paid mutator transaction binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) returns(bool)
func (_MultisigIsm *MultisigIsmTransactorSession) Verify(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _MultisigIsm.Contract.Verify(&_MultisigIsm.TransactOpts, _metadata, _message)
}
