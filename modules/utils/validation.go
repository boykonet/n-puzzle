package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type ArgTypeKey int

const (
	_ ArgTypeKey = iota
	FilenameArgument
	HeuristicFunctionArgument
)

const (
	_ int = iota
	ManhattanHeuristic
	EuclideanHeuristic
	ChebyshevHeuristic
)

var HeuristicNames = map[string]int{
	"manhattan": ManhattanHeuristic,
	"euclidean": EuclideanHeuristic,
	"chebyshev": ChebyshevHeuristic,
}

var (
	// ErrorIncorrectMap Incorrect map
	ErrorIncorrectMap = errors.New("incorrect map")
	// ErrorIncorrectAmountOfRows Incorrect amount of rows
	ErrorIncorrectAmountOfRows    = errors.New("incorrect amount of rows")
	ErrorIncorrectAmountOfColumns = errors.New("incorrect amount of columns")
	// ErrorIncorrectFileExtension Incorrect file extension (.txt)
	ErrorIncorrectFileExtension = errors.New("Incorrect file extension.\nSupported name is .txt")

	ErrorIncorrectAmountOfArgs = errors.New("Incorrect amount of arguments.\nExpected 2 arguments: the filename and the name of the heuristic function")

	ErrorHeuristic = errors.New("Incorrect name of the heuristic function.\nThe supported names are:\n    - manhattan\n    - euclidean\n    - chebyshev")

	ErrorInfoMessage = errors.New("The program takes 2 arguments:\n1. The filename with .txt extension\n2. The name of the heuristic function:\n    - manhattan\n    - euclidean\n    - chebyshev")
)

func checkFileExtension(filename string) error {
	s := strings.Split(filename, ".")
	if len(s) > 0 && s[len(s)-1] != "txt" {
		return ErrorIncorrectFileExtension
	}
	return nil
}

// CheckAndReturnArgs
// ./program_name                   - info message
// ./program_name map.txt manhattan - correct input
func CheckAndReturnArgs() (map[ArgTypeKey]string, error) {
	data := make(map[ArgTypeKey]string)
	switch len(os.Args) {
	case 1:
		return nil, ErrorInfoMessage
	case 2:
		return nil, ErrorIncorrectAmountOfArgs
	case 3:
		filename := os.Args[1]
		err := checkFileExtension(filename)
		if err != nil {
			return nil, err
		}
		data[FilenameArgument] = filename
		heuristic := os.Args[2]
		_, ok := HeuristicNames[heuristic]
		if ok == false {
			return nil, ErrorHeuristic
		}
		data[HeuristicFunctionArgument] = heuristic
	default:
		return nil, ErrorIncorrectAmountOfArgs
	}
	return data, nil
}

func ValidateInputData(data []string) error {
	if len(data) < 4 {
		return ErrorIncorrectMap
	}
	puzzleSize, err := strconv.ParseInt(data[0], 10, 32)
	if err != nil {
		return err
	}
	if len(data)-1 != int(puzzleSize) {
		return ErrorIncorrectAmountOfRows
	}
	data = data[1:]
	m := make(map[int64]struct{})
	for _, str := range data {
		sstrings := strings.FieldsFunc(
			str,
			func(c rune) bool {
				return c == ' '
			},
		)
		if len(sstrings) != int(puzzleSize) {
			return ErrorIncorrectAmountOfColumns
		}
		for i := 0; i < int(puzzleSize); i++ {
			number, err := strconv.ParseInt(sstrings[i], 10, 32)
			if err != nil {
				return err
			}
			if number >= puzzleSize*puzzleSize || number < 0 {
				return ErrorIncorrectMap
			}
			_, ok := m[number]
			if ok == true {
				return ErrorIncorrectMap
			}
			m[number] = struct{}{}
		}
	}
	return nil
}
