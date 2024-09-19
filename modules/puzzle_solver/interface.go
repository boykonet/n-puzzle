package puzzlesolver

type IPuzzleSolver interface {
	Solve(initialState, goalState [][]int) (bool, error)
}
