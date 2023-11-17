// code to describe what a deck is & how it works
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// creating a new type of 'deck'
// which is a slice of string
// type deck ==== []string (equal)
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"clubs", "spades", "diamonds", "heart"}
	cardValues := []string{"ace", "one", "two", "three"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

// printing cards of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (dealing the cards) multiple return values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// coverting deck type to string to use Writerfile func
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saving data to the harddrive
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// readfile from hard drive
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //return binnaryString and error
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffling of deck
func (d deck) shuffle() {
	// every single time we start up our program, we're going to take the current time
	//which will be reflected as a value in 60 of type in 64.
	// And we'll use that as kind of the seed value for our random number generator.
	// So every single time we start up our program, this starting time is going to be different,
	// which means we should get a different sequence of random numbers.

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

	// OR below code is also correct

	// source := rand.NewSource(time.Now().UnixNano()) //seed generating, returns source
	// r := rand.New(source)                           //returns rand type
	// for index := range d {
	// 	newPosition := r.Intn(len(d) - 1) // Intn will return int
	// 	d[index], d[newPosition] = d[newPosition], d[index]

	// }

}
