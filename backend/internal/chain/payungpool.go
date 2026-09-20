// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// PayungPoolPolicy is an auto generated low-level Go binding around an user-defined struct.
type PayungPoolPolicy struct {
	Holder   common.Address
	ZoneId   uint16
	StartDay uint32
	EndDay   uint32
	Closed   bool
}

// PayungPoolMetaData contains all meta data concerning the PayungPool contract.
var PayungPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"idrpToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_ACTIVE_PER_ZONE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ORACLE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activePolicyOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buyPolicy\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"weeksCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"policyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createZone\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"premiumPerWeek\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"thresholdMm\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"payoutPerDay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDaysPerWeek\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fundPool\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActivePolicies\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicy\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPayungPool.Policy\",\"components\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"startDay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"endDay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"closed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"idrp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextPolicyId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paidForDay\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payoutsInWeek\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"policies\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"startDay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"endDay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"closed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pruneExpired\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rainReported\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rainfallMm\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPremium\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"premiumPerWeek\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setZoneActive\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"dayIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRainfall\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"dayIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mm\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"today\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zones\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"premiumPerWeek\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"thresholdMm\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"payoutPerDay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDaysPerWeek\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PayoutSent\",\"inputs\":[{\"name\":\"policyId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zoneId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"dayIndex\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayoutSkipped\",\"inputs\":[{\"name\":\"policyId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"dayIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyBought\",\"inputs\":[{\"name\":\"policyId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zoneId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"startDay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"endDay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"premium\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PolicyExpired\",\"inputs\":[{\"name\":\"policyId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFunded\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PremiumUpdated\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"premiumPerWeek\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RainfallReported\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"dayIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"mm\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"isRainDay\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZoneCreated\",\"inputs\":[{\"name\":\"zoneId\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"DayNotFinished\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWeeks\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRainDay\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PolicyAlreadyActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RainAlreadyReported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RainNotReported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZoneFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZoneInactive\",\"inputs\":[]}]",
}

// PayungPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PayungPoolMetaData.ABI instead.
var PayungPoolABI = PayungPoolMetaData.ABI

// PayungPool is an auto generated Go binding around an Ethereum contract.
type PayungPool struct {
	PayungPoolCaller     // Read-only binding to the contract
	PayungPoolTransactor // Write-only binding to the contract
	PayungPoolFilterer   // Log filterer for contract events
}

// PayungPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayungPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayungPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayungPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayungPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayungPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayungPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PayungPoolSession struct {
	Contract     *PayungPool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayungPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayungPoolCallerSession struct {
	Contract *PayungPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PayungPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayungPoolTransactorSession struct {
	Contract     *PayungPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PayungPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayungPoolRaw struct {
	Contract *PayungPool // Generic contract binding to access the raw methods on
}

// PayungPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayungPoolCallerRaw struct {
	Contract *PayungPoolCaller // Generic read-only contract binding to access the raw methods on
}

// PayungPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayungPoolTransactorRaw struct {
	Contract *PayungPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayungPool creates a new instance of PayungPool, bound to a specific deployed contract.
func NewPayungPool(address common.Address, backend bind.ContractBackend) (*PayungPool, error) {
	contract, err := bindPayungPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PayungPool{PayungPoolCaller: PayungPoolCaller{contract: contract}, PayungPoolTransactor: PayungPoolTransactor{contract: contract}, PayungPoolFilterer: PayungPoolFilterer{contract: contract}}, nil
}

// NewPayungPoolCaller creates a new read-only instance of PayungPool, bound to a specific deployed contract.
func NewPayungPoolCaller(address common.Address, caller bind.ContractCaller) (*PayungPoolCaller, error) {
	contract, err := bindPayungPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayungPoolCaller{contract: contract}, nil
}

// NewPayungPoolTransactor creates a new write-only instance of PayungPool, bound to a specific deployed contract.
func NewPayungPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PayungPoolTransactor, error) {
	contract, err := bindPayungPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayungPoolTransactor{contract: contract}, nil
}

// NewPayungPoolFilterer creates a new log filterer instance of PayungPool, bound to a specific deployed contract.
func NewPayungPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PayungPoolFilterer, error) {
	contract, err := bindPayungPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayungPoolFilterer{contract: contract}, nil
}

// bindPayungPool binds a generic wrapper to an already deployed contract.
func bindPayungPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PayungPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayungPool *PayungPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayungPool.Contract.PayungPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayungPool *PayungPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayungPool.Contract.PayungPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayungPool *PayungPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayungPool.Contract.PayungPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayungPool *PayungPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayungPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayungPool *PayungPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayungPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayungPool *PayungPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayungPool.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PayungPool *PayungPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PayungPool *PayungPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PayungPool.Contract.DEFAULTADMINROLE(&_PayungPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PayungPool *PayungPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PayungPool.Contract.DEFAULTADMINROLE(&_PayungPool.CallOpts)
}

// MAXACTIVEPERZONE is a free data retrieval call binding the contract method 0x9bedff8f.
//
// Solidity: function MAX_ACTIVE_PER_ZONE() view returns(uint256)
func (_PayungPool *PayungPoolCaller) MAXACTIVEPERZONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "MAX_ACTIVE_PER_ZONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXACTIVEPERZONE is a free data retrieval call binding the contract method 0x9bedff8f.
//
// Solidity: function MAX_ACTIVE_PER_ZONE() view returns(uint256)
func (_PayungPool *PayungPoolSession) MAXACTIVEPERZONE() (*big.Int, error) {
	return _PayungPool.Contract.MAXACTIVEPERZONE(&_PayungPool.CallOpts)
}

