package service

import (
	"fmt"
	"testing"
)

func benchmarkRenekroonSet(key, val interface{}, t *TCacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Set(fmt.Sprintf("%d", i), val)
	}
}

func benchmarkRenekroonGet(key interface{}, t *TCacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Get(fmt.Sprintf("%d", i))
	}
}

func BenchmarkRenekroonSampleA(b *testing.B) {
	r := NewClientConn()
	benchmarkRenekroonSet(nil, "Hello", r, b)
	benchmarkRenekroonGet(nil, r, b)
}

func BenchmarkRenekroonSampleB(b *testing.B) {
	r := NewClientConn()
	benchmarkRenekroonSet(nil, "World", r, b)
	benchmarkRenekroonGet(nil, r, b)
}
