package rulesetfunctions

import "prushton.com/randochess/v2/board"

func Knook(self board.Board, start int, end int) ([]int, []int) {
	var knightMoves, knightTakes = DefaultKnight(self, start, end)
	var rookMoves, rookTakes = DefaultRook(self, start, end)

	return append(knightMoves, rookMoves...), append(knightTakes, rookTakes...)
}
