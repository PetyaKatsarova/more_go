package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func atomicUse() {
	var count int64
	var wg sync.WaitGroup

	wg.Add(1)
	// reader
	go func() {
		defer wg.Done() // The defer ensures that wg.Done() will be called at the end of the goroutine's execution, no matter where or how the goroutine exits (normally or
			// due to an error). This is crucial for proper synchronization using WaitGroup. Without defer, if the goroutine exits early or encounters a panic, wg.Done() might not 
			//be called, leading to the main function waiting indefinitely.
		time.Sleep(time.Millisecond * 10)
		fmt.Println("count atomicUs()e in go routine", atomic.LoadInt64(&count))
	}()

	wg.Add(50)
	// wrtiers
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 10)
			atomic.AddInt64(&count, 1) // avoids race conditions: safely reads the value of count
		}()
	}
	wg.Wait()
	fmt.Println("count atomcUse() in main", count)
}
