// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pickrouter

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

// PickrouterMetaData contains all meta data concerning the Pickrouter contract.
var PickrouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ClaimProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"adminWithdrawHT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"apy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_m1\",\"type\":\"uint256\"}],\"name\":\"decrMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCurrency\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergentWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergentWithdrawByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getDeficit\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"deficit\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getNewShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_vaultIndex\",\"type\":\"uint32\"}],\"name\":\"getTokenVaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_m1\",\"type\":\"uint256\"}],\"name\":\"incrMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_profitToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wrappedCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_referenceVaultAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"profit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"profitTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registedVaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_apy\",\"type\":\"uint256\"}],\"name\":\"setAPY\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_true\",\"type\":\"bool\"}],\"name\":\"setWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"tokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index0\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index1\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewAccumulatedProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"viewProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"p\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PickrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use PickrouterMetaData.ABI instead.
var PickrouterABI = PickrouterMetaData.ABI

// Pickrouter is an auto generated Go binding around an Ethereum contract.
type Pickrouter struct {
	PickrouterCaller     // Read-only binding to the contract
	PickrouterTransactor // Write-only binding to the contract
	PickrouterFilterer   // Log filterer for contract events
}

// PickrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PickrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PickrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PickrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PickrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PickrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PickrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PickrouterSession struct {
	Contract     *Pickrouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PickrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PickrouterCallerSession struct {
	Contract *PickrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PickrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PickrouterTransactorSession struct {
	Contract     *PickrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PickrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PickrouterRaw struct {
	Contract *Pickrouter // Generic contract binding to access the raw methods on
}

// PickrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PickrouterCallerRaw struct {
	Contract *PickrouterCaller // Generic read-only contract binding to access the raw methods on
}

// PickrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PickrouterTransactorRaw struct {
	Contract *PickrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPickrouter creates a new instance of Pickrouter, bound to a specific deployed contract.
func NewPickrouter(address common.Address, backend bind.ContractBackend) (*Pickrouter, error) {
	contract, err := bindPickrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pickrouter{PickrouterCaller: PickrouterCaller{contract: contract}, PickrouterTransactor: PickrouterTransactor{contract: contract}, PickrouterFilterer: PickrouterFilterer{contract: contract}}, nil
}

// NewPickrouterCaller creates a new read-only instance of Pickrouter, bound to a specific deployed contract.
func NewPickrouterCaller(address common.Address, caller bind.ContractCaller) (*PickrouterCaller, error) {
	contract, err := bindPickrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PickrouterCaller{contract: contract}, nil
}

// NewPickrouterTransactor creates a new write-only instance of Pickrouter, bound to a specific deployed contract.
func NewPickrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PickrouterTransactor, error) {
	contract, err := bindPickrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PickrouterTransactor{contract: contract}, nil
}

// NewPickrouterFilterer creates a new log filterer instance of Pickrouter, bound to a specific deployed contract.
func NewPickrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PickrouterFilterer, error) {
	contract, err := bindPickrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PickrouterFilterer{contract: contract}, nil
}

// bindPickrouter binds a generic wrapper to an already deployed contract.
func bindPickrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PickrouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pickrouter *PickrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pickrouter.Contract.PickrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pickrouter *PickrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pickrouter.Contract.PickrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pickrouter *PickrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pickrouter.Contract.PickrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pickrouter *PickrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pickrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pickrouter *PickrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pickrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pickrouter *PickrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pickrouter.Contract.contract.Transact(opts, method, params...)
}

