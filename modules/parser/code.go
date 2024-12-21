package parser

import (
	"errors"
	"strconv"
	"strings"
)

type mapParser struct {
}

var (
	ErrorIncorrectMapSize  = errors.New("incorrect map size")
	ErrorIncorrectMapValue = errors.New("incorrect value in the map")
	ErrorEmptyFile         = errors.New("empty file")
)

func NewMapParser() IMapParser {
	return &mapParser{}
}

// Parse TODO: doesn't work with the map without empty line in the end of the map
func (mp *mapParser) Parse(puzzleMap []string) ([][]int, error) {
	mp.RemoveComments(puzzleMap)
	mp.TrimRows(puzzleMap)
	puzzleMap = mp.RemoveEmptyStrings(puzzleMap)

	// The len of the array is checking after RemoveComments(..),
	// RemoveEmptyString(...) and TrimRow(...) for removing unnecessary information from the file
	if len(puzzleMap) == 0 {
		return nil, ErrorEmptyFile
	}

	size, err := strconv.Atoi(puzzleMap[0])
	if err != nil {
		return nil, ErrorIncorrectMapValue
	}

	// remove the map size
	puzzleMap = puzzleMap[1:]

	matrix, err := mp.ConvertToIntMatrix(puzzleMap)
	if err != nil {
		return nil, ErrorIncorrectMapValue
	}
	err = mp.Validate(matrix, size)
	if err != nil {
		return nil, err
	}
	return matrix, nil
}

func (mp *mapParser) RemoveComments(data []string) {
	for i, str := range data {
		index := strings.Index(str, "#")
		if index == -1 {
			continue
		}
		data[i] = str[:index]
	}
}

func (mp *mapParser) TrimRows(puzzleMap []string) {
	for index, row := range puzzleMap {
		row = strings.Trim(row, " \n")
		puzzleMap[index] = row
	}
}

func (mp *mapParser) RemoveEmptyStrings(data []string) []string {
	size := len(data)
	for i := 0; i < size; {
		if data[i] == "" {
			nd := append(data[:i], data[i+1:]...)
			data = append([]string{}, nd...)
			size--
			continue
		}
		i++
	}
	return data
}

func (mp *mapParser) ConvertToIntMatrix(puzzleMap []string) ([][]int, error) {
	matrix := make([][]int, 0, len(puzzleMap))
	for index, str := range puzzleMap {
		sstrings := strings.FieldsFunc(
			str,
			func(c rune) bool {
				return c == ' '
			},
		)
		mp.RemoveEmptyStrings(sstrings)

		matrix = append(matrix, []int{})
		for _, v := range sstrings {
			number, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			matrix[index] = append(matrix[index], number)
		}
	}
	return matrix, nil
}

func (mp *mapParser) Validate(puzzleMap [][]int, size int) error {
	if len(puzzleMap) != size {
		return ErrorIncorrectMapSize
	}
	elementsInMap := make(map[int]struct{})
	maxPossibleValue := size*size - 1
	for _, row := range puzzleMap {
		if len(row) != size {
			return ErrorIncorrectMapSize
		}
		for _, value := range row {
			_, ok := elementsInMap[value]
			if value < 0 || value > maxPossibleValue || ok == true {
				return ErrorIncorrectMapValue
			}
			elementsInMap[value] = struct{}{}
		}
	}
	return nil
}
