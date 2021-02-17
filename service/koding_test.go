package service

import "testing"

func benchmarkKodingSet(key, val interface{}, k *CacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		k.Set(key, val)
	}
}

func benchmarkKodingGet(key interface{}, k *CacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		k.Get(key)
	}
}

func BenchmarkKodingA(b *testing.B) {
	k := NewConnKCache()
	benchmarkKodingSet("A", "Hello", k, b)
	benchmarkKodingGet("A", k, b)
}

func BenchmarkKodingB(b *testing.B) {
	k := NewConnKCache()
	benchmarkKodingSet("B", "World", k, b)
	benchmarkKodingGet("B", k, b)
}

func benchmarkSharedSet(partition string, key, val interface{}, s *CacheSharedConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Set(partition, key, val)
	}
}

func benchmarkSharedGet(partition string, key interface{}, s *CacheSharedConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Get(partition, key)
	}
}

func BenchmarkSharedA(b *testing.B) {
	k := NewShardedCacheWithTTL()
	benchmarkSharedSet("node-1", "A", "Hello", k, b)
	benchmarkSharedGet("node-1", "A", k, b)
}

func BenchmarkSharedB(b *testing.B) {
	k := NewShardedCacheWithTTL()
	benchmarkSharedSet("node-2", "B", "World", k, b)
	benchmarkSharedGet("node-2", "B", k, b)
}
