package sudoku


func Process() {
  // load the puzzle
  load()
  // create the squares
  s1 := NewSquare(squares[0])
  s1.Print()
}
