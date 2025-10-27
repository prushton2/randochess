package rules

import "prushton.com/randochess/v2/board"

var allRulesets map[string]Ruleset = map[string]Ruleset{
	"Open World": {
		Name: "Open World",
		PieceRules: map[board.PieceType]func(board.Board, int, int) bool{
			board.Pawn:   DefaultPawn,
			board.Rook:   DefaultRook,
			board.Knight: DefaultKnight,
			board.Bishop: DefaultBishop,
			board.King:   DefaultKing,
			board.Queen:  DefaultQueen,
		},
		Width:  16,
		Height: 16,
	},
	"Default": {
		Name: "Default",
		PieceRules: map[board.PieceType]func(board.Board, int, int) bool{
			board.Pawn:   DefaultPawn,
			board.Rook:   DefaultRook,
			board.Knight: DefaultKnight,
			board.Bishop: DefaultBishop,
			board.King:   DefaultKing,
			board.Queen:  DefaultQueen,
		},
		Width:  8,
		Height: 8,
	},
}
