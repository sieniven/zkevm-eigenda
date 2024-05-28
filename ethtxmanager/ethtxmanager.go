// The ethtxmanager is the mock package for polygon CDK that handles ethereum transactions to
// the L1. It makes calls to send and aggregate batch, checks for possible errors, like wrong
// nonce or gas limit too low and make correct adjustments to request according to it.
//
// Also, it tracks transaction receipt and status of tx in case tx is rejected and send signals
// to sequencer/aggregator to resend sequence/batch
package ethtxmanager

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sieniven/zkevm-eigenda/etherman"
)

const (
	failureIntervalInSeconds = 5
)

var (
	// ErrNotFound when the object is not found
	ErrNotFound = errors.New("not found")
	// ErrAlreadyExists when the object already exists
	ErrAlreadyExists = errors.New("already exists")

	// ErrExecutionReverted returned when trying to get the revert message
	// but the call fails without revealing the revert reason
	ErrExecutionReverted = errors.New("execution reverted")
)

type MonitoredTxsStorage struct {
	inner map[string]monitoredTx
	mutex *sync.RWMutex
}

func (s *MonitoredTxsStorage) Get(ctx context.Context, owner *string, id string) (monitoredTx, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	mTx, ok := s.inner[id]
	if ok {
		return mTx, nil
	} else {
		if owner != nil {
			for _, mTx := range s.inner {
				if mTx.owner == *owner {
					return mTx, nil
				}
			}
		}
		return monitoredTx{}, ErrNotFound
	}
}

func (s *MonitoredTxsStorage) GetByStatus(ctx context.Context, owner *string, statusesFilter []MonitoredTxStatus) ([]monitoredTx, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	mTxs := []monitoredTx{}
	for _, mTx := range s.inner {
		if owner == nil || *owner == mTx.owner {
			for _, status := range statusesFilter {
				if mTx.status == status {
					mTxs = append(mTxs, mTx)
				}
			}
		}
	}
	return mTxs, nil
}

func (s *MonitoredTxsStorage) Add(ctx context.Context, mTx monitoredTx) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.inner[mTx.id] = mTx
	return nil
}

func (s *MonitoredTxsStorage) Update(ctx context.Context, mTx monitoredTx) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.inner[mTx.id] = mTx
	return nil
}

type Client struct {
	ctx      context.Context
	cancel   context.CancelFunc
	cfg      Config
	etherman *etherman.Client
	storage  MonitoredTxsStorage
}

// Factory method for a new eth tx manager instance
func New(cfg Config, etherMan *etherman.Client) *Client {
	// Initialize monitored txs in-memory storage
	s := MonitoredTxsStorage{
		inner: map[string]monitoredTx{},
		mutex: &sync.RWMutex{},
	}

	c := &Client{
		cfg:      cfg,
		etherman: etherMan,
		storage:  s,
	}
	return c
}

// Add a transaction to be sent and monitored
func (c *Client) Add(ctx context.Context, owner, id string, from common.Address, to *common.Address, value *big.Int, data []byte, gasOffset uint64) error {
	// get next nonce
	nonce, err := c.etherman.CurrentNonce(ctx, from)
	if err != nil {
		return fmt.Errorf("failed to get current nonce: %w", err)
	}
	// get gas
	gas, err := c.etherman.EstimateGas(ctx, from, to, value, data)
	if err != nil {
		err := fmt.Errorf("failed to estimate gas: %w, data: %v", err, common.Bytes2Hex(data))
		if c.cfg.ForcedGas > 0 {
			gas = c.cfg.ForcedGas
		} else {
			return err
		}
	}

	// get gas price
	gasPrice, err := c.suggestedGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to get suggested gas price: %w", err)
	}

	// create monitored tx
	mTx := monitoredTx{
		owner: owner, id: id, from: from, to: to,
		nonce: nonce, value: value, data: data,
		gas: gas, gasOffset: gasOffset, gasPrice: gasPrice,
		status: MonitoredTxStatusCreated,
		// initialize empty map
		history: map[common.Hash]bool{},
		// blockNumber is unused here
		blockNumber: big.NewInt(0),
		// createdAt is unused here
		createdAt: time.Now(),
		// updatedAt is unused here
		updatedAt: time.Now(),
	}

	// add to storage
	err = c.storage.Add(ctx, mTx)
	if err != nil {
		return fmt.Errorf("failed to add tx to get monitored: %w", err)
	}
	fmt.Printf("created monitored tx: %v\n", mTx.id)
	return nil
}

