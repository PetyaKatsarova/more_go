package main

// // gin: HTTP web framework written in go, install s package; provides a framework for
// // building web applications and services with a minimalistic setup
// // go get -u github.com/gin-gonic/gin

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // album represents data about a record album.
// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// // albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 			return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 			if a.ID == id {
// 					c.IndentedJSON(http.StatusOK, a)
// 					return
// 			}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/albums", getAlbums)
// 	router.GET("/albums/:id", getAlbumByID)
// 	router.POST("/albums", postAlbums)

// 	router.Run("localhost:8080")
// }

// // go run main.go
// // curl http://localhost:8080/albums

// /*
// curl http://localhost:8080/albums \
//     --include --header \
//     "Content-Type: application/json" \
//     --request "POST" --data \
//     '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
// */

// package main

// import (
// 	"time"
// 	"math/rand"
// 	"fmt"
// )

// func longTimeRequest() <-chan int32 {
// 	r := make(chan int32)

// 	go func() {
// 		// Simulate a workload.
// 		time.Sleep(time.Second * 3)
// 		r <- rand.Int31n(100) // receive only chan
// 	}()

// 	return r
// }

// func longTimeRequest2(r chan <- int32) {
// 	time.Sleep(time.Second * 3)
// 	r <- rand.Int31n(100)
// }

// func sumSquares(a, b int32) int32 {
// 	return a*a + b*b
// }

// func main() {
// 	a, b := longTimeRequest(), longTimeRequest()
// 	fmt.Println(sumSquares(<-a, <-b))

// 	results := make(chan int32, 2)
// 	go longTimeRequest2(results)
// 	go longTimeRequest2(results)
// 	fmt.Println(sumSquares(<-results, <- results))
// }
/*
In the following example, the values of two arguments of the sumSquares
 function call are requested concurrently. Each of the two channel
  receive operations will block until a send operation performs on the corresponding channel. It takes about three seconds instead of six seconds to return the final result.
*/

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"time"
// )

// func main() {
// 	values := make([]byte, 32*1024*1024) // many 0s
// 	if _, err := rand.Read(values); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	done := make(chan struct{}) // can be buffered or not

// 	// The sorting goroutine
// 	go func() {
// 		sort.Slice(values, func(i, j int) bool {
// 			return values[i] < values[j]
// 		})
// 		// Notify sorting is done.
// 		done <- struct{}{}
// 	}()

// 	// do some other things ...

// 	<-done // waiting here for notification
// 	fmt.Println(values[0], values[len(values)-1])

// 	done2 := make(chan struct{})

// 	// go func is like async in js
// 	go func() {
// 		fmt.Print("Hello")
// 		// Simulate a workload.
// 		time.Sleep(time.Second * 2)
// 		<-done2
// 	}()

// 	// Blocked here, wait for a notification.
// 	done2 <- struct{}{}
// 	fmt.Println(" world!")
// }
/*
ready <-chan T:

This is a receive-only channel of type T.
The <-chan syntax specifies that this channel can only be used to receive values.
You cannot send values to this channel; attempting to do so will result in a compile-time error.
This type of channel is typically used in function parameters to ensure that the function can only read from the channel and not write to it, thereby enforcing a specific direction of data flow and enhancing code safety.
Example Usage: Reading a value from the channel: value := <-ready.

done chan<- T:

This is a send-only channel of type T.
The chan<- syntax specifies that this channel can only be used to send values.
You cannot receive values from this channel; doing so will also result in a compile-time error.
Like the receive-only channel, this type is used in function parameters to ensure that the function can only write to the channel and not read from it.
Example Usage: Sending a value to the channel: done <- someValue.
*/

import "log"
import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait a notification
	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// Simulate an initialization phase.
	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications.
	ready <- T{}; ready <- T{}; ready <- T{}
	// Being N-to-1 notified.
	<-done; <-done; <-done
}

