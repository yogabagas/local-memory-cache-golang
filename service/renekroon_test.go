package service

import "testing"

func benchmarkRenekroonSet(key, val interface{}, t *TCacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Set(key, val)
	}
}

func benchmarkRenekroonGet(key interface{}, t *TCacheConn, b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Get(key)
	}
}

func BenchmarkRenekroonA(b *testing.B) {
	r := NewClientConn()
	benchmarkRenekroonSet("A", "Hello", r, b)
	benchmarkRenekroonGet("A", r, b)
}

func BenchmarkRenekroonB(b *testing.B) {
	r := NewClientConn()
	benchmarkRenekroonSet("B", "World", r, b)
	benchmarkRenekroonGet("B", r, b)
}
