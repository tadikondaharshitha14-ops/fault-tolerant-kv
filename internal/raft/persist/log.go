package persist

import (
	"bufio"
	"encoding/json"
	"os"
)

type Entry struct {
	Term  int    `json:"term"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Log struct {
	Path string
}

func New(path string) *Log {
	return &Log{
		Path: path,
	}
}

func (l *Log) Append(entry Entry) error {
	file, err := os.OpenFile(
		l.Path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))
	return err
}

func (l *Log) ReadAll() ([]Entry, error) {
	file, err := os.Open(l.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}

		return nil, err
	}
	defer file.Close()

	var entries []Entry

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var entry Entry

		if err := json.Unmarshal(
			scanner.Bytes(),
			&entry,
		); err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
