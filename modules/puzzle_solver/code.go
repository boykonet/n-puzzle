package puzzlesolver

import (
	"fmt"
	puzzletiles "n-puzzle/modules/puzzle_tiles"
	"n-puzzle/modules/queue"
	"time"
)

const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

type solver struct {
	States []puzzletiles.IPuzzleTiles
	Closed []puzzletiles.IPuzzleTiles
}

func NewPuzzleSolver() IPuzzleSolver {
	return &solver{
		States: make([]puzzletiles.IPuzzleTiles, 0, 4),
		Closed: make([]puzzletiles.IPuzzleTiles, 0),
	}
}

func (ps *solver) appendState(child puzzletiles.IPuzzleTiles) {
	ps.States = append(ps.States, child)
}

func (ps *solver) appendClosed(closed puzzletiles.IPuzzleTiles) {
	ps.Closed = append(ps.Closed, closed)
}

func (ps *solver) getLastClosed() puzzletiles.IPuzzleTiles {
	if len(ps.Closed) == 0 {
		return nil
	}
	return ps.Closed[len(ps.Closed)-1]
}

func (ps *solver) heuristic(currentState, goalState puzzletiles.IPuzzleTiles) (int, error) {
	size := currentState.GetSize()
	temp := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			startValue, err1 := currentState.GetValueByIndexes(i, j)
			goalValue, err2 := goalState.GetValueByIndexes(i, j)
			if err1 != nil || err2 != nil {
				return 0, err1
			}
			if startValue != goalValue && startValue != 0 {
				temp += 1
			}
		}
	}
	return temp, nil
}

func (ps *solver) bestChildIndex() int {
	minimum := ps.States[0].GetFval()
	index := 0
	for i, child := range ps.States {
		if minimum > child.GetFval() {
			minimum = child.GetFval()
			index = i
		}
	}
	return index
}

func (ps *solver) calcHeuristicVal(open, closed puzzletiles.IPuzzleTiles) (int, error) {
	heuristic, err := ps.heuristic(open, closed)
	if err != nil {
		return 0, err
	}
	return heuristic + open.GetLevel(), nil
}

func (ps *solver) Solve(initialStateArray, goalStateArray [][]int) (bool, error) {
	initialState := puzzletiles.NewPuzzleTiles(initialStateArray, len(initialStateArray), 0, 0)
	goalState := puzzletiles.NewPuzzleTiles(goalStateArray, len(goalStateArray), 0, 0)

	fval, err := ps.calcHeuristicVal(initialState, goalState)
	if err != nil {
		return false, err
	}
	initialState.SetFval(fval)

	frontier := queue.NewQueue[puzzletiles.IPuzzleTiles]()
	frontier.Push(initialState)
	for {
		if frontier.Empty() {
			return false, fmt.Errorf("no solution")
		}
		currentState := frontier.Back()
		frontier.Pop()
		//initialState.PrintPuzzle()
		//fmt.Println()
		h, err := ps.heuristic(currentState, goalState)
		if err != nil {
			fmt.Print(err)
			return false, err
		}
		if h == 0 {
			break
		}
		states, err := currentState.GenerateStates(ps.getLastClosed())
		if err != nil {
			fmt.Print(err)
			return false, err
		}
		counter := 0
		for _, state := range states {
			fval, err = ps.calcHeuristicVal(state, goalState)
			if err != nil {
				fmt.Print(err)
				return false, err
			}
			state.SetFval(fval)
			ps.appendState(state)
			counter += 1
		}
		ps.appendClosed(currentState)
		currentState = ps.States[ps.bestChildIndex()]
		ps.States = ps.States[0:0]
		time.Sleep(2 * time.Second)
	}
	//fmt.Println(ps.InitialState)
	return true, nil
}
