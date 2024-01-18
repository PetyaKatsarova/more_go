package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func throughputSendware(in chan<- string, throughput *uint64) chan<- string {
	previousThroughput := uint64(0)

	ticker	:= time.NewTicker(time.Second)
	quit 	:= make(chan struct{})
	go func ()  {
		defer ticker.Stop()
		for {
			select {
			case <- ticker.C:
//  atomic package is used to atomically update the value pointed to by the throughput pointer with the value of previousThroughput. 
				atomic.StoreUint64(throughput, previousThroughput) //first val is a pointer, second is val: update throughp with val of prevthroughp
				atomic.StoreUint64(&previousThroughput, 0)
			case <- quit:
				return
			}
		}	
	}()
	
	out := make(chan string, cap(in))
	go func ()  {
		defer close(out)
		for data := range out {
			atomic.AddUint64(&previousThroughput, 1)
			in <- data
		}
	}()
	return out
}

func main() {
	c := make(chan string, 100)
	fmt.Printf("chan c hs %d/%d items\n", len(c), cap(c))
	instantThroughput := uint64(0)
	ticker 			  := time.NewTicker(time.Second)
	quit 			  := make(chan struct{})
	go func ()  {
		defer ticker.Stop()
		for {
			select {
			case <- ticker.C:
				fmt.Printf("instant throughput %d", instantThroughput)
			case <- quit:
				return
			}
		}
	}()
	wrappedC := throughputSendware(c, &instantThroughput)
	var result []string
	go func ()  {
		for s := range c {
			result = append(result, s)
			time.Sleep(time.Duration(rand.Intn(100) * int(time.Microsecond)))
		}
	}()
	go func ()  {
		for i := 0; i < 100000; i++ {
			wrappedC <- ""
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Printf("finish %d results\n", len(result))
}