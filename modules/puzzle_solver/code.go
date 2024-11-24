package puzzlesolver

import (
	"fmt"
	"n-puzzle/modules/priority_queue"
	puzzlestate "n-puzzle/modules/puzzle_state"
	"n-puzzle/modules/utils"
	"time"
)

type solver struct {
	ClosedStatesMap  map[string]puzzlestate.IPuzzleState
	ClosedStatesKeys []string
}

func NewPuzzleSolver() IPuzzleSolver {
	return &solver{
		ClosedStatesMap:  make(map[string]puzzlestate.IPuzzleState),
		ClosedStatesKeys: make([]string, 0),
	}
}

func (ps *solver) addClosedState(state puzzlestate.IPuzzleState) {
	key := state.Encrypt()
	ps.ClosedStatesMap[key] = state
	ps.ClosedStatesKeys = append(ps.ClosedStatesKeys, key)
}

func (ps *solver) getLastExploredState() puzzlestate.IPuzzleState {
	if len(ps.ClosedStatesKeys) == 0 {
		return nil
	}
	lastKey := ps.ClosedStatesKeys[len(ps.ClosedStatesKeys)-1]
	return ps.ClosedStatesMap[lastKey]
}

// TODO: ???
//func (ps *solver) lessFvalElementIndex(states []puzzlestate.IPuzzleState) int {
//	//if len(states) == 0 {
//	//	return -1
//	//}
//	minFval := states[0].GetFval()
//	index := 0
//	for i, state := range states {
//		currFval := state.GetFval()
//		if minFval > currFval {
//			minFval = currFval
//			index = i
//		}
//	}
//	return index
//}

//func (ps *solver) calcHeuristicVal(currentState, goalState puzzlestate.IPuzzleState, heuristic func(s, g puzzlestate.IPuzzleState) int) int {
//	return heuristic(currentState, goalState) + currentState.GetLevel()
//}

// countInversions shows how the current puzzle state is close by the goal state, to the sorted array
func (ps *solver) countInversions(initialState puzzlestate.IPuzzleState) int {
	counter := 0
	array := initialState.ConvertToArray()
	size := initialState.GetSize()

	for i := 0; i < size*size-1; i++ {
		for j := i + 1; j < size*size; j++ {
			if array[i] > array[j] {
				counter += 1
			}
		}
	}
	return counter
}

// findXPosition helper function which helps to find the position of the empty tile by the row from bottom to top
// Returns the index of the empty tile from the bottom of the puzzle by the row
func (ps *solver) findXPosition(state puzzlestate.IPuzzleState) int {
	for i := state.GetSize() - 1; i >= 0; i-- {
		for j := state.GetSize() - 1; j >= 0; j-- {
			value, _ := state.GetValueByIndexes(i, j)
			if value == 0 {
				return state.GetSize() - i
			}
		}
	}
	return -1
}

func (ps *solver) ifSolvable(initialState puzzlestate.IPuzzleState) bool {
	size := initialState.GetSize()
	amountOfInversions := ps.countInversions(initialState)

	if size%2 == 1 {
		return amountOfInversions%2 == 0
	} else {
		pos := ps.findXPosition(initialState)
		if pos%2 == 1 {
			return amountOfInversions%2 == 0
		} else {
			return amountOfInversions%2 == 1
		}
	}
}

func (ps *solver) ClearExploredPuzzleStates() {
	for _, key := range ps.ClosedStatesKeys {
		delete(ps.ClosedStatesMap, key)
	}
	ps.ClosedStatesKeys = ps.ClosedStatesKeys[0:0]
}

func (ps *solver) printInfo(isSolvable bool, states [][][]int, complexityInTime, complexityInSize, numberOfMoves int) {
	if isSolvable == false {
		fmt.Println("The current puzzle isn't solvable")
		return
	}
	fmt.Println("The current puzzle is solvable")

	fmt.Println("Complexity in time: ", complexityInTime)
	fmt.Println("Complexity in size: ", complexityInSize)
	fmt.Println("Number of moves:    ", numberOfMoves)
	for _, state := range states {
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

func (ps *solver) Solve(initialStateArray, goalStateArray [][]int, heurictic int) (bool, error) {
	ps.ClearExploredPuzzleStates()

	heuristicFunc, ok := puzzlestate.DistanceFunctionNames[heurictic]
	if ok == false {
		return false, utils.ErrorHeuristic
	}

	goalState := puzzlestate.NewPuzzleState(goalStateArray, 0, nil, nil, nil)
	initialState := puzzlestate.NewPuzzleState(initialStateArray, 0, heuristicFunc, goalState, nil)

	// Check if the initial state is solvable
	if ps.ifSolvable(initialState) == false {
		return false, nil
	}

	complexityInTime := 0
	complexityInSize := 0
	numberOfMoves := 0
	orderedSequenceOfStates := [][][]int{}
	openStates := priority_queue.NewQueue[puzzlestate.IPuzzleState]()
	openStates.Push(initialState)
	complexityInTime += 1
	complexityInSize = 1
	for {
		if openStates.Size()+len(ps.ClosedStatesMap) > complexityInSize {
			complexityInSize = openStates.Size() + len(ps.ClosedStatesMap)
		}
		if openStates.Empty() == true {
			return false, nil
		}
		currentState := openStates.Front()
		openStates.Pop()
		if currentState.GetFval() == 0 {
			list := currentState.ListOfStates()
			for _, l := range list {
				orderedSequenceOfStates = append(orderedSequenceOfStates, l.CopyMatrix())
			}
			ps.printInfo(true, orderedSequenceOfStates, complexityInTime, complexityInSize, numberOfMoves)
			break
		}
		expandedNodes := puzzlestate.Actions(currentState, goalState, heuristicFunc)
		ps.addClosedState(currentState)

		for _, node := range expandedNodes {
			encrypted := node.Encrypt()
			_, ok = ps.ClosedStatesMap[encrypted]
			if ok == false {
				openStates.Push(node)
				complexityInTime += 1
			}
		}
		time.Sleep(2 * time.Second)
		numberOfMoves += 1
	}

	return true, nil
}
