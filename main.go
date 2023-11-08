package main

import (
	"fmt"
	"n-puzzle/modules/utils"
	"strconv"
	"strings"
)

func fromStringsToInts(data []string) ([][]int, error) {
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
	}
	utils.RemoveComments(ss)
	ss = utils.RemoveStringsFromSlice(ss, "")
	//utils.PrintSliceStrings(ss)
	err = utils.ValidateInputData(ss)
	if err != nil {
		return nil, err
	}
	array, err := fromStringsToInts(ss[1:])
	if err != nil {
		return nil, err
	}
	return array, nil
}

func main() {
	array, err := parsingAndValidation()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(array)
}
