package kv

import "testing"

func TestSetGet(t *testing.T) {
	store := NewStore()

	store.Set("name", "harshitha")

	value, ok := store.Get("name")

	if !ok {
		t.Fatal("expected key to exist")
	}

	if value != "harshitha" {
		t.Fatalf("expected harshitha, got %s", value)
	}
}

func TestDelete(t *testing.T) {
	store := NewStore()

	store.Set("name", "harshitha")
	store.Delete("name")

	_, ok := store.Get("name")

	if ok {
		t.Fatal("expected key to be deleted")
	}
}
