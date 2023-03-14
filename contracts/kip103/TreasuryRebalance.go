// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kip103

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IRetiredContractABI is the input ABI used to generate the binding from.
const IRetiredContractABI = "[{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"adminList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"quorom\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRetiredContractBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IRetiredContractBinRuntime = ``

// IRetiredContractFuncSigs maps the 4-byte function signature to its string representation.
var IRetiredContractFuncSigs = map[string]string{
	"1865c57d": "getState()",
}

// IRetiredContract is an auto generated Go binding around a Klaytn contract.
type IRetiredContract struct {
	IRetiredContractCaller     // Read-only binding to the contract
	IRetiredContractTransactor // Write-only binding to the contract
	IRetiredContractFilterer   // Log filterer for contract events
}

// IRetiredContractCaller is an auto generated read-only Go binding around a Klaytn contract.
type IRetiredContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRetiredContractTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IRetiredContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRetiredContractFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IRetiredContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRetiredContractSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IRetiredContractSession struct {
	Contract     *IRetiredContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRetiredContractCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IRetiredContractCallerSession struct {
	Contract *IRetiredContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IRetiredContractTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IRetiredContractTransactorSession struct {
	Contract     *IRetiredContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IRetiredContractRaw is an auto generated low-level Go binding around a Klaytn contract.
type IRetiredContractRaw struct {
	Contract *IRetiredContract // Generic contract binding to access the raw methods on
}

// IRetiredContractCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IRetiredContractCallerRaw struct {
	Contract *IRetiredContractCaller // Generic read-only contract binding to access the raw methods on
}

// IRetiredContractTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IRetiredContractTransactorRaw struct {
	Contract *IRetiredContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRetiredContract creates a new instance of IRetiredContract, bound to a specific deployed contract.
func NewIRetiredContract(address common.Address, backend bind.ContractBackend) (*IRetiredContract, error) {
	contract, err := bindIRetiredContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRetiredContract{IRetiredContractCaller: IRetiredContractCaller{contract: contract}, IRetiredContractTransactor: IRetiredContractTransactor{contract: contract}, IRetiredContractFilterer: IRetiredContractFilterer{contract: contract}}, nil
}

// NewIRetiredContractCaller creates a new read-only instance of IRetiredContract, bound to a specific deployed contract.
func NewIRetiredContractCaller(address common.Address, caller bind.ContractCaller) (*IRetiredContractCaller, error) {
	contract, err := bindIRetiredContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRetiredContractCaller{contract: contract}, nil
}

// NewIRetiredContractTransactor creates a new write-only instance of IRetiredContract, bound to a specific deployed contract.
func NewIRetiredContractTransactor(address common.Address, transactor bind.ContractTransactor) (*IRetiredContractTransactor, error) {
	contract, err := bindIRetiredContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRetiredContractTransactor{contract: contract}, nil
}

// NewIRetiredContractFilterer creates a new log filterer instance of IRetiredContract, bound to a specific deployed contract.
func NewIRetiredContractFilterer(address common.Address, filterer bind.ContractFilterer) (*IRetiredContractFilterer, error) {
	contract, err := bindIRetiredContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRetiredContractFilterer{contract: contract}, nil
}

// bindIRetiredContract binds a generic wrapper to an already deployed contract.
func bindIRetiredContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRetiredContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRetiredContract *IRetiredContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRetiredContract.Contract.IRetiredContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRetiredContract *IRetiredContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRetiredContract.Contract.IRetiredContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRetiredContract *IRetiredContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRetiredContract.Contract.IRetiredContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRetiredContract *IRetiredContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRetiredContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRetiredContract *IRetiredContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRetiredContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRetiredContract *IRetiredContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRetiredContract.Contract.contract.Transact(opts, method, params...)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(address[] adminList, uint256 quorom)
func (_IRetiredContract *IRetiredContractCaller) GetState(opts *bind.CallOpts) (struct {
	AdminList []common.Address
	Quorom    *big.Int
}, error,
) {
	ret := new(struct {
		AdminList []common.Address
		Quorom    *big.Int
	})
	out := ret
	err := _IRetiredContract.contract.Call(opts, out, "getState")
	return *ret, err
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(address[] adminList, uint256 quorom)
func (_IRetiredContract *IRetiredContractSession) GetState() (struct {
	AdminList []common.Address
	Quorom    *big.Int
}, error,
) {
	return _IRetiredContract.Contract.GetState(&_IRetiredContract.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(address[] adminList, uint256 quorom)
func (_IRetiredContract *IRetiredContractCallerSession) GetState() (struct {
	AdminList []common.Address
	Quorom    *big.Int
}, error,
) {
	return _IRetiredContract.Contract.GetState(&_IRetiredContract.CallOpts)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OwnableBinRuntime = `608060405234801561001057600080fd5b506004361061004c5760003560e01c8063715018a6146100515780638da5cb5b1461005b5780638f32d59b1461007b578063f2fde38b14610099575b600080fd5b6100596100ac565b005b6000546040516001600160a01b0390911681526020015b60405180910390f35b6000546001600160a01b031633146040519015158152602001610072565b6100596100a736600461027b565b610155565b6000546001600160a01b0316331461010b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031633146101af5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610102565b6101b8816101bb565b50565b6001600160a01b0381166102205760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610102565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006020828403121561028d57600080fd5b81356001600160a01b03811681146102a457600080fd5b939250505056fea264697066735822122044d0b28ebacaa88abe8a91da63ebd3ae689b203af3974796d70122e5309ab4f464736f6c63430008130033`

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163390811782556040519091907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a36102e18061005f6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063715018a6146100515780638da5cb5b1461005b5780638f32d59b1461007b578063f2fde38b14610099575b600080fd5b6100596100ac565b005b6000546040516001600160a01b0390911681526020015b60405180910390f35b6000546001600160a01b031633146040519015158152602001610072565b6100596100a736600461027b565b610155565b6000546001600160a01b0316331461010b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031633146101af5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610102565b6101b8816101bb565b50565b6001600160a01b0381166102205760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610102565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006020828403121561028d57600080fd5b81356001600160a01b03811681146102a457600080fd5b939250505056fea264697066735822122044d0b28ebacaa88abe8a91da63ebd3ae689b203af3974796d70122e5309ab4f464736f6c63430008130033"

// DeployOwnable deploys a new Klaytn contract, binding an instance of Ownable to it.
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

// Ownable is an auto generated Go binding around a Klaytn contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around a Klaytn contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around a Klaytn contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceABI is the input ABI used to generate the binding from.
const TreasuryRebalanceABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rebalanceBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"retired\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"approversCount\",\"type\":\"uint256\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumTreasuryRebalance.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rebalanceBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deployedBlockNumber\",\"type\":\"uint256\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumTreasuryRebalance.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newbie\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundAllocation\",\"type\":\"uint256\"}],\"name\":\"NewbieRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newbie\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newbieCount\",\"type\":\"uint256\"}],\"name\":\"NewbieRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"retired\",\"type\":\"address\"}],\"name\":\"RetiredRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"retired\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"retiredCount\",\"type\":\"uint256\"}],\"name\":\"RetiredRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumTreasuryRebalance.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_retiredAddress\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkRetiredsApproved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"finalizeContract\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"retirees\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"totalRetireesBalance\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"newbies\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"totalNewbiesFund\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newbieAddress\",\"type\":\"address\"}],\"name\":\"getNewbie\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewbieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newbieAddress\",\"type\":\"address\"}],\"name\":\"getNewbieIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_retiredAddress\",\"type\":\"address\"}],\"name\":\"getRetired\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRetiredCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_retiredAddress\",\"type\":\"address\"}],\"name\":\"getRetiredIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"treasuryAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isContractAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"memo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newbieAddress\",\"type\":\"address\"}],\"name\":\"newbieExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newbies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newbie\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newbieAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"registerNewbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_retiredAddress\",\"type\":\"address\"}],\"name\":\"registerRetired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newbieAddress\",\"type\":\"address\"}],\"name\":\"removeNewbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_retiredAddress\",\"type\":\"address\"}],\"name\":\"removeRetired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_retiredAddress\",\"type\":\"address\"}],\"name\":\"retiredExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"retirees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"retired\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumTreasuryRebalance.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sumOfRetiredBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"retireesBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TreasuryRebalanceBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const TreasuryRebalanceBinRuntime = `6080604052600436106101cd5760003560e01c80638da5cb5b116100f7578063d826f88f11610095578063ea6d4a9b11610064578063ea6d4a9b1461057d578063eb5a8e55146105ad578063f2fde38b146105cd578063faaf9ca6146105ed576101cd565b8063d826f88f14610512578063daea85c514610527578063e20fcf0014610547578063e2384cb31461055c576101cd565b806394393e11116100d157806394393e111461047b578063966e0794146104ba578063bf680590146104cf578063d1ed33fc146104fd576101cd565b80638da5cb5b146104285780638f32d59b1461044657806391734d8614610466576101cd565b806349a3fb451161016f578063681f6e7c1161013e578063681f6e7c146103b3578063683e13cb146103d35780636864b95b146103f3578063715018a614610413576101cd565b806349a3fb451461032357806358c3b870146103395780635a12667b1461035b578063652e27e014610393576101cd565b80631f8c1798116101ab5780631f8c1798146102b2578063200d2ed2146102d257806345205a6b146102f9578063484090961461030e576101cd565b806301784e051461022d57806311f5c466146102625780631c1dac5914610290575b60405162461bcd60e51b815260206004820152602a60248201527f5468697320636f6e747261637420646f6573206e6f742061636365707420616e60448201526979207061796d656e747360b01b60648201526084015b60405180910390fd5b34801561023957600080fd5b5061024d610248366004611f55565b610602565b60405190151581526020015b60405180910390f35b34801561026e57600080fd5b5061028261027d366004611f55565b6106b6565b604051908152602001610259565b34801561029c57600080fd5b506102b06102ab366004611f55565b610722565b005b3480156102be57600080fd5b506102b06102cd366004611f55565b6108ca565b3480156102de57600080fd5b506003546102ec9060ff1681565b6040516102599190611fb1565b34801561030557600080fd5b50610282610a0f565b34801561031a57600080fd5b506102b0610a6d565b34801561032f57600080fd5b5061028260045481565b34801561034557600080fd5b5061034e610b24565b6040516102599190611fc5565b34801561036757600080fd5b5061037b610376366004612013565b610bb2565b6040516001600160a01b039091168152602001610259565b34801561039f57600080fd5b506102b06103ae36600461202c565b610be1565b3480156103bf57600080fd5b506102826103ce366004611f55565b610dca565b3480156103df57600080fd5b5061024d6103ee366004611f55565b610e2c565b3480156103ff57600080fd5b506102b061040e366004611f55565b610eda565b34801561041f57600080fd5b506102b061108e565b34801561043457600080fd5b506000546001600160a01b031661037b565b34801561045257600080fd5b506000546001600160a01b0316331461024d565b34801561047257600080fd5b50600254610282565b34801561048757600080fd5b5061049b610496366004612013565b611102565b604080516001600160a01b039093168352602083019190915201610259565b3480156104c657600080fd5b506102b061113a565b3480156104db57600080fd5b506104ef6104ea366004611f55565b61131a565b60405161025992919061209c565b34801561050957600080fd5b50600154610282565b34801561051e57600080fd5b506102b0611401565b34801561053357600080fd5b506102b0610542366004611f55565b6114e0565b34801561055357600080fd5b506102826116c4565b34801561056857600080fd5b5061024d610577366004611f55565b3b151590565b34801561058957600080fd5b5061059d61059836600461210f565b611716565b60405161025994939291906121a4565b3480156105b957600080fd5b5061049b6105c8366004611f55565b611866565b3480156105d957600080fd5b506102b06105e8366004611f55565b611916565b3480156105f957600080fd5b506102b0611949565b60006001600160a01b03821661064c5760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606401610224565b60005b6001548110156106b057826001600160a01b031660018281548110610676576106766121e1565b60009182526020909120600290910201546001600160a01b03160361069e5750600192915050565b806106a88161220d565b91505061064f565b50919050565b6000805b60025481101561071857826001600160a01b0316600282815481106106e1576106e16121e1565b60009182526020909120600290910201546001600160a01b0316036107065792915050565b806107108161220d565b9150506106ba565b5060001992915050565b6000546001600160a01b0316331461074c5760405162461bcd60e51b815260040161022490612226565b6000806003805460ff169081111561076657610766611f79565b146107835760405162461bcd60e51b81526004016102249061225b565b600061078e83610dca565b905060001981036107b15760405162461bcd60e51b815260040161022490612292565b600180546107c09082906122c2565b815481106107d0576107d06121e1565b9060005260206000209060020201600182815481106107f1576107f16121e1565b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600180830180546108349284019190611df5565b509050506001805480610849576108496122d5565b60008281526020812060026000199093019283020180546001600160a01b03191681559061087a6001830182611e41565b50509055600154604080516001600160a01b038616815260208101929092527f13673a8c5e648e7bebdeb5f081555ca1843dba99063e9e55aa6399e61573c70d91015b60405180910390a1505050565b6000546001600160a01b031633146108f45760405162461bcd60e51b815260040161022490612226565b6000806003805460ff169081111561090e5761090e611f79565b1461092b5760405162461bcd60e51b81526004016102249061225b565b61093482610602565b1561098f5760405162461bcd60e51b815260206004820152602560248201527f52657469726564206164647265737320697320616c72656164792072656769736044820152641d195c995960da1b6064820152608401610224565b6001805480820182556000919091526002027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0384166001600160a01b0319909116811782556040519081527f7da2e87d0b02df1162d5736cc40dfcfffd17198aaf093ddff4a8f4eb26002fde906020016108bd565b6000805b600154811015610a695760018181548110610a3057610a306121e1565b6000918252602090912060029091020154610a55906001600160a01b031631836122eb565b915080610a618161220d565b915050610a13565b5090565b6000546001600160a01b03163314610a975760405162461bcd60e51b815260040161022490612226565b6000806003805460ff1690811115610ab157610ab1611f79565b14610ace5760405162461bcd60e51b81526004016102249061225b565b600380546001919060ff191682805b02179055506003546040517fafa725e7f44cadb687a7043853fa1a7e7b8f0da74ce87ec546e9420f04da8c1e91610b199160ff90911690611fb1565b60405180910390a150565b60058054610b31906122fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5d906122fe565b8015610baa5780601f10610b7f57610100808354040283529160200191610baa565b820191906000526020600020905b815481529060010190602001808311610b8d57829003601f168201915b505050505081565b60018181548110610bc257600080fd5b60009182526020909120600290910201546001600160a01b0316905081565b6000546001600160a01b03163314610c0b5760405162461bcd60e51b815260040161022490612226565b6000806003805460ff1690811115610c2557610c25611f79565b14610c425760405162461bcd60e51b81526004016102249061225b565b610c4b83610e2c565b15610ca45760405162461bcd60e51b8152602060048201526024808201527f4e6577626965206164647265737320697320616c726561647920726567697374604482015263195c995960e21b6064820152608401610224565b81600003610cf45760405162461bcd60e51b815260206004820152601960248201527f416d6f756e742063616e6e6f742062652073657420746f2030000000000000006044820152606401610224565b6040805180820182526001600160a01b038581168083526020808401878152600280546001810182556000829052865191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81018054929096166001600160a01b031990921691909117909455517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf90930192909255835190815290810185905290917fd261b37cd56b21cd1af841dca6331a133e5d8b9d55c2c6fe0ec822e2a303ef7491015b60405180910390a150505050565b6000805b60015481101561071857826001600160a01b031660018281548110610df557610df56121e1565b60009182526020909120600290910201546001600160a01b031603610e1a5792915050565b80610e248161220d565b915050610dce565b60006001600160a01b038216610e765760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606401610224565b60005b6002548110156106b057826001600160a01b031660028281548110610ea057610ea06121e1565b60009182526020909120600290910201546001600160a01b031603610ec85750600192915050565b80610ed28161220d565b915050610e79565b6000546001600160a01b03163314610f045760405162461bcd60e51b815260040161022490612226565b6000806003805460ff1690811115610f1e57610f1e611f79565b14610f3b5760405162461bcd60e51b81526004016102249061225b565b6000610f46836106b6565b90506000198103610f915760405162461bcd60e51b815260206004820152601560248201527413995dd89a59481b9bdd081c9959da5cdd195c9959605a1b6044820152606401610224565b60028054610fa1906001906122c2565b81548110610fb157610fb16121e1565b906000526020600020906002020160028281548110610fd257610fd26121e1565b600091825260209091208254600292830290910180546001600160a01b0319166001600160a01b0390921691909117815560019283015492019190915580548061101e5761101e6122d5565b6000828152602080822060001993909301600281810290940180546001600160a01b031916815560010192909255925554604080516001600160a01b0387168152928301919091527f16c9a1659c2467bd64845207a0990d56629de7a3955c0f93c19354cda213c93291016108bd565b6000546001600160a01b031633146110b85760405162461bcd60e51b815260040161022490612226565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6002818154811061111257600080fd5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b60005b6001548110156113175760006001828154811061115c5761115c6121e1565b6000918252602091829020604080518082018252600290930290910180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156111dc57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111be575b505050505081525050905060006111f782600001513b151590565b905080156112b85760008061120f8460000151611a5d565b915091508084602001515110156112385760405162461bcd60e51b815260040161022490612332565b60208401516000805b825181101561128e5761126d83828151811061125f5761125f6121e1565b602002602001015186611ad6565b50816112788161220d565b92505080806112869061220d565b915050611241565b50828110156112af5760405162461bcd60e51b815260040161022490612332565b50505050611302565b8160200151516001146113025760405162461bcd60e51b8152602060048201526012602482015271454f412073686f756c6420617070726f766560701b6044820152606401610224565b5050808061130f9061220d565b91505061113d565b50565b60006060600061132984610dca565b9050600019810361134c5760405162461bcd60e51b815260040161022490612292565b600060018281548110611361576113616121e1565b6000918252602091829020604080518082018252600290930290910180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156113e157602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113c3575b505050505081525050905080600001518160200151935093505050915091565b6000546001600160a01b0316331461142b5760405162461bcd60e51b815260040161022490612226565b6003805460ff168181111561144257611442611f79565b14158015611451575060045443105b6114b05760405162461bcd60e51b815260206004820152602a60248201527f436f6e74726163742069732066696e616c697a65642c2063616e6e6f742072656044820152697365742076616c75657360b01b6064820152608401610224565b6114bc60016000611e5f565b6114c860026000611e80565b6114d460056000611ea1565b6003805460ff19169055565b6001806003805460ff16908111156114fa576114fa611f79565b146115175760405162461bcd60e51b81526004016102249061225b565b61152082610602565b6115835760405162461bcd60e51b815260206004820152602e60248201527f72657469726564206e6565647320746f2062652072656769737465726564206260448201526d19599bdc9948185c1c1c9bdd985b60921b6064820152608401610224565b813b1515806115ff57336001600160a01b038416146115f05760405162461bcd60e51b8152602060048201526024808201527f7265746972656441646472657373206973206e6f7420746865206d73672e7365604482015263373232b960e11b6064820152608401610224565b6115fa8333611b33565b505050565b600061160a84611a5d565b509050805160000361165e5760405162461bcd60e51b815260206004820152601a60248201527f61646d696e206c6973742063616e6e6f7420626520656d7074790000000000006044820152606401610224565b6116683382611ad6565b6116b45760405162461bcd60e51b815260206004820152601b60248201527f6d73672e73656e646572206973206e6f74207468652061646d696e00000000006044820152606401610224565b6116be8433611b33565b50505050565b6000805b600254811015610a6957600281815481106116e5576116e56121e1565b9060005260206000209060020201600101548261170291906122eb565b91508061170e8161220d565b9150506116c8565b6060600060606000611734600054336001600160a01b039091161490565b6117505760405162461bcd60e51b815260040161022490612226565b6002806003805460ff169081111561176a5761176a611f79565b146117875760405162461bcd60e51b81526004016102249061225b565b600561179387826123c2565b506003805460ff1916811781556040517f8f8636c7757ca9b7d154e1d44ca90d8e8c885b9eac417c59bbf8eb7779ca6404916117d29160059190612482565b60405180910390a16117e2610a0f565b93506117ec6116c4565b9150600454431161185e5760405162461bcd60e51b815260206004820152603660248201527f436f6e74726163742063616e206f6e6c792066696e616c697a6520616674657260448201527520657865637574696e6720726562616c616e63696e6760501b6064820152608401610224565b509193509193565b6000806000611874846106b6565b905060001981036118bf5760405162461bcd60e51b815260206004820152601560248201527413995dd89a59481b9bdd081c9959da5cdd195c9959605a1b6044820152606401610224565b6000600282815481106118d4576118d46121e1565b60009182526020918290206040805180820190915260029092020180546001600160a01b03168083526001909101549190920181905290969095509350505050565b6000546001600160a01b031633146119405760405162461bcd60e51b815260040161022490612226565b61131781611d35565b6000546001600160a01b031633146119735760405162461bcd60e51b815260040161022490612226565b6001806003805460ff169081111561198d5761198d611f79565b146119aa5760405162461bcd60e51b81526004016102249061225b565b6119b2610a0f565b6119ba6116c4565b10611a415760405162461bcd60e51b815260206004820152604b60248201527f747265617375727920616d6f756e742073686f756c64206265206c657373207460448201527f68616e207468652073756d206f6620616c6c207265746972656420616464726560648201526a73732062616c616e63657360a81b608482015260a401610224565b611a4961113a565b600380546002919060ff1916600183610add565b6060600080839050806001600160a01b0316631865c57d6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611aa3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611acb9190810190612517565b909590945092505050565b6000805b8251811015611b2c57828181518110611af557611af56121e1565b60200260200101516001600160a01b0316846001600160a01b031603611b1a57600191505b80611b248161220d565b915050611ada565b5092915050565b6000611b3e83610dca565b90506000198103611b615760405162461bcd60e51b815260040161022490612292565b600060018281548110611b7657611b766121e1565b9060005260206000209060020201600101805480602002602001604051908101604052809291908181526020018280548015611bdb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611bbd575b5050505050905060005b8151811015611c6d57836001600160a01b0316828281518110611c0a57611c0a6121e1565b60200260200101516001600160a01b031603611c5b5760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e48185c1c1c9bdd995960821b6044820152606401610224565b80611c658161220d565b915050611be5565b5060018281548110611c8157611c816121e1565b600091825260208083206001600290930201820180548084018255908452922090910180546001600160a01b0386166001600160a01b031990911617905580547f80da462ebfbe41cfc9bc015e7a9a3c7a2a73dbccede72d8ceb583606c27f8f9091869186919086908110611cf857611cf86121e1565b600091825260209182902060016002909202010154604080516001600160a01b039586168152949093169184019190915290820152606001610dbc565b6001600160a01b038116611d9a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610224565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054828255906000526020600020908101928215611e355760005260206000209182015b82811115611e35578254825591600101919060010190611e1a565b50610a69929150611ed7565b50805460008255906000526020600020908101906113179190611ed7565b50805460008255600202906000526020600020908101906113179190611eec565b50805460008255600202906000526020600020908101906113179190611f1a565b508054611ead906122fe565b6000825580601f10611ebd575050565b601f01602090049060005260206000209081019061131791905b5b80821115610a695760008155600101611ed8565b80821115610a695780546001600160a01b03191681556000611f116001830182611e41565b50600201611eec565b5b80821115610a695780546001600160a01b031916815560006001820155600201611f1b565b6001600160a01b038116811461131757600080fd5b600060208284031215611f6757600080fd5b8135611f7281611f40565b9392505050565b634e487b7160e01b600052602160045260246000fd5b60048110611fad57634e487b7160e01b600052602160045260246000fd5b9052565b60208101611fbf8284611f8f565b92915050565b600060208083528351808285015260005b81811015611ff257858101830151858201604001528201611fd6565b506000604082860101526040601f19601f8301168501019250505092915050565b60006020828403121561202557600080fd5b5035919050565b6000806040838503121561203f57600080fd5b823561204a81611f40565b946020939093013593505050565b600081518084526020808501945080840160005b838110156120915781516001600160a01b03168752958201959082019060010161206c565b509495945050505050565b6001600160a01b03831681526040602082018190526000906120c090830184612058565b949350505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715612107576121076120c8565b604052919050565b6000602080838503121561212257600080fd5b823567ffffffffffffffff8082111561213a57600080fd5b818501915085601f83011261214e57600080fd5b813581811115612160576121606120c8565b612172601f8201601f191685016120de565b9150808252868482850101111561218857600080fd5b8084840185840137600090820190930192909252509392505050565b6080815260006121b76080830187612058565b85602084015282810360408401526121cf8186612058565b91505082606083015295945050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161221f5761221f6121f7565b5060010190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601c908201527f4e6f7420696e207468652064657369676e617465642073746174757300000000604082015260600190565b60208082526016908201527514995d1a5c9959081b9bdd081c9959da5cdd195c995960521b604082015260600190565b81810381811115611fbf57611fbf6121f7565b634e487b7160e01b600052603160045260246000fd5b80820180821115611fbf57611fbf6121f7565b600181811c9082168061231257607f821691505b6020821081036106b057634e487b7160e01b600052602260045260246000fd5b60208082526022908201527f6d696e2072657175697265642061646d696e732073686f756c6420617070726f604082015261766560f01b606082015260800190565b601f8211156115fa57600081815260208120601f850160051c8101602086101561239b5750805b601f850160051c820191505b818110156123ba578281556001016123a7565b505050505050565b815167ffffffffffffffff8111156123dc576123dc6120c8565b6123f0816123ea84546122fe565b84612374565b602080601f831160018114612425576000841561240d5750858301515b600019600386901b1c1916600185901b1785556123ba565b600085815260208120601f198616915b8281101561245457888601518255948401946001909101908401612435565b50858210156124725787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000808454612494816122fe565b80604086015260606001808416600081146124b657600181146124d057612501565b60ff1985168884015283151560051b880183019550612501565b8960005260208060002060005b868110156124f85781548b82018701529084019082016124dd565b8a018501975050505b505050505080915050611f726020830184611f8f565b6000806040838503121561252a57600080fd5b825167ffffffffffffffff8082111561254257600080fd5b818501915085601f83011261255657600080fd5b815160208282111561256a5761256a6120c8565b8160051b925061257b8184016120de565b828152928401810192818101908985111561259557600080fd5b948201945b848610156125bf57855193506125af84611f40565b838252948201949082019061259a565b9790910151969896975050505050505056fea2646970667358221220da56eec5106708083501d5245dfa79d245c3bec491b3ef29f233b0b2c88c2f1164736f6c63430008130033`

// TreasuryRebalanceFuncSigs maps the 4-byte function signature to its string representation.
var TreasuryRebalanceFuncSigs = map[string]string{
	"daea85c5": "approve(address)",
	"966e0794": "checkRetiredsApproved()",
	"faaf9ca6": "finalizeApproval()",
	"ea6d4a9b": "finalizeContract(string)",
	"48409096": "finalizeRegistration()",
	"eb5a8e55": "getNewbie(address)",
	"91734d86": "getNewbieCount()",
	"11f5c466": "getNewbieIndex(address)",
	"bf680590": "getRetired(address)",
	"d1ed33fc": "getRetiredCount()",
	"681f6e7c": "getRetiredIndex(address)",
	"e20fcf00": "getTreasuryAmount()",
	"e2384cb3": "isContractAddr(address)",
	"8f32d59b": "isOwner()",
	"58c3b870": "memo()",
	"683e13cb": "newbieExists(address)",
	"94393e11": "newbies(uint256)",
	"8da5cb5b": "owner()",
	"49a3fb45": "rebalanceBlockNumber()",
	"652e27e0": "registerNewbie(address,uint256)",
	"1f8c1798": "registerRetired(address)",
	"6864b95b": "removeNewbie(address)",
	"1c1dac59": "removeRetired(address)",
	"715018a6": "renounceOwnership()",
	"d826f88f": "reset()",
	"01784e05": "retiredExists(address)",
	"5a12667b": "retirees(uint256)",
	"200d2ed2": "status()",
	"45205a6b": "sumOfRetiredBalance()",
	"f2fde38b": "transferOwnership(address)",
}

// TreasuryRebalanceBin is the compiled bytecode used for deploying new contracts.
var TreasuryRebalanceBin = "0x60806040523480156200001157600080fd5b5060405162002730380380620027308339810160408190526200003491620000c8565b600080546001600160a01b0319163390811782556040519091907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a360048190556003805460ff191690556040517f6f182006c5a12fe70c0728eedb2d1b0628c41483ca6721c606707d778d22ed0a90620000b99060009084904290620000e2565b60405180910390a15062000119565b600060208284031215620000db57600080fd5b5051919050565b60608101600485106200010557634e487b7160e01b600052602160045260246000fd5b938152602081019290925260409091015290565b61260780620001296000396000f3fe6080604052600436106101cd5760003560e01c80638da5cb5b116100f7578063d826f88f11610095578063ea6d4a9b11610064578063ea6d4a9b1461057d578063eb5a8e55146105ad578063f2fde38b146105cd578063faaf9ca6146105ed576101cd565b8063d826f88f14610512578063daea85c514610527578063e20fcf0014610547578063e2384cb31461055c576101cd565b806394393e11116100d157806394393e111461047b578063966e0794146104ba578063bf680590146104cf578063d1ed33fc146104fd576101cd565b80638da5cb5b146104285780638f32d59b1461044657806391734d8614610466576101cd565b806349a3fb451161016f578063681f6e7c1161013e578063681f6e7c146103b3578063683e13cb146103d35780636864b95b146103f3578063715018a614610413576101cd565b806349a3fb451461032357806358c3b870146103395780635a12667b1461035b578063652e27e014610393576101cd565b80631f8c1798116101ab5780631f8c1798146102b2578063200d2ed2146102d257806345205a6b146102f9578063484090961461030e576101cd565b806301784e051461022d57806311f5c466146102625780631c1dac5914610290575b60405162461bcd60e51b815260206004820152602a60248201527f5468697320636f6e747261637420646f6573206e6f742061636365707420616e60448201526979207061796d656e747360b01b60648201526084015b60405180910390fd5b34801561023957600080fd5b5061024d610248366004611f55565b610602565b60405190151581526020015b60405180910390f35b34801561026e57600080fd5b5061028261027d366004611f55565b6106b6565b604051908152602001610259565b34801561029c57600080fd5b506102b06102ab366004611f55565b610722565b005b3480156102be57600080fd5b506102b06102cd366004611f55565b6108ca565b3480156102de57600080fd5b506003546102ec9060ff1681565b6040516102599190611fb1565b34801561030557600080fd5b50610282610a0f565b34801561031a57600080fd5b506102b0610a6d565b34801561032f57600080fd5b5061028260045481565b34801561034557600080fd5b5061034e610b24565b6040516102599190611fc5565b34801561036757600080fd5b5061037b610376366004612013565b610bb2565b6040516001600160a01b039091168152602001610259565b34801561039f57600080fd5b506102b06103ae36600461202c565b610be1565b3480156103bf57600080fd5b506102826103ce366004611f55565b610dca565b3480156103df57600080fd5b5061024d6103ee366004611f55565b610e2c565b3480156103ff57600080fd5b506102b061040e366004611f55565b610eda565b34801561041f57600080fd5b506102b061108e565b34801561043457600080fd5b506000546001600160a01b031661037b565b34801561045257600080fd5b506000546001600160a01b0316331461024d565b34801561047257600080fd5b50600254610282565b34801561048757600080fd5b5061049b610496366004612013565b611102565b604080516001600160a01b039093168352602083019190915201610259565b3480156104c657600080fd5b506102b061113a565b3480156104db57600080fd5b506104ef6104ea366004611f55565b61131a565b60405161025992919061209c565b34801561050957600080fd5b50600154610282565b34801561051e57600080fd5b506102b0611401565b34801561053357600080fd5b506102b0610542366004611f55565b6114e0565b34801561055357600080fd5b506102826116c4565b34801561056857600080fd5b5061024d610577366004611f55565b3b151590565b34801561058957600080fd5b5061059d61059836600461210f565b611716565b60405161025994939291906121a4565b3480156105b957600080fd5b5061049b6105c8366004611f55565b611866565b3480156105d957600080fd5b506102b06105e8366004611f55565b611916565b3480156105f957600080fd5b506102b0611949565b60006001600160a01b03821661064c5760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606401610224565b60005b6001548110156106b057826001600160a01b031660018281548110610676576106766121e1565b60009182526020909120600290910201546001600160a01b03160361069e5750600192915050565b806106a88161220d565b91505061064f565b50919050565b6000805b60025481101561071857826001600160a01b0316600282815481106106e1576106e16121e1565b60009182526020909120600290910201546001600160a01b0316036107065792915050565b806107108161220d565b9150506106ba565b5060001992915050565b6000546001600160a01b0316331461074c5760405162461bcd60e51b815260040161022490612226565b6000806003805460ff169081111561076657610766611f79565b146107835760405162461bcd60e51b81526004016102249061225b565b600061078e83610dca565b905060001981036107b15760405162461bcd60e51b815260040161022490612292565b600180546107c09082906122c2565b815481106107d0576107d06121e1565b9060005260206000209060020201600182815481106107f1576107f16121e1565b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b03909216919091178155600180830180546108349284019190611df5565b509050506001805480610849576108496122d5565b60008281526020812060026000199093019283020180546001600160a01b03191681559061087a6001830182611e41565b50509055600154604080516001600160a01b038616815260208101929092527f13673a8c5e648e7bebdeb5f081555ca1843dba99063e9e55aa6399e61573c70d91015b60405180910390a1505050565b6000546001600160a01b031633146108f45760405162461bcd60e51b815260040161022490612226565b6000806003805460ff169081111561090e5761090e611f79565b1461092b5760405162461bcd60e51b81526004016102249061225b565b61093482610602565b1561098f5760405162461bcd60e51b815260206004820152602560248201527f52657469726564206164647265737320697320616c72656164792072656769736044820152641d195c995960da1b6064820152608401610224565b6001805480820182556000919091526002027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0384166001600160a01b0319909116811782556040519081527f7da2e87d0b02df1162d5736cc40dfcfffd17198aaf093ddff4a8f4eb26002fde906020016108bd565b6000805b600154811015610a695760018181548110610a3057610a306121e1565b6000918252602090912060029091020154610a55906001600160a01b031631836122eb565b915080610a618161220d565b915050610a13565b5090565b6000546001600160a01b03163314610a975760405162461bcd60e51b815260040161022490612226565b6000806003805460ff1690811115610ab157610ab1611f79565b14610ace5760405162461bcd60e51b81526004016102249061225b565b600380546001919060ff191682805b02179055506003546040517fafa725e7f44cadb687a7043853fa1a7e7b8f0da74ce87ec546e9420f04da8c1e91610b199160ff90911690611fb1565b60405180910390a150565b60058054610b31906122fe565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5d906122fe565b8015610baa5780601f10610b7f57610100808354040283529160200191610baa565b820191906000526020600020905b815481529060010190602001808311610b8d57829003601f168201915b505050505081565b60018181548110610bc257600080fd5b60009182526020909120600290910201546001600160a01b0316905081565b6000546001600160a01b03163314610c0b5760405162461bcd60e51b815260040161022490612226565b6000806003805460ff1690811115610c2557610c25611f79565b14610c425760405162461bcd60e51b81526004016102249061225b565b610c4b83610e2c565b15610ca45760405162461bcd60e51b8152602060048201526024808201527f4e6577626965206164647265737320697320616c726561647920726567697374604482015263195c995960e21b6064820152608401610224565b81600003610cf45760405162461bcd60e51b815260206004820152601960248201527f416d6f756e742063616e6e6f742062652073657420746f2030000000000000006044820152606401610224565b6040805180820182526001600160a01b038581168083526020808401878152600280546001810182556000829052865191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81018054929096166001600160a01b031990921691909117909455517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf90930192909255835190815290810185905290917fd261b37cd56b21cd1af841dca6331a133e5d8b9d55c2c6fe0ec822e2a303ef7491015b60405180910390a150505050565b6000805b60015481101561071857826001600160a01b031660018281548110610df557610df56121e1565b60009182526020909120600290910201546001600160a01b031603610e1a5792915050565b80610e248161220d565b915050610dce565b60006001600160a01b038216610e765760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606401610224565b60005b6002548110156106b057826001600160a01b031660028281548110610ea057610ea06121e1565b60009182526020909120600290910201546001600160a01b031603610ec85750600192915050565b80610ed28161220d565b915050610e79565b6000546001600160a01b03163314610f045760405162461bcd60e51b815260040161022490612226565b6000806003805460ff1690811115610f1e57610f1e611f79565b14610f3b5760405162461bcd60e51b81526004016102249061225b565b6000610f46836106b6565b90506000198103610f915760405162461bcd60e51b815260206004820152601560248201527413995dd89a59481b9bdd081c9959da5cdd195c9959605a1b6044820152606401610224565b60028054610fa1906001906122c2565b81548110610fb157610fb16121e1565b906000526020600020906002020160028281548110610fd257610fd26121e1565b600091825260209091208254600292830290910180546001600160a01b0319166001600160a01b0390921691909117815560019283015492019190915580548061101e5761101e6122d5565b6000828152602080822060001993909301600281810290940180546001600160a01b031916815560010192909255925554604080516001600160a01b0387168152928301919091527f16c9a1659c2467bd64845207a0990d56629de7a3955c0f93c19354cda213c93291016108bd565b6000546001600160a01b031633146110b85760405162461bcd60e51b815260040161022490612226565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6002818154811061111257600080fd5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b60005b6001548110156113175760006001828154811061115c5761115c6121e1565b6000918252602091829020604080518082018252600290930290910180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156111dc57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111be575b505050505081525050905060006111f782600001513b151590565b905080156112b85760008061120f8460000151611a5d565b915091508084602001515110156112385760405162461bcd60e51b815260040161022490612332565b60208401516000805b825181101561128e5761126d83828151811061125f5761125f6121e1565b602002602001015186611ad6565b50816112788161220d565b92505080806112869061220d565b915050611241565b50828110156112af5760405162461bcd60e51b815260040161022490612332565b50505050611302565b8160200151516001146113025760405162461bcd60e51b8152602060048201526012602482015271454f412073686f756c6420617070726f766560701b6044820152606401610224565b5050808061130f9061220d565b91505061113d565b50565b60006060600061132984610dca565b9050600019810361134c5760405162461bcd60e51b815260040161022490612292565b600060018281548110611361576113616121e1565b6000918252602091829020604080518082018252600290930290910180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156113e157602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113c3575b505050505081525050905080600001518160200151935093505050915091565b6000546001600160a01b0316331461142b5760405162461bcd60e51b815260040161022490612226565b6003805460ff168181111561144257611442611f79565b14158015611451575060045443105b6114b05760405162461bcd60e51b815260206004820152602a60248201527f436f6e74726163742069732066696e616c697a65642c2063616e6e6f742072656044820152697365742076616c75657360b01b6064820152608401610224565b6114bc60016000611e5f565b6114c860026000611e80565b6114d460056000611ea1565b6003805460ff19169055565b6001806003805460ff16908111156114fa576114fa611f79565b146115175760405162461bcd60e51b81526004016102249061225b565b61152082610602565b6115835760405162461bcd60e51b815260206004820152602e60248201527f72657469726564206e6565647320746f2062652072656769737465726564206260448201526d19599bdc9948185c1c1c9bdd985b60921b6064820152608401610224565b813b1515806115ff57336001600160a01b038416146115f05760405162461bcd60e51b8152602060048201526024808201527f7265746972656441646472657373206973206e6f7420746865206d73672e7365604482015263373232b960e11b6064820152608401610224565b6115fa8333611b33565b505050565b600061160a84611a5d565b509050805160000361165e5760405162461bcd60e51b815260206004820152601a60248201527f61646d696e206c6973742063616e6e6f7420626520656d7074790000000000006044820152606401610224565b6116683382611ad6565b6116b45760405162461bcd60e51b815260206004820152601b60248201527f6d73672e73656e646572206973206e6f74207468652061646d696e00000000006044820152606401610224565b6116be8433611b33565b50505050565b6000805b600254811015610a6957600281815481106116e5576116e56121e1565b9060005260206000209060020201600101548261170291906122eb565b91508061170e8161220d565b9150506116c8565b6060600060606000611734600054336001600160a01b039091161490565b6117505760405162461bcd60e51b815260040161022490612226565b6002806003805460ff169081111561176a5761176a611f79565b146117875760405162461bcd60e51b81526004016102249061225b565b600561179387826123c2565b506003805460ff1916811781556040517f8f8636c7757ca9b7d154e1d44ca90d8e8c885b9eac417c59bbf8eb7779ca6404916117d29160059190612482565b60405180910390a16117e2610a0f565b93506117ec6116c4565b9150600454431161185e5760405162461bcd60e51b815260206004820152603660248201527f436f6e74726163742063616e206f6e6c792066696e616c697a6520616674657260448201527520657865637574696e6720726562616c616e63696e6760501b6064820152608401610224565b509193509193565b6000806000611874846106b6565b905060001981036118bf5760405162461bcd60e51b815260206004820152601560248201527413995dd89a59481b9bdd081c9959da5cdd195c9959605a1b6044820152606401610224565b6000600282815481106118d4576118d46121e1565b60009182526020918290206040805180820190915260029092020180546001600160a01b03168083526001909101549190920181905290969095509350505050565b6000546001600160a01b031633146119405760405162461bcd60e51b815260040161022490612226565b61131781611d35565b6000546001600160a01b031633146119735760405162461bcd60e51b815260040161022490612226565b6001806003805460ff169081111561198d5761198d611f79565b146119aa5760405162461bcd60e51b81526004016102249061225b565b6119b2610a0f565b6119ba6116c4565b10611a415760405162461bcd60e51b815260206004820152604b60248201527f747265617375727920616d6f756e742073686f756c64206265206c657373207460448201527f68616e207468652073756d206f6620616c6c207265746972656420616464726560648201526a73732062616c616e63657360a81b608482015260a401610224565b611a4961113a565b600380546002919060ff1916600183610add565b6060600080839050806001600160a01b0316631865c57d6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611aa3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611acb9190810190612517565b909590945092505050565b6000805b8251811015611b2c57828181518110611af557611af56121e1565b60200260200101516001600160a01b0316846001600160a01b031603611b1a57600191505b80611b248161220d565b915050611ada565b5092915050565b6000611b3e83610dca565b90506000198103611b615760405162461bcd60e51b815260040161022490612292565b600060018281548110611b7657611b766121e1565b9060005260206000209060020201600101805480602002602001604051908101604052809291908181526020018280548015611bdb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611bbd575b5050505050905060005b8151811015611c6d57836001600160a01b0316828281518110611c0a57611c0a6121e1565b60200260200101516001600160a01b031603611c5b5760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e48185c1c1c9bdd995960821b6044820152606401610224565b80611c658161220d565b915050611be5565b5060018281548110611c8157611c816121e1565b600091825260208083206001600290930201820180548084018255908452922090910180546001600160a01b0386166001600160a01b031990911617905580547f80da462ebfbe41cfc9bc015e7a9a3c7a2a73dbccede72d8ceb583606c27f8f9091869186919086908110611cf857611cf86121e1565b600091825260209182902060016002909202010154604080516001600160a01b039586168152949093169184019190915290820152606001610dbc565b6001600160a01b038116611d9a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610224565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054828255906000526020600020908101928215611e355760005260206000209182015b82811115611e35578254825591600101919060010190611e1a565b50610a69929150611ed7565b50805460008255906000526020600020908101906113179190611ed7565b50805460008255600202906000526020600020908101906113179190611eec565b50805460008255600202906000526020600020908101906113179190611f1a565b508054611ead906122fe565b6000825580601f10611ebd575050565b601f01602090049060005260206000209081019061131791905b5b80821115610a695760008155600101611ed8565b80821115610a695780546001600160a01b03191681556000611f116001830182611e41565b50600201611eec565b5b80821115610a695780546001600160a01b031916815560006001820155600201611f1b565b6001600160a01b038116811461131757600080fd5b600060208284031215611f6757600080fd5b8135611f7281611f40565b9392505050565b634e487b7160e01b600052602160045260246000fd5b60048110611fad57634e487b7160e01b600052602160045260246000fd5b9052565b60208101611fbf8284611f8f565b92915050565b600060208083528351808285015260005b81811015611ff257858101830151858201604001528201611fd6565b506000604082860101526040601f19601f8301168501019250505092915050565b60006020828403121561202557600080fd5b5035919050565b6000806040838503121561203f57600080fd5b823561204a81611f40565b946020939093013593505050565b600081518084526020808501945080840160005b838110156120915781516001600160a01b03168752958201959082019060010161206c565b509495945050505050565b6001600160a01b03831681526040602082018190526000906120c090830184612058565b949350505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715612107576121076120c8565b604052919050565b6000602080838503121561212257600080fd5b823567ffffffffffffffff8082111561213a57600080fd5b818501915085601f83011261214e57600080fd5b813581811115612160576121606120c8565b612172601f8201601f191685016120de565b9150808252868482850101111561218857600080fd5b8084840185840137600090820190930192909252509392505050565b6080815260006121b76080830187612058565b85602084015282810360408401526121cf8186612058565b91505082606083015295945050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161221f5761221f6121f7565b5060010190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601c908201527f4e6f7420696e207468652064657369676e617465642073746174757300000000604082015260600190565b60208082526016908201527514995d1a5c9959081b9bdd081c9959da5cdd195c995960521b604082015260600190565b81810381811115611fbf57611fbf6121f7565b634e487b7160e01b600052603160045260246000fd5b80820180821115611fbf57611fbf6121f7565b600181811c9082168061231257607f821691505b6020821081036106b057634e487b7160e01b600052602260045260246000fd5b60208082526022908201527f6d696e2072657175697265642061646d696e732073686f756c6420617070726f604082015261766560f01b606082015260800190565b601f8211156115fa57600081815260208120601f850160051c8101602086101561239b5750805b601f850160051c820191505b818110156123ba578281556001016123a7565b505050505050565b815167ffffffffffffffff8111156123dc576123dc6120c8565b6123f0816123ea84546122fe565b84612374565b602080601f831160018114612425576000841561240d5750858301515b600019600386901b1c1916600185901b1785556123ba565b600085815260208120601f198616915b8281101561245457888601518255948401946001909101908401612435565b50858210156124725787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000808454612494816122fe565b80604086015260606001808416600081146124b657600181146124d057612501565b60ff1985168884015283151560051b880183019550612501565b8960005260208060002060005b868110156124f85781548b82018701529084019082016124dd565b8a018501975050505b505050505080915050611f726020830184611f8f565b6000806040838503121561252a57600080fd5b825167ffffffffffffffff8082111561254257600080fd5b818501915085601f83011261255657600080fd5b815160208282111561256a5761256a6120c8565b8160051b925061257b8184016120de565b828152928401810192818101908985111561259557600080fd5b948201945b848610156125bf57855193506125af84611f40565b838252948201949082019061259a565b9790910151969896975050505050505056fea2646970667358221220da56eec5106708083501d5245dfa79d245c3bec491b3ef29f233b0b2c88c2f1164736f6c63430008130033"

// DeployTreasuryRebalance deploys a new Klaytn contract, binding an instance of TreasuryRebalance to it.
func DeployTreasuryRebalance(auth *bind.TransactOpts, backend bind.ContractBackend, _rebalanceBlockNumber *big.Int) (common.Address, *types.Transaction, *TreasuryRebalance, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryRebalanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TreasuryRebalanceBin), backend, _rebalanceBlockNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TreasuryRebalance{TreasuryRebalanceCaller: TreasuryRebalanceCaller{contract: contract}, TreasuryRebalanceTransactor: TreasuryRebalanceTransactor{contract: contract}, TreasuryRebalanceFilterer: TreasuryRebalanceFilterer{contract: contract}}, nil
}

// TreasuryRebalance is an auto generated Go binding around a Klaytn contract.
type TreasuryRebalance struct {
	TreasuryRebalanceCaller     // Read-only binding to the contract
	TreasuryRebalanceTransactor // Write-only binding to the contract
	TreasuryRebalanceFilterer   // Log filterer for contract events
}

// TreasuryRebalanceCaller is an auto generated read-only Go binding around a Klaytn contract.
type TreasuryRebalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryRebalanceTransactor is an auto generated write-only Go binding around a Klaytn contract.
type TreasuryRebalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryRebalanceFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type TreasuryRebalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryRebalanceSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type TreasuryRebalanceSession struct {
	Contract     *TreasuryRebalance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TreasuryRebalanceCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type TreasuryRebalanceCallerSession struct {
	Contract *TreasuryRebalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TreasuryRebalanceTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type TreasuryRebalanceTransactorSession struct {
	Contract     *TreasuryRebalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TreasuryRebalanceRaw is an auto generated low-level Go binding around a Klaytn contract.
type TreasuryRebalanceRaw struct {
	Contract *TreasuryRebalance // Generic contract binding to access the raw methods on
}

// TreasuryRebalanceCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type TreasuryRebalanceCallerRaw struct {
	Contract *TreasuryRebalanceCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryRebalanceTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type TreasuryRebalanceTransactorRaw struct {
	Contract *TreasuryRebalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasuryRebalance creates a new instance of TreasuryRebalance, bound to a specific deployed contract.
func NewTreasuryRebalance(address common.Address, backend bind.ContractBackend) (*TreasuryRebalance, error) {
	contract, err := bindTreasuryRebalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalance{TreasuryRebalanceCaller: TreasuryRebalanceCaller{contract: contract}, TreasuryRebalanceTransactor: TreasuryRebalanceTransactor{contract: contract}, TreasuryRebalanceFilterer: TreasuryRebalanceFilterer{contract: contract}}, nil
}

// NewTreasuryRebalanceCaller creates a new read-only instance of TreasuryRebalance, bound to a specific deployed contract.
func NewTreasuryRebalanceCaller(address common.Address, caller bind.ContractCaller) (*TreasuryRebalanceCaller, error) {
	contract, err := bindTreasuryRebalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceCaller{contract: contract}, nil
}

// NewTreasuryRebalanceTransactor creates a new write-only instance of TreasuryRebalance, bound to a specific deployed contract.
func NewTreasuryRebalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryRebalanceTransactor, error) {
	contract, err := bindTreasuryRebalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceTransactor{contract: contract}, nil
}

// NewTreasuryRebalanceFilterer creates a new log filterer instance of TreasuryRebalance, bound to a specific deployed contract.
func NewTreasuryRebalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryRebalanceFilterer, error) {
	contract, err := bindTreasuryRebalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceFilterer{contract: contract}, nil
}

// bindTreasuryRebalance binds a generic wrapper to an already deployed contract.
func bindTreasuryRebalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryRebalanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryRebalance *TreasuryRebalanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TreasuryRebalance.Contract.TreasuryRebalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryRebalance *TreasuryRebalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.TreasuryRebalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryRebalance *TreasuryRebalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.TreasuryRebalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasuryRebalance *TreasuryRebalanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TreasuryRebalance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasuryRebalance *TreasuryRebalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasuryRebalance *TreasuryRebalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.contract.Transact(opts, method, params...)
}

// CheckRetiredsApproved is a free data retrieval call binding the contract method 0x966e0794.
//
// Solidity: function checkRetiredsApproved() view returns()
func (_TreasuryRebalance *TreasuryRebalanceCaller) CheckRetiredsApproved(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _TreasuryRebalance.contract.Call(opts, out, "checkRetiredsApproved")
	return err
}

// CheckRetiredsApproved is a free data retrieval call binding the contract method 0x966e0794.
//
// Solidity: function checkRetiredsApproved() view returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) CheckRetiredsApproved() error {
	return _TreasuryRebalance.Contract.CheckRetiredsApproved(&_TreasuryRebalance.CallOpts)
}

// CheckRetiredsApproved is a free data retrieval call binding the contract method 0x966e0794.
//
// Solidity: function checkRetiredsApproved() view returns()
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) CheckRetiredsApproved() error {
	return _TreasuryRebalance.Contract.CheckRetiredsApproved(&_TreasuryRebalance.CallOpts)
}

// GetNewbie is a free data retrieval call binding the contract method 0xeb5a8e55.
//
// Solidity: function getNewbie(address _newbieAddress) view returns(address, uint256)
func (_TreasuryRebalance *TreasuryRebalanceCaller) GetNewbie(opts *bind.CallOpts, _newbieAddress common.Address) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TreasuryRebalance.contract.Call(opts, out, "getNewbie", _newbieAddress)
	return *ret0, *ret1, err
}

// GetNewbie is a free data retrieval call binding the contract method 0xeb5a8e55.
//
// Solidity: function getNewbie(address _newbieAddress) view returns(address, uint256)
func (_TreasuryRebalance *TreasuryRebalanceSession) GetNewbie(_newbieAddress common.Address) (common.Address, *big.Int, error) {
	return _TreasuryRebalance.Contract.GetNewbie(&_TreasuryRebalance.CallOpts, _newbieAddress)
}

// GetNewbie is a free data retrieval call binding the contract method 0xeb5a8e55.
//
// Solidity: function getNewbie(address _newbieAddress) view returns(address, uint256)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) GetNewbie(_newbieAddress common.Address) (common.Address, *big.Int, error) {
	return _TreasuryRebalance.Contract.GetNewbie(&_TreasuryRebalance.CallOpts, _newbieAddress)
}

// GetNewbieCount is a free data retrieval call binding the contract method 0x91734d86.
//
// Solidity: function getNewbieCount() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCaller) GetNewbieCount(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "getNewbieCount")
	return *ret0, err
}

// GetNewbieCount is a free data retrieval call binding the contract method 0x91734d86.
//
// Solidity: function getNewbieCount() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceSession) GetNewbieCount() (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetNewbieCount(&_TreasuryRebalance.CallOpts)
}

