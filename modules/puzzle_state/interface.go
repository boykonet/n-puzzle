package puzzlestate

type IPuzzleState interface {
	GetValueByIndexes(i, j int) (int, error)
	GetLevel() int
	GetFval() int
	GetSize() int
	Encrypt() string
	ConvertToArray() []int
	PrintPuzzle()
	Coordinates(number int) (y int, x int, e error)
	CopyMatrix() [][]int
	ListOfStates() []IPuzzleState
}
