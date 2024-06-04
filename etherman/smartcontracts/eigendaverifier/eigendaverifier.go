// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eigendaverifier

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// EigenDARollupUtilsBlobVerificationProof is an auto generated low-level Go binding around an user-defined struct.
type EigenDARollupUtilsBlobVerificationProof struct {
	BatchId        uint32
	BlobIndex      uint32
	BatchMetadata  IEigenDAServiceManagerBatchMetadata
	InclusionProof []byte
	QuorumIndices  []byte
}

// EigenDAVerifierBlobData is an auto generated low-level Go binding around an user-defined struct.
type EigenDAVerifierBlobData struct {
	BlobHeader            IEigenDAServiceManagerBlobHeader
	BlobVerificationProof EigenDARollupUtilsBlobVerificationProof
}

// IEigenDAServiceManagerBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBatchHeader struct {
	BlobHeadersRoot       [32]byte
	QuorumNumbers         []byte
	SignedStakeForQuorums []byte
	ReferenceBlockNumber  uint32
}

// IEigenDAServiceManagerBatchMetadata is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBatchMetadata struct {
	BatchHeader             IEigenDAServiceManagerBatchHeader
	SignatoryRecordHash     [32]byte
	ConfirmationBlockNumber uint32
}

// IEigenDAServiceManagerBlobHeader is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBlobHeader struct {
	Commitment       BN254G1Point
	DataLength       uint32
	QuorumBlobParams []IEigenDAServiceManagerQuorumBlobParam
}

// IEigenDAServiceManagerQuorumBlobParam is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerQuorumBlobParam struct {
	QuorumNumber                    uint8
	AdversaryThresholdPercentage    uint8
	ConfirmationThresholdPercentage uint8
	ChunkLength                     uint32
}

