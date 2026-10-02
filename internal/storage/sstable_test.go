package storage

import (
	"os"
	"testing"
)

func TestSSTableWriteAndGet(t *testing.T) {
	path := "test.sst"

	defer os.Remove(path)

	data := map[string]string{
		"name": "harshitha",
		"city": "Vijayawada",
	}

	err := WriteSSTable(path, data)

	if err != nil {
		t.Fatal(err)
	}

	table := OpenSSTable(path)

	value, ok, err := table.Get("name")

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

func TestSSTableMissingKey(t *testing.T) {
	path := "test-missing.sst"

	defer os.Remove(path)

	data := map[string]string{
		"name": "harshitha",
	}

	err := WriteSSTable(path, data)

	if err != nil {
		t.Fatal(err)
	}

	table := OpenSSTable(path)

	_, ok, err := table.Get("unknown")

	if err != nil {
		t.Fatal(err)
	}

	if ok {
		t.Fatal("expected key to be missing")
	}
}
