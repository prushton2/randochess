package rules

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"prushton.com/randochess/v2/board"
)

type Ruleset struct {
	Name       string `json:"name"`
	PieceRules map[board.PieceType]func(board.Board, int, int) bool
	Width      int
	Height     int
	Move       func(*board.Board, int, int)
	GetWinner  func(board.Board) board.Team
}

func (self Ruleset) MarshalJSON() ([]byte, error) {
	type MarshalableRuleset struct {
		Name string `json:"name"`
	}

	return json.Marshal(MarshalableRuleset{Name: self.Name})
}

func SelectRuleset(name string) (Ruleset, error) {
	if name == "Random" {
		keys := make([]string, 0, len(allRulesets))
		for k := range allRulesets {
			keys = append(keys, k)
		}
		randomKey := keys[rand.Intn(len(keys))]
		return allRulesets[randomKey], nil
	}

	ruleset, exists := allRulesets[name]
	if !exists {
		return Ruleset{}, fmt.Errorf("Invalid name")
	}

	return ruleset, nil
}

func CheckLineOfSight(self board.Board, start int, end int) bool {
	/*
		board is a large version of this:
		[0, 1, 2, 3
		 4, 5, 6, 7
		 8, 9,10,11
		12,13,14,15]
		X position is i%width
		y position is i/height
	*/

	var delta_x int = start%self.Width - end%self.Width
	var delta_y int = start/self.Height - end/self.Height

	// fmt.Printf("dx: %d\ndy: %d\n", delta_x, delta_y)

	// handle off cases like the horse where you dont move in a straight line, you are just allowed to move.
	if delta_x != 0 && delta_y != 0 {
		if Abs(delta_x) != Abs(delta_y) {
			return true
		}
	}

	iterations := Max(Abs(delta_x), Abs(delta_y))
	direction_x := -Clamp(delta_x, -1, 1)
	direction_y := -Clamp(delta_y, -1, 1)

	// dir x is negative for left, and positive for right
	// dir y is negative for up,   and positive for down
	// either is 0 for doesnt change
	hasLOS := true
	for n := range iterations {
		if n == 0 { // the first index is the piece itself
			continue
		}

		// we move the x and y coordinates by n in the directions defined
		index := start + (direction_x * n) + (direction_y * n * self.Width)
		fmt.Printf("Index %d: ", index)
		// oob? just skip it
		if index >= self.Width*self.Height || index < 0 {
			fmt.Print("\n")
			continue
		}
		fmt.Printf("%d-%d ", self.Pieces[index].GetPieceTeam(), self.Pieces[index].GetPieceType())

		// at any point if we encounter a piece we lose los
		if self.Pieces[index].GetPieceTeam() != board.NoTeam {
			hasLOS = false
		}
	}
	fmt.Print("\n")

	return hasLOS
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func Clamp(clampee int, low int, hi int) int {
	return Max(low, Min(clampee, hi))
}
