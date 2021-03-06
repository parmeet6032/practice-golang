package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) deal(handSize int) deck {
	hand := d[:handSize]
	d = d[handSize:]
	return hand
}

// converts deck --> []string --> string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saves the deck to file
// converts string --> []byte
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// reads deck from the saved file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

// shuffle deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// // in this, the default seed is provided always, so same random sequence generated always
		// newPosition := rand.Intn(len(d))

		// in this, custom source is given, which is always "different" as it depends on time
		newPosition := r.Intn(len(d))

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
