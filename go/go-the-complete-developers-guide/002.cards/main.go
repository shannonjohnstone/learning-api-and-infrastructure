package main

import "fmt"

func main() {
	cards := []string{"Ace if Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println("Card:", i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
f