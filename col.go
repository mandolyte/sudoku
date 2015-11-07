package sudoku
import "errors"
import "fmt"

type col struct {
  colnum int
  p *[9][9][10]int
}

func NewCol(colnum int, p *[9][9][10]int) col {
  var c col
  c.colnum = colnum
  c.p = p
  return c
}

func (c col) Print() {
  var i int
  for i=0; i < 9; i++ {
    fmt.Print(c.p[i][c.colnum][0])
  }
  fmt.Print("\n")
}

func (c col) is_in_col(candidate int) bool {
  var i int
  for i=0; i < 9; i++ {
    if c.p[i][c.colnum][0] == candidate {
      return true
    }
  }
  return false
}

func (c col) validate() error {
  var tester [10]int
  for i:=0; i < 9; i++ {
    tester[c.p[i][c.colnum][0]] = c.p[i][c.colnum][0]
  }
  for i:=1; i<10; i++ {
    if tester[i] == 0 {
      return errors.New("Not solved")
    }
  }
  return nil
}