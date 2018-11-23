// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SimpleLoanABI is the input ABI used to generate the binding from.
const SimpleLoanABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updateCollateralPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"stableCoinReceiver\",\"type\":\"bytes\"},{\"name\":\"request\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"sendCollateral\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"part\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"name\":\"debtAmount\",\"type\":\"uint256\"}],\"name\":\"collateralRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"addPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"bytes32\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"refundCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"name\":\"debtAmount\",\"type\":\"uint256\"}],\"name\":\"safelyCollateralized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"bytes32\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"liquidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updateAssetPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"loans\",\"outputs\":[{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"request\",\"type\":\"uint256\"},{\"name\":\"principle\",\"type\":\"uint256\"},{\"name\":\"interest\",\"type\":\"uint256\"},{\"name\":\"maturityDate\",\"type\":\"uint256\"},{\"name\":\"escrowDeadline\",\"type\":\"uint256\"},{\"name\":\"stableCoinReceiver\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"bytes32\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"rejectLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"params\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lid\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"acceptLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_lender\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__sendCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__acceptLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__rejectLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__refundCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__addPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__update\",\"type\":\"event\"}]"

// SimpleLoanBin is the compiled bytecode used for deploying new contracts.
const SimpleLoanBin = `0x608060405260646003819055614e2060045560055534801561002057600080fd5b506040516020806117418339810180604052602081101561004057600080fd5b505160008054600160a060020a031916600160a060020a0390921691909117815560026020526276a7007fdcece27f2f847054190196a10112a278f5c96dce59b21ea677f4141abd5fa386556202a3007fb95c855550ee3929cd717905a1f177a1b5585b8ce0ad32f737099643d62738f4556003547f1b1a10d3881692e0f4e4796096b4d3149390aa241eec76f15942f3502d0ab8c3819055609681027fa9fb08b0619fbf0fdbc0fd7e78cab49de4d989021285eb5b692c34100a19357a55606481027f1e0624433fc3f2022f2f932d64a0a1b79f40dbbf013f9e3b97342b66b2e11cf855600a81027fda2fc955906013a0058c7feb1082d3d112993cd51194ddd6551c69aa525dbc0a8190557f63e0784acdb4f42428d644c31f1ee8c32a0030019f4ac5dc204875714ae3ceba557f6d6178436f6d6d697373696f6e0000000000000000000000000000000000000082526014027f9a2c49f738157c707fd08a48ed3bf980dceadc9e427be548057ef9debac61dfa5561157a9081906101c790396000f3fe6080604052600436106101065763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166323592f86811461010b578063313ce567146101375780633adca01c1461015e5780634527864414610215578063532cd106146102455780635891de72146102755780635d43b0ae1461028a5780635e8bc142146102c057806367ebd110146102f05780637805ffa5146103345780638eaa6ac0146103645780638eccc8ce1461038e578063b3f52fec146103b8578063bcead63e146103ee578063c4a908151461041f578063d24378eb14610520578063d31eba5d14610535578063dc6ab52714610565578063e0d214371461058f575b600080fd5b34801561011757600080fd5b506101356004803603602081101561012e57600080fd5b50356105c5565b005b34801561014357600080fd5b5061014c6105e1565b60408051918252519081900360200190f35b610135600480360360a081101561017457600080fd5b81359160208101359181019060608101604082013564010000000081111561019b57600080fd5b8201836020820111156101ad57600080fd5b803590602001918460018302840111640100000000831117156101cf57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356105e7565b34801561022157600080fd5b5061014c6004803603604081101561023857600080fd5b5080359060200135610875565b34801561025157600080fd5b5061014c6004803603604081101561026857600080fd5b508035906020013561089a565b34801561028157600080fd5b5061014c6108d6565b34801561029657600080fd5b50610135600480360360608110156102ad57600080fd5b50803590602081013590604001356108dc565b3480156102cc57600080fd5b50610135600480360360408110156102e357600080fd5b5080359060200135610a96565b3480156102fc57600080fd5b506103206004803603604081101561031357600080fd5b5080359060200135610c14565b604080519115158252519081900360200190f35b34801561034057600080fd5b506101356004803603604081101561035757600080fd5b5080359060200135610c74565b34801561037057600080fd5b5061014c6004803603602081101561038757600080fd5b5035610f96565b34801561039a57600080fd5b50610135600480360360208110156103b157600080fd5b5035610fa8565b3480156103c457600080fd5b50610135600480360360608110156103db57600080fd5b5080359060208101359060400135610fc4565b3480156103fa57600080fd5b50610403611032565b60408051600160a060020a039092168252519081900360200190f35b34801561042b57600080fd5b506104496004803603602081101561044257600080fd5b5035611041565b604051808b600581111561045957fe5b60ff1681526020018a600160a060020a0316600160a060020a0316815260200189815260200188815260200187815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156104dc5781810151838201526020016104c4565b50505050905090810190601f1680156105095780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b34801561052c57600080fd5b5061014c61114d565b34801561054157600080fd5b506101356004803603604081101561055857600080fd5b5080359060200135611153565b34801561057157600080fd5b5061014c6004803603602081101561058857600080fd5b50356111ec565b34801561059b57600080fd5b50610135600480360360608110156105b257600080fd5b50803590602081013590604001356111fe565b600054600160a060020a031633146105dc57600080fd5b600455565b60035481565b841561065b57600160008681526001602052604090205460ff16600581111561060c57fe5b14806106345750600260008681526001602052604090205460ff16600581111561063257fe5b145b151561063f57600080fd5b600085815260016020526040902060020180543401905561082f565b8383836040516020018084815260200183805190602001908083835b602083106106965780518252601f199092019160209182019101610677565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201905282519201919091209750600092506106da915050565b60008681526001602052604090205460ff1660058111156106f757fe5b14801561070957506107093483610c14565b151561071457600080fd5b61071c611453565b6001808252336020808401919091526040808401889052346060850152608084018690527fb95c855550ee3929cd717905a1f177a1b5585b8ce0ad32f737099643d62738f454420161010085015261012084018790526000898152918390529020825181548493839160ff19169083600581111561079657fe5b0217905550602082810151825474ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a03909216820217835560408401516001840155606084015160028401556080840151600384015560a0840151600484015560c0840151600584015560e084015160068401558301516007830155610120830151805161082a92600885019201906114b3565b505050505b6040805186815234602082015280820183905290517f474f79dc33c29e99abd873b5114b4d572f88d3a28fbc53cbc252cbb145cf2e3f9181900360600190a15050505050565b6000606460035483850281151561088857fe5b0481151561089257fe5b049392505050565b6005546000908202670de0b6b3a7640000028015156108b7575060015b806003546004548602606402028115156108cd57fe5b04949350505050565b60045481565b600054600160a060020a031633146108f357600080fd5b600260008481526001602052604090205460ff16600581111561091257fe5b1461091c57600080fd5b6000838152600160205260409020600581015460048201547fdcece27f2f847054190196a10112a278f5c96dce59b21ea677f4141abd5fa3865460069093015491929091839190428201106109875785841161097957600061097d565b8584035b9150818403860395505b600086841161099757600061099b565b8684035b6000898152600160205260409020600401819055905082156109d0576000888152600160205260409020600501839055610a51565b600088815260016020908152604082206006018054850190557f696e746572657374526174650000000000000000000000000000000000000000909152600290527f1b1a10d3881692e0f4e4796096b4d3149390aa241eec76f15942f3502d0ab8c354610a3e908290610875565b6000898152600160205260409020600501555b604080518981526020810188905281517f6f17325e5a83dd2569c62cfbdba7cb9b19e73b7d46b7d37aa0acf1cc8ffb9d6a929181900390910190a15050505050505050565b600360008381526001602052604090205460ff166005811115610ab557fe5b1480610add5750600560008381526001602052604090205460ff166005811115610adb57fe5b145b80610b1f5750600160008381526001602052604090205460ff166005811115610b0257fe5b148015610b1f575060008281526001602052604090206007015442115b80610b605750600260008381526001602052604090205460ff166005811115610b4457fe5b148015610b605750600082815260016020526040902060040154155b1515610b6b57600080fd5b6000828152600160205260408082208054600460ff1990911617808255600290910180549084905591519192610100909104600160a060020a0316916108fc84150291849190818181858888f19350505050158015610bce573d6000803e3d6000fd5b50604080518481526020810183905280820184905290517f260ef70d5caed6f134782cfa7f0db59b4f7ce3726cdf88e644aeaa898edbae549181900360600190a1505050565b7f6c69717569646174696f6e537461727400000000000000000000000000000000600090815260026020527fa9fb08b0619fbf0fdbc0fd7e78cab49de4d989021285eb5b692c34100a19357a54610c6b848461089a565b10159392505050565b60008281526001602052604090206005808201546004830154925492019160029160ff90911690811115610ca457fe5b148015610cc05750600083815260016020526040812060040154115b8015610cff5750600083815260016020526040902060060154421180610cff5750600083815260016020526040902060020154610cfd9082610c14565b155b1515610d0a57600080fd5b60006004546005548302670de0b6b3a764000002811515610d2757fe5b0490506000610d6982600260007f6c69717569646174696f6e50656e616c74790000000000000000000000000000815260200190815260200160002054610875565b6000868152600160205260409020600201549091508282019080821115610d8e578091505b6000610d9a828761089a565b60026020527f9a2c49f738157c707fd08a48ed3bf980dceadc9e427be548057ef9debac61dfa547f63e0784acdb4f42428d644c31f1ee8c32a0030019f4ac5dc204875714ae3ceba547fa9fb08b0619fbf0fdbc0fd7e78cab49de4d989021285eb5b692c34100a19357a547f6c69717569646174696f6e456e6400000000000000000000000000000000000060009081527f1e0624433fc3f2022f2f932d64a0a1b79f40dbbf013f9e3b97342b66b2e11cf85494955092939192909190610e618986610875565b905081861015610e7c57610e758985610875565b9050610ea6565b82861015610ea657610ea389838503868803858a0302811515610e9b57fe5b048601610875565b90505b60008d8152600160205260408082206004810183905560058082018490556002820180548d90039055815460ff19161790555133916108fc841502918491818181858888f19350505050158015610f01573d6000803e3d6000fd5b5060008054604051600160a060020a0390911691838b0380156108fc02929091818181858888f19350505050158015610f3e573d6000803e3d6000fd5b50604080518e8152828a036020820152808201839052606081018e905290517f064e21237eec8c3f4bdf95fe7d379974f5ce7292b9454fa13301ed92a443c4ed9181900360800190a150505050505050505050505050565b60009081526002602052604090205490565b600054600160a060020a03163314610fbf57600080fd5b600555565b600054600160a060020a03163314610fdb57600080fd5b600083815260026020908152604091829020849055815185815290810184905280820183905290517f31eb539fa5855ce13b272ee9f9520cb20a498c755ce42a6f984fce0e19bd38b09181900360600190a1505050565b600054600160a060020a031681565b60016020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a9004600160a060020a031690806001015490806002015490806003015490806004015490806005015490806006015490806007015490806008018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111435780601f1061111857610100808354040283529160200191611143565b820191906000526020600020905b81548152906001019060200180831161112657829003601f168201915b505050505090508a565b60055481565b600054600160a060020a0316331461116a57600080fd5b600160008381526001602052604090205460ff16600581111561118957fe5b1461119357600080fd5b600082815260016020908152604091829020805460ff19166003179055815184815290810183905281517fc74197ec873ffbac3377394f9f2ac39b6b743f8592529f7e76500661ee3bd534929181900390910190a15050565b60026020526000908152604090205481565b600054600160a060020a0316331461121557600080fd5b600160008481526001602052604090205460ff16600581111561123457fe5b146112a057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7374617465206d75737420626520696e69746564000000000000000000000000604482015290519081900360640190fd5b600160008481526020019081526020016000206001015482604051602001808281526020019150506040516020818303038152906040528051906020012014151561134c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6b657920646f6573206e6f74206d617463682064696765737400000000000000604482015290519081900360640190fd5b60008381526001602090815260408220805460ff191660029081178255600382015460049092018290557f696e7465726573745261746500000000000000000000000000000000000000009093529190527f1b1a10d3881692e0f4e4796096b4d3149390aa241eec76f15942f3502d0ab8c3546113ca908290610875565b60008581526001602090815260409182902060058101939093557fdcece27f2f847054190196a10112a278f5c96dce59b21ea677f4141abd5fa3865442016006909301929092558051868152918201859052818101849052517f2caa7d07be7b7990462ca8a059568be0582ae0a101ccdcc1fdaba0da2cbf4beb9181900360600190a150505050565b60408051610140810190915280600081526020016000600160a060020a0316815260200160008019168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114f457805160ff1916838001178555611521565b82800160010185558215611521579182015b82811115611521578251825591602001919060010190611506565b5061152d929150611531565b5090565b61154b91905b8082111561152d5760008155600101611537565b9056fea165627a7a7230582013ff1a88b7d0940aeb7878e55acb5a8143015b717a16a220954ceea6ed58c54d0029`

