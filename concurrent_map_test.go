package go_benchmarks

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type mutexMap struct {
	m  map[int]int
	mu sync.RWMutex
}

func (mm *mutexMap) Get(k int) int {
	mm.mu.RLock()
	defer mm.mu.RUnlock()
	return mm.m[k]
}

func (mm *mutexMap) Set(k, v int) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	mm.m[k] = v
}

func BenchmarkMutexMap(b *testing.B) {

	b.Run("read", func(b *testing.B) {
		m := &mutexMap{m: map[int]int{
			0: 0,
		}, mu: sync.RWMutex{}}

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				v := m.Get(0)
				_ = v
			}
		})
	})

	b.Run("write", func(b *testing.B) {
		m := &mutexMap{m: map[int]int{
			0: 0,
		}, mu: sync.RWMutex{}}

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Set(0, 1)
			}
		})
	})
}

type atomicMap struct {
	m   map[int]int
	ops *uint64
}

func (am *atomicMap) acquire() {
	for !atomic.CompareAndSwapUint64(am.ops, 0, 1) {
		time.Sleep(time.Millisecond * 100)
	}
}

func (am *atomicMap) release() {
	atomic.AddUint64(am.ops, ^uint64(0))
}

func (am *atomicMap) Get(k int) int {
	am.acquire()
	defer am.release()

	return am.m[k]
}

func (am *atomicMap) Set(k, v int) {
	am.acquire()
	defer am.release()

	am.m[k] = v
}

func BenchmarkAtomicMap(b *testing.B) {

	b.Run("read", func(b *testing.B) {
		m := &atomicMap{
			m: map[int]int{
				0: 0,
			},
			ops: new(uint64),
		}

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				v := m.Get(0)
				_ = v
			}
		})
	})

	b.Run("write", func(b *testing.B) {
		m := &atomicMap{
			m: map[int]int{
				0: 0,
			},
			ops: new(uint64),
		}

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Set(0, 1)
			}
		})
	})
}
