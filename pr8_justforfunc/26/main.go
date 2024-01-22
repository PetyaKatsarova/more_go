//https://www.youtube.com/watch?v=t9bEg2A4jsw
// learning from : justforfunc this is #26

// trying to make it faster: go mod init github.com/PetyaKatsarova/more_go/pr8_justforfunc/26 didnt help

package main

import (
	"fmt"
	"time"
	"math/rand"
	"log"
)

func main() {
	a := asChan(1,3,5,7)
	b := asChan(2,4,6,8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

// chans only read from
func merge(a, b <- chan int) <- chan int {
	c := make(chan int)
	go func ()  {
		defer close(c)

		for a != nil || b != nil {
			select {
			case v, ok := <- a:  // ok: chan is open or closed!
				if !ok {
					a = nil // blocks a chan: disabled
					log.Printf("a is done")
					continue
				}
				c <- v
			case v, ok := <- b:
				if !ok {
				   b = nil	
				   log.Printf("b is done")
					continue	
				}	
				c <- v
			}
		}
	}()
	return c
}

func  asChan(vals ...int) <-chan int {
	c := make(chan int)
	go func ()  {
		for _, v := range vals {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
