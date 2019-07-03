package main

import "fmt"

func NQueens(n int) {
	Q := make([]int, n)
	NQueensUtil(Q, 0, 8)
}

func NQueensUtil(Q []int, k int, n int) {
	if k == n {
		Print(Q, n)
		return
	}
	for i := 0; i < n; i++ {
		Q[k] = i
		if Feasible(Q, k) {
			NQueensUtil(Q, k+1, n)
		}
	}
}

func Absolute(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Feasible(Q []int, k int) bool {
	for i := 0; i < k; i++ {
		if Q[k] == Q[i] || Absolute(Q[i]-Q[k]) == Absolute(i-k) {
			return false
		}
	}
	return true
}

func Print(Q []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ", Q[i])
	}
	fmt.Println()
}

func main2() {
	NQueens(8)
}