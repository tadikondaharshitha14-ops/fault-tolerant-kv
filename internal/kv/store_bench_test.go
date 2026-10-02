package kv

import "testing"

func BenchmarkSet(b *testing.B) {
	store := NewStore()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		store.Set("key", "value")
	}
}

func BenchmarkGet(b *testing.B) {
	store := NewStore()
	store.Set("key", "value")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		store.Get("key")
	}
}

func BenchmarkParallelSet(b *testing.B) {
	store := NewStore()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Set("key", "value")
		}
	})
}

func BenchmarkParallelGet(b *testing.B) {
	store := NewStore()
	store.Set("key", "value")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Get("key")
		}
	})
}
