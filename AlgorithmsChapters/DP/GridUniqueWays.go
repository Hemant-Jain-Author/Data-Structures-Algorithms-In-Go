package main

import "fmt"

// Diagonal movement allowed.
func Unique3Ways(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// Initialize first column.
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0]
	}
	// Initialize first row.
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
func UniqueWays(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// Initialize first column.
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0]
	}
	// Initialize first row.
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(UniqueWays(3, 3))
	fmt.Println(Unique3Ways(3, 3))
}

/*
6
13
*/
