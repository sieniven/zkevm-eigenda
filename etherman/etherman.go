package etherman

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/sieniven/polygoncdk-eigenda/etherman/smartcontracts/polygonrollupmanager"
	polygonzkevm "github.com/sieniven/polygoncdk-eigenda/etherman/smartcontracts/polygonvalidium_xlayer"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Minimal implementation of PolygonCDK's ether manager
type Client struct {
	EthClient     ethereumClient
	ZkEVM         *polygonzkevm.PolygonvalidiumXlayer
	RollupManager *polygonrollupmanager.Polygonrollupmanager
	SCAddresses   []common.Address
	RollupID      uint32
	l1Cfg         L1Config
	auth          map[common.Address]bind.TransactOpts // empty in case of read-only client
}

type ethereumClient interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.TransactionReader
	ethereum.TransactionSender
	bind.DeployBackend
}

// L1Config represents the configuration of the network used in L1
type L1Config struct {
	// Chain ID of the L1 network
	L1ChainID uint64 `json:"chainId"`
	// ZkEVMAddr Address of the L1 contract polygonZkEVMAddress
	ZkEVMAddr common.Address `json:"polygonZkEVMAddress"`
	// RollupManagerAddr Address of the L1 contract
	RollupManagerAddr common.Address `json:"polygonRollupManagerAddress"`
	// EigenDAVerifierManagerAddr Address of the L1 contract
	EigenDAVerifierManagerAddr common.Address `json:"eigenDAVerifierManagerAddress"`
}

func NewClient(url string, l1Config L1Config) (*Client, error) {
	// Connect to ethereum node
	ethClient, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("error connecting to %s: %+v\n", url, err)
		return nil, err
	}
	// Create smc clients
	zkevm, err := polygonzkevm.NewPolygonvalidiumXlayer(l1Config.ZkEVMAddr, ethClient)
	if err != nil {
		fmt.Printf("error creating Polygonzkevm client (%s)\n", l1Config.ZkEVMAddr.String())
		return nil, err
	}
	rollupManager, err := polygonrollupmanager.NewPolygonrollupmanager(l1Config.RollupManagerAddr, ethClient)
	if err != nil {
		fmt.Printf("error creating NewPolygonrollupmanager client (%s)\n", l1Config.RollupManagerAddr.String())
		return nil, err
	}
	var scAddresses []common.Address
	scAddresses = append(scAddresses, l1Config.ZkEVMAddr, l1Config.RollupManagerAddr, l1Config.EigenDAVerifierManagerAddr)

	// get RollupID
	rollupID, err := rollupManager.RollupAddressToID(&bind.CallOpts{Pending: false}, l1Config.ZkEVMAddr)
	if err != nil {
		fmt.Printf("error rollupManager.RollupAddressToID(%s)\n", l1Config.RollupManagerAddr)
	}

	return &Client{
		EthClient:     ethClient,
		ZkEVM:         zkevm,
		RollupManager: rollupManager,
		SCAddresses:   scAddresses,
		RollupID:      rollupID,
		l1Cfg:         l1Config,
		auth:          map[common.Address]bind.TransactOpts{},
	}, nil
}

// Mock function to replicate BuildSequenceBatchesTxData on PolygonCDK.
func (etherMan *Client) BuildMockSequenceBatchesTxData(sender common.Address, sequenceNum int, dataAvailabilityMessage []byte) (to *common.Address, data []byte, err error) {
	opts, err := etherMan.generateMockAuth(sender)
	if err != nil {
		return nil, nil, err
	}
	opts.NoSend = true
	// force nonce, gas limit and gas price to avoid querying it from the chain
	opts.Nonce = big.NewInt(1)
	opts.GasLimit = uint64(1)
	opts.GasPrice = big.NewInt(1)

	tx, err := etherMan.sequenceBatches(opts, sequenceNum, dataAvailabilityMessage)
	if err != nil {
		return nil, nil, err
	}
	return tx.To(), tx.Data(), nil
}

// Mock function to replicate sequenceBatches on PolygonCDK
// We will generate randomized []bytes to be sent to the mock PoE SC method SequenceBatchesValidium.
func (etherMan *Client) sequenceBatches(opts bind.TransactOpts, sequenceNum int, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	data := []byte("hihihihihihihihihihihihihihihihihihi")
	var defaultHash common.Hash
	var defaultAddr common.Address

	var batches []polygonzkevm.PolygonValidiumEtrogValidiumBatchData
	for idx := 0; idx < sequenceNum; idx++ {
		batch := polygonzkevm.PolygonValidiumEtrogValidiumBatchData{
			TransactionsHash:     crypto.Keccak256Hash(data),
			ForcedGlobalExitRoot: defaultHash,
			ForcedTimestamp:      uint64(0),
			ForcedBlockHashL1:    defaultHash,
		}
		batches = append(batches, batch)
	}
	tx, err := etherMan.ZkEVM.SequenceBatchesValidium(&opts, batches, uint64(0), uint64(0), defaultAddr, dataAvailabilityMessage)
	if err != nil {
		fmt.Println("sequenceBatches failed")
	}
	return tx, err
}

// LoadAuthFromKeyStoreXLayer loads an authorization from a key store file
func (etherMan *Client) LoadAuthFromKeyStore(path, password string) (*bind.TransactOpts, *ecdsa.PrivateKey, error) {
	auth, pk, err := newAuthFromKeystore(path, password, etherMan.l1Cfg.L1ChainID)
	if err != nil {
		return nil, nil, err
	}

	fmt.Printf("loaded authorization for address: %v\n", auth.From.String())
	etherMan.auth[auth.From] = auth
	return &auth, pk, nil
}

// newAuthFromKeystore an authorization instance from a keystore file
func newAuthFromKeystore(path, password string, chainID uint64) (bind.TransactOpts, *ecdsa.PrivateKey, error) {
	fmt.Printf("reading key from: %v\n", path)
	key, err := newKeyFromKeystore(path, password)
	if err != nil {
		return bind.TransactOpts{}, nil, err
	}
	if key == nil {
		return bind.TransactOpts{}, nil, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, new(big.Int).SetUint64(chainID))
	if err != nil {
		return bind.TransactOpts{}, nil, err
	}
	return *auth, key.PrivateKey, nil
}

// newKeyFromKeystore creates an instance of a keystore key from a keystore file
func newKeyFromKeystore(path, password string) (*keystore.Key, error) {
	if path == "" && password == "" {
		return nil, nil
	}
	keystoreEncrypted, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	fmt.Printf("decrypting key from: %v\n", path)
	key, err := keystore.DecryptKey(keystoreEncrypted, password)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// generateMockAuth generates an authorization instance from a randomly generated private key
// to be used to estimate gas for PoE operations NOT restricted to the Trusted Sequencer
func (etherMan *Client) generateMockAuth(sender common.Address) (bind.TransactOpts, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return bind.TransactOpts{}, fmt.Errorf("failed to generate a private key to estimate L1 txs")
	}
	chainID := big.NewInt(0).SetUint64(etherMan.l1Cfg.L1ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return bind.TransactOpts{}, fmt.Errorf("failed to generate a fake authorization to estimate L1 txs")
	}

	auth.From = sender
	auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		chainID := big.NewInt(0).SetUint64(etherMan.l1Cfg.L1ChainID)
		signer := types.LatestSignerForChainID(chainID)
		if err != nil {
			return nil, err
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}
	return *auth, nil
}
