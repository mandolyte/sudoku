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
	fmt.Printf("Validate() returns:%v\n",p.Validate())
}
