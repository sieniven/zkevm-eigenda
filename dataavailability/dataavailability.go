package dataavailability

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sieniven/zkevm-eigenda/etherman/types"
)

// DataAvailability is a mock implementation of the zkevm DataAvailability integration. It implements
// an abstract data availability integration that holds the DABackend implementation as composition.
type DataAvailability struct {
	ctx     context.Context
	backend DABackender
}

// Factory method for a new data availability manager instance
func New(cfg Config) *DataAvailability {
	da := &DataAvailability{
		ctx:     context.Background(),
		backend: NewDataProvider(cfg),
	}
	return da
}

// PostSequence sends the sequence data to the data availability backend, and returns the dataAvailabilityMessage
// as expected by the contract
func (d *DataAvailability) PostSequence(ctx context.Context, sequences []types.Sequence) ([]byte, error) {
	batchesData := [][]byte{}
	for _, batch := range sequences {
		// Do not send to the DA backend data that will be stored to L1
		if batch.ForcedBatchTimestamp == 0 {
			batchesData = append(batchesData, batch.BatchL2Data)
		}
	}
	return d.backend.PostSequence(ctx, batchesData)
}

// GetBatchL2Data in the zkevm node implementation tries to return the data from a batch in the following
// priorities:
// 1. From local DB
// 2. From Sequencer
// 3. From DA backend
//
// For this minimal mock implementation, we will test the lowest priority return method from the DA backend.
func (d *DataAvailability) GetBatchL2Data(batchNums []uint64, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	return d.backend.GetSequence(d.ctx, batchHashes, dataAvailabilityMessage)
}
