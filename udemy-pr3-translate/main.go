package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"./cli"
	// "github.com/HotPotatoC/go-translate/cli"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]")
	flag.StringVar(&sourceLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceLang, "st", "en", "Text to translate")
}

func main() {
	fmt.Println("ehllo world")
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

	processedStr := strings.ReplaceAll(<-strChan, "+", " ") // replace all occurences of + with " " in the receivd from chan
	close(strChan)
	wg.Wait()
}
