// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package online_test

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

// OnlineTestMetaData contains all meta data concerning the OnlineTest contract.
var OnlineTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"adminAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminEntranceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"}],\"name\":\"caculateWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_adminEntranceFee\",\"type\":\"uint256\"}],\"name\":\"chooseAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userSubmitFee\",\"type\":\"uint256\"}],\"name\":\"createExams\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"}],\"name\":\"endSubmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"examCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"examCodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"examFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exams\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"adminAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"userSubmitFee\",\"type\":\"uint256\"},{\"internalType\":\"enumOnlineTest.EXAM_STATE\",\"name\":\"examState\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_answer\",\"type\":\"string\"}],\"name\":\"submitAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_answer\",\"type\":\"string\"}],\"name\":\"submitTest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userExams\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"examCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submitTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600060055534801561001557600080fd5b50600080546001600160a01b031916331790556117af806100376000396000f3fe6080604052600436106101025760003560e01c80635ca0b3f311610095578063931486c611610064578063931486c6146102835780639bf2871f146102a3578063ce606ee0146102b9578063d0c9e799146102d9578063e567a4e0146102ef57600080fd5b80635ca0b3f3146101e75780636896fabf146102075780637cb2bb281461021b578063818305931461024b57600080fd5b80632e3f45af116100d15780632e3f45af146101785780633ccfd60b146101985780634198cb3f146101a0578063581c591b146101d157600080fd5b806301fa3bec1461010e57806312065fe0146101305780631506a8581461015257806317029f161461016557600080fd5b3661010957005b600080fd5b34801561011a57600080fd5b5061012e610129366004611308565b61030f565b005b34801561013c57600080fd5b50475b6040519081526020015b60405180910390f35b61012e6101603660046113c3565b610442565b61012e610173366004611308565b61065f565b34801561018457600080fd5b5061013f6101933660046113e5565b6108f3565b61012e610914565b3480156101ac57600080fd5b506101c06101bb3660046113e5565b6109c2565b60405161014995949392919061145a565b3480156101dd57600080fd5b5061013f60035481565b3480156101f357600080fd5b5061012e6102023660046113e5565b610a87565b34801561021357600080fd5b50333161013f565b34801561022757600080fd5b5061023b6102363660046113e5565b610fee565b60405161014994939291906114b7565b34801561025757600080fd5b5060015461026b906001600160a01b031681565b6040516001600160a01b039091168152602001610149565b34801561028f57600080fd5b5061012e61029e3660046114ef565b6110bd565b3480156102af57600080fd5b5061013f60055481565b3480156102c557600080fd5b5060005461026b906001600160a01b031681565b3480156102e557600080fd5b5061013f60025481565b3480156102fb57600080fd5b5061012e61030a3660046113e5565b611151565b6000828152600960205260409020546001600160a01b031633146103895760405162461bcd60e51b815260206004820152602660248201527f596f75206e65656420746f2062652041646d696e20746f207375626d697420616044820152656e737765722160d01b60648201526084015b60405180910390fd5b600260008381526009602052604090206004015460ff1660038111156103b1576103b1611444565b146103fe5760405162461bcd60e51b815260206004820152601e60248201527f41646d696e2043616e2774207375626d697420616e73776572207965742100006044820152606401610380565b600082815260096020526040902060020161041982826115b0565b50600082815260096020526040902060040180546003919060ff19166001835b02179055505050565b6001546001600160a01b031661045757600080fd5b60025460000361046657600080fd5b6001546001600160a01b031633146104c05760405162461bcd60e51b815260206004820181905260248201527f596f75206d7573742062652041646d696e20746f20437265617465204578616d6044820152606401610380565b6002546104cd9083611686565b3331116105335760405162461bcd60e51b815260206004820152602e60248201527f4e6f7420656e6f75676820457468657220746f20437265617465204578616d2060448201526d30b732102830bc902bb4b73732b960911b6064820152608401610380565b6002543410156105ab5760405162461bcd60e51b815260206004820152603860248201527f56616c75652061646d696e2073656e6420746f20636f6e7472616374206d757360448201527f742067726561746572207468616e2061646d696e2066656500000000000000006064820152608401610380565b600580549060006105bb8361169f565b90915550506005805460048054600180820183557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90910183905560009283526009602052604080842080546001600160a01b031916331790558454845280842090910186905583548352808320909101805460ff191690559154815290812060039081018390556002548154909290610656908490611686565b90915550505050565b60008281526009602052604081206004015460ff16600381111561068557610685611444565b146106c75760405162461bcd60e51b815260206004820152601260248201527143616e6e6f74207375626d6974207965742160701b6044820152606401610380565b600082815260096020526040902060030154333110156107375760405162461bcd60e51b815260206004820152602560248201527f4e6f7420656e6f75676820457468657220746f207375626d697420796f7572206044820152646578616d2160d81b6064820152608401610380565b6000828152600960205260409020600301543410156107be5760405162461bcd60e51b815260206004820152603d60248201527f56616c756520757365722073656e6420746f20636f6e7472616374206d75737460448201527f2067726561746572207468616e2075736572207375626d6974206665650000006064820152608401610380565b6107f2604051806080016040528060006001600160a01b031681526020016000815260200160608152602001600081525090565b33815260208101838152604082018381524260608401526008805460018101825560009190915283517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3600490920291820180546001600160a01b0319166001600160a01b0390921691909117815592517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee482015590518392917ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee501906108b990826115b0565b506060919091015160039182015560008481526009602052604081208201548254909291906108e9908490611686565b9091555050505050565b6004818154811061090357600080fd5b600091825260209091200154905081565b6000546001600160a01b031633146109845760405162461bcd60e51b815260206004820152602d60248201527f596f75206d75737420626520636f6e7472616374206f776e657220746f20776960448201526c746864726177206d6f6e65792160981b6064820152608401610380565b600080546003546040516001600160a01b039092169281156108fc029290818181858888f193505050501580156109bf573d6000803e3d6000fd5b50565b6009602052600090815260409020805460018201546002830180546001600160a01b039093169391926109f490611527565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2090611527565b8015610a6d5780601f10610a4257610100808354040283529160200191610a6d565b820191906000526020600020905b815481529060010190602001808311610a5057829003601f168201915b50505050600383015460049093015491929160ff16905085565b6000818152600960205260409020546001600160a01b03163314610aed5760405162461bcd60e51b815260206004820152601f60248201527f4f6e6c792061646d696e2063616e20636163756c6174652077696e6e657221006044820152606401610380565b600360008281526009602052604090206004015460ff166003811115610b1557610b15611444565b14610b6c5760405162461bcd60e51b815260206004820152602160248201527f41646d696e20646f6573206e6f74207375626d697420616e73776572207965746044820152602160f81b6064820152608401610380565b60005b600854811015610d2d578160088281548110610b8d57610b8d6116b8565b90600052602060002090600402016001015403610d1b576000828152600960209081526040918290209151610bc69260020191016116ce565b6040516020818303038152906040528051906020012060088281548110610bef57610bef6116b8565b9060005260206000209060040201600201604051602001610c1091906116ce565b6040516020818303038152906040528051906020012003610d1b57604080518082019091526000808252602082015260088281548110610c5257610c526116b8565b60009182526020909120600490910201546001600160a01b031681526008805483908110610c8257610c826116b8565b600091825260208083206004929092029091016003015490830190815260068054600181018255925291517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f600290920291820180546001600160a01b0319166001600160a01b0390921691909117905590517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909101555b80610d258161169f565b915050610b6f565b506000805b600654610d4190600190611744565b811015610e75576000610d55826001611686565b90505b600654811015610e625760068181548110610d7557610d756116b8565b90600052602060002090600202016001015460068381548110610d9a57610d9a6116b8565b9060005260206000209060020201600101541115610e505760068281548110610dc557610dc56116b8565b906000526020600020906002020160010154925060068181548110610dec57610dec6116b8565b90600052602060002090600202016001015460068381548110610e1157610e116116b8565b9060005260206000209060020201600101819055508260068281548110610e3a57610e3a6116b8565b9060005260206000209060020201600101819055505b80610e5a8161169f565b915050610d58565b5080610e6d8161169f565b915050610d32565b50600654600a11610f1e57600082815260096020526040812060010154610e9e90600a90611757565b905060005b600a811015610f175760068181548110610ebf57610ebf6116b8565b600091825260208220600290910201546040516001600160a01b039091169184156108fc02918591818181858888f19350505050158015610f04573d6000803e3d6000fd5b5080610f0f8161169f565b915050610ea3565b5050610fbb565b6006546000838152600960205260408120600101549091610f3e91611757565b905060005b600654811015610fb85760068181548110610f6057610f606116b8565b600091825260208220600290910201546040516001600160a01b039091169184156108fc02918591818181858888f19350505050158015610fa5573d6000803e3d6000fd5b5080610fb08161169f565b915050610f43565b50505b60078054610fcb91600691611254565b50600082815260096020526040902060040180546001919060ff19168280610439565b60088181548110610ffe57600080fd5b60009182526020909120600490910201805460018201546002830180546001600160a01b03909316945090929161103490611527565b80601f016020809104026020016040519081016040528092919081815260200182805461106090611527565b80156110ad5780601f10611082576101008083540402835291602001916110ad565b820191906000526020600020905b81548152906001019060200180831161109057829003601f168201915b5050505050908060030154905084565b6000546001600160a01b0316331461112b5760405162461bcd60e51b815260206004820152602b60248201527f546f2063686f6f73652041646d696e20796f75206d75737420626520636f6e7460448201526a72616374206f776e65722160a81b6064820152608401610380565b600180546001600160a01b0319166001600160a01b039390931692909217909155600255565b6000818152600960205260409020546001600160a01b031633146111c35760405162461bcd60e51b815260206004820152602360248201527f596f75206e65656420746f2062652041646d696e20746f20656e64207375626d60448201526269742160e81b6064820152608401610380565b60008181526009602052604081206004015460ff1660038111156111e9576111e9611444565b146112365760405162461bcd60e51b815260206004820152601b60248201527f41646d696e2043616e277420656e64207375626d6974207965742100000000006044820152606401610380565b6000908152600960205260409020600401805460ff19166002179055565b8280548282559060005260206000209060020281019282156112bc5760005260206000209160020282015b828111156112bc57825482546001600160a01b0319166001600160a01b03909116178255600180840154908301556002928301929091019061127f565b506112c89291506112cc565b5090565b5b808211156112c85780546001600160a01b0319168155600060018201556002016112cd565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561131b57600080fd5b82359150602083013567ffffffffffffffff8082111561133a57600080fd5b818501915085601f83011261134e57600080fd5b813581811115611360576113606112f2565b604051601f8201601f19908116603f01168101908382118183101715611388576113886112f2565b816040528281528860208487010111156113a157600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600080604083850312156113d657600080fd5b50508035926020909101359150565b6000602082840312156113f757600080fd5b5035919050565b6000815180845260005b8181101561142457602081850181015186830182015201611408565b506000602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b600052602160045260246000fd5b60018060a01b038616815284602082015260a06040820152600061148160a08301866113fe565b9050836060830152600483106114a757634e487b7160e01b600052602160045260246000fd5b8260808301529695505050505050565b60018060a01b03851681528360208201526080604082015260006114de60808301856113fe565b905082606083015295945050505050565b6000806040838503121561150257600080fd5b82356001600160a01b038116811461151957600080fd5b946020939093013593505050565b600181811c9082168061153b57607f821691505b60208210810361155b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156115ab57600081815260208120601f850160051c810160208610156115885750805b601f850160051c820191505b818110156115a757828155600101611594565b5050505b505050565b815167ffffffffffffffff8111156115ca576115ca6112f2565b6115de816115d88454611527565b84611561565b602080601f83116001811461161357600084156115fb5750858301515b600019600386901b1c1916600185901b1785556115a7565b600085815260208120601f198616915b8281101561164257888601518255948401946001909101908401611623565b50858210156116605787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b8082018082111561169957611699611670565b92915050565b6000600182016116b1576116b1611670565b5060010190565b634e487b7160e01b600052603260045260246000fd5b60008083546116dc81611527565b600182811680156116f4576001811461170957611738565b60ff1984168752821515830287019450611738565b8760005260208060002060005b8581101561172f5781548a820152908401908201611716565b50505082870194505b50929695505050505050565b8181038181111561169957611699611670565b60008261177457634e487b7160e01b600052601260045260246000fd5b50049056fea264697066735822122000892b1ecbfb6817ee4d2a2b02c8cd8ac76bc854b802cb8efa05d5af2b7d8e4464736f6c63430008110033",
}

