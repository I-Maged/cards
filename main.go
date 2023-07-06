package main

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_cards")

	newCards := newDeckFromFile("my_cards")

	newCards.print()

	// hand, remainingCards := deal(cards, 4)

	// hand.print()
	// remainingCards.print()

	// fmt.Println("------------")
	// fmt.Print(hand.toString())
}
