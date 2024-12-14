package puzzlestate

import (
	"n-puzzle/modules/utils"
)

// Result returns a new child puzzle state depends on the given action (swipe empty tile up, down, left or right)
func Result(s, g IPuzzleState, action int, heuristicFunc func(s, g [][]int) int) IPuzzleState {
	if action < SwapLeft && action > SwapDown {
		return nil
	}

	// Find the coordinate of the empty tile
	y0, x0, _ := s.Coordinates(0)

	childData := s.CopyMatrix()
	correction := coordCorrection[action]

	y, x := y0+correction.Y, x0+correction.X

	if !((x >= 0 && x < s.GetSize()) && (y >= 0 && y < s.GetSize())) {
		return nil
	}

	utils.Swap(&childData[y0][x0], &childData[y][x])
	return NewPuzzleState(childData, s.GetLevel()+1, heuristicFunc, g, s)
}

// Actions Takes the current state and returns the child states
func Actions(s, g IPuzzleState, heuristicFunc func(s, g [][]int) int) []IPuzzleState {
	states := make([]IPuzzleState, 0, 4)

	for _, action := range actions {
		newState := Result(s, g, action, heuristicFunc)
		if newState != nil {
			states = append(states, newState)
		}
	}
	return states
}
