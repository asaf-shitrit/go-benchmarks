package go_benchmarks

import (
	"testing"
	"unsafe"
)

// https://itnext.io/structure-size-optimization-in-golang-alignment-padding-more-effective-memory-layout-linters-fffdcba27c61

type memoryAlignedStruct struct {
	aaa [2]bool // 2 bytes
	bbb int32   // 4 bytes
	ccc [2]bool // 2 bytes
}

type memoryUnalignedStruct struct {
	aaa [2]bool // 2 bytes
	ccc [2]bool // 2 bytes
	bbb int32   // 4 bytes
}

func BenchmarkStructMemoryAlignment(b *testing.B) {
	b.Run("memory aligned", func(b *testing.B) {
		b.ReportMetric(float64(unsafe.Sizeof(memoryUnalignedStruct{})), "bits")
	})

	b.Run("memory unaligned", func(b *testing.B) {
		b.ReportMetric(float64(unsafe.Sizeof(memoryAlignedStruct{})), "bits")
	})
}
