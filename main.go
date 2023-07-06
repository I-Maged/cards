package main

func main() {
	newCards := newDeckFromFile("my_cards")

	newCards.shuffleDeck()
	newCards.print()
}
