package main

import "fmt"

func main() {
	// card := "Aces of Spades"
	// card := newCard()
	// fmt.Println(card)
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Diamond"
}
