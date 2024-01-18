package cli

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, strChan chan string, wg *sync.WaitGroup) {
	client := &http.Client{} // all explained at the bottom
	req, err := http.NewRequest("GET", translateUrl, nil)
	fmt.Println("req: ", req)
	if err != nil {	log.Fatalf("1.There was a problem: %s", err)}

	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t") // what dt stands for
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req) // make a req
	if err != nil {
		log.Fatalf("2.There was a problem: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		strChan <- "try again later, too many requests"
		wg.Done()
		return
	}

	parsedJson, err := gabs.ParseJSONBuffer(res.Body)
	fmt.Println("parsedJson: ", parsedJson)
	if err != nil {
		log.Fatalf("3.There was a problem: %s", err)
	}

	nestOne, err := parsedJson.ArrayElement(0)
	if err != nil {
		log.Fatalf("4.There was a problem: %s", err)
	}

	nestTwo, err := nestOne.ArrayElement(0)
	fmt.Println("nest two: ", nestTwo)
	if err != nil {
		log.Fatalf("5.There was a problem: %s", err)
	}

	translatedString, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatalf("6.There was a problem: %s", err)
	}
	
	strChan <- translatedString.Data().(string)
	wg.Done()
}

/* README

	client := &http.Client{} // {} means using 0 val init, Client{} is a built in type struct
	// &http creates a pointer to a new instance of the http.client

	!! when you declare a variable and assign a struct literal (like http.Client{}) to it, you are working with a copy of the struct,
	not the original struct itself. If you want to work with the original struct and potentially modify its properties, you use
	 a pointer to the struct.

	 query    := req.URL.Query() // gets the url query params: https://example.com/search?q=go&lang=en, the query parameters are q=go and
	lang=en.

parsedJson, err := gabs.ParseJSONBuffer(res.Body) is used to parse JSON data from an HTTP response body into a data structure using the
gabs package in Go. Let's break down what this line of code does:
gabs: gabs is a third-party Go package that provides a convenient way to work with JSON data. It allows you to easily parse,
manipulate, and access JSON objects and values.
ParseJSONBuffer(res.Body): This part of the code uses the gabs package's ParseJSONBuffer function to parse JSON data from the res.Body
 object. Here's what each part does:
res.Body: res is likely an http.Response object representing the HTTP response received from a server. res.Body is an io.ReadCloser
that provides access to the response body, which contains the JSON data.
ParseJSONBuffer: This function is called on the gabs package and takes an io.Reader as an argument, which is what res.Body is.
 It reads the JSON data from the provided io.Reader and parses it into a gabs.Container object.
*/

/*

func handler(w http.ResponseWriter, req *http.Request) {
	// Parse the query parameters from the request URL.
	queryParams := req.URL.Query()

	// Access specific query parameters by their keys.
	searchQuery := queryParams.Get("q")
	languages := queryParams["lang"] // Get all values associated with the "lang" key.

	// Print and display the parsed values.
	fmt.Printf("Search Query: %s\n", searchQuery)
	fmt.Printf("Languages: %v\n", languages)

	// Your application logic here.
}

func main() {
	http.HandleFunc("/search", handler)
	http.ListenAndServe(":8080", nil)
}
*/
