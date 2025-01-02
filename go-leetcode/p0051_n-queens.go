package leetcode

// https://leetcode.com/problems/n-queens/description/

func solveNQueens(n int) [][]string {
	var result [][]string

	curBoard := make([][]string, n)
	for i := range curBoard {
		curBoard[i] = make([]string, n)
		for j := range curBoard[i] {
			curBoard[i][j] = "."
		}
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			newBoard := make([]string, n)
			for i := range curBoard {
				newBoard[i] = ""
				for j := range curBoard[i] {
					newBoard[i] += curBoard[i][j]
				}
			}

			result = append(result, newBoard)

			return
		}

		for col := 0; col < n; col++ {
			if queenCanBePlaced(row, col, curBoard) {
				curBoard[row][col] = "Q"
				backtrack(row + 1)
				curBoard[row][col] = "."
			}
		}
	}

	backtrack(0)

	return result
}

func queenCanBePlaced(row int, col int, board [][]string) bool {
	for row := row - 1; row >= 0; row-- {
		if board[row][col] == "Q" {
			return false
		}
	}

	for row, col := row-1, col-1; row >= 0 && col >= 0; row, col = row-1, col-1 {
		if board[row][col] == "Q" {
			return false
		}
	}

	for row, col := row-1, col+1; row >= 0 && col < len(board); row, col = row-1, col+1 {
		if board[row][col] == "Q" {
			return false
		}
	}

	return true
}
