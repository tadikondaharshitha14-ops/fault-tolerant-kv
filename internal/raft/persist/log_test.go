package persist

import (
	"path/filepath"
	"testing"
)

func TestPersistentLog(t *testing.T) {
	path := filepath.Join(
		t.TempDir(),
		"raft.log",
	)

	log := New(path)

	err := log.Append(Entry{
		Term:  1,
		Key:   "name",
		Value: "harshitha",
	})

	if err != nil {
		t.Fatalf("append failed: %v", err)
	}

	err = log.Append(Entry{
		Term:  1,
		Key:   "course",
		Value: "MTECH CSE",
	})

	if err != nil {
		t.Fatalf("append failed: %v", err)
	}

	entries, err := log.ReadAll()

	if err != nil {
		t.Fatalf("read failed: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf(
			"expected 2 entries, got %d",
			len(entries),
		)
	}

	if entries[0].Key != "name" {
		t.Fatalf("unexpected first key: %s", entries[0].Key)
	}

	if entries[1].Value != "MTECH CSE" {
		t.Fatalf(
			"unexpected second value: %s",
			entries[1].Value,
		)
	}
}
