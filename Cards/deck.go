package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Squares"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var result deck
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			result = append(result, value+" of "+suit)
		}
	}
	return result
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//func deal(d deck, handSize int) (deck, deck) {
//	return d[:handSize], d[handSize:]
//}

func (d deck) toString() string {
	result := strings.Join(d, ",")
	return result
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swap position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
