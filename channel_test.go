package go_benchmarks

import "testing"

func BenchmarkChannelWrite(b *testing.B) {

	runOnAllCoreOptions(b, "channel write int", func(b *testing.B) {
		c := make(chan int)

		go func() {
			for {
				<-c
			}
		}()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c <- i
		}
	})

	runOnAllCoreOptions(b, "channel write string", func(b *testing.B) {
		c := make(chan string)

		go func() {
			for {
				<-c
			}
		}()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			c <- ""
		}
	})
}

func BenchmarkChannelRead(b *testing.B) {
	runOnAllCoreOptions(b, "channel read int", func(b *testing.B) {
		c := make(chan string)

		go func() {
			for {
				c <- ""
			}
		}()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			<-c
		}
	})
	runOnAllCoreOptions(b, "channel read string", func(b *testing.B) {
		c := make(chan int)

		go func() {
			for {
				c <- 0
			}
		}()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			<-c
		}
	})
}
