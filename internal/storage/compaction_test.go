package storage

import (
	"path/filepath"
	"testing"
)

func TestCompaction(t *testing.T) {
	dir := t.TempDir()

	store, err := OpenLSMStore(dir, 1)

	if err != nil {
		t.Fatal(err)
	}

	defer store.Close()

	if err := store.Set("name", "harshitha"); err != nil {
		t.Fatal(err)
	}

	if err := store.Set("city", "Vijayawada"); err != nil {
		t.Fatal(err)
	}

	if len(store.sstables) < 2 {
		t.Fatalf("expected multiple SSTables, got %d", len(store.sstables))
	}

	if err := store.Compact(); err != nil {
		t.Fatal(err)
	}

	if len(store.sstables) != 1 {
		t.Fatalf("expected 1 SSTable after compaction, got %d", len(store.sstables))
	}

	value, ok, err := store.Get("name")

	if err != nil {
		t.Fatal(err)
	}

	if !ok || value != "harshitha" {
		t.Fatalf("wrong value: %s", value)
	}

	compactFile := filepath.Join(dir, "compacted.sst")

	if _, err := filepath.Abs(compactFile); err != nil {
		t.Fatal(err)
	}
}
