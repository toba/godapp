// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BookmarkABI is the input ABI used to generate the binding from.
const BookmarkABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"show\",\"type\":\"string\"}],\"name\":\"bookmark\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBookmarks\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BookmarkBin is the compiled bytecode used for deploying new contracts.
const BookmarkBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a0316179055610402806100396000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166327243160811461005257806345d9bd151461011a5780638da5cb5b1461012d57600080fd5b341561005d57600080fd5b6100a360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061015c95505050505050565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100df5780820151838201526020016100c7565b50505050905090810190601f16801561010c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561012557600080fd5b6100a3610252565b341561013857600080fd5b61014061031d565b604051600160a060020a03909116815260200160405180910390f35b61016461032c565b600160a060020a033316600090815260208190526040902082805161018d92916020019061033e565b5060008033600160a060020a0316600160a060020a031681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102465780601f1061021b57610100808354040283529160200191610246565b820191906000526020600020905b81548152906001019060200180831161022957829003601f168201915b50505050509050919050565b61025a61032c565b60008033600160a060020a0316600160a060020a031681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103125780601f106102e757610100808354040283529160200191610312565b820191906000526020600020905b8154815290600101906020018083116102f557829003601f168201915b505050505090505b90565b600154600160a060020a031681565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061037f57805160ff19168380011785556103ac565b828001600101855582156103ac579182015b828111156103ac578251825591602001919060010190610391565b506103b89291506103bc565b5090565b61031a91905b808211156103b857600081556001016103c25600a165627a7a7230582003839fc975d2ca16e90ce75077b53a20a4dbe5d032999b70a7b8cb84f3b883590029`

// DeployBookmark deploys a new Ethereum contract, binding an instance of Bookmark to it.
func DeployBookmark(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bookmark, error) {
	parsed, err := abi.JSON(strings.NewReader(BookmarkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BookmarkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bookmark{BookmarkCaller: BookmarkCaller{contract: contract}, BookmarkTransactor: BookmarkTransactor{contract: contract}}, nil
}

// Bookmark is an auto generated Go binding around an Ethereum contract.
type Bookmark struct {
	BookmarkCaller     // Read-only binding to the contract
	BookmarkTransactor // Write-only binding to the contract
}

// BookmarkCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookmarkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookmarkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookmarkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookmarkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookmarkSession struct {
	Contract     *Bookmark         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BookmarkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookmarkCallerSession struct {
	Contract *BookmarkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BookmarkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookmarkTransactorSession struct {
	Contract     *BookmarkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BookmarkRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookmarkRaw struct {
	Contract *Bookmark // Generic contract binding to access the raw methods on
}

// BookmarkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookmarkCallerRaw struct {
	Contract *BookmarkCaller // Generic read-only contract binding to access the raw methods on
}

// BookmarkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookmarkTransactorRaw struct {
	Contract *BookmarkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBookmark creates a new instance of Bookmark, bound to a specific deployed contract.
func NewBookmark(address common.Address, backend bind.ContractBackend) (*Bookmark, error) {
	contract, err := bindBookmark(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bookmark{BookmarkCaller: BookmarkCaller{contract: contract}, BookmarkTransactor: BookmarkTransactor{contract: contract}}, nil
}

// NewBookmarkCaller creates a new read-only instance of Bookmark, bound to a specific deployed contract.
func NewBookmarkCaller(address common.Address, caller bind.ContractCaller) (*BookmarkCaller, error) {
	contract, err := bindBookmark(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BookmarkCaller{contract: contract}, nil
}

// NewBookmarkTransactor creates a new write-only instance of Bookmark, bound to a specific deployed contract.
func NewBookmarkTransactor(address common.Address, transactor bind.ContractTransactor) (*BookmarkTransactor, error) {
	contract, err := bindBookmark(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BookmarkTransactor{contract: contract}, nil
}

// bindBookmark binds a generic wrapper to an already deployed contract.
func bindBookmark(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BookmarkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookmark *BookmarkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bookmark.Contract.BookmarkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookmark *BookmarkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookmark.Contract.BookmarkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookmark *BookmarkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookmark.Contract.BookmarkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookmark *BookmarkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bookmark.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookmark *BookmarkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookmark.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookmark *BookmarkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookmark.Contract.contract.Transact(opts, method, params...)
}

// GetBookmarks is a free data retrieval call binding the contract method 0x45d9bd15.
//
// Solidity: function getBookmarks() constant returns(string)
func (_Bookmark *BookmarkCaller) GetBookmarks(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bookmark.contract.Call(opts, out, "getBookmarks")
	return *ret0, err
}

// GetBookmarks is a free data retrieval call binding the contract method 0x45d9bd15.
//
// Solidity: function getBookmarks() constant returns(string)
func (_Bookmark *BookmarkSession) GetBookmarks() (string, error) {
	return _Bookmark.Contract.GetBookmarks(&_Bookmark.CallOpts)
}

// GetBookmarks is a free data retrieval call binding the contract method 0x45d9bd15.
//
// Solidity: function getBookmarks() constant returns(string)
func (_Bookmark *BookmarkCallerSession) GetBookmarks() (string, error) {
	return _Bookmark.Contract.GetBookmarks(&_Bookmark.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bookmark *BookmarkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bookmark.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bookmark *BookmarkSession) Owner() (common.Address, error) {
	return _Bookmark.Contract.Owner(&_Bookmark.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bookmark *BookmarkCallerSession) Owner() (common.Address, error) {
	return _Bookmark.Contract.Owner(&_Bookmark.CallOpts)
}

// Bookmark is a paid mutator transaction binding the contract method 0x27243160.
//
// Solidity: function bookmark(show string) returns(string)
func (_Bookmark *BookmarkTransactor) Bookmark(opts *bind.TransactOpts, show string) (*types.Transaction, error) {
	return _Bookmark.contract.Transact(opts, "bookmark", show)
}

// Bookmark is a paid mutator transaction binding the contract method 0x27243160.
//
// Solidity: function bookmark(show string) returns(string)
func (_Bookmark *BookmarkSession) Bookmark(show string) (*types.Transaction, error) {
	return _Bookmark.Contract.Bookmark(&_Bookmark.TransactOpts, show)
}

// Bookmark is a paid mutator transaction binding the contract method 0x27243160.
//
// Solidity: function bookmark(show string) returns(string)
func (_Bookmark *BookmarkTransactorSession) Bookmark(show string) (*types.Transaction, error) {
	return _Bookmark.Contract.Bookmark(&_Bookmark.TransactOpts, show)
}
