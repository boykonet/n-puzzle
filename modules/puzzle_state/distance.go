package puzzlestate

import (
	"math"
	"n-puzzle/modules/utils"
)

var DistanceFunctionNames = map[int]func(s, g IPuzzleState) int{
	utils.ManhattanHeuristic: ManhattanDistance,
	utils.EuclideanHeuristic: EuclideanDistance,
	utils.ChebyshevHeuristic: ChebyshevDistance,
}

func ManhattanDistance(s, g IPuzzleState) int {
	var manhattan int

	size := s.GetSize()
	for i := 0; i < size*size; i++ {
		y1, x1, _ := s.Coordinates(i)
		y2, x2, _ := g.Coordinates(i)
		xSteps := math.Abs(float64(x1 - x2))
		ySteps := math.Abs(float64(y1 - y2))
		manhattan += int(xSteps + ySteps)
	}
	return manhattan
}

func EuclideanDistance(s, g IPuzzleState) int {
	var euclidean float64

	size := s.GetSize()
	for i := 0; i < size*size; i++ {
		y1, x1, _ := s.Coordinates(i)
		y2, x2, _ := g.Coordinates(i)
		xDistance := math.Abs(float64(x1 - x2))
		yDistance := math.Abs(float64(y1 - y2))
		euclidean += math.Sqrt(xDistance*xDistance + yDistance*yDistance)
	}
	return int(euclidean)
}

func ChebyshevDistance(s, g IPuzzleState) int {
	var chebyshev int

	size := s.GetSize()
	for i := 0; i < size*size; i++ {
		y1, x1, _ := s.Coordinates(i)
		y2, x2, _ := g.Coordinates(i)
		xSteps := math.Abs(float64(x1 - x2))
		ySteps := math.Abs(float64(y1 - y2))
		chebyshev += int(math.Max(xSteps, ySteps))
	}
	return chebyshev
}
