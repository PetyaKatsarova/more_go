package main

import "log"

func readThem(in, out chan string) {
	for i := range in {
		log.Println(i)
	}
	out <- "done" // make sure all from in is received, explained bellow
}

func main() {
	log.SetFlags(0)
	in, out := make(chan string), make(chan string)

	go readThem(in, out)

	strs := []string{"a","b", "c", "d", "e", "f"}
	for _, s := range strs {
		in <- s
	}
	close(in)
	<- out
}

/*
The <- out statement in the main function is used to block the execution of the main function until it receives a message from 
the out channel.It effectively waits for the readThem goroutine to send the "done" message on the out channel, indicating that
 it has finished processing all the input strings.This synchronization ensures that the main function doesn't exit prematurely
  and that the program's execution is coordinated properly.
*/