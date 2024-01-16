package main

import "fmt"

func FindWays(n int, m int, V int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, V+1)
	}

	// Table entries for only one dice.
	for j := 1; j <= m && j <= V; j++ {
		dp[1][j] = 1
	}

	// i is number of dice, j is Value, k value of dice.
	for i := 2; i <= n; i++ {
		for j := 1; j <= V; j++ {
			for k := 1; k <= j && k <= m; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	return dp[n][V]
}

func main() {
	fmt.Println(FindWays(3, 6, 6))
}

/*
10
*/