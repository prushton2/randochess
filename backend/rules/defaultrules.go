package rules

import "prushton.com/randochess/v2/board"

func DefaultPawn(board.Board, uint8, uint8) bool {

	return true
}

func DefaultRook(board.Board, uint8, uint8) bool {

	return true
}

func DefaultKnight(board.Board, uint8, uint8) bool {

	return true
}
func DefaultBishop(board.Board, uint8, uint8) bool {

	return true
}

func DefaultKing(board.Board, uint8, uint8) bool {

	return true
}

func DefaultQueen(board board.Board, start uint8, end uint8) bool {
	return CheckLineOfSight(board, int(start), int(end))
}
