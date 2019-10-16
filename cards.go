package main

import "fmt"

func main(){
//	var card string = "Ace of Spades"
//	card := "Ace of Spades" //only used for new var
	//card := newCard()
//	card = "Five of Diamonds"
cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	
	for i, card := range cards {
	fmt.Println(i, card)
}
}

//we meed to specify what data type newCared will work as
func newCard() string{
	return "Five of Diamonds"
}
