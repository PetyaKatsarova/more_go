package main

import (
	"flag"
	"fmt"
	"os"
	// "strings"
	"sync"

	"github.com/PetyaKatsarova/more_go/udemy-pr3-translate/cli"
)

// go run .\main.go -s en -st hello -t fr
//  go run .\main.go -s en -st "hello world" -t bg

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]") // flag is used for parsing command-line arguments provided when running the program
	flag.StringVar(&targetLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "en", "Text to translate")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("Options: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	strChan := make(chan string)
	wg.Add(1)

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	go cli.RequestTranslate(reqBody, strChan, &wg) // publish in the channel

	// processedStr := strings.ReplaceAll(<-strChan, "+", " ") // replace all occurences of + with " " in the receivd from chan
	// fmt.Println("processedStr: ", processedStr)
	fmt.Println(<-strChan)
	close(strChan)
	wg.Wait()
}
