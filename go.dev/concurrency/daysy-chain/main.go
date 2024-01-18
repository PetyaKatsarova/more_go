package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func f(left, right chan int) {
// 	left <- 1 + <-right
// }

// func main() {
// 	const n = 100000
// 	leftmost := make(chan int)
// 	right := leftmost
// 	left := leftmost
// 	for i := 0; i < n; i++ {
// 		right = make(chan int)
// 		go f(left, right)
// 		left = right
// 	}
// 	go func(c chan int) { c <- 1 }(right)
// 	fmt.Println(<- leftmost)
// }

// fake google search engine
// !! REDUCE TAIL LATENCY USING REPLICATED SEARCH SERVERS

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result struct {
	name string
}
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{name: fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

func main() {
	start := time.Now()
	// results := Google("golang")
	results := First("golang",
				fakeSearch("replica 1"),
				fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	// go func() { c <- First(query, Web1, Web2) }() // fanning pattern: get the data on the same chan and this is parallel!!
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query)}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}
