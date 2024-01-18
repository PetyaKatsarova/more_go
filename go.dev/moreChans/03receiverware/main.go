package main

import (
	"fmt"
	"strings"
)

// This receiverware will append string from original input data
func prefixReceiverware(prefix string, in <- chan string) <-chan string {
	out := make(chan string, cap(in))

	go func ()  {
		defer close(out)
		// it will fwd in data channel to out chan with prefix
		for data := range in {
			out <- prefix + data
		}	
	}()
	return out
}
// will uppercase str from original input data
func upperCaseReceiverware(in <-chan string) <-chan string {
	out := make(chan string, cap(in))

	go func ()  {
		defer close(out)
		for data := range in {
			out <- strings.ToUpper(data)
		}	
	}()
	return out
}

func main() {
	c := make(chan string, 5)
	fmt.Printf("chan c has %d/%d items\n", len(c), cap(c))

	go func ()  {
		defer close(c)
		c <- "string 01: "
		c <- "string 02: "	
	}()

	wrappedC := prefixReceiverware("kuku lala", c)
	wrappedC  = upperCaseReceiverware(wrappedC)

	for s := range wrappedC {
		fmt.Printf("read c val \"%s\"\n", s)
	}
	fmt.Println("finish")
}

