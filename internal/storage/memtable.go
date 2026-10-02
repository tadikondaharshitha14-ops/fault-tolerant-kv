package storage

import "sync"

type MemTable struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMemTable() *MemTable {
	return &MemTable{
		data: make(map[string]string),
	}
}

func (m *MemTable) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *MemTable) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.data[key]
	return value, ok
}

func (m *MemTable) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
}

func (m *MemTable) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}
