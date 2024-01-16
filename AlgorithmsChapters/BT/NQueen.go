package main

import (
	"fmt"
	"math"
)

func NQueens(n int) {
	board := make([]int, n)
	nQueensUtil(board, 0, n)
}

func nQueensUtil(board []int, col int, n int) {
	if col == n {
		fmt.Println(board)
		return
	}

	for row := 0; row < n; row++ {
		board[col] = row
		if isSafe(board, col) {
			nQueensUtil(board, col+1, n)
		}
	}
}

func isSafe(board []int, col int) bool {
	for i := 0; i < col; i++ {
		if board[col] == board[i] || math.Abs(float64(board[col]-board[i])) == float64(col-i) {
			return false
		}
	}
	return true
}

// Testing code.
func main() {
	NQueens(8)
}

/*
[0 4 7 5 2 6 1 3]
[0 5 7 2 6 3 1 4]
.......
[7 2 0 5 1 4 6 3]
[7 3 0 2 5 1 6 4]
*/
