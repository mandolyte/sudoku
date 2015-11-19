package sudoku

import (
  "fmt"
  "bufio"
  "errors"
  "strconv"
  "strings"
  "os"
)


type puzzle struct {
  // The essential puzzle 9x9 puzzle
  // with 9 more in 3rd dimension for pencil marks
  val [9][9][10]int

  // representation of each square in the puzzle
  ninesqs [9]*square

  // representation of the each row
  ninerows [9]row

  // representation of each column
  ninecols [9]col

  solutions map[string][9][9][10]int

}

func NewPuzzle() *puzzle {
  p := new(puzzle)
  p.solutions = make(map[string][9][9][10]int)
  return p
}


/* output the puzzle */
func (p *puzzle) print() {
  var i,j int
  for i = 0; i < 9; i++ {
    for j = 0; j < 9; j++ {
      fmt.Print(p.val[i][j][0])
      if j < 8 {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n")
  }
}

func (p *puzzle) String() string {
  var i,j int
  var sp string
  for i = 0; i < 9; i++ {
    for j = 0; j < 9; j++ {
      sp += fmt.Sprintf("%v",p.val[i][j][0])
      if j < 8 {
        sp += (" ")
      }
    }
    sp += "\n"
  }
  return sp
}

func (p *puzzle) GetSquare(i int) *square {
  return p.ninesqs[i]
}


func (p *puzzle) Load() error {
  /* read stdin for puzzle and load up */
  var i,j int
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    tokens := bufio.NewScanner(strings.NewReader(line))
    tokens.Split(bufio.ScanWords)
    j = 0
    for tokens.Scan() {
      cell := tokens.Text()
      n,nerr := strconv.Atoi(cell)

      if nerr == nil {
        p.val[i][j][0] = n
      } else {
        p.val[i][j][0] = 0
      }
      j++
    }
    i++
  }
  // check if error broke the loop
  if err := scanner.Err(); err != nil {
    return err
  }
  // create the squares
  dbg("DEBUG: load the squares\n")
  for i:=0; i<9; i++ {
    p.ninesqs[i] = NewSquare(squares[i], &p.val)
    //(p.ninesqs[i]).PencilMarks()
  }

  // create the rows
  dbg("DEBUG: load the rows\n")
  for i:=0; i<9; i++ {
    p.ninerows[i] = NewRow(i,&p.val)
  }
  // create the columns
  dbg("DEBUG: load the columns\n")
  for j:=0; j<9; j++ {
    p.ninecols[j] = NewCol(j,&p.val)
  }

  dbg("DEBUG: set single pencil marks\n")

  for {
    var change_count int
    for i:=0; i<9; i++ {
      c := (p.ninesqs[i]).ScanSetSinglePencilMarks()
      (p.ninesqs[i]).PencilMarks()
      change_count += c
    }
    dbg(fmt.Sprintf("changed squares is %v\n", change_count))
    // keep going until nothing is left to do
    if change_count == 0 {
      break
    }
  }
  // this will print puzzle after initial work
  //p.print()
  return nil
}

func (p *puzzle) Solve() (bool, error) {
  // generate all possible squares
  for i:=0; i<9; i++ {
    dbg(fmt.Sprintf("Permuting square %v\n",i))
    if debug {
      (p.ninesqs[i]).PrintPencilMarks()
    }
    (p.ninesqs[i]).Permutations()
  }
  // is it solved (must have easy puzzle!)
  if err := p.validate(); err == nil {
    p.print()
    return true, nil
  }

  // try to solve
  p.bruteForceSolve(0)

  num_sols := len(p.solutions)
  dbg(fmt.Sprintf("Number of solutions is %v\n",num_sols))
  if num_sols == 1 {
    p.print()
    return true, nil
  } else if num_sols > 1 {
    return false, errors.New("Error! more than one solution")
  }
  return false, errors.New("Not Solved :-(")
}


func (p *puzzle) validate() error {
  for i:=0; i<9; i++ {
    //log.Printf("Validating square %v\n",i)
    err := (p.ninesqs[i]).Validate()
    if err != nil {
      return err
    }
  }

  for i:=0; i<9; i++ {
    //log.Printf("Validating row %v\n",i)
    err := (p.ninerows[i]).validate()
    if err != nil {
      return err
    }
  }

  for i:=0; i<9; i++ {
    //log.Printf("Validating column %v\n",i)
    err := (p.ninecols[i]).validate()
    if err != nil {
      return err
    }
  }


  return nil
}

func (p *puzzle) bruteForceSolve(sq int) {
  dbg(fmt.Sprintf("bruteForceSolve() on %v\n",sq))
  if sq > 8 {
    return 
  }

  s := p.ninesqs[sq]
  num := len(s.candidates)
  if num > 0 {
    var count int = -1
    for _, psquare := range s.candidates {
      // copy psquare into puzzle
      count++
      dbg(fmt.Sprintf("Testing square %v: %v of %v\n",sq,count,num))
      for i:=0; i < 3; i++ {
        for j:=0; j < 3; j++ {
          p.val[s.corner[0]+i][s.corner[1]+j][0] = psquare[i][j]
        }
      }
      // ok, now it is copied... test to see if puzzle is solved
      if err := p.validate(); err == nil {
        // Yeah! puzzle is solved
        //dbg("Puzzle solved (a)!\n")
        p.solutions[p.String()] = p.val
      } 
      // with this candidate copied, let's proceed to the
      // other squares in turn to see they have the missing
      // combination
      // first - are we on last square?
      if sq == 8 {
        return  // bummer
      }
      // ok, not on last square, keep going
      p.bruteForceSolve(sq+1)
    }
  } else {
    p.bruteForceSolve(sq+1)
  }
  if err := p.validate(); err == nil {
    // Yeah! puzzle is solved
    //dbg("Puzzle solved (b)!\n")
    p.solutions[p.String()] = p.val
  }
  return 
}
