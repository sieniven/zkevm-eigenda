// The ethtxmanager is the mock package for polygon CDK that handles ethereum transactions to
// the L1. It makes calls to send and aggregate batch, checks for possible errors, like wrong
// nonce or gas limit too low and make correct adjustments to request according to it.
//
// Also, it tracks transaction receipt and status of tx in case tx is rejected and send signals
// to sequencer/aggregator to resend sequence/batch
package ethtxmanager

import (
	"context"
	"fmt"
	"time"

	"github.com/sieniven/polygoncdk-eigenda/etherman"
)

const (
	failureIntervalInSeconds = 5
)

type Client struct {
	ctx      context.Context
	cancel   context.CancelFunc
	cfg      Config
	etherman etherman.Client
}

// Factory method for a new eth tx manager instance
func New(cfg Config, etherMan etherman.Client) *Client {
	c := &Client{
		cfg:      cfg,
		etherman: etherMan,
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
}

// monitorTx does all the monitoring steps to the monitored tx
func (c *Client) monitorTx(ctx context.Context, mtx monitoredTx) {
}
