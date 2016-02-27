package sudoku
import "os"
import "sync"
import "github.com/mandolyte/simplelogger"
import "sync/atomic"
// misc

var mutex = &sync.Mutex{}
var solutions map[string]string
var sl *simplelogger.SimpleLogger
var Counter uint64

func init() {
    solutions = make(map[string]string)
    sl = &simplelogger.SimpleLogger {
        INFO:   true,
        DEBUG:  true,
        Writer: os.Stderr,
    }
    //sl.Info("init() @ time.now")
}

func inc_counter() {
	atomic.AddUint64(&Counter, 1)
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
	inc_counter()
	if err := p.Validate(); err == nil {
		add_solution(p.fingerprint(), p.String())
		return
	}
	// find first blank cell (zero)
	x := -1 
	for n,_ := range p.cells {
		if p.cells[n].value == 0 {
			x = n
			break
		}
	}
	if x == -1 {
		// none found, just return
		return
	}
	var wg sync.WaitGroup
	//	x,len(p.cells[x].marks)))
	for _,v := range p.cells[x].marks {
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
}