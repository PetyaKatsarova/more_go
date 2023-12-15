package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Sample text in which to search
    text := "Hello, world! Welcome to Golang."

    // Define a regular expression pattern
    // For example, to find words starting with 'W'
    pattern := `\bW\w*`

    // Compile the regular expression
    re, err := regexp.Compile(pattern)
    if err != nil {
        fmt.Println("Error compiling regex:", err)
    }

    // Find the first match
    match := re.FindString(text)
    fmt.Println("First match:", match)

    // Find all matches
    allMatches := re.FindAllString(text, -1)
    fmt.Println("All matches:", allMatches)
}

