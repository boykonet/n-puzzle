package puzzlesolver

import (
	"fmt"
	puzzletiles "n-puzzle/modules/puzzle_tiles"
	"time"
)

type solver struct {
	Size     int
	Parent   puzzletiles.IPuzzleTiles
	Children []puzzletiles.IPuzzleTiles
	Closed   []puzzletiles.IPuzzleTiles
	Goal     puzzletiles.IPuzzleTiles
}

func NewPuzzleSolver(size int, parent, goal [][]int) IPuzzleSolver {
	return &solver{
		Size:     size,
		Parent:   puzzletiles.NewPuzzleTiles(parent, size, 0, 0),
		Children: make([]puzzletiles.IPuzzleTiles, 0, 4),
		Closed:   make([]puzzletiles.IPuzzleTiles, 0),
		Goal:     puzzletiles.NewPuzzleTiles(goal, size, 0, 0),
	}
}

func (ps *solver) appendChild(child puzzletiles.IPuzzleTiles) {
	ps.Children = append(ps.Children, child)
}

func (ps *solver) appendClosed(closed puzzletiles.IPuzzleTiles) {
	ps.Closed = append(ps.Closed, closed)
}

func (ps *solver) h(start, goal puzzletiles.IPuzzleTiles) (int, error) {
	temp := 0
	for i := 0; i < ps.Size; i++ {
		for j := 0; j < ps.Size; j++ {
			startValue, err1 := start.GetValueByIndexes(i, j)
			goalValue, err2 := goal.GetValueByIndexes(i, j)
			if err1 != nil || err2 != nil {
				return 0, err1
			}
			if startValue != goalValue && startValue != 0 {
				temp += 1
			}
		}
	}
	return temp, nil
}

func (ps *solver) bestChildIndex() int {
	min := ps.Children[0].GetFval()
	index := 0
	for i, child := range ps.Children {
		if min > child.GetFval() {
			min = child.GetFval()
			index = i
		}
	}
	return index
}

func (ps *solver) calcHeuristicVal(open, closed puzzletiles.IPuzzleTiles) (int, error) {
	heuristic, err := ps.h(open, closed)
	if err != nil {
		return 0, err
	}
	return heuristic + open.GetLevel(), nil
}

func (ps *solver) Solve() {
	fval, err := ps.calcHeuristicVal(ps.Parent, ps.Goal)
	if err != nil {

	}
	ps.Parent.SetFval(fval)
	for {
		ps.Parent.PrintPuzzle()
		fmt.Println()
		h, err := ps.h(ps.Parent, ps.Goal)
		if err != nil {
			fmt.Print(err)
			return
		}
		if h == 0 {
			break
		}
		//curr := ps.Parent
		children, err := ps.Parent.GenerateChild()
		if err != nil {
			fmt.Print(err)
			return
		}
		counter := 0
		for _, child := range children {
			fval, err = ps.calcHeuristicVal(child, ps.Goal)
			if err != nil {
				fmt.Print(err)
				return
			}
			child.SetFval(fval)
			ps.appendChild(child)
			counter += 1
		}
		fmt.Println("counter = ", counter)
		ps.appendClosed(ps.Parent)
		ps.Parent = ps.Children[ps.bestChildIndex()]
		ps.Children = ps.Children[0:0]
		time.Sleep(2 * time.Second)
	}
	fmt.Println(ps.Parent)

}
