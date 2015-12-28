// sudoku solver
package main

import "log"
import "github.com/mandolyte/sudoku"
import "fmt"


func main() {
	p := sudoku.NewPuzzle()
	if err := p.Load(); err != nil {
		log.Fatal("Error on Load()")
	}
	s := p.String()
	fmt.Println(s)
	p.SetPencilMarks()
	p.SetSingleMarks()
	if err := p.Validate(); err != nil {
		fmt.Printf("%v\n",err)
	} else {
		fmt.Printf("Puzzle is solved.\n")
	}

}
