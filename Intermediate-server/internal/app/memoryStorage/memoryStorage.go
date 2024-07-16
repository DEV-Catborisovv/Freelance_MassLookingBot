package memorystorage

import (
	"errors"
	"sync"
)

type InMemoryStore struct {
	store map[interface{}]interface{}
	mu    sync.RWMutex
}

// NewInMemoryStore creates and returns a new instance of InMemoryStore.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		store: make(map[interface{}]interface{}),
	}
}

// Set stores a value associated with a key in the in-memory store.
func (s *InMemoryStore) Set(key, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}

// Get retrieves a value associated with a key from the in-memory store.
func (s *InMemoryStore) Get(key interface{}) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.store[key]
	return value, exists
}

// Delete removes a key-value pair from the in-memory store.
func (s *InMemoryStore) Delete(key interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}

// GetChannel retrieves a channel associated with a key from the in-memory store.
func (s *InMemoryStore) GetChannel(key interface{}) (chan interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.store[key]
	if !exists {
		return nil, errors.New("channel not found")
	}
	ch, ok := value.(chan interface{})
	if !ok {
		return nil, errors.New("value is not a channel")
	}
	return ch, nil
}
