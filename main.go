package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/cards"
	"myapp/player"
	"net/http"
)

type GameData struct {
	Player1 *player.Player
	Player2 *player.Player
}

func main() {

	// Récupération du deck de cartes
	deck := cards.GetInstance()

	// Instanciation des joueurs
	myPlayer := player.NewPlayer("Player 1", deck.DeckOfCards[:26], 0)
	cpuPlayer := player.NewPlayer("Monsieur Roboto", deck.DeckOfCards[26:], 0)

	// Stockage des informations des joueurs
	gameData := GameData{Player1: &myPlayer, Player2: &cpuPlayer}

	// Gestionnaire de requêtes pour la page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Création d'un template pour la page de formulaire HTML
		tmpl := template.Must(template.ParseFiles("form.html"))

		tmpl.Execute(w, nil)
	})

	// Gestionnaire de requêtes pour la page de jeu
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {

		// Vérification si le nom du joueur a été fourni
		if r.FormValue("nom") != "" {
			myPlayer.Name = r.FormValue("nom")
		}

		// Création d'un template pour la page de jeu HTML
		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.Execute(w, gameData)
	})

	// Gestionnaire de requêtes pour le résultat de jeu
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {

		// Appel de la fonction PlayRound pour obtenir le résultat de jeu
		result := player.PlayRound(&myPlayer, &cpuPlayer)

		// Encodage du résultat de jeu en JSON
		out, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Println(err)
			return
		}

		// Définition du type de contenu de la réponse
		w.Header().Set("Content-Type", "application/json")

		// Écriture du résultat de jeu encodé en JSON dans la réponse
		w.Write(out)
	})

	http.ListenAndServe(":8080", nil)
}
