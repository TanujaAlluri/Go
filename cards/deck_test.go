package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "AceofSpades" {
		t.Errorf("Expected first card to be Ace of spades, but got %v", d[0])
	}

	if d[len(d)-1] != "FourofClubs" {
		t.Errorf("Expected last card to be Four of clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected loaded deck length of 16, but got %v", len(d))
	}

	os.Remove("_decktesting")
}