// DeploySimpleLoan deploys a new Ethereum contract, binding an instance of SimpleLoan to it.
func DeploySimpleLoan(auth *bind.TransactOpts, backend bind.ContractBackend, _lender common.Address) (common.Address, *types.Transaction, *SimpleLoan, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleLoanABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleLoanBin), backend, _lender)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleLoan{SimpleLoanCaller: SimpleLoanCaller{contract: contract}, SimpleLoanTransactor: SimpleLoanTransactor{contract: contract}, SimpleLoanFilterer: SimpleLoanFilterer{contract: contract}}, nil
}

// SimpleLoan is an auto generated Go binding around an Ethereum contract.
type SimpleLoan struct {
	SimpleLoanCaller     // Read-only binding to the contract
	SimpleLoanTransactor // Write-only binding to the contract
	SimpleLoanFilterer   // Log filterer for contract events
}

// SimpleLoanCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleLoanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleLoanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleLoanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleLoanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleLoanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleLoanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleLoanSession struct {
	Contract     *SimpleLoan       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleLoanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleLoanCallerSession struct {
	Contract *SimpleLoanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimpleLoanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleLoanTransactorSession struct {
	Contract     *SimpleLoanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleLoanRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleLoanRaw struct {
	Contract *SimpleLoan // Generic contract binding to access the raw methods on
}

// SimpleLoanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleLoanCallerRaw struct {
	Contract *SimpleLoanCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleLoanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleLoanTransactorRaw struct {
	Contract *SimpleLoanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleLoan creates a new instance of SimpleLoan, bound to a specific deployed contract.
func NewSimpleLoan(address common.Address, backend bind.ContractBackend) (*SimpleLoan, error) {
	contract, err := bindSimpleLoan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleLoan{SimpleLoanCaller: SimpleLoanCaller{contract: contract}, SimpleLoanTransactor: SimpleLoanTransactor{contract: contract}, SimpleLoanFilterer: SimpleLoanFilterer{contract: contract}}, nil
}

// NewSimpleLoanCaller creates a new read-only instance of SimpleLoan, bound to a specific deployed contract.
func NewSimpleLoanCaller(address common.Address, caller bind.ContractCaller) (*SimpleLoanCaller, error) {
	contract, err := bindSimpleLoan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleLoanCaller{contract: contract}, nil
}

// NewSimpleLoanTransactor creates a new write-only instance of SimpleLoan, bound to a specific deployed contract.
func NewSimpleLoanTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleLoanTransactor, error) {
	contract, err := bindSimpleLoan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleLoanTransactor{contract: contract}, nil
}

// NewSimpleLoanFilterer creates a new log filterer instance of SimpleLoan, bound to a specific deployed contract.
func NewSimpleLoanFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleLoanFilterer, error) {
	contract, err := bindSimpleLoan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleLoanFilterer{contract: contract}, nil
}

// bindSimpleLoan binds a generic wrapper to an already deployed contract.
func bindSimpleLoan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleLoanABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleLoan *SimpleLoanRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleLoan.Contract.SimpleLoanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleLoan *SimpleLoanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleLoan.Contract.SimpleLoanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleLoan *SimpleLoanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleLoan.Contract.SimpleLoanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleLoan *SimpleLoanCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleLoan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleLoan *SimpleLoanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleLoan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleLoan *SimpleLoanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleLoan.Contract.contract.Transact(opts, method, params...)
}

