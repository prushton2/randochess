package game

import (
	"prushton.com/randochess/v2/board"
	"prushton.com/randochess/v2/rules"
)

type Game struct {
	Ruleset rules.Ruleset
	Board   board.Board
	Turn    board.Team
}

func New8x8() Game {
	game := Game{
		Board:   board.New(8, 8),
		Ruleset: rules.DefaultRuleset(),
		Turn:    board.White,
	}

	game.Board.Init8x8Board()

	return game
}