// OnlineTestABI is the input ABI used to generate the binding from.
// Deprecated: Use OnlineTestMetaData.ABI instead.
var OnlineTestABI = OnlineTestMetaData.ABI

// OnlineTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OnlineTestMetaData.Bin instead.
var OnlineTestBin = OnlineTestMetaData.Bin

// DeployOnlineTest deploys a new Ethereum contract, binding an instance of OnlineTest to it.
func DeployOnlineTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OnlineTest, error) {
	parsed, err := OnlineTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnlineTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnlineTest{OnlineTestCaller: OnlineTestCaller{contract: contract}, OnlineTestTransactor: OnlineTestTransactor{contract: contract}, OnlineTestFilterer: OnlineTestFilterer{contract: contract}}, nil
}

// OnlineTest is an auto generated Go binding around an Ethereum contract.
type OnlineTest struct {
	OnlineTestCaller     // Read-only binding to the contract
	OnlineTestTransactor // Write-only binding to the contract
	OnlineTestFilterer   // Log filterer for contract events
}

// OnlineTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnlineTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnlineTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnlineTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnlineTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnlineTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnlineTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnlineTestSession struct {
	Contract     *OnlineTest       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnlineTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnlineTestCallerSession struct {
	Contract *OnlineTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OnlineTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnlineTestTransactorSession struct {
	Contract     *OnlineTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OnlineTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnlineTestRaw struct {
	Contract *OnlineTest // Generic contract binding to access the raw methods on
}

// OnlineTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnlineTestCallerRaw struct {
	Contract *OnlineTestCaller // Generic read-only contract binding to access the raw methods on
}

// OnlineTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnlineTestTransactorRaw struct {
	Contract *OnlineTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnlineTest creates a new instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTest(address common.Address, backend bind.ContractBackend) (*OnlineTest, error) {
	contract, err := bindOnlineTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnlineTest{OnlineTestCaller: OnlineTestCaller{contract: contract}, OnlineTestTransactor: OnlineTestTransactor{contract: contract}, OnlineTestFilterer: OnlineTestFilterer{contract: contract}}, nil
}

// NewOnlineTestCaller creates a new read-only instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTestCaller(address common.Address, caller bind.ContractCaller) (*OnlineTestCaller, error) {
	contract, err := bindOnlineTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnlineTestCaller{contract: contract}, nil
}

