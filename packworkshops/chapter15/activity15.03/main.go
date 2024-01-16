package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// https://github.com/PacktWorkshops/The-Go-Workshop/blob/master/Chapter15/Activity15.03/main.go
/*
netstat -ano | findstr :8080 to find curr processes
*/

type Visitor struct {
	Name string
}

type Hello struct {
	tmpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlQuery 	:= r.URL.Query()
	cust		:= Visitor{}
	name, ok 	:= urlQuery["name"]

	if ok {
		fmt.Println(name)
		cust.Name = strings.Join(name, ",")
	}
	h.tmpl.Execute(w, cust)
}

func NewHello(tmplPath string) (*Hello, error) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil { return nil, err}
	return &Hello{tmpl}, nil
}

func main() {
	hello, err := NewHello("./index.html")
	if err != nil { log.Fatal(err)}
	http.Handle("/", hello)
	log.Fatal(http.ListenAndServe(":8088", nil))
}