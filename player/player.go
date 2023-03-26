package player

import "myapp/cards"

type Player struct {
	Name  string
	Score int
	Cards []cards.Card
}

func NewPlayer(name string, cards []cards.Card) Player {
	p := Player{Name: name, Cards: cards, Score: 0}
	return p
}
