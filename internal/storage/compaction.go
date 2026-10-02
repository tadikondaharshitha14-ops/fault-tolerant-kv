package storage

import (
	"os"
	"path/filepath"
)

func (s *LSMStore) Compact() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.sstables) <= 1 {
		return nil
	}

	merged := make(map[string]string)

	for _, table := range s.sstables {
		file, err := os.Open(table.path)

		if err != nil {
			return err
		}

		data, err := readSSTable(file)

		file.Close()

		if err != nil {
			return err
		}

		for key, value := range data {
			merged[key] = value
		}
	}

	compactPath := filepath.Join(s.dir, "compacted.sst")

	if err := WriteSSTable(compactPath, merged); err != nil {
		return err
	}

	for _, table := range s.sstables {
		os.Remove(table.path)
	}

	s.sstables = []*SSTable{
		OpenSSTable(compactPath),
	}

	return nil
}
