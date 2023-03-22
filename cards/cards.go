package cards

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Card struct {
	Type string
	Suit string
}

type Deck []Card

func createNewDeck() (deck Deck) {

	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return
}

func Shuffle(d Deck) Deck {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d); i++ {

		r := rand.Intn(i + 1)

		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
	return d
}

func Deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d[i])
	}
}

func Debug(d Deck) {
	if os.Getenv("DEBUG") != "" {
		for i := 0; i < len(d); i++ {
			fmt.Printf("Card #%d is a %s of %ss\n", i+1, d[i].Type, d[i].Suit)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
