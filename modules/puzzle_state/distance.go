package puzzlestate

import (
	"math"
	"n-puzzle/modules/utils"
)

var DistanceFunctionNames = map[int]func(s, g [][]int) int{
	utils.ManhattanHeuristic: ManhattanDistance,
	utils.EuclideanHeuristic: EuclideanDistance,
	utils.ChebyshevHeuristic: ChebyshevDistance,
}

func Coordinates(array [][]int, number int) (y, x int) {
	size := len(array)
	for y = 0; y < size; y++ {
		for x = 0; x < size; x++ {
			if array[y][x] == number {
				break
			}
		}
	}
	return y, x
}

func ManhattanDistance(s, g [][]int) int {
	var manhattan int

	size := len(s)
	for i := 0; i < size*size; i++ {
		y1, x1 := Coordinates(s, i)
		y2, x2 := Coordinates(g, i)
		if y1 == -1 || x1 == -1 || x2 == -1 || y2 == -1 {

		}
		xSteps := math.Abs(float64(x1 - x2))
		ySteps := math.Abs(float64(y1 - y2))
		manhattan += int(xSteps + ySteps)
	}
	return manhattan
}

func EuclideanDistance(s, g [][]int) int {
	var euclidean float64

	size := len(s)
	for i := 0; i < size*size; i++ {
		y1, x1 := Coordinates(s, i)
		y2, x2 := Coordinates(g, i)
		xDistance := math.Abs(float64(x1 - x2))
		yDistance := math.Abs(float64(y1 - y2))
		euclidean += math.Sqrt(xDistance*xDistance + yDistance*yDistance)
	}
	return int(euclidean)
}

func ChebyshevDistance(s, g [][]int) int {
	var chebyshev int

	size := len(s)
	for i := 0; i < size*size; i++ {
		y1, x1 := Coordinates(s, i)
		y2, x2 := Coordinates(g, i)
		xSteps := math.Abs(float64(x1 - x2))
		ySteps := math.Abs(float64(y1 - y2))
		chebyshev += int(math.Max(xSteps, ySteps))
	}
	return chebyshev
}
