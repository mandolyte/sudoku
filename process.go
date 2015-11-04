package sudoku
import "fmt"

func Process() {
  // load the puzzle
  load()
  print()
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
      c := (nine[i]).ScanSetSinglePencilMarks()
      change_count += c
    }

    // notes: create a row and col equiv to square scan
    // thus:
    // c := (row[i]).ScanForSingleMissing()
    // and
    // c := (col[j]).ScanForSingleMissing()
    for i:=0; i<9; i++ {
      var found [9]int = {0,0,0,0,0,0,0,0,0}
      for j:=0; j<9; j++ {
        if puzzle[i][j][0] != 0 {
          found[j] = j
        }
      }
      var count,single_col int = 0,0
      for j:=0; j<9; j++ {
        if found[j] == 0 {
          count++
          single_col = j
        }
      }
      if count == 1 {
        puzzle[i][j][single_col] =
      }
    }
    if change_count == 0 {
      break
    }
    fmt.Printf("changed squares is %v\n", change_count)
  }



}
