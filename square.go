package sudoku
import "fmt"

type square struct {
  corner [3]int
}

func NewSquare(corner_coords [3]int) square {
  var s square
  s.corner = corner_coords
  return s
}

func (s square) Print() {
  var i,j int
  for i=s.corner[0]; i < s.corner[0]+3; i++ {
    for j=s.corner[1]; j < s.corner[1]+3; j++ {
      fmt.Print(Puzzle[i][j][0])
    }
    fmt.Print("\n")
  }
}
