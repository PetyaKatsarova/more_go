package main
//learning from https://github.com/instabledesign/go-channel/blob/master/04_receiverware_rate_limiter_test.go

import (
    "fmt"
    "time"
)

func main() {
    interval := 2 * time.Second //2 seconds
    ticker := time.NewTicker(interval)

    // listen for ticks from the ticker
    go func() {
        for {
            select {
            case <-ticker.C:
                // This code will run every 2 seconds
                fmt.Println("Tick!")
            }
        }
    }()

    // Keep the program running for a while
    time.Sleep(11 * time.Second)

    // Stop the ticker when done to release resources
    ticker.Stop()
}