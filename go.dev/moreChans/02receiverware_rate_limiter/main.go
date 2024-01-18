package main

import (
	"fmt"
	"time"
	"math/rand"
)

// limit: The maximum number of strings allowed to pass through within a specified interval.
func rateLimiterReceiveware(in <-chan string, limit uint, interval time.Duration) <-chan string {
	currCount	:= uint64(0) // uses more memory than uint32 but can hold much bigger nums
	ticker 		:= time.NewTicker(interval)
	out 		:= make(chan string, cap(in)) // capacity/buffer size of chan in

	// get all strs from in chan, follow some conditions, send those strs to out
	go func ()  {
		defer ticker.Stop() // will perform just before exiting the func
		defer close(out)
		for str := range in {
			if currCount >= uint64(limit) { // limits the rate of incoming strings based on a specified limit and interval and
				// then sends the strings to an output channel (out). 
				<-ticker.C //  the code blocks and waits for a signal from the ticker.C channel.
				currCount = 0
			}
			currCount++
			out <- str
		}
	}()
	return out
}

func main() {
	c := make(chan string, 1000)
	fmt.Printf("channel c has %d/%d items\n", len(c), cap(c))

	wrappedC := rateLimiterReceiveware(c, 1, time.Second) // c chan, uint limit, interval time.Duration

	var result []string
	go func ()  {
		for s := range wrappedC	{
			result = append(result, s)
			fmt.Println(time.Now())
		}
	}()

	go func ()  {
		for i := 0; i < 100000; i++ {
			c <- ""
			// add entropy in sleep time in order to have a different throughput over the time
			time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
		}	
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("finish with ", len(result), " results")
}


/*
If in is a channel, cap(in) returns the capacity of the channel, which represents the maximum number of elements that can be held
 in the channel without blocking. It tells you how many elements the channel can buffer.
For an unbuffered channel, cap(in) will be 0 because it can hold only one value at a time. For a buffered channel, 
it will return the specified buffer size.
*/