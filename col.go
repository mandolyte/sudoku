package sudoku
import "fmt"

type col struct {
  colnum int
}

func NewCol(colnum int) col {
  var c col
  c.colnum = colnum
  return c
}

func (c col) Print() {
  var i int
  for i=0; i < 9; i++ {
    fmt.Print(puzzle[i][c.colnum][0])
  }
  fmt.Print("\n")
}

func (c col) is_in_col(candidate int) bool {
  var i int
  for i=0; i < 9; i++ {
    if puzzle[i][c.colnum][0] == candidate {
      return true
    }
  }
  return false
}
