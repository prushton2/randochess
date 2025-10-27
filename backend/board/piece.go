package board

import "encoding/json"

type PieceType int64
type Team int64

const ( // 3 bits
	Pawn   PieceType = 0
	Rook   PieceType = 1
	Knight PieceType = 2
	Bishop PieceType = 3
	Queen  PieceType = 4
	King   PieceType = 5
)

const ( // 2 bits
	NoTeam Team = 0
	White  Team = 2
	Black  Team = 3
)

type Piece uint8

// 000 00 000
//      |   |
//      |   PieceType
//      Team

func (p Piece) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint8(p))
}

func (self *Piece) SetPieceType(p PieceType) {
	new := (uint8(*self))&0b11111000 + uint8(p)
	*self = Piece(new)
}

func (self *Piece) SetPieceTeam(t Team) {
	new := (uint8(*self))&0b11100111 + (uint8(t) << 3)
	*self = Piece(new)
}

func (self *Piece) SetPieceMoved() {
	new := (uint8(*self)) | 0b00100000
	*self = Piece(new)
}

func (self *Piece) GetPieceType() PieceType {
	new := (uint8(*self) & 0b00000111)
	return PieceType(new)
}

func (self *Piece) GetPieceTeam() Team {
	new := (uint8(*self)) & 0b00011000
	return Team(new >> 3)
}

func (self *Piece) GetPieceMoved() bool {
	new := (uint8(*self)) & 0b00100000
	return new>>5 == 1
}

func (self Team) OtherTeam() Team {
	if self == NoTeam {
		return NoTeam
	}
	return Team(uint8(self) ^ 0b1)
}
