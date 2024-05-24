package dataavailability

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// DABackender is an interface for components that store and retrieve batch data
type DABackender interface {
	SequenceRetriever
	SequenceSender
	// Init initializes the DABackend
	Init() error
}

// SequenceSender is used to send provided sequence of batches
type SequenceSender interface {
	// PostSequence sends the sequence data to the data availability backend, and returns the dataAvailabilityMessage
	// as expected by the contract
	PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error)
}

// SequenceRetriever is used to retrieve batch data
type SequenceRetriever interface {
	// GetSequence retrieves the sequence data from the data availability backend
	GetSequence(ctx context.Context, batchHashes []common.Hash, blobInfo BlobInfo) ([][]byte, error)
}
