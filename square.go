package sudoku
import "errors"
import "fmt"
import "log"

type square struct {
  corner [3]int
  p *[9][9][10]int
  candidates []([3][3]int)
  level int
}

func NewSquare(corner_coords [3]int, p *[9][9][10]int) *square {
  s := new(square)
  s.corner = corner_coords
  s.p = p
  s.candidates = make([]([3][3]int),0)
  return s
}

func (s *square) print() {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      fmt.Print(s.p[i][j][0])
    }
    fmt.Print("\n")
  }
}

func (s *square) is_in_square(candidate int) bool {
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

func (s *square) is_in_row(row, candidate int) bool {
  var j int
  for j=0; j < 9; j++ {
    if s.p[row][j][0] == candidate {
      return true
    }
  }
  return false
}

func (s *square) is_in_col(col, candidate int) bool {
  var i int
  for i=0; i < 9; i++ {
    if s.p[i][col][0] == candidate {
      return true
    }
  }
  return false
}

func (s *square) pencilMarks() {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      // clear out any old pencil marks
      for k:=1; k < 10; k++ {
        s.p[i][j][k] = 0
      }
      // if coordinate in puzzle is not zeros continue
      // If zero then it needs pencil marks
      if s.p[i][j][0] == 0 {
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

func (s *square) printPencilMarks() {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      // if coordinate in puzzle is not zeros continue
      // If zero then it needs pencil marks
      if s.p[i][j][0] == 0 {
        log.Printf("Pencil marks for puzzle[%v][%v] are:",i,j)
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

func (s *square) scanSetSinglePencilMarks() int {
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

func (s *square) validate() error {
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

func testValidate(psquare [3][3]int) error {
  var tester [10]int
  var i,j int
  for i=0; i < 3; i++ {
    for j=0; j < 3; j++ {
      tester[psquare[i][j]] = psquare[i][j]
    }
  }
  for i=1; i<10; i++ {
    if tester[i] == 0 {
      return errors.New("Not solved")
    }
  }
  return nil
}

func (s *square) permutations() {
  var psquare [3][3]int
  // make a copy of our 3x3; counting the blanks
  var count, i2, j2 int
  for i:=s.corner[0]; i < s.corner[0]+3; i++ {
    j2 = 0
    for j:=s.corner[1]; j < s.corner[1]+3; j++ {
      psquare[i2][j2] = s.p[i][j][0]
      if psquare[i2][j2] == 0 {
        count++
      }
      j2++
    }
    i2++
  }
  log.Printf("psquare is:\n%v\n",psquare)
  if count == 0 {
    // nothing to do... square is solved
    s.candidates = append(s.candidates, psquare)
  } else {
    // now use this copy and generate all permutations 
    // this function is recursive
    s.permutate(s.corner[0],s.corner[1],1,psquare)
  }
  log.Printf("Square has %v possibilities\n",len(s.candidates))
  for _,ps := range s.candidates {
    for i:=0; i < 3; i++ {
      for j:=0; j < 3; j++ {
        fmt.Print(ps[i][j])
      }
      fmt.Print("\n")
    }
    fmt.Print("---\n")
  }
}

func (s *square) permutate(x,y,z int,psquare [3][3]int) {   
  s.level++
  if tverr := testValidate(psquare); tverr == nil {
    s.candidates = append(s.candidates,psquare)
  }
  var xwindup bool = true
  var ywindup bool = true
  for i:=s.corner[0]; i < s.corner[0]+3; i++ {
    if i < x && xwindup {
      continue
    } else {
      xwindup = false
    }
    for j:=s.corner[1]; j < s.corner[1]+3; j++ { 
      if j < y && ywindup {
        continue
      } else {
        ywindup = false
      }
      if s.p[i][j][0] == 0 { 
        for k:=z; k < 10; k++ { 
          if s.p[i][j][k] != 0 { 
            psquare[i-s.corner[0]][j-s.corner[1]] = s.p[i][j][k] 
            // start next level at next cell
            // k is easy, just set to 1
            // i and j are inter-related 
            if j < s.corner[1]+3 {
              // ok, just increment j
              s.permutate(i,j+1,1,psquare)
            } else if i < s.corner[0]+3 {
              // tricky bit
              // set j to back to corner column number
              // and increment i
              s.permutate(i+1,s.corner[1],1,psquare)
            } else {
              // this means we are at the end of the square
              // just wrap and go home (do nothing)
            }
          }
        }
      }    
    }   
  } 
  s.level-- 
}