// EigendaverifierMetaData contains all meta data concerning the Eigendaverifier contract.
var EigendaverifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eigenDAServiceManager\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptAdminRole\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decodeBlobData\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"blobData\",\"type\":\"tuple\",\"internalType\":\"structEigenDAVerifier.BlobData\",\"components\":[{\"name\":\"blobHeader\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BlobHeader\",\"components\":[{\"name\":\"commitment\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"dataLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBlobParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenDAServiceManager.QuorumBlobParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"adversaryThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"confirmationThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chunkLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]},{\"name\":\"blobVerificationProof\",\"type\":\"tuple\",\"internalType\":\"structEigenDARollupUtils.BlobVerificationProof\",\"components\":[{\"name\":\"batchId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"batchMetadata\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchMetadata\",\"components\":[{\"name\":\"batchHeader\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchHeader\",\"components\":[{\"name\":\"blobHeadersRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signedStakeForQuorums\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"confirmationBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"inclusionProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumIndices\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getProcotolName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pendingAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDataAvailabilityProtocol\",\"inputs\":[{\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferAdminRole\",\"inputs\":[{\"name\":\"newPendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AcceptAdminRole\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDataAvailabilityProtocol\",\"inputs\":[{\"name\":\"newTrustedSequencer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIEigenDAServiceManager\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferAdminRole\",\"inputs\":[{\"name\":\"newPendingAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BatchAlreadyVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedMaxVerifyBatches\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FinalPendingStateNumInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchTimeoutNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesAlreadyActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesDecentralized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForceBatchesOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForcedDataDoesNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalExitRootNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HaltTimeoutNotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HugeTokenMetadataNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InitSequencedBatchDoesNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializeTransaction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRangeBatchTimeTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRangeForceBatchTimeout\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRangeMultiplierBatchFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxTimestampSequenceInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewAccInputHashDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewPendingStateTimeoutMustBeLower\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewStateRootNotInsidePrime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughMaticAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughPOLAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OldAccInputHashDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OldStateRootDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPendingAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRollupManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTrustedAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTrustedSequencer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateNotConsolidable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequenceZeroBatches\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequencedTimestampInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransactionsLengthAboveMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610dc0380380610dc083398101604081905261002f91610078565b600180546001600160a01b039384166001600160a01b031991821617909155600080549290931691161790556100b2565b6001600160a01b038116811461007557600080fd5b50565b6000806040838503121561008b57600080fd5b825161009681610060565b60208401519092506100a781610060565b809150509250929050565b610cff806100c16000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063aba4c80d1161005b578063aba4c80d146100ed578063ada8f9191461010d578063e4f1712014610120578063f851a4401461014957600080fd5b8063267822471461008d5780633b51be4b146100bd5780637cd76b8b146100d25780638c3d7301146100e5575b600080fd5b6002546100a0906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100d06100cb366004610468565b61015c565b005b6100d06100e03660046104cc565b6101ee565b6100d061026e565b6101006100fb3660046104e9565b6102ef565b6040516100b49190610644565b6100d061011b3660046104cc565b61030a565b6040805180820182526007815266456967656e444160c81b602082015290516100b49190610721565b6001546100a0906001600160a01b031681565b600061016883836102ef565b8051600054602083015160405163219460e160e21b815293945073__$399f9ce8dd33a06d144ce1eb24d845e280$__936386518384936101b89390926001600160a01b0390911691600401610734565b60006040518083038186803b1580156101d057600080fd5b505af41580156101e4573d6000803e3d6000fd5b5050505050505050565b6001546001600160a01b0316331461021957604051634755657960e01b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a906020015b60405180910390a150565b6002546001600160a01b031633146102995760405163d1ec4b2360e01b815260040160405180910390fd5b600254600180546001600160a01b0319166001600160a01b0390921691821790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b6102f7610383565b61030382840184610b0b565b9392505050565b6001546001600160a01b0316331461033557604051634755657960e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610263565b6040805160e081018252600060a0820181815260c0830182905292820192835260608083019190915260808201529081526020810161041a6040805160a0808201835260008083526020808401829052845160e0810186526060808201848152608083018290529482015260c081018390529283528201819052818401529091820190815260200160608152602001606081525090565b905290565b60008083601f84011261043157600080fd5b50813567ffffffffffffffff81111561044957600080fd5b60208301915083602082850101111561046157600080fd5b9250929050565b60008060006040848603121561047d57600080fd5b83359250602084013567ffffffffffffffff81111561049b57600080fd5b6104a78682870161041f565b9497909650939450505050565b6001600160a01b03811681146104c957600080fd5b50565b6000602082840312156104de57600080fd5b8135610303816104b4565b600080602083850312156104fc57600080fd5b823567ffffffffffffffff81111561051357600080fd5b61051f8582860161041f565b90969095509350505050565b6000815180845260005b8181101561055157602081850181015186830182015201610535565b506000602082860101526020601f19601f83011685010191505092915050565b600063ffffffff808351168452806020840151166020850152604083015160a060408601528051606060a08701528051610100870152602081015160806101208801526105c261018088018261052b565b9050604082015160ff19888303016101408901526105e0828261052b565b91505083606083015116610160880152602083015160c08801528360408401511660e088015260608601519350868103606088015261061f818561052b565b93505050506080830151848203608086015261063b828261052b565b95945050505050565b60006020808352835160408285015260e0840161066f60608601835180518252602090810151910152565b8183015163ffffffff1660a0860152604090910151608060c0860181905281519283905290830191600091906101008701905b808410156106f9576106e582865160ff815116825260ff602082015116602083015260ff604082015116604083015263ffffffff60608201511660608301525050565b9385019360019390930192908201906106a2565b5093870151868503601f19016040880152936107158186610571565b98975050505050505050565b602081526000610303602083018461052b565b60608152600060e0820161075660608401875180518252602090810151910152565b60208681015163ffffffff1660a08501526040870151608060c0860181905281519384905290820192600091906101008701905b808410156107e1576107cd82875160ff815116825260ff602082015116602083015260ff604082015116604083015263ffffffff60608201511660608301525050565b94840194600193909301929082019061078a565b506001600160a01b0389168785015286810360408801526108028189610571565b9a9950505050505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561084957610849610810565b60405290565b6040516080810167ffffffffffffffff8111828210171561084957610849610810565b60405160a0810167ffffffffffffffff8111828210171561084957610849610810565b6040805190810167ffffffffffffffff8111828210171561084957610849610810565b604051601f8201601f1916810167ffffffffffffffff811182821017156108e1576108e1610810565b604052919050565b803563ffffffff811681146108fd57600080fd5b919050565b803560ff811681146108fd57600080fd5b600082601f83011261092457600080fd5b813567ffffffffffffffff81111561093e5761093e610810565b610951601f8201601f19166020016108b8565b81815284602083860101111561096657600080fd5b816020850160208301376000918101602001919091529392505050565b60006060828403121561099557600080fd5b61099d610826565b9050813567ffffffffffffffff808211156109b757600080fd5b90830190608082860312156109cb57600080fd5b6109d361084f565b823581526020830135828111156109e957600080fd5b6109f587828601610913565b602083015250604083013582811115610a0d57600080fd5b610a1987828601610913565b604083015250610a2b606084016108e9565b60608201528352505060208281013590820152610a4a604083016108e9565b604082015292915050565b600060a08284031215610a6757600080fd5b610a6f610872565b9050610a7a826108e9565b8152610a88602083016108e9565b6020820152604082013567ffffffffffffffff80821115610aa857600080fd5b610ab485838601610983565b60408401526060840135915080821115610acd57600080fd5b610ad985838601610913565b60608401526080840135915080821115610af257600080fd5b50610aff84828501610913565b60808301525092915050565b60006020808385031215610b1e57600080fd5b823567ffffffffffffffff80821115610b3657600080fd5b81850191506040808388031215610b4c57600080fd5b610b54610895565b833583811115610b6357600080fd5b84018089036080811215610b7657600080fd5b610b7e610826565b84821215610b8b57600080fd5b610b93610895565b9150823582528783013588830152818152610baf8584016108e9565b88820152606091508183013586811115610bc857600080fd5b8084019350508a601f840112610bdd57600080fd5b823586811115610bef57610bef610810565b610bfd898260051b016108b8565b81815260079190911b8401890190898101908d831115610c1c57600080fd5b948a01945b82861015610c8c576080868f031215610c3a5760008081fd5b610c4261084f565b610c4b87610902565b8152610c588c8801610902565b8c820152610c67898801610902565b89820152610c768688016108e9565b81870152825260809590950194908a0190610c21565b96830196909652508352505083850135915082821115610cab57600080fd5b610cb788838601610a55565b8582015280955050505050509291505056fea264697066735822122079ef10de9995af0b23bdbc891f35a77ad90c814e9fe5b97b49bbdbb5f500cf0b64736f6c63430008140033",
}

// EigendaverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use EigendaverifierMetaData.ABI instead.
var EigendaverifierABI = EigendaverifierMetaData.ABI

// EigendaverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EigendaverifierMetaData.Bin instead.
var EigendaverifierBin = EigendaverifierMetaData.Bin

// DeployEigendaverifier deploys a new Ethereum contract, binding an instance of Eigendaverifier to it.
func DeployEigendaverifier(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _eigenDAServiceManager common.Address) (common.Address, *types.Transaction, *Eigendaverifier, error) {
	parsed, err := EigendaverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EigendaverifierBin), backend, _admin, _eigenDAServiceManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eigendaverifier{EigendaverifierCaller: EigendaverifierCaller{contract: contract}, EigendaverifierTransactor: EigendaverifierTransactor{contract: contract}, EigendaverifierFilterer: EigendaverifierFilterer{contract: contract}}, nil
}

// Eigendaverifier is an auto generated Go binding around an Ethereum contract.
type Eigendaverifier struct {
	EigendaverifierCaller     // Read-only binding to the contract
	EigendaverifierTransactor // Write-only binding to the contract
	EigendaverifierFilterer   // Log filterer for contract events
}

// EigendaverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigendaverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigendaverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigendaverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigendaverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigendaverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigendaverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigendaverifierSession struct {
	Contract     *Eigendaverifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EigendaverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigendaverifierCallerSession struct {
	Contract *EigendaverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EigendaverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigendaverifierTransactorSession struct {
	Contract     *EigendaverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EigendaverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigendaverifierRaw struct {
	Contract *Eigendaverifier // Generic contract binding to access the raw methods on
}

// EigendaverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigendaverifierCallerRaw struct {
	Contract *EigendaverifierCaller // Generic read-only contract binding to access the raw methods on
}

// EigendaverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigendaverifierTransactorRaw struct {
	Contract *EigendaverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigendaverifier creates a new instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifier(address common.Address, backend bind.ContractBackend) (*Eigendaverifier, error) {
	contract, err := bindEigendaverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eigendaverifier{EigendaverifierCaller: EigendaverifierCaller{contract: contract}, EigendaverifierTransactor: EigendaverifierTransactor{contract: contract}, EigendaverifierFilterer: EigendaverifierFilterer{contract: contract}}, nil
}

// NewEigendaverifierCaller creates a new read-only instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifierCaller(address common.Address, caller bind.ContractCaller) (*EigendaverifierCaller, error) {
	contract, err := bindEigendaverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigendaverifierCaller{contract: contract}, nil
}

// NewEigendaverifierTransactor creates a new write-only instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*EigendaverifierTransactor, error) {
	contract, err := bindEigendaverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigendaverifierTransactor{contract: contract}, nil
}

