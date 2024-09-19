package puzzletiles

type IPuzzleTiles interface {
	GenerateStates(parent IPuzzleTiles) ([]IPuzzleTiles, error)
	SetFval(int)
	GetValueByIndexes(i, j int) (int, error)
	GetLevel() int
	GetFval() int
	GetSize() int
	PrintPuzzle()
}
