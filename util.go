package go_benchmarks

import (
	"fmt"
	"runtime"
	"testing"
)

func getCpuCoreOptions() []int {
	var coreOptions []int

	cores := runtime.NumCPU()
	for cores != 1 {
		coreOptions = append(coreOptions, cores)
		cores /= 2
	}

	return coreOptions
}

func runOnAllCoreOptions(b *testing.B, name string, run func(b *testing.B)) {
	options := getCpuCoreOptions()
	for _, op := range options {
		runtime.GOMAXPROCS(op)
		b.Run(fmt.Sprintf("[%d Cores] %s", op, name), run)
	}
}
