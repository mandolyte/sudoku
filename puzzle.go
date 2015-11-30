package sudoku

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "os"
)

type puzzle struct {
  cells []*cell
}

func NewPuzzle() *puzzle {
  p := new(puzzle)
  p.cells = make([]*cell,81,81)
  // initialize each cell
  for n,_ := range p.cells {
    p.cells[n] = NewCell()
  }
  return p
}

func (p *puzzle) String() string {
  var i,j,x int
  var sp string
  for i = 0; i < 9; i++ {
    for j = 0; j < 9; j++ {
      sp += fmt.Sprintf("%v",p.cells[x].value)
      x++
      if j < 8 {
        sp += (" ")
      }
    }
    sp += "\n"
  }
  return sp
}

func (p *puzzle) Load() error {
  /* read stdin for puzzle and load up */
  var i int
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    tokens := bufio.NewScanner(strings.NewReader(line))
    tokens.Split(bufio.ScanWords)
    for tokens.Scan() {
      c := tokens.Text()
      n,nerr := strconv.Atoi(c)

      if nerr == nil {
        p.cells[i].value = uint8(n)
      } else {
        p.cells[i].value = 0
      }
      i++
    }
  }
  // check if error broke the loop
  if err := scanner.Err(); err != nil {
    return err
  }

  return nil
}

func (p *puzzle) SetPencilMarks() {
  for n,_ := range p.cells {
    dbg(fmt.Sprintf("\n ==> Check Pencil Marks for %v\n",n))
    if p.cells[n].value != 0 {
      continue
    }
    for i:=1; i<10; i++ {
      if p.is_in_row(n,i) {
        continue
      }
      if p.is_in_col(n,i) {
        continue
      }
      if p.is_in_square(n,i) {
        continue
      }
      p.cells[n].marks = append(p.cells[n].marks,uint8(i))
    }
  }
}

func (p *puzzle) is_in_row(n,i int) bool {
  var r int = n/9
  dbg(fmt.Sprintf("is_in_row(%v,%v) is row %v\n",n,i,r))
  for c:=0; c<9; c++ {
    if p.cells[r*9+c].value == uint8(i) {
      dbg(fmt.Sprintf("True: %v is in row %v\n",i,r))
      return true
    }
  }
  return false
}

func (p *puzzle) is_in_col(n,i int) bool {
  var c int = n%9
  dbg(fmt.Sprintf("is_in_col(%v,%v) is col %v\n",n,i,c))
  for r:=0; r<9; r++ {
    if p.cells[r*9+c].value == uint8(i) {
      dbg(fmt.Sprintf("True: %v is in col %v\n",i,c))
      return true
    }
  }
  return false
}

func (p *puzzle) is_in_square(n,i int) bool {
  var r int = n/9
  var c int = n%9
  switch {
  case r < 3 && c < 3: r,c = 0,0
  case r < 3 && c < 6: r,c = 0,3
  case r < 3 && c < 9: r,c = 0,6

  case r < 6 && c < 3: r,c = 3,0
  case r < 6 && c < 6: r,c = 3,3
  case r < 6 && c < 9: r,c = 3,6

  case r < 9 && c < 3: r,c = 6,0
  case r < 9 && c < 6: r,c = 6,3
  case r < 9 && c < 9: r,c = 6,6

  default: dbg(fmt.Sprintf(
    "ERR: no square found for %v,%v",n,i))
    panic("see error")
  }
  dbg(fmt.Sprintf("is_in_square(%v,%v) is sq %v,%v\n",n,i,r,c))
  for x := r; x < r+3; x++ {
    for y := c; y < c+3; y++ {
      //dbg(fmt.Sprintf("X,Y=%v,%v\n",x,y))
      if p.cells[x*9+y].value == uint8(i) {
        dbg(fmt.Sprintf("True: %v is in sq %v,%v\n",i,r,c))
        return true
      }
    }
  }  

  return false
}

