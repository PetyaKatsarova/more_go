package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// func newDeck() []string{
func newDeck() deck {
	cards := deck{}
	// cards := make([]string, 0) // allocate memory to the new slice
	// var cards []string; // v2: without custom var: without type deck []string

	cardsSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardsValues := []string{"ace", "2", "3", "4"}

	for _, suit := range cardsValues {
		for _, val := range cardsSuits {
			cards = append(cards, suit+" of "+val)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func print(cards []string) {
// 	for i, card := range cards {
// 				fmt.Println(i, card)
// 			}
// }

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ") // join all els from the slice of strings separated by ,
}

func (d deck) saveToFile(filename string) error { // error is an actual type
	return os.WriteFile(filename,[]byte(d.toString()), 0666) // for unix: 3 numbers: first is for owner, second for group, third: all other users
	// 0 no permission, 1 execute permission, 2 write p, 4 read p: 4 + 2 read and write; 4 + 2 + 1: read, write and execute, 4 + 1 read and execute, 4+2: read and write
}

func newDEckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) // bs is byte slice
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // os.Exit() in pkg.go.dev : Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, 
		// non-zero an error. The program terminates immediately; deferred functions are not run.
	}
	s := strings.Split(string(bs), ", ")
	// turn slice of str into a deck
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // u cant do this in C#
	}
}
