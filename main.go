package main

import (
	"fmt"
	"os"

	"sudoku-solver/model"
	"sudoku-solver/observer"
)

func main() {
	grid, err := model.NewGridFromSequence("9,0,0,1,0,0,0,0,5.0,0,5,0,9,0,2,0,1.8,0,0,0,4,0,0,0,0.0,0,0,0,8,0,0,0,0.0,0,0,7,0,0,0,0,0.0,0,0,0,2,6,0,0,9.2,0,0,3,0,0,0,0,6.0,0,0,2,0,0,9,0,0.0,0,1,9,0,4,5,7,0")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	obs1 := observer.NewObserver(func(i interface{}) {
		u := i.(*model.Grid)
		u.PrintInConsole()
	})
	grid.Subscribe(obs1)
}
