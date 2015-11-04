# Go Challenge #8
... is to make a Sudoku solver. I'll probably never finish this being full time employed and active in many outside interests. But I love playing the game and I love using Go. Hard to resist at least thinking about the problem!

The challenge page is here:
http://golang-challenge.com/go-challenge8/

The page lists the rules and provides a sample input and output. Main constraints are:
- Read a puzzle on stdin using format in test1.txt
- Output solution on stdout using format in test1_expected.txt

# Design musings
## Essential data structure
*the puzzle content* Thinking of a 3D array of integers. The 2D slice would represent the input puzzle; with zeros for the blanks. The 3D slice would represent the pencil marks. Assuming a max number of pencil marks of 8, then the array would be 9 rows, by 9 columns, by 9 high: 9x9x9 = 729 slots. If the max number of pencil marks is 9, then 810 (9x9x10).

*the 3x3 square* Each square could be represented with a struct containing its data from the puzzle, the methods to interrogate the intersecting rows and columns, etc.

## Generating the pencil marks
If each 3x3 square is handled separately, perhaps by its own go routine, then each square can figure out its own pencil marks without interfering with the others.

Each 3x3 square:
- construct an array 0 to 8 (9 integers) that represent its non-zero content; for example, a 5 would go in the 4th slot
- For each blank, for each missing number (not in the array above), test the number for presence in the row and column of the blank. If not there insert into the slot in the pencil marks.

For example, based on the test puzzle below, in the first square (upper left), the pencil marks should be:
(0,1) --> 2
(1,0) --> 4
(1,2) --> 4,6
(2,1) --> 2,8

```
1 _ 3 _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8

```
