package puzzlesolver

type IPuzzleSolver interface {
	Solve(initialState, goalState [][]int, heurictic int) (bool, error)
}
