package main

func main() {
	cards := deck{"Ace of spades", newCard()}
	newCards := append(cards, "Six of Spades")

	cards.print()

	newCards.print()
}

func newCard() string {
	return "Five of diamonds"
}
