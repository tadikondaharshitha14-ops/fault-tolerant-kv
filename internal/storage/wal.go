package storage

import (
	"bufio"
	"os"
	"strings"
)

type WAL struct {
	file *os.File
}

func OpenWAL(path string) (*WAL, error) {
	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &WAL{
		file: file,
	}, nil
}

func (w *WAL) Append(key, value string) error {
	_, err := w.file.WriteString(key + "\t" + value + "\n")

	if err != nil {
		return err
	}

	return w.file.Sync()
}

func (w *WAL) Replay() (map[string]string, error) {
	result := make(map[string]string)

	_, err := w.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(w.file)

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "\t", 2)

		if len(parts) != 2 {
			continue
		}

		result[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	_, err = w.file.Seek(0, 2)

	return result, err
}

func (w *WAL) Close() error {
	return w.file.Close()
}
