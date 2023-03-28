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

func ShowCards(Deck Deck) {
	topDeck := len(Deck.DeckOfCards) - 1

	fmt.Printf("%s de %s", Deck.DeckOfCards[topDeck].Type, Deck.DeckOfCards[topDeck].Suit)

}

func removeCards(pDeck Deck, cDeck Deck) {
	topDeck := len(pDeck.DeckOfCards) - 1

	pDeck.DeckOfCards = append(pDeck.DeckOfCards[:topDeck], pDeck.DeckOfCards[topDeck+1])
	cDeck.DeckOfCards = append(cDeck.DeckOfCards[:topDeck], cDeck.DeckOfCards[topDeck+1])
}

func dealACard(pDeck Deck, cDeck Deck) {

	topDeck := len(pDeck.DeckOfCards) - 1

	if pDeck.DeckOfCards[topDeck].Value > cDeck.DeckOfCards[topDeck-1].Value {
		fmt.Println("Le Joueur gagne")
	} else if pDeck.DeckOfCards[topDeck].Value < cDeck.DeckOfCards[topDeck-1].Value {
		fmt.Println("Monsieur Roboto gagne")
	} else {
		fmt.Println("Égalité")
	}
	removeCards(pDeck, cDeck)
}

func deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d.DeckOfCards[i])
	}
}

func debug(d Deck) {
	if os.Getenv("DEBUG") != "" {
		for i := 0; i < len(d.DeckOfCards); i++ {
			fmt.Printf("Card #%d is a %s of %ss\n", i+1, d.DeckOfCards[i].Type, d.DeckOfCards[i].Suit)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	deckInstance = nil
}
