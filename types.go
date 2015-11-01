package sudoku

// The essential puzzle 9x9 puzzle
// with 9 more in 3rd dimension for pencil marks
var Puzzle [9][9][10]int
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
