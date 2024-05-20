package sequencesender

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	// SenderAddress defines which private key the eth tx manager needs to use
	// to sign the L1 txs
	SenderAddress common.Address
}
