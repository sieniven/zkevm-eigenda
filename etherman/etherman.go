package etherman

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/sieniven/polygoncdk-eigenda/etherman/smartcontracts/polygonrollupmanager"
	polygonzkevm "github.com/sieniven/polygoncdk-eigenda/etherman/smartcontracts/polygonvalidium_xlayer"
	ethmanTypes "github.com/sieniven/polygoncdk-eigenda/etherman/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type externalGasProviders struct {
	MultiGasProvider bool
	Providers        []ethereum.GasPricer
}

// Minimal implementation of PolygonCDK's ether manager
type Client struct {
	EthClient     ethereumClient
	ZkEVM         *polygonzkevm.PolygonvalidiumXlayer
	RollupManager *polygonrollupmanager.Polygonrollupmanager
	SCAddresses   []common.Address
	RollupID      uint32
	GasProviders  externalGasProviders
	l1Cfg         L1Config
	cfg           Config
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

var ErrNotFound = errors.New("not found")

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

func NewClient(cfg Config, l1Config L1Config) (*Client, error) {
	// Connect to ethereum node
	ethClient, err := ethclient.Dial(cfg.URL)
	if err != nil {
		fmt.Printf("error connecting to %s: %+v\n", cfg.URL, err)
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

	gProviders := []ethereum.GasPricer{ethClient}

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
		GasProviders: externalGasProviders{
			MultiGasProvider: false,
			Providers:        gProviders,
		},
		l1Cfg: l1Config,
		cfg:   cfg,
		auth:  map[common.Address]bind.TransactOpts{},
	}, nil
}

// EstimateGasSequenceBatchesXLayer estimates gas for sending batches
func (etherMan *Client) EstimateGasSequenceBatches(sender common.Address, sequences []ethmanTypes.Sequence, maxSequenceTimestamp uint64, lastSequencedBatchNumber uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	opts, err := etherMan.generateMockAuth(sender)
	if err == ErrNotFound {
		return nil, errors.New("can't find sender private key to sign tx")
	}
	opts.NoSend = true

	tx, err := etherMan.sequenceBatches(opts, sequences, maxSequenceTimestamp, lastSequencedBatchNumber, l2Coinbase, dataAvailabilityMessage)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Mock function to replicate BuildSequenceBatchesTxData on PolygonCDK.
func (etherMan *Client) BuildMockSequenceBatchesTxData(sender common.Address, sequences []ethmanTypes.Sequence, maxSequenceTimestamp uint64, lastSequencedBatchNumber uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (to *common.Address, data []byte, err error) {
	opts, err := etherMan.generateMockAuth(sender)
	if err != nil {
		return nil, nil, err
	}
	opts.NoSend = true
	// force nonce, gas limit and gas price to avoid querying it from the chain
	opts.Nonce = big.NewInt(1)
	opts.GasLimit = uint64(1)
	opts.GasPrice = big.NewInt(1)

	tx, err := etherMan.sequenceBatches(opts, sequences, maxSequenceTimestamp, lastSequencedBatchNumber, l2Coinbase, dataAvailabilityMessage)
	if err != nil {
		return nil, nil, err
	}
	return tx.To(), tx.Data(), nil
}

// Mock function to replicate sequenceBatches on PolygonCDK
// We will generate randomized []bytes to be sent to the mock PoE SC method SequenceBatchesValidium.
func (etherMan *Client) sequenceBatches(opts bind.TransactOpts, sequences []ethmanTypes.Sequence, maxSequenceTimestamp uint64, lastSequencedBatchNumber uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	var batches []polygonzkevm.PolygonValidiumEtrogValidiumBatchData
	for _, seq := range sequences {
		var ger common.Hash
		if seq.ForcedBatchTimestamp > 0 {
			ger = seq.GlobalExitRoot
		}
		batch := polygonzkevm.PolygonValidiumEtrogValidiumBatchData{
			TransactionsHash:     crypto.Keccak256Hash(seq.BatchL2Data),
			ForcedGlobalExitRoot: ger,
			ForcedTimestamp:      uint64(seq.ForcedBatchTimestamp),
			ForcedBlockHashL1:    seq.PrevBlockHash,
		}

		batches = append(batches, batch)
	}
	tx, err := etherMan.ZkEVM.SequenceBatchesValidium(&opts, batches, maxSequenceTimestamp, lastSequencedBatchNumber, l2Coinbase, dataAvailabilityMessage)
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

// getAuthByAddress tries to get an authorization from the authorizations map
func (etherMan *Client) getAuthByAddress(addr common.Address) (bind.TransactOpts, error) {
	auth, found := etherMan.auth[addr]
	if !found {
		return bind.TransactOpts{}, ErrNotFound
	}
	return auth, nil
}

// generateRandomAuth generates an authorization instance from a
// randomly generated private key to be used to estimate gas for PoE
// operations NOT restricted to the Trusted Sequencer
func (etherMan *Client) generateRandomAuth() (bind.TransactOpts, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return bind.TransactOpts{}, errors.New("failed to generate a private key to estimate L1 txs")
	}
	chainID := big.NewInt(0).SetUint64(etherMan.l1Cfg.L1ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return bind.TransactOpts{}, errors.New("failed to generate a fake authorization to estimate L1 txs")
	}

	return *auth, nil
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

// GetTx function get ethereum tx
func (etherMan *Client) GetTx(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	return etherMan.EthClient.TransactionByHash(ctx, txHash)
}

// GetTxReceipt function gets ethereum tx receipt
func (etherMan *Client) GetTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return etherMan.EthClient.TransactionReceipt(ctx, txHash)
}

// GetL1GasPrice gets the l1 gas price
func (etherMan *Client) GetL1GasPrice(ctx context.Context) *big.Int {
	// Get gasPrice from providers
	gasPrice := big.NewInt(0)
	for i, prov := range etherMan.GasProviders.Providers {
		gp, err := prov.SuggestGasPrice(ctx)
		if err != nil {
			fmt.Printf("error getting gas price from provider %d. Error: %s\n", i+1, err.Error())
		} else if gasPrice.Cmp(gp) == -1 { // gasPrice < gp
			gasPrice = gp
		}
	}
	fmt.Println("gasPrice chose: ", gasPrice)
	return gasPrice
}

// SendTx sends a tx to L1
func (etherMan *Client) SendTx(ctx context.Context, tx *types.Transaction) error {
	return etherMan.EthClient.SendTransaction(ctx, tx)
}

// SignTx tries to sign a transaction accordingly to the provided sender
func (etherMan *Client) SignTx(ctx context.Context, sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
	auth, err := etherMan.getAuthByAddress(sender)
	if err == ErrNotFound {
		return nil, ErrNotFound
	}
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

// CurrentNonce returns the current nonce for the provided account
func (etherMan *Client) CurrentNonce(ctx context.Context, account common.Address) (uint64, error) {
	return etherMan.EthClient.NonceAt(ctx, account, nil)
}

// SuggestedGasPrice returns the suggest nonce for the network at the moment
func (etherMan *Client) SuggestedGasPrice(ctx context.Context) (*big.Int, error) {
	suggestedGasPrice := etherMan.GetL1GasPrice(ctx)
	if suggestedGasPrice.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("failed to get the suggested gas price")
	}
	return suggestedGasPrice, nil
}

// EstimateGas returns the estimated gas for the tx
func (etherMan *Client) EstimateGas(ctx context.Context, from common.Address, to *common.Address, value *big.Int, data []byte) (uint64, error) {
	return etherMan.EthClient.EstimateGas(ctx, ethereum.CallMsg{
		From:  from,
		To:    to,
		Value: value,
		Data:  data,
	})
}

// CheckTxWasMined check if a tx was already mined
func (etherMan *Client) CheckTxWasMined(ctx context.Context, txHash common.Hash) (bool, *types.Receipt, error) {
	receipt, err := etherMan.EthClient.TransactionReceipt(ctx, txHash)
	if errors.Is(err, ethereum.NotFound) {
		return false, nil, nil
	} else if err != nil {
		return false, nil, err
	}
	return true, receipt, nil
}

// SetDataAvailabilityProtocol sets the address for the new data availability protocol
func (etherMan *Client) SetDataAvailabilityProtocol(from, daAddress common.Address) (*types.Transaction, error) {
	auth, err := etherMan.getAuthByAddress(from)
	if err != nil {
		return nil, err
	}
	return etherMan.ZkEVM.SetDataAvailabilityProtocol(&auth, daAddress)
}

// WaitTxToBeMined waits until a tx has been mined or the given timeout expires
func (etherMan *Client) WaitTxToBeMined(parentCtx context.Context, tx *types.Transaction, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, etherMan.EthClient, tx)
	if errors.Is(err, context.DeadlineExceeded) {
		return false, nil
	} else if err != nil {
		fmt.Printf("error waiting tx %s to be mined: %v\n", tx.Hash(), err)
		return false, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		// Get revert reason
		reason, reasonErr := RevertReason(ctx, etherMan.EthClient, tx, receipt.BlockNumber)
		if reasonErr != nil {
			reason = reasonErr.Error()
		}
		return false, fmt.Errorf("transaction has failed, reason: %s, receipt: %+v. tx: %+v, gas: %v", reason, receipt, tx, tx.Gas())
	}
	fmt.Printf("Transaction successfully mined: %v\n", tx.Hash())
	return true, nil
}

// GetRevertMessage tries to get a revert message of a transaction
func (etherMan *Client) GetRevertMessage(ctx context.Context, tx *types.Transaction) (string, error) {
	if tx == nil {
		return "", nil
	}

	receipt, err := etherMan.GetTxReceipt(ctx, tx.Hash())
	if err != nil {
		return "", err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		revertMessage, err := RevertReason(ctx, etherMan.EthClient, tx, receipt.BlockNumber)
		if err != nil {
			return "", err
		}
		return revertMessage, nil
	}
	return "", nil
}

// RevertReason returns the revert reason for a tx that has a receipt with failed status
func RevertReason(ctx context.Context, c ethereumClient, tx *types.Transaction, blockNumber *big.Int) (string, error) {
	if tx == nil {
		return "", nil
	}

	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		signer := types.LatestSignerForChainID(tx.ChainId())
		from, err = types.Sender(signer, tx)
		if err != nil {
			return "", err
		}
	}
	msg := ethereum.CallMsg{
		From: from,
		To:   tx.To(),
		Gas:  tx.Gas(),

		Value: tx.Value(),
		Data:  tx.Data(),
	}
	hex, err := c.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return "", err
	}

	unpackedMsg, err := abi.UnpackRevert(hex)
	if err != nil {
		fmt.Printf("failed to get the revert message for tx %v: %v\n", tx.Hash(), err)
		return "", errors.New("execution reverted")
	}

	return unpackedMsg, nil
}