// Apy is a free data retrieval call binding the contract method 0x2b3adc1c.
//
// Solidity: function apy(address ) view returns(uint256)
func (_Pickrouter *PickrouterCaller) Apy(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "apy", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apy is a free data retrieval call binding the contract method 0x2b3adc1c.
//
// Solidity: function apy(address ) view returns(uint256)
func (_Pickrouter *PickrouterSession) Apy(arg0 common.Address) (*big.Int, error) {
	return _Pickrouter.Contract.Apy(&_Pickrouter.CallOpts, arg0)
}

// Apy is a free data retrieval call binding the contract method 0x2b3adc1c.
//
// Solidity: function apy(address ) view returns(uint256)
func (_Pickrouter *PickrouterCallerSession) Apy(arg0 common.Address) (*big.Int, error) {
	return _Pickrouter.Contract.Apy(&_Pickrouter.CallOpts, arg0)
}

// FeeOwner is a free data retrieval call binding the contract method 0xb9818be1.
//
// Solidity: function feeOwner() view returns(address)
func (_Pickrouter *PickrouterCaller) FeeOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "feeOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeOwner is a free data retrieval call binding the contract method 0xb9818be1.
//
// Solidity: function feeOwner() view returns(address)
func (_Pickrouter *PickrouterSession) FeeOwner() (common.Address, error) {
	return _Pickrouter.Contract.FeeOwner(&_Pickrouter.CallOpts)
}

// FeeOwner is a free data retrieval call binding the contract method 0xb9818be1.
//
// Solidity: function feeOwner() view returns(address)
func (_Pickrouter *PickrouterCallerSession) FeeOwner() (common.Address, error) {
	return _Pickrouter.Contract.FeeOwner(&_Pickrouter.CallOpts)
}

// GetDeficit is a free data retrieval call binding the contract method 0xa74558e1.
//
// Solidity: function getDeficit(address _token) view returns(uint256[] deficit)
func (_Pickrouter *PickrouterCaller) GetDeficit(opts *bind.CallOpts, _token common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "getDeficit", _token)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDeficit is a free data retrieval call binding the contract method 0xa74558e1.
//
// Solidity: function getDeficit(address _token) view returns(uint256[] deficit)
func (_Pickrouter *PickrouterSession) GetDeficit(_token common.Address) ([]*big.Int, error) {
	return _Pickrouter.Contract.GetDeficit(&_Pickrouter.CallOpts, _token)
}

// GetDeficit is a free data retrieval call binding the contract method 0xa74558e1.
//
// Solidity: function getDeficit(address _token) view returns(uint256[] deficit)
func (_Pickrouter *PickrouterCallerSession) GetDeficit(_token common.Address) ([]*big.Int, error) {
	return _Pickrouter.Contract.GetDeficit(&_Pickrouter.CallOpts, _token)
}

// GetNewShares is a free data retrieval call binding the contract method 0xb4c45455.
//
// Solidity: function getNewShares(address _token, address _user) view returns(uint256 totalShare, uint256 userShare)
func (_Pickrouter *PickrouterCaller) GetNewShares(opts *bind.CallOpts, _token common.Address, _user common.Address) (struct {
	TotalShare *big.Int
	UserShare  *big.Int
}, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "getNewShares", _token, _user)

	outstruct := new(struct {
		TotalShare *big.Int
		UserShare  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UserShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNewShares is a free data retrieval call binding the contract method 0xb4c45455.
//
// Solidity: function getNewShares(address _token, address _user) view returns(uint256 totalShare, uint256 userShare)
func (_Pickrouter *PickrouterSession) GetNewShares(_token common.Address, _user common.Address) (struct {
	TotalShare *big.Int
	UserShare  *big.Int
}, error) {
	return _Pickrouter.Contract.GetNewShares(&_Pickrouter.CallOpts, _token, _user)
}

// GetNewShares is a free data retrieval call binding the contract method 0xb4c45455.
//
// Solidity: function getNewShares(address _token, address _user) view returns(uint256 totalShare, uint256 userShare)
func (_Pickrouter *PickrouterCallerSession) GetNewShares(_token common.Address, _user common.Address) (struct {
	TotalShare *big.Int
	UserShare  *big.Int
}, error) {
	return _Pickrouter.Contract.GetNewShares(&_Pickrouter.CallOpts, _token, _user)
}

// GetTokenVaults is a free data retrieval call binding the contract method 0x40766b11.
//
// Solidity: function getTokenVaults(address _token, uint32 _vaultIndex) view returns(address)
func (_Pickrouter *PickrouterCaller) GetTokenVaults(opts *bind.CallOpts, _token common.Address, _vaultIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "getTokenVaults", _token, _vaultIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenVaults is a free data retrieval call binding the contract method 0x40766b11.
//
// Solidity: function getTokenVaults(address _token, uint32 _vaultIndex) view returns(address)
func (_Pickrouter *PickrouterSession) GetTokenVaults(_token common.Address, _vaultIndex uint32) (common.Address, error) {
	return _Pickrouter.Contract.GetTokenVaults(&_Pickrouter.CallOpts, _token, _vaultIndex)
}

// GetTokenVaults is a free data retrieval call binding the contract method 0x40766b11.
//
// Solidity: function getTokenVaults(address _token, uint32 _vaultIndex) view returns(address)
func (_Pickrouter *PickrouterCallerSession) GetTokenVaults(_token common.Address, _vaultIndex uint32) (common.Address, error) {
	return _Pickrouter.Contract.GetTokenVaults(&_Pickrouter.CallOpts, _token, _vaultIndex)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pickrouter *PickrouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pickrouter *PickrouterSession) Owner() (common.Address, error) {
	return _Pickrouter.Contract.Owner(&_Pickrouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pickrouter *PickrouterCallerSession) Owner() (common.Address, error) {
	return _Pickrouter.Contract.Owner(&_Pickrouter.CallOpts)
}

// ProfitToken is a free data retrieval call binding the contract method 0x49343624.
//
// Solidity: function profitToken() view returns(address)
func (_Pickrouter *PickrouterCaller) ProfitToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "profitToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProfitToken is a free data retrieval call binding the contract method 0x49343624.
//
// Solidity: function profitToken() view returns(address)
func (_Pickrouter *PickrouterSession) ProfitToken() (common.Address, error) {
	return _Pickrouter.Contract.ProfitToken(&_Pickrouter.CallOpts)
}

// ProfitToken is a free data retrieval call binding the contract method 0x49343624.
//
// Solidity: function profitToken() view returns(address)
func (_Pickrouter *PickrouterCallerSession) ProfitToken() (common.Address, error) {
	return _Pickrouter.Contract.ProfitToken(&_Pickrouter.CallOpts)
}

// RegistedVaults is a free data retrieval call binding the contract method 0x58ccacfd.
//
// Solidity: function registedVaults(uint256 ) view returns(address)
func (_Pickrouter *PickrouterCaller) RegistedVaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "registedVaults", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistedVaults is a free data retrieval call binding the contract method 0x58ccacfd.
//
// Solidity: function registedVaults(uint256 ) view returns(address)
func (_Pickrouter *PickrouterSession) RegistedVaults(arg0 *big.Int) (common.Address, error) {
	return _Pickrouter.Contract.RegistedVaults(&_Pickrouter.CallOpts, arg0)
}

// RegistedVaults is a free data retrieval call binding the contract method 0x58ccacfd.
//
// Solidity: function registedVaults(uint256 ) view returns(address)
func (_Pickrouter *PickrouterCallerSession) RegistedVaults(arg0 *big.Int) (common.Address, error) {
	return _Pickrouter.Contract.RegistedVaults(&_Pickrouter.CallOpts, arg0)
}

// TokenState is a free data retrieval call binding the contract method 0x59b6f0dc.
//
// Solidity: function tokenState(address ) view returns(uint256 remain, uint256 max, uint256 shares, uint256 prevBlock, uint256 feeRate)
func (_Pickrouter *PickrouterCaller) TokenState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Remain    *big.Int
	Max       *big.Int
	Shares    *big.Int
	PrevBlock *big.Int
	FeeRate   *big.Int
}, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "tokenState", arg0)

	outstruct := new(struct {
		Remain    *big.Int
		Max       *big.Int
		Shares    *big.Int
		PrevBlock *big.Int
		FeeRate   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Remain = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Max = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrevBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FeeRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenState is a free data retrieval call binding the contract method 0x59b6f0dc.
//
// Solidity: function tokenState(address ) view returns(uint256 remain, uint256 max, uint256 shares, uint256 prevBlock, uint256 feeRate)
func (_Pickrouter *PickrouterSession) TokenState(arg0 common.Address) (struct {
	Remain    *big.Int
	Max       *big.Int
	Shares    *big.Int
	PrevBlock *big.Int
	FeeRate   *big.Int
}, error) {
	return _Pickrouter.Contract.TokenState(&_Pickrouter.CallOpts, arg0)
}

// TokenState is a free data retrieval call binding the contract method 0x59b6f0dc.
//
// Solidity: function tokenState(address ) view returns(uint256 remain, uint256 max, uint256 shares, uint256 prevBlock, uint256 feeRate)
func (_Pickrouter *PickrouterCallerSession) TokenState(arg0 common.Address) (struct {
	Remain    *big.Int
	Max       *big.Int
	Shares    *big.Int
	PrevBlock *big.Int
	FeeRate   *big.Int
}, error) {
	return _Pickrouter.Contract.TokenState(&_Pickrouter.CallOpts, arg0)
}

// UserState is a free data retrieval call binding the contract method 0x7281fb33.
//
// Solidity: function userState(address , address ) view returns(uint256 balance, uint256 prevBlock)
func (_Pickrouter *PickrouterCaller) UserState(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Balance   *big.Int
	PrevBlock *big.Int
}, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "userState", arg0, arg1)

	outstruct := new(struct {
		Balance   *big.Int
		PrevBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrevBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserState is a free data retrieval call binding the contract method 0x7281fb33.
//
// Solidity: function userState(address , address ) view returns(uint256 balance, uint256 prevBlock)
func (_Pickrouter *PickrouterSession) UserState(arg0 common.Address, arg1 common.Address) (struct {
	Balance   *big.Int
	PrevBlock *big.Int
}, error) {
	return _Pickrouter.Contract.UserState(&_Pickrouter.CallOpts, arg0, arg1)
}

// UserState is a free data retrieval call binding the contract method 0x7281fb33.
//
// Solidity: function userState(address , address ) view returns(uint256 balance, uint256 prevBlock)
func (_Pickrouter *PickrouterCallerSession) UserState(arg0 common.Address, arg1 common.Address) (struct {
	Balance   *big.Int
	PrevBlock *big.Int
}, error) {
	return _Pickrouter.Contract.UserState(&_Pickrouter.CallOpts, arg0, arg1)
}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address token0, address token1, uint32 index0, uint32 index1)
func (_Pickrouter *PickrouterCaller) VaultInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
	Index0 uint32
	Index1 uint32
}, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "vaultInfo", arg0)

	outstruct := new(struct {
		Token0 common.Address
		Token1 common.Address
		Index0 uint32
		Index1 uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Index0 = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Index1 = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address token0, address token1, uint32 index0, uint32 index1)
func (_Pickrouter *PickrouterSession) VaultInfo(arg0 common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
	Index0 uint32
	Index1 uint32
}, error) {
	return _Pickrouter.Contract.VaultInfo(&_Pickrouter.CallOpts, arg0)
}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address token0, address token1, uint32 index0, uint32 index1)
func (_Pickrouter *PickrouterCallerSession) VaultInfo(arg0 common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
	Index0 uint32
	Index1 uint32
}, error) {
	return _Pickrouter.Contract.VaultInfo(&_Pickrouter.CallOpts, arg0)
}

// ViewAccumulatedProfit is a free data retrieval call binding the contract method 0x3d391ec6.
//
// Solidity: function viewAccumulatedProfit() view returns(uint256)
func (_Pickrouter *PickrouterCaller) ViewAccumulatedProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "viewAccumulatedProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewAccumulatedProfit is a free data retrieval call binding the contract method 0x3d391ec6.
//
// Solidity: function viewAccumulatedProfit() view returns(uint256)
func (_Pickrouter *PickrouterSession) ViewAccumulatedProfit() (*big.Int, error) {
	return _Pickrouter.Contract.ViewAccumulatedProfit(&_Pickrouter.CallOpts)
}

// ViewAccumulatedProfit is a free data retrieval call binding the contract method 0x3d391ec6.
//
// Solidity: function viewAccumulatedProfit() view returns(uint256)
func (_Pickrouter *PickrouterCallerSession) ViewAccumulatedProfit() (*big.Int, error) {
	return _Pickrouter.Contract.ViewAccumulatedProfit(&_Pickrouter.CallOpts)
}

// ViewProfit is a free data retrieval call binding the contract method 0x404de794.
//
// Solidity: function viewProfit(address _user, address _token) view returns(uint256 p)
func (_Pickrouter *PickrouterCaller) ViewProfit(opts *bind.CallOpts, _user common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "viewProfit", _user, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewProfit is a free data retrieval call binding the contract method 0x404de794.
//
// Solidity: function viewProfit(address _user, address _token) view returns(uint256 p)
func (_Pickrouter *PickrouterSession) ViewProfit(_user common.Address, _token common.Address) (*big.Int, error) {
	return _Pickrouter.Contract.ViewProfit(&_Pickrouter.CallOpts, _user, _token)
}

// ViewProfit is a free data retrieval call binding the contract method 0x404de794.
//
// Solidity: function viewProfit(address _user, address _token) view returns(uint256 p)
func (_Pickrouter *PickrouterCallerSession) ViewProfit(_user common.Address, _token common.Address) (*big.Int, error) {
	return _Pickrouter.Contract.ViewProfit(&_Pickrouter.CallOpts, _user, _token)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Pickrouter *PickrouterCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Pickrouter *PickrouterSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Pickrouter.Contract.WhiteList(&_Pickrouter.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Pickrouter *PickrouterCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Pickrouter.Contract.WhiteList(&_Pickrouter.CallOpts, arg0)
}

// WrappedCurrency is a free data retrieval call binding the contract method 0xb0e1268e.
//
// Solidity: function wrappedCurrency() view returns(address)
func (_Pickrouter *PickrouterCaller) WrappedCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pickrouter.contract.Call(opts, &out, "wrappedCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedCurrency is a free data retrieval call binding the contract method 0xb0e1268e.
//
// Solidity: function wrappedCurrency() view returns(address)
func (_Pickrouter *PickrouterSession) WrappedCurrency() (common.Address, error) {
	return _Pickrouter.Contract.WrappedCurrency(&_Pickrouter.CallOpts)
}

// WrappedCurrency is a free data retrieval call binding the contract method 0xb0e1268e.
//
// Solidity: function wrappedCurrency() view returns(address)
func (_Pickrouter *PickrouterCallerSession) WrappedCurrency() (common.Address, error) {
	return _Pickrouter.Contract.WrappedCurrency(&_Pickrouter.CallOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xa2783299.
//
// Solidity: function adminWithdraw(address _token, address _to) returns()
func (_Pickrouter *PickrouterTransactor) AdminWithdraw(opts *bind.TransactOpts, _token common.Address, _to common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "adminWithdraw", _token, _to)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xa2783299.
//
// Solidity: function adminWithdraw(address _token, address _to) returns()
func (_Pickrouter *PickrouterSession) AdminWithdraw(_token common.Address, _to common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.AdminWithdraw(&_Pickrouter.TransactOpts, _token, _to)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xa2783299.
//
// Solidity: function adminWithdraw(address _token, address _to) returns()
func (_Pickrouter *PickrouterTransactorSession) AdminWithdraw(_token common.Address, _to common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.AdminWithdraw(&_Pickrouter.TransactOpts, _token, _to)
}

// AdminWithdrawHT is a paid mutator transaction binding the contract method 0x22b7bc02.
//
// Solidity: function adminWithdrawHT(address _to) returns()
func (_Pickrouter *PickrouterTransactor) AdminWithdrawHT(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "adminWithdrawHT", _to)
}

// AdminWithdrawHT is a paid mutator transaction binding the contract method 0x22b7bc02.
//
// Solidity: function adminWithdrawHT(address _to) returns()
func (_Pickrouter *PickrouterSession) AdminWithdrawHT(_to common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.AdminWithdrawHT(&_Pickrouter.TransactOpts, _to)
}

// AdminWithdrawHT is a paid mutator transaction binding the contract method 0x22b7bc02.
//
// Solidity: function adminWithdrawHT(address _to) returns()
func (_Pickrouter *PickrouterTransactorSession) AdminWithdrawHT(_to common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.AdminWithdrawHT(&_Pickrouter.TransactOpts, _to)
}

// DecrMax is a paid mutator transaction binding the contract method 0xf45848e1.
//
// Solidity: function decrMax(address _vault, uint256 _m0, uint256 _m1) returns()
func (_Pickrouter *PickrouterTransactor) DecrMax(opts *bind.TransactOpts, _vault common.Address, _m0 *big.Int, _m1 *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "decrMax", _vault, _m0, _m1)
}

// DecrMax is a paid mutator transaction binding the contract method 0xf45848e1.
//
// Solidity: function decrMax(address _vault, uint256 _m0, uint256 _m1) returns()
func (_Pickrouter *PickrouterSession) DecrMax(_vault common.Address, _m0 *big.Int, _m1 *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.DecrMax(&_Pickrouter.TransactOpts, _vault, _m0, _m1)
}

// DecrMax is a paid mutator transaction binding the contract method 0xf45848e1.
//
// Solidity: function decrMax(address _vault, uint256 _m0, uint256 _m1) returns()
func (_Pickrouter *PickrouterTransactorSession) DecrMax(_vault common.Address, _m0 *big.Int, _m1 *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.DecrMax(&_Pickrouter.TransactOpts, _vault, _m0, _m1)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.Deposit(&_Pickrouter.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.Deposit(&_Pickrouter.TransactOpts, _token, _amount)
}

// DepositCurrency is a paid mutator transaction binding the contract method 0x2d2d78f3.
//
// Solidity: function depositCurrency() payable returns()
func (_Pickrouter *PickrouterTransactor) DepositCurrency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "depositCurrency")
}

// DepositCurrency is a paid mutator transaction binding the contract method 0x2d2d78f3.
//
// Solidity: function depositCurrency() payable returns()
func (_Pickrouter *PickrouterSession) DepositCurrency() (*types.Transaction, error) {
	return _Pickrouter.Contract.DepositCurrency(&_Pickrouter.TransactOpts)
}

// DepositCurrency is a paid mutator transaction binding the contract method 0x2d2d78f3.
//
// Solidity: function depositCurrency() payable returns()
func (_Pickrouter *PickrouterTransactorSession) DepositCurrency() (*types.Transaction, error) {
	return _Pickrouter.Contract.DepositCurrency(&_Pickrouter.TransactOpts)
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x00202e1b.
//
// Solidity: function emergentWithdraw(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactor) EmergentWithdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "emergentWithdraw", _token, _amount)
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x00202e1b.
//
// Solidity: function emergentWithdraw(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterSession) EmergentWithdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.EmergentWithdraw(&_Pickrouter.TransactOpts, _token, _amount)
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x00202e1b.
//
// Solidity: function emergentWithdraw(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactorSession) EmergentWithdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.EmergentWithdraw(&_Pickrouter.TransactOpts, _token, _amount)
}

// EmergentWithdrawByAdmin is a paid mutator transaction binding the contract method 0x8cd85a15.
//
// Solidity: function emergentWithdrawByAdmin(address _user, address _vault, address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactor) EmergentWithdrawByAdmin(opts *bind.TransactOpts, _user common.Address, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "emergentWithdrawByAdmin", _user, _vault, _token, _amount)
}

// EmergentWithdrawByAdmin is a paid mutator transaction binding the contract method 0x8cd85a15.
//
// Solidity: function emergentWithdrawByAdmin(address _user, address _vault, address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterSession) EmergentWithdrawByAdmin(_user common.Address, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.EmergentWithdrawByAdmin(&_Pickrouter.TransactOpts, _user, _vault, _token, _amount)
}

// EmergentWithdrawByAdmin is a paid mutator transaction binding the contract method 0x8cd85a15.
//
// Solidity: function emergentWithdrawByAdmin(address _user, address _vault, address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactorSession) EmergentWithdrawByAdmin(_user common.Address, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.EmergentWithdrawByAdmin(&_Pickrouter.TransactOpts, _user, _vault, _token, _amount)
}

// IncrMax is a paid mutator transaction binding the contract method 0x616e7153.
//
// Solidity: function incrMax(address _vault, uint256 _m0, uint256 _m1) returns()
func (_Pickrouter *PickrouterTransactor) IncrMax(opts *bind.TransactOpts, _vault common.Address, _m0 *big.Int, _m1 *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "incrMax", _vault, _m0, _m1)
}