// GetNewbieCount is a free data retrieval call binding the contract method 0x91734d86.
//
// Solidity: function getNewbieCount() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) GetNewbieCount() (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetNewbieCount(&_TreasuryRebalance.CallOpts)
}

// GetNewbieIndex is a free data retrieval call binding the contract method 0x11f5c466.
//
// Solidity: function getNewbieIndex(address _newbieAddress) view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCaller) GetNewbieIndex(opts *bind.CallOpts, _newbieAddress common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "getNewbieIndex", _newbieAddress)
	return *ret0, err
}

// GetNewbieIndex is a free data retrieval call binding the contract method 0x11f5c466.
//
// Solidity: function getNewbieIndex(address _newbieAddress) view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceSession) GetNewbieIndex(_newbieAddress common.Address) (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetNewbieIndex(&_TreasuryRebalance.CallOpts, _newbieAddress)
}

// GetNewbieIndex is a free data retrieval call binding the contract method 0x11f5c466.
//
// Solidity: function getNewbieIndex(address _newbieAddress) view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) GetNewbieIndex(_newbieAddress common.Address) (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetNewbieIndex(&_TreasuryRebalance.CallOpts, _newbieAddress)
}

// GetRetired is a free data retrieval call binding the contract method 0xbf680590.
//
// Solidity: function getRetired(address _retiredAddress) view returns(address, address[])
func (_TreasuryRebalance *TreasuryRebalanceCaller) GetRetired(opts *bind.CallOpts, _retiredAddress common.Address) (common.Address, []common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TreasuryRebalance.contract.Call(opts, out, "getRetired", _retiredAddress)
	return *ret0, *ret1, err
}