// MAXACTIVEPERZONE is a free data retrieval call binding the contract method 0x9bedff8f.
//
// Solidity: function MAX_ACTIVE_PER_ZONE() view returns(uint256)
func (_PayungPool *PayungPoolCallerSession) MAXACTIVEPERZONE() (*big.Int, error) {
	return _PayungPool.Contract.MAXACTIVEPERZONE(&_PayungPool.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_PayungPool *PayungPoolCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_PayungPool *PayungPoolSession) ORACLEROLE() ([32]byte, error) {
	return _PayungPool.Contract.ORACLEROLE(&_PayungPool.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_PayungPool *PayungPoolCallerSession) ORACLEROLE() ([32]byte, error) {
	return _PayungPool.Contract.ORACLEROLE(&_PayungPool.CallOpts)
}

// ActivePolicyOf is a free data retrieval call binding the contract method 0xdbcb03e9.
//
// Solidity: function activePolicyOf(address , uint16 ) view returns(uint256)
func (_PayungPool *PayungPoolCaller) ActivePolicyOf(opts *bind.CallOpts, arg0 common.Address, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "activePolicyOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivePolicyOf is a free data retrieval call binding the contract method 0xdbcb03e9.
//
// Solidity: function activePolicyOf(address , uint16 ) view returns(uint256)
func (_PayungPool *PayungPoolSession) ActivePolicyOf(arg0 common.Address, arg1 uint16) (*big.Int, error) {
	return _PayungPool.Contract.ActivePolicyOf(&_PayungPool.CallOpts, arg0, arg1)
}

// ActivePolicyOf is a free data retrieval call binding the contract method 0xdbcb03e9.
//
// Solidity: function activePolicyOf(address , uint16 ) view returns(uint256)
func (_PayungPool *PayungPoolCallerSession) ActivePolicyOf(arg0 common.Address, arg1 uint16) (*big.Int, error) {
	return _PayungPool.Contract.ActivePolicyOf(&_PayungPool.CallOpts, arg0, arg1)
}

// GetActivePolicies is a free data retrieval call binding the contract method 0xa968c0a6.
//
// Solidity: function getActivePolicies(uint16 zoneId) view returns(uint256[])
func (_PayungPool *PayungPoolCaller) GetActivePolicies(opts *bind.CallOpts, zoneId uint16) ([]*big.Int, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "getActivePolicies", zoneId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActivePolicies is a free data retrieval call binding the contract method 0xa968c0a6.
//
// Solidity: function getActivePolicies(uint16 zoneId) view returns(uint256[])
func (_PayungPool *PayungPoolSession) GetActivePolicies(zoneId uint16) ([]*big.Int, error) {
	return _PayungPool.Contract.GetActivePolicies(&_PayungPool.CallOpts, zoneId)
}

// GetActivePolicies is a free data retrieval call binding the contract method 0xa968c0a6.
//
// Solidity: function getActivePolicies(uint16 zoneId) view returns(uint256[])
func (_PayungPool *PayungPoolCallerSession) GetActivePolicies(zoneId uint16) ([]*big.Int, error) {
	return _PayungPool.Contract.GetActivePolicies(&_PayungPool.CallOpts, zoneId)
}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 id) view returns((address,uint16,uint32,uint32,bool))
func (_PayungPool *PayungPoolCaller) GetPolicy(opts *bind.CallOpts, id *big.Int) (PayungPoolPolicy, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "getPolicy", id)

	if err != nil {
		return *new(PayungPoolPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(PayungPoolPolicy)).(*PayungPoolPolicy)

	return out0, err

}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 id) view returns((address,uint16,uint32,uint32,bool))
func (_PayungPool *PayungPoolSession) GetPolicy(id *big.Int) (PayungPoolPolicy, error) {
	return _PayungPool.Contract.GetPolicy(&_PayungPool.CallOpts, id)
}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 id) view returns((address,uint16,uint32,uint32,bool))
func (_PayungPool *PayungPoolCallerSession) GetPolicy(id *big.Int) (PayungPoolPolicy, error) {
	return _PayungPool.Contract.GetPolicy(&_PayungPool.CallOpts, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PayungPool *PayungPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PayungPool *PayungPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PayungPool.Contract.GetRoleAdmin(&_PayungPool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PayungPool *PayungPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PayungPool.Contract.GetRoleAdmin(&_PayungPool.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PayungPool *PayungPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PayungPool *PayungPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PayungPool.Contract.HasRole(&_PayungPool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PayungPool *PayungPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PayungPool.Contract.HasRole(&_PayungPool.CallOpts, role, account)
}

// Idrp is a free data retrieval call binding the contract method 0xce63bca0.
//
// Solidity: function idrp() view returns(address)
func (_PayungPool *PayungPoolCaller) Idrp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "idrp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Idrp is a free data retrieval call binding the contract method 0xce63bca0.
//
// Solidity: function idrp() view returns(address)
func (_PayungPool *PayungPoolSession) Idrp() (common.Address, error) {
	return _PayungPool.Contract.Idrp(&_PayungPool.CallOpts)
}

// Idrp is a free data retrieval call binding the contract method 0xce63bca0.
//
// Solidity: function idrp() view returns(address)
func (_PayungPool *PayungPoolCallerSession) Idrp() (common.Address, error) {
	return _PayungPool.Contract.Idrp(&_PayungPool.CallOpts)
}

// NextPolicyId is a free data retrieval call binding the contract method 0xcad0b8db.
//
// Solidity: function nextPolicyId() view returns(uint256)
func (_PayungPool *PayungPoolCaller) NextPolicyId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "nextPolicyId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPolicyId is a free data retrieval call binding the contract method 0xcad0b8db.
//
// Solidity: function nextPolicyId() view returns(uint256)
func (_PayungPool *PayungPoolSession) NextPolicyId() (*big.Int, error) {
	return _PayungPool.Contract.NextPolicyId(&_PayungPool.CallOpts)
}

// NextPolicyId is a free data retrieval call binding the contract method 0xcad0b8db.
//
// Solidity: function nextPolicyId() view returns(uint256)
func (_PayungPool *PayungPoolCallerSession) NextPolicyId() (*big.Int, error) {
	return _PayungPool.Contract.NextPolicyId(&_PayungPool.CallOpts)
}

// PaidForDay is a free data retrieval call binding the contract method 0xf84ca62c.
//
// Solidity: function paidForDay(uint256 , uint32 ) view returns(bool)
func (_PayungPool *PayungPoolCaller) PaidForDay(opts *bind.CallOpts, arg0 *big.Int, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "paidForDay", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PaidForDay is a free data retrieval call binding the contract method 0xf84ca62c.
//
// Solidity: function paidForDay(uint256 , uint32 ) view returns(bool)
func (_PayungPool *PayungPoolSession) PaidForDay(arg0 *big.Int, arg1 uint32) (bool, error) {
	return _PayungPool.Contract.PaidForDay(&_PayungPool.CallOpts, arg0, arg1)
}

// PaidForDay is a free data retrieval call binding the contract method 0xf84ca62c.
//
// Solidity: function paidForDay(uint256 , uint32 ) view returns(bool)
func (_PayungPool *PayungPoolCallerSession) PaidForDay(arg0 *big.Int, arg1 uint32) (bool, error) {
	return _PayungPool.Contract.PaidForDay(&_PayungPool.CallOpts, arg0, arg1)
}

// PayoutsInWeek is a free data retrieval call binding the contract method 0xa000f2f7.
//
// Solidity: function payoutsInWeek(uint256 , uint32 ) view returns(uint8)
func (_PayungPool *PayungPoolCaller) PayoutsInWeek(opts *bind.CallOpts, arg0 *big.Int, arg1 uint32) (uint8, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "payoutsInWeek", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PayoutsInWeek is a free data retrieval call binding the contract method 0xa000f2f7.
//
// Solidity: function payoutsInWeek(uint256 , uint32 ) view returns(uint8)
func (_PayungPool *PayungPoolSession) PayoutsInWeek(arg0 *big.Int, arg1 uint32) (uint8, error) {
	return _PayungPool.Contract.PayoutsInWeek(&_PayungPool.CallOpts, arg0, arg1)
}

// PayoutsInWeek is a free data retrieval call binding the contract method 0xa000f2f7.
//
// Solidity: function payoutsInWeek(uint256 , uint32 ) view returns(uint8)
func (_PayungPool *PayungPoolCallerSession) PayoutsInWeek(arg0 *big.Int, arg1 uint32) (uint8, error) {
	return _PayungPool.Contract.PayoutsInWeek(&_PayungPool.CallOpts, arg0, arg1)
}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(address holder, uint16 zoneId, uint32 startDay, uint32 endDay, bool closed)
func (_PayungPool *PayungPoolCaller) Policies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Holder   common.Address
	ZoneId   uint16
	StartDay uint32
	EndDay   uint32
	Closed   bool
}, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "policies", arg0)

	outstruct := new(struct {
		Holder   common.Address
		ZoneId   uint16
		StartDay uint32
		EndDay   uint32
		Closed   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Holder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ZoneId = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.StartDay = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.EndDay = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Closed = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(address holder, uint16 zoneId, uint32 startDay, uint32 endDay, bool closed)
func (_PayungPool *PayungPoolSession) Policies(arg0 *big.Int) (struct {
	Holder   common.Address
	ZoneId   uint16
	StartDay uint32
	EndDay   uint32
	Closed   bool
}, error) {
	return _PayungPool.Contract.Policies(&_PayungPool.CallOpts, arg0)
}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(address holder, uint16 zoneId, uint32 startDay, uint32 endDay, bool closed)
func (_PayungPool *PayungPoolCallerSession) Policies(arg0 *big.Int) (struct {
	Holder   common.Address
	ZoneId   uint16
	StartDay uint32
	EndDay   uint32
	Closed   bool
}, error) {
	return _PayungPool.Contract.Policies(&_PayungPool.CallOpts, arg0)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_PayungPool *PayungPoolCaller) PoolBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "poolBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_PayungPool *PayungPoolSession) PoolBalance() (*big.Int, error) {
	return _PayungPool.Contract.PoolBalance(&_PayungPool.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_PayungPool *PayungPoolCallerSession) PoolBalance() (*big.Int, error) {
	return _PayungPool.Contract.PoolBalance(&_PayungPool.CallOpts)
}

// RainReported is a free data retrieval call binding the contract method 0x8b1f6ab5.
//
// Solidity: function rainReported(uint16 , uint32 ) view returns(bool)
func (_PayungPool *PayungPoolCaller) RainReported(opts *bind.CallOpts, arg0 uint16, arg1 uint32) (bool, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "rainReported", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RainReported is a free data retrieval call binding the contract method 0x8b1f6ab5.
//
// Solidity: function rainReported(uint16 , uint32 ) view returns(bool)
func (_PayungPool *PayungPoolSession) RainReported(arg0 uint16, arg1 uint32) (bool, error) {
	return _PayungPool.Contract.RainReported(&_PayungPool.CallOpts, arg0, arg1)
}

// RainReported is a free data retrieval call binding the contract method 0x8b1f6ab5.
//
// Solidity: function rainReported(uint16 , uint32 ) view returns(bool)
func (_PayungPool *PayungPoolCallerSession) RainReported(arg0 uint16, arg1 uint32) (bool, error) {
	return _PayungPool.Contract.RainReported(&_PayungPool.CallOpts, arg0, arg1)
}

// RainfallMm is a free data retrieval call binding the contract method 0x3ac32b8f.
//
// Solidity: function rainfallMm(uint16 , uint32 ) view returns(uint16)
func (_PayungPool *PayungPoolCaller) RainfallMm(opts *bind.CallOpts, arg0 uint16, arg1 uint32) (uint16, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "rainfallMm", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RainfallMm is a free data retrieval call binding the contract method 0x3ac32b8f.
//
// Solidity: function rainfallMm(uint16 , uint32 ) view returns(uint16)
func (_PayungPool *PayungPoolSession) RainfallMm(arg0 uint16, arg1 uint32) (uint16, error) {
	return _PayungPool.Contract.RainfallMm(&_PayungPool.CallOpts, arg0, arg1)
}

// RainfallMm is a free data retrieval call binding the contract method 0x3ac32b8f.
//
// Solidity: function rainfallMm(uint16 , uint32 ) view returns(uint16)
func (_PayungPool *PayungPoolCallerSession) RainfallMm(arg0 uint16, arg1 uint32) (uint16, error) {
	return _PayungPool.Contract.RainfallMm(&_PayungPool.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PayungPool *PayungPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PayungPool *PayungPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PayungPool.Contract.SupportsInterface(&_PayungPool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PayungPool *PayungPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PayungPool.Contract.SupportsInterface(&_PayungPool.CallOpts, interfaceId)
}

// Today is a free data retrieval call binding the contract method 0xb74e452b.
//
// Solidity: function today() view returns(uint32)
func (_PayungPool *PayungPoolCaller) Today(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "today")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Today is a free data retrieval call binding the contract method 0xb74e452b.
//
// Solidity: function today() view returns(uint32)
func (_PayungPool *PayungPoolSession) Today() (uint32, error) {
	return _PayungPool.Contract.Today(&_PayungPool.CallOpts)
}

// Today is a free data retrieval call binding the contract method 0xb74e452b.
//
// Solidity: function today() view returns(uint32)
func (_PayungPool *PayungPoolCallerSession) Today() (uint32, error) {
	return _PayungPool.Contract.Today(&_PayungPool.CallOpts)
}

// Zones is a free data retrieval call binding the contract method 0x00ccccce.
//
// Solidity: function zones(uint16 ) view returns(string name, uint256 premiumPerWeek, uint16 thresholdMm, uint256 payoutPerDay, uint8 maxDaysPerWeek, bool active)
func (_PayungPool *PayungPoolCaller) Zones(opts *bind.CallOpts, arg0 uint16) (struct {
	Name           string
	PremiumPerWeek *big.Int
	ThresholdMm    uint16
	PayoutPerDay   *big.Int
	MaxDaysPerWeek uint8
	Active         bool
}, error) {
	var out []interface{}
	err := _PayungPool.contract.Call(opts, &out, "zones", arg0)

	outstruct := new(struct {
		Name           string
		PremiumPerWeek *big.Int
		ThresholdMm    uint16
		PayoutPerDay   *big.Int
		MaxDaysPerWeek uint8
		Active         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.PremiumPerWeek = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ThresholdMm = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.PayoutPerDay = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxDaysPerWeek = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Active = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Zones is a free data retrieval call binding the contract method 0x00ccccce.
//
// Solidity: function zones(uint16 ) view returns(string name, uint256 premiumPerWeek, uint16 thresholdMm, uint256 payoutPerDay, uint8 maxDaysPerWeek, bool active)
func (_PayungPool *PayungPoolSession) Zones(arg0 uint16) (struct {
	Name           string
	PremiumPerWeek *big.Int
	ThresholdMm    uint16
	PayoutPerDay   *big.Int
	MaxDaysPerWeek uint8
	Active         bool
}, error) {
	return _PayungPool.Contract.Zones(&_PayungPool.CallOpts, arg0)
}

// Zones is a free data retrieval call binding the contract method 0x00ccccce.
//
// Solidity: function zones(uint16 ) view returns(string name, uint256 premiumPerWeek, uint16 thresholdMm, uint256 payoutPerDay, uint8 maxDaysPerWeek, bool active)
func (_PayungPool *PayungPoolCallerSession) Zones(arg0 uint16) (struct {
	Name           string
	PremiumPerWeek *big.Int
	ThresholdMm    uint16
	PayoutPerDay   *big.Int
	MaxDaysPerWeek uint8
	Active         bool
}, error) {
	return _PayungPool.Contract.Zones(&_PayungPool.CallOpts, arg0)
}

// BuyPolicy is a paid mutator transaction binding the contract method 0xc8a7e3fd.
//
// Solidity: function buyPolicy(uint16 zoneId, uint8 weeksCount) returns(uint256 policyId)
func (_PayungPool *PayungPoolTransactor) BuyPolicy(opts *bind.TransactOpts, zoneId uint16, weeksCount uint8) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "buyPolicy", zoneId, weeksCount)
}

// BuyPolicy is a paid mutator transaction binding the contract method 0xc8a7e3fd.
//
// Solidity: function buyPolicy(uint16 zoneId, uint8 weeksCount) returns(uint256 policyId)
func (_PayungPool *PayungPoolSession) BuyPolicy(zoneId uint16, weeksCount uint8) (*types.Transaction, error) {
	return _PayungPool.Contract.BuyPolicy(&_PayungPool.TransactOpts, zoneId, weeksCount)
}

// BuyPolicy is a paid mutator transaction binding the contract method 0xc8a7e3fd.
//
// Solidity: function buyPolicy(uint16 zoneId, uint8 weeksCount) returns(uint256 policyId)
func (_PayungPool *PayungPoolTransactorSession) BuyPolicy(zoneId uint16, weeksCount uint8) (*types.Transaction, error) {
	return _PayungPool.Contract.BuyPolicy(&_PayungPool.TransactOpts, zoneId, weeksCount)
}

// CreateZone is a paid mutator transaction binding the contract method 0xf5957e39.
//
// Solidity: function createZone(uint16 zoneId, string name, uint256 premiumPerWeek, uint16 thresholdMm, uint256 payoutPerDay, uint8 maxDaysPerWeek) returns()
func (_PayungPool *PayungPoolTransactor) CreateZone(opts *bind.TransactOpts, zoneId uint16, name string, premiumPerWeek *big.Int, thresholdMm uint16, payoutPerDay *big.Int, maxDaysPerWeek uint8) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "createZone", zoneId, name, premiumPerWeek, thresholdMm, payoutPerDay, maxDaysPerWeek)
}

// CreateZone is a paid mutator transaction binding the contract method 0xf5957e39.
//
// Solidity: function createZone(uint16 zoneId, string name, uint256 premiumPerWeek, uint16 thresholdMm, uint256 payoutPerDay, uint8 maxDaysPerWeek) returns()
func (_PayungPool *PayungPoolSession) CreateZone(zoneId uint16, name string, premiumPerWeek *big.Int, thresholdMm uint16, payoutPerDay *big.Int, maxDaysPerWeek uint8) (*types.Transaction, error) {
	return _PayungPool.Contract.CreateZone(&_PayungPool.TransactOpts, zoneId, name, premiumPerWeek, thresholdMm, payoutPerDay, maxDaysPerWeek)
}

// CreateZone is a paid mutator transaction binding the contract method 0xf5957e39.
//
// Solidity: function createZone(uint16 zoneId, string name, uint256 premiumPerWeek, uint16 thresholdMm, uint256 payoutPerDay, uint8 maxDaysPerWeek) returns()
func (_PayungPool *PayungPoolTransactorSession) CreateZone(zoneId uint16, name string, premiumPerWeek *big.Int, thresholdMm uint16, payoutPerDay *big.Int, maxDaysPerWeek uint8) (*types.Transaction, error) {
	return _PayungPool.Contract.CreateZone(&_PayungPool.TransactOpts, zoneId, name, premiumPerWeek, thresholdMm, payoutPerDay, maxDaysPerWeek)
}

// FundPool is a paid mutator transaction binding the contract method 0xda3a5f72.
//
// Solidity: function fundPool(uint256 amount) returns()
func (_PayungPool *PayungPoolTransactor) FundPool(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "fundPool", amount)
}

// FundPool is a paid mutator transaction binding the contract method 0xda3a5f72.
//
// Solidity: function fundPool(uint256 amount) returns()
func (_PayungPool *PayungPoolSession) FundPool(amount *big.Int) (*types.Transaction, error) {
	return _PayungPool.Contract.FundPool(&_PayungPool.TransactOpts, amount)
}

// FundPool is a paid mutator transaction binding the contract method 0xda3a5f72.
//
// Solidity: function fundPool(uint256 amount) returns()
func (_PayungPool *PayungPoolTransactorSession) FundPool(amount *big.Int) (*types.Transaction, error) {
	return _PayungPool.Contract.FundPool(&_PayungPool.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PayungPool *PayungPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PayungPool *PayungPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PayungPool.Contract.GrantRole(&_PayungPool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PayungPool *PayungPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PayungPool.Contract.GrantRole(&_PayungPool.TransactOpts, role, account)
}

// PruneExpired is a paid mutator transaction binding the contract method 0x1d78e69d.
//
// Solidity: function pruneExpired(uint16 zoneId) returns()
func (_PayungPool *PayungPoolTransactor) PruneExpired(opts *bind.TransactOpts, zoneId uint16) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "pruneExpired", zoneId)
}

// PruneExpired is a paid mutator transaction binding the contract method 0x1d78e69d.
//
// Solidity: function pruneExpired(uint16 zoneId) returns()
func (_PayungPool *PayungPoolSession) PruneExpired(zoneId uint16) (*types.Transaction, error) {
	return _PayungPool.Contract.PruneExpired(&_PayungPool.TransactOpts, zoneId)
}

// PruneExpired is a paid mutator transaction binding the contract method 0x1d78e69d.
//
// Solidity: function pruneExpired(uint16 zoneId) returns()
func (_PayungPool *PayungPoolTransactorSession) PruneExpired(zoneId uint16) (*types.Transaction, error) {
	return _PayungPool.Contract.PruneExpired(&_PayungPool.TransactOpts, zoneId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PayungPool *PayungPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PayungPool *PayungPoolSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PayungPool.Contract.RenounceRole(&_PayungPool.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PayungPool *PayungPoolTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PayungPool.Contract.RenounceRole(&_PayungPool.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PayungPool *PayungPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PayungPool *PayungPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PayungPool.Contract.RevokeRole(&_PayungPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PayungPool *PayungPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PayungPool.Contract.RevokeRole(&_PayungPool.TransactOpts, role, account)
}

// SetPremium is a paid mutator transaction binding the contract method 0xa4240bc9.
//
// Solidity: function setPremium(uint16 zoneId, uint256 premiumPerWeek) returns()
func (_PayungPool *PayungPoolTransactor) SetPremium(opts *bind.TransactOpts, zoneId uint16, premiumPerWeek *big.Int) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "setPremium", zoneId, premiumPerWeek)
}

// SetPremium is a paid mutator transaction binding the contract method 0xa4240bc9.
//
// Solidity: function setPremium(uint16 zoneId, uint256 premiumPerWeek) returns()
func (_PayungPool *PayungPoolSession) SetPremium(zoneId uint16, premiumPerWeek *big.Int) (*types.Transaction, error) {
	return _PayungPool.Contract.SetPremium(&_PayungPool.TransactOpts, zoneId, premiumPerWeek)
}

// SetPremium is a paid mutator transaction binding the contract method 0xa4240bc9.
//
// Solidity: function setPremium(uint16 zoneId, uint256 premiumPerWeek) returns()
func (_PayungPool *PayungPoolTransactorSession) SetPremium(zoneId uint16, premiumPerWeek *big.Int) (*types.Transaction, error) {
	return _PayungPool.Contract.SetPremium(&_PayungPool.TransactOpts, zoneId, premiumPerWeek)
}

// SetZoneActive is a paid mutator transaction binding the contract method 0x22215254.
//
// Solidity: function setZoneActive(uint16 zoneId, bool active) returns()
func (_PayungPool *PayungPoolTransactor) SetZoneActive(opts *bind.TransactOpts, zoneId uint16, active bool) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "setZoneActive", zoneId, active)
}

// SetZoneActive is a paid mutator transaction binding the contract method 0x22215254.
//
// Solidity: function setZoneActive(uint16 zoneId, bool active) returns()
func (_PayungPool *PayungPoolSession) SetZoneActive(zoneId uint16, active bool) (*types.Transaction, error) {
	return _PayungPool.Contract.SetZoneActive(&_PayungPool.TransactOpts, zoneId, active)
}

// SetZoneActive is a paid mutator transaction binding the contract method 0x22215254.
//
// Solidity: function setZoneActive(uint16 zoneId, bool active) returns()
func (_PayungPool *PayungPoolTransactorSession) SetZoneActive(zoneId uint16, active bool) (*types.Transaction, error) {
	return _PayungPool.Contract.SetZoneActive(&_PayungPool.TransactOpts, zoneId, active)
}

// Settle is a paid mutator transaction binding the contract method 0x3901ec3b.
//
// Solidity: function settle(uint16 zoneId, uint32 dayIndex) returns()
func (_PayungPool *PayungPoolTransactor) Settle(opts *bind.TransactOpts, zoneId uint16, dayIndex uint32) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "settle", zoneId, dayIndex)
}

// Settle is a paid mutator transaction binding the contract method 0x3901ec3b.
//
// Solidity: function settle(uint16 zoneId, uint32 dayIndex) returns()
func (_PayungPool *PayungPoolSession) Settle(zoneId uint16, dayIndex uint32) (*types.Transaction, error) {
	return _PayungPool.Contract.Settle(&_PayungPool.TransactOpts, zoneId, dayIndex)
}

// Settle is a paid mutator transaction binding the contract method 0x3901ec3b.
//
// Solidity: function settle(uint16 zoneId, uint32 dayIndex) returns()
func (_PayungPool *PayungPoolTransactorSession) Settle(zoneId uint16, dayIndex uint32) (*types.Transaction, error) {
	return _PayungPool.Contract.Settle(&_PayungPool.TransactOpts, zoneId, dayIndex)
}

// SubmitRainfall is a paid mutator transaction binding the contract method 0x62399d38.
//
// Solidity: function submitRainfall(uint16 zoneId, uint32 dayIndex, uint16 mm) returns()
func (_PayungPool *PayungPoolTransactor) SubmitRainfall(opts *bind.TransactOpts, zoneId uint16, dayIndex uint32, mm uint16) (*types.Transaction, error) {
	return _PayungPool.contract.Transact(opts, "submitRainfall", zoneId, dayIndex, mm)
}

// SubmitRainfall is a paid mutator transaction binding the contract method 0x62399d38.
//
// Solidity: function submitRainfall(uint16 zoneId, uint32 dayIndex, uint16 mm) returns()
func (_PayungPool *PayungPoolSession) SubmitRainfall(zoneId uint16, dayIndex uint32, mm uint16) (*types.Transaction, error) {
	return _PayungPool.Contract.SubmitRainfall(&_PayungPool.TransactOpts, zoneId, dayIndex, mm)
}

// SubmitRainfall is a paid mutator transaction binding the contract method 0x62399d38.
//
// Solidity: function submitRainfall(uint16 zoneId, uint32 dayIndex, uint16 mm) returns()
func (_PayungPool *PayungPoolTransactorSession) SubmitRainfall(zoneId uint16, dayIndex uint32, mm uint16) (*types.Transaction, error) {
	return _PayungPool.Contract.SubmitRainfall(&_PayungPool.TransactOpts, zoneId, dayIndex, mm)
}

// PayungPoolPayoutSentIterator is returned from FilterPayoutSent and is used to iterate over the raw logs and unpacked data for PayoutSent events raised by the PayungPool contract.
type PayungPoolPayoutSentIterator struct {
	Event *PayungPoolPayoutSent // Event containing the contract specifics and raw log

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
func (it *PayungPoolPayoutSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolPayoutSent)
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
		it.Event = new(PayungPoolPayoutSent)
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
func (it *PayungPoolPayoutSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolPayoutSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolPayoutSent represents a PayoutSent event raised by the PayungPool contract.
type PayungPoolPayoutSent struct {
	PolicyId *big.Int
	Holder   common.Address
	ZoneId   uint16
	DayIndex uint32
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayoutSent is a free log retrieval operation binding the contract event 0x1dcbf782c8e614b5dcafda7154c04585e394e769a58d20954e98c191c605dd61.
//
// Solidity: event PayoutSent(uint256 indexed policyId, address indexed holder, uint16 indexed zoneId, uint32 dayIndex, uint256 amount)
func (_PayungPool *PayungPoolFilterer) FilterPayoutSent(opts *bind.FilterOpts, policyId []*big.Int, holder []common.Address, zoneId []uint16) (*PayungPoolPayoutSentIterator, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "PayoutSent", policyIdRule, holderRule, zoneIdRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolPayoutSentIterator{contract: _PayungPool.contract, event: "PayoutSent", logs: logs, sub: sub}, nil
}

// WatchPayoutSent is a free log subscription operation binding the contract event 0x1dcbf782c8e614b5dcafda7154c04585e394e769a58d20954e98c191c605dd61.
//
// Solidity: event PayoutSent(uint256 indexed policyId, address indexed holder, uint16 indexed zoneId, uint32 dayIndex, uint256 amount)
func (_PayungPool *PayungPoolFilterer) WatchPayoutSent(opts *bind.WatchOpts, sink chan<- *PayungPoolPayoutSent, policyId []*big.Int, holder []common.Address, zoneId []uint16) (event.Subscription, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "PayoutSent", policyIdRule, holderRule, zoneIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolPayoutSent)
				if err := _PayungPool.contract.UnpackLog(event, "PayoutSent", log); err != nil {
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

// ParsePayoutSent is a log parse operation binding the contract event 0x1dcbf782c8e614b5dcafda7154c04585e394e769a58d20954e98c191c605dd61.
//
// Solidity: event PayoutSent(uint256 indexed policyId, address indexed holder, uint16 indexed zoneId, uint32 dayIndex, uint256 amount)
func (_PayungPool *PayungPoolFilterer) ParsePayoutSent(log types.Log) (*PayungPoolPayoutSent, error) {
	event := new(PayungPoolPayoutSent)
	if err := _PayungPool.contract.UnpackLog(event, "PayoutSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolPayoutSkippedIterator is returned from FilterPayoutSkipped and is used to iterate over the raw logs and unpacked data for PayoutSkipped events raised by the PayungPool contract.
type PayungPoolPayoutSkippedIterator struct {
	Event *PayungPoolPayoutSkipped // Event containing the contract specifics and raw log

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
func (it *PayungPoolPayoutSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolPayoutSkipped)
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
		it.Event = new(PayungPoolPayoutSkipped)
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
func (it *PayungPoolPayoutSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolPayoutSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolPayoutSkipped represents a PayoutSkipped event raised by the PayungPool contract.
type PayungPoolPayoutSkipped struct {
	PolicyId *big.Int
	DayIndex uint32
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPayoutSkipped is a free log retrieval operation binding the contract event 0x8b59234143e5521c1bdb41e3a12a50212d8b3745a0457d355d3c7d90f7c9a22b.
//
// Solidity: event PayoutSkipped(uint256 indexed policyId, uint32 indexed dayIndex, string reason)
func (_PayungPool *PayungPoolFilterer) FilterPayoutSkipped(opts *bind.FilterOpts, policyId []*big.Int, dayIndex []uint32) (*PayungPoolPayoutSkippedIterator, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}
	var dayIndexRule []interface{}
	for _, dayIndexItem := range dayIndex {
		dayIndexRule = append(dayIndexRule, dayIndexItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "PayoutSkipped", policyIdRule, dayIndexRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolPayoutSkippedIterator{contract: _PayungPool.contract, event: "PayoutSkipped", logs: logs, sub: sub}, nil
}

// WatchPayoutSkipped is a free log subscription operation binding the contract event 0x8b59234143e5521c1bdb41e3a12a50212d8b3745a0457d355d3c7d90f7c9a22b.
//
// Solidity: event PayoutSkipped(uint256 indexed policyId, uint32 indexed dayIndex, string reason)
func (_PayungPool *PayungPoolFilterer) WatchPayoutSkipped(opts *bind.WatchOpts, sink chan<- *PayungPoolPayoutSkipped, policyId []*big.Int, dayIndex []uint32) (event.Subscription, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}
	var dayIndexRule []interface{}
	for _, dayIndexItem := range dayIndex {
		dayIndexRule = append(dayIndexRule, dayIndexItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "PayoutSkipped", policyIdRule, dayIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolPayoutSkipped)
				if err := _PayungPool.contract.UnpackLog(event, "PayoutSkipped", log); err != nil {
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

// ParsePayoutSkipped is a log parse operation binding the contract event 0x8b59234143e5521c1bdb41e3a12a50212d8b3745a0457d355d3c7d90f7c9a22b.
//
// Solidity: event PayoutSkipped(uint256 indexed policyId, uint32 indexed dayIndex, string reason)
func (_PayungPool *PayungPoolFilterer) ParsePayoutSkipped(log types.Log) (*PayungPoolPayoutSkipped, error) {
	event := new(PayungPoolPayoutSkipped)
	if err := _PayungPool.contract.UnpackLog(event, "PayoutSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolPolicyBoughtIterator is returned from FilterPolicyBought and is used to iterate over the raw logs and unpacked data for PolicyBought events raised by the PayungPool contract.
type PayungPoolPolicyBoughtIterator struct {
	Event *PayungPoolPolicyBought // Event containing the contract specifics and raw log

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
func (it *PayungPoolPolicyBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolPolicyBought)
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
		it.Event = new(PayungPoolPolicyBought)
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
func (it *PayungPoolPolicyBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolPolicyBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolPolicyBought represents a PolicyBought event raised by the PayungPool contract.
type PayungPoolPolicyBought struct {
	PolicyId *big.Int
	Holder   common.Address
	ZoneId   uint16
	StartDay uint32
	EndDay   uint32
	Premium  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPolicyBought is a free log retrieval operation binding the contract event 0x882880044d7eeef5e3c8c196f76c25ecc0ff8aad1d8a100deb4df737ed2d2318.
//
// Solidity: event PolicyBought(uint256 indexed policyId, address indexed holder, uint16 indexed zoneId, uint32 startDay, uint32 endDay, uint256 premium)
func (_PayungPool *PayungPoolFilterer) FilterPolicyBought(opts *bind.FilterOpts, policyId []*big.Int, holder []common.Address, zoneId []uint16) (*PayungPoolPolicyBoughtIterator, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "PolicyBought", policyIdRule, holderRule, zoneIdRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolPolicyBoughtIterator{contract: _PayungPool.contract, event: "PolicyBought", logs: logs, sub: sub}, nil
}

// WatchPolicyBought is a free log subscription operation binding the contract event 0x882880044d7eeef5e3c8c196f76c25ecc0ff8aad1d8a100deb4df737ed2d2318.
//
// Solidity: event PolicyBought(uint256 indexed policyId, address indexed holder, uint16 indexed zoneId, uint32 startDay, uint32 endDay, uint256 premium)
func (_PayungPool *PayungPoolFilterer) WatchPolicyBought(opts *bind.WatchOpts, sink chan<- *PayungPoolPolicyBought, policyId []*big.Int, holder []common.Address, zoneId []uint16) (event.Subscription, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "PolicyBought", policyIdRule, holderRule, zoneIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolPolicyBought)
				if err := _PayungPool.contract.UnpackLog(event, "PolicyBought", log); err != nil {
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

// ParsePolicyBought is a log parse operation binding the contract event 0x882880044d7eeef5e3c8c196f76c25ecc0ff8aad1d8a100deb4df737ed2d2318.
//
// Solidity: event PolicyBought(uint256 indexed policyId, address indexed holder, uint16 indexed zoneId, uint32 startDay, uint32 endDay, uint256 premium)
func (_PayungPool *PayungPoolFilterer) ParsePolicyBought(log types.Log) (*PayungPoolPolicyBought, error) {
	event := new(PayungPoolPolicyBought)
	if err := _PayungPool.contract.UnpackLog(event, "PolicyBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolPolicyExpiredIterator is returned from FilterPolicyExpired and is used to iterate over the raw logs and unpacked data for PolicyExpired events raised by the PayungPool contract.
type PayungPoolPolicyExpiredIterator struct {
	Event *PayungPoolPolicyExpired // Event containing the contract specifics and raw log

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
func (it *PayungPoolPolicyExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolPolicyExpired)
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
		it.Event = new(PayungPoolPolicyExpired)
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
func (it *PayungPoolPolicyExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolPolicyExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolPolicyExpired represents a PolicyExpired event raised by the PayungPool contract.
type PayungPoolPolicyExpired struct {
	PolicyId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPolicyExpired is a free log retrieval operation binding the contract event 0xb22aa541487fc93d989bb5b17d1e3a967f649c4979cfd658c0849e070ae0dca0.
//
// Solidity: event PolicyExpired(uint256 indexed policyId)
func (_PayungPool *PayungPoolFilterer) FilterPolicyExpired(opts *bind.FilterOpts, policyId []*big.Int) (*PayungPoolPolicyExpiredIterator, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "PolicyExpired", policyIdRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolPolicyExpiredIterator{contract: _PayungPool.contract, event: "PolicyExpired", logs: logs, sub: sub}, nil
}

// WatchPolicyExpired is a free log subscription operation binding the contract event 0xb22aa541487fc93d989bb5b17d1e3a967f649c4979cfd658c0849e070ae0dca0.
//
// Solidity: event PolicyExpired(uint256 indexed policyId)
func (_PayungPool *PayungPoolFilterer) WatchPolicyExpired(opts *bind.WatchOpts, sink chan<- *PayungPoolPolicyExpired, policyId []*big.Int) (event.Subscription, error) {

	var policyIdRule []interface{}
	for _, policyIdItem := range policyId {
		policyIdRule = append(policyIdRule, policyIdItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "PolicyExpired", policyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolPolicyExpired)
				if err := _PayungPool.contract.UnpackLog(event, "PolicyExpired", log); err != nil {
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

// ParsePolicyExpired is a log parse operation binding the contract event 0xb22aa541487fc93d989bb5b17d1e3a967f649c4979cfd658c0849e070ae0dca0.
//
// Solidity: event PolicyExpired(uint256 indexed policyId)
func (_PayungPool *PayungPoolFilterer) ParsePolicyExpired(log types.Log) (*PayungPoolPolicyExpired, error) {
	event := new(PayungPoolPolicyExpired)
	if err := _PayungPool.contract.UnpackLog(event, "PolicyExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolPoolFundedIterator is returned from FilterPoolFunded and is used to iterate over the raw logs and unpacked data for PoolFunded events raised by the PayungPool contract.
type PayungPoolPoolFundedIterator struct {
	Event *PayungPoolPoolFunded // Event containing the contract specifics and raw log

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
func (it *PayungPoolPoolFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolPoolFunded)
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
		it.Event = new(PayungPoolPoolFunded)
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
func (it *PayungPoolPoolFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolPoolFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolPoolFunded represents a PoolFunded event raised by the PayungPool contract.
type PayungPoolPoolFunded struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolFunded is a free log retrieval operation binding the contract event 0x32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e838.
//
// Solidity: event PoolFunded(address indexed from, uint256 amount)
func (_PayungPool *PayungPoolFilterer) FilterPoolFunded(opts *bind.FilterOpts, from []common.Address) (*PayungPoolPoolFundedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "PoolFunded", fromRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolPoolFundedIterator{contract: _PayungPool.contract, event: "PoolFunded", logs: logs, sub: sub}, nil
}

// WatchPoolFunded is a free log subscription operation binding the contract event 0x32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e838.
//
// Solidity: event PoolFunded(address indexed from, uint256 amount)
func (_PayungPool *PayungPoolFilterer) WatchPoolFunded(opts *bind.WatchOpts, sink chan<- *PayungPoolPoolFunded, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "PoolFunded", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolPoolFunded)
				if err := _PayungPool.contract.UnpackLog(event, "PoolFunded", log); err != nil {
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

// ParsePoolFunded is a log parse operation binding the contract event 0x32173d8e51cec3a6fe484b7a1c3febe760cdf96e03d4cca36a43563a4333e838.
//
// Solidity: event PoolFunded(address indexed from, uint256 amount)
func (_PayungPool *PayungPoolFilterer) ParsePoolFunded(log types.Log) (*PayungPoolPoolFunded, error) {
	event := new(PayungPoolPoolFunded)
	if err := _PayungPool.contract.UnpackLog(event, "PoolFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolPremiumUpdatedIterator is returned from FilterPremiumUpdated and is used to iterate over the raw logs and unpacked data for PremiumUpdated events raised by the PayungPool contract.
type PayungPoolPremiumUpdatedIterator struct {
	Event *PayungPoolPremiumUpdated // Event containing the contract specifics and raw log

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
func (it *PayungPoolPremiumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolPremiumUpdated)
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
		it.Event = new(PayungPoolPremiumUpdated)
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
func (it *PayungPoolPremiumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolPremiumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolPremiumUpdated represents a PremiumUpdated event raised by the PayungPool contract.
type PayungPoolPremiumUpdated struct {
	ZoneId         uint16
	PremiumPerWeek *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPremiumUpdated is a free log retrieval operation binding the contract event 0xe173514f96c06f42ab23a57f86c0dab4648aaa8e90d56f8102d7e5b454b2ca33.
//
// Solidity: event PremiumUpdated(uint16 indexed zoneId, uint256 premiumPerWeek)
func (_PayungPool *PayungPoolFilterer) FilterPremiumUpdated(opts *bind.FilterOpts, zoneId []uint16) (*PayungPoolPremiumUpdatedIterator, error) {

	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "PremiumUpdated", zoneIdRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolPremiumUpdatedIterator{contract: _PayungPool.contract, event: "PremiumUpdated", logs: logs, sub: sub}, nil
}

// WatchPremiumUpdated is a free log subscription operation binding the contract event 0xe173514f96c06f42ab23a57f86c0dab4648aaa8e90d56f8102d7e5b454b2ca33.
//
// Solidity: event PremiumUpdated(uint16 indexed zoneId, uint256 premiumPerWeek)
func (_PayungPool *PayungPoolFilterer) WatchPremiumUpdated(opts *bind.WatchOpts, sink chan<- *PayungPoolPremiumUpdated, zoneId []uint16) (event.Subscription, error) {

	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "PremiumUpdated", zoneIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolPremiumUpdated)
				if err := _PayungPool.contract.UnpackLog(event, "PremiumUpdated", log); err != nil {
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

// ParsePremiumUpdated is a log parse operation binding the contract event 0xe173514f96c06f42ab23a57f86c0dab4648aaa8e90d56f8102d7e5b454b2ca33.
//
// Solidity: event PremiumUpdated(uint16 indexed zoneId, uint256 premiumPerWeek)
func (_PayungPool *PayungPoolFilterer) ParsePremiumUpdated(log types.Log) (*PayungPoolPremiumUpdated, error) {
	event := new(PayungPoolPremiumUpdated)
	if err := _PayungPool.contract.UnpackLog(event, "PremiumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolRainfallReportedIterator is returned from FilterRainfallReported and is used to iterate over the raw logs and unpacked data for RainfallReported events raised by the PayungPool contract.
type PayungPoolRainfallReportedIterator struct {
	Event *PayungPoolRainfallReported // Event containing the contract specifics and raw log

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
func (it *PayungPoolRainfallReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolRainfallReported)
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
		it.Event = new(PayungPoolRainfallReported)
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
func (it *PayungPoolRainfallReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolRainfallReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolRainfallReported represents a RainfallReported event raised by the PayungPool contract.
type PayungPoolRainfallReported struct {
	ZoneId    uint16
	DayIndex  uint32
	Mm        uint16
	IsRainDay bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRainfallReported is a free log retrieval operation binding the contract event 0x47ff5b1805ac3ac6341e7e3056e5a7d3c6996fc489fcaa14226abb1c75cd6005.
//
// Solidity: event RainfallReported(uint16 indexed zoneId, uint32 indexed dayIndex, uint16 mm, bool isRainDay)
func (_PayungPool *PayungPoolFilterer) FilterRainfallReported(opts *bind.FilterOpts, zoneId []uint16, dayIndex []uint32) (*PayungPoolRainfallReportedIterator, error) {

	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}
	var dayIndexRule []interface{}
	for _, dayIndexItem := range dayIndex {
		dayIndexRule = append(dayIndexRule, dayIndexItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "RainfallReported", zoneIdRule, dayIndexRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolRainfallReportedIterator{contract: _PayungPool.contract, event: "RainfallReported", logs: logs, sub: sub}, nil
}

// WatchRainfallReported is a free log subscription operation binding the contract event 0x47ff5b1805ac3ac6341e7e3056e5a7d3c6996fc489fcaa14226abb1c75cd6005.
//
// Solidity: event RainfallReported(uint16 indexed zoneId, uint32 indexed dayIndex, uint16 mm, bool isRainDay)
func (_PayungPool *PayungPoolFilterer) WatchRainfallReported(opts *bind.WatchOpts, sink chan<- *PayungPoolRainfallReported, zoneId []uint16, dayIndex []uint32) (event.Subscription, error) {

	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}
	var dayIndexRule []interface{}
	for _, dayIndexItem := range dayIndex {
		dayIndexRule = append(dayIndexRule, dayIndexItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "RainfallReported", zoneIdRule, dayIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolRainfallReported)
				if err := _PayungPool.contract.UnpackLog(event, "RainfallReported", log); err != nil {
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

// ParseRainfallReported is a log parse operation binding the contract event 0x47ff5b1805ac3ac6341e7e3056e5a7d3c6996fc489fcaa14226abb1c75cd6005.
//
// Solidity: event RainfallReported(uint16 indexed zoneId, uint32 indexed dayIndex, uint16 mm, bool isRainDay)
func (_PayungPool *PayungPoolFilterer) ParseRainfallReported(log types.Log) (*PayungPoolRainfallReported, error) {
	event := new(PayungPoolRainfallReported)
	if err := _PayungPool.contract.UnpackLog(event, "RainfallReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PayungPool contract.
type PayungPoolRoleAdminChangedIterator struct {
	Event *PayungPoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PayungPoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolRoleAdminChanged)
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
		it.Event = new(PayungPoolRoleAdminChanged)
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
func (it *PayungPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolRoleAdminChanged represents a RoleAdminChanged event raised by the PayungPool contract.
type PayungPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PayungPool *PayungPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PayungPoolRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolRoleAdminChangedIterator{contract: _PayungPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PayungPool *PayungPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PayungPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolRoleAdminChanged)
				if err := _PayungPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PayungPool *PayungPoolFilterer) ParseRoleAdminChanged(log types.Log) (*PayungPoolRoleAdminChanged, error) {
	event := new(PayungPoolRoleAdminChanged)
	if err := _PayungPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PayungPool contract.
type PayungPoolRoleGrantedIterator struct {
	Event *PayungPoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *PayungPoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolRoleGranted)
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
		it.Event = new(PayungPoolRoleGranted)
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
func (it *PayungPoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolRoleGranted represents a RoleGranted event raised by the PayungPool contract.
type PayungPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PayungPool *PayungPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PayungPoolRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolRoleGrantedIterator{contract: _PayungPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PayungPool *PayungPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PayungPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolRoleGranted)
				if err := _PayungPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PayungPool *PayungPoolFilterer) ParseRoleGranted(log types.Log) (*PayungPoolRoleGranted, error) {
	event := new(PayungPoolRoleGranted)
	if err := _PayungPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PayungPool contract.
type PayungPoolRoleRevokedIterator struct {
	Event *PayungPoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PayungPoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolRoleRevoked)
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
		it.Event = new(PayungPoolRoleRevoked)
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
func (it *PayungPoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolRoleRevoked represents a RoleRevoked event raised by the PayungPool contract.
type PayungPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PayungPool *PayungPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PayungPoolRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolRoleRevokedIterator{contract: _PayungPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PayungPool *PayungPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PayungPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolRoleRevoked)
				if err := _PayungPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PayungPool *PayungPoolFilterer) ParseRoleRevoked(log types.Log) (*PayungPoolRoleRevoked, error) {
	event := new(PayungPoolRoleRevoked)
	if err := _PayungPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayungPoolZoneCreatedIterator is returned from FilterZoneCreated and is used to iterate over the raw logs and unpacked data for ZoneCreated events raised by the PayungPool contract.
type PayungPoolZoneCreatedIterator struct {
	Event *PayungPoolZoneCreated // Event containing the contract specifics and raw log

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
func (it *PayungPoolZoneCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayungPoolZoneCreated)
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
		it.Event = new(PayungPoolZoneCreated)
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
func (it *PayungPoolZoneCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayungPoolZoneCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayungPoolZoneCreated represents a ZoneCreated event raised by the PayungPool contract.
type PayungPoolZoneCreated struct {
	ZoneId uint16
	Name   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterZoneCreated is a free log retrieval operation binding the contract event 0x2da346d3337ba631daf0758afccc0fbf2013c400760bc620e3b3078e866c8e5c.
//
// Solidity: event ZoneCreated(uint16 indexed zoneId, string name)
func (_PayungPool *PayungPoolFilterer) FilterZoneCreated(opts *bind.FilterOpts, zoneId []uint16) (*PayungPoolZoneCreatedIterator, error) {

	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.FilterLogs(opts, "ZoneCreated", zoneIdRule)
	if err != nil {
		return nil, err
	}
	return &PayungPoolZoneCreatedIterator{contract: _PayungPool.contract, event: "ZoneCreated", logs: logs, sub: sub}, nil
}

// WatchZoneCreated is a free log subscription operation binding the contract event 0x2da346d3337ba631daf0758afccc0fbf2013c400760bc620e3b3078e866c8e5c.
//
// Solidity: event ZoneCreated(uint16 indexed zoneId, string name)
func (_PayungPool *PayungPoolFilterer) WatchZoneCreated(opts *bind.WatchOpts, sink chan<- *PayungPoolZoneCreated, zoneId []uint16) (event.Subscription, error) {

	var zoneIdRule []interface{}
	for _, zoneIdItem := range zoneId {
		zoneIdRule = append(zoneIdRule, zoneIdItem)
	}

	logs, sub, err := _PayungPool.contract.WatchLogs(opts, "ZoneCreated", zoneIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayungPoolZoneCreated)
				if err := _PayungPool.contract.UnpackLog(event, "ZoneCreated", log); err != nil {
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

// ParseZoneCreated is a log parse operation binding the contract event 0x2da346d3337ba631daf0758afccc0fbf2013c400760bc620e3b3078e866c8e5c.
//
// Solidity: event ZoneCreated(uint16 indexed zoneId, string name)
func (_PayungPool *PayungPoolFilterer) ParseZoneCreated(log types.Log) (*PayungPoolZoneCreated, error) {
	event := new(PayungPoolZoneCreated)
	if err := _PayungPool.contract.UnpackLog(event, "ZoneCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
