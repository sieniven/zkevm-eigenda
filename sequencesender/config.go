package sequencesender

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sieniven/zkevm-eigenda/config/types"
)

type Config struct {
	// WaitPeriodSendSequence is the time the sequencer waits until
	// trying to send a sequence to L1
	WaitPeriodSendSequence types.Duration `mapstructure:"WaitPeriodSendSequence"`

	// MaxTxSizeForL1 is the maximum size a single transaction can have. This field has
	// non-trivial consequences: larger transactions than 128KB are significantly harder and
	// more expensive to propagate; larger transactions also take more resources
	// to validate whether they fit into the pool or not.
	MaxTxSizeForL1 uint64 `mapstructure:"MaxTxSizeForL1"`

	// SenderAddress defines which private key the eth tx manager needs to use
	// to sign the L1 txs
	SenderAddress common.Address `mapstructure:"SenderAddress"`

	// L2Coinbase defines which address is going to receive the fees
	L2Coinbase common.Address `mapstructure:"L2Coinbase"`

	// GasOffset is the amount of gas to be added to the gas estimation in order
	// to provide an amount that is higher than the estimated one. This is used
	// to avoid the TX getting reverted in case something has changed in the network
	// state after the estimation which can cause the TX to require more gas to be
	// executed.
	//
	// ex:
	// gas estimation: 1000
	// gas offset: 100
	// final gas: 1100
	GasOffset uint64 `mapstructure:"GasOffset"`

	// MaxBatchesForL1 is the maximum amount of batches to be sequenced in a single L1 tx
	MaxBatchesForL1 uint64 `mapstructure:"MaxBatchesForL1"`

	// Mock config to replicate the BatchConstraintsCfg in the zkevm node.
	MaxBatchBytesSize uint64 `mapstructure:"MaxBatchBytesSize"`
}
