package main

import "fmt"


func createDeck() {
	var suits = [4]string{"Clubs", "Spades", "Hearts", "Diamonds"}
	var values = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := []string{}
	// fmt.Println("Len of values: ", len(values))
	// fmt.Printf("%T\n", suits)
	// fmt.Println(len(suits))
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			cardName := suits[i] + " of " + values[j]
			cards = append(cards, cardName)
		}
	}
	fmt.Println(len(cards))
}

func main() {
	fmt.Println("blackjack")
	createDeck()
}