// ResultsByStatus returns all the results for all the monitored txs related to the owner and matching the provided statuses
// if the statuses are empty, all the statuses are considered.
//
// the slice is returned is in order by created_at field ascending
func (c *Client) ResultsByStatus(ctx context.Context, owner string, statuses []MonitoredTxStatus) ([]MonitoredTxResult, error) {
	mTxs, err := c.storage.GetByStatus(ctx, &owner, statuses)
	if err != nil {
		return nil, err
	}

	results := make([]MonitoredTxResult, 0, len(mTxs))

	for _, mTx := range mTxs {
		result, err := c.buildResult(ctx, mTx)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

// Result returns the current result of the transaction execution with all the details
func (c *Client) Result(ctx context.Context, owner, id string) (MonitoredTxResult, error) {
	mTx, err := c.storage.Get(ctx, &owner, id)
	if err != nil {
		return MonitoredTxResult{}, err
	}

	return c.buildResult(ctx, mTx)
}

// SetStatusDone sets the status of a monitored tx to MonitoredStatusDone.
// this method is provided to the callers to decide when a monitored tx should be
// considered done, so they can start to ignore it when querying it by Status.
func (c *Client) setStatusDone(ctx context.Context, owner, id string) error {
	mTx, err := c.storage.Get(ctx, &owner, id)
	if err != nil {
		return err
	}

	mTx.status = MonitoredTxStatusDone

	return c.storage.Update(ctx, mTx)
}

func (c *Client) buildResult(ctx context.Context, mTx monitoredTx) (MonitoredTxResult, error) {
	history := mTx.historyHashSlice()
	txs := make(map[common.Hash]TxResult, len(history))

	for _, txHash := range history {
		tx, _, err := c.etherman.GetTx(ctx, txHash)
		if !errors.Is(err, ethereum.NotFound) && err != nil {
			return MonitoredTxResult{}, err
		}

		receipt, err := c.etherman.GetTxReceipt(ctx, txHash)
		if !errors.Is(err, ethereum.NotFound) && err != nil {
			return MonitoredTxResult{}, err
		}

		revertMessage, err := c.etherman.GetRevertMessage(ctx, tx)
		if !errors.Is(err, ethereum.NotFound) && err != nil && err.Error() != ErrExecutionReverted.Error() {
			return MonitoredTxResult{}, err
		}

		txs[txHash] = TxResult{
			Tx:            tx,
			Receipt:       receipt,
			RevertMessage: revertMessage,
		}
	}

	result := MonitoredTxResult{
		ID:     mTx.id,
		Status: mTx.status,
		Txs:    txs,
	}

	return result, nil
}

// Start will start the tx management, reading txs from the storage and
// send them to the L1 blockchain. It will keep monitoring them until
// they get minted.
func (c *Client) Start() {
	// Infinite loop to manage txs as they arrive
	c.ctx, c.cancel = context.WithCancel(context.Background())

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.cfg.FrequenceToMonitorTxs.Duration):
			err := c.monitorTxs(context.Background())
			if err != nil {
				c.logErrorAndWait("failed to monitor txs: %v", err)
			}
		}
	}
}

// Stop will stops the monitored tx management
func (c *Client) Stop() {
	c.cancel()
}

// logErrorAndWait used when an error is detected before trying again
func (c *Client) logErrorAndWait(msg string, err error) {
	fmt.Println(msg, err)
	time.Sleep(failureIntervalInSeconds * time.Second)
}

// monitorTxs process all pending monitored transactions
func (c *Client) monitorTxs(ctx context.Context) error {
	statusesFilter := []MonitoredTxStatus{MonitoredTxStatusCreated, MonitoredTxStatusSent}
	mTxs, err := c.storage.GetByStatus(ctx, nil, statusesFilter)
	if err != nil {
		return fmt.Errorf("failed to get created monitored txs: %v", err)
	}
	fmt.Printf("Found %v monitored tx to process\n", len(mTxs))

	wg := sync.WaitGroup{}
	wg.Add(len(mTxs))
	for _, mTx := range mTxs {
		mTx := mTx // force variable shadowing to avoid pointer conflicts
		go func(c *Client, mTx monitoredTx) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("monitoring recovered from this err: %v\n", err)
				}
				wg.Done()
			}()
			c.monitorTx(ctx, mTx)
		}(c, mTx)
	}
	wg.Wait()
	return nil
}

