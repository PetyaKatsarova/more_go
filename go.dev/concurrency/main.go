package main

import (
	"fmt"
	"sync"
	"time"
)

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// 	"strings"
// )

// func handleConnection(conn net.Conn, done chan<- string) {
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn) // read from the connection
// 	for {
// 		msg, err := reader.ReadString('\n')
// 		if err != nil {
// 			done <- fmt.Sprintf("Client disconnected: %s", conn.RemoteAddr().String())
// 			return
// 		}
// 		fmt.Printf("Received: %s", msg)
// 		conn.Write([]byte(strings.ToUpper(msg)))
// 	}
// }

// func main() {
// 	listener, err := net.Listen("tcp", ":8080") // creates a server
// 	if err != nil {
// 		fmt.Println("Error starting the server:", err)
// 		os.Exit(1)
// 	}
// 	defer listener.Close()
// 	fmt.Println("Server started on port 8080")
// 	done := make(chan string)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}
// 		fmt.Printf("Accepted connection from %s", conn.RemoteAddr().String())
// 		go handleConnection(conn, done)
// 	}

// 	for msg := range done {
// 		fmt.Println(msg)
// 	}
// }

var (
	myResource	string
	rwMutex		sync.RWMutex
)

func readResource(id int) {
	rwMutex.RLock()
	fmt.Printf("Goroutine: %d: reading res val: %s\n", id, myResource)
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()
}

func writeResource(value string) {
	rwMutex.Lock()
	fmt.Println("writing to res")
	myResource = value
	time.Sleep(1 * time.Second)
	rwMutex.Unlock()
}



func main() {
	for i := 1; i <= 5; i++ {
		go readResource(i)
	}
	go writeResource("new val")
	time.Sleep(5 * time.Second)

	var wg sync.WaitGroup
	for _, salutiation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func (salutation string)  {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutiation)
	}
		wg.Wait()
}
