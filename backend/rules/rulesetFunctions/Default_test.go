package rulesetfunctions

import (
	"slices"
	"testing"

	"prushton.com/randochess/v2/board"
)

func compareExpectedOut(t *testing.T, calculatedMove []int, calculatedTake []int, expectedMove []int, expectedTake []int) {
	if len(calculatedMove) != len(expectedMove) {
		t.Errorf("Available move spots length is %d instead of %d", len(calculatedMove), len(expectedMove))
	}

	if len(calculatedTake) != len(expectedTake) {
		t.Errorf("Available take spots length is %d instead of %d", len(calculatedTake), len(expectedTake))
	}

	for _, e := range expectedMove {
		if !slices.Contains(calculatedMove[:], e) {
			t.Errorf("Move does not contain %d", e)
		}
	}

	for _, e := range expectedTake {
		if !slices.Contains(calculatedTake[:], e) {
			t.Errorf("Take does not contain %d", e)
		}
	}
}

func TestPawn_Move(t *testing.T) {
	// This is testing all of the move functionality

	brd := board.New(8, 8)
	brd.Pieces[43].SetPieceType(board.Pawn)

	// White side no move
	brd.Pieces[43].SetPieceTeam(board.White)
	var move, take = DefaultPawn(brd, 43) // end pos doesnt change the output

	compareExpectedOut(t, move, take, []int{27, 35}, []int{})

	// Black side no move
	brd.Pieces[43].SetPieceTeam(board.Black)

	move, take = DefaultPawn(brd, 43)

	compareExpectedOut(t, move, take, []int{51, 59}, []int{})

	// White side moved
	brd.Pieces[43].SetPieceMoved()
	brd.Pieces[43].SetPieceTeam(board.White)
	move, take = DefaultPawn(brd, 43) // end pos doesnt change the output

	compareExpectedOut(t, move, take, []int{35}, []int{})

	// Black side moved
	brd.Pieces[43].SetPieceTeam(board.Black)
	move, take = DefaultPawn(brd, 43) // end pos doesnt change the output

	compareExpectedOut(t, move, take, []int{51}, []int{})
}

func TestPawn_Take(t *testing.T) {
	brd := board.New(8, 8)
	brd.Pieces[43].SetPieceType(board.Pawn)

	brd.Pieces[52].SetPieceType(board.Pawn)
	brd.Pieces[52].SetPieceTeam(board.White)

	brd.Pieces[50].SetPieceType(board.Pawn)
	brd.Pieces[50].SetPieceTeam(board.White)

	brd.Pieces[34].SetPieceType(board.Pawn)
	brd.Pieces[34].SetPieceTeam(board.Black)

	brd.Pieces[36].SetPieceType(board.Pawn)
	brd.Pieces[36].SetPieceTeam(board.Black)

	// White piece taking
	brd.Pieces[43].SetPieceTeam(board.White)
	var move, take = DefaultPawn(brd, 43) // end pos doesnt change the output

	compareExpectedOut(t, move, take, []int{35, 27}, []int{36, 34})

	// Black piece taking
	brd.Pieces[43].SetPieceTeam(board.Black)
	move, take = DefaultPawn(brd, 43) // end pos doesnt change the output

	compareExpectedOut(t, move, take, []int{51, 59}, []int{52, 50})
}

func TestKnight(t *testing.T) {
	var expected_out []int = []int{1, 3, 8, 12, 24, 28, 33, 35}

	brd := board.New(8, 8)
	brd.Pieces[17].SetPieceTeam(board.White)
	brd.Pieces[17].SetPieceType(board.Knight)
	var move, take = DefaultKnight(brd, 18) // end pos doesnt change the output

	compareExpectedOut(t, move, take, expected_out, expected_out)
}

func TestRook(t *testing.T) {
	var expected_out []int = []int{27, 19, 11, 36, 37, 38, 39, 34, 33, 32, 43, 51}

	brd := board.New(8, 8)
	brd.Pieces[35].SetPieceTeam(board.White)
	brd.Pieces[35].SetPieceType(board.Rook)

	brd.Pieces[11].SetPieceTeam(board.Black)
	brd.Pieces[11].SetPieceType(board.Pawn)

	brd.Pieces[51].SetPieceTeam(board.White)
	brd.Pieces[51].SetPieceType(board.King)

	var move, take = DefaultRook(brd, 35) // end pos doesnt change the output

	// fmt.Printf("t: %v\n", move)

	compareExpectedOut(t, move, take, expected_out, expected_out)
}

func TestBishop(t *testing.T) {
	var expected_out []int = []int{8, 17, 21, 26, 28, 42, 44, 49, 53, 62}

	brd := board.New(8, 8)
	brd.Pieces[35].SetPieceTeam(board.White)
	brd.Pieces[35].SetPieceType(board.Bishop)

	brd.Pieces[21].SetPieceTeam(board.Black)
	brd.Pieces[21].SetPieceType(board.Pawn)

	brd.Pieces[49].SetPieceTeam(board.White)
	brd.Pieces[49].SetPieceType(board.King)

	var move, take = DefaultBishop(brd, 35) // end pos doesnt change the output

	compareExpectedOut(t, move, take, expected_out, expected_out)
}

func TestKing(t *testing.T) {
	var expected_out []int = []int{26, 27, 28, 34, 36, 42, 43, 44}

	brd := board.New(8, 8)
	brd.Pieces[35].SetPieceTeam(board.White)
	brd.Pieces[35].SetPieceType(board.King)

	var move, take = DefaultKing(brd, 35) // end pos doesnt change the output

	compareExpectedOut(t, move, take, expected_out, expected_out)
}

func TestQueen(t *testing.T) {
	var expected_out []int = []int{8, 17, 21, 26, 28, 42, 44, 49, 53, 62, 27, 19, 11, 36, 37, 38, 39, 34, 33, 32, 43, 51}

	brd := board.New(8, 8)
	brd.Pieces[35].SetPieceTeam(board.White)
	brd.Pieces[35].SetPieceType(board.Bishop)

	brd.Pieces[21].SetPieceTeam(board.Black)
	brd.Pieces[21].SetPieceType(board.Pawn)

	brd.Pieces[49].SetPieceTeam(board.White)
	brd.Pieces[49].SetPieceType(board.King)

	brd.Pieces[11].SetPieceTeam(board.Black)
	brd.Pieces[11].SetPieceType(board.Pawn)

	brd.Pieces[51].SetPieceTeam(board.White)
	brd.Pieces[51].SetPieceType(board.King)

	var move, take = DefaultQueen(brd, 35)

	compareExpectedOut(t, move, take, expected_out, expected_out)
}
