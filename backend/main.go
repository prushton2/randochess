package main

import (
	"fmt"
	"net/http"

	"prushton.com/randochess/v2/board"
	"prushton.com/randochess/v2/game"
)

type RequestCode struct {
	Code string `json:"code"`
}

type RequestMove struct {
	Code  string `json:"code"`
	Start int    `json:"start_pos"`
	End   int    `json:"end_pos"`
}

type CodeInfo struct {
	GameIndex int
	Team      board.Team
}

var games map[int]game.Game
var codes map[string]CodeInfo

func fetch(w http.ResponseWriter, r *http.Request) {

}

func move(w http.ResponseWriter, r *http.Request) {

}

func main() {

	games = make(map[int]game.Game)
	codes = make(map[string]CodeInfo)

	games[0] = game.New8x8()
	codes["0"] = CodeInfo{GameIndex: 0, Team: board.White}
	codes["1"] = CodeInfo{GameIndex: 0, Team: board.Black}

	http.HandleFunc("/game/fetch", fetch)
	http.HandleFunc("/game/move", move)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
