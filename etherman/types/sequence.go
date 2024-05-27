package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	polygonzkevm "github.com/sieniven/zkevm-eigenda/etherman/smartcontracts/polygonvalidium_xlayer"
)

// Sequence represents an operation sent to the PoE smart contract to be
// processed.
type Sequence struct {
	GlobalExitRoot, StateRoot, LocalExitRoot common.Hash
	AccInputHash                             common.Hash
	LastL2BLockTimestamp                     int64
	BatchL2Data                              []byte
	IsSequenceTooBig                         bool
	BatchNumber                              uint64
	ForcedBatchTimestamp                     int64
	PrevBlockHash                            common.Hash
}

// IsEmpty checks is sequence struct is empty
func (s Sequence) IsEmpty() bool {
	return reflect.DeepEqual(s, Sequence{})
}

// VerifiedBatch represents a VerifiedBatch
type VerifiedBatch struct {
	BlockNumber uint64
	BatchNumber uint64
	Aggregator  common.Address
	StateRoot   common.Hash
	TxHash      common.Hash
}

// SequencedBatchElderberryData represents an Elderberry sequenced batch data
type SequencedBatchElderberryData struct {
	MaxSequenceTimestamp     uint64
	InitSequencedBatchNumber uint64 // Last sequenced batch number
}

// SequencedBatch represents virtual batch
type SequencedBatch struct {
	BatchNumber   uint64
	L1InfoRoot    *common.Hash
	SequencerAddr common.Address
	TxHash        common.Hash
	Nonce         uint64
	Coinbase      common.Address
	// Struct used in Etrog
	*polygonzkevm.PolygonRollupBaseEtrogBatchData
	// Struct used in Elderberry
	*SequencedBatchElderberryData
}

// UpdateEtrogSequence represents the first etrog sequence
type UpdateEtrogSequence struct {
	BatchNumber   uint64
	SequencerAddr common.Address
	TxHash        common.Hash
	Nonce         uint64
	// Struct used in Etrog
	*polygonzkevm.PolygonRollupBaseEtrogBatchData
}
