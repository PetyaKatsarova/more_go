package main

import (
	"fmt"
	// "math/rand"
	"reflect"
	"sync"
	// "time"
)

func main() {
	a := asChan(0,1,2,3,4,5,6,7,8,9) // a is a chan with no buffer receiving all those ints
	b := asChan(10,11,12,13,14,15,16,17,18,19)
	c := asChan(20,21,22,23,24,25,26,27,28,29)
	for v := range mergeReflect(a,b,c) {   //merge(a,b,c) {
		fmt.Println(v)
	}
}

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	go func ()  {
		var wg sync.WaitGroup
		wg.Add((len(chans)))

		for _, c := range chans {
			go func (c <-chan int)  { // cant simply use c because of race or sth: will send only on last chan
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func asChan(vals ...int) <-chan int {
	c := make(chan int)
	go func ()  {
		for _, v := range vals {
			c <- v
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func mergeReflect(chans ...<-chan int) <- chan int {
	out := make(chan int)

	go func ()  {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase {
				Dir: reflect.SelectRecv, // receive direction	
				Chan: reflect.ValueOf(c),
			})
		}
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases) // is the index of the chosen chan, v is val and bool
			if !ok {  // chan is closed
				cases = append(cases[:i], cases[i+1:]...) // remove the closed chan
				continue
			}
			out <- v.Interface().(int)
		}
	}()
	return out
}

// type SelectCase struct {
// 	Dir		reflect.SelectDir// direction of case
// 	Chan	reflect.Value // chan to use for send or receive
// 	Send	reflect.Value // val to send
// }

/*
The reflect package provides a set of functions and types for working with reflection, which is the ability to inspect
 and manipulate the type, structure, and values of variables at runtime.  Reflection allows you to write more generic and
  flexible code because you can work with types and values dynamically without knowing their specifics at compile time.
   However, it should be used with caution as it can make your code less type-safe and less efficient.
   Type Conversion: You can convert between different types using reflection. This can be useful when working with interfaces 
   and need to convert interface values to their
    concrete types.
	reflect.Select: used for performing non-blocking reflection
	 on multiple channels. It allows you to wait for a value to be sent or received on one of several channels, whichever
	  happens first. This is particularly useful when you need to coordinate multiple Goroutines or select from multiple
	   channels without blocking indefinitely.
	   func Select(cases []SelectCase) (chosen int, recv reflect.Value, recvOK bool)

*/


// import (
// 	"fmt"
// 	"reflect"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch1 <- 42
// 	}()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch2 <- "hello"
// 	}()

// 	cases := []reflect.SelectCase{
// 		{
// 			Dir:  reflect.SelectRecv,
// 			Chan: reflect.ValueOf(ch1),
// 		},
// 		{
// 			Dir:  reflect.SelectRecv,
// 			Chan: reflect.ValueOf(ch2),
// 		},
// 	}

// 	chosen, value, recvOK := reflect.Select(cases)

// 	if recvOK {
// 		fmt.Printf("Selected case %d: %v\n", chosen, value.Interface())
// 	} else {
// 		fmt.Println("All cases blocked.")
// 	}
// }
