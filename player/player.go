package player

import "myapp/cards"

type Player struct {
	Name  string
	Score int
	Cards []cards.Card
}

func NewPlayer(name string, halfCards []cards.Card) Player {
	p := Player{Name: name, Cards: halfCards, Score: 0}
	return p
}
