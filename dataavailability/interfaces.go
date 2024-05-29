package dataavailability

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// DABackender is an interface for components that store and retrieve batch data
type DABackender interface {
	SequenceRetriever
	SequenceSender
	BlobRetriever
	DataAvailabilityStorage
	// Init initializes the DABackend
	Init() error
}

// SequenceSender is used to send provided sequence of batches
type SequenceSender interface {
	// PostSequence sends the sequence data to the data availability backend, and returns the blob information
	// as expected by the contract
	PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error)
}

// SequenceRetriever is used to retrieve batch data
type SequenceRetriever interface {
	// GetSequence retrieves the sequence data from the data availability backend
	GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error)
}

// BlobRetriever is used to retrieve data availability message from EigenDA blob request ID
type BlobRetriever interface {
	GetDataAvailabilityMessageFromId(ctx context.Context, requestID []byte) ([]byte, error)
}

// BatchDataProvider is used to retrieve batch data
type BatchDataProvider interface {
	// GetBatchL2Data retrieve the data of a batch from the DA backend. The returned data must be the
	// pre-image of the hash
	GetBatchL2Data(batchNum []uint64, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error)
}

type DataAvailabilityStorage interface {
	// Stores the batch's blob information sent to the DA layer to the backend storage
	StoreDataAvailabilityMessage(ctx context.Context, batchHash common.Hash, dataAvailabilityMessage []byte) error
}
