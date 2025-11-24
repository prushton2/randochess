package game

import (
	"fmt"
	"slices"
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

	game.Ruleset.InitBoard(&game.Board)

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

	// depending on where the piece is moving to, we get either the valid move position or valid take positioon
	var validMoveSpots []int

	if self.Board.Pieces[start].GetPieceTeam().OtherTeam() == self.Board.Pieces[end].GetPieceTeam() {
		// we are taking another piece, so check the places the piece can take
		_, validMoveSpots = rule(self.Board, start)
	} else {
		// else check where they can move
		validMoveSpots, _ = rule(self.Board, start)
	}

	if !slices.Contains(validMoveSpots[:], end) {
		// fmt.Printf("%v\n%d\n", validMoveSpots, end)
		return fmt.Errorf("Cannot move to specified spot")
	}

	if self.Board.Pieces[start].GetPieceTeam() == self.Board.Pieces[end].GetPieceTeam() {
		return fmt.Errorf("Cannot take own team's piece")
	}

	self.Turn = self.Ruleset.Move(&self.Board, start, end, self.Turn)

	self.Winner = self.Ruleset.GetWinner(self.Board)

	return nil
}
