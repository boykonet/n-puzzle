package puzzlestate

type IPuzzleState interface {
	//Actions() []IPuzzleState
	SetFval(int)
	GetValueByIndexes(i, j int) (int, error)
	GetLevel() int
	GetFval() int
	GetSize() int
	Encrypt() string
	ConvertToArray() []int
	PrintPuzzle()
	Coordinates(number int) (int, int, error)
	CopyPuzzle() [][]int
	ListOfStates() []IPuzzleState
}
