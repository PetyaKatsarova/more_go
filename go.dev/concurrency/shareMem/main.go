package main

import (
	"log"
	"net/http"
	// "net/url"
	"time"
)

const (
	numPollers		= 2 // num of poll goroutines to launch
	pollInterval	= 60 * time.Second // how often to poll each URL
	statusInterval	= 10 * time.Second
	errTimeout		= 10 * time.Second 
)

var urls = []string {
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

type State struct {
	url		string
	status	string
}

type Resourse struct {
	url			string
	errCount	int
}


func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s.status
			}
		}
	}()
	return updates
}

func logState(s map[string]string) {
	log.Println("Current state:")
	for k, v := range s {
		log.Printf(" %s %s", k, v)
	}
}

func (r *Resourse) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
}

func (r *Resourse) Sleep(done chan<- *Resourse) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <- chan *Resourse, out chan <- *Resourse, status chan <- State) {
	for r := range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

func main() {
	pending, complete := make(chan *Resourse), make(chan *Resourse)

	status := StateMonitor(statusInterval)
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}

	// Send some Resources to the pending queue.
	go func() {
		for _, url := range urls {
			pending <- &Resourse{url: url}
		}
	}()

	for r := range complete {
		go r.Sleep(pending)
	}
}
