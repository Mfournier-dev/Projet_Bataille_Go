package cards

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Card struct {
	Suit  string
	Type  string
	Value int
}

type Deck struct {
	FullCards []Card
}

var deckInstance *Deck

func GetInstance() *Deck {
	if deckInstance == nil {
		deckInstance = &Deck{
			FullCards: shuffle(createNewDeck()),
		}
	}
	return deckInstance
}

func createNewDeck() (deck Deck) {

	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	suits := []string{"HEART", "DIAMOND", "CLUB", "SPADE"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type:  types[i],
				Suit:  suits[n],
				Value: i + 2,
			}
			deck.FullCards = append(deck.FullCards, card)
		}
	}
	return deck
}

func shuffle(d Deck) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d.FullCards); i++ {

		r := rand.Intn(i + 1)

		if i != r {
			d.FullCards[r], d.FullCards[i] = d.FullCards[i], d.FullCards[r]
		}
	}
	return d.FullCards
}

func Deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d.FullCards[i])
	}
}

func Debug(d Deck) {
	if os.Getenv("DEBUG") != "" {
		for i := 0; i < len(d.FullCards); i++ {
			fmt.Printf("Card #%d is a %s of %ss\n", i+1, d.FullCards[i].Type, d.FullCards[i].Suit)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	deckInstance = nil
}