// monitorTx does all the monitoring steps to the monitored tx
func (c *Client) monitorTx(ctx context.Context, mTx monitoredTx) {
	// check if any of the txs in the history was confirmed
	var lastReceiptChecked types.Receipt
	// monitored tx is confirmed until we find a successful receipt
	confirmed := false
	// monitored tx doesn't have a failed receipt until we find a failed receipt for any
	// tx in the monitored tx history
	hasFailedReceipts := false
	// all history txs are considered mined until we can't find a receipt for any
	// tx in the monitored tx history
	allHistoryTxsWereMined := true
	for txHash := range mTx.history {
		mined, receipt, err := c.etherman.CheckTxWasMined(ctx, txHash)
		if err != nil {
			fmt.Printf("failed to check if tx %v was mined: %v\n", txHash.String(), err)
			continue
		}

		// If the tx is not mined yet, check that not all the tx were mined and go to the next
		if !mined {
			allHistoryTxsWereMined = false
			continue
		}
		lastReceiptChecked = *receipt

		// If the tx was mined successfully then we can set it as confirmed and break the loop
		if lastReceiptChecked.Status == types.ReceiptStatusSuccessful {
			confirmed = true
			break
		}

		// If the tx was mined but failed, we continue to consider it was not confirmed
		// and set that we found a failed receipt. This info will be used later to check
		// if nonce needs to be reviewed
		confirmed = false
		hasFailedReceipts = true
	}

	// We need to check if we need to review the nonce carefully, to avoid sending duplicate data
	// to the roll-up and causing unnecessary trusted state reorg.
	//
	// If we have failed receipts, this means at least one of the generated txs was mined.
	// In this case, maybe the curent nonce was already consumed (if this is the first iteration of
	// this cycle, next iteration might have the nonce already updated by the previous one), then
	// we need to check if there are tx that were not mined yet.
	//
	// If so, we just need to wait because maybe one of them will get mined successfully.
	if !confirmed && hasFailedReceipts && allHistoryTxsWereMined {
		fmt.Println("nonce needs to be updated")
		err := c.reviewMonitoredTxNonce(ctx, &mTx)
		if err != nil {
			fmt.Printf("failed to review monitored tx nonce: %v\n", err)
			return
		}
		err = c.storage.Update(ctx, mTx)
		if err != nil {
			fmt.Printf("failed to update the monitored tx nonce change: %v\n", err)
			return
		}
	}

	// If the history size reaches the max history size, this means that something is really wrong
	// with this tx and we are not able to identify automatically, so we can mark this as failed to
	// let the caller know something is not right and needs to be reviewed. We also do not want to
	// be reviewing and monitoring this tx indefinitely.
	// if len(mTx.history) == maxHistorySize {
	// 	mTx.status = MonitoredTxStatusFailed
	// 	fmt.Printf("marked as failed because reached the history size limit: %v", err)
	// 	// update monitored tx changes into storage
	// 	err = c.storage.Update(ctx, mTx)
	// 	if err != nil {
	// 		fmt.Printf("failed to update monitored tx when max history size limit reached: %v", err)
	// 		continue
	// 	}
	// }

	var signedTx *types.Transaction
	var err error
	if !confirmed {
		// review tx and increase gas and gas price if needed
		if mTx.status == MonitoredTxStatusSent {
			err := c.reviewMonitoredTx(ctx, &mTx)
			if err != nil {
				fmt.Printf("failed to review monitored tx: %v\n", err)
				return
			}
			err = c.storage.Update(ctx, mTx)
			if err != nil {
				fmt.Printf("failed to update monitored tx review change: %v\n", err)
				return
			}
		}

		// rebuild transaction
		tx := mTx.Tx()
		fmt.Printf("unsigned tx %v created\n", tx.Hash().String())

		// sign tx
		signedTx, err = c.etherman.SignTx(ctx, mTx.from, tx)
		if err != nil {
			fmt.Printf("failed to sign tx %v: %v", tx.Hash().String(), err)
			return
		}

		// add tx to monitored tx history
		err = mTx.AddHistory(signedTx)
		if errors.Is(err, ErrAlreadyExists) {
			fmt.Println("signed tx already existed in the history")
		} else if err != nil {
			fmt.Printf("failed to add signed tx %v to monitored tx history: %v\n", signedTx.Hash().String(), err)
			return
		} else {
			// Update monitored tx changes into storage
			err = c.storage.Update(ctx, mTx)
			if err != nil {
				fmt.Printf("failed to update monitored tx: %v\n", err)
				return
			}
			fmt.Println("signed tx added to the monitored tx history")
		}

		// Check if the tx is already in the network. If not, send it
		_, _, err = c.etherman.GetTx(ctx, signedTx.Hash())
		// if not found, send it tx to the network
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println("signed tx not found in the network")
			err := c.etherman.SendTx(ctx, signedTx)
			if err != nil {
				fmt.Printf("failed to send tx %v to network: %v\n", signedTx.Hash().String(), err)
				return
			}
			fmt.Printf("signed tx sent to the network: %v\n", signedTx.Hash().String())
			if mTx.status == MonitoredTxStatusCreated {
				// update tx status to sent
				mTx.status = MonitoredTxStatusSent
				fmt.Printf("status changed to %v\n", string(mTx.status))
				// update monitored tx changes into storage
				err = c.storage.Update(ctx, mTx)
				if err != nil {
					fmt.Printf("failed to update monitored tx changes: %v\n", err)
					return
				}
			}
		} else {
			fmt.Println("signed tx already found in the network")
		}

		fmt.Println("waiting signedTx to be mined...")

		// wait tx to get mined
		confirmed, err = c.etherman.WaitTxToBeMined(ctx, signedTx, c.cfg.WaitTxToBeMined.Duration)
		if err != nil {
			fmt.Printf("failed to wait tx to be mined: %v\n", err)
			return
		}
		if !confirmed {
			fmt.Println("signedTx not mined yet and timeout has been reached")
			return
		}

		// get tx receipt
		var txReceipt *types.Receipt
		txReceipt, err = c.etherman.GetTxReceipt(ctx, signedTx.Hash())
		if err != nil {
			fmt.Printf("failed to get tx receipt for tx %v: %v\n", signedTx.Hash().String(), err)
			return
		}
		lastReceiptChecked = *txReceipt

		// if mined, check receipt and mark as Failed or Confirmed
		if lastReceiptChecked.Status == types.ReceiptStatusSuccessful {
			mTx.status = MonitoredTxStatusConfirmed
			mTx.blockNumber = lastReceiptChecked.BlockNumber
			fmt.Printf("Tx hash %v confirmed\n", signedTx.Hash())
		} else {
			// if we should continue to monitor, we move to the next one and this will
			// be reviewed in the next monitoring cycle
			if c.shouldContinueToMonitorThisTx(ctx, lastReceiptChecked) {
				return
			}
			// otherwise we understand this monitored tx has failed
			mTx.status = MonitoredTxStatusFailed
			mTx.blockNumber = lastReceiptChecked.BlockNumber
			fmt.Printf("Tx hash %v failed\n", signedTx.Hash())
		}

		// update monitored tx changes into storage
		err = c.storage.Update(ctx, mTx)
		if err != nil {
			fmt.Printf("failed to update monitored tx: %v\n", err)
			return
		}
	}
}

