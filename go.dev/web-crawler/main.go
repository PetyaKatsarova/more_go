package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// returns the body of URL and a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

type SaveVisited struct {
	mu 		sync.Mutex
	visited	map[string]bool
}

func (v *SaveVisited) Visit(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.visited[url] {
		return true
	}
	v.visited[url] = true
	return false
}

var visitedUrls = &SaveVisited{visited: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 || visitedUrls.Visit(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	var wg sync.WaitGroup
	for _, val := range urls {
		wg.Add(1)
		go func (val string)  {
			defer wg.Done()
			Crawl(val, depth - 1, fetcher)
		}(val)
	}
	wg.Wait()
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}
 type fakeFetcher map[string]*fakeResult

 type fakeResult struct {
	body string
	urls []string
 }

 func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
 }
 // fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}