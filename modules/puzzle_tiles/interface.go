package puzzletiles

type IPuzzleTiles interface {
	GenerateChild() ([]IPuzzleTiles, error)
}
