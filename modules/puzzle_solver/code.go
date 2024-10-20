package puzzlesolver

import (
	puzzlestate "n-puzzle/modules/puzzle_state"
	"n-puzzle/modules/queue"
	"n-puzzle/modules/utils"
	"time"
)

type solver struct {
	ExploredStatesMap  map[string]puzzlestate.IPuzzleState
	ExploredStatesKeys []string
}

func NewPuzzleSolver() IPuzzleSolver {
	return &solver{
		ExploredStatesMap:  make(map[string]puzzlestate.IPuzzleState),
		ExploredStatesKeys: make([]string, 0),
	}
}

func (ps *solver) addExploredState(state puzzlestate.IPuzzleState) {
	key := state.Encrypt()
	ps.ExploredStatesMap[key] = state
	ps.ExploredStatesKeys = append(ps.ExploredStatesKeys, key)
}

func (ps *solver) getLastExploredState() puzzlestate.IPuzzleState {
	if len(ps.ExploredStatesKeys) == 0 {
		return nil
	}
	lastKey := ps.ExploredStatesKeys[len(ps.ExploredStatesKeys)-1]
	return ps.ExploredStatesMap[lastKey]
}

//func ManhattanDistance(s, g puzzlestate.IPuzzleState) int {
//	var manhattan int
//
//	size := s.GetSize()
//	for i := 0; i < size*size; i++ {
//		y1, x1, _ := s.Coordinates(i)
//		y2, x2, _ := g.Coordinates(i)
//		xSteps := math.Abs(float64(x1 - x2))
//		ySteps := math.Abs(float64(y1 - y2))
//		manhattan += int(xSteps + ySteps)
//	}
//	return manhattan
//}
//
//func EuclideanDistance(s, g puzzlestate.IPuzzleState) int {
//	var euclidean float64
//
//	size := s.GetSize()
//	for i := 0; i < size*size; i++ {
//		y1, x1, _ := s.Coordinates(i)
//		y2, x2, _ := g.Coordinates(i)
//		xDistance := math.Abs(float64(x1 - x2))
//		yDistance := math.Abs(float64(y1 - y2))
//		euclidean += math.Sqrt(xDistance*xDistance + yDistance*yDistance)
//	}
//	return int(euclidean)
//}
//
//func ChebyshevDistance(s, g puzzlestate.IPuzzleState) int {
//	var chebyshev int
//
//	size := s.GetSize()
//	for i := 0; i < size*size; i++ {
//		y1, x1, _ := s.Coordinates(i)
//		y2, x2, _ := g.Coordinates(i)
//		xSteps := math.Abs(float64(x1 - x2))
//		ySteps := math.Abs(float64(y1 - y2))
//		chebyshev += int(math.Max(xSteps, ySteps))
//	}
//	return chebyshev
//}
//
//var ale = map[int]func(s, g puzzlestate.IPuzzleState) int{
//	utils.ManhattanHeuristic: ManhattanDistance,
//	utils.EuclideanHeuristic: EuclideanDistance,
//	utils.ChebyshevHeuristic: ChebyshevDistance,
//}

func (ps *solver) lessFvalElementIndex(states []puzzlestate.IPuzzleState) int {
	//if len(states) == 0 {
	//	return -1
	//}
	minFval := states[0].GetFval()
	index := 0
	for i, state := range states {
		currFval := state.GetFval()
		if minFval > currFval {
			minFval = currFval
			index = i
		}
	}
	return index
}

func (ps *solver) calcHeuristicVal(currentState, goalState puzzlestate.IPuzzleState, heuristic func(s, g puzzlestate.IPuzzleState) int) int {
	return heuristic(currentState, goalState) + currentState.GetLevel()
}

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
	for _, key := range ps.ExploredStatesKeys {
		delete(ps.ExploredStatesMap, key)
	}
	ps.ExploredStatesKeys = ps.ExploredStatesKeys[0:0]
}

func (ps *solver) Solve(initialStateArray, goalStateArray [][]int, heurictic int) (bool, error) {
	ps.ClearExploredPuzzleStates()

	var totalNumberOfStates int

	initialState := puzzlestate.NewPuzzleTiles(initialStateArray, len(initialStateArray), 0, 0)
	goalState := puzzlestate.NewPuzzleTiles(goalStateArray, len(goalStateArray), 0, 0)

	if ps.ifSolvable(initialState) == false {
		return false, nil
	}

	heuristicFunc, ok := puzzlestate.DistanceFunctionNames[heurictic]
	if ok == false {
		return false, utils.ErrorHeuristic
	}

	fval := ps.calcHeuristicVal(initialState, goalState, heuristicFunc)
	initialState.SetFval(fval)

	frontier := queue.NewQueue[puzzlestate.IPuzzleState]()
	frontier.Push(initialState)
	totalNumberOfStates += 1
	for {
		if frontier.Empty() == true {
			return false, nil
		}
		currentState := frontier.Front()
		frontier.Pop()
		h := heuristicFunc(currentState, goalState)
		if h == 0 {
			break
		}
		expandedNodes := currentState.Actions()
		for _, node := range expandedNodes {
			fval = ps.calcHeuristicVal(node, goalState, heuristicFunc)
			node.SetFval(fval)
		}
		totalNumberOfStates += len(expandedNodes)

		ps.addExploredState(currentState)
		bestState := expandedNodes[ps.lessFvalElementIndex(expandedNodes)]
		encrypted := bestState.Encrypt()
		_, ok := ps.ExploredStatesMap[encrypted]
		if ok == false {
			frontier.Push(bestState)
		}
		time.Sleep(2 * time.Second)
	}
	//fmt.Println(ps.InitialState)
	return true, nil
}
