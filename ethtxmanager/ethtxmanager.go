// The ethtxmanager is the mock package for polygon CDK that handles ethereum transactions to
// the L1. It makes calls to send and aggregate batch, checks for possible errors, like wrong
// nonce or gas limit too low and make correct adjustments to request according to it.
//
// Also, it tracks transaction receipt and status of tx in case tx is rejected and send signals
// to sequencer/aggregator to resend sequence/batch
package ethtxmanager

type Client struct {
}