// IncrMax is a paid mutator transaction binding the contract method 0x616e7153.
//
// Solidity: function incrMax(address _vault, uint256 _m0, uint256 _m1) returns()
func (_Pickrouter *PickrouterSession) IncrMax(_vault common.Address, _m0 *big.Int, _m1 *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.IncrMax(&_Pickrouter.TransactOpts, _vault, _m0, _m1)
}

// IncrMax is a paid mutator transaction binding the contract method 0x616e7153.
//
// Solidity: function incrMax(address _vault, uint256 _m0, uint256 _m1) returns()
func (_Pickrouter *PickrouterTransactorSession) IncrMax(_vault common.Address, _m0 *big.Int, _m1 *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.IncrMax(&_Pickrouter.TransactOpts, _vault, _m0, _m1)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _feeOwner, address _profitToken, address _wrappedCurrency, address _referenceVaultAddr) returns()
func (_Pickrouter *PickrouterTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _feeOwner common.Address, _profitToken common.Address, _wrappedCurrency common.Address, _referenceVaultAddr common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "initialize", _owner, _feeOwner, _profitToken, _wrappedCurrency, _referenceVaultAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _feeOwner, address _profitToken, address _wrappedCurrency, address _referenceVaultAddr) returns()
func (_Pickrouter *PickrouterSession) Initialize(_owner common.Address, _feeOwner common.Address, _profitToken common.Address, _wrappedCurrency common.Address, _referenceVaultAddr common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Initialize(&_Pickrouter.TransactOpts, _owner, _feeOwner, _profitToken, _wrappedCurrency, _referenceVaultAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _feeOwner, address _profitToken, address _wrappedCurrency, address _referenceVaultAddr) returns()
func (_Pickrouter *PickrouterTransactorSession) Initialize(_owner common.Address, _feeOwner common.Address, _profitToken common.Address, _wrappedCurrency common.Address, _referenceVaultAddr common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Initialize(&_Pickrouter.TransactOpts, _owner, _feeOwner, _profitToken, _wrappedCurrency, _referenceVaultAddr)
}

// Profit is a paid mutator transaction binding the contract method 0x02550e4d.
//
// Solidity: function profit(address _token) returns()
func (_Pickrouter *PickrouterTransactor) Profit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "profit", _token)
}

