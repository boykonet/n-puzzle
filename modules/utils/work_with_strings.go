package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//func PrintSliceStrings(ss []string) {
//	for i := 0; i < len(ss); i++ {
//		fmt.Printf("string: [%v], len string: [%v]\n", ss[i], len(ss[i]))
//	}
//}

//func RemoveComments(data []string) {
//	for i, str := range data {
//		index := strings.Index(str, "#")
//		if index == -1 {
//			continue
//		}
//		data[i] = str[0:index]
//		data[i] = strings.Trim(data[i], " ")
//	}
//}

//func RemoveEmptyStringsFromSlice(data []string, remove string) []string {
//	var res []string
//	for i := 0; i < len(data); i++ {
//		if data[i] == remove {
//			continue
//		}
//		res = append(res, data[i])
//	}
//	return res
//}

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

func ReadFromStdin() ([]string, error) {
	return nil, fmt.Errorf("Stdin doesn't support\n")
}
