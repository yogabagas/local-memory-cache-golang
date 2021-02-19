package service

import (
	"fmt"
	"testing"
)

func benchmarkFreeCacheSet(key, val interface{}, f *FCacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		f.Set(fmt.Sprintf("%d", i), i)
	}
}

func benchmarkFreeCacheGet(key interface{}, f *FCacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		f.Get(fmt.Sprintf("%d", i))
	}
}

func BenchmarkFreeCacheSampleA(b *testing.B) {
	f := NewFreeCacheConn()
	benchmarkFreeCacheSet(nil, "Hello", f, b)
	benchmarkFreeCacheGet(nil, f, b)
}

func BenchmarkFreeCacheSampleB(b *testing.B) {
	f := NewFreeCacheConn()
	benchmarkFreeCacheSet(nil, "World", f, b)
	benchmarkFreeCacheGet(nil, f, b)
}
