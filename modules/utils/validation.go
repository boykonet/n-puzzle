package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ArgTypeKey int

const (
	ReadingModeArgumentType ArgTypeKey = iota
	FilenameArgumentType
)

const (
	ReadingModeStdin string = "stdin"
	ReadingModeFile  string = "file"
)

var (
	// ErrorVDNotEnoughInfo Validated data function
	ErrorVDNotEnoughInfo = errors.New("validateData: Not less than 4 rows, comments excluded")
	// ErrorVDIncorrectPuzzleSize Validated data function
	ErrorVDIncorrectPuzzleSize = errors.New("validateData: Incorrect puzzle size")
	// ErrorVDIncorrectCountElemsInRow Validated data function
	ErrorVDIncorrectCountElemsInRow = errors.New("validateData: Incorrect count elements in row")
	// ErrorVDIncorrectPuzzleNumber Validated data function
	ErrorVDIncorrectPuzzleNumber = errors.New("validateData: Incorrect puzzle number")
	// ErrorVDRepeatedNumber Validated data function
	ErrorVDRepeatedNumber = errors.New("validateData: Repeated number")
)

func ValidateArgs() (map[ArgTypeKey]string, error) {
	data := make(map[ArgTypeKey]string, 0)
	if len(os.Args) == 1 {
		data[ReadingModeArgumentType] = ReadingModeStdin
	} else if len(os.Args) == 2 {
		data[ReadingModeArgumentType] = ReadingModeFile
		data[FilenameArgumentType] = os.Args[1]
		//file := strings.Split(os.Args[1], "=")
		//if len(file) != 2 || file[0] != "file" {
		//	return nil, fmt.Errorf("Incorrect argument: incorrect key-value pair\n")
		//} else {
		//	data[FilenameArgumentType] = file[1]
		//}
	} else {
		return nil, fmt.Errorf("Incorrect argument: the program allow only 1 added argument\n")
	}
	return data, nil
}

func ValidateInputData(data []string) error {
	if len(data) < 4 {
		return ErrorVDNotEnoughInfo
	}
	puzzleSize, err := strconv.ParseInt(data[0], 10, 32)
	if err != nil {
		return fmt.Errorf("validate data: %v\n", err)
	}
	if len(data)-1 != int(puzzleSize) {
		return ErrorVDIncorrectPuzzleSize
	}
	data = data[1:]
	m := make(map[int64]bool, 0)
	for _, str := range data {
		sstrings := strings.FieldsFunc(
			str,
			func(c rune) bool {
				return c == ' '
			},
		)
		if len(sstrings) != int(puzzleSize) {
			return ErrorVDIncorrectCountElemsInRow
		}
		for i := 0; i < int(puzzleSize); i++ {
			number, err := strconv.ParseInt(sstrings[i], 10, 32)
			if err != nil {
				return fmt.Errorf("validate data: %v\n", err)
			}
			if number >= puzzleSize*puzzleSize || number < 0 {
				return ErrorVDIncorrectPuzzleNumber
			}
			_, ok := m[number]
			if ok == true {
				return ErrorVDRepeatedNumber
			}
			m[number] = true
		}
	}
	return nil
}
