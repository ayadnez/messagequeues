package main

import (
	"fmt"
	"sync"
)

type StoreProducerFunc func() Storer
type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type MemoryStorage struct {
	// to make the storage concurrent safe , using mutex locks
	mu   sync.RWMutex
	Data [][]byte
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Data: make([][]byte, 0),
	}
}

func (m *MemoryStorage) Push(b []byte) (int, error) {
	// to make the Push concurrent safe
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Data = append(m.Data, b)
	return len(m.Data) - 1, nil
}

func (m *MemoryStorage) Fetch(offset int) ([]byte, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()
	if len(m.Data) < offset {
		return nil, fmt.Errorf("offset %d too high ", offset)
	}
	return m.Data[offset], nil
}