// GetRetired is a free data retrieval call binding the contract method 0xbf680590.
//
// Solidity: function getRetired(address _retiredAddress) view returns(address, address[])
func (_TreasuryRebalance *TreasuryRebalanceSession) GetRetired(_retiredAddress common.Address) (common.Address, []common.Address, error) {
	return _TreasuryRebalance.Contract.GetRetired(&_TreasuryRebalance.CallOpts, _retiredAddress)
}

// GetRetired is a free data retrieval call binding the contract method 0xbf680590.
//
// Solidity: function getRetired(address _retiredAddress) view returns(address, address[])
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) GetRetired(_retiredAddress common.Address) (common.Address, []common.Address, error) {
	return _TreasuryRebalance.Contract.GetRetired(&_TreasuryRebalance.CallOpts, _retiredAddress)
}

// GetRetiredCount is a free data retrieval call binding the contract method 0xd1ed33fc.
//
// Solidity: function getRetiredCount() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCaller) GetRetiredCount(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "getRetiredCount")
	return *ret0, err
}

// GetRetiredCount is a free data retrieval call binding the contract method 0xd1ed33fc.
//
// Solidity: function getRetiredCount() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceSession) GetRetiredCount() (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetRetiredCount(&_TreasuryRebalance.CallOpts)
}

