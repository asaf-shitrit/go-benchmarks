package go_benchmarks

import (
	"testing"
)

func BenchmarkReadArray(b *testing.B) {
	b.Run("array of size 16", func(b *testing.B) {
		arr := [16]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			a := arr[i%len(arr)]
			_ = a
		}
	})

	b.Run("array of size 128", func(b *testing.B) {
		arr := [128]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			a := arr[i%len(arr)]
			_ = a
		}
	})

	b.Run("array of size 1024", func(b *testing.B) {
		arr := [1024]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			a := arr[i%len(arr)]
			_ = a
		}
	})
}

func BenchmarkWriteArray(b *testing.B) {
	b.Run("array of size 16", func(b *testing.B) {
		arr := [16]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr[i%len(arr)] = i
		}
	})

	b.Run("array of size 128", func(b *testing.B) {
		arr := [128]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr[i%len(arr)] = i
		}
	})

	b.Run("array of size 1024", func(b *testing.B) {
		arr := [1024]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr[i%len(arr)] = i
		}
	})

	b.Run("array of size 1024", func(b *testing.B) {
		arr := [1024]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr[i%len(arr)] = i
		}
	})

	b.Run("array of size 65536", func(b *testing.B) {
		arr := [65536]int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr[i%len(arr)] = i
		}
	})

}
