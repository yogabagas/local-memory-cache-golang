package service

import (
	"fmt"
	"testing"
)

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
		lru.Set(fmt.Sprintf("%d", i), val)
	}
}

func benchmarkLRUGet(key interface{}, lru *LRUConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		lru.Get(fmt.Sprintf("%d", i))
	}
}

func BenchmarkLRUSampleA(b *testing.B) {
	lru := NewLRUConnection()
	benchmarkLRUSet(nil, "Hello", lru, b)
	benchmarkLRUGet(nil, lru, b)
}

func BenchmarkLRUSampleB(b *testing.B) {
	lru := NewLRUConnection()
	benchmarkLRUSet(nil, "World", lru, b)
	benchmarkLRUGet(nil, lru, b)
}
