package storage

import (
	"os"
	"path/filepath"
	"sync"
)

type LSMStore struct {
	mu sync.RWMutex

	dir string

	wal      *WAL
	memtable *MemTable

	sstables []*SSTable

	maxMemTableSize int
}

func OpenLSMStore(dir string, maxMemTableSize int) (*LSMStore, error) {
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return nil, err
	}

	walPath := filepath.Join(dir, "wal.log")

	wal, err := OpenWAL(walPath)

	if err != nil {
		return nil, err
	}

	store := &LSMStore{
		dir:             dir,
		wal:             wal,
		memtable:        NewMemTable(),
		sstables:        make([]*SSTable, 0),
		maxMemTableSize: maxMemTableSize,
	}

	recovered, err := wal.Replay()

	if err != nil {
		wal.Close()
		return nil, err
	}

	for key, value := range recovered {
		store.memtable.Set(key, value)
	}

	return store, nil
}

func (s *LSMStore) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.wal.Append(key, value)

	if err != nil {
		return err
	}

	s.memtable.Set(key, value)

	if s.memtable.Size() >= s.maxMemTableSize {
		return s.flush()
	}

	return nil
}

func (s *LSMStore) Get(key string) (string, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.memtable.Get(key)

	if ok {
		return value, true, nil
	}

	for i := len(s.sstables) - 1; i >= 0; i-- {
		value, ok, err := s.sstables[i].Get(key)

		if err != nil {
			return "", false, err
		}

		if ok {
			return value, true, nil
		}
	}

	return "", false, nil
}

func (s *LSMStore) flush() error {
	data := make(map[string]string)

	for key, value := range s.memtable.data {
		data[key] = value
	}

	filename := filepath.Join(
		s.dir,
		"sstable-"+string(rune(len(s.sstables)+'0'))+".sst",
	)

	err := WriteSSTable(filename, data)

	if err != nil {
		return err
	}

	s.sstables = append(
		s.sstables,
		OpenSSTable(filename),
	)

	s.memtable = NewMemTable()

	return nil
}

func (s *LSMStore) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.wal.Close()
}
