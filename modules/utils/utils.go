package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Duplicate2DArray[T int](data [][]T) [][]T {
	duplicate := make([][]T, len(data))
	for i := range data {
		duplicate[i] = make([]T, len(data[i]))
		copy(duplicate[i], data[i])
	}
	return duplicate
}

func Swap[T int](first *T, second *T) {
	if first == nil || second == nil {
		return
	}
	tmp := *first
	*first = *second
	*second = tmp
}

func ReadFromFile(filename string) ([]string, error) {
	var res []string

	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() == true {
		return nil, fmt.Errorf("the file is a directory")
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		str = strings.Trim(str, "\n")
		res = append(res, str)
	}
	return res, nil
}

func ConvertPuzzleToArrayOfStrings(puzzle [][]int) []string {
	convertedPuzzle := make([]string, 0)

	size := len(puzzle)
	lenMaxElem := len(strconv.Itoa(size*size-1)) + 2
	for index, row := range puzzle {
		convertedPuzzle = append(convertedPuzzle, "")
		for _, elem := range row {
			pattern := strconv.Itoa(elem)
			for len(pattern) < lenMaxElem {
				pattern = " " + pattern
			}
			convertedPuzzle[index] += pattern
		}
		convertedPuzzle[index] += "\n"
	}
	return convertedPuzzle
}

func CreateDirectories(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0700)
		if err != nil {
			return err
		}
	}
	return nil
}
