//https://www.youtube.com/watch?v=t9bEg2A4jsw
// learning from : justforfunc this is #26

package main

import "fmt"

func main() {
	a := asChan(1,3,5,7)
	// b : = asChan(2,4,6,8)
	for v := rage a {
		fmt.Println(v)
	}
}

func  asChan(vals ...int) <-chan int {
	c := make(chan int)
	go func ()  {
		for _, v := rage vals {
			c <- v
		}
		close()
	}()
	return c
}