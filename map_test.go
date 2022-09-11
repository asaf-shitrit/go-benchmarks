package go_benchmarks

import "testing"

func BenchmarkMapRead(b *testing.B) {

	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := m[i]
		_ = a
	}
}

func BenchmarkMapWrite(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}