// AssetPrice is a free data retrieval call binding the contract method 0xd24378eb.
//
// Solidity: function assetPrice() constant returns(uint256)
func (_SimpleLoan *SimpleLoanCaller) AssetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "assetPrice")
	return *ret0, err
}

// AssetPrice is a free data retrieval call binding the contract method 0xd24378eb.
//
// Solidity: function assetPrice() constant returns(uint256)
func (_SimpleLoan *SimpleLoanSession) AssetPrice() (*big.Int, error) {
	return _SimpleLoan.Contract.AssetPrice(&_SimpleLoan.CallOpts)
}

// AssetPrice is a free data retrieval call binding the contract method 0xd24378eb.
//
// Solidity: function assetPrice() constant returns(uint256)
func (_SimpleLoan *SimpleLoanCallerSession) AssetPrice() (*big.Int, error) {
	return _SimpleLoan.Contract.AssetPrice(&_SimpleLoan.CallOpts)
}

// CollateralPrice is a free data retrieval call binding the contract method 0x5891de72.
//
// Solidity: function collateralPrice() constant returns(uint256)
func (_SimpleLoan *SimpleLoanCaller) CollateralPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "collateralPrice")
	return *ret0, err
}

// CollateralPrice is a free data retrieval call binding the contract method 0x5891de72.
//
// Solidity: function collateralPrice() constant returns(uint256)
func (_SimpleLoan *SimpleLoanSession) CollateralPrice() (*big.Int, error) {
	return _SimpleLoan.Contract.CollateralPrice(&_SimpleLoan.CallOpts)
}

