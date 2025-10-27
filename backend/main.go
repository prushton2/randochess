package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"prushton.com/randochess/v2/board"
	"prushton.com/randochess/v2/game"
)

type CodeInfo struct {
	GameIndex int
	Team      board.Team
}

var games map[int]game.Game
var codes map[string]CodeInfo

type RequestCode struct {
	Code string `json:"code"`
}

type RequestMove struct {
	Code  string `json:"code"`
	Start int    `json:"start_pos"`
	End   int    `json:"end_pos"`
}

type ResponseFetch struct {
	Team board.Team `json:"team"`
	Game game.Game  `json:"game"`
}

type RequestNew struct {
	RuleName string `json:"ruleName"`
}

func new(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body (is this a post request?)", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	var parsedBody RequestNew
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		http.Error(w, "Body is not valid JSON", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

}

func fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body (is this a post request?)", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	var parsedBody RequestCode
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		http.Error(w, "Body is not valid JSON", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	playerInfo, exists := codes[parsedBody.Code]
	if !exists {
		http.Error(w, "Code isnt valid", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	gameInfo, exists := games[playerInfo.GameIndex]
	if !exists {
		http.Error(w, "Player points to invalid game", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	var response = ResponseFetch{
		Team: playerInfo.Team,
		Game: gameInfo,
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to serialize JSON", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	// fmt.Printf("%s\n", data)

	io.Writer.Write(w, data)
}

func move(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body (is this a post request?)", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	var parsedBody RequestMove
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		http.Error(w, "Body is not valid JSON", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	playerInfo, exists := codes[parsedBody.Code]
	if !exists {
		http.Error(w, "Code isnt valid", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	gameInfo, exists := games[playerInfo.GameIndex]
	if !exists {
		http.Error(w, "Player points to invalid game", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	err = gameInfo.Move(parsedBody.Start, parsedBody.End)
	if err != nil {
		http.Error(w, "Invalid Move", http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}

	games[playerInfo.GameIndex] = gameInfo

	io.WriteString(w, "{\"status\": \"success\"}")
}

func main() {

	games = make(map[int]game.Game)
	codes = make(map[string]CodeInfo)

	games[0] = game.New("PREPARE THYSELF")
	codes["0"] = CodeInfo{GameIndex: 0, Team: board.White}
	codes["1"] = CodeInfo{GameIndex: 0, Team: board.Black}

	http.HandleFunc("/game/new", new)
	http.HandleFunc("/game/fetch", fetch)
	http.HandleFunc("/game/move", move)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
