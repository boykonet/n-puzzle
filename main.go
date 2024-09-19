package main

import (
	"errors"
	"fmt"
	parser "n-puzzle/modules/parser"
	puzzleSolver "n-puzzle/modules/puzzle_solver"
	"n-puzzle/modules/utils"
)

var (
	ErrorIncorrectReadingMode = errors.New("Not supported reading mode\n")
)

func ReadAndParseMap() ([][]int, error) {
	data, err := utils.CheckAndReturnArgs()
	if err != nil {
		return nil, err
	}

	var ss []string
	if data[utils.ReadingModeArgumentType] == utils.ReadingModeFile {
		ss, err = utils.ReadFromFile(data[utils.FilenameArgumentType])
		if err != nil {
			return nil, err
		}
	} else if data[utils.ReadingModeArgumentType] == utils.ReadingModeStdin {
		_, err = utils.ReadFromStdin()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, ErrorIncorrectReadingMode
	}

	customParser := parser.NewMapParser()

	return customParser.Parse(ss)
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
	startPuzzle, err := ReadAndParseMap()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("start puzzle: ", startPuzzle)

	puzzleSolver := puzzleSolver.NewPuzzleSolver()
	ok, err := puzzleSolver.Solve(startPuzzle, generateGoalStateMatrix(len(startPuzzle)))
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
