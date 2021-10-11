package main

import "fmt"

func main() {
	cards := []string{"Ace of spades", newCard()}
	newCards := append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	for i, newCard := range newCards {
		fmt.Println(i, newCard)
	}
}

func newCard() string {
	return "Five of diamonds"
}
