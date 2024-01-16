package main

import (
	"fmt"
	"log"
	"net/http"
)

//https://github.com/PacktWorkshops/The-Go-Workshop/blob/master/Chapter15/Activity15.01/main.go
type PageWithCounter struct{
	counter int
	heading string
	content string
}

func(h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.counter++
	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s</p>\n<p>Views: %d</p>", h.heading, h.content, h.counter)
	w.Write([]byte(msg))
}

func main() {
	hello := PageWithCounter{heading: "Hello World", content: "This is the main page"}
	ch1	  := PageWithCounter{heading: "Chapter 1", content: "This is the first chapter"}
	ch2	  := PageWithCounter{heading: "Chapter 2", content: "This is the second chapter"}

	http.Handle("/", &hello)
	http.Handle("/chapter1", &ch1)
	http.Handle("/chapter2", &ch2)

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("listening on port 8080")
}
/*
You have defined a type PageWithCounter, and you've attached the ServeHTTP method to it. By doing this, you've made PageWithCounter 
satisfy the http.Handler interface.

to run:
1. go run .
2. in browser: localhost:8080/
*/
