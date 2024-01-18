package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in, out chan int
	sth     int // keeps track of the num of active worker goroutines
	mtx     *sync.Mutex
}

func (w *Worker) readThem() {
	w.sth++
	go func ()  {
		partial := 0
		for i := range w.in {
			partial += i
		}
		w.out <- partial

		w.mtx.Lock()
		w.sth--
		if w.sth == 0 { close(w.out) }
		w.mtx.Unlock()
	}()
}

func (w *Worker) getResult() int {
	total	:= 0
	wg 		:= &sync.WaitGroup{}
	wg.Add(1)
	go func ()  {
		for i := range w.out {
			total += i
		}
		wg.Done()
	}()
	wg.Wait()
	return total
}

func main() {
	mtx 	:= &sync.Mutex{}
	in  	:= make(chan int, 100)
	out		:= make(chan int)
	wrNum	:= 10
	worker	:= Worker{in: in, out: out, mtx: mtx}

	for i := 1; i <= wrNum; i++ {
		worker.readThem()
	}

	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)
fmt.Println("result: ", worker.getResult(), "sth: ", worker.sth)
}