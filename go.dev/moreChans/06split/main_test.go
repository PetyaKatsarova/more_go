package main

import (
	"testing"
)


/*
The Benchmark_Split function you've shown is part of a Go benchmark test, which is a special type of Go test used for
 performance profiling. Benchmarks are executed using the go test command with the -bench flag.
 go test -bench=.
 go test -bench=. -benchmem
-- above line for memory allocation 
 In the benchmark code, we have two goroutines (out1 and out2) that are reading from channels but not printing anything.
  This is expected behavior for a benchmark test because benchmarking requires precise measurement of performance and doesn't
   typically involve printing output to the console.
*/

func Benchmart_Split(b *testing.B) {
	source := make(chan string, 3)
	out1, out2 := split(source)
	go func() {
		for _ = range out1 {
		}
	}()
	go func() {
		for _ = range out2 {
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		source <- "aaa"
	}
}