package puzzlesolver

import (
	"n-puzzle/modules/priority_queue"
	puzzlestate "n-puzzle/modules/puzzle_state"
	"n-puzzle/modules/utils"
	"os"
	"strconv"
	"strings"
)

type solver struct {
	ClosedStatesMap         map[string]puzzlestate.IPuzzleState
	ClosedStatesKeys        []string
	OrderedSequenceOfStates [][][]int
	ComplexityInTime        int
	ComplexityInSize        int
	NumberOfMoves           int
}

func NewPuzzleSolver() IPuzzleSolver {
	return &solver{
		ClosedStatesMap:         make(map[string]puzzlestate.IPuzzleState),
		ClosedStatesKeys:        make([]string, 0),
		OrderedSequenceOfStates: make([][][]int, 0),
		ComplexityInTime:        0,
		ComplexityInSize:        0,
		NumberOfMoves:           0,
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
func (ps *solver) findXPosition(state [][]int) int {
	size := len(state)
	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j >= 0; j-- {
			if state[i][j] == 0 {
				return size - i
			}
		}
	}
	return -1
}

func (ps *solver) ifSolvable(initialState puzzlestate.IPuzzleState) bool {
	size := initialState.GetSize()
	inversions := ps.countInversions(initialState.ConvertToArray())

	if size%2 == 1 {
		return inversions%2 == 0
	} else {
		pos := ps.findXPosition(initialState.CopyMatrix())
		if pos%2 == 1 {
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
	ps.OrderedSequenceOfStates = ps.OrderedSequenceOfStates[0:0]

	ps.ComplexityInTime = 0
	ps.ComplexityInSize = 0
	ps.NumberOfMoves = 0
}

func (ps *solver) ConvertPuzzleInfoIntoString(initialState [][]int, isSolvable bool) string {
	info := ""
	for _, row := range utils.ConvertPuzzleToArrayOfStrings(initialState) {
		info += row
	}
	if isSolvable == false {
		info += "Puzzle is NOT SOLVABLE\n"
		return info
	}
	info += "Puzzle is SOLVABLE\n\n"

	info += "Complexity in time: " + strconv.Itoa(ps.ComplexityInTime) + "\n"
	info += "Complexity in size: " + strconv.Itoa(ps.ComplexityInSize) + "\n"
	info += "Number of moves:    " + strconv.Itoa(ps.NumberOfMoves) + "\n"
	return info
}

func (ps *solver) WriteToFile(initialState [][]int, filepath string, isSolvable bool) error {
	dirPathIndex := strings.LastIndex(filepath, "/")
	dirPath := filepath[0:dirPathIndex]
	// Create the output file directory of not exists
	err := utils.CreateDirectories(dirPath)
	if err != nil {
		return err
	}

	// Get complexity, number of moves, etc. and convert it into the string
	info := ps.ConvertPuzzleInfoIntoString(initialState, isSolvable)

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	// Write complexity, number of moves, etc. into the file
	_, err = f.Write([]byte(info))
	if err != nil {
		_ = f.Close()
		return err
	}

	// Write puzzle states into the file
	for _, state := range ps.OrderedSequenceOfStates {
		convertedPuzzle := utils.ConvertPuzzleToArrayOfStrings(state)
		convertedPuzzle = append(convertedPuzzle, "\n")
		for _, row := range convertedPuzzle {
			_, err = f.Write([]byte(row))
			if err != nil {
				_ = f.Close()
				return err
			}
		}
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ps *solver) Solve(initialStateArray, goalStateArray [][]int, hFunction func(s, g [][]int) int, outputFilePath string) (bool, error) {
	ps.Clear()

	goalState := puzzlestate.NewPuzzleState(goalStateArray, 0, nil, nil, nil)
	initialState := puzzlestate.NewPuzzleState(initialStateArray, 0, hFunction, goalState, nil)

	// Check if the initial state is solvable
	if ps.ifSolvable(initialState) == false {
		err := ps.WriteToFile(initialState.CopyMatrix(), outputFilePath, false)
		return false, err
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
			err = ps.WriteToFile(initialState.CopyMatrix(), outputFilePath, true)
			if err != nil {
				return false, err
			}
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
