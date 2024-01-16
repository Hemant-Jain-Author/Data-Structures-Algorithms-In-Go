package main

import "fmt"

func MinCost(cost [][]int, m int, n int) int {
	if m == 0 || n == 0 {
		return 99999
	}
	if m == 1 && n == 1 {
		return cost[0][0]
	}
	return cost[m-1][n-1] + min(MinCost(cost, m-1, n-1), MinCost(cost, m-1, n), MinCost(cost, m, n-1))
}

func MinCostBU(cost [][]int, m int, n int) int {
	tc := make([][]int, m)
	for i := range tc {
		tc[i] = make([]int, n)
	}
	tc[0][0] = cost[0][0]

	// Initialize first column.
	for i := 1; i < m; i++ {
		tc[i][0] = tc[i-1][0] + cost[i][0]
	}

	// Initialize first row.
	for j := 1; j < n; j++ {
		tc[0][j] = tc[0][j-1] + cost[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			tc[i][j] = cost[i][j] + min(tc[i-1][j-1], tc[i-1][j], tc[i][j-1])
		}
	}
	return tc[m-1][n-1]
}

func min(a, b, c int) int {
	minVal := a
	if b < a {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	return minVal
}

func main() {
	cost := [][]int{{1, 3, 4}, {4, 7, 5}, {1, 5, 3}}
	fmt.Println(MinCost(cost, 3, 3))
	fmt.Println(MinCostBU(cost, 3, 3))
}

/*
11
11
*/
