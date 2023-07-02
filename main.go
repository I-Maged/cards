package main

import "fmt"

func main() {
	cards := []string{"Five of Diamonds", newCard()}
	cards = append(cards, "Queen of Spades")

	for _, card := range cards {
		fmt.Println(card)
	}
}
func newCard() string {
	return "Ace of Spades"
}
