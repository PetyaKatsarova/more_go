package main

import (
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_Main(t *testing.T) {
	main()
	bts, err := os.ReadFile("result.txt")
	if err != nil {t.Error(err)}

	if len(string(bts)) != 14 { t.Error("Wrong string", string(bts))}
}

func TestSplitter(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	in		:= make(chan int, 10)
	odd		:= make(chan int)
	even	:= make(chan int)
	go splitter(in, odd, even, wg)

	for i := 1; i <= 10; i++ {
		in <- i
	}
	close(in)

	oddSum, evenSum := 0, 0
	for i := 1; i <= 10; i++ {
		select {
		case i := <- odd:
			oddSum += i
		case i := <- even:
			evenSum += i
		}
	}
	wg.Wait()

	if oddSum != 1+3+5+7+9 { t.Error("Odd sum should not be: ", oddSum)}
	if evenSum != 2+4+6+8+10 { t.Error("Odd sum should not be: ", evenSum)}
}

func TestSum(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	in	:= make(chan int, 10) // receives all the nums
	out	:= make(chan int, 1) // just receives the sum

//launch a goroutine func, which listens on the in
// chan and calculates the sum of the numbers received.
	go sum(in, out, wg)

	for i := 0; i < 10; i++ {
		in <- 1
	}
	close(in)

	sm := <- out // out already has the sum from the method sum()
	if sm != 10 { t.Errorf("Expected 10 but received %d", sm )}
}

func TestMerger(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	even	:= make(chan int, 10)
	odd		:= make(chan int, 1)
	outF1	:= "test_output.txt"

	go merger(even, odd, wg, outF1)

	even <- 10
	odd <- 20
	wg.Wait()

	bts, err := os.ReadFile(outF1)
	if err != nil { t.Error(t) }

	if len(string(bts)) != 15 {
		t.Error("Wrong str", string(bts), len(string(bts)))
	}
	
	evens := strings.ReplaceAll(string(bts), "Odd 20\n", "")
	if len(evens) >= 14 {
		t.Error("No odds removed")
	}

	empty := strings.ReplaceAll(evens, "Even 10\n", "")
	if len(empty) !=0 {
		t.Error("non empty string", empty)
	}
}