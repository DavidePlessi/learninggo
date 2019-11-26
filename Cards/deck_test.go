package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card equal to Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Squares" {
		t.Errorf("Expected first card equal to King of Squares, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_deckTesting.txt"

	_ = os.Remove(fileName)
	d := newDeck()
	_ = d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected deck length of %v, but got %v", len(d), len(loadedDeck))
	}

	_ = os.Remove(fileName)
}
