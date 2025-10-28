package rules

import "prushton.com/randochess/v2/board"

var allRulesets map[string]Ruleset = map[string]Ruleset{
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
	"Oops! All Knights!": {
		Name: "Oops! All Knights!",
		PieceRules: map[board.PieceType]func(board.Board, int, int) bool{
			board.Pawn:   DefaultKnight,
			board.Rook:   DefaultKnight,
			board.Knight: DefaultKnight,
			board.Bishop: DefaultKnight,
			board.King:   DefaultKnight,
			board.Queen:  DefaultKnight,
		},
		Width:  8,
		Height: 8,
	},
	"PREPARE THYSELF": {
		Name: "PREPARE THYSELF",
		PieceRules: map[board.PieceType]func(board.Board, int, int) bool{
			board.Pawn:   PrepareThyselfPawn,
			board.Rook:   DefaultRook,
			board.Knight: DefaultKnight,
			board.Bishop: DefaultBishop,
			board.King:   DefaultKing,
			board.Queen:  DefaultQueen,
		},
		Width:  8,
		Height: 8,
	},
	"Have a plan to kill everyone you meet": {
		Name: "Have a plan to kill everyone you meet",
		PieceRules: map[board.PieceType]func(board.Board, int, int) bool{
			board.Pawn:   DefaultPawn,
			board.Rook:   DefaultRook,
			board.Knight: DefaultKnight,
			board.Bishop: PlanBishop,
			board.King:   DefaultKing,
			board.Queen:  DefaultQueen,
		},
		Width:  8,
		Height: 8,
	},
}