// NewEigendaverifierFilterer creates a new log filterer instance of Eigendaverifier, bound to a specific deployed contract.
func NewEigendaverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*EigendaverifierFilterer, error) {
	contract, err := bindEigendaverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigendaverifierFilterer{contract: contract}, nil
}

// bindEigendaverifier binds a generic wrapper to an already deployed contract.
func bindEigendaverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigendaverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eigendaverifier *EigendaverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eigendaverifier.Contract.EigendaverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eigendaverifier *EigendaverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.EigendaverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eigendaverifier *EigendaverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.EigendaverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eigendaverifier *EigendaverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eigendaverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eigendaverifier *EigendaverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eigendaverifier *EigendaverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Eigendaverifier *EigendaverifierCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Eigendaverifier *EigendaverifierSession) Admin() (common.Address, error) {
	return _Eigendaverifier.Contract.Admin(&_Eigendaverifier.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Eigendaverifier *EigendaverifierCallerSession) Admin() (common.Address, error) {
	return _Eigendaverifier.Contract.Admin(&_Eigendaverifier.CallOpts)
}

// DecodeBlobData is a free data retrieval call binding the contract method 0xaba4c80d.
//
// Solidity: function decodeBlobData(bytes data) pure returns((((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[]),(uint32,uint32,((bytes32,bytes,bytes,uint32),bytes32,uint32),bytes,bytes)) blobData)
func (_Eigendaverifier *EigendaverifierCaller) DecodeBlobData(opts *bind.CallOpts, data []byte) (EigenDAVerifierBlobData, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "decodeBlobData", data)

	if err != nil {
		return *new(EigenDAVerifierBlobData), err
	}

	out0 := *abi.ConvertType(out[0], new(EigenDAVerifierBlobData)).(*EigenDAVerifierBlobData)

	return out0, err

}

// DecodeBlobData is a free data retrieval call binding the contract method 0xaba4c80d.
//
// Solidity: function decodeBlobData(bytes data) pure returns((((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[]),(uint32,uint32,((bytes32,bytes,bytes,uint32),bytes32,uint32),bytes,bytes)) blobData)
func (_Eigendaverifier *EigendaverifierSession) DecodeBlobData(data []byte) (EigenDAVerifierBlobData, error) {
	return _Eigendaverifier.Contract.DecodeBlobData(&_Eigendaverifier.CallOpts, data)
}

