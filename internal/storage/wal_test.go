package storage

import (
	"os"
	"testing"
)

func TestWALReplay(t *testing.T) {
	path := "test.wal"

	defer os.Remove(path)

	wal, err := OpenWAL(path)
	if err != nil {
		t.Fatal(err)
	}

	defer wal.Close()

	if err := wal.Append("name", "harshitha"); err != nil {
		t.Fatal(err)
	}

	if err := wal.Append("city", "Vijayawada"); err != nil {
		t.Fatal(err)
	}

	data, err := wal.Replay()
	if err != nil {
		t.Fatal(err)
	}

	if data["name"] != "harshitha" {
		t.Fatalf("wrong name: %s", data["name"])
	}

	if data["city"] != "Vijayawada" {
		t.Fatalf("wrong city: %s", data["city"])
	}
}
