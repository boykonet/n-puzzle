package puzzlestate

import (
	"fmt"
	"n-puzzle/modules/utils"
	"strconv"
)

const (
	SwapLeft = iota
	SwapRight
	SwapUp
	SwapDown
)

// coordinate correction
var coordCorrection = map[int]struct{ X, Y int }{
	SwapLeft:  {X: -1, Y: 0},
	SwapRight: {X: 1, Y: 0},
	SwapUp:    {X: 0, Y: -1},
	SwapDown:  {X: 0, Y: 1},
}

var actions = []int{SwapLeft, SwapRight, SwapUp, SwapDown}

type state struct {
	Data   [][]int
	Size   int
	Level  int
	Fval   int
	Parent IPuzzleState
}

func NewPuzzleState(data [][]int, level int, heuristicFunc func(s, g [][]int) int, goal, parent IPuzzleState) IPuzzleState {
	var fval int
	if goal == nil {
		// current state is equivalent of the goal state
		fval = 0
	} else {
		fval = heuristicFunc(data, goal.CopyMatrix())
	}
	return &state{
		Data:   utils.Duplicate2DArray(data),
		Size:   len(data),
		Level:  level,
		Fval:   fval,
		Parent: parent,
	}
}

// Encrypt based on the following principle: 1.2.3.4.5.6.7.8.0
func (s *state) Encrypt() string {
	var encryptedState string

	for i, row := range s.Data {
		for j, value := range row {
			encryptedState += strconv.Itoa(value)
			if j < s.Size-1 {
				encryptedState += "."
			}
		}
		if i < s.Size-1 {
			encryptedState += "."
		}
	}
	return encryptedState
}

func (s *state) CopyMatrix() [][]int {
	return utils.Duplicate2DArray(s.Data)
}

func (s *state) GetFval() int {
	return s.Fval
}

func (s *state) GetSize() int {
	return s.Size
}

func (s *state) GetValueByIndexes(i, j int) (int, error) {
	if !(i >= 0 && i < s.Size && j >= 0 && j < s.Size) {
		return -1, fmt.Errorf("indexes are out of range")
	}
	return s.Data[i][j], nil
}

func (s *state) GetLevel() int {
	return s.Level
}

// ConvertToArray Converts the matrix to array of integers
func (s *state) ConvertToArray() []int {
	var array []int

	for _, row := range s.Data {
		for _, elem := range row {
			array = append(array, elem)
		}
	}
	return array
}

// Coordinates Returns the coordinates (y, x, nil) of the number in matrix if the current number exists.
// Otherwise, returns (-1, -1, error)
func (s *state) Coordinates(number int) (y int, x int, e error) {
	for y = 0; y < s.Size; y++ {
		for x = 0; x < s.Size; x++ {
			if s.Data[y][x] == number {
				return y, x, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("incorrect number: %v", number)
}

func (s *state) ListOfStates() [][][]int {
	nodes := make([][][]int, 0)
	if s.Parent != nil {
		nodes = s.Parent.ListOfStates()
	}
	return append(nodes, s.CopyMatrix())
}

func (s *state) PrintPuzzle() {
	for _, row := range s.Data {
		for index, val := range row {
			fmt.Print(val)
			if index < s.Size-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
