package main

import (
	"fmt"
	"sync"
)

func split(from chan string) (chan string, chan string) {
	to1 := make(chan string, cap(from))
	to2 := make(chan string, cap(from))

	go func ()  {
		defer close(to1)
		defer close(to2)
		for data := range from {
			to1 <- data
			to2 <- data
		}
	}()
	return to1, to2
}

func main() {
	c := make(chan string, 5)
	c1, c2 := split(c)
	var wg sync.WaitGroup
	wg.Add(2) 

	go func ()  {
		defer wg.Done()
		for str := range c1 {
			fmt.Printf("read c1 val \"%s\"\n", str)
		}
		fmt.Println("c1 closed")
	}()
	go func ()  {
		defer wg.Done()
		for str := range c2 {
			fmt.Printf("read c2 val \"%s\"\n", str)
		}
		fmt.Println("c2 closed")
	}()

	c <- "str 1"
	c <- "str 2"
	close(c)
	// go func() {
	// 	defer wg.Done()
	// 	c <- "str 1"
	// 	c <- "str 2"
	// 	close(c)
	// }()

	wg.Wait()
}

/*
The Benchmark_Split function you've shown is part of a Go benchmark test, which is a special type of Go test used for
 performance profiling. Benchmarks are executed using the go test command with the -bench flag.
 go test -bench=.
 In the benchmark code, we have two goroutines (out1 and out2) that are reading from channels but not printing anything.
  This is expected behavior for a benchmark test because benchmarking requires precise measurement of performance and doesn't
   typically involve printing output to the console.
*/

// func Benchmart_Split(b *testing.B) {
// 	source := make(chan string, 3)
// 	out1, out2 := split(source)
// 	go func() {
// 		for _ = range out1 {
// 		}
// 	}()
// 	go func() {
// 		for _ = range out2 {
// 		}
// 	}()

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		source <- "aaa"
// 	}
// }