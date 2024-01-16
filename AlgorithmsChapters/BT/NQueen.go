package main

import (
	"fmt"
	"math"
)

func NQueens(Q []int, k int, n int) {
	if k == n {
		fmt.Println(Q)
		return
	}
	for i := 0; i < n; i++ {
		Q[k] = i
		if feasible(Q, k) {
			NQueens(Q, k+1, n)
		}
	}
}

func feasible(Q []int, k int) bool {
	for i := 0; i < k; i++ {
		if Q[k] == Q[i] || math.Abs(float64(Q[i]-Q[k])) == math.Abs(float64(i-k)) {
			return false
		}
	}
	return true
}

// Testing code.
func main() {
	Q := make([]int, 8)
	NQueens(Q, 0, 8)
}

/*
[0 4 7 5 2 6 1 3]
[0 5 7 2 6 3 1 4]
.......
[7 2 0 5 1 4 6 3]
[7 3 0 2 5 1 6 4]
*/