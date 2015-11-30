package sudoku
import "fmt"
import "os"
// misc

var DEBUG bool = true

func dbg (msg string) {
  if DEBUG {
    fmt.Fprint(os.Stderr,msg)
  }
}