// CollateralPrice is a free data retrieval call binding the contract method 0x5891de72.
//
// Solidity: function collateralPrice() constant returns(uint256)
func (_SimpleLoan *SimpleLoanCallerSession) CollateralPrice() (*big.Int, error) {
	return _SimpleLoan.Contract.CollateralPrice(&_SimpleLoan.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0x532cd106.
//
// Solidity: function collateralRatio(collateralAmount uint256, debtAmount uint256) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCaller) CollateralRatio(opts *bind.CallOpts, collateralAmount *big.Int, debtAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "collateralRatio", collateralAmount, debtAmount)
	return *ret0, err
}

// CollateralRatio is a free data retrieval call binding the contract method 0x532cd106.
//
// Solidity: function collateralRatio(collateralAmount uint256, debtAmount uint256) constant returns(uint256)
func (_SimpleLoan *SimpleLoanSession) CollateralRatio(collateralAmount *big.Int, debtAmount *big.Int) (*big.Int, error) {
	return _SimpleLoan.Contract.CollateralRatio(&_SimpleLoan.CallOpts, collateralAmount, debtAmount)
}

// CollateralRatio is a free data retrieval call binding the contract method 0x532cd106.
//
// Solidity: function collateralRatio(collateralAmount uint256, debtAmount uint256) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCallerSession) CollateralRatio(collateralAmount *big.Int, debtAmount *big.Int) (*big.Int, error) {
	return _SimpleLoan.Contract.CollateralRatio(&_SimpleLoan.CallOpts, collateralAmount, debtAmount)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SimpleLoan *SimpleLoanCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SimpleLoan *SimpleLoanSession) Decimals() (*big.Int, error) {
	return _SimpleLoan.Contract.Decimals(&_SimpleLoan.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SimpleLoan *SimpleLoanCallerSession) Decimals() (*big.Int, error) {
	return _SimpleLoan.Contract.Decimals(&_SimpleLoan.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(name bytes32) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCaller) Get(opts *bind.CallOpts, name [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "get", name)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(name bytes32) constant returns(uint256)
func (_SimpleLoan *SimpleLoanSession) Get(name [32]byte) (*big.Int, error) {
	return _SimpleLoan.Contract.Get(&_SimpleLoan.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(name bytes32) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCallerSession) Get(name [32]byte) (*big.Int, error) {
	return _SimpleLoan.Contract.Get(&_SimpleLoan.CallOpts, name)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_SimpleLoan *SimpleLoanCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "lender")
	return *ret0, err
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_SimpleLoan *SimpleLoanSession) Lender() (common.Address, error) {
	return _SimpleLoan.Contract.Lender(&_SimpleLoan.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() constant returns(address)
func (_SimpleLoan *SimpleLoanCallerSession) Lender() (common.Address, error) {
	return _SimpleLoan.Contract.Lender(&_SimpleLoan.CallOpts)
}

// Loans is a free data retrieval call binding the contract method 0xc4a90815.
//
// Solidity: function loans( bytes32) constant returns(state uint8, borrower address, digest bytes32, amount uint256, request uint256, principle uint256, interest uint256, maturityDate uint256, escrowDeadline uint256, stableCoinReceiver bytes)
func (_SimpleLoan *SimpleLoanCaller) Loans(opts *bind.CallOpts, arg0 [32]byte) (struct {
	State              uint8
	Borrower           common.Address
	Digest             [32]byte
	Amount             *big.Int
	Request            *big.Int
	Principle          *big.Int
	Interest           *big.Int
	MaturityDate       *big.Int
	EscrowDeadline     *big.Int
	StableCoinReceiver []byte
}, error) {
	ret := new(struct {
		State              uint8
		Borrower           common.Address
		Digest             [32]byte
		Amount             *big.Int
		Request            *big.Int
		Principle          *big.Int
		Interest           *big.Int
		MaturityDate       *big.Int
		EscrowDeadline     *big.Int
		StableCoinReceiver []byte
	})
	out := ret
	err := _SimpleLoan.contract.Call(opts, out, "loans", arg0)
	return *ret, err
}

// Loans is a free data retrieval call binding the contract method 0xc4a90815.
//
// Solidity: function loans( bytes32) constant returns(state uint8, borrower address, digest bytes32, amount uint256, request uint256, principle uint256, interest uint256, maturityDate uint256, escrowDeadline uint256, stableCoinReceiver bytes)
func (_SimpleLoan *SimpleLoanSession) Loans(arg0 [32]byte) (struct {
	State              uint8
	Borrower           common.Address
	Digest             [32]byte
	Amount             *big.Int
	Request            *big.Int
	Principle          *big.Int
	Interest           *big.Int
	MaturityDate       *big.Int
	EscrowDeadline     *big.Int
	StableCoinReceiver []byte
}, error) {
	return _SimpleLoan.Contract.Loans(&_SimpleLoan.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xc4a90815.
//
// Solidity: function loans( bytes32) constant returns(state uint8, borrower address, digest bytes32, amount uint256, request uint256, principle uint256, interest uint256, maturityDate uint256, escrowDeadline uint256, stableCoinReceiver bytes)
func (_SimpleLoan *SimpleLoanCallerSession) Loans(arg0 [32]byte) (struct {
	State              uint8
	Borrower           common.Address
	Digest             [32]byte
	Amount             *big.Int
	Request            *big.Int
	Principle          *big.Int
	Interest           *big.Int
	MaturityDate       *big.Int
	EscrowDeadline     *big.Int
	StableCoinReceiver []byte
}, error) {
	return _SimpleLoan.Contract.Loans(&_SimpleLoan.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xdc6ab527.
//
// Solidity: function params( bytes32) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCaller) Params(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "params", arg0)
	return *ret0, err
}

// Params is a free data retrieval call binding the contract method 0xdc6ab527.
//
// Solidity: function params( bytes32) constant returns(uint256)
func (_SimpleLoan *SimpleLoanSession) Params(arg0 [32]byte) (*big.Int, error) {
	return _SimpleLoan.Contract.Params(&_SimpleLoan.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xdc6ab527.
//
// Solidity: function params( bytes32) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCallerSession) Params(arg0 [32]byte) (*big.Int, error) {
	return _SimpleLoan.Contract.Params(&_SimpleLoan.CallOpts, arg0)
}

// Part is a free data retrieval call binding the contract method 0x45278644.
//
// Solidity: function part(value uint256, percent uint256) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCaller) Part(opts *bind.CallOpts, value *big.Int, percent *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "part", value, percent)
	return *ret0, err
}

// Part is a free data retrieval call binding the contract method 0x45278644.
//
// Solidity: function part(value uint256, percent uint256) constant returns(uint256)
func (_SimpleLoan *SimpleLoanSession) Part(value *big.Int, percent *big.Int) (*big.Int, error) {
	return _SimpleLoan.Contract.Part(&_SimpleLoan.CallOpts, value, percent)
}

// Part is a free data retrieval call binding the contract method 0x45278644.
//
// Solidity: function part(value uint256, percent uint256) constant returns(uint256)
func (_SimpleLoan *SimpleLoanCallerSession) Part(value *big.Int, percent *big.Int) (*big.Int, error) {
	return _SimpleLoan.Contract.Part(&_SimpleLoan.CallOpts, value, percent)
}

// SafelyCollateralized is a free data retrieval call binding the contract method 0x67ebd110.
//
// Solidity: function safelyCollateralized(collateralAmount uint256, debtAmount uint256) constant returns(bool)
func (_SimpleLoan *SimpleLoanCaller) SafelyCollateralized(opts *bind.CallOpts, collateralAmount *big.Int, debtAmount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SimpleLoan.contract.Call(opts, out, "safelyCollateralized", collateralAmount, debtAmount)
	return *ret0, err
}

// SafelyCollateralized is a free data retrieval call binding the contract method 0x67ebd110.
//
// Solidity: function safelyCollateralized(collateralAmount uint256, debtAmount uint256) constant returns(bool)
func (_SimpleLoan *SimpleLoanSession) SafelyCollateralized(collateralAmount *big.Int, debtAmount *big.Int) (bool, error) {
	return _SimpleLoan.Contract.SafelyCollateralized(&_SimpleLoan.CallOpts, collateralAmount, debtAmount)
}

// SafelyCollateralized is a free data retrieval call binding the contract method 0x67ebd110.
//
// Solidity: function safelyCollateralized(collateralAmount uint256, debtAmount uint256) constant returns(bool)
func (_SimpleLoan *SimpleLoanCallerSession) SafelyCollateralized(collateralAmount *big.Int, debtAmount *big.Int) (bool, error) {
	return _SimpleLoan.Contract.SafelyCollateralized(&_SimpleLoan.CallOpts, collateralAmount, debtAmount)
}

// AcceptLoan is a paid mutator transaction binding the contract method 0xe0d21437.
//
// Solidity: function acceptLoan(lid bytes32, key bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactor) AcceptLoan(opts *bind.TransactOpts, lid [32]byte, key [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "acceptLoan", lid, key, offchain)
}

// AcceptLoan is a paid mutator transaction binding the contract method 0xe0d21437.
//
// Solidity: function acceptLoan(lid bytes32, key bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanSession) AcceptLoan(lid [32]byte, key [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.AcceptLoan(&_SimpleLoan.TransactOpts, lid, key, offchain)
}

// AcceptLoan is a paid mutator transaction binding the contract method 0xe0d21437.
//
// Solidity: function acceptLoan(lid bytes32, key bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) AcceptLoan(lid [32]byte, key [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.AcceptLoan(&_SimpleLoan.TransactOpts, lid, key, offchain)
}

// AddPayment is a paid mutator transaction binding the contract method 0x5d43b0ae.
//
// Solidity: function addPayment(lid bytes32, amount uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactor) AddPayment(opts *bind.TransactOpts, lid [32]byte, amount *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "addPayment", lid, amount, offchain)
}

// AddPayment is a paid mutator transaction binding the contract method 0x5d43b0ae.
//
// Solidity: function addPayment(lid bytes32, amount uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanSession) AddPayment(lid [32]byte, amount *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.AddPayment(&_SimpleLoan.TransactOpts, lid, amount, offchain)
}

// AddPayment is a paid mutator transaction binding the contract method 0x5d43b0ae.
//
// Solidity: function addPayment(lid bytes32, amount uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) AddPayment(lid [32]byte, amount *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.AddPayment(&_SimpleLoan.TransactOpts, lid, amount, offchain)
}

// Liquidate is a paid mutator transaction binding the contract method 0x7805ffa5.
//
// Solidity: function liquidate(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactor) Liquidate(opts *bind.TransactOpts, lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "liquidate", lid, offchain)
}

// Liquidate is a paid mutator transaction binding the contract method 0x7805ffa5.
//
// Solidity: function liquidate(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanSession) Liquidate(lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.Liquidate(&_SimpleLoan.TransactOpts, lid, offchain)
}