func (p *puzzle) PencilMarks() string {
  var i,j,x int
  var sp string
  for i = 0; i < 9; i++ {
    for j = 0; j < 9; j++ {
      if len(p.cells[x].marks) > 0 {
        sp += fmt.Sprintf("i,j,n=%v,%v,%v:%v\n",i,j,x,p.cells[x].marks)
      }
      x++
    }
  }
  return sp
}

func (p *puzzle) Validate() bool {
  // quick check any blanks/zeros
  for n,_ := range p.cells {
    if p.cells[n].value == 0 {
      dbg(fmt.Sprintf("Validate() found zero at n=%v",n))
      return false
    }
  }

  // check rows for dups, missing
  for r := 0; r < 9; r++ {
    rowdups := make(map[uint8]int)
    for c :=0; c < 9; c++ {
      rowdups[p.cells[r*9 + c].value]++
    }
    for i := 1; i < 10; i++ {
      if val,ok := rowdups[uint8(i)]; ok {
        if val > 1 {
          dbg(fmt.Sprintf("Validate() row %v, dup %v\n",r,i))
          return false
        }
      } else {
        dbg(fmt.Sprintf("Validate() row %v, missing val %v\n",r,i))
        return false
      }
    }
  }

  // check columns for dups, missing
  for c := 0; c < 9; c++ {
    coldups := make(map[uint8]int)
    for r :=0; r < 9; r++ {
      coldups[p.cells[r*9 + c].value]++
    }
    for i := 1; i < 10; i++ {
      if val,ok := coldups[uint8(i)]; ok {
        if val > 1 {
          dbg(fmt.Sprintf("Validate() col %v, dup %v\n",c,i))
          return false
        }
      } else {
        dbg(fmt.Sprintf("Validate() col %v, missing val %v\n",c,i))
        return false
      }
    }
  }

  // check squares for dups, missing
  // square 1 - case r < 3 && c < 3: r,c = 0,0
  if valid := p.validate_square(0,0); !valid {
    return false
  }
  // square 2 - case r < 3 && c < 6: r,c = 0,3
  if valid := p.validate_square(0,3); !valid {
    return false
  }
  // square 3 - case r < 3 && c < 9: r,c = 0,6
  if valid := p.validate_square(0,6); !valid {
    return false
  }
  // square 4 - case r < 6 && c < 3: r,c = 3,0
  if valid := p.validate_square(3,0); !valid {
    return false
  }
  // square 5 - case r < 6 && c < 6: r,c = 3,3
  if valid := p.validate_square(3,3); !valid {
    return false
  }
  // square 6 - case r < 6 && c < 9: r,c = 3,6
  if valid := p.validate_square(3,6); !valid {
    return false
  }
  // square 7 - case r < 9 && c < 3: r,c = 6,0
  if valid := p.validate_square(6,0); !valid {
    return false
  }
  // square 8 - case r < 9 && c < 6: r,c = 6,3
  if valid := p.validate_square(6,3); !valid {
    return false
  }
  // square 9 case r < 9 && c < 9: r,c = 6,6
  if valid := p.validate_square(6,6); !valid {
    return false
  }
  return true
}

func (p *puzzle) validate_square(x,y int) bool {
  //dbg(fmt.Sprintf("validate_square(%v,%v)\n",x,y))
  sqdups := make(map[uint8]int)
  for r := x; r < x+3; r++ {
    for c := y; c < y+3; c++ {
      sqdups[p.cells[r*9 + c].value]++
    }
  }
  for i := 1; i < 10; i++ {
    if val,ok := sqdups[uint8(i)]; ok {
      if val > 1 {
        dbg(fmt.Sprintf("Validate() square %v,%v, dup %v\n",x,y,i))
        return false
      }
    } else {
      dbg(fmt.Sprintf("Validate() square %v,%v, missing val %v\n",x,y,i))
      return false
    }
  }
  return true
}
