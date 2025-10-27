package main

import (
	"fmt"
	"net/http"

	"prushton.com/randochess/v2/board"
)

func handler(w http.ResponseWriter, r *http.Request) {
	board := board.New(8, 8)
	board.Init8x8Board()

}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
