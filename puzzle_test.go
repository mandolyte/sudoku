package sudoku

import (
	"testing"
)

var case1 = `
_ _ _ 5 _ _ _ _ _
_ 7 _ _ _ 2 8 6 _
4 _ 2 _ 7 1 3 _ _
_ 9 _ _ _ _ _ _ 8
_ _ 7 _ 8 _ 1 _ _
8 _ _ _ _ _ _ 3 _
_ _ 3 9 4 _ 6 _ 1
_ 4 8 2 _ _ _ 9 _
_ _ _ _ _ 7 _ _ _
`



func Test_inital_processing(t *testing.T) {
	p := NewPuzzle()

	if err := p.Fload([]byte(case1)); err != nil {
		t.Fatalf("[fail]p.Fload() error:%v",err)
	}

	expected := "000500000070002860402071300090000008007080100800000030003940601048200090000007000"
	result := p.fingerprint()
	if result != expected {
		t.Fatalf("[fail] fingerprint error! Expected:\n%v\nReceived:\n%v\n",expected,result)

	}

	p.SetPencilMarks()
	p.SetSingleMarks()

	expected = "000564000070392864462871359090000008007080100800000030003940601048200090000007000"
	result = p.fingerprint()
	if result != expected {
		t.Fatalf("[fail] Pencil marks error! Expected:\n%v\nReceived:\n%v\n",expected,result)

	}

	Solve(p)
	expected = "389564712571392864462871359695413278237689145814725936753948621148236597926157483"
	solutions := GetSolutions()
	//number_solutions := len(solutions)
	//fmt.Printf("Number of solutions is:%v\n",number_solutions)
	if len(solutions) != 1 {
		t.Fatalf("[fail] Solver error! Expected:1, Received:%v\n",len(solutions) )

	}
	// get the result
	for result,_ = range solutions {
	}

	if result != expected {
		t.Fatalf("[fail] Pencil marks error! Expected:\n%v\nReceived:\n%v\n",expected,result)
	}
}
