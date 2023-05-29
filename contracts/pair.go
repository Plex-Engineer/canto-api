// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pair

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

// PairMetaData contains all meta data concerning the Pair contract.
var PairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"_k\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"observationLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"window\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"window\",\"type\":\"uint256\"}],\"name\":\"sampleReserves\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"window\",\"type\":\"uint256\"}],\"name\":\"sampleSupply\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PairABI is the input ABI used to generate the binding from.
// Deprecated: Use PairMetaData.ABI instead.
var PairABI = PairMetaData.ABI

// Pair is an auto generated Go binding around an Ethereum contract.
type Pair struct {
	PairCaller     // Read-only binding to the contract
	PairTransactor // Write-only binding to the contract
	PairFilterer   // Log filterer for contract events
}

// PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairSession struct {
	Contract     *Pair             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairCallerSession struct {
	Contract *PairCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairTransactorSession struct {
	Contract     *PairTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairRaw struct {
	Contract *Pair // Generic contract binding to access the raw methods on
}

// PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairCallerRaw struct {
	Contract *PairCaller // Generic read-only contract binding to access the raw methods on
}

// PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairTransactorRaw struct {
	Contract *PairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPair creates a new instance of Pair, bound to a specific deployed contract.
func NewPair(address common.Address, backend bind.ContractBackend) (*Pair, error) {
	contract, err := bindPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pair{PairCaller: PairCaller{contract: contract}, PairTransactor: PairTransactor{contract: contract}, PairFilterer: PairFilterer{contract: contract}}, nil
}

// NewPairCaller creates a new read-only instance of Pair, bound to a specific deployed contract.
func NewPairCaller(address common.Address, caller bind.ContractCaller) (*PairCaller, error) {
	contract, err := bindPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairCaller{contract: contract}, nil
}

// NewPairTransactor creates a new write-only instance of Pair, bound to a specific deployed contract.
func NewPairTransactor(address common.Address, transactor bind.ContractTransactor) (*PairTransactor, error) {
	contract, err := bindPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairTransactor{contract: contract}, nil
}

// NewPairFilterer creates a new log filterer instance of Pair, bound to a specific deployed contract.
func NewPairFilterer(address common.Address, filterer bind.ContractFilterer) (*PairFilterer, error) {
	contract, err := bindPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairFilterer{contract: contract}, nil
}

// bindPair binds a generic wrapper to an already deployed contract.
func bindPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pair.Contract.PairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.Contract.PairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pair.Contract.PairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pair.Contract.contract.Transact(opts, method, params...)
}

// K is a free data retrieval call binding the contract method 0x9cf42614.
//
// Solidity: function _k(uint256 x, uint256 y) view returns(uint256)
func (_Pair *PairCaller) K(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "_k", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0x9cf42614.
//
// Solidity: function _k(uint256 x, uint256 y) view returns(uint256)
func (_Pair *PairSession) K(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Pair.Contract.K(&_Pair.CallOpts, x, y)
}

// K is a free data retrieval call binding the contract method 0x9cf42614.
//
// Solidity: function _k(uint256 x, uint256 y) view returns(uint256)
func (_Pair *PairCallerSession) K(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Pair.Contract.K(&_Pair.CallOpts, x, y)
}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256)
func (_Pair *PairCaller) Current(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "current", tokenIn, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256)
func (_Pair *PairSession) Current(tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Pair.Contract.Current(&_Pair.CallOpts, tokenIn, amountIn)
}

