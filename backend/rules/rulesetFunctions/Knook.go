package rulesetfunctions

import "prushton.com/randochess/v2/board"

func Knook(self board.Board, start int) ([]int, []int) {
	var knightMoves, knightTakes = DefaultKnight(self, start)
	var rookMoves, rookTakes = DefaultRook(self, start)

	return append(knightMoves, rookMoves...), append(knightTakes, rookTakes...)
}
