package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"sudoku-solver/observer"
)

const rowSequenceSeparator = "."
const lineSequenceSeparator = ","

var sequencePatternReg, _ = regexp.Compile("^[0-9,.]*$")

type Grid struct {
	c         [9][9]uint
	observers []observer.IObserver
}

func NewGrid() Grid {
	return Grid{
		c: [9][9]uint{},
	}
}

// NewGridFromSequence
// Expected sequence format :
//
//	1,2,3,4,5,6,7,8,9.2,3,4,5,6,7,8...
//
// Each line separator : .
// Each column separator : ,
//
// Format describes the grid row by row.
func NewGridFromSequence(sequence string) (Grid, error) {
	grid := NewGrid()
	if err := checkSequenceLength(sequence); err != nil {
		return grid, err
	}
	rows := strings.Split(sequence, rowSequenceSeparator)
	for i, row := range rows {
		cells := strings.Split(row, lineSequenceSeparator)
		for j, cell := range cells {
			v, err := strconv.Atoi(cell)
			if err != nil {
				return grid, Error{e: err}
			}
			grid.SetValue(uint(i), uint(j), uint(v))
		}
	}
	return grid, nil
}

func (g *Grid) IsInRow(row, number uint) bool {
	for _, n := range g.c[row] {
		if number == n {
			return true
		}
	}
	return false
}

func (g *Grid) IsInColumn(column, number uint) bool {
	for _, row := range g.c {
		if row[column] == number {
			return true
		}
	}
	return false
}

func (g *Grid) IsInBloc(row, column, number uint) bool {
	iOffset := row - (row % 3)
	jOffset := column - (column % 3)
	for i := iOffset; i < iOffset+3; i++ {
		for j := jOffset; j < jOffset+3; j++ {
			if g.c[i][j] == number {
				return true
			}
		}
	}
	return false
}

func (g *Grid) SetValue(row, column, v uint) {
	g.c[row][column] = v
	g.Notify()
}

func (g *Grid) PrintInConsole() {
	for _, row := range g.c {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Print("\n")
	}
}

func checkSequenceLength(sequence string) error {
	n := 9
	if len(sequence) != n*n+(n-1)*n+n-1 {
		return Error{m: "sequence is invalid, does not contains the required number of items"}
	}
	if !sequencePatternReg.MatchString(sequence) {
		return Error{m: "sequence is invalid, contains invalid characters"}
	}
	return nil
}

func (g *Grid) Notify() {
	for _, o := range g.observers {
		o.Update(g)
	}
}

func (g *Grid) Subscribe(observer observer.Observer) {
	g.observers = append(g.observers, observer)
}
