// sudoku solver
package main

import "fmt"
import "log"
import "github.com/mandolyte/sudoku"


func main() {
	p := sudoku.NewPuzzle()
	solved, err := p.Solve()
	if err != nil {
		log.Fatalf("Error from NewPuzzle():\n%v" , err)
	}
	if solved {
		fmt.Println("Puzzle Solved!")
	} else {
		fmt.Println("Puzzle not solved :-(")
	}

}