// Current is a free data retrieval call binding the contract method 0x517b3f82.
//
// Solidity: function current(address tokenIn, uint256 amountIn) view returns(uint256)
func (_Pair *PairCallerSession) Current(tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Pair.Contract.Current(&_Pair.CallOpts, tokenIn, amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 , address ) view returns(uint256)
func (_Pair *PairCaller) GetAmountOut(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "getAmountOut", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 , address ) view returns(uint256)
func (_Pair *PairSession) GetAmountOut(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Pair.Contract.GetAmountOut(&_Pair.CallOpts, arg0, arg1)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xf140a35a.
//
// Solidity: function getAmountOut(uint256 , address ) view returns(uint256)
func (_Pair *PairCallerSession) GetAmountOut(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Pair.Contract.GetAmountOut(&_Pair.CallOpts, arg0, arg1)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pair *PairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pair *PairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pair *PairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_Pair *PairCaller) ObservationLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "observationLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_Pair *PairSession) ObservationLength() (*big.Int, error) {
	return _Pair.Contract.ObservationLength(&_Pair.CallOpts)
}

// ObservationLength is a free data retrieval call binding the contract method 0xebeb31db.
//
// Solidity: function observationLength() view returns(uint256)
func (_Pair *PairCallerSession) ObservationLength() (*big.Int, error) {
	return _Pair.Contract.ObservationLength(&_Pair.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256)
func (_Pair *PairCaller) Quote(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "quote", tokenIn, amountIn, granularity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256)
func (_Pair *PairSession) Quote(tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	return _Pair.Contract.Quote(&_Pair.CallOpts, tokenIn, amountIn, granularity)
}

// Quote is a free data retrieval call binding the contract method 0x9e8cc04b.
//
// Solidity: function quote(address tokenIn, uint256 amountIn, uint256 granularity) view returns(uint256)
func (_Pair *PairCallerSession) Quote(tokenIn common.Address, amountIn *big.Int, granularity *big.Int) (*big.Int, error) {
	return _Pair.Contract.Quote(&_Pair.CallOpts, tokenIn, amountIn, granularity)
}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairCaller) Sample(opts *bind.CallOpts, tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "sample", tokenIn, amountIn, points, window)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairSession) Sample(tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.Sample(&_Pair.CallOpts, tokenIn, amountIn, points, window)
}

// Sample is a free data retrieval call binding the contract method 0x13345fe1.
//
// Solidity: function sample(address tokenIn, uint256 amountIn, uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairCallerSession) Sample(tokenIn common.Address, amountIn *big.Int, points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.Sample(&_Pair.CallOpts, tokenIn, amountIn, points, window)
}

// SampleReserves is a free data retrieval call binding the contract method 0xf99f51a6.
//
// Solidity: function sampleReserves(uint256 points, uint256 window) view returns(uint256[], uint256[])
func (_Pair *PairCaller) SampleReserves(opts *bind.CallOpts, points *big.Int, window *big.Int) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "sampleReserves", points, window)

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// SampleReserves is a free data retrieval call binding the contract method 0xf99f51a6.
//
// Solidity: function sampleReserves(uint256 points, uint256 window) view returns(uint256[], uint256[])
func (_Pair *PairSession) SampleReserves(points *big.Int, window *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Pair.Contract.SampleReserves(&_Pair.CallOpts, points, window)
}

// SampleReserves is a free data retrieval call binding the contract method 0xf99f51a6.
//
// Solidity: function sampleReserves(uint256 points, uint256 window) view returns(uint256[], uint256[])
func (_Pair *PairCallerSession) SampleReserves(points *big.Int, window *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Pair.Contract.SampleReserves(&_Pair.CallOpts, points, window)
}

// SampleSupply is a free data retrieval call binding the contract method 0xeba6aef9.
//
// Solidity: function sampleSupply(uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairCaller) SampleSupply(opts *bind.CallOpts, points *big.Int, window *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "sampleSupply", points, window)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SampleSupply is a free data retrieval call binding the contract method 0xeba6aef9.
//
// Solidity: function sampleSupply(uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairSession) SampleSupply(points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.SampleSupply(&_Pair.CallOpts, points, window)
}

// SampleSupply is a free data retrieval call binding the contract method 0xeba6aef9.
//
// Solidity: function sampleSupply(uint256 points, uint256 window) view returns(uint256[])
func (_Pair *PairCallerSession) SampleSupply(points *big.Int, window *big.Int) ([]*big.Int, error) {
	return _Pair.Contract.SampleSupply(&_Pair.CallOpts, points, window)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_Pair *PairCaller) Stable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_Pair *PairSession) Stable() (bool, error) {
	return _Pair.Contract.Stable(&_Pair.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(bool)
func (_Pair *PairCallerSession) Stable() (bool, error) {
	return _Pair.Contract.Stable(&_Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairSession) Token0() (common.Address, error) {
	return _Pair.Contract.Token0(&_Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCallerSession) Token0() (common.Address, error) {
	return _Pair.Contract.Token0(&_Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairSession) Token1() (common.Address, error) {
	return _Pair.Contract.Token1(&_Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCallerSession) Token1() (common.Address, error) {
	return _Pair.Contract.Token1(&_Pair.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Pair *PairTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Pair *PairSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Pair *PairTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, src, dst, amount)
}
