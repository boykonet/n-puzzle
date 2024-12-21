package puzzlesolver

type IPuzzleSolver interface {
	Solve(initialState, goalState [][]int, hFunction func(s, g [][]int) int, filepath string) (bool, error)
}
