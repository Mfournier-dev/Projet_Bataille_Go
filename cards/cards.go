package cards

import (
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Type  string
	Value int
}

type Deck struct {
	DeckOfCards []Card
}

var deckInstance *Deck

func GetInstance() *Deck {
	if deckInstance == nil {
		deckInstance = &Deck{
			DeckOfCards: shuffle(createNewDeck()),
		}
	}
	return deckInstance
}

func createNewDeck() (deck Deck) {

	types := []string{"Deux", "Trois", "Quatre", "Cinq", "Six", "Sept",
		"Huit", "Neuf", "Dix", "Valet", "Dame", "Roi", "As"}

	suits := []string{"Coeur", "Carreau", "Trèfle", "Pique"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type:  types[i],
				Suit:  suits[n],
				Value: i + 2,
			}
			deck.DeckOfCards = append(deck.DeckOfCards, card)
		}
	}
	return deck
}

func shuffle(d Deck) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d.DeckOfCards); i++ {

		r := rand.Intn(i + 1)

		if i != r {
			d.DeckOfCards[r], d.DeckOfCards[i] = d.DeckOfCards[i], d.DeckOfCards[r]
		}
	}
	return d.DeckOfCards
}

func RemovePlayedCard(arr []Card) []Card {
	topDeck := len(arr) - 1

	newArray := make([]Card, topDeck)
	copy(newArray, arr[:topDeck])

	return newArray
}

func init() {
	deckInstance = nil
}
