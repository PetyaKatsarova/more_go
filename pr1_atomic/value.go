package main 

// import (
// 	"sync"
// 	"sync/atomic"
// )

// type Config struct {
// 	a []int
// }

// func valueV() {
// 	var wg sync.WaitGroup
// 	var v atomic.Value

// 	// writer
// 	go func ()  {
// 		var i int
// 		for {
// 			i++
// 			cfg := Config {
// 				a: []int{i+1, i+2, i+3, i+4, i+5},
// 			}
// 			v.Store(cfg)
// 		}
// 	}()

// 		// readers
// 	wg.Add(5)
// 	for i:=0; i < 5; i++ {
// 		go func ()  {
// 			defer wg.Done()
// 			cfg := v.Load().(Config) // type assertion	
// 		}()
// 	}
// 	wg.Wait()
// }