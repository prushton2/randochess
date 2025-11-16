package rulesetfunctions

import (
	"fmt"

	"prushton.com/randochess/v2/board"
)

func DefaultMove(self *board.Board, start int, end int, team board.Team) board.Team {
	self.Pieces[end] = self.Pieces[start]
	self.Pieces[start].SetPieceTeam(board.NoTeam)
	self.Pieces[end].SetPieceMoved()
	return team.OtherTeam()
}

func DefaultGetWinner(self board.Board) board.Team {
	WhiteInPlay := false
	BlackInPlay := false

	for i := range self.Width * self.Height {
		if self.Pieces[i].GetPieceType() == board.King {
			if self.Pieces[i].GetPieceTeam() == board.White {
				WhiteInPlay = true
			}
			if self.Pieces[i].GetPieceTeam() == board.Black {
				BlackInPlay = true
			}
		}
	}

	if WhiteInPlay != BlackInPlay && (WhiteInPlay || BlackInPlay) {
		if WhiteInPlay {
			return board.White
		}
		return board.Black
	}
	return board.NoTeam
}

func DefaultInitBoard(self *board.Board) error {
	if self.Height%2 == 1 || self.Width%2 == 1 {
		return fmt.Errorf("Cannot init board with odd width or height")
	}
	heightOffset := (self.Height - 8) / 2
	widthOffset := (self.Width - 8) / 2

	offset := heightOffset*self.Width + widthOffset

	backRow := [8]board.PieceType{board.Rook, board.Knight, board.Bishop, board.Queen, board.King, board.Bishop, board.Knight, board.Rook}

	for i := range 8 {
		self.Pieces[offset+i].SetPieceTeam(board.Black)
		self.Pieces[offset+i].SetPieceType(backRow[i])

		self.Pieces[offset+i+self.Width].SetPieceTeam(board.Black)
		self.Pieces[offset+i+self.Width].SetPieceType(board.Pawn)

		self.Pieces[offset+i+self.Width*6].SetPieceTeam(board.White)
		self.Pieces[offset+i+self.Width*6].SetPieceType(board.Pawn)

		self.Pieces[offset+i+self.Width*7].SetPieceTeam(board.White)
		self.Pieces[offset+i+self.Width*7].SetPieceType(backRow[i])
	}

	return nil
}

func DefaultPawn(self board.Board, start int, end int) ([]int, []int) {
	var validMoveLocations []int = make([]int, 0)
	var validTakeLocations []int = make([]int, 0)

	var team = self.Pieces[start].GetPieceTeam()
	// white pawns move towards index 0, black pawns move away. This lets us combine the checks for each piece into one function
	var direction = 0

	if team == board.White {
		direction = -1
	} else {
		direction = 1
	}

	if self.Pieces[start+direction*self.Width].GetPieceTeam() == board.NoTeam {
		validMoveLocations = append(validMoveLocations, start+direction*self.Width)
	}

	if self.Pieces[start+direction*2*self.Width].GetPieceTeam() == board.NoTeam && !self.Pieces[start].GetPieceMoved() {
		validMoveLocations = append(validMoveLocations, start+direction*2*self.Width)
	}

	if self.Pieces[start+direction*self.Width+1].GetPieceTeam() == team.OtherTeam() {
		validTakeLocations = append(validTakeLocations, start+direction*self.Width+1)
	}

	if self.Pieces[start+direction*self.Width-1].GetPieceTeam() == team.OtherTeam() {
		validTakeLocations = append(validTakeLocations, start+direction*self.Width-1)
	}

	return validMoveLocations, validTakeLocations
}

func DefaultRook(self board.Board, start int, end int) ([]int, []int) {
	var directions [][2]int = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var validMoveLocations []int = GetMoveLocationsFromDirections(self, start, end, directions, true)
	// most pieces can take at the same spots they can move to, so i just return them both
	return validMoveLocations, validMoveLocations
}

func DefaultKnight(self board.Board, start int, end int) ([]int, []int) {
	var moveLocations [][2]int = [][2]int{{2, 1}, {2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}
	var validMoveLocations []int = GetMoveLocationsFromOffset(self, start, end, moveLocations)

	return validMoveLocations, validMoveLocations
}

func DefaultBishop(self board.Board, start int, end int) ([]int, []int) {
	var directions [][2]int = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	var validMoveLocations []int = GetMoveLocationsFromDirections(self, start, end, directions, true)

	// most pieces can take at the same spots they can move to, so i just return them both
	return validMoveLocations, validMoveLocations
}

func DefaultKing(self board.Board, start int, end int) ([]int, []int) {
	var moveLocations [8][2]int = [8][2]int{{1, 1}, {0, 1}, {-1, 1}, {1, 0}, {-1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	var validMoveLocations []int = make([]int, 0)

	for _, location := range moveLocations {
		var destination = start + location[0] + location[1]*self.Width
		if destination >= self.Width*self.Height || destination < 0 {
			continue
		}

		validMoveLocations = append(validMoveLocations, destination)
	}

	return validMoveLocations, validMoveLocations
}

func DefaultQueen(self board.Board, start int, end int) ([]int, []int) {
	var bishopMoves, bishopTakes = DefaultBishop(self, start, end)
	var rookMoves, rookTakes = DefaultRook(self, start, end)

	return append(bishopMoves, rookMoves...), append(bishopTakes, rookTakes...)
}
