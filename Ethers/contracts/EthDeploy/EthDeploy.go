// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EthDeploy

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

// EthDeployMetaData contains all meta data concerning the EthDeploy contract.
var EthDeployMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GetData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"SetData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061059a8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806376b8e52814610038578063b7e8260c14610056575b5f5ffd5b610040610072565b60405161004d9190610186565b60405180910390f35b610070600480360381019061006b919061020f565b610101565b005b60605f805461008090610287565b80601f01602080910402602001604051908101604052809291908181526020018280546100ac90610287565b80156100f75780601f106100ce576101008083540402835291602001916100f7565b820191905f5260205f20905b8154815290600101906020018083116100da57829003601f168201915b5050505050905090565b81815f9182610111929190610497565b505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61015882610116565b6101628185610120565b9350610172818560208601610130565b61017b8161013e565b840191505092915050565b5f6020820190508181035f83015261019e818461014e565b905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126101cf576101ce6101ae565b5b8235905067ffffffffffffffff8111156101ec576101eb6101b2565b5b602083019150836001820283011115610208576102076101b6565b5b9250929050565b5f5f60208385031215610225576102246101a6565b5b5f83013567ffffffffffffffff811115610242576102416101aa565b5b61024e858286016101ba565b92509250509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061029e57607f821691505b6020821081036102b1576102b061025a565b5b50919050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261034a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261030f565b610354868361030f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61039861039361038e8461036c565b610375565b61036c565b9050919050565b5f819050919050565b6103b18361037e565b6103c56103bd8261039f565b84845461031b565b825550505050565b5f5f905090565b6103dc6103cd565b6103e78184846103a8565b505050565b5b8181101561040a576103ff5f826103d4565b6001810190506103ed565b5050565b601f82111561044f57610420816102ee565b61042984610300565b81016020851015610438578190505b61044c61044485610300565b8301826103ec565b50505b505050565b5f82821c905092915050565b5f61046f5f1984600802610454565b1980831691505092915050565b5f6104878383610460565b9150826002028217905092915050565b6104a183836102b7565b67ffffffffffffffff8111156104ba576104b96102c1565b5b6104c48254610287565b6104cf82828561040e565b5f601f8311600181146104fc575f84156104ea578287013590505b6104f4858261047c565b86555061055b565b601f19841661050a866102ee565b5f5b828110156105315784890135825560018201915060208501945060208101905061050c565b8683101561054e578489013561054a601f891682610460565b8355505b6001600288020188555050505b5050505050505056fea26469706673582212208ee6cf447b7895649d888b18450894ef42bdecd7e244cf5428804ad11526f07264736f6c634300081e0033",
}

// EthDeployABI is the input ABI used to generate the binding from.
// Deprecated: Use EthDeployMetaData.ABI instead.
var EthDeployABI = EthDeployMetaData.ABI

// EthDeployBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthDeployMetaData.Bin instead.
var EthDeployBin = EthDeployMetaData.Bin

// DeployEthDeploy deploys a new Ethereum contract, binding an instance of EthDeploy to it.
func DeployEthDeploy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthDeploy, error) {
	parsed, err := EthDeployMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthDeployBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthDeploy{EthDeployCaller: EthDeployCaller{contract: contract}, EthDeployTransactor: EthDeployTransactor{contract: contract}, EthDeployFilterer: EthDeployFilterer{contract: contract}}, nil
}

// EthDeploy is an auto generated Go binding around an Ethereum contract.
type EthDeploy struct {
	EthDeployCaller     // Read-only binding to the contract
	EthDeployTransactor // Write-only binding to the contract
	EthDeployFilterer   // Log filterer for contract events
}

// EthDeployCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthDeployCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDeployTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthDeployTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDeployFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthDeployFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthDeploySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthDeploySession struct {
	Contract     *EthDeploy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthDeployCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthDeployCallerSession struct {
	Contract *EthDeployCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthDeployTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthDeployTransactorSession struct {
	Contract     *EthDeployTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthDeployRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthDeployRaw struct {
	Contract *EthDeploy // Generic contract binding to access the raw methods on
}

// EthDeployCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthDeployCallerRaw struct {
	Contract *EthDeployCaller // Generic read-only contract binding to access the raw methods on
}

// EthDeployTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthDeployTransactorRaw struct {
	Contract *EthDeployTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthDeploy creates a new instance of EthDeploy, bound to a specific deployed contract.
func NewEthDeploy(address common.Address, backend bind.ContractBackend) (*EthDeploy, error) {
	contract, err := bindEthDeploy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthDeploy{EthDeployCaller: EthDeployCaller{contract: contract}, EthDeployTransactor: EthDeployTransactor{contract: contract}, EthDeployFilterer: EthDeployFilterer{contract: contract}}, nil
}

// NewEthDeployCaller creates a new read-only instance of EthDeploy, bound to a specific deployed contract.
func NewEthDeployCaller(address common.Address, caller bind.ContractCaller) (*EthDeployCaller, error) {
	contract, err := bindEthDeploy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthDeployCaller{contract: contract}, nil
}

// NewEthDeployTransactor creates a new write-only instance of EthDeploy, bound to a specific deployed contract.
func NewEthDeployTransactor(address common.Address, transactor bind.ContractTransactor) (*EthDeployTransactor, error) {
	contract, err := bindEthDeploy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthDeployTransactor{contract: contract}, nil
}

// NewEthDeployFilterer creates a new log filterer instance of EthDeploy, bound to a specific deployed contract.
func NewEthDeployFilterer(address common.Address, filterer bind.ContractFilterer) (*EthDeployFilterer, error) {
	contract, err := bindEthDeploy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthDeployFilterer{contract: contract}, nil
}

// bindEthDeploy binds a generic wrapper to an already deployed contract.
func bindEthDeploy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthDeployMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDeploy *EthDeployRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthDeploy.Contract.EthDeployCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDeploy *EthDeployRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDeploy.Contract.EthDeployTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDeploy *EthDeployRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDeploy.Contract.EthDeployTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthDeploy *EthDeployCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthDeploy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthDeploy *EthDeployTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthDeploy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthDeploy *EthDeployTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthDeploy.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(string)
func (_EthDeploy *EthDeployCaller) GetData(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthDeploy.contract.Call(opts, &out, "GetData")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(string)
func (_EthDeploy *EthDeploySession) GetData() (string, error) {
	return _EthDeploy.Contract.GetData(&_EthDeploy.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(string)
func (_EthDeploy *EthDeployCallerSession) GetData() (string, error) {
	return _EthDeploy.Contract.GetData(&_EthDeploy.CallOpts)
}

// SetData is a paid mutator transaction binding the contract method 0xb7e8260c.
//
// Solidity: function SetData(string _data) returns()
func (_EthDeploy *EthDeployTransactor) SetData(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _EthDeploy.contract.Transact(opts, "SetData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0xb7e8260c.
//
// Solidity: function SetData(string _data) returns()
func (_EthDeploy *EthDeploySession) SetData(_data string) (*types.Transaction, error) {
	return _EthDeploy.Contract.SetData(&_EthDeploy.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0xb7e8260c.
//
// Solidity: function SetData(string _data) returns()
func (_EthDeploy *EthDeployTransactorSession) SetData(_data string) (*types.Transaction, error) {
	return _EthDeploy.Contract.SetData(&_EthDeploy.TransactOpts, _data)
}
