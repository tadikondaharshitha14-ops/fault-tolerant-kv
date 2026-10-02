package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLSMSetGet(t *testing.T) {
	dir := t.TempDir()

	store, err := OpenLSMStore(dir, 10)

	if err != nil {
		t.Fatal(err)
	}

	defer store.Close()

	err = store.Set("name", "harshitha")

	if err != nil {
		t.Fatal(err)
	}

	value, ok, err := store.Get("name")

	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("expected key to exist")
	}

	if value != "harshitha" {
		t.Fatalf("expected harshitha, got %s", value)
	}
}

func TestLSMFlush(t *testing.T) {
	dir := t.TempDir()

	store, err := OpenLSMStore(dir, 2)

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

	files, err := filepath.Glob(filepath.Join(dir, "*.sst"))

	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 SSTable, got %d", len(files))
	}

	value, ok, err := store.Get("name")

	if err != nil {
		t.Fatal(err)
	}

	if !ok || value != "harshitha" {
		t.Fatalf("wrong value: %s", value)
	}
}

func TestLSMRecovery(t *testing.T) {
	dir := t.TempDir()

	store, err := OpenLSMStore(dir, 100)

	if err != nil {
		t.Fatal(err)
	}

	if err := store.Set("name", "harshitha"); err != nil {
		t.Fatal(err)
	}

	store.Close()

	store, err = OpenLSMStore(dir, 100)

	if err != nil {
		t.Fatal(err)
	}

	defer store.Close()

	value, ok, err := store.Get("name")

	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("expected recovered key")
	}

	if value != "harshitha" {
		t.Fatalf("expected harshitha, got %s", value)
	}

	_ = os.RemoveAll(dir)
}