// Profit is a paid mutator transaction binding the contract method 0x02550e4d.
//
// Solidity: function profit(address _token) returns()
func (_Pickrouter *PickrouterSession) Profit(_token common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Profit(&_Pickrouter.TransactOpts, _token)
}

// Profit is a paid mutator transaction binding the contract method 0x02550e4d.
//
// Solidity: function profit(address _token) returns()
func (_Pickrouter *PickrouterTransactorSession) Profit(_token common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Profit(&_Pickrouter.TransactOpts, _token)
}

// ProfitTokens is a paid mutator transaction binding the contract method 0x8e92042e.
//
// Solidity: function profitTokens(address[] _tokens) returns()
func (_Pickrouter *PickrouterTransactor) ProfitTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "profitTokens", _tokens)
}

// ProfitTokens is a paid mutator transaction binding the contract method 0x8e92042e.
//
// Solidity: function profitTokens(address[] _tokens) returns()
func (_Pickrouter *PickrouterSession) ProfitTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.ProfitTokens(&_Pickrouter.TransactOpts, _tokens)
}

// ProfitTokens is a paid mutator transaction binding the contract method 0x8e92042e.
//
// Solidity: function profitTokens(address[] _tokens) returns()
func (_Pickrouter *PickrouterTransactorSession) ProfitTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.ProfitTokens(&_Pickrouter.TransactOpts, _tokens)
}

