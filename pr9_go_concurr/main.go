package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// each book will have id
var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg 		:= &sync.WaitGroup{}
	mtx		:= &sync.RWMutex{}
	cacheCh	:= make(chan Book)
	dbCh	:= make(chan Book)

	for i := 0; i < 5; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2) // on each iteration add 2 wg

		// concur: check if book id in cache and if so send to ch.
		go func (id int, wg *sync.WaitGroup, mtx *sync.RWMutex, ch chan<- Book )  {
			if b, ok := queryCache(id, mtx); ok {
				ch <- b // pass the book to the channel
			}
			wg.Done()
		}(id, wg, mtx, cacheCh)

		// check fi book si in db and if so send through ch
		go func (id int, wg *sync.WaitGroup, mtx *sync.RWMutex, ch chan<- Book)  {
			if b, ok := queryDb(id, mtx); ok {
				mtx.Lock()
				cache[id] = b
				mtx.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, mtx, dbCh)

		// create go routien per query to handle resp
		go func (cacheCh, dbCh <- chan Book)  {
			select {
			case b := <- cacheCh:
				fmt.Println("Source: Cache")
				fmt.Println(b)
				<- dbCh // what do we do here?????
			case b := <- dbCh:
				fmt.Println("Source: Database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

// returns book by given id and bool if found
func queryCache(id int, mtx *sync.RWMutex)  (Book, bool) {
	mtx.RLock()
	b, ok := cache[id]
	mtx.RUnlock()
	return b, ok
}

func  queryDb(id int, mtx *sync.RWMutex) (Book, bool)  {
	// time.Sleep(100 * time.Millisecond)

	for _, b := range books { // the books slice of Book from book.go
		if b.ID == id {
			mtx.Lock()
			cache[id] = b
			mtx.Unlock()
			return b, true
		}
	}
	return Book{}, false
}