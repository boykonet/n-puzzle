package puzzlestate

import (
	"fmt"
	"n-puzzle/modules/utils"
)

// Result returns a new child puzzle state depends on the given action (swipe empty tile up, down, left or right)
func Result(state IPuzzleState, action int) IPuzzleState {
	if action < SwapLeft && action > SwapDown {
		return nil
	}

	x0, y0, _ := state.Coordinates(0)
	childData := state.CopyPuzzle()
	correction := cc[action]

	x, y := x0+correction.X, y0+correction.Y

	if !((x >= 0 && x < state.GetSize()) && (y >= 0 && y < state.GetSize())) {
		return nil
	}

	fmt.Println("x =", x, "y =", y)

	utils.Swap(&childData[x0][y0], &childData[x][y])
	return NewPuzzleTiles(childData, state.GetSize(), state.GetLevel(), state.GetFval(), state)
}

// Actions Takes the current state and returns the child states
func Actions(state IPuzzleState) []IPuzzleState {
	states := make([]IPuzzleState, 0, 4)

	for key, _ := range ccKeys {
		newState := Result(state, key)
		if newState != nil {
			states = append(states, newState)
		}
	}
	return states
}
