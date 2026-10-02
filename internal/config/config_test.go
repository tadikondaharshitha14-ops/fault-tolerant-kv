package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "nodes.json")

	data := []byte(`{
		"nodes": [
			{"id": 1, "address": "127.0.0.1:50051"},
			{"id": 2, "address": "127.0.0.1:50052"},
			{"id": 3, "address": "127.0.0.1:50053"}
		]
	}`)

	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Nodes) != 3 {
		t.Fatalf(
			"expected 3 nodes, got %d",
			len(cfg.Nodes),
		)
	}

	if cfg.Nodes[0].ID != 1 {
		t.Fatalf("expected node 1")
	}

	if cfg.Nodes[2].Address != "127.0.0.1:50053" {
		t.Fatalf("unexpected node 3 address")
	}
}