// NewOnlineTestTransactor creates a new write-only instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTestTransactor(address common.Address, transactor bind.ContractTransactor) (*OnlineTestTransactor, error) {
	contract, err := bindOnlineTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnlineTestTransactor{contract: contract}, nil
}

// NewOnlineTestFilterer creates a new log filterer instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTestFilterer(address common.Address, filterer bind.ContractFilterer) (*OnlineTestFilterer, error) {
	contract, err := bindOnlineTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnlineTestFilterer{contract: contract}, nil
}

// bindOnlineTest binds a generic wrapper to an already deployed contract.
func bindOnlineTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnlineTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnlineTest *OnlineTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnlineTest.Contract.OnlineTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnlineTest *OnlineTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.Contract.OnlineTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnlineTest *OnlineTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnlineTest.Contract.OnlineTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnlineTest *OnlineTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnlineTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnlineTest *OnlineTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnlineTest *OnlineTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnlineTest.Contract.contract.Transact(opts, method, params...)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_OnlineTest *OnlineTestCaller) AdminAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "adminAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_OnlineTest *OnlineTestSession) AdminAddr() (common.Address, error) {
	return _OnlineTest.Contract.AdminAddr(&_OnlineTest.CallOpts)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_OnlineTest *OnlineTestCallerSession) AdminAddr() (common.Address, error) {
	return _OnlineTest.Contract.AdminAddr(&_OnlineTest.CallOpts)
}

