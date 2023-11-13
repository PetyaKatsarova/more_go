package main

// import (
// 	"fmt"
// 	"sync"
// )
// /*
// u need to install CGo before running with the -race flag
// check if race condition: go run -race main.go
// ! to enable the -race flag u need 
// Integration of C and Go: CGo allows Go programs to call C libraries and use C code. This is useful when you need to use libraries that are only available in C,
//  or for certain tasks where C's lower-level access to hardware is necessary.
//  -- CONSIDERATIONS: 
//  Security and Stability: While powerful, using C code can negate some of the safety guarantees of Go, particularly around memory safety. C code can introduce security 
//  vulnerabilities or cause crashes.Build Complexity: Programs using CGo can be more complex to build, as they may require a C compiler and other tools that are not 
//  typically needed for pure Go programs.Cross-Compilation: CGo can make cross-compiling your program more complicated, as you'll need a C cross-compiler for the
//   target platform.
//   NB!! ATOMIC: Load_ Add_ Store_ Swap_ CompareAndSwap_

//   ----------- GO RUN -RACE MAIN.GO
// */

// func main() {
// 	var count int32
// 	var wg sync.WaitGroup
// 	wg.Add(5)

// 	go func ()  {
// 		defer wg.Done()
// 		count += 10
// 	}()
// 	go func ()  {
// 		defer wg.Done()
// 		count -= 15
// 	}()
// 	go func ()  {
// 		defer wg.Done()
// 		count++
// 	}()
// 	go func ()  {
// 		defer wg.Done()
// 		count = 0
// 	}()
// 	go func ()  {
// 		defer wg.Done()
// 		count = 100
// 	}()

// 	// code from atomic-use.go
// 	done := make(chan bool)
// 	go func() {
// 		atomicUse()
// 		done <- true
// 	}()
// 	<- done
// 	wg.Wait()

// 	fmt.Println("count in main...", count)
// }

func main() {
	// raceFix()
	// valueV()
	atomicUse()
}