package service

import "testing"

// func TestLRUCache(t *testing.T) {

// 	lru := NewLRUConnection()

// 	bSet := testing.Benchmark(func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			lru.Set("A", "Test")
// 		}
// 	})
// 	t.Log("time-set", time.Duration(bSet.NsPerOp()))

// 	bGet := testing.Benchmark(func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			lru.Get("A")
// 		}
// 	})
// 	t.Log("time-get", time.Duration(bGet.NsPerOp()))

// }

func benchmarkLRUSet(key, val interface{}, lru *LRUConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		lru.Set(key, val)
	}
}

func benchmarkLRUGet(key interface{}, lru *LRUConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		lru.Get(key)
	}
}

func BenchmarkLRUA(b *testing.B) {
	lru := NewLRUConnection()
	benchmarkLRUSet("A", "Hello", lru, b)
	benchmarkLRUGet("A", lru, b)
}

func BenchmarkLRUB(b *testing.B) {
	lru := NewLRUConnection()
	benchmarkLRUSet("B", "World", lru, b)
	benchmarkLRUGet("B", lru, b)
}
