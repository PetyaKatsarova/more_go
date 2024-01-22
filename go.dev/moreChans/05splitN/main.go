package main

import (
	"fmt"
	"time"
)

func splitN(from chan string, n uint) []chan string {
	chansWithStr := make([]chan string, int(n)) // allocate memory 
	// for slice of chans str with len or buffer? n
	for i := uint(0); i < n; i++ {
		// populate it with chans str
		chansWithStr[i] = make(chan string, cap(from))
	}

	go func ()  {
		defer func ()  {
			for _, val := range chansWithStr {
				close(val)
			}
		}()

// every time 'from' chan receives a val, loop through all chansWithStr
// chans end send to all that val
		for val := range from {
			for _, val2 := range chansWithStr {
				val2 <- val
			}
		}
	}()
	return chansWithStr
}

func main() {
	c 				:= make(chan string, 5)
	chansWithstr	:= splitN(c, 2) // create slice of 2 chans str
	go func ()  {
		for s := range chansWithstr[0] {
			fmt.Printf("read chanWithstr[0] val\"%s\"\n", s)
		}
		fmt.Println("chanWithstr[0] closed")
	}()

	go func ()  {
		for s := range chansWithstr[1] {
			fmt.Printf("read chanWithstr[1] val \"%s\"\n", s)
		}
		fmt.Println("chanWithstr[1] closed")
	}()
	c <- "str 1"
	c <- "str 2"
	c <- "str 3"
	fmt.Printf("channel c has %d/%d items\n", len(c), cap(c))
	fmt.Printf("channel chansWithstr[0] has %d/%d items\n", len(chansWithstr[0]), cap(chansWithstr[0]))
	fmt.Printf("channel chansWithstr[1] has %d/%d items\n", len(chansWithstr[1]), cap(chansWithstr[1]))

	close(c)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("finish")
}