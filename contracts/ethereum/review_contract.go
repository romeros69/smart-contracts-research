// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// ReviewContractReview is an auto generated low-level Go binding around an user-defined struct.
type ReviewContractReview struct {
	ProductId string
	Reviewer  common.Address
	Rating    uint8
	Comment   string
}

// EthereumMetaData contains all meta data concerning the Ethereum contract.
var EthereumMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reviewer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rating\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"ReviewSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"}],\"name\":\"getAverageRating\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"}],\"name\":\"getProductReviews\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"reviewer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"rating\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"internalType\":\"structReviewContract.Review[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"hasReviewed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"productReviews\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"reviewer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"rating\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"rating\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"submitReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610dce8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80632f5f42b914610059578063998863fd146100855780639cef29bf146100aa578063a8cb9a42146100f8578063f1ba707214610118575b5f80fd5b61006c61006736600461094d565b61012d565b60405161007c94939291906109bd565b60405180910390f35b610098610093366004610a04565b6102a1565b60405160ff909116815260200161007c565b6100e86100b8366004610a36565b600160209081525f9283526040909220815180830184018051928152908401929093019190912091525460ff1681565b604051901515815260200161007c565b61010b610106366004610a04565b6104c0565b60405161007c9190610a8e565b61012b610126366004610b37565b61067f565b005b81516020818401810180515f825292820191850191909120919052805482908110610156575f80fd5b905f5260205f2090600302015f9150915050805f01805461017690610baf565b80601f01602080910402602001604051908101604052809291908181526020018280546101a290610baf565b80156101ed5780601f106101c4576101008083540402835291602001916101ed565b820191905f5260205f20905b8154815290600101906020018083116101d057829003601f168201915b50505050600183015460028401805493946001600160a01b03831694600160a01b90930460ff1693509161022090610baf565b80601f016020809104026020016040519081016040528092919081815260200182805461024c90610baf565b80156102975780601f1061026e57610100808354040283529160200191610297565b820191905f5260205f20905b81548152906001019060200180831161027a57829003601f168201915b5050505050905084565b5f805f836040516102b29190610be7565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020015f905b82821015610455578382905f5260205f2090600302016040518060800160405290815f8201805461030f90610baf565b80601f016020809104026020016040519081016040528092919081815260200182805461033b90610baf565b80156103865780601f1061035d57610100808354040283529160200191610386565b820191905f5260205f20905b81548152906001019060200180831161036957829003601f168201915b505050918352505060018201546001600160a01b0381166020830152600160a01b900460ff1660408201526002820180546060909201916103c690610baf565b80601f01602080910402602001604051908101604052809291908181526020018280546103f290610baf565b801561043d5780601f106104145761010080835404028352916020019161043d565b820191905f5260205f20905b81548152906001019060200180831161042057829003601f168201915b505050505081525050815260200190600101906102df565b5050505090505f805b825181101561049d5782818151811061047957610479610bfd565b60200260200101516040015160ff16826104939190610c11565b915060010161045e565b505f8251116104ac575f6104b8565b81516104b89082610c36565b949350505050565b60605f826040516104d19190610be7565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020015f905b82821015610674578382905f5260205f2090600302016040518060800160405290815f8201805461052e90610baf565b80601f016020809104026020016040519081016040528092919081815260200182805461055a90610baf565b80156105a55780601f1061057c576101008083540402835291602001916105a5565b820191905f5260205f20905b81548152906001019060200180831161058857829003601f168201915b505050918352505060018201546001600160a01b0381166020830152600160a01b900460ff1660408201526002820180546060909201916105e590610baf565b80601f016020809104026020016040519081016040528092919081815260200182805461061190610baf565b801561065c5780601f106106335761010080835404028352916020019161065c565b820191905f5260205f20905b81548152906001019060200180831161063f57829003601f168201915b505050505081525050815260200190600101906104fe565b505050509050919050565b60018260ff1610158015610697575060058260ff1611155b6106e85760405162461bcd60e51b815260206004820152601e60248201527f526174696e67206d757374206265206265747765656e203120616e642035000060448201526064015b60405180910390fd5b335f90815260016020526040908190209051610705908590610be7565b9081526040519081900360200190205460ff16156107745760405162461bcd60e51b815260206004820152602660248201527f596f75206861766520616c7265616479207265766965776564207468697320706044820152651c9bd91d58dd60d21b60648201526084016106df565b6040805160808101825284815233602082015260ff8416818301526060810183905290515f906107a5908690610be7565b9081526040516020918190038201902080546001810182555f918252919020825183926003029091019081906107db9082610ca1565b506020820151600182018054604085015160ff16600160a01b026001600160a81b03199091166001600160a01b03909316929092179190911790556060820151600282019061082a9082610ca1565b5050335f90815260016020819052604091829020915190925061084e908790610be7565b908152604051908190036020018120805492151560ff199093169290921790915533907f4a278f4079c2c80f15994199897bd8aedfb0f89e2e876a7644751b9099ba941e906108a290879087908790610d61565b60405180910390a250505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126108d3575f80fd5b813567ffffffffffffffff808211156108ee576108ee6108b0565b604051601f8301601f19908116603f01168101908282118183101715610916576109166108b0565b8160405283815286602085880101111561092e575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f806040838503121561095e575f80fd5b823567ffffffffffffffff811115610974575f80fd5b610980858286016108c4565b95602094909401359450505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b608081525f6109cf608083018761098f565b6001600160a01b038616602084015260ff8516604084015282810360608401526109f9818561098f565b979650505050505050565b5f60208284031215610a14575f80fd5b813567ffffffffffffffff811115610a2a575f80fd5b6104b8848285016108c4565b5f8060408385031215610a47575f80fd5b82356001600160a01b0381168114610a5d575f80fd5b9150602083013567ffffffffffffffff811115610a78575f80fd5b610a84858286016108c4565b9150509250929050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b83811015610b2957603f19898403018552815160808151818652610adb8287018261098f565b838b01516001600160a01b0316878c01528984015160ff168a880152606093840151878203948801949094529150610b159050818361098f565b968901969450505090860190600101610ab5565b509098975050505050505050565b5f805f60608486031215610b49575f80fd5b833567ffffffffffffffff80821115610b60575f80fd5b610b6c878388016108c4565b94506020860135915060ff82168214610b83575f80fd5b90925060408501359080821115610b98575f80fd5b50610ba5868287016108c4565b9150509250925092565b600181811c90821680610bc357607f821691505b602082108103610be157634e487b7160e01b5f52602260045260245ffd5b50919050565b5f82518060208501845e5f920191825250919050565b634e487b7160e01b5f52603260045260245ffd5b80820180821115610c3057634e487b7160e01b5f52601160045260245ffd5b92915050565b5f82610c5057634e487b7160e01b5f52601260045260245ffd5b500490565b601f821115610c9c57805f5260205f20601f840160051c81016020851015610c7a5750805b601f840160051c820191505b81811015610c99575f8155600101610c86565b50505b505050565b815167ffffffffffffffff811115610cbb57610cbb6108b0565b610ccf81610cc98454610baf565b84610c55565b602080601f831160018114610d02575f8415610ceb5750858301515b5f19600386901b1c1916600185901b178555610d59565b5f85815260208120601f198616915b82811015610d3057888601518255948401946001909101908401610d11565b5085821015610d4d57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b606081525f610d73606083018661098f565b60ff851660208401528281036040840152610d8e818561098f565b969550505050505056fea2646970667358221220e57dc39f740fc9cd13e2c739a4a9e7b87056ce083e87eec9423adf19112677ab64736f6c63430008190033",
}

// EthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumMetaData.ABI instead.
var EthereumABI = EthereumMetaData.ABI

// EthereumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthereumMetaData.Bin instead.
var EthereumBin = EthereumMetaData.Bin

// DeployEthereum deploys a new Ethereum contract, binding an instance of Ethereum to it.
func DeployEthereum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethereum, error) {
	parsed, err := EthereumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthereumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// Ethereum is an auto generated Go binding around an Ethereum contract.
type Ethereum struct {
	EthereumCaller     // Read-only binding to the contract
	EthereumTransactor // Write-only binding to the contract
	EthereumFilterer   // Log filterer for contract events
}

// EthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumSession struct {
	Contract     *Ethereum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumCallerSession struct {
	Contract *EthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumTransactorSession struct {
	Contract     *EthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumRaw struct {
	Contract *Ethereum // Generic contract binding to access the raw methods on
}

// EthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumCallerRaw struct {
	Contract *EthereumCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumTransactorRaw struct {
	Contract *EthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereum creates a new instance of Ethereum, bound to a specific deployed contract.
func NewEthereum(address common.Address, backend bind.ContractBackend) (*Ethereum, error) {
	contract, err := bindEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// NewEthereumCaller creates a new read-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumCaller(address common.Address, caller bind.ContractCaller) (*EthereumCaller, error) {
	contract, err := bindEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumCaller{contract: contract}, nil
}

// NewEthereumTransactor creates a new write-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumTransactor, error) {
	contract, err := bindEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumTransactor{contract: contract}, nil
}

// NewEthereumFilterer creates a new log filterer instance of Ethereum, bound to a specific deployed contract.
func NewEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumFilterer, error) {
	contract, err := bindEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumFilterer{contract: contract}, nil
}

// bindEthereum binds a generic wrapper to an already deployed contract.
func bindEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthereumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.EthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transact(opts, method, params...)
}

// GetAverageRating is a free data retrieval call binding the contract method 0x998863fd.
//
// Solidity: function getAverageRating(string productId) view returns(uint8)
func (_Ethereum *EthereumCaller) GetAverageRating(opts *bind.CallOpts, productId string) (uint8, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getAverageRating", productId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAverageRating is a free data retrieval call binding the contract method 0x998863fd.
//
// Solidity: function getAverageRating(string productId) view returns(uint8)
func (_Ethereum *EthereumSession) GetAverageRating(productId string) (uint8, error) {
	return _Ethereum.Contract.GetAverageRating(&_Ethereum.CallOpts, productId)
}

// GetAverageRating is a free data retrieval call binding the contract method 0x998863fd.
//
// Solidity: function getAverageRating(string productId) view returns(uint8)
func (_Ethereum *EthereumCallerSession) GetAverageRating(productId string) (uint8, error) {
	return _Ethereum.Contract.GetAverageRating(&_Ethereum.CallOpts, productId)
}

// GetProductReviews is a free data retrieval call binding the contract method 0xa8cb9a42.
//
// Solidity: function getProductReviews(string productId) view returns((string,address,uint8,string)[])
func (_Ethereum *EthereumCaller) GetProductReviews(opts *bind.CallOpts, productId string) ([]ReviewContractReview, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getProductReviews", productId)

	if err != nil {
		return *new([]ReviewContractReview), err
	}

	out0 := *abi.ConvertType(out[0], new([]ReviewContractReview)).(*[]ReviewContractReview)

	return out0, err

}

// GetProductReviews is a free data retrieval call binding the contract method 0xa8cb9a42.
//
// Solidity: function getProductReviews(string productId) view returns((string,address,uint8,string)[])
func (_Ethereum *EthereumSession) GetProductReviews(productId string) ([]ReviewContractReview, error) {
	return _Ethereum.Contract.GetProductReviews(&_Ethereum.CallOpts, productId)
}

// GetProductReviews is a free data retrieval call binding the contract method 0xa8cb9a42.
//
// Solidity: function getProductReviews(string productId) view returns((string,address,uint8,string)[])
func (_Ethereum *EthereumCallerSession) GetProductReviews(productId string) ([]ReviewContractReview, error) {
	return _Ethereum.Contract.GetProductReviews(&_Ethereum.CallOpts, productId)
}

// HasReviewed is a free data retrieval call binding the contract method 0x9cef29bf.
//
// Solidity: function hasReviewed(address , string ) view returns(bool)
func (_Ethereum *EthereumCaller) HasReviewed(opts *bind.CallOpts, arg0 common.Address, arg1 string) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "hasReviewed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasReviewed is a free data retrieval call binding the contract method 0x9cef29bf.
//
// Solidity: function hasReviewed(address , string ) view returns(bool)
func (_Ethereum *EthereumSession) HasReviewed(arg0 common.Address, arg1 string) (bool, error) {
	return _Ethereum.Contract.HasReviewed(&_Ethereum.CallOpts, arg0, arg1)
}

// HasReviewed is a free data retrieval call binding the contract method 0x9cef29bf.
//
// Solidity: function hasReviewed(address , string ) view returns(bool)
func (_Ethereum *EthereumCallerSession) HasReviewed(arg0 common.Address, arg1 string) (bool, error) {
	return _Ethereum.Contract.HasReviewed(&_Ethereum.CallOpts, arg0, arg1)
}

// ProductReviews is a free data retrieval call binding the contract method 0x2f5f42b9.
//
// Solidity: function productReviews(string , uint256 ) view returns(string productId, address reviewer, uint8 rating, string comment)
func (_Ethereum *EthereumCaller) ProductReviews(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	ProductId string
	Reviewer  common.Address
	Rating    uint8
	Comment   string
}, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "productReviews", arg0, arg1)

	outstruct := new(struct {
		ProductId string
		Reviewer  common.Address
		Rating    uint8
		Comment   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProductId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Reviewer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Rating = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Comment = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// ProductReviews is a free data retrieval call binding the contract method 0x2f5f42b9.
//
// Solidity: function productReviews(string , uint256 ) view returns(string productId, address reviewer, uint8 rating, string comment)
func (_Ethereum *EthereumSession) ProductReviews(arg0 string, arg1 *big.Int) (struct {
	ProductId string
	Reviewer  common.Address
	Rating    uint8
	Comment   string
}, error) {
	return _Ethereum.Contract.ProductReviews(&_Ethereum.CallOpts, arg0, arg1)
}

// ProductReviews is a free data retrieval call binding the contract method 0x2f5f42b9.
//
// Solidity: function productReviews(string , uint256 ) view returns(string productId, address reviewer, uint8 rating, string comment)
func (_Ethereum *EthereumCallerSession) ProductReviews(arg0 string, arg1 *big.Int) (struct {
	ProductId string
	Reviewer  common.Address
	Rating    uint8
	Comment   string
}, error) {
	return _Ethereum.Contract.ProductReviews(&_Ethereum.CallOpts, arg0, arg1)
}

// SubmitReview is a paid mutator transaction binding the contract method 0xf1ba7072.
//
// Solidity: function submitReview(string productId, uint8 rating, string comment) returns()
func (_Ethereum *EthereumTransactor) SubmitReview(opts *bind.TransactOpts, productId string, rating uint8, comment string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "submitReview", productId, rating, comment)
}

// SubmitReview is a paid mutator transaction binding the contract method 0xf1ba7072.
//
// Solidity: function submitReview(string productId, uint8 rating, string comment) returns()
func (_Ethereum *EthereumSession) SubmitReview(productId string, rating uint8, comment string) (*types.Transaction, error) {
	return _Ethereum.Contract.SubmitReview(&_Ethereum.TransactOpts, productId, rating, comment)
}

// SubmitReview is a paid mutator transaction binding the contract method 0xf1ba7072.
//
// Solidity: function submitReview(string productId, uint8 rating, string comment) returns()
func (_Ethereum *EthereumTransactorSession) SubmitReview(productId string, rating uint8, comment string) (*types.Transaction, error) {
	return _Ethereum.Contract.SubmitReview(&_Ethereum.TransactOpts, productId, rating, comment)
}

// EthereumReviewSubmittedIterator is returned from FilterReviewSubmitted and is used to iterate over the raw logs and unpacked data for ReviewSubmitted events raised by the Ethereum contract.
type EthereumReviewSubmittedIterator struct {
	Event *EthereumReviewSubmitted // Event containing the contract specifics and raw log

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
func (it *EthereumReviewSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumReviewSubmitted)
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
		it.Event = new(EthereumReviewSubmitted)
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
func (it *EthereumReviewSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumReviewSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumReviewSubmitted represents a ReviewSubmitted event raised by the Ethereum contract.
type EthereumReviewSubmitted struct {
	ProductId string
	Reviewer  common.Address
	Rating    uint8
	Comment   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReviewSubmitted is a free log retrieval operation binding the contract event 0x4a278f4079c2c80f15994199897bd8aedfb0f89e2e876a7644751b9099ba941e.
//
// Solidity: event ReviewSubmitted(string productId, address indexed reviewer, uint8 rating, string comment)
func (_Ethereum *EthereumFilterer) FilterReviewSubmitted(opts *bind.FilterOpts, reviewer []common.Address) (*EthereumReviewSubmittedIterator, error) {

	var reviewerRule []interface{}
	for _, reviewerItem := range reviewer {
		reviewerRule = append(reviewerRule, reviewerItem)
	}

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "ReviewSubmitted", reviewerRule)
	if err != nil {
		return nil, err
	}
	return &EthereumReviewSubmittedIterator{contract: _Ethereum.contract, event: "ReviewSubmitted", logs: logs, sub: sub}, nil
}

// WatchReviewSubmitted is a free log subscription operation binding the contract event 0x4a278f4079c2c80f15994199897bd8aedfb0f89e2e876a7644751b9099ba941e.
//
// Solidity: event ReviewSubmitted(string productId, address indexed reviewer, uint8 rating, string comment)
func (_Ethereum *EthereumFilterer) WatchReviewSubmitted(opts *bind.WatchOpts, sink chan<- *EthereumReviewSubmitted, reviewer []common.Address) (event.Subscription, error) {

	var reviewerRule []interface{}
	for _, reviewerItem := range reviewer {
		reviewerRule = append(reviewerRule, reviewerItem)
	}

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "ReviewSubmitted", reviewerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumReviewSubmitted)
				if err := _Ethereum.contract.UnpackLog(event, "ReviewSubmitted", log); err != nil {
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

// ParseReviewSubmitted is a log parse operation binding the contract event 0x4a278f4079c2c80f15994199897bd8aedfb0f89e2e876a7644751b9099ba941e.
//
// Solidity: event ReviewSubmitted(string productId, address indexed reviewer, uint8 rating, string comment)
func (_Ethereum *EthereumFilterer) ParseReviewSubmitted(log types.Log) (*EthereumReviewSubmitted, error) {
	event := new(EthereumReviewSubmitted)
	if err := _Ethereum.contract.UnpackLog(event, "ReviewSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
