package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("ERROR: Expected Length 52, but got: %v", len(d))
	}

	if d[0] != "Two of Clubs" {
		t.Errorf("ERROR: Expected 'Two of Clubs' but got: %v", d[0])
	}

	if d[51] != "Ace of Diamonds" {
		t.Errorf("ERROR: Expected 'Ace of Diamonds' but got: %v", d[51])
	}
}

func TestSaveDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("ERROR: Expected Length 52, but got: %v", len(loadedDeck))
	}
}
