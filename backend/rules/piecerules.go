package rules

import (
	"fmt"

	"prushton.com/randochess/v2/board"
)

func DefaultPawn(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	// moving backwards
	if delta_y < 0 && self.Pieces[start].GetPieceTeam() == board.White {
		return false
	}

	if delta_y > 0 && self.Pieces[start].GetPieceTeam() == board.Black {
		return false
	}

	// moving 1 space or 2 on first turn
	if (Abs(delta_y) == 1 && delta_x == 0) || (Abs(delta_y) == 2 && !self.Pieces[start].GetPieceMoved()) {
		return self.Pieces[end].GetPieceTeam() == board.NoTeam && CheckLineOfSight(self, start, end)
	}

	// taking
	if Abs(delta_y) == 1 && Abs(delta_x) == 1 && self.Pieces[end].GetPieceTeam() != board.NoTeam {
		return true
	}

	return false
}

func DefaultRook(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return ((delta_x == 0) != (delta_y == 0)) && CheckLineOfSight(self, start, end)
}

func DefaultKnight(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return (Abs(delta_x) == 1 && Abs(delta_y) == 2) || (Abs(delta_x) == 2 && Abs(delta_y) == 1)
}
func DefaultBishop(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return Abs(delta_x) == Abs(delta_y) && CheckLineOfSight(self, start, end)
}

func DefaultKing(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	return delta_x >= -1 && delta_x <= 1 && delta_y >= -1 && delta_y <= 1
}

func DefaultQueen(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	if delta_x == 0 || delta_y == 0 {
		return CheckLineOfSight(self, start, end)
	}

	if Abs(delta_x) == Abs(delta_y) {
		return CheckLineOfSight(self, start, end)
	}

	return false
}

func PrepareThyselfPawn(self board.Board, start int, end int) bool {
	if self.Pieces[start].GetPieceMoved() {
		return DefaultPawn(self, start, end)
	}

	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	fmt.Printf("dx: %d, dy: %d\n", delta_x, delta_y)

	if delta_x != 0 {
		return false
	}

	if delta_y > 0 && self.Pieces[start].GetPieceTeam() == board.White || delta_y < 0 && self.Pieces[start].GetPieceTeam() == board.Black {
		return CheckLineOfSight(self, start, end)
	}

	return false

}

func PlanBishop(self board.Board, start int, end int) bool {
	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	// cant move well, but can take without LOS
	if self.Pieces[end].GetPieceTeam() == board.NoTeam {
		return Abs(delta_x) <= 1 && Abs(delta_y) <= 1
	} else {
		return Abs(delta_x) == Abs(delta_y)
	}
}
