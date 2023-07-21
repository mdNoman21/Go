package main

import (
	"fmt"
	"math/rand"
	"strconv"
)
// extra functionality ,prompt for user affirmation and input,then only code can proceed
type Quote struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

type QuoteCollection struct {
	Quotes []Quote `json:"quotes"`
}

func (collection *QuoteCollection) GetQuote(id string) {
	for _, quote := range collection.Quotes {
		if quote.ID == id {
			fmt.Println(quote)
		}
	}
}
// func DisplayMenu() {
// 	fmt.Println("Need a quote for the day,press enter.")
	
// }
func main() {
	// DisplayMenu()
	collection := QuoteCollection{}
	collection.Quotes = []Quote{
		{
			ID:      "1",
			Message: "Quote 1 message",
			Author:  "Author 1",
		},
		{
			ID:      "2",
			Message: "Quote 2 message",
			Author:  "Author 2",
		},
		{
			ID:      "3",
			Message: "Quote 3 message",
			Author:  "Author 3",
		},
		{
			ID:      "4",
			Message: "Quote 4 message",
			Author:  "Author 4",
		},
		{
			ID:      "5",
			Message: "Quote 5 message",
			Author:  "Author 5",
		},
	}
	randomQuoteId := strconv.Itoa(rand.Intn(len(collection.Quotes)) + 1)
	collection.GetQuote(randomQuoteId)
	// fmt.Println("Press enter to generate a random quote.")
	// fmt.Scanln()

}
