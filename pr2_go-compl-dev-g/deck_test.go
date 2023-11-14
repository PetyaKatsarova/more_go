package main

import (
	"fmt"
	"os"
	"testing"
)

// func will be run with *testing.T: t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("\nexpected deck len of 16 but got %v", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("\nExpected first card 'ace of spades' but got: %v", d[0])
	}

	if d[len(d) - 1] != "4 of clubs" {
		t.Errorf("\nExpected 'last 4 of clubs' but got: %v", d[len(d) - 1])
	}
}
 
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesing")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	// c is a byte slice from the readfile
	c , _ := os.ReadFile("_decktesting")
	fmt.Printf("%+s\n", string(c))
	loadedDeck := newDEckFromFile("_decktesting")
	// make sure it fails at least once
	if len(loadedDeck) != 160 {
		t.Errorf("\nexpected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
