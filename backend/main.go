package main

import (
	"fmt"
	"net/http"

	"prushton.com/randochess/v2/game"
)

func handler(w http.ResponseWriter, r *http.Request) {
	game := game.New8x8()
	fmt.Printf("%s\n", game.Ruleset.Name)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
