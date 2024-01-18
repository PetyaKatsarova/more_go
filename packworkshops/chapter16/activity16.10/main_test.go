package main

import (
	"sync"
	"testing"
)

func TestSum(t *testing.T) {

	for c := 1; c <= 100; c++ {
		mtx	:= &sync.Mutex{}
		in	:= make(chan int, 100)
		out := make(chan int)
		wrk := Worker{in: in, out: out, mtx: mtx}

		// inside the 100 loop
		for w := 1; w <= c; w++ {
			wrk.readThem() // in go routine sums the int received in the chan and sends the sum to w.out
		}
		for i := 1; i <= 100; i++ {
			in <- i
		}
		close(in)
		res := wrk.getResult()
		if res != 5050 { t.Errorf("Expexted 5050 received %d", res) }
	}
}