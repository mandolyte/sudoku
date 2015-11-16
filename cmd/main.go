// sudoku solver
package main

import "log"
import "github.com/mandolyte/sudoku"


func main() {
	p := sudoku.NewPuzzle()
	solved, _ := p.Solve()
	if solved {
		log.Println("Puzzle Solved!")
	} else {
		log.Println("Puzzle not solved :-(")
	}

}
