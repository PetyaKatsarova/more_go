package main

import (
	"fmt"
	"sync"
	"time"
)

type Splitter struct {
	from chan string
	tos  []chan string
	mu   sync.RWMutex
}

// return new splitter where data from 'from ' is sent to each to chan from 'tos' slice of chans str:
// no data yet, just opening the chans and directions
func (splitter *Splitter) start() *Splitter {
	go func ()  {
		defer splitter.close() // closes all tos chans
		for data := range splitter.from { // send to each to chan from tos the data from from
			splitter.mu.RLock()
			for _, to := range splitter.tos {
				to <- data
			}
			splitter.mu.RUnlock()
		}
	}()
	return splitter
}

func  (splitter *Splitter) close()  {
	splitter.mu.Lock()
	defer splitter.mu.Unlock()
	for _, to := range splitter.tos {
		close(to)
	}
}

// remove last chan from toes and replace with newly created
/*
It returns two values: the to channel and a closure (a function).
The to channel is returned as the first value. This is the channel where you can send data to.
The closure is returned as the second value. This closure can be executed later to close the to channel and remove it from
 the tos slice.
*/
func (splitter *Splitter) Split() (chan string, func()) {
	to := make(chan string, cap(splitter.from))
	splitter.mu.Lock()
	splitter.tos = append(splitter.tos, to) // append new chan 'to'with buffer size 'from'
	splitter.mu.Unlock()
	return to, func() {
		splitter.mu.Lock()
		defer splitter.mu.Unlock()
		for i, ch := range splitter.tos {
			if ch == to {
				close(ch)
				splitter.tos = append(splitter.tos[:i], splitter.tos[i+1:]...)
				break
			}
		}
	}
}

func newSplitter(from chan string) *Splitter {
	return (&Splitter{from: from}).start()
}

func main() {
	c 				:= make(chan string, 5)
	splitter		:= newSplitter(c)
	c1, c1Closefn 	:= splitter.Split() // c1 and c2 chans that are added and then can be closed and removed from the tos
	                                    // slice with the closer
	c2, _ 			:= splitter.Split()

	go func ()  {
		for s := range c1 {
			fmt.Println("read c1 val: ", s)
		}
		fmt.Println("c1 closed")
	}()
	go func ()  {
		for s := range c2 {
			fmt.Println("read c2 val: ", s)
		}
		fmt.Println("c2 closed")
	}()

	c <- "str 1"
	c <- "str 2"
	time.Sleep(100*time.Millisecond)
	c1Closefn() // read from c chan and then close and remove prev chan from toes

	c <- "str 3"
	fmt.Printf("chan c has %d/%d items\n", len(c), cap(c))
	fmt.Printf("chan c1 has %d/%d items\n", len(c1), cap(c1))
	fmt.Printf("chan c2 has %d/%d items\n", len(c2), cap(c2))
	close(c)

	time.Sleep(100*time.Millisecond)
	fmt.Println("finish")
}