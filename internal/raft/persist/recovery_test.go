package persist

import (
	"path/filepath"
	"testing"
)

func TestRecoverAfterRestart(t *testing.T) {
	path := filepath.Join(
		t.TempDir(),
		"raft.log",
	)

	// First process/node instance.
	first := New(path)

	err := first.Append(Entry{
		Term:  1,
		Key:   "name",
		Value: "harshitha",
	})
	if err != nil {
		t.Fatalf("first append failed: %v", err)
	}

	err = first.Append(Entry{
		Term:  1,
		Key:   "course",
		Value: "MTECH CSE",
	})
	if err != nil {
		t.Fatalf("second append failed: %v", err)
	}

	// Simulate restart by creating a new log object.
	recovered, err := Recover(path)
	if err != nil {
		t.Fatalf("recovery failed: %v", err)
	}

	if len(recovered) != 2 {
		t.Fatalf(
			"expected 2 recovered entries, got %d",
			len(recovered),
		)
	}

	if recovered[0].Value != "harshitha" {
		t.Fatalf("first value was not recovered")
	}

	if recovered[1].Value != "MTECH CSE" {
		t.Fatalf("second value was not recovered")
	}
}
