package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// sends number read from a file: one at a time(one line) through a channel: in endless loop, chan is not closed
func source(filename string, out chan int, wg *sync.WaitGroup) {
	f, err := os.Open(filename)
	if err != nil { panic(err) }

	rd := bufio.NewReader(f)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				i, err := strconv.Atoi(str) // ugly code- repetition but the job is done
				fmt.Println(i)
				if err != nil { panic(err) }
				out <- i
				f.Close()
				wg.Done()
				return
			} else { panic(err) }
		}
		str     = strings.TrimRight(str, "\r\n")
		i, err := strconv.Atoi(str)
		fmt.Println(i)
		if err != nil { panic(err) }
		out <- i
	}
}

func splitter(in, odd, even chan int, wg *sync.WaitGroup) {
	for i := range in {
		switch i % 2 {
		case 0:
			even <- i
		case 1:
			odd <- i
		}
	}
	close(even)
	close(odd)
	wg.Done()
}

func sum(in, out chan int, wg *sync.WaitGroup) {
	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
	wg.Done()
}

// create a file resultF and write the received even and odd sums
func merger(even, odd chan int, wg *sync.WaitGroup, resultFile string) {
	rs, err := os.Create(resultFile)
	if err != nil { panic(err) }
	for i := 0; i < 2; i++ {
		select {
		case i := <-even:
			rs.Write([]byte(fmt.Sprintf("Even %d\n", i)))
		case i := <- odd:
			rs.Write([]byte(fmt.Sprintf("Odd %d\n", i)))
		}
	}
	rs.Close()
	wg.Done()
}

func main() {
	wg 		:= &sync.WaitGroup{} // wg is a pointer to waitgroup
	wg.Add(2)
	wg2 	:= &sync.WaitGroup{}
	wg2.Add(4)
	odd 	:= make(chan int)
	even 	:= make(chan int)
	out 	:= make(chan int)
	sumod	:= make(chan int)
	sumeven := make(chan int)
	go source("./input1.dat", out, wg) // sends to out num read 1 by 1 from the file
	go source("./input2.dat", out, wg)
	go splitter(out, odd, even, wg2) // sends either odd or even nums to the corresp. chan, looping through the in chan
	go sum(odd, sumod, wg2)
	go sum(even, sumeven, wg2) // sends from even the sum of all nums to sumeven
	go merger(sumeven, sumod, wg2, "./result.txt") // create new file and write on it even or odd plus the sum nums from the corresponding
	// 2 chans even, odd
	wg.Wait()
	close(out)
	wg2.Wait()
}