// shouldContinueToMonitorThisTx checks the the tx receipt and decides if it should
// continue or not to monitor the monitored tx related to the tx from this receipt
func (c *Client) shouldContinueToMonitorThisTx(ctx context.Context, receipt types.Receipt) bool {
	// if the receipt has a is successful result, stop monitoring
	if receipt.Status == types.ReceiptStatusSuccessful {
		return false
	}

	tx, _, err := c.etherman.GetTx(ctx, receipt.TxHash)
	if err != nil {
		fmt.Printf("failed to get tx when monitored tx identified as failed, tx : %v: %v\n", receipt.TxHash.String(), err)
		return false
	}
	_, err = c.etherman.GetRevertMessage(ctx, tx)
	if err != nil {
		// if the error when getting the revert message is not identified, continue to monitor
		if err.Error() == ErrExecutionReverted.Error() {
			return true
		} else {
			fmt.Printf("failed to get revert message for monitored tx identified as failed, tx %v: %v\n", receipt.TxHash.String(), err)
		}
	}
	// if nothing weird was found, stop monitoring
	return false
}

// reviewMonitoredTxNonce checks if the nonce needs to be updated accordingly to the current
// nonce of the sender account.
//
// IMPORTANT: Nonce is reviewed apart from the other fields because it is a very sensible
// information and can make duplicated data to be sent to the blockchain, causing possible
// side effects and wasting resources.
func (c *Client) reviewMonitoredTxNonce(ctx context.Context, mTx *monitoredTx) error {
	fmt.Println("reviewing nonce")
	nonce, err := c.etherman.CurrentNonce(ctx, mTx.from)
	if err != nil {
		err := fmt.Errorf("failed to load current nonce for acc %v: %w", mTx.from.String(), err)
		return err
	}
	if nonce > mTx.nonce {
		fmt.Printf("monitored tx nonce updated from %v to %v", mTx.nonce, nonce)
		mTx.nonce = nonce
	}
	return nil
}