// Liquidate is a paid mutator transaction binding the contract method 0x7805ffa5.
//
// Solidity: function liquidate(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) Liquidate(lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.Liquidate(&_SimpleLoan.TransactOpts, lid, offchain)
}

// RefundCollateral is a paid mutator transaction binding the contract method 0x5e8bc142.
//
// Solidity: function refundCollateral(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactor) RefundCollateral(opts *bind.TransactOpts, lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "refundCollateral", lid, offchain)
}

// RefundCollateral is a paid mutator transaction binding the contract method 0x5e8bc142.
//
// Solidity: function refundCollateral(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanSession) RefundCollateral(lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.RefundCollateral(&_SimpleLoan.TransactOpts, lid, offchain)
}

// RefundCollateral is a paid mutator transaction binding the contract method 0x5e8bc142.
//
// Solidity: function refundCollateral(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) RefundCollateral(lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.RefundCollateral(&_SimpleLoan.TransactOpts, lid, offchain)
}

// RejectLoan is a paid mutator transaction binding the contract method 0xd31eba5d.
//
// Solidity: function rejectLoan(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactor) RejectLoan(opts *bind.TransactOpts, lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "rejectLoan", lid, offchain)
}

// RejectLoan is a paid mutator transaction binding the contract method 0xd31eba5d.
//
// Solidity: function rejectLoan(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanSession) RejectLoan(lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.RejectLoan(&_SimpleLoan.TransactOpts, lid, offchain)
}