// AdminEntranceFee is a free data retrieval call binding the contract method 0xd0c9e799.
//
// Solidity: function adminEntranceFee() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) AdminEntranceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "adminEntranceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminEntranceFee is a free data retrieval call binding the contract method 0xd0c9e799.
//
// Solidity: function adminEntranceFee() view returns(uint256)
func (_OnlineTest *OnlineTestSession) AdminEntranceFee() (*big.Int, error) {
	return _OnlineTest.Contract.AdminEntranceFee(&_OnlineTest.CallOpts)
}

// AdminEntranceFee is a free data retrieval call binding the contract method 0xd0c9e799.
//
// Solidity: function adminEntranceFee() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) AdminEntranceFee() (*big.Int, error) {
	return _OnlineTest.Contract.AdminEntranceFee(&_OnlineTest.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_OnlineTest *OnlineTestCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "contractOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_OnlineTest *OnlineTestSession) ContractOwner() (common.Address, error) {
	return _OnlineTest.Contract.ContractOwner(&_OnlineTest.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_OnlineTest *OnlineTestCallerSession) ContractOwner() (common.Address, error) {
	return _OnlineTest.Contract.ContractOwner(&_OnlineTest.CallOpts)
}

// ExamCode is a free data retrieval call binding the contract method 0x9bf2871f.
//
// Solidity: function examCode() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) ExamCode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "examCode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamCode is a free data retrieval call binding the contract method 0x9bf2871f.
//
// Solidity: function examCode() view returns(uint256)
func (_OnlineTest *OnlineTestSession) ExamCode() (*big.Int, error) {
	return _OnlineTest.Contract.ExamCode(&_OnlineTest.CallOpts)
}

// ExamCode is a free data retrieval call binding the contract method 0x9bf2871f.
//
// Solidity: function examCode() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) ExamCode() (*big.Int, error) {
	return _OnlineTest.Contract.ExamCode(&_OnlineTest.CallOpts)
}

