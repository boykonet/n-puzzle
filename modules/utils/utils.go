package utils

import (
	"bufio"
	"io"
	"os"
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
