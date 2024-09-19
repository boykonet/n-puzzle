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
	// ErrorVDIncorrectPuzzleSize Incorrect puzzle size
	ErrorVDIncorrectPuzzleSize = errors.New("incorrect puzzle size")
	// ErrorVDIncorrectPuzzleNumber Validated data function
	ErrorVDIncorrectPuzzleNumber = errors.New("validateData: Incorrect puzzle number")
	// ErrorVDRepeatedNumber Repeated number in the map
	ErrorVDRepeatedNumber = errors.New("repeated number in the map")
	// ErrorIncorrectFileExtension Incorrect file extension (.txt)
	ErrorIncorrectFileExtension = errors.New("incorrect file extension")

	ErrorIncorrectAmountOfArgs = errors.New("incorrect amount of arguments")
)

func checkFileExtension(filename string) error {
	s := strings.Split(filename, ".")
	if len(s) > 0 && s[len(s)-1] != "txt" {
		return ErrorIncorrectFileExtension
	}
	return nil
}

func CheckAndReturnArgs() (map[ArgTypeKey]string, error) {
	data := make(map[ArgTypeKey]string)
	switch len(os.Args) {
	case 1:
		data[ReadingModeArgumentType] = ReadingModeStdin
	case 2:
		data[ReadingModeArgumentType] = ReadingModeFile
		filename := os.Args[1]
		err := checkFileExtension(filename)
		if err != nil {
			return nil, err
		}
		data[FilenameArgumentType] = filename
	default:
		return nil, ErrorIncorrectAmountOfArgs
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
			return ErrorVDIncorrectPuzzleSize
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
