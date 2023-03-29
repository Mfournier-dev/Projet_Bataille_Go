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
	RoundNumber    int    `json:"round_number"`
}

type Player struct {
	Name  string
	Score int
	Cards []cards.Card
}

func NewPlayer(name string, halfCards []cards.Card, score int) Player {
	p := Player{Name: name, Cards: halfCards, Score: score}
	return p
}

func (p Player) GetCard() cards.Card {
	topDeck := len(p.Cards) - 1

	return p.Cards[topDeck]
}

func PlayRound(humanPlayer *Player, cpuPlayer *Player) Round {
	rand.Seed(time.Now().UnixNano())

	//récupération de la valeur de la carte jouée
	humanValue := humanPlayer.GetCard().Value
	computerValue := cpuPlayer.GetCard().Value

	//récupération de la carte jouée pour l'affichage
	humanChoice := humanPlayer.GetCard().Type + " de " + humanPlayer.GetCard().Suit
	computerChoice := cpuPlayer.GetCard().Type + " de " + cpuPlayer.GetCard().Suit

	//supprimer la carte jouée (pour les deux joueurs)
	humanPlayer.Cards = cards.RemovePlayedCard(humanPlayer.Cards)
	cpuPlayer.Cards = cards.RemovePlayedCard(cpuPlayer.Cards)

	allRounds := 26
	roundResult := ""

	//comparaison pour déterminer le vainqueur de la manche
	if humanValue > computerValue {

		roundResult = humanPlayer.Name + " wins !"
		humanPlayer.Score++

	} else if humanValue < computerValue {

		roundResult = "Computer wins !"
		cpuPlayer.Score++

	} else {

		roundResult = "It's a draw !"
	}

	var result Round

	if len(humanPlayer.Cards) == 0 && len(cpuPlayer.Cards) == 0 {

		if humanPlayer.Score > cpuPlayer.Score {
			result.Winner = PLAYERWINS
		} else if humanPlayer.Score < cpuPlayer.Score {
			result.Winner = COMPUTERWINS
		} else {
			result.Winner = DRAW
		}
	}

	result.ComputerScore = cpuPlayer.Score
	result.HumanScore = humanPlayer.Score
	result.ComputerChoice = computerChoice
	result.HumanChoice = humanChoice
	result.RoundResult = roundResult
	result.RoundNumber = allRounds - len(humanPlayer.Cards)
	return result
}
