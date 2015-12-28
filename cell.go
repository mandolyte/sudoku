package sudoku

type cell struct {
  value uint8
  marks []uint8
}

func NewCell() *cell {
  c := new(cell)
  c.marks = make([]uint8,0,9)
  return c
}

