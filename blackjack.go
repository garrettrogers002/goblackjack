package main

import "fmt"


func createDeck() {
	var suits = [4]string{"Clubs", "Spades", "Hearts", "Diamonds"}
	var values = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := []map[string]interface{}{}



	// fmt.Println("Len of values: ", len(values))
	// fmt.Printf("%T\n", suits)
	// fmt.Println(len(suits))
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			// cardName := suits[i] + " of " + values[j]
			card := map[string]interface{}{
				"suit": suits[i],
				"value": values[j],
				"number": j + 1,

				"gameValue": "getCardValue(j + 1)",
			}
				
			cards = append(cards, card)
		}
	}
	fmt.Println(len(cards))
}

func main() {
	fmt.Println("blackjack")
	createDeck()
}
