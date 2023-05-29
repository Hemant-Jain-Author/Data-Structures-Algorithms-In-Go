package main

import "fmt"

func MinStairCost(cost []int, n int) int {
	// base case
	if n == 1 {
		return cost[0]
	}
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-2], dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	a := []int{1, 5, 6, 3, 4, 7, 9, 1, 2, 11}
	n := len(a)
	fmt.Print(MinStairCost(a, n))
}

/*
18
*/