// GetRetiredCount is a free data retrieval call binding the contract method 0xd1ed33fc.
//
// Solidity: function getRetiredCount() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) GetRetiredCount() (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetRetiredCount(&_TreasuryRebalance.CallOpts)
}

// GetRetiredIndex is a free data retrieval call binding the contract method 0x681f6e7c.
//
// Solidity: function getRetiredIndex(address _retiredAddress) view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCaller) GetRetiredIndex(opts *bind.CallOpts, _retiredAddress common.Address) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "getRetiredIndex", _retiredAddress)
	return *ret0, err
}

// GetRetiredIndex is a free data retrieval call binding the contract method 0x681f6e7c.
//
// Solidity: function getRetiredIndex(address _retiredAddress) view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceSession) GetRetiredIndex(_retiredAddress common.Address) (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetRetiredIndex(&_TreasuryRebalance.CallOpts, _retiredAddress)
}

// GetRetiredIndex is a free data retrieval call binding the contract method 0x681f6e7c.
//
// Solidity: function getRetiredIndex(address _retiredAddress) view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) GetRetiredIndex(_retiredAddress common.Address) (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetRetiredIndex(&_TreasuryRebalance.CallOpts, _retiredAddress)
}

// GetTreasuryAmount is a free data retrieval call binding the contract method 0xe20fcf00.
//
// Solidity: function getTreasuryAmount() view returns(uint256 treasuryAmount)
func (_TreasuryRebalance *TreasuryRebalanceCaller) GetTreasuryAmount(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "getTreasuryAmount")
	return *ret0, err
}

