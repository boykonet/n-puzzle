package parser

import (
	"errors"
	"strconv"
	"strings"
)

type mapParser struct {
}

func NewMapParser() IMapParser {
	return &mapParser{}
}

func (mp *mapParser) Parse(puzzleMap []string) ([][]int, error) {
	mp.RemoveComments(puzzleMap)
	puzzleMap = mp.RemoveEmptyStrings(puzzleMap)
	mp.TrimRows(puzzleMap)

	size, err := strconv.Atoi(puzzleMap[0])
	if err != nil {
		return nil, err
	}

	// remove the string with the map size from map
	puzzleMap = puzzleMap[1:]

	matrix, err := mp.ConvertToIntMatrix(puzzleMap)
	if err != nil {
		return nil, err
	}
	err = mp.Validate(matrix, size)
	if err != nil {
		return nil, err
	}
	return matrix, nil
}

func (mp *mapParser) RemoveComments(data []string) /*[]string*/ {
	for i, str := range data {
		index := strings.Index(str, "#")
		if index == -1 {
			continue
		}
		data[i] = str[:index]
	}
	//return data
}

func (mp *mapParser) TrimRows(puzzleMap []string) /*[]string*/ {
	for index, row := range puzzleMap {
		row = strings.Trim(row, " ")
		puzzleMap[index] = row
	}
	//return puzzleMap
}

func (mp *mapParser) RemoveEmptyStrings(data []string) []string {
	//var res []string
	size := len(data)
	for i := 0; i < size; {
		if data[i] == "" {
			nd := append(data[:i], data[i+1:]...)
			data = append([]string{}, nd...)
			size--
			continue
		}
		i++
		//res = append(res, data[i])
	}
	return data
	//data = res
	//return res
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

var (
	ErrorIncorrectMapSize  = errors.New("incorrect map size")
	ErrorIncorrectMapValue = errors.New("incorrect value in the map")
)

func (mp *mapParser) Validate(puzzleMap [][]int, size int) error {
	elementsInMap := make(map[int]struct{})
	maxPossibleValue := size * size
	for _, row := range puzzleMap {
		if len(row) != size {
			return ErrorIncorrectMapSize
		}
		for _, value := range row {
			_, ok := elementsInMap[value]
			if value < 0 || value >= maxPossibleValue || ok == true {
				return ErrorIncorrectMapValue
			}
			elementsInMap[value] = struct{}{}
		}
	}
	return nil
}