// RejectLoan is a paid mutator transaction binding the contract method 0xd31eba5d.
//
// Solidity: function rejectLoan(lid bytes32, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) RejectLoan(lid [32]byte, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.RejectLoan(&_SimpleLoan.TransactOpts, lid, offchain)
}

// SendCollateral is a paid mutator transaction binding the contract method 0x3adca01c.
//
// Solidity: function sendCollateral(lid bytes32, digest bytes32, stableCoinReceiver bytes, request uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactor) SendCollateral(opts *bind.TransactOpts, lid [32]byte, digest [32]byte, stableCoinReceiver []byte, request *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "sendCollateral", lid, digest, stableCoinReceiver, request, offchain)
}

// SendCollateral is a paid mutator transaction binding the contract method 0x3adca01c.
//
// Solidity: function sendCollateral(lid bytes32, digest bytes32, stableCoinReceiver bytes, request uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanSession) SendCollateral(lid [32]byte, digest [32]byte, stableCoinReceiver []byte, request *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.SendCollateral(&_SimpleLoan.TransactOpts, lid, digest, stableCoinReceiver, request, offchain)
}

// SendCollateral is a paid mutator transaction binding the contract method 0x3adca01c.
//
// Solidity: function sendCollateral(lid bytes32, digest bytes32, stableCoinReceiver bytes, request uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) SendCollateral(lid [32]byte, digest [32]byte, stableCoinReceiver []byte, request *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.SendCollateral(&_SimpleLoan.TransactOpts, lid, digest, stableCoinReceiver, request, offchain)
}

// Update is a paid mutator transaction binding the contract method 0xb3f52fec.
//
// Solidity: function update(name bytes32, value uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactor) Update(opts *bind.TransactOpts, name [32]byte, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "update", name, value, offchain)
}

// Update is a paid mutator transaction binding the contract method 0xb3f52fec.
//
// Solidity: function update(name bytes32, value uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanSession) Update(name [32]byte, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.Update(&_SimpleLoan.TransactOpts, name, value, offchain)
}

// Update is a paid mutator transaction binding the contract method 0xb3f52fec.
//
// Solidity: function update(name bytes32, value uint256, offchain bytes32) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) Update(name [32]byte, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _SimpleLoan.Contract.Update(&_SimpleLoan.TransactOpts, name, value, offchain)
}

// UpdateAssetPrice is a paid mutator transaction binding the contract method 0x8eccc8ce.
//
// Solidity: function updateAssetPrice(newPrice uint256) returns()
func (_SimpleLoan *SimpleLoanTransactor) UpdateAssetPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "updateAssetPrice", newPrice)
}

// UpdateAssetPrice is a paid mutator transaction binding the contract method 0x8eccc8ce.
//
// Solidity: function updateAssetPrice(newPrice uint256) returns()
func (_SimpleLoan *SimpleLoanSession) UpdateAssetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _SimpleLoan.Contract.UpdateAssetPrice(&_SimpleLoan.TransactOpts, newPrice)
}

// UpdateAssetPrice is a paid mutator transaction binding the contract method 0x8eccc8ce.
//
// Solidity: function updateAssetPrice(newPrice uint256) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) UpdateAssetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _SimpleLoan.Contract.UpdateAssetPrice(&_SimpleLoan.TransactOpts, newPrice)
}

// UpdateCollateralPrice is a paid mutator transaction binding the contract method 0x23592f86.
//
// Solidity: function updateCollateralPrice(newPrice uint256) returns()
func (_SimpleLoan *SimpleLoanTransactor) UpdateCollateralPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _SimpleLoan.contract.Transact(opts, "updateCollateralPrice", newPrice)
}

// UpdateCollateralPrice is a paid mutator transaction binding the contract method 0x23592f86.
//
// Solidity: function updateCollateralPrice(newPrice uint256) returns()
func (_SimpleLoan *SimpleLoanSession) UpdateCollateralPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _SimpleLoan.Contract.UpdateCollateralPrice(&_SimpleLoan.TransactOpts, newPrice)
}

// UpdateCollateralPrice is a paid mutator transaction binding the contract method 0x23592f86.
//
// Solidity: function updateCollateralPrice(newPrice uint256) returns()
func (_SimpleLoan *SimpleLoanTransactorSession) UpdateCollateralPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _SimpleLoan.Contract.UpdateCollateralPrice(&_SimpleLoan.TransactOpts, newPrice)
}

