package puzzlesolver

type solver struct {
	Data   [][]int
	Open   [][]int
	Closed [][]int
}

func NewPuzzleSolver(data [][]int, open, closed [][]int) IPuzzleSolver {
	return &solver{
		Data:   data,
		Open:   open,
		Closed: closed,
	}
}

func (ps *solver) Solve() {

}
