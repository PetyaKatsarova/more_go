package main

import (
	"fmt"
	"sync"
	"time"
)

// safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // only 1 gorooutine at a time can access the given key
	c.v[key]++
	c.mu.Unlock()
}

// returns the curr val of the counter for the given key
func(c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Print(c.Value("somekey"))
}