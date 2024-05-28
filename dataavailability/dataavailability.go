package dataavailability

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sieniven/zkevm-eigenda/etherman/types"
)

// DataAvailability is a mock implementation of the zkevm DataAvailability integration. It implements
// an abstract data availability integration that holds the DABackend implementation as composition.
type DataAvailability struct {
	BatchDataProvider
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
	// Pre-process sequence data to send to the DA layer
	batchesData := [][]byte{}
	batchesHash := []common.Hash{}
	for _, batch := range sequences {
		hash := crypto.Keccak256Hash(batch.BatchL2Data)
		batchesData = append(batchesData, batch.BatchL2Data)
		batchesHash = append(batchesHash, hash)
	}

	msg, err := d.backend.PostSequence(ctx, batchesData)
	if err != nil {
		return nil, err
	}

	// Index the DA blob information to the batch hash in storage
	for _, hash := range batchesHash {
		err = d.backend.StoreDataAvailabilityMessage(ctx, hash, msg)
		if err != nil {
			return nil, err
		}
	}

	return msg, nil
}

// GetBatchL2Data in the zkevm node implementation tries to return the data from a batch in the following
// priorities:
// 1. From local DB
// 2. From Sequencer
// 3. From DA backend
//
// For this minimal mock implementation, we will test the lowest priority return method from the DA backend.
func (d *DataAvailability) GetBatchL2Data(batchNums []uint64, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	fmt.Printf("trying to get data from DA backend for batches %v\n", batchNums)
	return d.backend.GetSequence(d.ctx, batchHashes, dataAvailabilityMessage)
}

// Get batch data from EigenDA request ID
func (d *DataAvailability) GetBatchL2DataFromRequestId(ctx context.Context, id []byte) ([][]byte, error) {
	fmt.Println("trying to get batch data from EigenDA request id: ", string(id))
	info, err := d.backend.GetBlobInformationFromId(ctx, id)
	if err != nil {
		return nil, err
	} else {
		return d.GetBatchL2Data([]uint64{}, []common.Hash{}, info)
	}
}