// GetTreasuryAmount is a free data retrieval call binding the contract method 0xe20fcf00.
//
// Solidity: function getTreasuryAmount() view returns(uint256 treasuryAmount)
func (_TreasuryRebalance *TreasuryRebalanceSession) GetTreasuryAmount() (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetTreasuryAmount(&_TreasuryRebalance.CallOpts)
}

// GetTreasuryAmount is a free data retrieval call binding the contract method 0xe20fcf00.
//
// Solidity: function getTreasuryAmount() view returns(uint256 treasuryAmount)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) GetTreasuryAmount() (*big.Int, error) {
	return _TreasuryRebalance.Contract.GetTreasuryAmount(&_TreasuryRebalance.CallOpts)
}

// IsContractAddr is a free data retrieval call binding the contract method 0xe2384cb3.
//
// Solidity: function isContractAddr(address _addr) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCaller) IsContractAddr(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "isContractAddr", _addr)
	return *ret0, err
}

// IsContractAddr is a free data retrieval call binding the contract method 0xe2384cb3.
//
// Solidity: function isContractAddr(address _addr) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceSession) IsContractAddr(_addr common.Address) (bool, error) {
	return _TreasuryRebalance.Contract.IsContractAddr(&_TreasuryRebalance.CallOpts, _addr)
}

