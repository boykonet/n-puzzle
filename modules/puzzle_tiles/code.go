package puzzletiles

import "n-puzzle/modules/utils"

type tiles struct {
	Data  [][]int
	Level int
	Fval  int
}

func NewPuzzleTiles(data [][]int, level int, fval int) IPuzzleTiles {
	return &tiles{
		Data:  utils.Duplicate2Darray(data),
		Level: level,
		Fval:  fval,
	}
}

func (t *tiles) GenerateChild() {

}
