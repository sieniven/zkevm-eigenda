package dataavailability

import (
	"errors"
	"sync"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/ethereum/go-ethereum/common"
)

var ErrNotFound = errors.New("not found")

type BlobInfo struct {
	BlobIndex            uint32
	BatchHeaderHash      []byte
	BatchRoot            []byte
	ReferenceBlockNumber uint
}

// In-memory data availability storage for the mock implementation.
type DAStorage struct {
	inner map[common.Hash]BlobInfo
	mutex *sync.RWMutex
}

func (s *DAStorage) Get(hash common.Hash) (BlobInfo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	info, ok := s.inner[hash]
	if ok {
		return info, nil
	} else {
		return BlobInfo{}, ErrNotFound
	}
}

func (s *DAStorage) Add(hash common.Hash, blob *disperser_rpc.BlobVerificationProof) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	info := BlobInfo{
		BlobIndex:            blob.BlobIndex,
		BatchHeaderHash:      blob.BatchMetadata.BatchHeaderHash,
		BatchRoot:            blob.BatchMetadata.BatchHeader.BatchRoot,
		ReferenceBlockNumber: uint(blob.BatchMetadata.ConfirmationBlockNumber),
	}
	s.inner[hash] = info
	return nil
}

func (s *DAStorage) Update(hash common.Hash, blob *disperser_rpc.BlobVerificationProof) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	info := BlobInfo{
		BlobIndex:            blob.BlobIndex,
		BatchHeaderHash:      blob.BatchMetadata.BatchHeaderHash,
		BatchRoot:            blob.BatchMetadata.BatchHeader.BatchRoot,
		ReferenceBlockNumber: uint(blob.BatchMetadata.ConfirmationBlockNumber),
	}
	s.inner[hash] = info
	return nil
}
