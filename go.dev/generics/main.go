package main

import (
	"fmt"
	"time"
)

/*
 In Go, a comparable type is one that can be compared with another value of the same
 type using operators like == and !=. This is important for map keys, as they need to
  be comparable: bool, numeric types, string types, pointer types, channel, interface,
  struct, arr;
  !NB: non comparable: slices, maps, funcs (all those can be compared to nil)
*/

	/*
	   constraints:
	*/

type Num interface {
	int64 | float64
}

// func SumIntOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	func SumIntOrFloats[K comparable, V Num](m map[K]V) V {
	var s V 
	for _, v := range m {
		s += v
	}
	return s
}


func main() {
	ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

	start := time.Now();
	fmt.Printf("Generic Sums: %v and %v\n",
	SumIntOrFloats[string, int64](ints),
	SumIntOrFloats[string, float64](floats))
	fmt.Println("took: ", time.Since(start).Microseconds(), "microseconds")
}