// Rebalance is a paid mutator transaction binding the contract method 0x21c28191.
//
// Solidity: function rebalance(address _token) returns()
func (_Pickrouter *PickrouterTransactor) Rebalance(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "rebalance", _token)
}

// Rebalance is a paid mutator transaction binding the contract method 0x21c28191.
//
// Solidity: function rebalance(address _token) returns()
func (_Pickrouter *PickrouterSession) Rebalance(_token common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Rebalance(&_Pickrouter.TransactOpts, _token)
}

// Rebalance is a paid mutator transaction binding the contract method 0x21c28191.
//
// Solidity: function rebalance(address _token) returns()
func (_Pickrouter *PickrouterTransactorSession) Rebalance(_token common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Rebalance(&_Pickrouter.TransactOpts, _token)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _vault) returns()
func (_Pickrouter *PickrouterTransactor) Register(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "register", _vault)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _vault) returns()
func (_Pickrouter *PickrouterSession) Register(_vault common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Register(&_Pickrouter.TransactOpts, _vault)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _vault) returns()
func (_Pickrouter *PickrouterTransactorSession) Register(_vault common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Register(&_Pickrouter.TransactOpts, _vault)
}

// SetAPY is a paid mutator transaction binding the contract method 0xc4160fd3.
//
// Solidity: function setAPY(address _token, uint256 _apy) returns()
func (_Pickrouter *PickrouterTransactor) SetAPY(opts *bind.TransactOpts, _token common.Address, _apy *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "setAPY", _token, _apy)
}

