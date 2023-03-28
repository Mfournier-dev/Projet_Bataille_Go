package player

import (
	"myapp/cards"
)

type Player struct {
	Name  string
	Score int
	Cards []cards.Card
}

func NewPlayer(name string, halfCards []cards.Card) Player {
	p := Player{Name: name, Cards: halfCards, Score: 0}
	return p
}

func (p Player) ShowCards() string {
	topDeck := len(p.Cards) - 1

	return p.Cards[topDeck].Type + " de " + p.Cards[topDeck].Suit
	//return (deck.DeckOfCards[topDeck].Type + deck.DeckOfCards[topDeck].Suit)

}
