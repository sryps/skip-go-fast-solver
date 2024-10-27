// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator_announce

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

// ValidatorAnnounceMetaData contains all meta data concerning the ValidatorAnnounce contract.
var ValidatorAnnounceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_storageLocation\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"announce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"getAnnouncedStorageLocations\",\"outputs\":[{\"internalType\":\"string[][]\",\"name\":\"\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAnnouncedValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mailbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ValidatorAnnounceABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorAnnounceMetaData.ABI instead.
var ValidatorAnnounceABI = ValidatorAnnounceMetaData.ABI

// ValidatorAnnounce is an auto generated Go binding around an Ethereum contract.
type ValidatorAnnounce struct {
	ValidatorAnnounceCaller     // Read-only binding to the contract
	ValidatorAnnounceTransactor // Write-only binding to the contract
	ValidatorAnnounceFilterer   // Log filterer for contract events
}

// ValidatorAnnounceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorAnnounceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorAnnounceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorAnnounceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorAnnounceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorAnnounceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorAnnounceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorAnnounceSession struct {
	Contract     *ValidatorAnnounce // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorAnnounceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorAnnounceCallerSession struct {
	Contract *ValidatorAnnounceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorAnnounceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorAnnounceTransactorSession struct {
	Contract     *ValidatorAnnounceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorAnnounceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorAnnounceRaw struct {
	Contract *ValidatorAnnounce // Generic contract binding to access the raw methods on
}

// ValidatorAnnounceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorAnnounceCallerRaw struct {
	Contract *ValidatorAnnounceCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorAnnounceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorAnnounceTransactorRaw struct {
	Contract *ValidatorAnnounceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorAnnounce creates a new instance of ValidatorAnnounce, bound to a specific deployed contract.
func NewValidatorAnnounce(address common.Address, backend bind.ContractBackend) (*ValidatorAnnounce, error) {
	contract, err := bindValidatorAnnounce(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorAnnounce{ValidatorAnnounceCaller: ValidatorAnnounceCaller{contract: contract}, ValidatorAnnounceTransactor: ValidatorAnnounceTransactor{contract: contract}, ValidatorAnnounceFilterer: ValidatorAnnounceFilterer{contract: contract}}, nil
}

// NewValidatorAnnounceCaller creates a new read-only instance of ValidatorAnnounce, bound to a specific deployed contract.
func NewValidatorAnnounceCaller(address common.Address, caller bind.ContractCaller) (*ValidatorAnnounceCaller, error) {
	contract, err := bindValidatorAnnounce(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorAnnounceCaller{contract: contract}, nil
}

// NewValidatorAnnounceTransactor creates a new write-only instance of ValidatorAnnounce, bound to a specific deployed contract.
func NewValidatorAnnounceTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorAnnounceTransactor, error) {
	contract, err := bindValidatorAnnounce(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorAnnounceTransactor{contract: contract}, nil
}

// NewValidatorAnnounceFilterer creates a new log filterer instance of ValidatorAnnounce, bound to a specific deployed contract.
func NewValidatorAnnounceFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorAnnounceFilterer, error) {
	contract, err := bindValidatorAnnounce(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorAnnounceFilterer{contract: contract}, nil
}

// bindValidatorAnnounce binds a generic wrapper to an already deployed contract.
func bindValidatorAnnounce(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorAnnounceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorAnnounce *ValidatorAnnounceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorAnnounce.Contract.ValidatorAnnounceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorAnnounce *ValidatorAnnounceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorAnnounce.Contract.ValidatorAnnounceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorAnnounce *ValidatorAnnounceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorAnnounce.Contract.ValidatorAnnounceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorAnnounce *ValidatorAnnounceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorAnnounce.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorAnnounce *ValidatorAnnounceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorAnnounce.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorAnnounce *ValidatorAnnounceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorAnnounce.Contract.contract.Transact(opts, method, params...)
}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_ValidatorAnnounce *ValidatorAnnounceCaller) GetAnnouncedStorageLocations(opts *bind.CallOpts, _validators []common.Address) ([][]string, error) {
	var out []interface{}
	err := _ValidatorAnnounce.contract.Call(opts, &out, "getAnnouncedStorageLocations", _validators)

	if err != nil {
		return *new([][]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][]string)).(*[][]string)

	return out0, err

}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_ValidatorAnnounce *ValidatorAnnounceSession) GetAnnouncedStorageLocations(_validators []common.Address) ([][]string, error) {
	return _ValidatorAnnounce.Contract.GetAnnouncedStorageLocations(&_ValidatorAnnounce.CallOpts, _validators)
}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_ValidatorAnnounce *ValidatorAnnounceCallerSession) GetAnnouncedStorageLocations(_validators []common.Address) ([][]string, error) {
	return _ValidatorAnnounce.Contract.GetAnnouncedStorageLocations(&_ValidatorAnnounce.CallOpts, _validators)
}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_ValidatorAnnounce *ValidatorAnnounceCaller) GetAnnouncedValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorAnnounce.contract.Call(opts, &out, "getAnnouncedValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_ValidatorAnnounce *ValidatorAnnounceSession) GetAnnouncedValidators() ([]common.Address, error) {
	return _ValidatorAnnounce.Contract.GetAnnouncedValidators(&_ValidatorAnnounce.CallOpts)
}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_ValidatorAnnounce *ValidatorAnnounceCallerSession) GetAnnouncedValidators() ([]common.Address, error) {
	return _ValidatorAnnounce.Contract.GetAnnouncedValidators(&_ValidatorAnnounce.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ValidatorAnnounce *ValidatorAnnounceCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ValidatorAnnounce.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ValidatorAnnounce *ValidatorAnnounceSession) LocalDomain() (uint32, error) {
	return _ValidatorAnnounce.Contract.LocalDomain(&_ValidatorAnnounce.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ValidatorAnnounce *ValidatorAnnounceCallerSession) LocalDomain() (uint32, error) {
	return _ValidatorAnnounce.Contract.LocalDomain(&_ValidatorAnnounce.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_ValidatorAnnounce *ValidatorAnnounceCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorAnnounce.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_ValidatorAnnounce *ValidatorAnnounceSession) Mailbox() (common.Address, error) {
	return _ValidatorAnnounce.Contract.Mailbox(&_ValidatorAnnounce.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_ValidatorAnnounce *ValidatorAnnounceCallerSession) Mailbox() (common.Address, error) {
	return _ValidatorAnnounce.Contract.Mailbox(&_ValidatorAnnounce.CallOpts)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_ValidatorAnnounce *ValidatorAnnounceTransactor) Announce(opts *bind.TransactOpts, _validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _ValidatorAnnounce.contract.Transact(opts, "announce", _validator, _storageLocation, _signature)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_ValidatorAnnounce *ValidatorAnnounceSession) Announce(_validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _ValidatorAnnounce.Contract.Announce(&_ValidatorAnnounce.TransactOpts, _validator, _storageLocation, _signature)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_ValidatorAnnounce *ValidatorAnnounceTransactorSession) Announce(_validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _ValidatorAnnounce.Contract.Announce(&_ValidatorAnnounce.TransactOpts, _validator, _storageLocation, _signature)
}
