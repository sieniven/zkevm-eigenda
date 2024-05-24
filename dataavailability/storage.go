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
// Will need a thorough design to think through the backend storage for DA information on
// the node. For now, this mock node PoC quickly implements this by indexing block hashes
// to the index of the sequence data that is submitted on the EigenDA layer.
type DAStorage struct {
	inner    map[common.Hash]BlobInfo
	da_inner map[common.Hash]int
	mutex    *sync.RWMutex
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

func (s *DAStorage) GetIndex(hash common.Hash) (int, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	idx, ok := s.da_inner[hash]
	if ok {
		return idx, nil
	} else {
		return 0, ErrNotFound
	}
}

func (s *DAStorage) AddIndex(hash common.Hash, idx int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.da_inner[hash] = idx
	return nil
}

func (s *DAStorage) UpdateIndex(hash common.Hash, idx int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.da_inner[hash] = idx
	return nil
}
