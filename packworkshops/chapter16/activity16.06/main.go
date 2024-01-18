package main

import (
	"log"
	"time"
)

func push(from, to int, out chan int) {
	for i := from; i <= to; i++ {
		out <- i
		time.Sleep(time.Millisecond)
	}
}

func main() {
	res := 0
	ch  := make(chan int, 100)

	go push(1, 25, ch)
	go push(26,50, ch)
	go push(51, 75, ch)
	go push(76, 100, ch)

	for j := 0; j <100; j++ {
		i := <- ch
		log.Println(i)
		res += i
	}
	log.Println(res)
}