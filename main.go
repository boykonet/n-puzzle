package main

import (
	"fmt"
	parser "n-puzzle/modules/parser"
	puzzleSolver "n-puzzle/modules/puzzle_solver"
	"n-puzzle/modules/utils"
)

func ReadAndParseMap() ([][]int, int, error) {
	args, err := utils.CheckAndReturnArgs()
	if err != nil {
		return nil, 0, err
	}

	rawMap, err := utils.ReadFromFile(args[utils.FilenameArgument])
	if err != nil {
		return nil, 0, err
	}

	heuristic := utils.HeuristicNames[args[utils.HeuristicFunctionArgument]]

	customParser := parser.NewMapParser()

	parsedMap, err := customParser.Parse(rawMap)
	return parsedMap, heuristic, err
}

func generateGoalStateMatrix(size int) [][]int {
	arr := make([][]int, size)
	counter := 1
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == size-1 && j == size-1 {
				arr[i] = append(arr[i], 0)
			} else {
				arr[i] = append(arr[i], counter)
			}
			counter += 1
		}
	}
	return arr
}

func main() {
	startPuzzle, heuristic, err := ReadAndParseMap()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("start puzzle: ", startPuzzle)

	puzzleSolver := puzzleSolver.NewPuzzleSolver()
	ok, err := puzzleSolver.Solve(startPuzzle, generateGoalStateMatrix(len(startPuzzle)), heuristic)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok == true {
		fmt.Println("The current puzzle is solvable")
	} else {
		fmt.Println("The current puzzle isn't solvable")
	}
}
