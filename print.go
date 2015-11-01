package sudoku

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "os"
  "log"
)
/* output the puzzle */

func Print() {
  var i,j int
  for i = 0; i < 9; i++ {
    for j = 0; j < 9; j++ {
      fmt.Print(Puzzle[i][j][0])
      fmt.Print(" ")
    }
    fmt.Print("\n")
  }

}

func load() {
  /* read stdin for puzzle and load up */
  var i,j int
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    tokens := bufio.NewScanner(strings.NewReader(line))
    tokens.Split(bufio.ScanWords)
    j = 0
    for tokens.Scan() {
      cell := tokens.Text()
      if n,nerr := strconv.Atoi(cell); nerr == nil {
        Puzzle[i][j][0] = n
      } else {
        Puzzle[i][j][0] = 0
      }
      j++
    }
    i++
  }
  // check if error broke the loop
  if err := scanner.Err(); err != nil {
    log.Fatal("Error on stdin: %v\n", err)
  }
}