// ExamCodes is a free data retrieval call binding the contract method 0x2e3f45af.
//
// Solidity: function examCodes(uint256 ) view returns(uint256)
func (_OnlineTest *OnlineTestCaller) ExamCodes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "examCodes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamCodes is a free data retrieval call binding the contract method 0x2e3f45af.
//
// Solidity: function examCodes(uint256 ) view returns(uint256)
func (_OnlineTest *OnlineTestSession) ExamCodes(arg0 *big.Int) (*big.Int, error) {
	return _OnlineTest.Contract.ExamCodes(&_OnlineTest.CallOpts, arg0)
}

// ExamCodes is a free data retrieval call binding the contract method 0x2e3f45af.
//
// Solidity: function examCodes(uint256 ) view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) ExamCodes(arg0 *big.Int) (*big.Int, error) {
	return _OnlineTest.Contract.ExamCodes(&_OnlineTest.CallOpts, arg0)
}

// ExamFund is a free data retrieval call binding the contract method 0x581c591b.
//
// Solidity: function examFund() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) ExamFund(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "examFund")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamFund is a free data retrieval call binding the contract method 0x581c591b.
//
// Solidity: function examFund() view returns(uint256)
func (_OnlineTest *OnlineTestSession) ExamFund() (*big.Int, error) {
	return _OnlineTest.Contract.ExamFund(&_OnlineTest.CallOpts)
}

// ExamFund is a free data retrieval call binding the contract method 0x581c591b.
//
// Solidity: function examFund() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) ExamFund() (*big.Int, error) {
	return _OnlineTest.Contract.ExamFund(&_OnlineTest.CallOpts)
}

