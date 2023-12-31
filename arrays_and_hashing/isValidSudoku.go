package arraysandhashing

// 36. Valid Sudoku
// Medium
// 9.7K
// 1K
// Companies
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Example 1:

// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true
// Example 2:

// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

// Constraints:

// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.

// this solution doesn't account for the two other conditions
func (ArrayAlg) isValidSudoku(board [][]byte) bool {

	boardExpectedLength := 9
	if len(board) != boardExpectedLength {
		return false
	}
	for _, value := range board {
		if len(value) != boardExpectedLength {
			return false
		}
	}
	var cStepper int
	var rStepper int
	// var counterDivider int
	var counterDividerTracker int
	checker := make(map[byte]bool)
	cStepper = 0
	rStepper = 0
	for cStepper != boardExpectedLength && rStepper != boardExpectedLength {

		for i := rStepper; i < rStepper+3; i++ {
			data := board[cStepper][i]
			if value := checker[data]; value {
				// fmt.Println(counterDivider)
				return false
			}
			if data != '.' {
				checker[data] = true
			}
		}
		cStepper++
		counterDividerTracker++
		if counterDividerTracker%3 == 0 {
			// counterDivider++

			checker = make(map[byte]bool)
			// cStepper = counterDivider

		}

		if cStepper == boardExpectedLength-1 && rStepper != boardExpectedLength-1 {
			rStepper = rStepper + 3
			cStepper = 0
		}

	}

	return true
}

// this solution solves the problem with nxn = n^2
func (ArrayAlg) isValidSudokuS(board [][]byte) bool {
	boardExpectedLength := 9
	if len(board) != boardExpectedLength {
		return false
	}
	for _, val := range board {
		if len(val) != boardExpectedLength {
			return false
		}
	}
	cStepper := 0
	rStepper := 0
	divider := 0
	checker := make(map[byte]bool)

	for cStepper != boardExpectedLength && rStepper != boardExpectedLength {
		for i := rStepper; i < rStepper+3; i++ {
			val := board[cStepper][i]
			if exist := checker[val]; exist {
				return false
			}
			if val != '.' {
				checker[val] = true
			}
		}
		cStepper++
		divider++
		if divider%3 == 0 {
			checker = make(map[byte]bool)
		}

		if cStepper == boardExpectedLength && rStepper != boardExpectedLength-1 {
			cStepper = 0
			rStepper += 3
		}

	}

	for _, value := range board {
		checker := make(map[byte]bool)
		for _, val := range value {
			if exist := checker[val]; exist {
				return false
			}
			if val != '.' {
				checker[val] = true
			}
		}
	}
	for i := 0; i < boardExpectedLength; i++ {
		checker := make(map[byte]bool)
		for j := 0; j < boardExpectedLength; j++ {
			if exist := checker[board[j][i]]; exist {
				return false
			}
			if board[j][i] != '.' {
				checker[board[j][i]] = true
			}

		}
	}
	return true
}

func (ArrayAlg) isValidSudokuL(board [][]byte) bool {
	var row, column, mat [9][9]bool
	boardlenth := len(board)
	for i := 0; i < boardlenth; i++ {
		for j := 0; j < boardlenth; j++ {
			if board[i][j] != '.' {
				k := int(board[i][j]) - 49
				if row[i][k] || column[j][k] || mat[i/3*3+j/3][k] {
					return false
				}
				row[i][k], column[j][k], mat[i/3*3+j/3][k] = true, true, true
			}
		}
	}
	return true
}
