package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	hand, remainingCard := deal(cards, 3)
	hand.print()
	fmt.Println("")
	remainingCard.print()

	fmt.Println(cards.toString())

	cards.saveToFile("my_card")

	freshDeck := newDeckFromFile("my_card")
	freshDeck.print()

	cards.shuffle()
	cards.print()

}
