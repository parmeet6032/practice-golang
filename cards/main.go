package main

import "fmt"

func main() {
	// declareVariables()
	// usingSlice()
	// usingTypeDeck()
	usingFile()
}

func declareVariables() {
	// DECLARING VARIABLES

	// (1)
	// var card string = "Ace of Spades"

	// (2)
	card := "Ace of Spades"
	fmt.Println(card)

	card = returnCard()
	fmt.Println(card)

	card = "Seven of Hearts"
	fmt.Println(card)

	// prints length of string
	bytesPrinted, err := fmt.Println(len(card))
	fmt.Println(bytesPrinted)
	fmt.Println(err)

	fmt.Println("\n" + "PRINTING CHARACTERS")
	for i, ch := range card {
		fmt.Println(i, "-->", ch, "-->", string(ch))
	}
}

func returnCard() string {
	return "Five of Diamonds"
}

func usingSlice() {
	cards := []string{"Ace of Spades", returnCard(), "Three of Clubs"}
	fmt.Println(cards)

	cards = append(cards, "Seven of Hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func usingTypeDeck() {
	cards := newDeck()

	fmt.Println("\n 1st WAY -------->")
	hand, remainingCards := deal(cards, 5)

	fmt.Println("HAND --> ")
	hand.print()

	fmt.Println("REMAINING --> ")
	cards = remainingCards
	cards.print()

	// fmt.Println("\n 2nd WAY (not working) -------->")
	// hand = cards.deal(4)

	// fmt.Println("HAND --> ")
	// hand.print()

	// fmt.Println("REMAINING --> ")
	// cards.print()
}

func usingFile() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
}
