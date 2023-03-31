package player

import (
	"myapp/cards"
)

const (
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
	ALLROUNDS    = 26
)

// definit les éléments d'une manche du jeu
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

// Crée un nouveau joueur en lui attribuant un nom et la moitié du paquet de cartes
func NewPlayer(name string, halfCards []cards.Card, score int) Player {
	p := Player{Name: name, Cards: halfCards, Score: score}
	return p
}

// Récupère la valeur de la carte au dessus du paquet.
func (p Player) GetCard() cards.Card {
	topDeck := len(p.Cards) - 1

	return p.Cards[topDeck]
}

// Simule une manche du jeu
func PlayRound(humanPlayer *Player, cpuPlayer *Player) Round {

	var result Round
	roundResult := ""

	//récupération de la valeur de la carte jouée
	humanValue := humanPlayer.GetCard().Value
	computerValue := cpuPlayer.GetCard().Value

	//récupération de la carte jouée pour l'affichage
	humanChoice := humanPlayer.GetCard().Type + " de " + humanPlayer.GetCard().Suit
	computerChoice := cpuPlayer.GetCard().Type + " de " + cpuPlayer.GetCard().Suit

	//supprimer la carte jouée (pour les deux joueurs)
	humanPlayer.Cards = cards.RemovePlayedCard(humanPlayer.Cards)
	cpuPlayer.Cards = cards.RemovePlayedCard(cpuPlayer.Cards)

	//comparaison pour déterminer le vainqueur de la manche
	if humanValue > computerValue {

		roundResult = humanPlayer.Name + " remporte la manche !"
		humanPlayer.Score++

	} else if humanValue < computerValue {

		roundResult = "Monsieur Roboto remporte la manche !"
		cpuPlayer.Score++

	} else {

		roundResult = "Égalité !"
	}
	//Lorsque les paquets sont vides, on annonce le gagnant de la partie
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
	result.RoundNumber = ALLROUNDS - len(humanPlayer.Cards)

	return result
}
