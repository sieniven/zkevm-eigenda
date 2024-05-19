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

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sieniven/polygoncdk-eigenda/etherman"
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

func (s *MonitoredTxsStorage) GetByStatus(ctx context.Context, owner *string, statusesFilter []MonitoredTxStatus) ([]monitoredTx, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	mTxs := []monitoredTx{}
	for sender, mTx := range s.inner {
		if sender == *owner {
			mTxs = append(mTxs, mTx)
		} else {
			for _, status := range statusesFilter {
				if mTx.status == status {
					mTxs = append(mTxs, mTx)
				}
			}
		}
	}
	return mTxs, nil
}

func (s *MonitoredTxsStorage) Update(ctx context.Context, mTx monitoredTx) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.inner[mTx.id] = mTx
	return nil
}

type Client struct {
	ctx      context.Context
	cancel   context.CancelFunc
	cfg      Config
	etherman etherman.Client
	storage  MonitoredTxsStorage
}

// Factory method for a new eth tx manager instance
func New(cfg Config, etherMan etherman.Client) *Client {
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
		case <-time.After(c.cfg.FrequenceToMonitorTxs):
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
	statusesFilter := []MonitoredTxStatus{MonitoredTxStatusCreated, MonitoredTxStatusSent, MonitoredTxStatusReorged}
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
	var err error
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
		confirmed = true
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
	if !confirmed {
		// review tx and increase gas and gas price if needed
		// review tx and increase gas and gas price if needed
		if mTx.status == MonitoredTxStatusSent {
			err := c.reviewMonitoredTx(ctx, &mTx)
			if err != nil {
				fmt.Errorf("failed to review monitored tx: %v", err)
				return
			}
			err = c.storage.Update(ctx, mTx)
			if err != nil {
				fmt.Errorf("failed to update monitored tx review change: %v", err)
				return
			}
		}

		// rebuild transaction
		// TODO: continue
	}
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
