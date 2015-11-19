package sudoku
import "errors"
import "fmt"

type square struct {
  corner [3]int
  p *[9][9][10]int
  candidates map[string][3][3]int
  linearsq [9]int
  linearpm [][]int
  level int
}

func NewSquare(corner_coords [3]int, p *[9][9][10]int) *square {
  s := new(square)
  s.corner = corner_coords
  s.p = p
  s.candidates = make(map[string][3][3]int)
  return s
}

/* this method is only for testing purposes */
func SetSquare(a,b,c,d,e,f,g,h,i int) *square {
  var val [9][9][10]int

  val[0][0][0] = a
  val[0][1][0] = b
  val[0][2][0] = c

  val[1][0][0] = d
  val[1][1][0] = e
  val[1][2][0] = f

  val[2][0][0] = g
  val[2][1][0] = h
  val[2][2][0] = i

  s := new(square)
  s.corner = [3]int{0,0,0}
  s.p = &val
  s.candidates = make(map[string][3][3]int)
  return s
}

func (s *square) String() string {
  var i,j int
  var sp string
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      sp += fmt.Sprintf("%v",s.p[i][j][0])
    }
    sp += "\n"
  }
  return sp
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

func (s *square) PencilMarks() {
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

func (s *square) PrintPencilMarks() {
  dbg(fmt.Sprintf("Pencil Marks for square:\n%v",s.String()))
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      // if coordinate in puzzle is not zeros continue
      // If zero then it needs pencil marks
      dbg(fmt.Sprintf("Pencil marks for cell [%v %v]:",i,j))
      var x int
      for x=1; x < 10; x++ {
        if s.p[i][j][x] == 0 {
          continue
        }
        dbg(fmt.Sprintf("%v,",x))
      }
      dbg(fmt.Sprint("\n"))
    }
  }
}

func (s *square) ScanSetSinglePencilMarks() int {
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
          if s.p[i][j][x]  != 0 {
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

func (s *square) Validate() error {
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

func (s *square) linearValidate() error {
  //dbg(fmt.Sprintf("linearValidate() testing %v",s.linearsq))
  var tester [10]int
  for i:=0; i<9; i++ {
    tester[s.linearsq[i]] = s.linearsq[i]
  }
  for i:=1; i<10; i++ {
    if tester[i] == 0 {
      return errors.New("Not solved")
    }
  }
  return nil
}

func (s *square) Permutations() {
  s.Convert2DTo1D()
  s.BruteForcePermute()
}

func (s *square) Convert2DTo1D() {
  s.linearsq[0] = s.p[s.corner[0]][s.corner[1]][0]
  s.linearsq[1] = s.p[s.corner[0]][s.corner[1]+1][0]
  s.linearsq[2] = s.p[s.corner[0]][s.corner[1]+2][0]

  s.linearsq[3] = s.p[s.corner[0]+1][s.corner[1]][0]
  s.linearsq[4] = s.p[s.corner[0]+1][s.corner[1]+1][0]
  s.linearsq[5] = s.p[s.corner[0]+1][s.corner[1]+2][0]

  s.linearsq[6] = s.p[s.corner[0]+2][s.corner[1]][0]
  s.linearsq[7] = s.p[s.corner[0]+2][s.corner[1]+1][0]
  s.linearsq[8] = s.p[s.corner[0]+2][s.corner[1]+2][0]

  s.linearpm = make([][]int,0)
  // fill up the slices with pencil marks

  for i:=s.corner[0]; i<s.corner[0]+3; i++ {
    for j:=s.corner[1]; j<s.corner[1]+3; j++ {
      tmp := make([]int,0)
      for k:=1; k<10; k++ {
        if s.p[i][j][k] == 0 {
          continue
        }
        tmp = append(tmp,k) // k is 1 to 9, the P/M itself
      }
      s.linearpm = append(s.linearpm,tmp)
    }
  }

  /*
  // debug
  for n,sq := range s.linearpm {
    dbg(fmt.Sprintf("Convert2DTo1D() n=%v, pencilmarks=%v, len=%v\n",
      n,sq,len(sq)))
  }
  */

}

func (s *square) BruteForcePermute() {
  s.bruteForceCandidates(0)
  dbg(fmt.Sprintf("Square has %v possibilities\n",len(s.candidates)))
  /*
  if debug {
    for _,ps := range s.candidates {
      dbg(fmt.Sprintf("%v\n",ps))
    }
  }
  */
}

func (s *square) bruteForceCandidates(loc int) {
  //dbg(fmt.Sprintf("bruteForceCandidates() with loc=%v\n",loc))
  if loc > 8 {
    return
  }
  if len(s.linearpm[loc]) > 0 {
    
    //dbg(fmt.Sprintf("loc=%v, value=%v, pm_len=%v\n",
      //loc,s.linearsq[loc],len(s.linearpm[loc])))
    
    for _, cell := range s.linearpm[loc] {
      //dbg(fmt.Sprintf("Working on loc=%v,pm=%v\n",loc,cell))
      // copy pencil marks into array/slice
      s.linearsq[loc] = cell
      if err := s.linearValidate(); err == nil {
        // good to go; this is a candidate
        // copy to map of candidates
        // first, convert back to a square
        var psquare[3][3]int
        psquare[0][0] = s.linearsq[0]
        psquare[0][1] = s.linearsq[1]
        psquare[0][2] = s.linearsq[2]
        psquare[1][0] = s.linearsq[3]
        psquare[1][1] = s.linearsq[4]
        psquare[1][2] = s.linearsq[5]
        psquare[2][0] = s.linearsq[6]
        psquare[2][1] = s.linearsq[7]
        psquare[2][2] = s.linearsq[8]
        //dbg(fmt.Sprintf("linearValidate() valid:%v\n",psquare))
        s.candidates[squareFingerprint(psquare)] = psquare
      } else { 
        // ok, not on last cell in Square, keep going
        //dbg(fmt.Sprintf("linearValidate() not valid:%v\n",s.linearsq))
        s.bruteForceCandidates(loc + 1)
      }
    }
    //dbg(fmt.Sprintf("... done with all pencilmarks at %v\n",loc))
    s.bruteForceCandidates(loc+1)
  }
  //dbg(fmt.Sprintf("... no pencil marks at %v, keep going\n",loc))
  s.bruteForceCandidates(loc+1)
}


