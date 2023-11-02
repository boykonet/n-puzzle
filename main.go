package main

import (
	"fmt"
	"os"
	"strings"
)

type ArgTypeKey int

const (
	ReadingModeArgumentType ArgTypeKey = 0
	FilenameArgumentType
)

const (
	ReadingModeStdin string = "stdin"
	ReadingModeFile  string = "file"
)

func readPuzzle(size int) ([][]string, error) {
	var tile string
	var puzzle [][]string
	for i := 0; i < size; i++ {
		var line []string
		for j := 0; j < size; j++ {
			_, err := fmt.Scanf("%v", &tile)
			if err != nil {
				return nil, fmt.Errorf("Incorrect map: wrong data\n")
			}
			line = append(line, tile)
		}
		puzzle = append(puzzle, line)
	}
	return puzzle, nil
}

func validateArgs() (map[ArgTypeKey]string, error) {
	data := make(map[ArgTypeKey]string, 2)
	if len(os.Args) == 1 {
		data[ReadingModeArgumentType] = ReadingModeStdin
	} else if len(os.Args) == 2 {
		data[ReadingModeArgumentType] = ReadingModeFile
		file := strings.Split(os.Args[1], "=")
		if len(file) != 2 || file[0] != "file" {
			return nil, fmt.Errorf("Incorrect argument: incorrect key-value pair\n")
		} else {
			data[FilenameArgumentType] = file[1]
		}
	} else {
		return nil, fmt.Errorf("Incorrect argument: the program allow only 1 added argument\n")
	}
	return data, nil
}

func main() {
	data, err := validateArgs()
	if err != nil {
		fmt.Print(err)
		return
	}
	if data[ReadingModeArgumentType] == ReadingModeFile {
		//f, err := os.Open(string(data[FilenameArgumentType]))
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//reader := bufio.NewReader(f)
		//defer f.Close()
	} else if data[ReadingModeArgumentType] == ReadingModeFile {

	}
}
