package parser

import "testing"

func TestMapParser_ClearComments(t *testing.T) {
	mapWithTheComments := []string{
		"# this is a comment\n",
		"3 #another comment\n",
		"3 2 6\n",
		"1 4 0\n",
		"8 7 5",
	}
	result := []string{
		"3",
		"3 2 6",
		"1 4 0",
		"8 7 5",
	}

	mapParser := NewMapParser()
	mapParser.Parse(mapWithTheComments)
}
