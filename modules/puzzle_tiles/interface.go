package puzzletiles

type IPuzzleTiles interface {
	GenerateChild() ([][][]int, error)
}
