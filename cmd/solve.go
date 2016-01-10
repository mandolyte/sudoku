// sudoku solver
package main

import "log"
import "github.com/mandolyte/sudoku"
import "fmt"
import "flag"
import "os"
import "io/ioutil"

// setup command line arg requirements
var input *string = flag.String("i", "", 
	"Input file with puzzle; default is STDIN")
var otput *string = flag.String("o", "", 
	"Output file for solution; default is STDOUT")
var help  *bool   = flag.Bool("h", false, "Shows help/usage")

func main() {
	// parse command line
	flag.Parse() 

	// help?
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	p := sudoku.NewPuzzle()
	if flag.NFlag() == 0 { 
		// load from stdin
		if err := p.Load(); err != nil {
			log.Fatalf("Error processing puzzle:%v",err)
		}
	} else {
		// load from input file
		// must have an input argument
		if *input == "" {
			flag.PrintDefaults()
			os.Exit(0)
		}
		fpuz, err := ioutil.ReadFile(*input)
		if err != nil {
			log.Fatalf("Cannot read %v, err=%v",*input,err)
		}
		if err = p.Fload(fpuz); err != nil {
			log.Fatalf("Error processing puzzle:%v",err)
		}
	}

	// now that we have the puzzle constructed,
	// set the pencil marks and set any cells to those with
	// only one pencil mark possible
	p.SetPencilMarks()
	p.SetSingleMarks()


	sudoku.Solve(p)
	solutions := sudoku.GetSolutions()
	//number_solutions := len(solutions)
	//fmt.Printf("Number of solutions is:%v\n",number_solutions)
	for _,s := range solutions {
		fmt.Printf("%v",s)
	}
}
