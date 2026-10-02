package cluster

import "testing"

func TestNodeSetGet(t *testing.T) {
	node := NewNode(1, ":50051")

	node.Set("name", "harshitha")

	value, ok := node.Get("name")

	if !ok {
		t.Fatal("expected key to exist")
	}

	if value != "harshitha" {
		t.Fatalf("expected harshitha, got %s", value)
	}
}

func TestNodeDelete(t *testing.T) {
	node := NewNode(1, ":50051")

	node.Set("name", "harshitha")
	node.Delete("name")

	_, ok := node.Get("name")

	if ok {
		t.Fatal("expected key to be deleted")
	}
}
