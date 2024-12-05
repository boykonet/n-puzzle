package puzzlesolver

import (
	"fmt"
	"n-puzzle/modules/priority_queue"
	puzzlestate "n-puzzle/modules/puzzle_state"
)

type solver struct {
	ClosedStatesMap         map[string]puzzlestate.IPuzzleState
	ClosedStatesKeys        []string
	ComplexityInTime        int
	ComplexityInSize        int
	NumberOfMoves           int
	OrderedSequenceOfStates [][][]int
}

func NewPuzzleSolver() IPuzzleSolver {
	return &solver{
		ClosedStatesMap:         make(map[string]puzzlestate.IPuzzleState),
		ClosedStatesKeys:        make([]string, 0),
		ComplexityInTime:        0,
		ComplexityInSize:        0,
		NumberOfMoves:           0,
		OrderedSequenceOfStates: make([][][]int, 0),
	}
}

func (ps *solver) addClosedState(state puzzlestate.IPuzzleState) {
	key := state.Encrypt()
	ps.ClosedStatesMap[key] = state
	ps.ClosedStatesKeys = append(ps.ClosedStatesKeys, key)
}

// countInversions shows how the current puzzle state is close by the goal state, to the sorted array
func (ps *solver) countInversions(puzzle []int) int {
	amountOfInversions := 0

	for i := 0; i < len(puzzle); i++ {
		for j := i + 1; j < len(puzzle); j++ {
			if puzzle[i] != 0 && puzzle[j] != 0 && puzzle[i] > puzzle[j] {
				amountOfInversions += 1
			}
		}
	}
	return amountOfInversions
}

// findXPosition helper function which helps to find the position of the empty tile by the row from bottom to top
// Returns the index of the empty tile from the bottom of the puzzle by the row
func (ps *solver) findXPosition(puzzle []int, size int) int {
	//for i := state.GetSize() - 1; i >= 0; i-- {
	//	for j := state.GetSize() - 1; j >= 0; j-- {
	//		value, _ := state.GetValueByIndexes(i, j)
	//		if value == 0 {
	//			return state.GetSize() - i
	//		}
	//	}
	//}
	//return -1
	blankIndex := -1
	for i := 0; i < len(puzzle); i++ {
		if puzzle[i] == 0 {
			blankIndex = i
			break
		}
	}
	return size - (blankIndex / size)
}

func (ps *solver) ifSolvable(initialState puzzlestate.IPuzzleState) bool {
	size := initialState.GetSize()
	inversions := ps.countInversions(initialState.ConvertToArray())
	blankRow := ps.findXPosition(initialState.ConvertToArray(), size)

	if size%2 == 1 {
		return inversions%2 == 0
	} else {
		if (size-blankRow)%2 == 1 {
			return inversions%2 == 0
		} else {
			return inversions%2 == 1
		}
	}
}

func (ps *solver) Clear() {
	for _, key := range ps.ClosedStatesKeys {
		delete(ps.ClosedStatesMap, key)
	}
	ps.ClosedStatesKeys = ps.ClosedStatesKeys[0:0]
	ps.ComplexityInTime = 0
	ps.ComplexityInSize = 0
	ps.NumberOfMoves = 0
	ps.OrderedSequenceOfStates = ps.OrderedSequenceOfStates[0:0]
}

func (ps *solver) printInfo(isSolvable bool) {
	if isSolvable == false {
		fmt.Println("The current puzzle isn't solvable")
		return
	}
	fmt.Println("The current puzzle is solvable")

	fmt.Println("Complexity in time: ", ps.ComplexityInTime)
	fmt.Println("Complexity in size: ", ps.ComplexityInSize)
	fmt.Println("Number of moves:    ", ps.NumberOfMoves)
	for _, state := range ps.OrderedSequenceOfStates {
		for _, row := range state {
			for _, elem := range row {
				fmt.Print(elem)
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
		fmt.Print("\n\n")
	}
}

func (ps *solver) Solve(initialStateArray, goalStateArray [][]int, hFunction func(s, g [][]int) int) (bool, error) {
	ps.Clear()

	goalState := puzzlestate.NewPuzzleState(goalStateArray, 0, nil, nil, nil)
	initialState := puzzlestate.NewPuzzleState(initialStateArray, 0, hFunction, goalState, nil)

	// Check if the initial state is solvable
	if ps.ifSolvable(initialState) == false {
		fmt.Println("UNSOLVABLE")
		return false, nil
	}

	openStates := priority_queue.NewPriorityQueue[puzzlestate.IPuzzleState]()
	openStates.Enqueue(initialState, initialState.GetFval())
	ps.ComplexityInTime += 1
	ps.ComplexityInSize = 1
	for {
		if openStates.Len()+len(ps.ClosedStatesMap) > ps.ComplexityInSize {
			ps.ComplexityInSize = openStates.Len() + len(ps.ClosedStatesMap)
		}
		if openStates.Len() == 0 {
			return false, nil
		}
		currentState, err := openStates.Dequeue()
		if err != nil {
			return false, err
		}
		if currentState.GetFval() == 0 {
			ps.OrderedSequenceOfStates = currentState.ListOfStates()
			ps.NumberOfMoves = len(ps.OrderedSequenceOfStates)
			ps.printInfo(true)
			return true, nil
		}
		expandedNodes := puzzlestate.Actions(currentState, goalState, hFunction)
		ps.addClosedState(currentState)

		for _, node := range expandedNodes {
			encrypted := node.Encrypt()
			_, ok := ps.ClosedStatesMap[encrypted]
			if ok == false {
				openStates.Enqueue(node, node.GetFval())
				ps.ComplexityInTime += 1
			}
		}
	}
}
