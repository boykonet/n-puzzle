package main

import (
	"fmt"
	parser "n-puzzle/modules/parser"
	puzzlesolver "n-puzzle/modules/puzzle_solver"
	puzzlestate "n-puzzle/modules/puzzle_state"
	"n-puzzle/modules/utils"
	"os"
	"strings"
)

const (
	outputDirectoryName = "./output/"
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
		fmt.Println("Error:", err)
		return
	}
	hFunc, ok := puzzlestate.DistanceFunctionNames[heuristic]
	if ok == false {
		fmt.Println(utils.ErrorHeuristic)
	}

	fmt.Println("Start puzzle: ", startPuzzle)

	// Create the path for output file
	outputFilePath := os.Args[1]
	outputFilePath = strings.Trim(outputFilePath, "./")

	outputFilePath = outputDirectoryName + outputFilePath

	// Create the puzzle solver object
	puzzleSolver := puzzlesolver.NewPuzzleSolver()
	ok, err = puzzleSolver.Solve(startPuzzle, generateGoalStateMatrix(len(startPuzzle)), hFunc, outputFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if ok == true {
		fmt.Println("The current puzzle is SOLVABLE")
	} else {
		fmt.Println("The current puzzle is NOT SOLVABLE")
	}
	fmt.Println("You can find the information into the", outputFilePath)
}