// IsContractAddr is a free data retrieval call binding the contract method 0xe2384cb3.
//
// Solidity: function isContractAddr(address _addr) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) IsContractAddr(_addr common.Address) (bool, error) {
	return _TreasuryRebalance.Contract.IsContractAddr(&_TreasuryRebalance.CallOpts, _addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceSession) IsOwner() (bool, error) {
	return _TreasuryRebalance.Contract.IsOwner(&_TreasuryRebalance.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) IsOwner() (bool, error) {
	return _TreasuryRebalance.Contract.IsOwner(&_TreasuryRebalance.CallOpts)
}

// Memo is a free data retrieval call binding the contract method 0x58c3b870.
//
// Solidity: function memo() view returns(string)
func (_TreasuryRebalance *TreasuryRebalanceCaller) Memo(opts *bind.CallOpts) (string, error) {
	ret0 := new(string)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "memo")
	return *ret0, err
}

// Memo is a free data retrieval call binding the contract method 0x58c3b870.
//
// Solidity: function memo() view returns(string)
func (_TreasuryRebalance *TreasuryRebalanceSession) Memo() (string, error) {
	return _TreasuryRebalance.Contract.Memo(&_TreasuryRebalance.CallOpts)
}

// Memo is a free data retrieval call binding the contract method 0x58c3b870.
//
// Solidity: function memo() view returns(string)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) Memo() (string, error) {
	return _TreasuryRebalance.Contract.Memo(&_TreasuryRebalance.CallOpts)
}

// NewbieExists is a free data retrieval call binding the contract method 0x683e13cb.
//
// Solidity: function newbieExists(address _newbieAddress) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCaller) NewbieExists(opts *bind.CallOpts, _newbieAddress common.Address) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "newbieExists", _newbieAddress)
	return *ret0, err
}

// NewbieExists is a free data retrieval call binding the contract method 0x683e13cb.
//
// Solidity: function newbieExists(address _newbieAddress) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceSession) NewbieExists(_newbieAddress common.Address) (bool, error) {
	return _TreasuryRebalance.Contract.NewbieExists(&_TreasuryRebalance.CallOpts, _newbieAddress)
}

// NewbieExists is a free data retrieval call binding the contract method 0x683e13cb.
//
// Solidity: function newbieExists(address _newbieAddress) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) NewbieExists(_newbieAddress common.Address) (bool, error) {
	return _TreasuryRebalance.Contract.NewbieExists(&_TreasuryRebalance.CallOpts, _newbieAddress)
}

// Newbies is a free data retrieval call binding the contract method 0x94393e11.
//
// Solidity: function newbies(uint256 ) view returns(address newbie, uint256 amount)
func (_TreasuryRebalance *TreasuryRebalanceCaller) Newbies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Newbie common.Address
	Amount *big.Int
}, error,
) {
	ret := new(struct {
		Newbie common.Address
		Amount *big.Int
	})
	out := ret
	err := _TreasuryRebalance.contract.Call(opts, out, "newbies", arg0)
	return *ret, err
}

// Newbies is a free data retrieval call binding the contract method 0x94393e11.
//
// Solidity: function newbies(uint256 ) view returns(address newbie, uint256 amount)
func (_TreasuryRebalance *TreasuryRebalanceSession) Newbies(arg0 *big.Int) (struct {
	Newbie common.Address
	Amount *big.Int
}, error,
) {
	return _TreasuryRebalance.Contract.Newbies(&_TreasuryRebalance.CallOpts, arg0)
}

// Newbies is a free data retrieval call binding the contract method 0x94393e11.
//
// Solidity: function newbies(uint256 ) view returns(address newbie, uint256 amount)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) Newbies(arg0 *big.Int) (struct {
	Newbie common.Address
	Amount *big.Int
}, error,
) {
	return _TreasuryRebalance.Contract.Newbies(&_TreasuryRebalance.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasuryRebalance *TreasuryRebalanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasuryRebalance *TreasuryRebalanceSession) Owner() (common.Address, error) {
	return _TreasuryRebalance.Contract.Owner(&_TreasuryRebalance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) Owner() (common.Address, error) {
	return _TreasuryRebalance.Contract.Owner(&_TreasuryRebalance.CallOpts)
}

// RebalanceBlockNumber is a free data retrieval call binding the contract method 0x49a3fb45.
//
// Solidity: function rebalanceBlockNumber() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCaller) RebalanceBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "rebalanceBlockNumber")
	return *ret0, err
}

// RebalanceBlockNumber is a free data retrieval call binding the contract method 0x49a3fb45.
//
// Solidity: function rebalanceBlockNumber() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceSession) RebalanceBlockNumber() (*big.Int, error) {
	return _TreasuryRebalance.Contract.RebalanceBlockNumber(&_TreasuryRebalance.CallOpts)
}

// RebalanceBlockNumber is a free data retrieval call binding the contract method 0x49a3fb45.
//
// Solidity: function rebalanceBlockNumber() view returns(uint256)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) RebalanceBlockNumber() (*big.Int, error) {
	return _TreasuryRebalance.Contract.RebalanceBlockNumber(&_TreasuryRebalance.CallOpts)
}

// RetiredExists is a free data retrieval call binding the contract method 0x01784e05.
//
// Solidity: function retiredExists(address _retiredAddress) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCaller) RetiredExists(opts *bind.CallOpts, _retiredAddress common.Address) (bool, error) {
	ret0 := new(bool)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "retiredExists", _retiredAddress)
	return *ret0, err
}

// RetiredExists is a free data retrieval call binding the contract method 0x01784e05.
//
// Solidity: function retiredExists(address _retiredAddress) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceSession) RetiredExists(_retiredAddress common.Address) (bool, error) {
	return _TreasuryRebalance.Contract.RetiredExists(&_TreasuryRebalance.CallOpts, _retiredAddress)
}

// RetiredExists is a free data retrieval call binding the contract method 0x01784e05.
//
// Solidity: function retiredExists(address _retiredAddress) view returns(bool)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) RetiredExists(_retiredAddress common.Address) (bool, error) {
	return _TreasuryRebalance.Contract.RetiredExists(&_TreasuryRebalance.CallOpts, _retiredAddress)
}

// Retirees is a free data retrieval call binding the contract method 0x5a12667b.
//
// Solidity: function retirees(uint256 ) view returns(address retired)
func (_TreasuryRebalance *TreasuryRebalanceCaller) Retirees(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "retirees", arg0)
	return *ret0, err
}

// Retirees is a free data retrieval call binding the contract method 0x5a12667b.
//
// Solidity: function retirees(uint256 ) view returns(address retired)
func (_TreasuryRebalance *TreasuryRebalanceSession) Retirees(arg0 *big.Int) (common.Address, error) {
	return _TreasuryRebalance.Contract.Retirees(&_TreasuryRebalance.CallOpts, arg0)
}

// Retirees is a free data retrieval call binding the contract method 0x5a12667b.
//
// Solidity: function retirees(uint256 ) view returns(address retired)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) Retirees(arg0 *big.Int) (common.Address, error) {
	return _TreasuryRebalance.Contract.Retirees(&_TreasuryRebalance.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_TreasuryRebalance *TreasuryRebalanceCaller) Status(opts *bind.CallOpts) (uint8, error) {
	ret0 := new(uint8)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_TreasuryRebalance *TreasuryRebalanceSession) Status() (uint8, error) {
	return _TreasuryRebalance.Contract.Status(&_TreasuryRebalance.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) Status() (uint8, error) {
	return _TreasuryRebalance.Contract.Status(&_TreasuryRebalance.CallOpts)
}

// SumOfRetiredBalance is a free data retrieval call binding the contract method 0x45205a6b.
//
// Solidity: function sumOfRetiredBalance() view returns(uint256 retireesBalance)
func (_TreasuryRebalance *TreasuryRebalanceCaller) SumOfRetiredBalance(opts *bind.CallOpts) (*big.Int, error) {
	ret0 := new(*big.Int)
	out := ret0
	err := _TreasuryRebalance.contract.Call(opts, out, "sumOfRetiredBalance")
	return *ret0, err
}

// SumOfRetiredBalance is a free data retrieval call binding the contract method 0x45205a6b.
//
// Solidity: function sumOfRetiredBalance() view returns(uint256 retireesBalance)
func (_TreasuryRebalance *TreasuryRebalanceSession) SumOfRetiredBalance() (*big.Int, error) {
	return _TreasuryRebalance.Contract.SumOfRetiredBalance(&_TreasuryRebalance.CallOpts)
}

// SumOfRetiredBalance is a free data retrieval call binding the contract method 0x45205a6b.
//
// Solidity: function sumOfRetiredBalance() view returns(uint256 retireesBalance)
func (_TreasuryRebalance *TreasuryRebalanceCallerSession) SumOfRetiredBalance() (*big.Int, error) {
	return _TreasuryRebalance.Contract.SumOfRetiredBalance(&_TreasuryRebalance.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) Approve(opts *bind.TransactOpts, _retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "approve", _retiredAddress)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) Approve(_retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.Approve(&_TreasuryRebalance.TransactOpts, _retiredAddress)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) Approve(_retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.Approve(&_TreasuryRebalance.TransactOpts, _retiredAddress)
}

// FinalizeApproval is a paid mutator transaction binding the contract method 0xfaaf9ca6.
//
// Solidity: function finalizeApproval() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) FinalizeApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "finalizeApproval")
}

