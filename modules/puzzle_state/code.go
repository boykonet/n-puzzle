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
var cc = map[int]struct{ X, Y int }{
	SwapLeft:  {X: -1, Y: 0},
	SwapRight: {X: 1, Y: 0},
	SwapUp:    {X: 0, Y: -1},
	SwapDown:  {X: 0, Y: 1},
}

var ccKeys = []int{SwapLeft, SwapDown, SwapUp, SwapDown}

type tiles struct {
	Data   [][]int
	Size   int
	Level  int
	Fval   int
	Parent *tiles
}

func NewPuzzleTiles(data [][]int, size int, level int, fval int) IPuzzleState {
	return &tiles{
		Data:   utils.Duplicate2DArray(data),
		Size:   size,
		Level:  level,
		Fval:   fval,
		Parent: nil,
	}
}

func (t *tiles) isValidIndexes(x, y int) bool {
	if x >= 0 && x < t.Size && y >= 0 && y < t.Size {
		return true
	}
	return false
}

func (t *tiles) compareStates(parent, child IPuzzleState) int {
	for i := 0; i < t.Size; i++ {
		for j := 0; j < t.Size; j++ {
			pval, _ := parent.GetValueByIndexes(i, j)
			cval, _ := child.GetValueByIndexes(i, j)
			if pval != cval {
				return -1
			}
		}
	}
	return 0
}

// Encrypt based on the following principle: 1.2.3-4.5.6-7.8.0
func (t *tiles) Encrypt() string {
	var encryptedState string

	for i, row := range t.Data {
		for j, value := range row {
			encryptedState += strconv.Itoa(value)
			if j < t.Size-1 {
				encryptedState += "."
			}
		}
		if i < t.Size-1 {
			encryptedState += "-"
		}
	}
	return encryptedState
}

func (t *tiles) Result(action int) IPuzzleState {
	if action < SwapLeft && action > SwapDown {
		return nil
	}

	x0, y0, _ := t.findElemCoordinates(0)
	childData := utils.Duplicate2DArray(t.Data)
	correction := cc[action]

	if !t.isValidIndexes(x0+correction.X, y0+correction.Y) {
		return nil
	}

	utils.Swap(&childData[x0][y0], &childData[x0+correction.X][y0+correction.Y])
	return &tiles{
		Data:  childData,
		Size:  t.Size,
		Level: t.Level + 1,
		Fval:  t.Fval,
	}
}

func (t *tiles) Actions() []IPuzzleState {
	states := make([]IPuzzleState, 0, 4)

	for key, _ := range ccKeys {
		newState := t.Result(key)
		if newState != nil {
			states = append(states, newState)
		}
	}
	return states
}

func (t *tiles) SetFval(fval int) {
	t.Fval = fval
}

func (t *tiles) GetFval() int {
	return t.Fval
}

func (t *tiles) GetSize() int {
	return t.Size
}

func (t *tiles) GetValueByIndexes(i, j int) (int, error) {
	ok := t.isValidIndexes(i, j)
	if ok == false {
		return -1, fmt.Errorf("indexes are out of range")
	}
	return t.Data[i][j], nil
}

func (t *tiles) GetLevel() int {
	return t.Level
}

func (t *tiles) findElemCoordinates(element int) (int, int, error) {
	for i := 0; i < t.Size; i++ {
		for j := 0; j < t.Size; j++ {
			if t.Data[i][j] == element {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("element '%v' is not in the array", element)
}

func (t *tiles) PrintPuzzle() {
	for _, value := range t.Data {
		fmt.Println(value)
	}
}

// ConvertToArray Converts the matrix to array of integers
func (t *tiles) ConvertToArray() []int {
	var array []int

	for _, row := range t.Data {
		for _, elem := range row {
			array = append(array, elem)
		}
	}
	return array
}

// Coordinates Returns the coordinates (y, x, nil) of the number in matrix if the current number exists
// Otherwise, returns (-1, -1, error)
func (t *tiles) Coordinates(number int) (int, int, error) {
	for y := 0; y < t.Size; y++ {
		for x := 0; x < t.Size; x++ {
			if t.Data[y][x] == number {
				return y, x, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("Incorrect number: %v", number)
}

//func (t *tiles) calcHeuristicVal(goalState IPuzzleState, heuristic func(s, g IPuzzleState) int) int {
//	return heuristic(t., goalState) + t.Level
//}
