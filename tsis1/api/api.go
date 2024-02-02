package api

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type FCB_Players struct {
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

var players = []FCB_Players{
	{Name: "Manuel Neuer", Number: 1, Age: 37, Position: "Goalkeeper"},
	{Name: "Matthjis  De Ligt", Number: 2, Age: 24, Position: "Center-Back"},
	{Name: "Alphonso Davies", Number: 19, Age: 23, Position: "Left-Back"},
	{Name: "Joshua Kimmich", Number: 6, Age: 28, Position: "Defensive Midlefield"},
	{Name: "Jamal Musiala", Number: 42, Age: 20, Position: "Attacking Midfield"},
	{Name: "Leroy Sane", Number: 10, Age: 28, Position: "Right Winger"},
	{Name: "Harry Kane", Number: 9, Age: 30, Position: "Centre-Forward"},
	{Name: "Thomas Muller", Number: 25, Age: 34, Position: "Second Striker"},
	{Name: "Serge Gnabry", Number: 7, Age: 28, Position: "Right Winger"},
}

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	playernumber := params["number"]

	number, err := strconv.Atoi(playernumber)
	if err != nil {
		http.Error(w, "Invalid player number", http.StatusBadRequest)
		return
	}
	for _, player := range players {
		if player.Number == number {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.NotFound(w, r)
}

func young_players(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sort.Slice(players, func(i, j int) bool { return players[i].Age < players[j].Age })
	youngPlayers := players[:5]
	json.NewEncoder(w).Encode(youngPlayers)
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/players", ListPlayers).Methods("GET")
	r.HandleFunc("/players/{number}", GetPlayer).Methods("GET")
	r.HandleFunc("/youngPlayers", young_players).Methods("GET")
}
