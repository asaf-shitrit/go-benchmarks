package go_benchmarks

import (
	"fmt"
	"math"
	"testing"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func getSliceSizes(cap int) []int {
	sizes := make([]int, 0)
	for i := 0; i < cap; i++ {
		sizes = append(sizes, powInt(2, i))
	}
	return sizes
}

func BenchmarkReadSlice(b *testing.B) {
	sizes := getSliceSizes(16)

	for _, size := range sizes {
		b.Run(fmt.Sprintf("read from slice of size: %d", size), func(b *testing.B) {
			s := make([]int, size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				a := s[i%len(s)]
				_ = a
			}
		})
	}
}

func BenchmarkWriteSlice(b *testing.B) {
	b.Run("no cap", func(b *testing.B) {
		s := make([]int, 0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s = append(s, b.N)
		}
	})

	b.Run("with cap", func(b *testing.B) {
		s := make([]int, 0, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s = append(s, b.N)
		}
	})

	b.Run("same size", func(b *testing.B) {
		s := make([]int, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s[i] = i
		}
	})
}

func BenchmarkSliceLengthOptimization(b *testing.B) {

	s := make([]int, 0)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}

	b.Run("iterate over all", func(b *testing.B) {
		_ = s[len(s)-1]
		for i := 0; i < b.N; i++ {
			_ = s[b.N]
		}
	})
	b.Run("iterate over existing", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = s[b.N]
		}
	})
}
