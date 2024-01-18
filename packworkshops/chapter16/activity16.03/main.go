package main

import (
	"sync"
	"sync/atomic"
	"log"
)

func sum(from, to int, wg *sync.WaitGroup, res *int32) {
	for i := from; i <= to; i++ {
		atomic.AddInt32(res, int32(i)) // *res = *res + int32(i)
	}
	wg.Done()
	return
}

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go sum(1,25, wg, &s1)
	go sum(26, 50, wg, &s1)
	go sum(51, 75, wg, &s1)
	go sum(76, 100, wg, &s1)
	wg.Wait()
	log.Println(s1)
}

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