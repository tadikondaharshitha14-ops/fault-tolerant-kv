package storage

import (
	"fmt"
	"testing"
)

func BenchmarkLSMSet(b *testing.B) {
	dir := b.TempDir()

	store, err := OpenLSMStore(dir, 1000000)

	if err != nil {
		b.Fatal(err)
	}

	defer store.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)

		if err := store.Set(key, value); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLSMGet(b *testing.B) {
	dir := b.TempDir()

	store, err := OpenLSMStore(dir, 1000000)

	if err != nil {
		b.Fatal(err)
	}

	defer store.Close()

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)

		if err := store.Set(key, value); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i%1000)

		if _, _, err := store.Get(key); err != nil {
			b.Fatal(err)
		}
	}
}
