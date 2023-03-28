package player

import (
	"math/rand"
	"myapp/cards"
	"time"
)

const (
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	HumanChoice    string `json:"human_choice"`
	HumanScore     int    `json:"human_score"`
	ComputerChoice string `json:"computer_choice"`
	ComputerScore  int    `json:"computer_score"`
	RoundResult    string `json:"round_result"`
}

type Player struct {
	Name  string
	Score int
	Cards []cards.Card
}

func NewPlayer(name string, halfCards []cards.Card) Player {
	p := Player{Name: name, Cards: halfCards, Score: 0}
	return p
}

func (p Player) ShowCards() cards.Card {
	topDeck := len(p.Cards) - 1

	return p.Cards[topDeck]
}

func PlayRound(humanPlayer Player, cpuPlayer Player) Round {
	rand.Seed(time.Now().UnixNano())

	humanValue := humanPlayer.ShowCards().Value
	computerValue := cpuPlayer.ShowCards().Value

	humanChoice := humanPlayer.ShowCards().Type + " de " + humanPlayer.ShowCards().Suit
	computerChoice := cpuPlayer.ShowCards().Type + " de " + cpuPlayer.ShowCards().Suit

	roundResult := ""
	winner := 0

	if humanValue > computerValue {
		roundResult = humanPlayer.Name + " wins !"
		winner = PLAYERWINS
		humanPlayer.Score++

	} else if humanValue < computerValue {
		roundResult = "Computer wins !"
		winner = COMPUTERWINS
		cpuPlayer.Score++
	} else {
		roundResult = "It's a draw !"
		winner = DRAW
	}

	var result Round
	result.Winner = winner
	result.ComputerScore = cpuPlayer.Score
	result.HumanScore = humanPlayer.Score
	result.ComputerChoice = computerChoice
	result.HumanChoice = humanChoice
	result.RoundResult = roundResult
	return result
}
