package rules

import (
	"math"

	"prushton.com/randochess/v2/board"
)

type Ruleset struct {
	Name       string `json:"name"`
	PieceRules map[board.PieceType]func(board.Board, uint8, uint8) bool
}

func DefaultRuleset() Ruleset {
	return Ruleset{
		Name:       "Default",
		PieceRules: make(map[board.PieceType]func(board.Board, uint8, uint8) bool),
	}
}

func CheckLineOfSight(board board.Board, start int, end int) bool {
	/*
		board is a large version of this:
		[0, 1, 2, 3
		 4, 5, 6, 7
		 8, 9,10,11
		12,13,14,15]
		X position is i%width
		y position is i/height
	*/

	var delta_x int = start%board.Width - end%board.Width
	var delta_y int = start/board.Height - end/board.Height

	// handle off cases like the horse where you dont move in a straight line, you are just allowed to move.
	if delta_x != 0 && delta_y != 0 {
		if math.Abs(float64(delta_x)) != math.Abs(float64(delta_y)) {
			return true
		}
	}

	return true
}
