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
	return func(n_inds [][]int) [][]int {
		neighbours := make([][]int, 0, 4)
		for _, indxs := range n_inds {
			if indxs[0] >= 0 && indxs[0] < t.PuzzleSize &&
				indxs[1] >= 0 && indxs[1] < t.PuzzleSize {
				neighbours = append(neighbours, indxs)
			}
		}
		return neighbours
	}([][]int{{i, j - 1}, {i, j + 1}, {i - 1, j}, {i + 1, j}})
}

func (t *tiles) compareParentAndChild(parent, child IPuzzleTiles) int {
	for i := 0; i < t.PuzzleSize; i++ {
		for j := 0; j < t.PuzzleSize; j++ {
			pval, err := parent.GetValueByIndexes(i, j)
			if err != nil {
				return -1
			}
			cval, err := child.GetValueByIndexes(i, j)
			if err != nil {
				return -1
			}
			if pval != cval {
				return -1
			}
		}
	}
	return 0
}

func (t *tiles) GenerateStates(parent IPuzzleTiles) ([]IPuzzleTiles, error) {
	children := make([]IPuzzleTiles, 0, 4)
	i, j, err := t.findCoordinates(0)
	if err != nil {
		return nil, err
	}
	neighbours := t.getValidNeighboursIndexes(i, j)
	for _, neighbour := range neighbours {
		children = append(
			children,
			&tiles{
				PuzzleData: t.generateNewChild(i, j, neighbour[0], neighbour[1]),
				PuzzleSize: t.PuzzleSize,
				Level:      t.Level + 1,
				Fval:       0,
			},
		)
	}

	if parent != nil {
		for i, child := range children {
			if t.compareParentAndChild(parent, child) == 0 {
				if i == 0 {
					return children[1:], nil
				} else if i == len(children)-1 {
					return children[0:i], nil
				}
				return append(children[0:i], children[i+1:]...), nil
			}
		}
	}
	return children, nil
}

func (t *tiles) SetFval(fval int) {
	t.Fval = fval
}

func (t *tiles) GetFval() int {
	return t.Fval
}

func (t *tiles) GetSize() int {
	return t.PuzzleSize
}

func (t *tiles) isSafeCoords(i, j int) error {
	if (i >= 0 && i < t.PuzzleSize) && (j >= 0 && j < t.PuzzleSize) {
		return nil
	}
	return fmt.Errorf("Indexes are out of range\n")
}

func (t *tiles) GetValueByIndexes(i, j int) (int, error) {
	err := t.isSafeCoords(i, j)
	if err != nil {
		return -1, err
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

func (t *tiles) PrintPuzzle() {
	for _, value := range t.PuzzleData {
		fmt.Println(value)
	}
}
