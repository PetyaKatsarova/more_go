package main

import "fmt"

type deck []string

// func newDeck() []string{
	func  newDeck() deck {
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
