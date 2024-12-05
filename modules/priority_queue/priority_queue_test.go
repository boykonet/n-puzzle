package priority_queue

import (
	"fmt"
	puzzlestate "n-puzzle/modules/puzzle_state"
	"testing"
)

// 1 2 3
// 4 5 6 - goal state
// 7 8 0
//
//   ^
//   |
//                          5         6
// 1 2 3                  1 2 0     1 2 3    (1 2 3)
// 4 5 0                  4 5 3     4 5 6    (4 0 5)
// 7 8 6                  7 8 6     7 8 0    (7 8 6)
//
//   ^
//   |
//                          1         2        3        4
// 1 2 3                  1 0 3     1 2 3    1 2 3    1 2 3
// 4 0 5 - initial state  4 2 5     4 5 0    4 8 5    0 4 5
// 7 8 6                  7 8 6     7 8 6    7 0 6    7 8 6

func TestPriorityQueue(t *testing.T) {
	distance := puzzlestate.ManhattanDistance
	goalState := puzzlestate.NewPuzzleState([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}, 0, nil, nil, nil)
	initialState := puzzlestate.NewPuzzleState([][]int{{1, 2, 3}, {4, 0, 5}, {7, 8, 6}}, 0, distance, goalState, nil)

	generatedState1 := puzzlestate.NewPuzzleState([][]int{{1, 0, 3}, {4, 2, 5}, {7, 8, 6}}, 1, distance, goalState, initialState)
	generatedState2 := puzzlestate.NewPuzzleState([][]int{{1, 2, 3}, {4, 5, 0}, {7, 8, 6}}, 1, distance, goalState, initialState)
	generatedState3 := puzzlestate.NewPuzzleState([][]int{{1, 2, 3}, {4, 8, 5}, {7, 0, 6}}, 1, distance, goalState, initialState)
	generatedState4 := puzzlestate.NewPuzzleState([][]int{{1, 2, 3}, {0, 4, 5}, {7, 8, 6}}, 1, distance, goalState, initialState)

	generatedState5 := puzzlestate.NewPuzzleState([][]int{{1, 2, 0}, {4, 5, 3}, {7, 8, 6}}, 2, distance, goalState, generatedState2)
	generatedState6 := puzzlestate.NewPuzzleState([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}, 2, distance, goalState, generatedState2)

	pq := NewPriorityQueue[puzzlestate.IPuzzleState]()

	pq.Enqueue(initialState, initialState.GetFval())
	pq.Enqueue(generatedState1, generatedState1.GetFval())
	pq.Enqueue(generatedState2, generatedState2.GetFval())
	pq.Enqueue(generatedState3, generatedState3.GetFval())
	pq.Enqueue(generatedState4, generatedState4.GetFval())
	pq.Enqueue(generatedState5, generatedState5.GetFval())
	pq.Enqueue(generatedState6, generatedState6.GetFval())
	pq.Enqueue(generatedState6, generatedState6.GetFval())

	//fmt.Println(initialState.GetFval())
	//fmt.Println(generatedState1.GetFval())
	//fmt.Println(generatedState2.GetFval())
	//fmt.Println(generatedState3.GetFval())
	//fmt.Println(generatedState4.GetFval())
	//fmt.Println(generatedState5.GetFval())
	//fmt.Println(generatedState6.GetFval())

	for pq.Len() != 0 {
		elem, err := pq.Dequeue()
		if err != nil {
			fmt.Println("Error:", err)
			t.Fail()
			return
		}
		elem.PrintPuzzle()
		fmt.Println(elem.GetFval())
	}
}
