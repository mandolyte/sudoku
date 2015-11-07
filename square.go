package sudoku
import "errors"
import "fmt"

type square struct {
  corner [3]int
  p *[9][9][10]int
}

func NewSquare(corner_coords [3]int, p *[9][9][10]int) square {
  var s square
  s.corner = corner_coords
  s.p = p
  return s
}

func (s square) Print() {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      fmt.Print(s.p[i][j][0])
    }
    fmt.Print("\n")
  }
}

func (s square) is_in_square(candidate int) bool {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      if s.p[i][j][0] == candidate {
        return true
      }
    }
  }
  return false
}

func (s square) is_in_row(row, candidate int) bool {
  var j int
  for j=0; j < 9; j++ {
    if s.p[row][j][0] == candidate {
      return true
    }
  }
  return false
}

func (s square) is_in_col(col, candidate int) bool {
  var i int
  for i=0; i < 9; i++ {
    if s.p[i][col][0] == candidate {
      return true
    }
  }
  return false
}

func (s square) PencilMarks() {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      // if coordinate in puzzle is not zeros continue
      // If zero then it needs pencil marks
      if s.p[i][j][0] == 0 {
        //fmt.Printf("puzzle[%v][%v] needs pencil marks\n",i,j)
        var x int
        for x=1; x < 10; x++ {
          // loop thru all 9 candidates (1..9)
          // if the candidate is not in the square already
          in_sq := s.is_in_square(x)
          if in_sq {
            continue
          }
          // and if not in the row
          in_row := s.is_in_row(i, x)
          if in_row {
            continue
          }
          // and if not in the column
          in_col := s.is_in_col(j, x)
          if in_col {
            continue
          }
          // then we have a pencil mark for this number
          // assign it to the nth position of the 3rd
          // dimension of this location in the square.
          s.p[i][j][x] = x
        }
      }
    }
  }
}

func (s square) PrintPencilMarks() {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      // if coordinate in puzzle is not zeros continue
      // If zero then it needs pencil marks
      if s.p[i][j][0] == 0 {
        fmt.Printf("Pencil marks for puzzle[%v][%v] are:",i,j)
        var x int
        for x=1; x < 10; x++ {
          if s.p[i][j][x] == 0 {
            continue
          }
          fmt.Printf("%v,",x)
        }
        fmt.Print("\n")
      }
    }
  }
}

func (s square) ScanSetSinglePencilMarks() int {
  var i,j int
  var total int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      // if coordinate in puzzle is not zeros continue
      // If zero then it needs pencil marks
      var y,count int = 0,0
      if s.p[i][j][0] == 0 {
        var x int
        for x=1; x < 10; x++ {
          if s.p[i][j][x] != 0 {
            count++
            y = x
          }
        }
        if count == 1 {
          s.p[i][j][0] = y
          total++
        }
      }
    }
  }
  return total
}

func (s square) validate() error {
  var tester [10]int
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      tester[s.p[i][j][0]] = s.p[i][j][0]
    }
  }
  for i=1; i<10; i++ {
    if tester[i] == 0 {
      return errors.New("Not solved")
    }
  }
  return nil
}