// Exams is a free data retrieval call binding the contract method 0x4198cb3f.
//
// Solidity: function exams(uint256 ) view returns(address adminAddr, uint256 reward, string answer, uint256 userSubmitFee, uint8 examState)
func (_OnlineTest *OnlineTestCaller) Exams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AdminAddr     common.Address
	Reward        *big.Int
	Answer        string
	UserSubmitFee *big.Int
	ExamState     uint8
}, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "exams", arg0)

	outstruct := new(struct {
		AdminAddr     common.Address
		Reward        *big.Int
		Answer        string
		UserSubmitFee *big.Int
		ExamState     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AdminAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.UserSubmitFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExamState = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// Exams is a free data retrieval call binding the contract method 0x4198cb3f.
//
// Solidity: function exams(uint256 ) view returns(address adminAddr, uint256 reward, string answer, uint256 userSubmitFee, uint8 examState)
func (_OnlineTest *OnlineTestSession) Exams(arg0 *big.Int) (struct {
	AdminAddr     common.Address
	Reward        *big.Int
	Answer        string
	UserSubmitFee *big.Int
	ExamState     uint8
}, error) {
	return _OnlineTest.Contract.Exams(&_OnlineTest.CallOpts, arg0)
}

// Exams is a free data retrieval call binding the contract method 0x4198cb3f.
//
// Solidity: function exams(uint256 ) view returns(address adminAddr, uint256 reward, string answer, uint256 userSubmitFee, uint8 examState)
func (_OnlineTest *OnlineTestCallerSession) Exams(arg0 *big.Int) (struct {
	AdminAddr     common.Address
	Reward        *big.Int
	Answer        string
	UserSubmitFee *big.Int
	ExamState     uint8
}, error) {
	return _OnlineTest.Contract.Exams(&_OnlineTest.CallOpts, arg0)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) GetAccountBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "getAccountBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_OnlineTest *OnlineTestSession) GetAccountBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetAccountBalance(&_OnlineTest.CallOpts)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) GetAccountBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetAccountBalance(&_OnlineTest.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_OnlineTest *OnlineTestSession) GetBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetBalance(&_OnlineTest.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) GetBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetBalance(&_OnlineTest.CallOpts)
}

// UserExams is a free data retrieval call binding the contract method 0x7cb2bb28.
//
// Solidity: function userExams(uint256 ) view returns(address userAddr, uint256 examCode, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestCaller) UserExams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UserAddr   common.Address
	ExamCode   *big.Int
	Answer     string
	SubmitTime *big.Int
}, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "userExams", arg0)

	outstruct := new(struct {
		UserAddr   common.Address
		ExamCode   *big.Int
		Answer     string
		SubmitTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ExamCode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SubmitTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserExams is a free data retrieval call binding the contract method 0x7cb2bb28.
//
// Solidity: function userExams(uint256 ) view returns(address userAddr, uint256 examCode, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestSession) UserExams(arg0 *big.Int) (struct {
	UserAddr   common.Address
	ExamCode   *big.Int
	Answer     string
	SubmitTime *big.Int
}, error) {
	return _OnlineTest.Contract.UserExams(&_OnlineTest.CallOpts, arg0)
}

// UserExams is a free data retrieval call binding the contract method 0x7cb2bb28.
//
// Solidity: function userExams(uint256 ) view returns(address userAddr, uint256 examCode, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestCallerSession) UserExams(arg0 *big.Int) (struct {
	UserAddr   common.Address
	ExamCode   *big.Int
	Answer     string
	SubmitTime *big.Int
}, error) {
	return _OnlineTest.Contract.UserExams(&_OnlineTest.CallOpts, arg0)
}

// CaculateWinner is a paid mutator transaction binding the contract method 0x5ca0b3f3.
//
// Solidity: function caculateWinner(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactor) CaculateWinner(opts *bind.TransactOpts, _examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "caculateWinner", _examCode)
}

// CaculateWinner is a paid mutator transaction binding the contract method 0x5ca0b3f3.
//
// Solidity: function caculateWinner(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestSession) CaculateWinner(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CaculateWinner(&_OnlineTest.TransactOpts, _examCode)
}

// CaculateWinner is a paid mutator transaction binding the contract method 0x5ca0b3f3.
//
// Solidity: function caculateWinner(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactorSession) CaculateWinner(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CaculateWinner(&_OnlineTest.TransactOpts, _examCode)
}

// ChooseAdmin is a paid mutator transaction binding the contract method 0x931486c6.
//
// Solidity: function chooseAdmin(address _adminAddr, uint256 _adminEntranceFee) returns()
func (_OnlineTest *OnlineTestTransactor) ChooseAdmin(opts *bind.TransactOpts, _adminAddr common.Address, _adminEntranceFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "chooseAdmin", _adminAddr, _adminEntranceFee)
}

// ChooseAdmin is a paid mutator transaction binding the contract method 0x931486c6.
//
// Solidity: function chooseAdmin(address _adminAddr, uint256 _adminEntranceFee) returns()
func (_OnlineTest *OnlineTestSession) ChooseAdmin(_adminAddr common.Address, _adminEntranceFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.ChooseAdmin(&_OnlineTest.TransactOpts, _adminAddr, _adminEntranceFee)
}

