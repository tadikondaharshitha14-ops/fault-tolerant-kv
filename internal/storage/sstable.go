package storage

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type SSTable struct {
	path string
}

func WriteSSTable(path string, data map[string]string) error {
	keys := make([]string, 0, len(data))

	for key := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, key := range keys {
		_, err := writer.WriteString(key + "\t" + data[key] + "\n")

		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func (s *SSTable) Get(key string) (string, bool, error) {
	file, err := os.Open(s.path)

	if err != nil {
		return "", false, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "\t", 2)

		if len(parts) != 2 {
			continue
		}

		if parts[0] == key {
			return parts[1], true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", false, err
	}

	return "", false, nil
}

func OpenSSTable(path string) *SSTable {
	return &SSTable{
		path: path,
	}
}
