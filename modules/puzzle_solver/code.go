package puzzlesolver

import (
	"n-puzzle/modules/utils"
)

type solver struct {
	data [][]int
}

func NewPuzzleSolver(data [][]int) IPuzzleSolver {
	return &solver{
		data: utils.Duplicate2Darray[int](data),
	}
}

func (ps *solver) Solve() {

}
