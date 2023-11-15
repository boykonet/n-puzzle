package puzzletiles

import (
	"fmt"
	"n-puzzle/modules/utils"
)

type tiles struct {
	PuzzleData [][]int
	PuzzleSize int
	Level      int
	Fval       int
}

func NewPuzzleTiles(data [][]int, size int, level int, fval int) IPuzzleTiles {
	return &tiles{
		PuzzleData: utils.Duplicate2Darray(data),
		PuzzleSize: size,
		Level:      level,
		Fval:       fval,
	}
}

func (t *tiles) getValidNeighboursIndexes(i, j int) [][]int {
	return func(combs [][]int) [][]int {
		neighbours := make([][]int, 0, 4)
		for _, comb := range combs {
			if comb[0] >= 0 && comb[0] < t.PuzzleSize &&
				comb[1] >= 0 && comb[1] < t.PuzzleSize {
				neighbours = append(neighbours, comb)
			}
		}
		return neighbours
	}([][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}})
}

func (t *tiles) GenerateChild() ([]IPuzzleTiles, error) {
	children := make([]IPuzzleTiles, 0, 4)
	i, j, err := t.findCoordinates('_')
	if err != nil {
		return nil, err
	}
	neighbours := t.getValidNeighboursIndexes(i, j)
	for _, neighbour := range neighbours {
		children = append(children, &tiles{
			PuzzleData: t.generateNewChild(i, j, neighbour[0], neighbour[1]),
			PuzzleSize: t.PuzzleSize,
			Level:      t.Level + 1,
			Fval:       0,
		})
	}
	return children, nil
}

func (t *tiles) SetFval(fval int) {
	t.Fval = fval
}

func (t *tiles) GetFval() int {
	return t.Fval
}

func (t *tiles) GetValueByIndexes(i, j int) (int, error) {
	if i >= t.PuzzleSize || j >= t.PuzzleSize {
		return -1, fmt.Errorf("The indexes out of range\n")
	}
	return t.PuzzleData[i][j], nil
}

func (t *tiles) GetLevel() int {
	return t.Level
}

func (t *tiles) findCoordinates(element int) (int, int, error) {
	for i := 0; i < t.PuzzleSize; i++ {
		for j := 0; j < t.PuzzleSize; j++ {
			if t.PuzzleData[i][j] == element {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("Invalid map: element '%v' is not in the array\n", element)
}

func (t *tiles) generateNewChild(i0, j0, i, j int) [][]int {
	newPuzzle := utils.Duplicate2Darray(t.PuzzleData)
	utils.Swap(&newPuzzle[i0][j0], &newPuzzle[i][j])
	return newPuzzle
}
