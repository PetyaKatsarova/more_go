package main

import "fmt"

func main() {
	cards := newDeck()
	//cards.print()
	// print(newDeck())
	hand, remainingGards := deal(cards, 5)
	fmt.Println("------------------------hand cards----------------------------------- ")
	hand.print()
	fmt.Println("--------------------- remainining cards -------------------------")
	remainingGards.print()
}