// FinalizeApproval is a paid mutator transaction binding the contract method 0xfaaf9ca6.
//
// Solidity: function finalizeApproval() returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) FinalizeApproval() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.FinalizeApproval(&_TreasuryRebalance.TransactOpts)
}

// FinalizeApproval is a paid mutator transaction binding the contract method 0xfaaf9ca6.
//
// Solidity: function finalizeApproval() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) FinalizeApproval() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.FinalizeApproval(&_TreasuryRebalance.TransactOpts)
}

// FinalizeContract is a paid mutator transaction binding the contract method 0xea6d4a9b.
//
// Solidity: function finalizeContract(string _memo) returns(address[] retirees, uint256 totalRetireesBalance, address[] newbies, uint256 totalNewbiesFund)
func (_TreasuryRebalance *TreasuryRebalanceTransactor) FinalizeContract(opts *bind.TransactOpts, _memo string) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "finalizeContract", _memo)
}

// FinalizeContract is a paid mutator transaction binding the contract method 0xea6d4a9b.
//
// Solidity: function finalizeContract(string _memo) returns(address[] retirees, uint256 totalRetireesBalance, address[] newbies, uint256 totalNewbiesFund)
func (_TreasuryRebalance *TreasuryRebalanceSession) FinalizeContract(_memo string) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.FinalizeContract(&_TreasuryRebalance.TransactOpts, _memo)
}

// FinalizeContract is a paid mutator transaction binding the contract method 0xea6d4a9b.
//
// Solidity: function finalizeContract(string _memo) returns(address[] retirees, uint256 totalRetireesBalance, address[] newbies, uint256 totalNewbiesFund)
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) FinalizeContract(_memo string) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.FinalizeContract(&_TreasuryRebalance.TransactOpts, _memo)
}

// FinalizeRegistration is a paid mutator transaction binding the contract method 0x48409096.
//
// Solidity: function finalizeRegistration() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) FinalizeRegistration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "finalizeRegistration")
}

// FinalizeRegistration is a paid mutator transaction binding the contract method 0x48409096.
//
// Solidity: function finalizeRegistration() returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) FinalizeRegistration() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.FinalizeRegistration(&_TreasuryRebalance.TransactOpts)
}

// FinalizeRegistration is a paid mutator transaction binding the contract method 0x48409096.
//
// Solidity: function finalizeRegistration() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) FinalizeRegistration() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.FinalizeRegistration(&_TreasuryRebalance.TransactOpts)
}

// RegisterNewbie is a paid mutator transaction binding the contract method 0x652e27e0.
//
// Solidity: function registerNewbie(address _newbieAddress, uint256 _amount) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) RegisterNewbie(opts *bind.TransactOpts, _newbieAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "registerNewbie", _newbieAddress, _amount)
}

// RegisterNewbie is a paid mutator transaction binding the contract method 0x652e27e0.
//
// Solidity: function registerNewbie(address _newbieAddress, uint256 _amount) returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) RegisterNewbie(_newbieAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RegisterNewbie(&_TreasuryRebalance.TransactOpts, _newbieAddress, _amount)
}

// RegisterNewbie is a paid mutator transaction binding the contract method 0x652e27e0.
//
// Solidity: function registerNewbie(address _newbieAddress, uint256 _amount) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) RegisterNewbie(_newbieAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RegisterNewbie(&_TreasuryRebalance.TransactOpts, _newbieAddress, _amount)
}

// RegisterRetired is a paid mutator transaction binding the contract method 0x1f8c1798.
//
// Solidity: function registerRetired(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) RegisterRetired(opts *bind.TransactOpts, _retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "registerRetired", _retiredAddress)
}

// RegisterRetired is a paid mutator transaction binding the contract method 0x1f8c1798.
//
// Solidity: function registerRetired(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) RegisterRetired(_retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RegisterRetired(&_TreasuryRebalance.TransactOpts, _retiredAddress)
}

// RegisterRetired is a paid mutator transaction binding the contract method 0x1f8c1798.
//
// Solidity: function registerRetired(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) RegisterRetired(_retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RegisterRetired(&_TreasuryRebalance.TransactOpts, _retiredAddress)
}

// RemoveNewbie is a paid mutator transaction binding the contract method 0x6864b95b.
//
// Solidity: function removeNewbie(address _newbieAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) RemoveNewbie(opts *bind.TransactOpts, _newbieAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "removeNewbie", _newbieAddress)
}

// RemoveNewbie is a paid mutator transaction binding the contract method 0x6864b95b.
//
// Solidity: function removeNewbie(address _newbieAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) RemoveNewbie(_newbieAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RemoveNewbie(&_TreasuryRebalance.TransactOpts, _newbieAddress)
}

// RemoveNewbie is a paid mutator transaction binding the contract method 0x6864b95b.
//
// Solidity: function removeNewbie(address _newbieAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) RemoveNewbie(_newbieAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RemoveNewbie(&_TreasuryRebalance.TransactOpts, _newbieAddress)
}

// RemoveRetired is a paid mutator transaction binding the contract method 0x1c1dac59.
//
// Solidity: function removeRetired(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) RemoveRetired(opts *bind.TransactOpts, _retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "removeRetired", _retiredAddress)
}

// RemoveRetired is a paid mutator transaction binding the contract method 0x1c1dac59.
//
// Solidity: function removeRetired(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) RemoveRetired(_retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RemoveRetired(&_TreasuryRebalance.TransactOpts, _retiredAddress)
}

// RemoveRetired is a paid mutator transaction binding the contract method 0x1c1dac59.
//
// Solidity: function removeRetired(address _retiredAddress) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) RemoveRetired(_retiredAddress common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RemoveRetired(&_TreasuryRebalance.TransactOpts, _retiredAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RenounceOwnership(&_TreasuryRebalance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.RenounceOwnership(&_TreasuryRebalance.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) Reset() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.Reset(&_TreasuryRebalance.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) Reset() (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.Reset(&_TreasuryRebalance.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.TransferOwnership(&_TreasuryRebalance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.TransferOwnership(&_TreasuryRebalance.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TreasuryRebalance.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TreasuryRebalance *TreasuryRebalanceSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.Fallback(&_TreasuryRebalance.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TreasuryRebalance *TreasuryRebalanceTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TreasuryRebalance.Contract.Fallback(&_TreasuryRebalance.TransactOpts, calldata)
}

// TreasuryRebalanceApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the TreasuryRebalance contract.
type TreasuryRebalanceApprovedIterator struct {
	Event *TreasuryRebalanceApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceApproved)
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
		it.Event = new(TreasuryRebalanceApproved)
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
func (it *TreasuryRebalanceApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceApproved represents a Approved event raised by the TreasuryRebalance contract.
type TreasuryRebalanceApproved struct {
	Retired        common.Address
	Approver       common.Address
	ApproversCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x80da462ebfbe41cfc9bc015e7a9a3c7a2a73dbccede72d8ceb583606c27f8f90.
//
// Solidity: event Approved(address retired, address approver, uint256 approversCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterApproved(opts *bind.FilterOpts) (*TreasuryRebalanceApprovedIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "Approved")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceApprovedIterator{contract: _TreasuryRebalance.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0x80da462ebfbe41cfc9bc015e7a9a3c7a2a73dbccede72d8ceb583606c27f8f90.
//
// Solidity: event Approved(address retired, address approver, uint256 approversCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceApproved) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "Approved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceApproved)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0x80da462ebfbe41cfc9bc015e7a9a3c7a2a73dbccede72d8ceb583606c27f8f90.
//
// Solidity: event Approved(address retired, address approver, uint256 approversCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseApproved(log types.Log) (*TreasuryRebalanceApproved, error) {
	event := new(TreasuryRebalanceApproved)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceContractDeployedIterator is returned from FilterContractDeployed and is used to iterate over the raw logs and unpacked data for ContractDeployed events raised by the TreasuryRebalance contract.
type TreasuryRebalanceContractDeployedIterator struct {
	Event *TreasuryRebalanceContractDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceContractDeployed)
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
		it.Event = new(TreasuryRebalanceContractDeployed)
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
func (it *TreasuryRebalanceContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceContractDeployed represents a ContractDeployed event raised by the TreasuryRebalance contract.
type TreasuryRebalanceContractDeployed struct {
	Status               uint8
	RebalanceBlockNumber *big.Int
	DeployedBlockNumber  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0x6f182006c5a12fe70c0728eedb2d1b0628c41483ca6721c606707d778d22ed0a.
//
// Solidity: event ContractDeployed(uint8 status, uint256 rebalanceBlockNumber, uint256 deployedBlockNumber)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterContractDeployed(opts *bind.FilterOpts) (*TreasuryRebalanceContractDeployedIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceContractDeployedIterator{contract: _TreasuryRebalance.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0x6f182006c5a12fe70c0728eedb2d1b0628c41483ca6721c606707d778d22ed0a.
//
// Solidity: event ContractDeployed(uint8 status, uint256 rebalanceBlockNumber, uint256 deployedBlockNumber)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceContractDeployed) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceContractDeployed)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
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

// ParseContractDeployed is a log parse operation binding the contract event 0x6f182006c5a12fe70c0728eedb2d1b0628c41483ca6721c606707d778d22ed0a.
//
// Solidity: event ContractDeployed(uint8 status, uint256 rebalanceBlockNumber, uint256 deployedBlockNumber)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseContractDeployed(log types.Log) (*TreasuryRebalanceContractDeployed, error) {
	event := new(TreasuryRebalanceContractDeployed)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the TreasuryRebalance contract.
type TreasuryRebalanceFinalizedIterator struct {
	Event *TreasuryRebalanceFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceFinalized)
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
		it.Event = new(TreasuryRebalanceFinalized)
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
func (it *TreasuryRebalanceFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceFinalized represents a Finalized event raised by the TreasuryRebalance contract.
type TreasuryRebalanceFinalized struct {
	Memo   string
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x8f8636c7757ca9b7d154e1d44ca90d8e8c885b9eac417c59bbf8eb7779ca6404.
//
// Solidity: event Finalized(string memo, uint8 status)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterFinalized(opts *bind.FilterOpts) (*TreasuryRebalanceFinalizedIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceFinalizedIterator{contract: _TreasuryRebalance.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x8f8636c7757ca9b7d154e1d44ca90d8e8c885b9eac417c59bbf8eb7779ca6404.
//
// Solidity: event Finalized(string memo, uint8 status)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceFinalized) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceFinalized)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// ParseFinalized is a log parse operation binding the contract event 0x8f8636c7757ca9b7d154e1d44ca90d8e8c885b9eac417c59bbf8eb7779ca6404.
//
// Solidity: event Finalized(string memo, uint8 status)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseFinalized(log types.Log) (*TreasuryRebalanceFinalized, error) {
	event := new(TreasuryRebalanceFinalized)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "Finalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceNewbieRegisteredIterator is returned from FilterNewbieRegistered and is used to iterate over the raw logs and unpacked data for NewbieRegistered events raised by the TreasuryRebalance contract.
type TreasuryRebalanceNewbieRegisteredIterator struct {
	Event *TreasuryRebalanceNewbieRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceNewbieRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceNewbieRegistered)
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
		it.Event = new(TreasuryRebalanceNewbieRegistered)
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
func (it *TreasuryRebalanceNewbieRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceNewbieRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceNewbieRegistered represents a NewbieRegistered event raised by the TreasuryRebalance contract.
type TreasuryRebalanceNewbieRegistered struct {
	Newbie         common.Address
	FundAllocation *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewbieRegistered is a free log retrieval operation binding the contract event 0xd261b37cd56b21cd1af841dca6331a133e5d8b9d55c2c6fe0ec822e2a303ef74.
//
// Solidity: event NewbieRegistered(address newbie, uint256 fundAllocation)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterNewbieRegistered(opts *bind.FilterOpts) (*TreasuryRebalanceNewbieRegisteredIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "NewbieRegistered")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceNewbieRegisteredIterator{contract: _TreasuryRebalance.contract, event: "NewbieRegistered", logs: logs, sub: sub}, nil
}

// WatchNewbieRegistered is a free log subscription operation binding the contract event 0xd261b37cd56b21cd1af841dca6331a133e5d8b9d55c2c6fe0ec822e2a303ef74.
//
// Solidity: event NewbieRegistered(address newbie, uint256 fundAllocation)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchNewbieRegistered(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceNewbieRegistered) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "NewbieRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceNewbieRegistered)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "NewbieRegistered", log); err != nil {
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

// ParseNewbieRegistered is a log parse operation binding the contract event 0xd261b37cd56b21cd1af841dca6331a133e5d8b9d55c2c6fe0ec822e2a303ef74.
//
// Solidity: event NewbieRegistered(address newbie, uint256 fundAllocation)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseNewbieRegistered(log types.Log) (*TreasuryRebalanceNewbieRegistered, error) {
	event := new(TreasuryRebalanceNewbieRegistered)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "NewbieRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceNewbieRemovedIterator is returned from FilterNewbieRemoved and is used to iterate over the raw logs and unpacked data for NewbieRemoved events raised by the TreasuryRebalance contract.
type TreasuryRebalanceNewbieRemovedIterator struct {
	Event *TreasuryRebalanceNewbieRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceNewbieRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceNewbieRemoved)
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
		it.Event = new(TreasuryRebalanceNewbieRemoved)
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
func (it *TreasuryRebalanceNewbieRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceNewbieRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceNewbieRemoved represents a NewbieRemoved event raised by the TreasuryRebalance contract.
type TreasuryRebalanceNewbieRemoved struct {
	Newbie      common.Address
	NewbieCount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewbieRemoved is a free log retrieval operation binding the contract event 0x16c9a1659c2467bd64845207a0990d56629de7a3955c0f93c19354cda213c932.
//
// Solidity: event NewbieRemoved(address newbie, uint256 newbieCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterNewbieRemoved(opts *bind.FilterOpts) (*TreasuryRebalanceNewbieRemovedIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "NewbieRemoved")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceNewbieRemovedIterator{contract: _TreasuryRebalance.contract, event: "NewbieRemoved", logs: logs, sub: sub}, nil
}

// WatchNewbieRemoved is a free log subscription operation binding the contract event 0x16c9a1659c2467bd64845207a0990d56629de7a3955c0f93c19354cda213c932.
//
// Solidity: event NewbieRemoved(address newbie, uint256 newbieCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchNewbieRemoved(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceNewbieRemoved) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "NewbieRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceNewbieRemoved)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "NewbieRemoved", log); err != nil {
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

// ParseNewbieRemoved is a log parse operation binding the contract event 0x16c9a1659c2467bd64845207a0990d56629de7a3955c0f93c19354cda213c932.
//
// Solidity: event NewbieRemoved(address newbie, uint256 newbieCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseNewbieRemoved(log types.Log) (*TreasuryRebalanceNewbieRemoved, error) {
	event := new(TreasuryRebalanceNewbieRemoved)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "NewbieRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TreasuryRebalance contract.
type TreasuryRebalanceOwnershipTransferredIterator struct {
	Event *TreasuryRebalanceOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceOwnershipTransferred)
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
		it.Event = new(TreasuryRebalanceOwnershipTransferred)
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
func (it *TreasuryRebalanceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceOwnershipTransferred represents a OwnershipTransferred event raised by the TreasuryRebalance contract.
type TreasuryRebalanceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TreasuryRebalanceOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceOwnershipTransferredIterator{contract: _TreasuryRebalance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceOwnershipTransferred)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseOwnershipTransferred(log types.Log) (*TreasuryRebalanceOwnershipTransferred, error) {
	event := new(TreasuryRebalanceOwnershipTransferred)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceRetiredRegisteredIterator is returned from FilterRetiredRegistered and is used to iterate over the raw logs and unpacked data for RetiredRegistered events raised by the TreasuryRebalance contract.
type TreasuryRebalanceRetiredRegisteredIterator struct {
	Event *TreasuryRebalanceRetiredRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceRetiredRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceRetiredRegistered)
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
		it.Event = new(TreasuryRebalanceRetiredRegistered)
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
func (it *TreasuryRebalanceRetiredRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceRetiredRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceRetiredRegistered represents a RetiredRegistered event raised by the TreasuryRebalance contract.
type TreasuryRebalanceRetiredRegistered struct {
	Retired common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRetiredRegistered is a free log retrieval operation binding the contract event 0x7da2e87d0b02df1162d5736cc40dfcfffd17198aaf093ddff4a8f4eb26002fde.
//
// Solidity: event RetiredRegistered(address retired)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterRetiredRegistered(opts *bind.FilterOpts) (*TreasuryRebalanceRetiredRegisteredIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "RetiredRegistered")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceRetiredRegisteredIterator{contract: _TreasuryRebalance.contract, event: "RetiredRegistered", logs: logs, sub: sub}, nil
}

// WatchRetiredRegistered is a free log subscription operation binding the contract event 0x7da2e87d0b02df1162d5736cc40dfcfffd17198aaf093ddff4a8f4eb26002fde.
//
// Solidity: event RetiredRegistered(address retired)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchRetiredRegistered(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceRetiredRegistered) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "RetiredRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceRetiredRegistered)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "RetiredRegistered", log); err != nil {
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

// ParseRetiredRegistered is a log parse operation binding the contract event 0x7da2e87d0b02df1162d5736cc40dfcfffd17198aaf093ddff4a8f4eb26002fde.
//
// Solidity: event RetiredRegistered(address retired)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseRetiredRegistered(log types.Log) (*TreasuryRebalanceRetiredRegistered, error) {
	event := new(TreasuryRebalanceRetiredRegistered)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "RetiredRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceRetiredRemovedIterator is returned from FilterRetiredRemoved and is used to iterate over the raw logs and unpacked data for RetiredRemoved events raised by the TreasuryRebalance contract.
type TreasuryRebalanceRetiredRemovedIterator struct {
	Event *TreasuryRebalanceRetiredRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceRetiredRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceRetiredRemoved)
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
		it.Event = new(TreasuryRebalanceRetiredRemoved)
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
func (it *TreasuryRebalanceRetiredRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceRetiredRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceRetiredRemoved represents a RetiredRemoved event raised by the TreasuryRebalance contract.
type TreasuryRebalanceRetiredRemoved struct {
	Retired      common.Address
	RetiredCount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRetiredRemoved is a free log retrieval operation binding the contract event 0x13673a8c5e648e7bebdeb5f081555ca1843dba99063e9e55aa6399e61573c70d.
//
// Solidity: event RetiredRemoved(address retired, uint256 retiredCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterRetiredRemoved(opts *bind.FilterOpts) (*TreasuryRebalanceRetiredRemovedIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "RetiredRemoved")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceRetiredRemovedIterator{contract: _TreasuryRebalance.contract, event: "RetiredRemoved", logs: logs, sub: sub}, nil
}

// WatchRetiredRemoved is a free log subscription operation binding the contract event 0x13673a8c5e648e7bebdeb5f081555ca1843dba99063e9e55aa6399e61573c70d.
//
// Solidity: event RetiredRemoved(address retired, uint256 retiredCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchRetiredRemoved(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceRetiredRemoved) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "RetiredRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceRetiredRemoved)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "RetiredRemoved", log); err != nil {
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

// ParseRetiredRemoved is a log parse operation binding the contract event 0x13673a8c5e648e7bebdeb5f081555ca1843dba99063e9e55aa6399e61573c70d.
//
// Solidity: event RetiredRemoved(address retired, uint256 retiredCount)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseRetiredRemoved(log types.Log) (*TreasuryRebalanceRetiredRemoved, error) {
	event := new(TreasuryRebalanceRetiredRemoved)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "RetiredRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TreasuryRebalanceStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the TreasuryRebalance contract.
type TreasuryRebalanceStatusChangedIterator struct {
	Event *TreasuryRebalanceStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TreasuryRebalanceStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryRebalanceStatusChanged)
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
		it.Event = new(TreasuryRebalanceStatusChanged)
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
func (it *TreasuryRebalanceStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryRebalanceStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryRebalanceStatusChanged represents a StatusChanged event raised by the TreasuryRebalance contract.
type TreasuryRebalanceStatusChanged struct {
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0xafa725e7f44cadb687a7043853fa1a7e7b8f0da74ce87ec546e9420f04da8c1e.
//
// Solidity: event StatusChanged(uint8 status)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) FilterStatusChanged(opts *bind.FilterOpts) (*TreasuryRebalanceStatusChangedIterator, error) {
	logs, sub, err := _TreasuryRebalance.contract.FilterLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return &TreasuryRebalanceStatusChangedIterator{contract: _TreasuryRebalance.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0xafa725e7f44cadb687a7043853fa1a7e7b8f0da74ce87ec546e9420f04da8c1e.
//
// Solidity: event StatusChanged(uint8 status)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *TreasuryRebalanceStatusChanged) (event.Subscription, error) {
	logs, sub, err := _TreasuryRebalance.contract.WatchLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryRebalanceStatusChanged)
				if err := _TreasuryRebalance.contract.UnpackLog(event, "StatusChanged", log); err != nil {
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

// ParseStatusChanged is a log parse operation binding the contract event 0xafa725e7f44cadb687a7043853fa1a7e7b8f0da74ce87ec546e9420f04da8c1e.
//
// Solidity: event StatusChanged(uint8 status)
func (_TreasuryRebalance *TreasuryRebalanceFilterer) ParseStatusChanged(log types.Log) (*TreasuryRebalanceStatusChanged, error) {
	event := new(TreasuryRebalanceStatusChanged)
	if err := _TreasuryRebalance.contract.UnpackLog(event, "StatusChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
