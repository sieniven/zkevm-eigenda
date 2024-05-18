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
	}
}
