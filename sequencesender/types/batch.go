package types

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Batch struct
type Batch struct {
	BatchNumber   uint64
	Coinbase      common.Address
	BatchL2Data   []byte
	StateRoot     common.Hash
	LocalExitRoot common.Hash
	AccInputHash  common.Hash
	// Timestamp (<=incaberry) -> batch time
	// 			 (>incaberry) -> minTimestamp used in batch creation, real timestamp is in virtual_batch.batch_timestamp
	Timestamp      time.Time
	Transactions   []types.Transaction
	GlobalExitRoot common.Hash
	ForcedBatchNum *uint64
	Resources      BatchResources
	// WIP: if WIP == true is a openBatch
	WIP bool
}

// BatchResources is a struct that contains the limited resources of a batch
type BatchResources struct {
	ZKCounters ZKCounters
	Bytes      uint64
}

// ZKCounters counters for the tx
type ZKCounters struct {
	GasUsed          uint64
	KeccakHashes     uint32
	PoseidonHashes   uint32
	PoseidonPaddings uint32
	MemAligns        uint32
	Arithmetics      uint32
	Binaries         uint32
	Steps            uint32
	Sha256Hashes_V2  uint32
}
