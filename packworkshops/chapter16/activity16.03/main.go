package main




/*
The atomic operations are typically more efficient than using traditional locking mechanisms like mutexes in cases where simple
 read and write operations are required.
*/
// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// )

// func main() {
// 	var counter int32 // Counter to be incremented atomically
// 	var wg sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			atomic.AddInt32(&counter, 1)
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()

// 	fmt.Printf("Counter: %d\n", atomic.LoadInt32(&counter))
// }