// SimpleLoanAcceptLoanIterator is returned from FilterAcceptLoan and is used to iterate over the raw logs and unpacked data for AcceptLoan events raised by the SimpleLoan contract.
type SimpleLoanAcceptLoanIterator struct {
	Event *SimpleLoanAcceptLoan // Event containing the contract specifics and raw log

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
func (it *SimpleLoanAcceptLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLoanAcceptLoan)
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
		it.Event = new(SimpleLoanAcceptLoan)
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
func (it *SimpleLoanAcceptLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLoanAcceptLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLoanAcceptLoan represents a AcceptLoan event raised by the SimpleLoan contract.
type SimpleLoanAcceptLoan struct {
	Lid      [32]byte
	Key      [32]byte
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptLoan is a free log retrieval operation binding the contract event 0x2caa7d07be7b7990462ca8a059568be0582ae0a101ccdcc1fdaba0da2cbf4beb.
//
// Solidity: e __acceptLoan(lid bytes32, key bytes32, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) FilterAcceptLoan(opts *bind.FilterOpts) (*SimpleLoanAcceptLoanIterator, error) {

	logs, sub, err := _SimpleLoan.contract.FilterLogs(opts, "__acceptLoan")
	if err != nil {
		return nil, err
	}
	return &SimpleLoanAcceptLoanIterator{contract: _SimpleLoan.contract, event: "__acceptLoan", logs: logs, sub: sub}, nil
}

// WatchAcceptLoan is a free log subscription operation binding the contract event 0x2caa7d07be7b7990462ca8a059568be0582ae0a101ccdcc1fdaba0da2cbf4beb.
//
// Solidity: e __acceptLoan(lid bytes32, key bytes32, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) WatchAcceptLoan(opts *bind.WatchOpts, sink chan<- *SimpleLoanAcceptLoan) (event.Subscription, error) {

	logs, sub, err := _SimpleLoan.contract.WatchLogs(opts, "__acceptLoan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLoanAcceptLoan)
				if err := _SimpleLoan.contract.UnpackLog(event, "__acceptLoan", log); err != nil {
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

// SimpleLoanAddPaymentIterator is returned from FilterAddPayment and is used to iterate over the raw logs and unpacked data for AddPayment events raised by the SimpleLoan contract.
type SimpleLoanAddPaymentIterator struct {
	Event *SimpleLoanAddPayment // Event containing the contract specifics and raw log

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
func (it *SimpleLoanAddPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLoanAddPayment)
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
		it.Event = new(SimpleLoanAddPayment)
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
func (it *SimpleLoanAddPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLoanAddPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLoanAddPayment represents a AddPayment event raised by the SimpleLoan contract.
type SimpleLoanAddPayment struct {
	Lid      [32]byte
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddPayment is a free log retrieval operation binding the contract event 0x6f17325e5a83dd2569c62cfbdba7cb9b19e73b7d46b7d37aa0acf1cc8ffb9d6a.
//
// Solidity: e __addPayment(lid bytes32, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) FilterAddPayment(opts *bind.FilterOpts) (*SimpleLoanAddPaymentIterator, error) {

	logs, sub, err := _SimpleLoan.contract.FilterLogs(opts, "__addPayment")
	if err != nil {
		return nil, err
	}
	return &SimpleLoanAddPaymentIterator{contract: _SimpleLoan.contract, event: "__addPayment", logs: logs, sub: sub}, nil
}

// WatchAddPayment is a free log subscription operation binding the contract event 0x6f17325e5a83dd2569c62cfbdba7cb9b19e73b7d46b7d37aa0acf1cc8ffb9d6a.
//
// Solidity: e __addPayment(lid bytes32, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) WatchAddPayment(opts *bind.WatchOpts, sink chan<- *SimpleLoanAddPayment) (event.Subscription, error) {

	logs, sub, err := _SimpleLoan.contract.WatchLogs(opts, "__addPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLoanAddPayment)
				if err := _SimpleLoan.contract.UnpackLog(event, "__addPayment", log); err != nil {
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

// SimpleLoanLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the SimpleLoan contract.
type SimpleLoanLiquidateIterator struct {
	Event *SimpleLoanLiquidate // Event containing the contract specifics and raw log

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
func (it *SimpleLoanLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLoanLiquidate)
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
		it.Event = new(SimpleLoanLiquidate)
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
func (it *SimpleLoanLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLoanLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLoanLiquidate represents a Liquidate event raised by the SimpleLoan contract.
type SimpleLoanLiquidate struct {
	Lid        [32]byte
	Amount     *big.Int
	Commission *big.Int
	Offchain   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x064e21237eec8c3f4bdf95fe7d379974f5ce7292b9454fa13301ed92a443c4ed.
//
// Solidity: e __liquidate(lid bytes32, amount uint256, commission uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) FilterLiquidate(opts *bind.FilterOpts) (*SimpleLoanLiquidateIterator, error) {

	logs, sub, err := _SimpleLoan.contract.FilterLogs(opts, "__liquidate")
	if err != nil {
		return nil, err
	}
	return &SimpleLoanLiquidateIterator{contract: _SimpleLoan.contract, event: "__liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x064e21237eec8c3f4bdf95fe7d379974f5ce7292b9454fa13301ed92a443c4ed.
//
// Solidity: e __liquidate(lid bytes32, amount uint256, commission uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *SimpleLoanLiquidate) (event.Subscription, error) {

	logs, sub, err := _SimpleLoan.contract.WatchLogs(opts, "__liquidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLoanLiquidate)
				if err := _SimpleLoan.contract.UnpackLog(event, "__liquidate", log); err != nil {
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

// SimpleLoanRefundCollateralIterator is returned from FilterRefundCollateral and is used to iterate over the raw logs and unpacked data for RefundCollateral events raised by the SimpleLoan contract.
type SimpleLoanRefundCollateralIterator struct {
	Event *SimpleLoanRefundCollateral // Event containing the contract specifics and raw log

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
func (it *SimpleLoanRefundCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLoanRefundCollateral)
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
		it.Event = new(SimpleLoanRefundCollateral)
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
func (it *SimpleLoanRefundCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLoanRefundCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLoanRefundCollateral represents a RefundCollateral event raised by the SimpleLoan contract.
type SimpleLoanRefundCollateral struct {
	Lid      [32]byte
	Amount   *big.Int
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundCollateral is a free log retrieval operation binding the contract event 0x260ef70d5caed6f134782cfa7f0db59b4f7ce3726cdf88e644aeaa898edbae54.
//
// Solidity: e __refundCollateral(lid bytes32, amount uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) FilterRefundCollateral(opts *bind.FilterOpts) (*SimpleLoanRefundCollateralIterator, error) {

	logs, sub, err := _SimpleLoan.contract.FilterLogs(opts, "__refundCollateral")
	if err != nil {
		return nil, err
	}
	return &SimpleLoanRefundCollateralIterator{contract: _SimpleLoan.contract, event: "__refundCollateral", logs: logs, sub: sub}, nil
}

// WatchRefundCollateral is a free log subscription operation binding the contract event 0x260ef70d5caed6f134782cfa7f0db59b4f7ce3726cdf88e644aeaa898edbae54.
//
// Solidity: e __refundCollateral(lid bytes32, amount uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) WatchRefundCollateral(opts *bind.WatchOpts, sink chan<- *SimpleLoanRefundCollateral) (event.Subscription, error) {

	logs, sub, err := _SimpleLoan.contract.WatchLogs(opts, "__refundCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLoanRefundCollateral)
				if err := _SimpleLoan.contract.UnpackLog(event, "__refundCollateral", log); err != nil {
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

// SimpleLoanRejectLoanIterator is returned from FilterRejectLoan and is used to iterate over the raw logs and unpacked data for RejectLoan events raised by the SimpleLoan contract.
type SimpleLoanRejectLoanIterator struct {
	Event *SimpleLoanRejectLoan // Event containing the contract specifics and raw log

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
func (it *SimpleLoanRejectLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLoanRejectLoan)
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
		it.Event = new(SimpleLoanRejectLoan)
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
func (it *SimpleLoanRejectLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLoanRejectLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLoanRejectLoan represents a RejectLoan event raised by the SimpleLoan contract.
type SimpleLoanRejectLoan struct {
	Lid      [32]byte
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRejectLoan is a free log retrieval operation binding the contract event 0xc74197ec873ffbac3377394f9f2ac39b6b743f8592529f7e76500661ee3bd534.
//
// Solidity: e __rejectLoan(lid bytes32, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) FilterRejectLoan(opts *bind.FilterOpts) (*SimpleLoanRejectLoanIterator, error) {

	logs, sub, err := _SimpleLoan.contract.FilterLogs(opts, "__rejectLoan")
	if err != nil {
		return nil, err
	}
	return &SimpleLoanRejectLoanIterator{contract: _SimpleLoan.contract, event: "__rejectLoan", logs: logs, sub: sub}, nil
}

// WatchRejectLoan is a free log subscription operation binding the contract event 0xc74197ec873ffbac3377394f9f2ac39b6b743f8592529f7e76500661ee3bd534.
//
// Solidity: e __rejectLoan(lid bytes32, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) WatchRejectLoan(opts *bind.WatchOpts, sink chan<- *SimpleLoanRejectLoan) (event.Subscription, error) {

	logs, sub, err := _SimpleLoan.contract.WatchLogs(opts, "__rejectLoan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLoanRejectLoan)
				if err := _SimpleLoan.contract.UnpackLog(event, "__rejectLoan", log); err != nil {
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

// SimpleLoanSendCollateralIterator is returned from FilterSendCollateral and is used to iterate over the raw logs and unpacked data for SendCollateral events raised by the SimpleLoan contract.
type SimpleLoanSendCollateralIterator struct {
	Event *SimpleLoanSendCollateral // Event containing the contract specifics and raw log

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
func (it *SimpleLoanSendCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLoanSendCollateral)
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
		it.Event = new(SimpleLoanSendCollateral)
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
func (it *SimpleLoanSendCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLoanSendCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLoanSendCollateral represents a SendCollateral event raised by the SimpleLoan contract.
type SimpleLoanSendCollateral struct {
	Lid      [32]byte
	Amount   *big.Int
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSendCollateral is a free log retrieval operation binding the contract event 0x474f79dc33c29e99abd873b5114b4d572f88d3a28fbc53cbc252cbb145cf2e3f.
//
// Solidity: e __sendCollateral(lid bytes32, amount uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) FilterSendCollateral(opts *bind.FilterOpts) (*SimpleLoanSendCollateralIterator, error) {

	logs, sub, err := _SimpleLoan.contract.FilterLogs(opts, "__sendCollateral")
	if err != nil {
		return nil, err
	}
	return &SimpleLoanSendCollateralIterator{contract: _SimpleLoan.contract, event: "__sendCollateral", logs: logs, sub: sub}, nil
}

// WatchSendCollateral is a free log subscription operation binding the contract event 0x474f79dc33c29e99abd873b5114b4d572f88d3a28fbc53cbc252cbb145cf2e3f.
//
// Solidity: e __sendCollateral(lid bytes32, amount uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) WatchSendCollateral(opts *bind.WatchOpts, sink chan<- *SimpleLoanSendCollateral) (event.Subscription, error) {

	logs, sub, err := _SimpleLoan.contract.WatchLogs(opts, "__sendCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLoanSendCollateral)
				if err := _SimpleLoan.contract.UnpackLog(event, "__sendCollateral", log); err != nil {
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

// SimpleLoanUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the SimpleLoan contract.
type SimpleLoanUpdateIterator struct {
	Event *SimpleLoanUpdate // Event containing the contract specifics and raw log

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
func (it *SimpleLoanUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLoanUpdate)
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
		it.Event = new(SimpleLoanUpdate)
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
func (it *SimpleLoanUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLoanUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLoanUpdate represents a Update event raised by the SimpleLoan contract.
type SimpleLoanUpdate struct {
	Name     [32]byte
	Value    *big.Int
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x31eb539fa5855ce13b272ee9f9520cb20a498c755ce42a6f984fce0e19bd38b0.
//
// Solidity: e __update(name bytes32, value uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) FilterUpdate(opts *bind.FilterOpts) (*SimpleLoanUpdateIterator, error) {

	logs, sub, err := _SimpleLoan.contract.FilterLogs(opts, "__update")
	if err != nil {
		return nil, err
	}
	return &SimpleLoanUpdateIterator{contract: _SimpleLoan.contract, event: "__update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x31eb539fa5855ce13b272ee9f9520cb20a498c755ce42a6f984fce0e19bd38b0.
//
// Solidity: e __update(name bytes32, value uint256, offchain bytes32)
func (_SimpleLoan *SimpleLoanFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *SimpleLoanUpdate) (event.Subscription, error) {

	logs, sub, err := _SimpleLoan.contract.WatchLogs(opts, "__update")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLoanUpdate)
				if err := _SimpleLoan.contract.UnpackLog(event, "__update", log); err != nil {
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
