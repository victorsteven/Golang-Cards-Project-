package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck to have a length of 16, but got %v", len(d))
		// print("Expected deck to have a length of 16, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to have four ofclubs, but git %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Exptected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
