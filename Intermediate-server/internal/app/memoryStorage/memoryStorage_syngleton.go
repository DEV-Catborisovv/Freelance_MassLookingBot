package memorystorage

import "sync"

var (
	instance *InMemoryStore
	once     sync.Once
)

// GetInstance returns the singleton instance of InMemoryStore.
func GetInstance() *InMemoryStore {
	once.Do(func() {
		instance = NewInMemoryStore()
	})
	return instance
}
