package main

import "log"

func push(from, to int, in chan bool, out chan int) {
	for i := from; i <= to; i++ {
		<-in // reads a value from the in channel. Since in is of type chan bool, it's likely being used as a synchronization
		// signal. The function will wait here until it receives a value from the in channel before proceeding.
		out <- i
	}
}

func main() {
	res := 0

	in := make(chan bool, 100)
	out := make(chan int, 100)

	go push(1, 25, in, out)
	go push(26, 50, in, out)
	go push(51, 75, in, out)
	go push(76, 100, in, out)

	for j := 0; j < 100; j++ {
		in <- true
		i := <-out
		log.Println(i)
		res += i
	}
	log.Println(res)
}