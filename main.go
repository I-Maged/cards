package main

func main() {
	d := newDeck()
	d.saveToFile("_newcards")
	newCards := newDeckFromFile("_newcards")

	newCards.shuffleDeck()
	newCards.print()
}
