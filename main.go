package main

import "fmt"

func main() {
	// cards := shuffle(newDeck())
	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println("Hand")
	// hand.print()

	// fmt.Println("remainingDeck")
	// remainingDeck.print()

	// cards := newDeck()
	cards := newDeckFromFile("myCards")

	fmt.Println(cards)

}
