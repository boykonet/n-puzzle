package parser

import (
	"errors"
	"testing"
)

func TestMapParser_CorrectMap(t *testing.T) {
	mapWithTheComments := []string{
		"# this is a comment\n",
		"3 #another comment\n",
		"3 2 6\n",
		"1 4 0\n",
		"8 7 5",
	}
	output := [][]int{
		{3, 2, 6},
		{1, 4, 0},
		{8, 7, 5},
	}

	myParser := NewMapParser()
	result, err := myParser.Parse(mapWithTheComments)
	if err != nil {
		return
	}
	if len(result) != len(output) {
		t.Fail()
	}

	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output); j++ {
			if result[i][j] != output[i][j] {
				t.Fail()
			}
		}
	}
}

func TestMapParser_IncorrectMap_WrongSize(t *testing.T) {
	mapWithTheComments := []string{
		"# this is a comment\n",
		"4 #another comment\n",
		"3 2 6\n",
		"1 4 0\n",
		"8 7 5\n",
	}

	myParser := NewMapParser()
	_, err := myParser.Parse(mapWithTheComments)
	if err != nil {
		if !errors.Is(err, ErrorIncorrectMapSize) {
			t.Fail()
		}
	}
}

func TestMapParser_IncorrectMap_RepeatedNumber(t *testing.T) {
	mapWithTheComments := []string{
		"# this is a comment\n",
		"3 #another comment\n",
		"3 2 6\n",
		"1 4 0\n",
		"8 6 5 # 6 is repeated number\n",
	}

	myParser := NewMapParser()
	_, err := myParser.Parse(mapWithTheComments)
	if err != nil {
		if !errors.Is(err, ErrorIncorrectMapValue) {
			t.Fail()
		}
	}
}

func TestMapParser_IncorrectMap_NoZero(t *testing.T) {
	mapWithTheComments := []string{
		"# this is a comment\n",
		"3 #another comment\n",
		"3 2 6\n",
		"1 4 9 # 9 is too much\n",
		"8 7 5\n",
	}

	myParser := NewMapParser()
	_, err := myParser.Parse(mapWithTheComments)
	if err != nil {
		if !errors.Is(err, ErrorIncorrectMapValue) {
			t.Fail()
		}
	}
}

func TestMapParser_IncorrectMap_AlphabeticSymbolInMap(t *testing.T) {
	mapWithTheComments := []string{
		"# this is a comment\n",
		"3 #another comment\n",
		"3 2 6\n",
		"1 4a 0\n",
		"8 7 5\n",
	}

	myParser := NewMapParser()
	_, err := myParser.Parse(mapWithTheComments)
	if err != nil {
		if !errors.Is(err, ErrorIncorrectMapValue) {
			t.Fail()
		}
	}
}

func TestMapParser_IncorrectMap_AlphabeticSymbolInSize(t *testing.T) {
	mapWithTheComments := []string{
		"# this is a comment\n",
		"3nnn #another comment\n",
		"3 2 6\n",
		"1 4 0\n",
		"8 7 5\n",
	}

	myParser := NewMapParser()
	_, err := myParser.Parse(mapWithTheComments)
	if err != nil {
		if !errors.Is(err, ErrorIncorrectMapValue) {
			t.Fail()
		}
	}

	mapWithTheComments2 := []string{
		"# this is a comment\n",
		"3 nnn #another comment\n",
		"3 2 6\n",
		"1 4 0\n",
		"8 7 5\n",
	}

	_, err = myParser.Parse(mapWithTheComments2)
	if err != nil {
		if !errors.Is(err, ErrorIncorrectMapValue) {
			t.Fail()
		}
	}
}
