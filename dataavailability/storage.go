package dataavailability

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

var ErrNotFound = errors.New("not found")

// In-memory data availability storage for the mock implementation.
type DAStorage struct {
	inner map[common.Hash][]byte
	mutex *sync.RWMutex
}

func (s *DAStorage) Get(hash common.Hash) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	id, ok := s.inner[hash]
	if ok {
		return id, nil
	} else {
		return nil, ErrNotFound
	}
}

func (s *DAStorage) Add(hash common.Hash, requestId []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.inner[hash] = requestId
	return nil
}

func (s *DAStorage) Update(hash common.Hash, requestId []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.inner[hash] = requestId
	return nil
}
