package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length to be 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be - 'Ace of Spades' but got - %v ", d[0])
	}

	if d[len(d)-1] != "Queen of Clubs" {
		t.Errorf("Expected last card to be - 'Queen of Clubs' but got - %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_deckTesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of loaded deck to be 52 but found to be %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
