package sudoku
import "fmt"
import "os"
import "sync"
// misc

var DEBUG bool = true
var mutex = &sync.Mutex{}
var solutions map[string]string


func dbg (msg string) {
  if DEBUG {
    fmt.Fprint(os.Stderr,msg)
  }
}

func init() {
    solutions = make(map[string]string)
}

func add_solution(sol, puz string) {
	mutex.Lock()
	solutions[sol] = puz
	mutex.Unlock()
}

func GetSolutions() map[string]string {
	return solutions
}

func Copy (p *puzzle) *puzzle {
	q := NewPuzzle()
	for n,_ := range p.cells {
		q.cells[n].value = p.cells[n].value
		q.cells[n].marks = append(q.cells[n].marks, 
			p.cells[n].marks...)
	}
	/* test only remove later
	f1 := p.fingerprint()
	f2 := q.fingerprint()
	if f1 == f2 {
		fmt.Printf("fingerprints do match:\n%v\n%v\n",f1,f2)
	} else {
		fmt.Printf("fingerprints do not match:\n%v\n%v\n",f1,f2)
	}
	*/
	return q
}

func Solve(p *puzzle) {
	//dbg(fmt.Sprintf("Solve() entered with:\n%v\n",p.String()))
	//dbg(fmt.Sprintf("Pencil marks are:\n%v\n",p.PencilMarks()))
	if err := p.Validate(); err == nil {
		add_solution(p.fingerprint(), p.String())
		//dbg("Solve() return via add_solution()\n")
		return
	}
	// find first blank cell (zero)
	x := -1 
	for n,_ := range p.cells {
		if p.cells[n].value == 0 {
			x = n
			//dbg(fmt.Sprintf("Blank found at:%v\n",x))
			break
		}
	}
	if x == -1 {
		// none found, just return
		//dbg("Solve() return via no blanks found\n")
		return
	}
	var wg sync.WaitGroup
	//dbg(fmt.Sprintf("Number of marks at %v is %v\n",
	//	x,len(p.cells[x].marks)))
	for _,v := range p.cells[x].marks {
		//dbg(fmt.Sprintf("Cell x=%v; trying mark %v\n",x,v))
		q := Copy(p)
		q.cells[x].value = v
		q.Remove_used_mark(v, x)
		wg.Add(1)
		go func (q *puzzle) {
			defer wg.Done()
			Solve(q)
		}(q)
	}
	wg.Wait()
	//dbg("Solve() return via child go routines done\n")
}