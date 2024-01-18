package main

import "log"

func worker(in, out chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
}

func sum(workers, from, to int) int {
	in	:= make(chan int)
	out := make(chan int, workers) // receive sums of all workers

	for i := 0; i < workers; i++ {
		go worker(in, out) // run worker for all workers cuncurrently, at the same time
	}

	for i := from; i <= to; i++ {
		in <- i
	}
	close(in)
	
	sum := 0
	for i := 0; i < workers; i++ {
		sum += <-out
	}
	close(out)
	return sum
}

func main() {
	log.Println(sum(100, 1, 200))
}