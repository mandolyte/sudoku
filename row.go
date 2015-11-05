package sudoku
import "fmt"

type row struct {
  rownum int
}

func NewRow(rownum int) row {
  var r row
  r.rownum = rownum
  return r
}

func (r row) Print() {
  var j int
  for j=0; j < 9; j++ {
    fmt.Print(puzzle[r.rownum][j][0])
  }
  fmt.Print("\n")
}

func (r row) is_in_row(candidate int) bool {
  var j int
  for j=0; j < 9; j++ {
    if puzzle[r.rownum][j][0] == candidate {
      return true
    }
  }
  return false
}