// ChooseAdmin is a paid mutator transaction binding the contract method 0x931486c6.
//
// Solidity: function chooseAdmin(address _adminAddr, uint256 _adminEntranceFee) returns()
func (_OnlineTest *OnlineTestTransactorSession) ChooseAdmin(_adminAddr common.Address, _adminEntranceFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.ChooseAdmin(&_OnlineTest.TransactOpts, _adminAddr, _adminEntranceFee)
}

// CreateExams is a paid mutator transaction binding the contract method 0x1506a858.
//
// Solidity: function createExams(uint256 _reward, uint256 _userSubmitFee) payable returns()
func (_OnlineTest *OnlineTestTransactor) CreateExams(opts *bind.TransactOpts, _reward *big.Int, _userSubmitFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "createExams", _reward, _userSubmitFee)
}

// CreateExams is a paid mutator transaction binding the contract method 0x1506a858.
//
// Solidity: function createExams(uint256 _reward, uint256 _userSubmitFee) payable returns()
func (_OnlineTest *OnlineTestSession) CreateExams(_reward *big.Int, _userSubmitFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CreateExams(&_OnlineTest.TransactOpts, _reward, _userSubmitFee)
}

// CreateExams is a paid mutator transaction binding the contract method 0x1506a858.
//
// Solidity: function createExams(uint256 _reward, uint256 _userSubmitFee) payable returns()
func (_OnlineTest *OnlineTestTransactorSession) CreateExams(_reward *big.Int, _userSubmitFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CreateExams(&_OnlineTest.TransactOpts, _reward, _userSubmitFee)
}

// EndSubmit is a paid mutator transaction binding the contract method 0xe567a4e0.
//
// Solidity: function endSubmit(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactor) EndSubmit(opts *bind.TransactOpts, _examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "endSubmit", _examCode)
}

// EndSubmit is a paid mutator transaction binding the contract method 0xe567a4e0.
//
// Solidity: function endSubmit(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestSession) EndSubmit(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.EndSubmit(&_OnlineTest.TransactOpts, _examCode)
}

// EndSubmit is a paid mutator transaction binding the contract method 0xe567a4e0.
//
// Solidity: function endSubmit(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactorSession) EndSubmit(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.EndSubmit(&_OnlineTest.TransactOpts, _examCode)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 _examCode, string _answer) returns()
func (_OnlineTest *OnlineTestTransactor) SubmitAnswer(opts *bind.TransactOpts, _examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "submitAnswer", _examCode, _answer)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 _examCode, string _answer) returns()
func (_OnlineTest *OnlineTestSession) SubmitAnswer(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitAnswer(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 _examCode, string _answer) returns()
func (_OnlineTest *OnlineTestTransactorSession) SubmitAnswer(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitAnswer(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// SubmitTest is a paid mutator transaction binding the contract method 0x17029f16.
//
// Solidity: function submitTest(uint256 _examCode, string _answer) payable returns()
func (_OnlineTest *OnlineTestTransactor) SubmitTest(opts *bind.TransactOpts, _examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "submitTest", _examCode, _answer)
}

// SubmitTest is a paid mutator transaction binding the contract method 0x17029f16.
//
// Solidity: function submitTest(uint256 _examCode, string _answer) payable returns()
func (_OnlineTest *OnlineTestSession) SubmitTest(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitTest(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// SubmitTest is a paid mutator transaction binding the contract method 0x17029f16.
//
// Solidity: function submitTest(uint256 _examCode, string _answer) payable returns()
func (_OnlineTest *OnlineTestTransactorSession) SubmitTest(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitTest(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_OnlineTest *OnlineTestTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_OnlineTest *OnlineTestSession) Withdraw() (*types.Transaction, error) {
	return _OnlineTest.Contract.Withdraw(&_OnlineTest.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_OnlineTest *OnlineTestTransactorSession) Withdraw() (*types.Transaction, error) {
	return _OnlineTest.Contract.Withdraw(&_OnlineTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnlineTest *OnlineTestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnlineTest *OnlineTestSession) Receive() (*types.Transaction, error) {
	return _OnlineTest.Contract.Receive(&_OnlineTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnlineTest *OnlineTestTransactorSession) Receive() (*types.Transaction, error) {
	return _OnlineTest.Contract.Receive(&_OnlineTest.TransactOpts)
}
