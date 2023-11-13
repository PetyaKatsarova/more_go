package main

import "fmt"

type deck []string

func newDeck() {

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

