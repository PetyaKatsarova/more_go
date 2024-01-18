package main

import (
	"fmt"
	"time"
	"io"
	"log"
	"net"
)

//// Clock1 is a TCP server that periodically writes the time.
// to start the program: in windows: tick in settings/programs/telnet , run the program, in another terminal write
// telnet localhost 8000
func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil { log.Fatal(err) }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err) // e.g., connection aborted
            continue
        }
        fmt.Println("server listens on port 8000")
        handleConn(conn) // handle one connection at a time
    }
}

func handleConn(c net.Conn) {
    defer c.Close()
    for {
        _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
        if err != nil {
            return // e.g., client disconnected
        }
        time.Sleep(1 * time.Second)
    }
}

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("worker", id, "started job", j) // only send to jobs
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finished job", j)
// 		results <- j * 2 // only receive in results
// 	}
// }

// func main() {
// 	const numjobs = 5
// 	jobs := make(chan int, numjobs)
// 	results := make(chan int, numjobs)

// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}

// 	for j := 1; j <= numjobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for a := 1; a <= numjobs; a++ {
// 		fmt.Printf("results[%d] %d\n" , a, <-results)
// 	}
// }
// func main() {
//     go spinner(100 * time.Millisecond)
//     const n = 45
//     fibN := fib(n) // slow
//     fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// }

// func spinner(delay time.Duration) {
//     for {
//         for _, r := range `-\|/` {
//             fmt.Printf("\r%c", r)
//             time.Sleep(delay)
//         }
//     }
// }

// func fib(x int) int {
//     if x < 2 {
//         return x
//     }
//     return fib(x-1) + fib(x-2)
// }
