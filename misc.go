package sudoku
import "errors"
import "fmt"
import "os"
// misc

var debug bool = true
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


func squareValidate(psquare [3][3]int) error {
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

func squareFingerprint(s [3][3]int) string {
  return fmt.Sprintf("%v%v%v%v%v%v%v%v%v",
    s[0][0],s[0][1],s[0][2],
    s[1][0],s[1][1],s[1][2],
    s[2][0],s[2][1],s[2][2])
}

func dbg (msg string) {
  if debug {
    fmt.Fprint(os.Stderr,msg)
  }
}
