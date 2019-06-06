package main

import "testing"

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
