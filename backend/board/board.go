package board

import "fmt"

type Board struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Pieces []Piece `json:"pieces"`
}

func New(w int, h int) Board {
	return Board{
		Width:  int(w),
		Height: int(h),
		Pieces: make([]Piece, w*h),
	}
}

func (self *Board) InitBoard() error {
	if self.Height%2 == 1 || self.Width%2 == 1 {
		return fmt.Errorf("Cannot init board with odd width or height")
	}
	heightOffset := (self.Height - 8) / 2
	widthOffset := (self.Width - 8) / 2

	offset := heightOffset*self.Width + widthOffset

	backRow := [8]PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for i := range 8 {
		self.Pieces[offset+i].SetPieceTeam(Black)
		self.Pieces[offset+i].SetPieceType(backRow[i])

		self.Pieces[offset+56+i].SetPieceTeam(White)
		self.Pieces[offset+56+i].SetPieceType(backRow[i])

		self.Pieces[offset+8+i].SetPieceTeam(Black)
		self.Pieces[offset+8+i].SetPieceType(Pawn)

		self.Pieces[offset+48+i].SetPieceTeam(White)
		self.Pieces[offset+48+i].SetPieceType(Pawn)
	}

	return nil
}
