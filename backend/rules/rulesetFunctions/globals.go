package rulesetfunctions

import "prushton.com/randochess/v2/board"

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
		// fmt.Printf("Index %d: ", index)

		// oob? no you cant move there
		if index >= self.Width*self.Height || index < 0 {
			// fmt.Print("\n")
			hasLOS = false
			break
		}
		// fmt.Printf("%d-%d ", self.Pieces[index].GetPieceTeam(), self.Pieces[index].GetPieceType())

		// at any point if we encounter a piece we lose los
		if self.Pieces[index].GetPieceTeam() != board.NoTeam {
			hasLOS = false
		}
	}
	// fmt.Print("\n")

	return hasLOS
}

// directions is an array of directions to check. A directions looks like {1, 0} or {-1, 1} where index 0 is the x and index 1 is the y, determining how to move along each axis.
// This function will, for each direction, search until it hits the edge of the board or loses line of sight (if enabled). all board positions that are valid are returned.
func GetMoveLocationsFromDirections(self board.Board, start int, directions [][2]int, checkLOS bool, limit int) []int {
	var validMoveLocations []int = make([]int, 0)
	// construct an array of spaces where the piece can move.
	// Iterate over every direction and look until we reach the edge of the board or a piece
	for _, direction := range directions {
		var distance int = 1
		var last_destination = start
		var destination = start

		for distance <= limit || limit == -1 {
			last_destination = destination
			destination = start + direction[0]*distance + direction[1]*self.Width*distance

			// oob? exit
			if destination >= self.Width*self.Height || destination < 0 {
				break
			}

			var delta_x int = destination%self.Width - last_destination%self.Width
			var delta_y int = destination/self.Height - last_destination/self.Height

			// this condition is true if we are not on the first iteration and a few things happen:
			//    we overflow from one row to the next
			//    we overflow from one column to the next (idk how this would happen tbh)
			// if either is true, we stop crawling forward
			if (delta_x != direction[0] || delta_y != direction[1]) && destination != last_destination {
				break
			}

			if (CheckLineOfSight(self, start, destination) && checkLOS) || !checkLOS {
				validMoveLocations = append(validMoveLocations, destination)
			} else {
				break
			}

			distance += 1
		}
	}
	return validMoveLocations
}

// offsets is an array like directions, but contains offsets from the current piece's position. ie {2, 1}, {1, -2}. It returns all spaces that are inbounds. LOS checking is not possible here
// because LOS checking only works on cardinal and diagonal directions, not directions where movement is different and nonzero in each direction
func GetMoveLocationsFromOffset(self board.Board, start int, offsets [][2]int) []int {
	var validMoveLocations []int = make([]int, 0)

	for _, location := range offsets {
		var destination = start + location[0] + location[1]*self.Width
		if destination >= self.Width*self.Height || destination < 0 {
			continue
		}

		validMoveLocations = append(validMoveLocations, destination)
	}

	return validMoveLocations
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
