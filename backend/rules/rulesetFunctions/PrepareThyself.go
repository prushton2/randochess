package rulesetfunctions

import "prushton.com/randochess/v2/board"

func PrepareThyselfPawn(self board.Board, start int) bool {
	return true // for now
	// if self.Pieces[start].GetPieceMoved() {
	// 	return DefaultPawn(self, start, end)
	// }

	// var delta_x int = start%self.Width - end%self.Width
	// var delta_y int = start/self.Height - end/self.Height

	// if delta_x != 0 {
	// 	return false
	// }

	// if delta_y > 0 && self.Pieces[start].GetPieceTeam() == board.White || delta_y < 0 && self.Pieces[start].GetPieceTeam() == board.Black {
	// 	return CheckLineOfSight(self, start, end)
	// }

	// return false

}
