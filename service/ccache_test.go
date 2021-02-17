package service

import (
	"testing"
)

// func TestCCache(t *testing.T) {

// 	cc := NewCCConn()

// 	bSet := testing.Benchmark(func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			cc.Set("A", "Test")
// 		}
// 	})

// 	t.Log("time-get", time.Duration(bSet.NsPerOp()))

// 	bGet := testing.Benchmark(func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			cc.Get("A")
// 		}
// 	})
// 	t.Log("time-set", time.Duration(bGet.NsPerOp()))
// }

func benchmarkCCSet(key, val interface{}, cc *CcacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.Set(key, val)
	}
}

func benchmarkCCGet(key interface{}, cc *CcacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.Get(key)
	}
}

func BenchmarkCCA(b *testing.B) {
	cc := NewCCConn()
	benchmarkCCSet("A", "Hello", cc, b)
	benchmarkCCGet("A", cc, b)
}

func BenchmarkCCB(b *testing.B) {
	cc := NewCCConn()
	benchmarkCCSet("B", "World", cc, b)
	benchmarkCCGet("B", cc, b)
}
