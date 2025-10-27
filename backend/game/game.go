package game

import (
	"fmt"
	"time"

	"prushton.com/randochess/v2/board"
	"prushton.com/randochess/v2/rules"
)

type Game struct {
	Ruleset         rules.Ruleset `json:"ruleset"`
	Board           board.Board   `json:"board"`
	Turn            board.Team    `json:"turn"`
	Winner          board.Team    `json:"winner"`
	LastRequestedAt int64         `json:"lastRequestedAt"`
}

func New(rulesetName string) (Game, error) {

	ruleset, err := rules.SelectRuleset(rulesetName)
	if err != nil {
		return Game{}, err
	}

	game := Game{
		Board:           board.New(ruleset.Width, ruleset.Height),
		Ruleset:         ruleset,
		Turn:            board.White,
		Winner:          board.NoTeam,
		LastRequestedAt: time.Now().Unix(),
	}

	game.Board.InitBoard()

	return game, nil
}

func (self *Game) Move(start int, end int) error {
	self.LastRequestedAt = time.Now().Unix()

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
	self.Turn = self.Turn.OtherTeam()

	if self.Board.Pieces[end].GetPieceType() == board.King {
		self.Winner = self.Board.Pieces[end].GetPieceTeam().OtherTeam()
	}

	self.Board.Pieces[end] = self.Board.Pieces[start]
	self.Board.Pieces[start].SetPieceTeam(board.NoTeam)
	self.Board.Pieces[end].SetPieceMoved()

	return nil
}