// SetAPY is a paid mutator transaction binding the contract method 0xc4160fd3.
//
// Solidity: function setAPY(address _token, uint256 _apy) returns()
func (_Pickrouter *PickrouterSession) SetAPY(_token common.Address, _apy *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetAPY(&_Pickrouter.TransactOpts, _token, _apy)
}

// SetAPY is a paid mutator transaction binding the contract method 0xc4160fd3.
//
// Solidity: function setAPY(address _token, uint256 _apy) returns()
func (_Pickrouter *PickrouterTransactorSession) SetAPY(_token common.Address, _apy *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetAPY(&_Pickrouter.TransactOpts, _token, _apy)
}

// SetFeeOwner is a paid mutator transaction binding the contract method 0x4b104eff.
//
// Solidity: function setFeeOwner(address _addr) returns()
func (_Pickrouter *PickrouterTransactor) SetFeeOwner(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "setFeeOwner", _addr)
}

// SetFeeOwner is a paid mutator transaction binding the contract method 0x4b104eff.
//
// Solidity: function setFeeOwner(address _addr) returns()
func (_Pickrouter *PickrouterSession) SetFeeOwner(_addr common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetFeeOwner(&_Pickrouter.TransactOpts, _addr)
}

// SetFeeOwner is a paid mutator transaction binding the contract method 0x4b104eff.
//
// Solidity: function setFeeOwner(address _addr) returns()
func (_Pickrouter *PickrouterTransactorSession) SetFeeOwner(_addr common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetFeeOwner(&_Pickrouter.TransactOpts, _addr)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x942dc573.
//
// Solidity: function setFeeRate(address _token, uint256 _rate) returns()
func (_Pickrouter *PickrouterTransactor) SetFeeRate(opts *bind.TransactOpts, _token common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "setFeeRate", _token, _rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x942dc573.
//
// Solidity: function setFeeRate(address _token, uint256 _rate) returns()
func (_Pickrouter *PickrouterSession) SetFeeRate(_token common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetFeeRate(&_Pickrouter.TransactOpts, _token, _rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x942dc573.
//
// Solidity: function setFeeRate(address _token, uint256 _rate) returns()
func (_Pickrouter *PickrouterTransactorSession) SetFeeRate(_token common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetFeeRate(&_Pickrouter.TransactOpts, _token, _rate)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_Pickrouter *PickrouterTransactor) SetOwner(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "setOwner", _addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_Pickrouter *PickrouterSession) SetOwner(_addr common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetOwner(&_Pickrouter.TransactOpts, _addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_Pickrouter *PickrouterTransactorSession) SetOwner(_addr common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetOwner(&_Pickrouter.TransactOpts, _addr)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _addr, bool _true) returns()
func (_Pickrouter *PickrouterTransactor) SetWhiteList(opts *bind.TransactOpts, _addr common.Address, _true bool) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "setWhiteList", _addr, _true)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _addr, bool _true) returns()
func (_Pickrouter *PickrouterSession) SetWhiteList(_addr common.Address, _true bool) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetWhiteList(&_Pickrouter.TransactOpts, _addr, _true)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _addr, bool _true) returns()
func (_Pickrouter *PickrouterTransactorSession) SetWhiteList(_addr common.Address, _true bool) (*types.Transaction, error) {
	return _Pickrouter.Contract.SetWhiteList(&_Pickrouter.TransactOpts, _addr, _true)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0xfb08937c.
//
// Solidity: function tokenTransfer(address _token, address _vaultFrom, address _vaultTo, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactor) TokenTransfer(opts *bind.TransactOpts, _token common.Address, _vaultFrom common.Address, _vaultTo common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "tokenTransfer", _token, _vaultFrom, _vaultTo, _amount)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0xfb08937c.
//
// Solidity: function tokenTransfer(address _token, address _vaultFrom, address _vaultTo, uint256 _amount) returns()
func (_Pickrouter *PickrouterSession) TokenTransfer(_token common.Address, _vaultFrom common.Address, _vaultTo common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.TokenTransfer(&_Pickrouter.TransactOpts, _token, _vaultFrom, _vaultTo, _amount)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0xfb08937c.
//
// Solidity: function tokenTransfer(address _token, address _vaultFrom, address _vaultTo, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactorSession) TokenTransfer(_token common.Address, _vaultFrom common.Address, _vaultTo common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.TokenTransfer(&_Pickrouter.TransactOpts, _token, _vaultFrom, _vaultTo, _amount)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address _vault) returns()
func (_Pickrouter *PickrouterTransactor) Unregister(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "unregister", _vault)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address _vault) returns()
func (_Pickrouter *PickrouterSession) Unregister(_vault common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Unregister(&_Pickrouter.TransactOpts, _vault)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address _vault) returns()
func (_Pickrouter *PickrouterTransactorSession) Unregister(_vault common.Address) (*types.Transaction, error) {
	return _Pickrouter.Contract.Unregister(&_Pickrouter.TransactOpts, _vault)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "withdraw", _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.Withdraw(&_Pickrouter.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactorSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.Withdraw(&_Pickrouter.TransactOpts, _token, _amount)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x3c00a36c.
//
// Solidity: function withdrawCurrency(uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactor) WithdrawCurrency(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.contract.Transact(opts, "withdrawCurrency", _amount)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x3c00a36c.
//
// Solidity: function withdrawCurrency(uint256 _amount) returns()
func (_Pickrouter *PickrouterSession) WithdrawCurrency(_amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.WithdrawCurrency(&_Pickrouter.TransactOpts, _amount)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x3c00a36c.
//
// Solidity: function withdrawCurrency(uint256 _amount) returns()
func (_Pickrouter *PickrouterTransactorSession) WithdrawCurrency(_amount *big.Int) (*types.Transaction, error) {
	return _Pickrouter.Contract.WithdrawCurrency(&_Pickrouter.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pickrouter *PickrouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pickrouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pickrouter *PickrouterSession) Receive() (*types.Transaction, error) {
	return _Pickrouter.Contract.Receive(&_Pickrouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pickrouter *PickrouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Pickrouter.Contract.Receive(&_Pickrouter.TransactOpts)
}

// PickrouterClaimProfitIterator is returned from FilterClaimProfit and is used to iterate over the raw logs and unpacked data for ClaimProfit events raised by the Pickrouter contract.
type PickrouterClaimProfitIterator struct {
	Event *PickrouterClaimProfit // Event containing the contract specifics and raw log

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
func (it *PickrouterClaimProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PickrouterClaimProfit)
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
		it.Event = new(PickrouterClaimProfit)
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
func (it *PickrouterClaimProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PickrouterClaimProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PickrouterClaimProfit represents a ClaimProfit event raised by the Pickrouter contract.
type PickrouterClaimProfit struct {
	User  common.Address
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterClaimProfit is a free log retrieval operation binding the contract event 0x4f7cfd131b52299d8b144ecbde1508e9159e764890acce884f9c8267ab53ed67.
//
// Solidity: event ClaimProfit(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) FilterClaimProfit(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*PickrouterClaimProfitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pickrouter.contract.FilterLogs(opts, "ClaimProfit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PickrouterClaimProfitIterator{contract: _Pickrouter.contract, event: "ClaimProfit", logs: logs, sub: sub}, nil
}

// WatchClaimProfit is a free log subscription operation binding the contract event 0x4f7cfd131b52299d8b144ecbde1508e9159e764890acce884f9c8267ab53ed67.
//
// Solidity: event ClaimProfit(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) WatchClaimProfit(opts *bind.WatchOpts, sink chan<- *PickrouterClaimProfit, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pickrouter.contract.WatchLogs(opts, "ClaimProfit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PickrouterClaimProfit)
				if err := _Pickrouter.contract.UnpackLog(event, "ClaimProfit", log); err != nil {
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

// ParseClaimProfit is a log parse operation binding the contract event 0x4f7cfd131b52299d8b144ecbde1508e9159e764890acce884f9c8267ab53ed67.
//
// Solidity: event ClaimProfit(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) ParseClaimProfit(log types.Log) (*PickrouterClaimProfit, error) {
	event := new(PickrouterClaimProfit)
	if err := _Pickrouter.contract.UnpackLog(event, "ClaimProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PickrouterDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Pickrouter contract.
type PickrouterDepositIterator struct {
	Event *PickrouterDeposit // Event containing the contract specifics and raw log

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
func (it *PickrouterDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PickrouterDeposit)
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
		it.Event = new(PickrouterDeposit)
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
func (it *PickrouterDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PickrouterDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PickrouterDeposit represents a Deposit event raised by the Pickrouter contract.
type PickrouterDeposit struct {
	User  common.Address
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*PickrouterDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pickrouter.contract.FilterLogs(opts, "Deposit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PickrouterDepositIterator{contract: _Pickrouter.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PickrouterDeposit, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pickrouter.contract.WatchLogs(opts, "Deposit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PickrouterDeposit)
				if err := _Pickrouter.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) ParseDeposit(log types.Log) (*PickrouterDeposit, error) {
	event := new(PickrouterDeposit)
	if err := _Pickrouter.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PickrouterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Pickrouter contract.
type PickrouterWithdrawIterator struct {
	Event *PickrouterWithdraw // Event containing the contract specifics and raw log

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
func (it *PickrouterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PickrouterWithdraw)
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
		it.Event = new(PickrouterWithdraw)
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
func (it *PickrouterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PickrouterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PickrouterWithdraw represents a Withdraw event raised by the Pickrouter contract.
type PickrouterWithdraw struct {
	User  common.Address
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*PickrouterWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pickrouter.contract.FilterLogs(opts, "Withdraw", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PickrouterWithdrawIterator{contract: _Pickrouter.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PickrouterWithdraw, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pickrouter.contract.WatchLogs(opts, "Withdraw", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PickrouterWithdraw)
				if err := _Pickrouter.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 value)
func (_Pickrouter *PickrouterFilterer) ParseWithdraw(log types.Log) (*PickrouterWithdraw, error) {
	event := new(PickrouterWithdraw)
	if err := _Pickrouter.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
