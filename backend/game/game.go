package game

import (
	"fmt"

	"prushton.com/randochess/v2/board"
	"prushton.com/randochess/v2/rules"
)

type Game struct {
	Ruleset rules.Ruleset `json:"ruleset"`
	Board   board.Board   `json:"board"`
	Turn    board.Team    `json:"turn"`
}

func New(rulesetName string) Game {

	ruleset := rules.SelectRuleset(rulesetName)

	game := Game{
		Board:   board.New(ruleset.Width, ruleset.Height),
		Ruleset: ruleset,
		Turn:    board.White,
	}

	game.Board.InitBoard()

	return game
}

func (self *Game) Move(start int, end int) error {
	if self.Board.Pieces[start].GetPieceTeam() != self.Turn {
		return fmt.Errorf("Incorrect Turn")
	}

	if start >= self.Board.Height*self.Board.Width || end >= self.Board.Height*self.Board.Width || start < 0 || end < 0 {
		return fmt.Errorf("Invalid start/end pos")
	}

	rule, exists := self.Ruleset.PieceRules[self.Board.Pieces[start].GetPieceType()]
	if !exists {
		return fmt.Errorf("No rule found for piece")
	}

	if !rule(self.Board, start, end) {
		return fmt.Errorf("Invalid move")
	}

	if self.Board.Pieces[start].GetPieceTeam() == self.Board.Pieces[end].GetPieceTeam() {
		return fmt.Errorf("Cannot take own team's piece")
	}

	// switch turn
	if self.Turn == board.White {
		self.Turn = board.Black
	} else {
		self.Turn = board.White
	}

	self.Board.Pieces[end] = self.Board.Pieces[start]
	self.Board.Pieces[start].SetPieceTeam(board.NoTeam)
	self.Board.Pieces[end].SetPieceMoved()

	return nil
}
