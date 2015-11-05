package sudoku
import "log"

var squares [9][3]int
func init() {
  // set the corners of each square in the puzzle
  squares[0] = [3]int{0,0,0}
  squares[1] = [3]int{0,3,0}
  squares[2] = [3]int{0,6,0}
  squares[3] = [3]int{3,0,0}
  squares[4] = [3]int{3,3,0}
  squares[5] = [3]int{3,6,0}
  squares[6] = [3]int{6,0,0}
  squares[7] = [3]int{6,3,0}
  squares[8] = [3]int{6,6,0}
}

func Process() {
  // load the puzzle
  load()
  debug()
  // create the squares
  var ninesqs [9]square
  for i:=0; i<9; i++ {
    ninesqs[i] = NewSquare(squares[i])
    //(nine[i]).Print()
    (ninesqs[i]).PencilMarks()
    //(nine[i]).PrintPencilMarks()
  }
  // create the rows
  var ninerows [9]row
  for i:=0; i<9; i++ {
    ninerows[i] = NewRow(i)
  }
  // create the columns
  var ninecols [9]col
  for j:=0; j<9; j++ {
    ninecols[j] = NewCol(j)
  }


  for {
    var change_count int
    for i:=0; i<9; i++ {
      c := (ninesqs[i]).ScanSetSinglePencilMarks()
      change_count += c
    }
    log.Printf("changed squares is %v\n", change_count)
    // keep going until nothing is left to do
    if change_count == 0 {
      break
    }
  }



}
