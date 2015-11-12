// sudoku solver
package main

import "log"
import "github.com/mandolyte/sudoku"


func main() {
	p := sudoku.NewPuzzle()
	solved, err := p.Solve()
	if err != nil {
		log.Fatalf("Error from NewPuzzle():\n%v" , err)
	}
	if solved {
		log.Println("Puzzle Solved!")
	} else {
		log.Println("Puzzle not solved :-(")
	}

}
