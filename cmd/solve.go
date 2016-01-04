// sudoku solver
package main

import "log"
import "github.com/mandolyte/sudoku"
import "fmt"
//import "time"


func main() {
	p := sudoku.NewPuzzle()
	if err := p.Load(); err != nil {
		log.Fatal("Error on Load()")
	}
	p.SetPencilMarks()
	p.SetSingleMarks()
	//s := p.String()
	//fmt.Println(s)

	sudoku.Solve(p)
	//time.Sleep(30 * time.Second)
	solutions := sudoku.GetSolutions()
	//number_solutions := len(solutions)
	//fmt.Printf("Number of solutions is:%v\n",number_solutions)
	for _,s := range solutions {
		fmt.Printf("%v",s)
	}
}
