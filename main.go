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

	deck := cards.GetInstance()

	myPlayer := player.NewPlayer("Player 1", deck.DeckOfCards[:26], 0)
	cpuPlayer := player.NewPlayer("CPU", deck.DeckOfCards[26:], 0)

	gameData := GameData{Player1: &myPlayer, Player2: &cpuPlayer}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Création d'un template
		tmpl := template.Must(template.ParseFiles("form.html"))

		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("nom") != "" {
			myPlayer.Name = r.FormValue("nom")
		}

		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.Execute(w, gameData)
	})

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {

		result := player.PlayRound(&myPlayer, &cpuPlayer)

		out, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

		log.Println(len(myPlayer.Cards))
	})

	http.ListenAndServe(":8080", nil)
}
