// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gov

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

// GovParamParamView is an auto generated low-level Go binding around an user-defined struct.
type GovParamParamView struct {
	Name  string
	Value []byte
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// ContextBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ContextBinRuntime = ``

// Context is an auto generated Go binding around a Klaytn contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around a Klaytn contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around a Klaytn contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// GovParamABI is the input ABI used to generate the binding from.
const GovParamABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"ParamVotableUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SetParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"VotableUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"VoteContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"addParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getAllParams\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structGovParam.ParamView[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getParam\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_activationBlock\",\"type\":\"uint64\"}],\"name\":\"setParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovParamBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const GovParamBinRuntime = `608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100c8578063c147b8c2146100e3578063f2fde38b14610103578063f548319b1461011657600080fd5b80632c974cb914610082578063715018a614610097578063783df7361461009f575b600080fd5b610095610090366004610bd1565b610129565b005b610095610384565b6100b26100ad366004610c66565b6103ba565b6040516100bf9190610ccc565b60405180910390f35b6000546040516001600160a01b0390911681526020016100bf565b6100f66100f1366004610d51565b6106c0565b6040516100bf9190610d73565b610095610111366004610d8d565b6107a9565b610095610124366004610e42565b610844565b8361013c6000546001600160a01b031690565b6001600160a01b0316336001600160a01b0316146101955760405162461bcd60e51b81526020600482015260116024820152701c195c9b5a5cdcda5bdb8819195b9a5959607a1b60448201526064015b60405180910390fd5b600085815260016020819052604090912001544367ffffffffffffffff909116106102025760405162461bcd60e51b815260206004820152601d60248201527f616c7265616479206861766520612070656e64696e67206368616e6765000000604482015260640161018c565b8167ffffffffffffffff1643106102665760405162461bcd60e51b815260206004820152602260248201527f63616e6e6f74207365742061637469766174696f6e426c6f636b20746f2070616044820152611cdd60f21b606482015260840161018c565b6000858152600160205260408120805461027f90610eda565b9050116102c25760405162461bcd60e51b815260206004820152601160248201527037379039bab1b4103830b930b6b2ba32b960791b604482015260640161018c565b60008581526001602052604090206003810180546002909201916102e590610eda565b6102f0929190610a49565b506000858152600160208190526040909120908101805467ffffffffffffffff191667ffffffffffffffff851617905561032e906003018585610ad4565b506000858152600160205260409081902090517fd7d6c246fce475d3f08d03e87447c7ee332e6b50ade11b9259f47b4b60541df69161037591889190889088908890610f3e565b60405180910390a15050505050565b6000546001600160a01b031633146103ae5760405162461bcd60e51b815260040161018c9061101c565b6103b860006109f9565b565b60025460609060009067ffffffffffffffff8111156103db576103db610db6565b60405190808252806020026020018201604052801561042057816020015b60408051808201909152606080825260208201528152602001906001900390816103f95790505b50905060005b6002548110156106b9576000600160006002848154811061044957610449611051565b9060005260206000200154815260200190815260200160002060405180608001604052908160008201805461047d90610eda565b80601f01602080910402602001604051908101604052809291908181526020018280546104a990610eda565b80156104f65780601f106104cb576101008083540402835291602001916104f6565b820191906000526020600020905b8154815290600101906020018083116104d957829003601f168201915b5050509183525050600182015467ffffffffffffffff16602082015260028201805460409092019161052790610eda565b80601f016020809104026020016040519081016040528092919081815260200182805461055390610eda565b80156105a05780601f10610575576101008083540402835291602001916105a0565b820191906000526020600020905b81548152906001019060200180831161058357829003601f168201915b505050505081526020016003820180546105b990610eda565b80601f01602080910402602001604051908101604052809291908181526020018280546105e590610eda565b80156106325780601f1061060757610100808354040283529160200191610632565b820191906000526020600020905b81548152906001019060200180831161061557829003601f168201915b50505050508152505090506000610666866002858154811061065657610656611051565b90600052602060002001546106c0565b905060006040518060400160405280846000015181526020018381525090508085858151811061069857610698611051565b602002602001018190525050505080806106b190611067565b915050610426565b5092915050565b6000818152600160208190526040909120015460609067ffffffffffffffff168310610787576000828152600160205260409020600301805461070290610eda565b80601f016020809104026020016040519081016040528092919081815260200182805461072e90610eda565b801561077b5780601f106107505761010080835404028352916020019161077b565b820191906000526020600020905b81548152906001019060200180831161075e57829003601f168201915b505050505090506107a3565b6000828152600160205260409020600201805461070290610eda565b92915050565b6000546001600160a01b031633146107d35760405162461bcd60e51b815260040161018c9061101c565b6001600160a01b0381166108385760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161018c565b610841816109f9565b50565b6000546001600160a01b0316331461086e5760405162461bcd60e51b815260040161018c9061101c565b60008251116108b65760405162461bcd60e51b81526020600482015260146024820152736e616d652063616e6e6f7420626520656d70747960601b604482015260640161018c565b600083815260016020526040902080546108cf90610eda565b1590506109145760405162461bcd60e51b8152602060048201526013602482015272185b1c9958591e48195e1a5cdd1a5b99c81a59606a1b604482015260640161018c565b604080516080810182528381526000602080830182905283518082018552828152838501526060830185905286825260018152929020815180519293919261095f9284920190610b48565b5060208281015160018301805467ffffffffffffffff191667ffffffffffffffff909216919091179055604083015180516109a09260028501920190610b48565b50606082015180516109bc916003840191602090910190610b48565b5050600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0193909355505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054610a5590610eda565b90600052602060002090601f016020900481019282610a775760008555610ac4565b82601f10610a885780548555610ac4565b82800160010185558215610ac457600052602060002091601f016020900482015b82811115610ac4578254825591600101919060010190610aa9565b50610ad0929150610bbc565b5090565b828054610ae090610eda565b90600052602060002090601f016020900481019282610b025760008555610ac4565b82601f10610b1b5782800160ff19823516178555610ac4565b82800160010185558215610ac4579182015b82811115610ac4578235825591602001919060010190610b2d565b828054610b5490610eda565b90600052602060002090601f016020900481019282610b765760008555610ac4565b82601f10610b8f57805160ff1916838001178555610ac4565b82800160010185558215610ac4579182015b82811115610ac4578251825591602001919060010190610ba1565b5b80821115610ad05760008155600101610bbd565b60008060008060608587031215610be757600080fd5b84359350602085013567ffffffffffffffff80821115610c0657600080fd5b818701915087601f830112610c1a57600080fd5b813581811115610c2957600080fd5b886020828501011115610c3b57600080fd5b60208301955080945050604087013591508082168214610c5a57600080fd5b50939692955090935050565b600060208284031215610c7857600080fd5b5035919050565b6000815180845260005b81811015610ca557602081850181015186830182015201610c89565b81811115610cb7576000602083870101525b50601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015610d4357888303603f1901855281518051878552610d1788860182610c7f565b91890151858303868b0152919050610d2f8183610c7f565b968901969450505090860190600101610cf3565b509098975050505050505050565b60008060408385031215610d6457600080fd5b50508035926020909101359150565b602081526000610d866020830184610c7f565b9392505050565b600060208284031215610d9f57600080fd5b81356001600160a01b0381168114610d8657600080fd5b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115610de757610de7610db6565b604051601f8501601f19908116603f01168101908282118183101715610e0f57610e0f610db6565b81604052809350858152868686011115610e2857600080fd5b858560208301376000602087830101525050509392505050565b600080600060608486031215610e5757600080fd5b83359250602084013567ffffffffffffffff80821115610e7657600080fd5b818601915086601f830112610e8a57600080fd5b610e9987833560208501610dcc565b93506040860135915080821115610eaf57600080fd5b508401601f81018613610ec157600080fd5b610ed086823560208401610dcc565b9150509250925092565b600181811c90821680610eee57607f821691505b60208210811415610f0f57634e487b7160e01b600052602260045260246000fd5b50919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b858152600060206080818401526000875481600182811c915080831680610f6657607f831692505b858310811415610f8457634e487b7160e01b85526022600452602485fd5b6080880183905260a08801818015610fa35760018114610fb457610fdf565b60ff19861682528782019650610fdf565b60008e81526020902060005b86811015610fd957815484820152908501908901610fc0565b83019750505b5050505050508381036040850152610ff8818789610f15565b92505050611012606083018467ffffffffffffffff169052565b9695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600060001982141561108957634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122040d13b2da3dd1c6ab037174cc7023bf2c2f203b407b6b83ae65b22e06087035864736f6c634300080b0033`

// GovParamFuncSigs maps the 4-byte function signature to its string representation.
var GovParamFuncSigs = map[string]string{
	"f548319b": "addParam(uint256,string,bytes)",
	"783df736": "getAllParams(uint256)",
	"c147b8c2": "getParam(uint256,uint256)",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"2c974cb9": "setParam(uint256,bytes,uint64)",
	"f2fde38b": "transferOwnership(address)",
}

// GovParamBin is the compiled bytecode used for deploying new contracts.
var GovParamBin = "0x608060405234801561001057600080fd5b506040516111ab3803806111ab83398101604081905261002f916100a6565b61003833610056565b6001600160a01b038116156100505761005081610056565b506100d6565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100b857600080fd5b81516001600160a01b03811681146100cf57600080fd5b9392505050565b6110c6806100e56000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100c8578063c147b8c2146100e3578063f2fde38b14610103578063f548319b1461011657600080fd5b80632c974cb914610082578063715018a614610097578063783df7361461009f575b600080fd5b610095610090366004610bd1565b610129565b005b610095610384565b6100b26100ad366004610c66565b6103ba565b6040516100bf9190610ccc565b60405180910390f35b6000546040516001600160a01b0390911681526020016100bf565b6100f66100f1366004610d51565b6106c0565b6040516100bf9190610d73565b610095610111366004610d8d565b6107a9565b610095610124366004610e42565b610844565b8361013c6000546001600160a01b031690565b6001600160a01b0316336001600160a01b0316146101955760405162461bcd60e51b81526020600482015260116024820152701c195c9b5a5cdcda5bdb8819195b9a5959607a1b60448201526064015b60405180910390fd5b600085815260016020819052604090912001544367ffffffffffffffff909116106102025760405162461bcd60e51b815260206004820152601d60248201527f616c7265616479206861766520612070656e64696e67206368616e6765000000604482015260640161018c565b8167ffffffffffffffff1643106102665760405162461bcd60e51b815260206004820152602260248201527f63616e6e6f74207365742061637469766174696f6e426c6f636b20746f2070616044820152611cdd60f21b606482015260840161018c565b6000858152600160205260408120805461027f90610eda565b9050116102c25760405162461bcd60e51b815260206004820152601160248201527037379039bab1b4103830b930b6b2ba32b960791b604482015260640161018c565b60008581526001602052604090206003810180546002909201916102e590610eda565b6102f0929190610a49565b506000858152600160208190526040909120908101805467ffffffffffffffff191667ffffffffffffffff851617905561032e906003018585610ad4565b506000858152600160205260409081902090517fd7d6c246fce475d3f08d03e87447c7ee332e6b50ade11b9259f47b4b60541df69161037591889190889088908890610f3e565b60405180910390a15050505050565b6000546001600160a01b031633146103ae5760405162461bcd60e51b815260040161018c9061101c565b6103b860006109f9565b565b60025460609060009067ffffffffffffffff8111156103db576103db610db6565b60405190808252806020026020018201604052801561042057816020015b60408051808201909152606080825260208201528152602001906001900390816103f95790505b50905060005b6002548110156106b9576000600160006002848154811061044957610449611051565b9060005260206000200154815260200190815260200160002060405180608001604052908160008201805461047d90610eda565b80601f01602080910402602001604051908101604052809291908181526020018280546104a990610eda565b80156104f65780601f106104cb576101008083540402835291602001916104f6565b820191906000526020600020905b8154815290600101906020018083116104d957829003601f168201915b5050509183525050600182015467ffffffffffffffff16602082015260028201805460409092019161052790610eda565b80601f016020809104026020016040519081016040528092919081815260200182805461055390610eda565b80156105a05780601f10610575576101008083540402835291602001916105a0565b820191906000526020600020905b81548152906001019060200180831161058357829003601f168201915b505050505081526020016003820180546105b990610eda565b80601f01602080910402602001604051908101604052809291908181526020018280546105e590610eda565b80156106325780601f1061060757610100808354040283529160200191610632565b820191906000526020600020905b81548152906001019060200180831161061557829003601f168201915b50505050508152505090506000610666866002858154811061065657610656611051565b90600052602060002001546106c0565b905060006040518060400160405280846000015181526020018381525090508085858151811061069857610698611051565b602002602001018190525050505080806106b190611067565b915050610426565b5092915050565b6000818152600160208190526040909120015460609067ffffffffffffffff168310610787576000828152600160205260409020600301805461070290610eda565b80601f016020809104026020016040519081016040528092919081815260200182805461072e90610eda565b801561077b5780601f106107505761010080835404028352916020019161077b565b820191906000526020600020905b81548152906001019060200180831161075e57829003601f168201915b505050505090506107a3565b6000828152600160205260409020600201805461070290610eda565b92915050565b6000546001600160a01b031633146107d35760405162461bcd60e51b815260040161018c9061101c565b6001600160a01b0381166108385760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161018c565b610841816109f9565b50565b6000546001600160a01b0316331461086e5760405162461bcd60e51b815260040161018c9061101c565b60008251116108b65760405162461bcd60e51b81526020600482015260146024820152736e616d652063616e6e6f7420626520656d70747960601b604482015260640161018c565b600083815260016020526040902080546108cf90610eda565b1590506109145760405162461bcd60e51b8152602060048201526013602482015272185b1c9958591e48195e1a5cdd1a5b99c81a59606a1b604482015260640161018c565b604080516080810182528381526000602080830182905283518082018552828152838501526060830185905286825260018152929020815180519293919261095f9284920190610b48565b5060208281015160018301805467ffffffffffffffff191667ffffffffffffffff909216919091179055604083015180516109a09260028501920190610b48565b50606082015180516109bc916003840191602090910190610b48565b5050600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0193909355505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054610a5590610eda565b90600052602060002090601f016020900481019282610a775760008555610ac4565b82601f10610a885780548555610ac4565b82800160010185558215610ac457600052602060002091601f016020900482015b82811115610ac4578254825591600101919060010190610aa9565b50610ad0929150610bbc565b5090565b828054610ae090610eda565b90600052602060002090601f016020900481019282610b025760008555610ac4565b82601f10610b1b5782800160ff19823516178555610ac4565b82800160010185558215610ac4579182015b82811115610ac4578235825591602001919060010190610b2d565b828054610b5490610eda565b90600052602060002090601f016020900481019282610b765760008555610ac4565b82601f10610b8f57805160ff1916838001178555610ac4565b82800160010185558215610ac4579182015b82811115610ac4578251825591602001919060010190610ba1565b5b80821115610ad05760008155600101610bbd565b60008060008060608587031215610be757600080fd5b84359350602085013567ffffffffffffffff80821115610c0657600080fd5b818701915087601f830112610c1a57600080fd5b813581811115610c2957600080fd5b886020828501011115610c3b57600080fd5b60208301955080945050604087013591508082168214610c5a57600080fd5b50939692955090935050565b600060208284031215610c7857600080fd5b5035919050565b6000815180845260005b81811015610ca557602081850181015186830182015201610c89565b81811115610cb7576000602083870101525b50601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015610d4357888303603f1901855281518051878552610d1788860182610c7f565b91890151858303868b0152919050610d2f8183610c7f565b968901969450505090860190600101610cf3565b509098975050505050505050565b60008060408385031215610d6457600080fd5b50508035926020909101359150565b602081526000610d866020830184610c7f565b9392505050565b600060208284031215610d9f57600080fd5b81356001600160a01b0381168114610d8657600080fd5b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115610de757610de7610db6565b604051601f8501601f19908116603f01168101908282118183101715610e0f57610e0f610db6565b81604052809350858152868686011115610e2857600080fd5b858560208301376000602087830101525050509392505050565b600080600060608486031215610e5757600080fd5b83359250602084013567ffffffffffffffff80821115610e7657600080fd5b818601915086601f830112610e8a57600080fd5b610e9987833560208501610dcc565b93506040860135915080821115610eaf57600080fd5b508401601f81018613610ec157600080fd5b610ed086823560208401610dcc565b9150509250925092565b600181811c90821680610eee57607f821691505b60208210811415610f0f57634e487b7160e01b600052602260045260246000fd5b50919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b858152600060206080818401526000875481600182811c915080831680610f6657607f831692505b858310811415610f8457634e487b7160e01b85526022600452602485fd5b6080880183905260a08801818015610fa35760018114610fb457610fdf565b60ff19861682528782019650610fdf565b60008e81526020902060005b86811015610fd957815484820152908501908901610fc0565b83019750505b5050505050508381036040850152610ff8818789610f15565b92505050611012606083018467ffffffffffffffff169052565b9695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600060001982141561108957634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122040d13b2da3dd1c6ab037174cc7023bf2c2f203b407b6b83ae65b22e06087035864736f6c634300080b0033"

// DeployGovParam deploys a new Klaytn contract, binding an instance of GovParam to it.
func DeployGovParam(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *GovParam, error) {
	parsed, err := abi.JSON(strings.NewReader(GovParamABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovParamBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovParam{GovParamCaller: GovParamCaller{contract: contract}, GovParamTransactor: GovParamTransactor{contract: contract}, GovParamFilterer: GovParamFilterer{contract: contract}}, nil
}

// GovParam is an auto generated Go binding around a Klaytn contract.
type GovParam struct {
	GovParamCaller     // Read-only binding to the contract
	GovParamTransactor // Write-only binding to the contract
	GovParamFilterer   // Log filterer for contract events
}

// GovParamCaller is an auto generated read-only Go binding around a Klaytn contract.
type GovParamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovParamTransactor is an auto generated write-only Go binding around a Klaytn contract.
type GovParamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovParamFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type GovParamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovParamSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type GovParamSession struct {
	Contract     *GovParam         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovParamCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type GovParamCallerSession struct {
	Contract *GovParamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GovParamTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type GovParamTransactorSession struct {
	Contract     *GovParamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GovParamRaw is an auto generated low-level Go binding around a Klaytn contract.
type GovParamRaw struct {
	Contract *GovParam // Generic contract binding to access the raw methods on
}

// GovParamCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type GovParamCallerRaw struct {
	Contract *GovParamCaller // Generic read-only contract binding to access the raw methods on
}

// GovParamTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type GovParamTransactorRaw struct {
	Contract *GovParamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovParam creates a new instance of GovParam, bound to a specific deployed contract.
func NewGovParam(address common.Address, backend bind.ContractBackend) (*GovParam, error) {
	contract, err := bindGovParam(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovParam{GovParamCaller: GovParamCaller{contract: contract}, GovParamTransactor: GovParamTransactor{contract: contract}, GovParamFilterer: GovParamFilterer{contract: contract}}, nil
}

// NewGovParamCaller creates a new read-only instance of GovParam, bound to a specific deployed contract.
func NewGovParamCaller(address common.Address, caller bind.ContractCaller) (*GovParamCaller, error) {
	contract, err := bindGovParam(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovParamCaller{contract: contract}, nil
}

// NewGovParamTransactor creates a new write-only instance of GovParam, bound to a specific deployed contract.
func NewGovParamTransactor(address common.Address, transactor bind.ContractTransactor) (*GovParamTransactor, error) {
	contract, err := bindGovParam(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovParamTransactor{contract: contract}, nil
}

// NewGovParamFilterer creates a new log filterer instance of GovParam, bound to a specific deployed contract.
func NewGovParamFilterer(address common.Address, filterer bind.ContractFilterer) (*GovParamFilterer, error) {
	contract, err := bindGovParam(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovParamFilterer{contract: contract}, nil
}

// bindGovParam binds a generic wrapper to an already deployed contract.
func bindGovParam(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovParamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovParam *GovParamRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovParam.Contract.GovParamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovParam *GovParamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovParam.Contract.GovParamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovParam *GovParamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovParam.Contract.GovParamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovParam *GovParamCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovParam.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovParam *GovParamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovParam.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovParam *GovParamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovParam.Contract.contract.Transact(opts, method, params...)
}

// GetAllParams is a free data retrieval call binding the contract method 0x783df736.
//
// Solidity: function getAllParams(uint256 n) view returns((string,bytes)[])
func (_GovParam *GovParamCaller) GetAllParams(opts *bind.CallOpts, n *big.Int) ([]GovParamParamView, error) {
	var (
		ret0 = new([]GovParamParamView)
	)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "getAllParams", n)
	return *ret0, err
}

// GetAllParams is a free data retrieval call binding the contract method 0x783df736.
//
// Solidity: function getAllParams(uint256 n) view returns((string,bytes)[])
func (_GovParam *GovParamSession) GetAllParams(n *big.Int) ([]GovParamParamView, error) {
	return _GovParam.Contract.GetAllParams(&_GovParam.CallOpts, n)
}

// GetAllParams is a free data retrieval call binding the contract method 0x783df736.
//
// Solidity: function getAllParams(uint256 n) view returns((string,bytes)[])
func (_GovParam *GovParamCallerSession) GetAllParams(n *big.Int) ([]GovParamParamView, error) {
	return _GovParam.Contract.GetAllParams(&_GovParam.CallOpts, n)
}

// GetParam is a free data retrieval call binding the contract method 0xc147b8c2.
//
// Solidity: function getParam(uint256 n, uint256 id) view returns(bytes)
func (_GovParam *GovParamCaller) GetParam(opts *bind.CallOpts, n *big.Int, id *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "getParam", n, id)
	return *ret0, err
}

// GetParam is a free data retrieval call binding the contract method 0xc147b8c2.
//
// Solidity: function getParam(uint256 n, uint256 id) view returns(bytes)
func (_GovParam *GovParamSession) GetParam(n *big.Int, id *big.Int) ([]byte, error) {
	return _GovParam.Contract.GetParam(&_GovParam.CallOpts, n, id)
}

// GetParam is a free data retrieval call binding the contract method 0xc147b8c2.
//
// Solidity: function getParam(uint256 n, uint256 id) view returns(bytes)
func (_GovParam *GovParamCallerSession) GetParam(n *big.Int, id *big.Int) ([]byte, error) {
	return _GovParam.Contract.GetParam(&_GovParam.CallOpts, n, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovParam *GovParamCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovParam *GovParamSession) Owner() (common.Address, error) {
	return _GovParam.Contract.Owner(&_GovParam.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovParam *GovParamCallerSession) Owner() (common.Address, error) {
	return _GovParam.Contract.Owner(&_GovParam.CallOpts)
}

// AddParam is a paid mutator transaction binding the contract method 0xf548319b.
//
// Solidity: function addParam(uint256 id, string _name, bytes value) returns()
func (_GovParam *GovParamTransactor) AddParam(opts *bind.TransactOpts, id *big.Int, _name string, value []byte) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "addParam", id, _name, value)
}

// AddParam is a paid mutator transaction binding the contract method 0xf548319b.
//
// Solidity: function addParam(uint256 id, string _name, bytes value) returns()
func (_GovParam *GovParamSession) AddParam(id *big.Int, _name string, value []byte) (*types.Transaction, error) {
	return _GovParam.Contract.AddParam(&_GovParam.TransactOpts, id, _name, value)
}

// AddParam is a paid mutator transaction binding the contract method 0xf548319b.
//
// Solidity: function addParam(uint256 id, string _name, bytes value) returns()
func (_GovParam *GovParamTransactorSession) AddParam(id *big.Int, _name string, value []byte) (*types.Transaction, error) {
	return _GovParam.Contract.AddParam(&_GovParam.TransactOpts, id, _name, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovParam *GovParamTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovParam *GovParamSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovParam.Contract.RenounceOwnership(&_GovParam.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovParam *GovParamTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovParam.Contract.RenounceOwnership(&_GovParam.TransactOpts)
}

// SetParam is a paid mutator transaction binding the contract method 0x2c974cb9.
//
// Solidity: function setParam(uint256 id, bytes value, uint64 _activationBlock) returns()
func (_GovParam *GovParamTransactor) SetParam(opts *bind.TransactOpts, id *big.Int, value []byte, _activationBlock uint64) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "setParam", id, value, _activationBlock)
}

// SetParam is a paid mutator transaction binding the contract method 0x2c974cb9.
//
// Solidity: function setParam(uint256 id, bytes value, uint64 _activationBlock) returns()
func (_GovParam *GovParamSession) SetParam(id *big.Int, value []byte, _activationBlock uint64) (*types.Transaction, error) {
	return _GovParam.Contract.SetParam(&_GovParam.TransactOpts, id, value, _activationBlock)
}

// SetParam is a paid mutator transaction binding the contract method 0x2c974cb9.
//
// Solidity: function setParam(uint256 id, bytes value, uint64 _activationBlock) returns()
func (_GovParam *GovParamTransactorSession) SetParam(id *big.Int, value []byte, _activationBlock uint64) (*types.Transaction, error) {
	return _GovParam.Contract.SetParam(&_GovParam.TransactOpts, id, value, _activationBlock)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovParam *GovParamTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovParam *GovParamSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.TransferOwnership(&_GovParam.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovParam *GovParamTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.TransferOwnership(&_GovParam.TransactOpts, newOwner)
}

// GovParamOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GovParam contract.
type GovParamOwnershipTransferredIterator struct {
	Event *GovParamOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovParamOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamOwnershipTransferred)
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
		it.Event = new(GovParamOwnershipTransferred)
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
func (it *GovParamOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamOwnershipTransferred represents a OwnershipTransferred event raised by the GovParam contract.
type GovParamOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovParam *GovParamFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovParamOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovParamOwnershipTransferredIterator{contract: _GovParam.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovParam *GovParamFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovParamOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamOwnershipTransferred)
				if err := _GovParam.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GovParam *GovParamFilterer) ParseOwnershipTransferred(log types.Log) (*GovParamOwnershipTransferred, error) {
	event := new(GovParamOwnershipTransferred)
	if err := _GovParam.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamParamVotableUpdatedIterator is returned from FilterParamVotableUpdated and is used to iterate over the raw logs and unpacked data for ParamVotableUpdated events raised by the GovParam contract.
type GovParamParamVotableUpdatedIterator struct {
	Event *GovParamParamVotableUpdated // Event containing the contract specifics and raw log

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
func (it *GovParamParamVotableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamParamVotableUpdated)
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
		it.Event = new(GovParamParamVotableUpdated)
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
func (it *GovParamParamVotableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamParamVotableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamParamVotableUpdated represents a ParamVotableUpdated event raised by the GovParam contract.
type GovParamParamVotableUpdated struct {
	Arg0 *big.Int
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterParamVotableUpdated is a free log retrieval operation binding the contract event 0x474fcc0988d951be0d382e04d07c3709137e8cb05389453a1badc188f1c2ffc3.
//
// Solidity: event ParamVotableUpdated(uint256 arg0, bool arg1)
func (_GovParam *GovParamFilterer) FilterParamVotableUpdated(opts *bind.FilterOpts) (*GovParamParamVotableUpdatedIterator, error) {

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "ParamVotableUpdated")
	if err != nil {
		return nil, err
	}
	return &GovParamParamVotableUpdatedIterator{contract: _GovParam.contract, event: "ParamVotableUpdated", logs: logs, sub: sub}, nil
}

// WatchParamVotableUpdated is a free log subscription operation binding the contract event 0x474fcc0988d951be0d382e04d07c3709137e8cb05389453a1badc188f1c2ffc3.
//
// Solidity: event ParamVotableUpdated(uint256 arg0, bool arg1)
func (_GovParam *GovParamFilterer) WatchParamVotableUpdated(opts *bind.WatchOpts, sink chan<- *GovParamParamVotableUpdated) (event.Subscription, error) {

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "ParamVotableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamParamVotableUpdated)
				if err := _GovParam.contract.UnpackLog(event, "ParamVotableUpdated", log); err != nil {
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

// ParseParamVotableUpdated is a log parse operation binding the contract event 0x474fcc0988d951be0d382e04d07c3709137e8cb05389453a1badc188f1c2ffc3.
//
// Solidity: event ParamVotableUpdated(uint256 arg0, bool arg1)
func (_GovParam *GovParamFilterer) ParseParamVotableUpdated(log types.Log) (*GovParamParamVotableUpdated, error) {
	event := new(GovParamParamVotableUpdated)
	if err := _GovParam.contract.UnpackLog(event, "ParamVotableUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamSetParamIterator is returned from FilterSetParam and is used to iterate over the raw logs and unpacked data for SetParam events raised by the GovParam contract.
type GovParamSetParamIterator struct {
	Event *GovParamSetParam // Event containing the contract specifics and raw log

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
func (it *GovParamSetParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamSetParam)
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
		it.Event = new(GovParamSetParam)
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
func (it *GovParamSetParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamSetParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamSetParam represents a SetParam event raised by the GovParam contract.
type GovParamSetParam struct {
	Arg0 *big.Int
	Arg1 string
	Arg2 []byte
	Arg3 uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetParam is a free log retrieval operation binding the contract event 0xd7d6c246fce475d3f08d03e87447c7ee332e6b50ade11b9259f47b4b60541df6.
//
// Solidity: event SetParam(uint256 arg0, string arg1, bytes arg2, uint64 arg3)
func (_GovParam *GovParamFilterer) FilterSetParam(opts *bind.FilterOpts) (*GovParamSetParamIterator, error) {

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "SetParam")
	if err != nil {
		return nil, err
	}
	return &GovParamSetParamIterator{contract: _GovParam.contract, event: "SetParam", logs: logs, sub: sub}, nil
}

// WatchSetParam is a free log subscription operation binding the contract event 0xd7d6c246fce475d3f08d03e87447c7ee332e6b50ade11b9259f47b4b60541df6.
//
// Solidity: event SetParam(uint256 arg0, string arg1, bytes arg2, uint64 arg3)
func (_GovParam *GovParamFilterer) WatchSetParam(opts *bind.WatchOpts, sink chan<- *GovParamSetParam) (event.Subscription, error) {

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "SetParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamSetParam)
				if err := _GovParam.contract.UnpackLog(event, "SetParam", log); err != nil {
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

// ParseSetParam is a log parse operation binding the contract event 0xd7d6c246fce475d3f08d03e87447c7ee332e6b50ade11b9259f47b4b60541df6.
//
// Solidity: event SetParam(uint256 arg0, string arg1, bytes arg2, uint64 arg3)
func (_GovParam *GovParamFilterer) ParseSetParam(log types.Log) (*GovParamSetParam, error) {
	event := new(GovParamSetParam)
	if err := _GovParam.contract.UnpackLog(event, "SetParam", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the GovParam contract.
type GovParamValidatorAddedIterator struct {
	Event *GovParamValidatorAdded // Event containing the contract specifics and raw log

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
func (it *GovParamValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamValidatorAdded)
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
		it.Event = new(GovParamValidatorAdded)
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
func (it *GovParamValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamValidatorAdded represents a ValidatorAdded event raised by the GovParam contract.
type GovParamValidatorAdded struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address arg0)
func (_GovParam *GovParamFilterer) FilterValidatorAdded(opts *bind.FilterOpts) (*GovParamValidatorAddedIterator, error) {

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &GovParamValidatorAddedIterator{contract: _GovParam.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address arg0)
func (_GovParam *GovParamFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *GovParamValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamValidatorAdded)
				if err := _GovParam.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address arg0)
func (_GovParam *GovParamFilterer) ParseValidatorAdded(log types.Log) (*GovParamValidatorAdded, error) {
	event := new(GovParamValidatorAdded)
	if err := _GovParam.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the GovParam contract.
type GovParamValidatorRemovedIterator struct {
	Event *GovParamValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *GovParamValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamValidatorRemoved)
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
		it.Event = new(GovParamValidatorRemoved)
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
func (it *GovParamValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamValidatorRemoved represents a ValidatorRemoved event raised by the GovParam contract.
type GovParamValidatorRemoved struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address arg0)
func (_GovParam *GovParamFilterer) FilterValidatorRemoved(opts *bind.FilterOpts) (*GovParamValidatorRemovedIterator, error) {

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &GovParamValidatorRemovedIterator{contract: _GovParam.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address arg0)
func (_GovParam *GovParamFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *GovParamValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamValidatorRemoved)
				if err := _GovParam.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address arg0)
func (_GovParam *GovParamFilterer) ParseValidatorRemoved(log types.Log) (*GovParamValidatorRemoved, error) {
	event := new(GovParamValidatorRemoved)
	if err := _GovParam.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamVotableUpdatedIterator is returned from FilterVotableUpdated and is used to iterate over the raw logs and unpacked data for VotableUpdated events raised by the GovParam contract.
type GovParamVotableUpdatedIterator struct {
	Event *GovParamVotableUpdated // Event containing the contract specifics and raw log

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
func (it *GovParamVotableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamVotableUpdated)
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
		it.Event = new(GovParamVotableUpdated)
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
func (it *GovParamVotableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamVotableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamVotableUpdated represents a VotableUpdated event raised by the GovParam contract.
type GovParamVotableUpdated struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterVotableUpdated is a free log retrieval operation binding the contract event 0xc561407216cee7a48319239d9317e6f025c76442add64869c3229fcc2cc532fa.
//
// Solidity: event VotableUpdated(bool arg0)
func (_GovParam *GovParamFilterer) FilterVotableUpdated(opts *bind.FilterOpts) (*GovParamVotableUpdatedIterator, error) {

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "VotableUpdated")
	if err != nil {
		return nil, err
	}
	return &GovParamVotableUpdatedIterator{contract: _GovParam.contract, event: "VotableUpdated", logs: logs, sub: sub}, nil
}

// WatchVotableUpdated is a free log subscription operation binding the contract event 0xc561407216cee7a48319239d9317e6f025c76442add64869c3229fcc2cc532fa.
//
// Solidity: event VotableUpdated(bool arg0)
func (_GovParam *GovParamFilterer) WatchVotableUpdated(opts *bind.WatchOpts, sink chan<- *GovParamVotableUpdated) (event.Subscription, error) {

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "VotableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamVotableUpdated)
				if err := _GovParam.contract.UnpackLog(event, "VotableUpdated", log); err != nil {
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

// ParseVotableUpdated is a log parse operation binding the contract event 0xc561407216cee7a48319239d9317e6f025c76442add64869c3229fcc2cc532fa.
//
// Solidity: event VotableUpdated(bool arg0)
func (_GovParam *GovParamFilterer) ParseVotableUpdated(log types.Log) (*GovParamVotableUpdated, error) {
	event := new(GovParamVotableUpdated)
	if err := _GovParam.contract.UnpackLog(event, "VotableUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamVoteContractUpdatedIterator is returned from FilterVoteContractUpdated and is used to iterate over the raw logs and unpacked data for VoteContractUpdated events raised by the GovParam contract.
type GovParamVoteContractUpdatedIterator struct {
	Event *GovParamVoteContractUpdated // Event containing the contract specifics and raw log

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
func (it *GovParamVoteContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamVoteContractUpdated)
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
		it.Event = new(GovParamVoteContractUpdated)
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
func (it *GovParamVoteContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamVoteContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamVoteContractUpdated represents a VoteContractUpdated event raised by the GovParam contract.
type GovParamVoteContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterVoteContractUpdated is a free log retrieval operation binding the contract event 0xbb74b2013639f069748049dfe8b049c50e999babc72c1754c7bcd39c6a93920b.
//
// Solidity: event VoteContractUpdated(address arg0)
func (_GovParam *GovParamFilterer) FilterVoteContractUpdated(opts *bind.FilterOpts) (*GovParamVoteContractUpdatedIterator, error) {

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "VoteContractUpdated")
	if err != nil {
		return nil, err
	}
	return &GovParamVoteContractUpdatedIterator{contract: _GovParam.contract, event: "VoteContractUpdated", logs: logs, sub: sub}, nil
}

// WatchVoteContractUpdated is a free log subscription operation binding the contract event 0xbb74b2013639f069748049dfe8b049c50e999babc72c1754c7bcd39c6a93920b.
//
// Solidity: event VoteContractUpdated(address arg0)
func (_GovParam *GovParamFilterer) WatchVoteContractUpdated(opts *bind.WatchOpts, sink chan<- *GovParamVoteContractUpdated) (event.Subscription, error) {

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "VoteContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamVoteContractUpdated)
				if err := _GovParam.contract.UnpackLog(event, "VoteContractUpdated", log); err != nil {
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

// ParseVoteContractUpdated is a log parse operation binding the contract event 0xbb74b2013639f069748049dfe8b049c50e999babc72c1754c7bcd39c6a93920b.
//
// Solidity: event VoteContractUpdated(address arg0)
func (_GovParam *GovParamFilterer) ParseVoteContractUpdated(log types.Log) (*GovParamVoteContractUpdated, error) {
	event := new(GovParamVoteContractUpdated)
	if err := _GovParam.contract.UnpackLog(event, "VoteContractUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OwnableBinRuntime = ``

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
