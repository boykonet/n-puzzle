package main

import (
	"fmt"
	puzzlesolver "n-puzzle/modules/puzzle_solver"
	"n-puzzle/modules/utils"
	"strconv"
	"strings"
)

func convertStringsToInts(data []string) ([][]int, error) {
	matrix := make([][]int, 0, len(data))
	for index, str := range data {
		sstrings := strings.FieldsFunc(
			str,
			func(c rune) bool {
				return c == ' '
			},
		)
		matrix = append(matrix, []int{})
		for _, v := range sstrings {
			number, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("from strings to ints: %v", err)
			}
			matrix[index] = append(matrix[index], number)
		}
	}
	return matrix, nil
}

func parsingAndValidation() ([][]int, error) {
	data, err := utils.ValidateArgs()
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
		err = utils.ReadFromStdin()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Not supported reading mode\n")
	}
	utils.RemoveComments(ss)
	ss = utils.RemoveStringsFromSlice(ss, "")
	//utils.PrintSliceStrings(ss)
	err = utils.ValidateInputData(ss)
	if err != nil {
		return nil, err
	}
	array, err := convertStringsToInts(ss[1:])
	if err != nil {
		return nil, err
	}
	return array, nil
}

func genGoalStateMatrix(size int) [][]int {
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
	startPuzzle, err := parsingAndValidation()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("start puzzle: ", startPuzzle)

	puzzleSolver := puzzlesolver.NewPuzzleSolver(len(startPuzzle), startPuzzle, genGoalStateMatrix(len(startPuzzle)))
	puzzleSolver.Solve()
}
