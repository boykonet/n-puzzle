package parser

type IMapParser interface {
	Parse(puzzleMap []string) ([][]int, error)
}
