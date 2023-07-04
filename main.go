package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 4)

	// hand.print()
	// remainingCards.print()

	// fmt.Println("------------")
	// fmt.Print(hand.toString())
}
