package board

import "fmt"

type Board struct {
	Width  int
	Height int
	Pieces []Piece
}

func New(w uint8, h uint8) Board {
	return Board{
		Width:  int(w),
		Height: int(h),
		Pieces: make([]Piece, w*h),
	}
}

func (self *Board) Init8x8Board() error {
	if self.Width != 8 || self.Height != 8 {
		return fmt.Errorf("Cannot init board, not 8x8")
	}

	backRow := [8]PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for i := range 8 {
		self.Pieces[i].SetPieceTeam(Black)
		self.Pieces[i].SetPieceType(backRow[i])

		self.Pieces[56+i].SetPieceTeam(White)
		self.Pieces[56+i].SetPieceType(backRow[i])

		self.Pieces[8+i].SetPieceTeam(Black)
		self.Pieces[8+i].SetPieceType(Pawn)

		self.Pieces[48+i].SetPieceTeam(White)
		self.Pieces[48+i].SetPieceType(Pawn)
	}

	return nil
}
