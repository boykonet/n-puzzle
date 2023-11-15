package puzzletiles

type IPuzzleTiles interface {
	GenerateChild() ([]IPuzzleTiles, error)
	SetFval(int)
	GetValueByIndexes(i, j int) (int, error)
	GetLevel() int
	GetFval() int
}