// reviewMonitoredTx checks if some field needs to be updated accordingly to the current
// information stored and the current state of the blockchain
func (c *Client) reviewMonitoredTx(ctx context.Context, mTx *monitoredTx) error {
	fmt.Println("reviewing")
	// get gas
	gas, err := c.etherman.EstimateGas(ctx, mTx.from, mTx.to, mTx.value, mTx.data)
	if err != nil {
		err := fmt.Errorf("failed to estimate gas: %w", err)
		return err
	}

	// check gas
	if gas > mTx.gas {
		fmt.Printf("monitored tx gas updated from %v to %v\n", mTx.gas, gas)
		mTx.gas = gas
	}

	// get gas price
	gasPrice, err := c.suggestedGasPrice(ctx)
	if err != nil {
		err := fmt.Errorf("failed to get suggested gas price: %w", err)
		return err
	}

	// check gas price
	if gasPrice.Cmp(mTx.gasPrice) == 1 {
		fmt.Printf("monitored tx gas price updated from %v to %v\n", mTx.gasPrice.String(), gasPrice.String())
		mTx.gasPrice = gasPrice
	}
	return nil
}

func (c *Client) suggestedGasPrice(ctx context.Context) (*big.Int, error) {
	// get gas price
	gasPrice, err := c.etherman.SuggestedGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	// adjust the gas price by the margin factor
	marginFactor := big.NewFloat(0).SetFloat64(c.cfg.GasPriceMarginFactor)
	fGasPrice := big.NewFloat(0).SetInt(gasPrice)
	adjustedGasPrice, _ := big.NewFloat(0).Mul(fGasPrice, marginFactor).Int(big.NewInt(0))

	// if there is a max gas price limit configured and the current
	// adjusted gas price is over this limit, set the gas price as the limit
	if c.cfg.MaxGasPriceLimit > 0 {
		maxGasPrice := big.NewInt(0).SetUint64(c.cfg.MaxGasPriceLimit)
		if adjustedGasPrice.Cmp(maxGasPrice) == 1 {
			adjustedGasPrice.Set(maxGasPrice)
		}
	}

	return adjustedGasPrice, nil
}

// ResultHandler used by the caller to handle results when processing monitored txs
type ResultHandler func(MonitoredTxResult)

// ProcessPendingMonitoredTxs will check all monitored txs of this owner and wait until
// all of them are either confirmed or failed before confirming
//
// For the confirmed and failed ones, the resultHandler will be triggered
func (c *Client) ProcessPendingMonitoredTxs(ctx context.Context, owner string, resultHandler ResultHandler) {
	statusesFilter := []MonitoredTxStatus{
		MonitoredTxStatusCreated,
		MonitoredTxStatusSent,
		MonitoredTxStatusFailed,
		MonitoredTxStatusConfirmed,
	}
	// Keep running until there are no pending monitored txs
	for {
		results, err := c.ResultsByStatus(ctx, owner, statusesFilter)
		if err != nil {
			// If something goes wrong here, we log and wait for abit and keep it in the infinite loop to not
			// unlock the caller.
			fmt.Printf("failed to get results by statuses from eth tx manager to monitored txs err: %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		if len(results) == 0 {
			// if there are not pending monitored txs, stop
			return
		}

		for _, result := range results {
			// If the result is confirmed, we set it as done. We do not stop looking into the current
			// monitored tx
			if result.Status == MonitoredTxStatusConfirmed {
				err := c.setStatusDone(ctx, owner, result.ID)
				if err != nil {
					fmt.Printf("failed to set monitored tx as done, err: %v\n", err)
					continue
				} else {
					fmt.Println("monitored tx confirmed")
				}
				resultHandler(result)
				continue
			}

			// If the result is failed, we need to go around it and rebuild a batch verification
			if result.Status == MonitoredTxStatusFailed {
				resultHandler(result)
				continue
			}

			// If the result is neither confirmed or failed, it means we need to wait until it gets
			// confirmed or failed.
			for {
				// wait before refreshing the result info
				time.Sleep(time.Second)

				// refresh the result info
				result, err := c.Result(ctx, owner, result.ID)
				if err != nil {
					fmt.Printf("failed to get monitored tx result, err: %v\n", err)
					continue
				}
				// if the result status is confirmed or failed, breaks the wait loop
				if result.Status == MonitoredTxStatusConfirmed || result.Status == MonitoredTxStatusFailed {
					break
				}
				fmt.Printf("waiting for monitored tx to get confirmed, status: %v\n", result.Status.String())
			}
		}
	}
}
