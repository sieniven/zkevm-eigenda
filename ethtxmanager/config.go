package ethtxmanager

import "time"

type Config struct {
	// FrequencyToMonitorTxs frequency of the resending failed txs
	FrequenceToMonitorTxs time.Duration `mapstructure:"FrequencyToMonitorTxs"`
	// WaitTxToBeMined time to wait after transaction was sent to the ethereum
	WaitTxToBeMined time.Duration `mapstructure:"WaitTxToBeMined"`
}
