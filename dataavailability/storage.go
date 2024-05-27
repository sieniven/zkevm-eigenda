package dataavailability

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

var ErrNotFound = errors.New("not found")

// In-memory data availability storage for the mock implementation.
// Will need a thorough design to think through the backend storage for DA information on
// the node. For now, this mock node PoC quickly implements this by indexing block hashes
// to the index of the sequence data that is submitted on the EigenDA layer.
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

func (s *DAStorage) Add(hash common.Hash, info BlobInfo) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.inner[hash] = info
	return nil
}

func (s *DAStorage) Update(hash common.Hash, info BlobInfo) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.inner[hash] = info
	return nil
}
