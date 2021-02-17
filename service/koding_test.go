package service

import (
	"fmt"
	"testing"
)

func benchmarkKodingSet(key, val interface{}, k *CacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		k.Set(fmt.Sprintf("%d", i), val)
	}
}

func benchmarkKodingGet(key interface{}, k *CacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		k.Get(fmt.Sprintf("%d", i))
	}
}

func BenchmarkKodingA(b *testing.B) {
	k := NewConnKCache()
	benchmarkKodingSet(nil, "Hello", k, b)
	benchmarkKodingGet(nil, k, b)
}

func BenchmarkKodingB(b *testing.B) {
	k := NewConnKCache()
	benchmarkKodingSet(nil, "World", k, b)
	benchmarkKodingGet(nil, k, b)
}

func benchmarkSharedSet(partition string, key, val interface{}, s *CacheSharedConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Set(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i), val)
	}
}

func benchmarkSharedGet(partition string, key interface{}, s *CacheSharedConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Get(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
	}
}

func BenchmarkSharedSampleA(b *testing.B) {
	k := NewShardedCacheWithTTL()
	benchmarkSharedSet("", nil, "Hello", k, b)
	benchmarkSharedGet("", nil, k, b)
}

func BenchmarkSharedSampleB(b *testing.B) {
	k := NewShardedCacheWithTTL()
	benchmarkSharedSet("", nil, "World", k, b)
	benchmarkSharedGet("", nil, k, b)
}
