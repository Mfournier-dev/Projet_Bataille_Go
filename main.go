package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/cards"
	"myapp/player"
	"myapp/rps"
	"net/http"
	"strconv"
)

type GameData struct {
	Player1 *player.Player
	Player2 *player.Player
}

func main() {

	deck := cards.GetInstance()

	myPlayer := player.NewPlayer("Player 1", deck.FullCards[:26])
	cpuPlayer := player.NewPlayer("CPU", deck.FullCards[26:])

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

	//partie DARA
	//http.HandleFunc("/play", playRound)
	//http.HandleFunc("/", homePage)

	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	//i have added here playerChoice variable
	//from object request, we getting variable "c" content
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
