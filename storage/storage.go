package storage

import "sync"

// very basic key-value store

type Storage struct {
	mu    *sync.RWMutex
	table map[string]interface{}
}

func New() *Storage {
	return &Storage{
		mu:    new(sync.RWMutex),
		table: make(map[string]interface{}),
	}
}

func (s *Storage) Put(key string, val interface{}) error {
	s.mu.Lock()
	s.table[key] = val
	s.mu.Unlock()
	return nil
}

func (s *Storage) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	val, ok := s.table[key]
	s.mu.RUnlock()
	return val, ok
}
