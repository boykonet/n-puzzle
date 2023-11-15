package puzzlesolver

import (
	"fmt"
	puzzletiles "n-puzzle/modules/puzzle_tiles"
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

func (ps *solver) h(open, closed puzzletiles.IPuzzleTiles) (int, error) {
	temp := 0
	for i := 0; i < ps.Size; i++ {
		for j := 0; j < ps.Size; j++ {
			openValue, err1 := open.GetValueByIndexes(i, j)
			goalValue, err2 := closed.GetValueByIndexes(i, j)
			if err1 != nil || err2 != nil {
				return 0, err1
			}
			if openValue != goalValue && openValue != 0 {
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
		curr := ps.Parent
		h, _ := ps.h(curr, ps.Goal)
		if h == 0 {
			break
		}
		children, _ := curr.GenerateChild()
		for _, child := range children {
			fval, err = ps.calcHeuristicVal(child, ps.Goal)
			if err != nil {

			}
			child.SetFval(fval)
			ps.appendChild(child)
		}
		ps.appendClosed(curr)
		ps.Parent = ps.Children[ps.bestChildIndex()]
		ps.Children = ps.Children[:0]
	}
	fmt.Println(ps.Parent)

}
