package storage

import "testing"

func TestMemTableSetGet(t *testing.T) {
	mem := NewMemTable()

	mem.Set("name", "harshitha")

	value, ok := mem.Get("name")

	if !ok {
		t.Fatal("expected key to exist")
	}

	if value != "harshitha" {
		t.Fatalf("expected harshitha, got %s", value)
	}
}

func TestMemTableDelete(t *testing.T) {
	mem := NewMemTable()

	mem.Set("name", "harshitha")
	mem.Delete("name")

	_, ok := mem.Get("name")

	if ok {
		t.Fatal("expected key to be deleted")
	}
}

func TestMemTableSize(t *testing.T) {
	mem := NewMemTable()

	mem.Set("name", "harshitha")
	mem.Set("city", "Vijayawada")

	if mem.Size() != 2 {
		t.Fatalf("expected size 2, got %d", mem.Size())
	}
}
