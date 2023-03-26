package cards

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Card struct {
	Type  string
	Suit  string
	Value int
}

type Deck struct {
	cards []Card
}

var deckInstance *Deck

func GetInstance() *Deck {
	if deckInstance == nil {
		deckInstance = &Deck{
			cards: shuffle(createNewDeck()),
		}
	}
	return deckInstance
}

func createNewDeck() (deck Deck) {

	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type:  types[i],
				Suit:  suits[n],
				Value: i,
			}
			deck.cards = append(deck.cards, card)
		}
	}
	return deck
}

func shuffle(d Deck) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d.cards); i++ {

		r := rand.Intn(i + 1)

		if i != r {
			d.cards[r], d.cards[i] = d.cards[i], d.cards[r]
		}
	}
	return d.cards
}

func Deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d.cards[i])
	}
}

func Debug(d Deck) {
	if os.Getenv("DEBUG") != "" {
		for i := 0; i < len(d.cards); i++ {
			fmt.Printf("Card #%d is a %s of %ss\n", i+1, d.cards[i].Type, d.cards[i].Suit)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	deckInstance = nil
}
