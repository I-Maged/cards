package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	cards := newDeckFromFile("_decktesting")

	if len(cards) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}

	os.Remove("_decktesting")
}
