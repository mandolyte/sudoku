package sudoku

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "os"
  "log"
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

}

func NewPuzzle() *puzzle {
  return new(puzzle)
}


/* output the puzzle */
func (p *puzzle) print() {
  var i,j int
  for i = 0; i < 9; i++ {
    for j = 0; j < 9; j++ {
      fmt.Print(p.val[i][j][0])
      fmt.Print(" ")
    }
    fmt.Print("\n")
  }
}


func (p *puzzle) load() error {
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
  return nil
}

func (p *puzzle) Solve() (bool, error) {
  // load the puzzle
  err := p.load()
  if err != nil {
    return false, err
  }


  // create the squares
  log.Println("DEBUG: load the squares")
  for i:=0; i<9; i++ {
    p.ninesqs[i] = NewSquare(squares[i], &p.val)
    //(nine[i]).Print()
    (p.ninesqs[i]).pencilMarks()
    //(nine[i]).PrintPencilMarks()
  }

  // create the rows
  log.Println("DEBUG: load the rows")
  for i:=0; i<9; i++ {
    p.ninerows[i] = NewRow(i,&p.val)
  }
  // create the columns
  log.Println("DEBUG: load the columns")
  for j:=0; j<9; j++ {
    p.ninecols[j] = NewCol(j,&p.val)
  }

  log.Println("DEBUG: set single pencil marks")

  for {
    var change_count int
    for i:=0; i<9; i++ {
      c := (p.ninesqs[i]).scanSetSinglePencilMarks()
      (p.ninesqs[i]).pencilMarks()
      change_count += c
    }
    log.Printf("changed squares is %v\n", change_count)
    // keep going until nothing is left to do
    if change_count == 0 {
      break
    }
  }

  p.print()
  // generate all possible squares
  for i:=0; i<9; i++ {
    log.Printf("Permuting square %v\n",i)
    (p.ninesqs[i]).printPencilMarks()
    (p.ninesqs[i]).permutations()
  }



  if err := p.validate(); err != nil {
    return false, nil
  }
  return true, nil
}


func (p *puzzle) validate() error {
  for i:=0; i<9; i++ {
    log.Printf("Validating square %v\n",i)
    err := (p.ninesqs[i]).validate()
    if err != nil {
      return err
    }
  }

  for i:=0; i<9; i++ {
    log.Printf("Validating row %v\n",i)
    err := (p.ninerows[i]).validate()
    if err != nil {
      return err
    }
  }

  for i:=0; i<9; i++ {
    log.Printf("Validating column %v\n",i)
    err := (p.ninecols[i]).validate()
    if err != nil {
      return err
    }
  }


  return nil
}

