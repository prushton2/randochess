package board

import "fmt"

type Board struct {
	width  uint8
	height uint8
	pieces []Piece
	turn   Team
}

func New(w uint8, h uint8) Board {
	board := Board{
		width:  w,
		height: h,
		pieces: make([]Piece, w*h),
		turn:   White,
	}

	return board
}

func (self *Board) Init8x8Board() error {
	if self.width != 8 || self.height != 8 {
		return fmt.Errorf("Cannot init board, not 8x8")
	}

	backRow := [8]PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for i := range 8 {
		self.pieces[i].SetPieceTeam(Black)
		self.pieces[i].SetPieceType(backRow[i])

		self.pieces[56+i].SetPieceTeam(White)
		self.pieces[56+i].SetPieceType(backRow[i])

		self.pieces[8+i].SetPieceTeam(Black)
		self.pieces[8+i].SetPieceType(Pawn)

		self.pieces[48+i].SetPieceTeam(White)
		self.pieces[48+i].SetPieceType(Pawn)
	}

	return nil
}