// DecodeBlobData is a free data retrieval call binding the contract method 0xaba4c80d.
//
// Solidity: function decodeBlobData(bytes data) pure returns((((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[]),(uint32,uint32,((bytes32,bytes,bytes,uint32),bytes32,uint32),bytes,bytes)) blobData)
func (_Eigendaverifier *EigendaverifierCallerSession) DecodeBlobData(data []byte) (EigenDAVerifierBlobData, error) {
	return _Eigendaverifier.Contract.DecodeBlobData(&_Eigendaverifier.CallOpts, data)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Eigendaverifier *EigendaverifierCaller) GetProcotolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "getProcotolName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Eigendaverifier *EigendaverifierSession) GetProcotolName() (string, error) {
	return _Eigendaverifier.Contract.GetProcotolName(&_Eigendaverifier.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() pure returns(string)
func (_Eigendaverifier *EigendaverifierCallerSession) GetProcotolName() (string, error) {
	return _Eigendaverifier.Contract.GetProcotolName(&_Eigendaverifier.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Eigendaverifier *EigendaverifierCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Eigendaverifier *EigendaverifierSession) PendingAdmin() (common.Address, error) {
	return _Eigendaverifier.Contract.PendingAdmin(&_Eigendaverifier.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Eigendaverifier *EigendaverifierCallerSession) PendingAdmin() (common.Address, error) {
	return _Eigendaverifier.Contract.PendingAdmin(&_Eigendaverifier.CallOpts)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes data) view returns()
func (_Eigendaverifier *EigendaverifierCaller) VerifyMessage(opts *bind.CallOpts, arg0 [32]byte, data []byte) error {
	var out []interface{}
	err := _Eigendaverifier.contract.Call(opts, &out, "verifyMessage", arg0, data)

	if err != nil {
		return err
	}

	return err

}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes data) view returns()
func (_Eigendaverifier *EigendaverifierSession) VerifyMessage(arg0 [32]byte, data []byte) error {
	return _Eigendaverifier.Contract.VerifyMessage(&_Eigendaverifier.CallOpts, arg0, data)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 , bytes data) view returns()
func (_Eigendaverifier *EigendaverifierCallerSession) VerifyMessage(arg0 [32]byte, data []byte) error {
	return _Eigendaverifier.Contract.VerifyMessage(&_Eigendaverifier.CallOpts, arg0, data)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Eigendaverifier *EigendaverifierTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eigendaverifier.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Eigendaverifier *EigendaverifierSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Eigendaverifier.Contract.AcceptAdminRole(&_Eigendaverifier.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Eigendaverifier *EigendaverifierTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Eigendaverifier.Contract.AcceptAdminRole(&_Eigendaverifier.TransactOpts)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Eigendaverifier *EigendaverifierTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Eigendaverifier *EigendaverifierSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.SetDataAvailabilityProtocol(&_Eigendaverifier.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Eigendaverifier *EigendaverifierTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.SetDataAvailabilityProtocol(&_Eigendaverifier.TransactOpts, newDataAvailabilityProtocol)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Eigendaverifier *EigendaverifierTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Eigendaverifier *EigendaverifierSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.TransferAdminRole(&_Eigendaverifier.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Eigendaverifier *EigendaverifierTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Eigendaverifier.Contract.TransferAdminRole(&_Eigendaverifier.TransactOpts, newPendingAdmin)
}

// EigendaverifierAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Eigendaverifier contract.
type EigendaverifierAcceptAdminRoleIterator struct {
	Event *EigendaverifierAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *EigendaverifierAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigendaverifierAcceptAdminRole)
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
		it.Event = new(EigendaverifierAcceptAdminRole)
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
func (it *EigendaverifierAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigendaverifierAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigendaverifierAcceptAdminRole represents a AcceptAdminRole event raised by the Eigendaverifier contract.
type EigendaverifierAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*EigendaverifierAcceptAdminRoleIterator, error) {

	logs, sub, err := _Eigendaverifier.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &EigendaverifierAcceptAdminRoleIterator{contract: _Eigendaverifier.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *EigendaverifierAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Eigendaverifier.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigendaverifierAcceptAdminRole)
				if err := _Eigendaverifier.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) ParseAcceptAdminRole(log types.Log) (*EigendaverifierAcceptAdminRole, error) {
	event := new(EigendaverifierAcceptAdminRole)
	if err := _Eigendaverifier.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigendaverifierSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Eigendaverifier contract.
type EigendaverifierSetDataAvailabilityProtocolIterator struct {
	Event *EigendaverifierSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

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
func (it *EigendaverifierSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigendaverifierSetDataAvailabilityProtocol)
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
		it.Event = new(EigendaverifierSetDataAvailabilityProtocol)
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
func (it *EigendaverifierSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigendaverifierSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigendaverifierSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Eigendaverifier contract.
type EigendaverifierSetDataAvailabilityProtocol struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newTrustedSequencer)
func (_Eigendaverifier *EigendaverifierFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*EigendaverifierSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Eigendaverifier.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &EigendaverifierSetDataAvailabilityProtocolIterator{contract: _Eigendaverifier.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newTrustedSequencer)
func (_Eigendaverifier *EigendaverifierFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *EigendaverifierSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Eigendaverifier.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigendaverifierSetDataAvailabilityProtocol)
				if err := _Eigendaverifier.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
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

// ParseSetDataAvailabilityProtocol is a log parse operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newTrustedSequencer)
func (_Eigendaverifier *EigendaverifierFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*EigendaverifierSetDataAvailabilityProtocol, error) {
	event := new(EigendaverifierSetDataAvailabilityProtocol)
	if err := _Eigendaverifier.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigendaverifierTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Eigendaverifier contract.
type EigendaverifierTransferAdminRoleIterator struct {
	Event *EigendaverifierTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *EigendaverifierTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigendaverifierTransferAdminRole)
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
		it.Event = new(EigendaverifierTransferAdminRole)
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
func (it *EigendaverifierTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigendaverifierTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigendaverifierTransferAdminRole represents a TransferAdminRole event raised by the Eigendaverifier contract.
type EigendaverifierTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*EigendaverifierTransferAdminRoleIterator, error) {

	logs, sub, err := _Eigendaverifier.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &EigendaverifierTransferAdminRoleIterator{contract: _Eigendaverifier.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *EigendaverifierTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Eigendaverifier.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigendaverifierTransferAdminRole)
				if err := _Eigendaverifier.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Eigendaverifier *EigendaverifierFilterer) ParseTransferAdminRole(log types.Log) (*EigendaverifierTransferAdminRole, error) {
	event := new(EigendaverifierTransferAdminRole)
	if err := _Eigendaverifier.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
