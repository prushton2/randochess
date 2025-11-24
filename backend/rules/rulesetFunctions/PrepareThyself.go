package rulesetfunctions

import "prushton.com/randochess/v2/board"

func PrepareThyselfPawn(self board.Board, start int) ([]int, []int) {
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

	if !self.Pieces[start].GetPieceMoved() {
		validMoveLocations = append(validMoveLocations, GetMoveLocationsFromDirections(self, start, [][2]int{{0, direction}}, true, 4)...)
		validTakeLocations = validMoveLocations
	} else {
		if self.Pieces[start+direction*self.Width].GetPieceTeam() == board.NoTeam {
			validMoveLocations = append(validMoveLocations, start+direction*self.Width)
		}
	}

	if self.Pieces[start+direction*self.Width+1].GetPieceTeam() == team.OtherTeam() {
		validTakeLocations = append(validTakeLocations, start+direction*self.Width+1)
	}

	if self.Pieces[start+direction*self.Width-1].GetPieceTeam() == team.OtherTeam() {
		validTakeLocations = append(validTakeLocations, start+direction*self.Width-1)
	}

	return validMoveLocations, validTakeLocations
}
