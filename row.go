package sudoku
import "errors"
import "fmt"

type row struct {
  rownum int
  p *[9][9][10]int
}

func NewRow(rownum int, p *[9][9][10]int) row {
  var r row
  r.rownum = rownum
  r.p = p
  return r
}

func (r row) Print() {
  var j int
  for j=0; j < 9; j++ {
    fmt.Print(r.p[r.rownum][j][0])
  }
  fmt.Print("\n")
}

func (r row) is_in_row(candidate int) bool {
  var j int
  for j=0; j < 9; j++ {
    if r.p[r.rownum][j][0] == candidate {
      return true
    }
  }
  return false
}

func (r row) validate() error {
  var tester [10]int
  for j:=0; j < 9; j++ {
    tester[r.p[r.rownum][j][0]] = r.p[r.rownum][j][0]
  }
  for i:=1; i<10; i++ {
    if tester[i] == 0 {
      return errors.New("Not solved")
    }
  }
  return nil
}