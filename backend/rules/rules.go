package rules

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"prushton.com/randochess/v2/board"
)

type Ruleset struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PieceRules  map[board.PieceType]func(board.Board, int, int) ([]int, []int)
	Width       int
	Height      int
	Move        func(*board.Board, int, int, board.Team) board.Team
	GetWinner   func(board.Board) board.Team
	InitBoard   func(*board.Board) error
}

func (rs Ruleset) MarshalJSON() ([]byte, error) { //changing name a bit for consitency
	type MarshalableRuleset struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	return json.Marshal(MarshalableRuleset{Name: rs.Name, Description: rs.Description})
}

func SelectRuleset(name string) (Ruleset, error) {
	if name == "Random" {
		keys := make([]string, 0, len(AllRulesets))
		for k := range AllRulesets {
			keys = append(keys, k)
		}
		randomKey := keys[rand.Intn(len(keys))]
		return AllRulesets[randomKey], nil
	}

	ruleset, exists := AllRulesets[name]
	if !exists {
		return Ruleset{}, fmt.Errorf("invalid name")
	}

	return ruleset, nil
}
